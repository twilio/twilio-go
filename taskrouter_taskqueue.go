package twilio

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// TaskRouterTaskQueue allow you to categorize Tasks and describe which Workers are eligible to handle those Tasks.
// refer: https://www.twilio.com/docs/taskrouter/api/task-queue
type TaskRouterTaskQueue struct {
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

// TaskRouterTaskQueueParams taskRouterTaskQueue parameters.
type TaskRouterTaskQueueParams struct {
	FriendlyName           *string `form:",omitempty"`
	AssignmentActivitySID  *string `form:"AssignmentActivitySid,omitempty"`
	MaxReservedWorkers     *int    `form:",omitempty"`
	TargetWorkers          *string `form:",omitempty"`
	TaskOrder              *string `form:"TaskOrder,omitempty"`
	ReservationActivitySID *string `form:"ReservationActivitySid,omitempty"`
}

// TaskRouterTaskQueueList struct to parse response of taskRouterTaskqueue read.
type TaskRouterTaskQueueList struct {
	TaskRouterTaskQueues *[]TaskRouterTaskQueue `json:"task_queuess"`
	Meta                 *Meta                  `json:"meta,omitempty"`
}

// TaskRouterTaskQueueQueryParams query params to read taskRouterTaskqueues.
type TaskRouterTaskQueueQueryParams struct {
	FriendlyName             *string `form:"FriendlyName,omitempty"`
	EvaluateWorkerAttributes *string `form:"EvaluateWorkerAttributes,omitempty"`
	PageSize                 *int    `form:"PageSize,omitempty"`
}

// TaskRouterTaskQueueClient is the entrypoint for taskRouterTaskqueue CRUD.
type TaskRouterTaskQueueClient struct {
	ServiceURL string
	client     *Twilio
}

// NewTaskRouterTaskQueueClient constructs a new TaskRouterTaskQueue Client.
func NewTaskRouterTaskQueueClient(client *Twilio) *TaskRouterTaskQueueClient {
	c := new(TaskRouterTaskQueueClient)
	c.client = client
	c.ServiceURL = fmt.Sprintf("https://taskrouter.%s/v1/Workspaces", client.BaseURL)

	return c
}

// Create creates taskRouterTaskqueue with the given the config.
func (c *TaskRouterTaskQueueClient) Create(workspaceSID string, taskRouterTaskqueueparams *TaskRouterTaskQueueParams) (*TaskRouterTaskQueue, error) {
	url := fmt.Sprintf("%s/%s/%s", c.ServiceURL, workspaceSID, "TaskRouterTaskQueues")

	if len(*taskRouterTaskqueueparams.FriendlyName) == 0 {
		return nil, errors.New("friendlyname is required in taskQqueuesParams")
	}

	resp, err := c.client.Post(url, taskRouterTaskqueueparams)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	taskRouterTaskQueue := &TaskRouterTaskQueue{}

	if err := json.NewDecoder(resp.Body).Decode(taskRouterTaskQueue); err != nil {
		return nil, err
	}

	return taskRouterTaskQueue, nil
}

// Fetch fetches taskRouterTaskqueue for the given SID.
func (c *TaskRouterTaskQueueClient) Fetch(workspaceSID string, taskRouterTaskQueueSID string) (*TaskRouterTaskQueue, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", c.ServiceURL, workspaceSID, "TaskRouterTaskQueues", taskRouterTaskQueueSID)
	resp, err := c.client.Get(url, nil)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	taskRouterTaskqueue := &TaskRouterTaskQueue{}

	if err := json.NewDecoder(resp.Body).Decode(taskRouterTaskqueue); err != nil {
		return nil, err
	}

	return taskRouterTaskqueue, nil
}

// Read returns all existing taskRouterTaskqueues.
func (c *TaskRouterTaskQueueClient) Read(workspaceSID string, queryParams *TaskRouterTaskQueueQueryParams) (*TaskRouterTaskQueueList, error) {
	url := fmt.Sprintf("%s/%s/%s", c.ServiceURL, workspaceSID, "TaskRouterTaskQueues")

	resp, err := c.client.Get(url, queryParams)

	if err != nil {
		return nil, err
	}

	taskRouterTaskqueues := &TaskRouterTaskQueueList{}

	if decodeErr := json.NewDecoder(resp.Body).Decode(&taskRouterTaskqueues); decodeErr != nil {
		return nil, decodeErr
	}

	return taskRouterTaskqueues, nil
}

// Update updates taskRouterTaskqueue with given config.
func (c *TaskRouterTaskQueueClient) Update(workspaceSID string,
	taskRouterTaskQueueSID string, taskRouterTaskQueueParams *TaskRouterTaskQueueParams) (*TaskRouterTaskQueue, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", c.ServiceURL, workspaceSID, "TaskRouterTaskQueues", taskRouterTaskQueueSID)

	resp, err := c.client.Post(url, taskRouterTaskQueueParams)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	taskRouterTaskQueue := &TaskRouterTaskQueue{}

	if err := json.NewDecoder(resp.Body).Decode(taskRouterTaskQueue); err != nil {
		return nil, err
	}

	return taskRouterTaskQueue, nil
}

// Delete deletes taskRouterTaskQueue with the given SID.
func (c *TaskRouterTaskQueueClient) Delete(workspaceSID string, taskRouterTaskqueueSID string) error {
	url := fmt.Sprintf("%s/%s/%s/%s", c.ServiceURL, workspaceSID, "TaskRouterTaskQueues", taskRouterTaskqueueSID)
	resp, err := c.client.Delete(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
