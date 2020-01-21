package taskrouter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	twilio "github.com/twilio/twilio-go"
)

// NewTaskrouterClient creates and returns taskrouter client
func NewTaskrouterClient(accountSid string, authToken string, domain string) Client {
	credentials := twilio.Credentials{AccountSid: accountSid,
		AuthToken: authToken,
	}
	request := twilio.Request{Credentials: credentials,
		BaseURL: "https://taskrouter." + domain,
		Client:  &http.Client{},
	}

	return Client{Request: request}
}

// Client taskrouter client
type Client struct {
	Request twilio.Request
}

// WorkspaceList list of workspaces
type WorkspaceList struct {
	Workspaces []twilio.Workspace `json:"workspaces,omitempty"`
}

// CreateWorkflow create workflow with the given friendly name
func (c Client) CreateWorkflow(workflowParams twilio.WorkflowParams) (*twilio.Workflow, error) {
	url := "/v1/Workspaces/" + workflowParams.WorkspaceSid + "/Workflows"
	resp, err := c.Request.Post(url, workflowParams)

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

// CreateWorkspace create workspace with the given friendly name
func (c Client) CreateWorkspace(workspaceParams twilio.WorkspaceParams) (*twilio.Workspace, error) {
	url := "/v1/Workspaces"
	resp, err := c.Request.Post(url, workspaceParams)

	if err != nil {
		log.Printf("error while creating workspace: %s", err)
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

// CreateNewActivity create a new activity
func (c Client) CreateNewActivity(activityParams twilio.ActivityParams, workspaceSid string) (*twilio.Activity, error) {
	url := "/v1/Workspaces/" + workspaceSid + "/Activities"

	resp, err := c.Request.Post(url, activityParams)

	if err != nil {
		log.Printf("error decoding the output of activity create: %s", err)
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

// FetchWorkflow fetch
func (c Client) FetchWorkflow(workspaceSid string, workflowSid string) (*twilio.Workflow, error) {
	url := "/v1/Workspaces/" + workspaceSid + "/Workflows/" + workflowSid
	resp, err := c.Request.Get(url)

	if err != nil {
		log.Printf("error while getting workflow: %s", err)
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

// CreateTaskQueue create a new activity
func (c Client) CreateTaskQueue(taskQueuParams twilio.TaskQueueParams, workspaceSid string) (*twilio.TaskQueue, error) {
	url := "/v1/Workspaces/" + workspaceSid + "/TaskQueues"
	resp, err := c.Request.Post(url, taskQueuParams)

	if err != nil {
		log.Printf("error while creating taskqueue: %s", err)
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

// SetupTaskRouterWithDefaultConfig sets up taskrouter
func (c Client) SetupTaskRouterWithDefaultConfig() (*twilio.Workspace, error) {
	workspaceParams := twilio.WorkspaceParams{
		FriendlyName: twilio.WorkspaceDefaultFriendlyName,
	}
	ws, err := c.SetupWorkspace(workspaceParams)

	if err != nil {
		return nil, err
	}

	taskQueueParams := twilio.TaskQueueParams{
		FriendlyName: twilio.TaskQueueDefaultFriendlyName,
	}

	tq, err := c.CreateTaskQueue(taskQueueParams, ws.Sid)

	if err != nil {
		return ws, err
	}

	conf := `{"task_routing":{"default_filter":{"task_queue_sid":"` + tq.Sid + `"}}}`

	workflowParams := twilio.WorkflowParams{
		FriendlyName:  twilio.WorkflowDefaultFriendlyName,
		WorkspaceSid:  ws.Sid,
		Configuration: conf,
	}

	_, err = c.CreateWorkflow(workflowParams)

	if err != nil {
		return ws, err
	}

	activityParams := twilio.ActivityParams{
		FriendlyName: twilio.ActivityDefaultFriendlyName,
	}

	_, err = c.CreateNewActivity(activityParams, ws.Sid)

	return ws, err
}

// SetupWorkspace sets up taskrouter
func (c Client) SetupWorkspace(workspaceParams twilio.WorkspaceParams) (*twilio.Workspace, error) {
	ws, err := c.FindWorkspace(workspaceParams.FriendlyName)

	if err != nil {
		log.Panic("error while searching for existing workspace", err)
	}

	if ws != nil {
		fmt.Println("workspace already exists, skipping workspace creation")
		return ws, nil
	}

	return c.CreateWorkspace(workspaceParams)
}

// FindWorkspace find a workspace by a friendly name.
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

	resp, err := c.Request.Get(uri)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	wslist, err := parseWorkspaceResponse(body)

	if err != nil {
		return nil, err
	}

	return wslist, err
}

func parseWorkspaceResponse(body []byte) (*WorkspaceList, error) {
	var wslist = new(WorkspaceList)
	err := json.Unmarshal(body, &wslist)

	return wslist, err
}

// DeleteWorkspaces delete
func (c Client) DeleteWorkspaces(sid string) error {
	url := "/v1/Workspaces/" + sid
	resp, err := c.Request.Delete(url)

	if err != nil {
		log.Printf("Error deleting workspace %s", err)
	}

	defer resp.Body.Close()

	return err
}
