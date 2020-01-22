// Package taskrouter - package for CRUD operations on taskrouter (workspace, workflow, taskqueues, activities)
package taskrouter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	twilio "github.com/twilio/twilio-go"
)

// NewTaskrouterClient creates and returns taskrouter client (constructor).
func NewTaskrouterClient(accountSid string, authToken string, domain string) Client {
	credentials := twilio.Credentials{AccountSID: accountSid,
		AuthToken: authToken,
	}
	request := twilio.Client{Credentials: credentials,
		BaseURL:    "https://taskrouter." + domain,
		HttpClient: &http.Client{},
	}

	return Client{Client: request}
}

// Client taskrouter client struct.
type Client struct {
	Client twilio.Client
}

// WorkspaceList struct for holding list of workspaces.
type WorkspaceList struct {
	Workspaces []twilio.Workspace `json:"workspaces"`
}

// CreateWorkflow creates workflow with the given config.
func (c Client) CreateWorkflow(workflowParams twilio.WorkflowParams) (*twilio.Workflow, error) {
	url := "/v1/Workspaces/" + workflowParams.WorkspaceSid + "/Workflows"
	resp, err := c.Client.Post(url, workflowParams)

	if err != nil {
		log.Printf("error creating workflow: %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	wf := &twilio.Workflow{}
	err = json.NewDecoder(resp.Body).Decode(wf)

	if err != nil {
		log.Printf("error decoding the output of workflow create: %s", err)
		return nil, err
	}

	return wf, nil
}

// CreateWorkspace creates workspace with the given the config.
func (c Client) CreateWorkspace(workspaceParams twilio.WorkspaceParams) (*twilio.Workspace, error) {
	url := "/v1/Workspaces"
	resp, err := c.Client.Post(url, workspaceParams)

	if err != nil {
		log.Printf("error creating workspace: %s", err)
	}

	defer resp.Body.Close()

	ws := &twilio.Workspace{}
	err = json.NewDecoder(resp.Body).Decode(ws)

	if err != nil {
		log.Printf("error decoding the output of workspace create: %s", err)
		return nil, err
	}

	return ws, nil
}

// CreateNewActivity creates a new activity with the given config.
func (c Client) CreateNewActivity(activityParams twilio.ActivityParams, workspaceSid string) (*twilio.Activity, error) {
	url := "/v1/Workspaces/" + workspaceSid + "/Activities"

	resp, err := c.Client.Post(url, activityParams)

	if err != nil {
		log.Printf("error creating a new activity: %s", err)
	}

	defer resp.Body.Close()

	activity := &twilio.Activity{}
	err = json.NewDecoder(resp.Body).Decode(activity)

	if err != nil {
		log.Printf("error decoding the output of acitivity create: %s", err)
		return nil, err
	}

	return activity, nil
}

// FetchWorkflow fetches workflow with given workspace sid and workflow sid.
func (c Client) FetchWorkflow(workspaceSid string, workflowSid string) (*twilio.Workflow, error) {
	url := "/v1/Workspaces/" + workspaceSid + "/Workflows/" + workflowSid
	resp, err := c.Client.Get(url)

	if err != nil {
		log.Printf("error getting existing workflow: %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	workflow := &twilio.Workflow{}
	err = json.NewDecoder(resp.Body).Decode(workflow)

	if err != nil {
		log.Printf("error decoding the output of workflow fetch: %s", err)
		return nil, err
	}

	return workflow, nil
}

// CreateTaskQueue creatse a new task-queue with the given config.
func (c Client) CreateTaskQueue(taskQueuParams twilio.TaskQueueParams, workspaceSid string) (*twilio.TaskQueue, error) {
	url := "/v1/Workspaces/" + workspaceSid + "/TaskQueues"
	resp, err := c.Client.Post(url, taskQueuParams)

	if err != nil {
		log.Printf("error creating taskqueue: %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	taskQueue := &twilio.TaskQueue{}
	err = json.NewDecoder(resp.Body).Decode(taskQueue)

	if err != nil {
		log.Printf("error decoding the output of task queue create: %s", err)
		return nil, err
	}

	return taskQueue, nil
}

// SetupTaskRouterWithDefaultConfig sets up taskrouter(Flex setup).
// 1.Checks if a workspace exists with the default friendly name, if not then creates one.
// 2.Creates a task-queue with default friendly name for the workspace sid.
// 3.Creates workflow with default friendly name using the created taskqueue sid and workspace sid.
// 4.Creates a new activity with default friendly name using the sid of workspace.
func (c Client) SetupTaskRouterWithDefaultConfig() (*twilio.Workspace, *twilio.TaskQueue, *twilio.Workflow,
	*twilio.Activity, error) {
	workspaceParams := twilio.WorkspaceParams{
		FriendlyName: twilio.WorkspaceDefaultFriendlyName,
	}
	ws, err := c.SetupWorkspace(workspaceParams)

	if err != nil {
		return nil, nil, nil, nil, err
	}

	taskQueueParams := twilio.TaskQueueParams{
		FriendlyName: twilio.TaskQueueDefaultFriendlyName,
	}

	tq, err := c.CreateTaskQueue(taskQueueParams, ws.Sid)

	if err != nil {
		return ws, nil, nil, nil, err
	}

	conf := `{"task_routing":{"default_filter":{"task_queue_sid":"` + tq.Sid + `"}}}`

	workflowParams := twilio.WorkflowParams{
		FriendlyName:  twilio.WorkflowDefaultFriendlyName,
		WorkspaceSid:  ws.Sid,
		Configuration: conf,
	}

	wf, err := c.CreateWorkflow(workflowParams)

	if err != nil {
		return ws, tq, nil, nil, err
	}

	activityParams := twilio.ActivityParams{
		FriendlyName: twilio.ActivityDefaultFriendlyName,
	}

	ac, err := c.CreateNewActivity(activityParams, ws.Sid)

	return ws, tq, wf, ac, err
}

// SetupWorkspace checks if workspace exists, returns if found, else creates one.
func (c Client) SetupWorkspace(workspaceParams twilio.WorkspaceParams) (*twilio.Workspace, error) {
	ws, err := c.FindWorkspace(workspaceParams.FriendlyName)

	if err != nil {
		return nil, err
	}

	if ws != nil {
		fmt.Println("workspace already exists, skipping workspace creation")
		return ws, nil
	}

	return c.CreateWorkspace(workspaceParams)
}

// FindWorkspace finds a workspace by a friendly name.
func (c Client) FindWorkspace(friendlyName string) (*twilio.Workspace, error) {
	wslist, err := c.RetreiveAllWorkspaces()

	if err != nil {
		log.Printf("error retrieving existing workspaaces: %s", err)
		return nil, err
	}

	for _, element := range wslist.Workspaces {
		if element.FriendlyName == friendlyName {
			return &element, nil
		}
	}

	return nil, nil
}

// RetreiveAllWorkspaces retrieves all workspaces (upto 1000) for a given account/sid.
func (c Client) RetreiveAllWorkspaces() (*WorkspaceList, error) {
	uri := "/v1/Workspaces?pagesize=1000"

	resp, err := c.Client.Get(uri)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	wslist, err := workspaceListFromResponseBody(body)

	if err != nil {
		return nil, err
	}

	return wslist, err
}

func workspaceListFromResponseBody(body []byte) (*WorkspaceList, error) {
	var wslist = new(WorkspaceList)
	err := json.Unmarshal(body, &wslist)

	return wslist, err
}

// DeleteWorkspaces deletes a workspace for a given sid.
func (c Client) DeleteWorkspaces(sid string) error {
	url := "/v1/Workspaces/" + sid
	resp, err := c.Client.Delete(url)

	if err != nil {
		log.Printf("Error deleting workspace %s", err)
	}

	defer resp.Body.Close()

	return err
}
