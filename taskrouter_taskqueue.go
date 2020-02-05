package twilio

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	twilio "github.com/twilio/twilio-go/internal"
)

// TaskQueue allow you to categorize Tasks and describe which Workers are eligible to handle those Tasks.
// refer: https://www.twilio.com/docs/taskrouter/api/task-queue
type TaskQueue struct {
	AccountSid              *string            `json:"account_sid"`
	AssignmentActivitySid   *string            `json:"assignment_activity_sid"`
	AssignmentActivityName  *string            `json:"assignment_activity_name"`
	DateCreated             *time.Time         `json:"date_created"`
	DateUpdated             *time.Time         `json:"date_updated"`
	FriendlyName            *string            `json:"friendly_name"`
	MaxReservedWorkers      *int               `json:"max_reserved_workers"`
	ReservationActivitySid  *string            `json:"reservation_activity_sid"`
	ReservationActivityName *string            `json:"reservation_activity_name"`
	Sid                     *string            `json:"sid"`
	TargetWorkers           *string            `json:"target_workers"`
	TaskOrder               *string            `json:"task_order"`
	URI                     *string            `json:"url"`
	WorkspaceSid            *string            `json:"workspace_sid"`
	Links                   map[string]*string `json:"links"`
}

// TaskQueueParams taskQueue parameters.
type TaskQueueParams struct {
	FriendlyName           *string `form:",omitempty"`
	AssignmentActivitySID  *string `form:"AssignmentActivitySid,omitempty"`
	MaxReservedWorkers     *int    `form:",omitempty"`
	TargetWorkers          *string `form:",omitempty"`
	TaskOrder              *string `form:"TaskOrder,omitempty"`
	ReservationActivitySID *string `form:"ReservationActivitySid,omitempty"`
}

// TaskQueueList struct to parse response of taskqueue read.
type TaskQueueList struct {
	TaskQueues *[]TaskQueue `json:"task_queues"`
	Meta       *Meta        `json:"meta,omitempty"`
}

// TaskQueueQueryParams query params to read taskqueues.
type TaskQueueQueryParams struct {
	FriendlyName             *string `form:"FriendlyName,omitempty"`
	EvaluateWorkerAttributes *string `form:"EvaluateWorkerAttributes,omitempty"`
	PageSize                 *int    `form:"PageSize,omitempty"`
}

// TaskQueueClient is the entrypoint for taskqueue CRUD.
type TaskQueueClient struct {
	ServiceURL string
	Client     *twilio.Client
}

// NewTaskQueueClient constructs a new TaskQueue Client.
func NewTaskQueueClient(twilioClient *twilio.Client) *TaskQueueClient {
	c := new(TaskQueueClient)
	c.Client = twilioClient
	c.ServiceURL = fmt.Sprintf("https://taskrouter.%s/v1/Workspaces", twilioClient.BaseURL)

	return c
}

// Create creates taskqueue with the given the config.
func (ws *TaskQueueClient) Create(workspaceSID string, taskqueueparams *TaskQueueParams) (*TaskQueue, error) {
	url := fmt.Sprintf("%s/%s/%s", ws.ServiceURL, workspaceSID, "TaskQueues")

	if len(*taskqueueparams.FriendlyName) == 0 {
		return nil, errors.New("friendlyname is required in taskQqueueParams")
	}

	resp, err := ws.Client.Post(url, taskqueueparams)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	taskQueue := &TaskQueue{}

	if err := json.NewDecoder(resp.Body).Decode(taskQueue); err != nil {
		return nil, err
	}

	return taskQueue, nil
}

// Fetch fetches taskqueue for the given SID.
func (ws *TaskQueueClient) Fetch(workspaceSID string, taskQueueSID string) (*TaskQueue, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", ws.ServiceURL, workspaceSID, "TaskQueues", taskQueueSID)
	resp, err := ws.Client.Get(url, nil)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	taskqueue := &TaskQueue{}

	if err := json.NewDecoder(resp.Body).Decode(taskqueue); err != nil {
		return nil, err
	}

	return taskqueue, nil
}

// Read returns all existing taskqueues.
func (ws *TaskQueueClient) Read(workspaceSID string, queryParams *TaskQueueQueryParams) (*TaskQueueList, error) {
	url := fmt.Sprintf("%s/%s/%s", ws.ServiceURL, workspaceSID, "TaskQueues")

	resp, err := ws.Client.Get(url, queryParams)

	if err != nil {
		return nil, err
	}

	taskqueues := &TaskQueueList{}

	if decodeErr := json.NewDecoder(resp.Body).Decode(&taskqueues); decodeErr != nil {
		return nil, decodeErr
	}

	return taskqueues, nil
}

// Update updates taskqueue with given config.
func (ws *TaskQueueClient) Update(workspaceSID string,
	taskQueueSID string, taskQueueParams *TaskQueueParams) (*TaskQueue, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", ws.ServiceURL, workspaceSID, "TaskQueues", taskQueueSID)

	resp, err := ws.Client.Post(url, taskQueueParams)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	taskQueue := &TaskQueue{}

	if err := json.NewDecoder(resp.Body).Decode(taskQueue); err != nil {
		return nil, err
	}

	return taskQueue, nil
}

// Delete deletes taskQueue with the given SID.
func (ws *TaskQueueClient) Delete(workspaceSID string, taskqueueSID string) error {
	url := fmt.Sprintf("%s/%s/%s/%s", ws.ServiceURL, workspaceSID, "TaskQueues", taskqueueSID)
	resp, err := ws.Client.Delete(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
