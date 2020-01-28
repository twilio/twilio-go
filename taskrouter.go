package twilio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	twilio "github.com/twilio/twilio-go/internal"
)

//Workspace refer: https://www.twilio.com/docs/taskrouter/api/workspace
type Workspace struct {
	Sid                  string            `json:"sid"`
	AccountSid           string            `json:"account_sid"`
	DateCreated          time.Time         `json:"date_created"`
	DateUpdated          time.Time         `json:"date_updated"`
	DefaultActivityName  string            `json:"default_activity_name"`
	DefaultActivitySid   string            `json:"default_activity_sid"`
	TimeoutActivityName  string            `json:"timeout_activity_name"`
	TimeoutActivitySid   string            `json:"timeout_activity_sid"`
	URL                  string            `json:"url"`
	Links                map[string]string `json:"links"`
	FriendlyName         string            `json:"friendly_name"`
	EventCallbackURL     string            `json:"event_callback_url"`
	EventsFilter         string            `json:"events_filter"`
	MultitaskEnabled     bool              `json:"multi_task_enabled"`
	Template             string            `json:"template"`
	PrioritizeQueueOrder string            `json:"prioritize_queue_order"`
}

//WorkspaceParams refer: https://www.twilio.com/docs/taskrouter/api/workspace
type WorkspaceParams struct {
	FriendlyName         string `url:"FriendlyName,omitempty"`
	EventCallbackURL     string `url:"EventCallbackUrl,omitempty"`
	EventsFilter         string `url:"EventsFilter,omitempty"`
	MultitaskEnabled     bool   `url:"MultitaskEnabled,omitempty"`
	Template             string `url:"Template,omitempty"`
	PrioritizeQueueOrder string `url:"PrioritizeQueueOrder,omitempty"`
}

// constants
const (
	//WorkspaceDefaultFriendlyName = "Flex Task Assignment"
	WorkspaceDefaultFriendlyName = "Flex Task Assignment"
	WorkflowDefaultFriendlyName  = "Assign to Anyone"
	TaskQueueDefaultFriendlyName = "Everyone"
	ActivityDefaultFriendlyName  = "Break"
)

// Workflow workflow refer: https://www.twilio.com/docs/taskrouter/api/workflow
type Workflow struct {
	AssignmentCallbackURL         string            `json:"assignment_callback_url"`
	Configuration                 string            `json:"configuration"`
	AccountSid                    string            `json:"account_sid"`
	DateCreated                   time.Time         `json:"date_created"`
	DateUpdated                   time.Time         `json:"date_updated"`
	DocumentContentType           string            `json:"document_content_type"`
	FallbackAssignmentCallbackURL string            `json:"fallback_assignment_callback_url"`
	FriendlyName                  string            `json:"friendly_name"`
	Sid                           string            `json:"sid"`
	TaskReservationTimeout        int               `json:"task_reservation_timeout"`
	WorkspaceSid                  string            `json:"workspace_sid"`
	URL                           string            `json:"url"`
	Links                         map[string]string `json:"links"`
}

// WorkflowParams workflow parameters
type WorkflowParams struct {
	WorkspaceSid                  string `url:"WorkspaceSid,omitempty"`
	FriendlyName                  string `url:"FriendlyName,omitempty"`
	Configuration                 string `url:"Configuration,omitempty"`
	AssignmentCallbackURL         string `url:"AssignmentCallbackURL,omitempty"`
	FallbackAssignmentCallbackURL string `url:"FallbackAssignmentCallbackURL,omitempty"`
	TaskReservationTimeout        int    `url:"TaskReservationTimeout,omitempty"`
}

// TaskQueue taskqueue refer: https://www.twilio.com/docs/taskrouter/api/task-queue
type TaskQueue struct {
	AccountSid              string            `json:"account_sid"`
	AssignmentActivitySid   string            `json:"assignment_activity_sid"`
	AssignmentActivityName  string            `json:"assignment_activity_name"`
	DateCreated             time.Time         `json:"date_created"`
	DateUpdated             time.Time         `json:"date_updated"`
	FriendlyName            string            `json:"friendly_name"`
	MaxReservedWorkers      int               `json:"max_reserved_workers"`
	ReservationActivitySid  string            `json:"reservation_activity_sid"`
	ReservationActivityName string            `json:"reservation_activity_name"`
	Sid                     string            `json:"sid"`
	TargetWorkers           string            `json:"target_workers"`
	TaskOrder               string            `json:"task_order"`
	URI                     string            `json:"url"`
	WorkspaceSid            string            `json:"workspace_sid"`
	Links                   map[string]string `json:"links"`
}

// TaskQueueParams taskQueue parameters refer: https://www.twilio.com/docs/taskrouter/api/task-queue
type TaskQueueParams struct {
	FriendlyName           string `url:"FriendlyName,omitempty"`
	AssignmentActivitySid  string `url:"AssignmentActivitySid,omitempty"`
	MaxReservedWorkers     int    `url:"MaxReservedWorkers,omitempty"`
	TargetWorkers          string `url:"TargetWorkers,omitempty"`
	TaskOrder              string `url:"TaskOrder,omitempty"`
	ReservationActivitySid string `url:"ReservationActivitySid,omitempty"`
}

// Activity activity refer: https://www.twilio.com/docs/taskrouter/api/activity
type Activity struct {
	AccountSid   string    `json:"account_sid"`
	Available    string    `json:"available"`
	DateCreated  time.Time `json:"date_created"`
	DateUpdated  time.Time `json:"date_updated"`
	FriendlyName string    `json:"friendly_name"`
	Sid          string    `json:"sid"`
	WorkspaceSid string    `json:"workspace_sid"`
	URI          string    `json:"url"`
}

// ActivityParams activity refer: https://www.twilio.com/docs/taskrouter/api/activity
type ActivityParams struct {
	Available    string `url:"Available,omitempty"`
	FriendlyName string `url:"FriendlyName,omitempty"`
	WorkspaceSid string `url:"WorkspaceSid,omitempty"`
}

// TaskRouter is the entrypoint for the TaskRouter API.
type TaskRouter struct {
	serviceURL string
	client     *twilio.Client
}

// WorkspaceList struct for holding list of workspaces.
type WorkspaceList struct {
	Workspaces []Workspace `json:"workspaces"`
}

// Initialize constructs a new TaskRouter client.
func (tr *TaskRouter) Initialize(client *twilio.Client) {
	tr.client = client
	tr.serviceURL = fmt.Sprintf("https://taskrouter.%s/v1", tr.client.BaseURL)
}

// CreateWorkflow creates workflow with the given config.
func (tr TaskRouter) CreateWorkflow(workflowParams WorkflowParams) (*Workflow, error) {
	url := tr.serviceURL + "/Workspaces/" + workflowParams.WorkspaceSid + "/Workflows"
	resp, err := tr.client.Post(url, workflowParams)

	if err != nil {
		log.Printf("error creating workflow: %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	wf := &Workflow{}
	err = json.NewDecoder(resp.Body).Decode(wf)

	if err != nil {
		log.Printf("error decoding the output of workflow create: %s", err)
		return nil, err
	}

	return wf, nil
}

// CreateWorkspace creates workspace with the given the config.
func (tr TaskRouter) CreateWorkspace(workspaceParams WorkspaceParams) (*Workspace, error) {
	url := tr.serviceURL + "/Workspaces"
	resp, err := tr.client.Post(url, workspaceParams)

	if err != nil {
		log.Printf("error creating workspace: %s", err)
	}

	defer resp.Body.Close()

	ws := &Workspace{}
	err = json.NewDecoder(resp.Body).Decode(ws)

	if err != nil {
		log.Printf("error decoding the output of workspace create: %s", err)
		return nil, err
	}

	return ws, nil
}

// CreateNewActivity creates a new activity with the given config.
func (tr TaskRouter) CreateNewActivity(activityParams ActivityParams, workspaceSid string) (*Activity, error) {
	url := tr.serviceURL + "/Workspaces/" + workspaceSid + "/Activities"

	resp, err := tr.client.Post(url, activityParams)

	if err != nil {
		log.Printf("error creating a new activity: %s", err)
	}

	defer resp.Body.Close()

	activity := &Activity{}
	err = json.NewDecoder(resp.Body).Decode(activity)

	if err != nil {
		log.Printf("error decoding the output of acitivity create: %s", err)
		return nil, err
	}

	return activity, nil
}

// FetchWorkflow fetches workflow with given workspace sid and workflow sid.
func (tr TaskRouter) FetchWorkflow(workspaceSid string, workflowSid string) (*Workflow, error) {
	url := tr.serviceURL + "/Workspaces/" + workspaceSid + "/Workflows/" + workflowSid
	resp, err := tr.client.Get(url, nil)

	if err != nil {
		log.Printf("error getting existing workflow: %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	workflow := &Workflow{}
	err = json.NewDecoder(resp.Body).Decode(workflow)

	if err != nil {
		log.Printf("error decoding the output of workflow fetch: %s", err)
		return nil, err
	}

	return workflow, nil
}

// CreateTaskQueue creatse a new task-queue with the given config.
func (tr TaskRouter) CreateTaskQueue(taskQueuParams TaskQueueParams, workspaceSid string) (*TaskQueue, error) {
	url := tr.serviceURL + "/Workspaces/" + workspaceSid + "/TaskQueues"
	resp, err := tr.client.Post(url, taskQueuParams)

	if err != nil {
		log.Printf("error creating taskqueue: %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	taskQueue := &TaskQueue{}
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
func (tr TaskRouter) SetupTaskRouterWithDefaultConfig() (*Workspace, *TaskQueue, *Workflow,
	*Activity, error) {
	workspaceParams := WorkspaceParams{
		FriendlyName: WorkspaceDefaultFriendlyName,
	}
	ws, err := tr.SetupWorkspace(workspaceParams)

	if err != nil {
		return nil, nil, nil, nil, err
	}

	taskQueueParams := TaskQueueParams{
		FriendlyName: TaskQueueDefaultFriendlyName,
	}

	tq, err := tr.CreateTaskQueue(taskQueueParams, ws.Sid)

	if err != nil {
		return ws, nil, nil, nil, err
	}

	conf := `{"task_routing":{"default_filter":{"task_queue_sid":"` + tq.Sid + `"}}}`

	workflowParams := WorkflowParams{
		FriendlyName:  WorkflowDefaultFriendlyName,
		WorkspaceSid:  ws.Sid,
		Configuration: conf,
	}

	wf, err := tr.CreateWorkflow(workflowParams)

	if err != nil {
		return ws, tq, nil, nil, err
	}

	activityParams := ActivityParams{
		FriendlyName: ActivityDefaultFriendlyName,
	}

	ac, err := tr.CreateNewActivity(activityParams, ws.Sid)

	return ws, tq, wf, ac, err
}

// SetupWorkspace checks if workspace exists, returns if found, else creates one.
func (tr TaskRouter) SetupWorkspace(workspaceParams WorkspaceParams) (*Workspace, error) {
	ws, err := tr.FindWorkspace(workspaceParams.FriendlyName)

	if err != nil {
		return nil, err
	}

	if ws != nil {
		fmt.Println("workspace already exists, skipping workspace creation")
		return ws, nil
	}

	return tr.CreateWorkspace(workspaceParams)
}

// FindWorkspace finds a workspace by a friendly name.
func (tr TaskRouter) FindWorkspace(friendlyName string) (*Workspace, error) {
	wslist, err := tr.RetrieveAllWorkspaces()

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

// RetrieveAllWorkspaces retrieves all workspaces (upto 1000) for a given account/sid.
func (tr TaskRouter) RetrieveAllWorkspaces() (*WorkspaceList, error) {
	uri := tr.serviceURL + "/Workspaces?pagesize=1000"

	resp, err := tr.client.Get(uri, nil)

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
func (tr TaskRouter) DeleteWorkspaces(sid string) error {
	url := tr.serviceURL + "/Workspaces/" + sid
	resp, err := tr.client.Delete(url)

	if err != nil {
		log.Printf("Error deleting workspace %s", err)
	}

	defer resp.Body.Close()

	return err
}
