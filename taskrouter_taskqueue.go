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
	AccountSID              *string            `json:"account_sid"`
	AssignmentActivitySID   *string            `json:"assignment_activity_sid"`
	AssignmentActivityName  *string            `json:"assignment_activity_name"`
	DateCreated             *time.Time         `json:"date_created"`
	DateUpdated             *time.Time         `json:"date_updated"`
	FriendlyName            *string            `json:"friendly_name"`
	MaxReservedWorkers      *int               `json:"max_reserved_workers"`
	ReservationActivitySID  *string            `json:"reservation_activity_sid"`
	ReservationActivityName *string            `json:"reservation_activity_name"`
	SID                     *string            `json:"sid"`
	TargetWorkers           *string            `json:"target_workers"`
	TaskOrder               *string            `json:"task_order"`
	URL                     *string            `json:"url"`
	WorkspaceSID            *string            `json:"workspace_sid"`
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
	TaskQueues []*TaskRouterTaskQueue `json:"task_queues"`
	Meta       *Meta                  `json:"meta,omitempty"`
}

// TaskRouterTaskQueueQueryParams query params to read TaskQueues.
type TaskRouterTaskQueueQueryParams struct {
	FriendlyName             *string `form:"FriendlyName,omitempty"`
	EvaluateWorkerAttributes *string `form:"EvaluateWorkerAttributes,omitempty"`
	PageSize                 *int    `form:"PageSize,omitempty"`
}

// taskRouterTaskQueueClient is the entrypoint for taskRouterTaskqueue CRUD.
type taskRouterTaskQueueClient struct {
	baseURL string
	client  *Twilio
}

// newTaskRouterTaskQueueClient constructs a new TaskRouterTaskQueue Client.
func newTaskRouterTaskQueueClient(client *Twilio) *taskRouterTaskQueueClient {
	c := new(taskRouterTaskQueueClient)
	c.client = client
	c.baseURL = c.url(fmt.Sprintf("https://taskrouter.%s/v1", client.BaseURL))

	return c
}

// Create creates taskRouterTaskqueue with the given the config.
func (c *taskRouterTaskQueueClient) Create(
	workspaceSID string,
	params *TaskRouterTaskQueueParams,
) (*TaskRouterTaskQueue, error) {
	url := c.url(fmt.Sprintf("/Workspaces/%s/TaskQueues", workspaceSID))

	if len(*params.FriendlyName) == 0 {
		return nil, errors.New("friendlyname is required in taskQqueuesParams")
	}

	resp, err := c.client.Post(url, params)

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
func (c *taskRouterTaskQueueClient) Fetch(workspaceSID string, sid string) (*TaskRouterTaskQueue, error) {
	url := c.url(fmt.Sprintf("/Workspaces/%s/TaskQueues/%s", workspaceSID, sid))
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

// Read returns all existing TaskQueues.
func (c *taskRouterTaskQueueClient) Read(
	workspaceSID string,
	params *TaskRouterTaskQueueQueryParams,
) (*TaskRouterTaskQueueList, error) {
	url := c.url(fmt.Sprintf("/Workspaces/%s/TaskQueues", workspaceSID))

	resp, err := c.client.Get(url, params)

	if err != nil {
		return nil, err
	}

	q := &TaskRouterTaskQueueList{}
	if err := json.NewDecoder(resp.Body).Decode(q); err != nil {
		return nil, err
	}

	return q, nil
}

// Update updates taskRouterTaskqueue with given config.
func (c *taskRouterTaskQueueClient) Update(
	workspaceSID string,
	sid string,
	params *TaskRouterTaskQueueParams,
) (*TaskRouterTaskQueue, error) {
	url := c.url(fmt.Sprintf("/Workspaces/%s/TaskQueues/%s", workspaceSID, sid))
	resp, err := c.client.Post(url, params)

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
func (c *taskRouterTaskQueueClient) Delete(workspaceSID string, sid string) error {
	url := c.url(fmt.Sprintf("/Workspaces/%s/TaskQueues/%s", workspaceSID, sid))
	resp, err := c.client.Delete(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}

func (c *taskRouterTaskQueueClient) url(path string) string {
	if c.client.defaultbaseURL != nil {
		return *c.client.defaultbaseURL + path
	}

	return c.baseURL + path
}
