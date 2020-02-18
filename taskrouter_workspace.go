package twilio

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// A TaskRouterWorkspace is a container for your Tasks, Workers, TaskQueues, Workflows, and Activities.
// refer: https://www.twilio.com/docs/taskrouter/api/workspace.
type TaskRouterWorkspace struct {
	SID                  *string            `json:"sid"`
	AccountSID           *string            `json:"account_sid"`
	DateCreated          *time.Time         `json:"date_created"`
	DateUpdated          *time.Time         `json:"date_updated"`
	DefaultActivityName  *string            `json:"default_activity_name"`
	DefaultActivitySID   *string            `json:"default_activity_sid"`
	TimeoutActivityName  *string            `json:"timeout_activity_name"`
	TimeoutActivitySID   *string            `json:"timeout_activity_sid"`
	URL                  *string            `json:"url"`
	Links                map[string]*string `json:"links"`
	FriendlyName         *string            `json:"friendly_name"`
	EventCallbackURL     *string            `json:"event_callback_url"`
	EventsFilter         *string            `json:"events_filter"`
	MultitaskEnabled     *bool              `json:"multi_task_enabled"`
	Template             *string            `json:"template"`
	PrioritizeQueueOrder *string            `json:"prioritize_queue_order"`
}

// TaskRouterWorkspaceParams TaskRouterWorkspace params for CRUD.
type TaskRouterWorkspaceParams struct {
	FriendlyName         *string `form:"FriendlyName,omitempty"`
	EventCallbackURL     *string `form:"EventCallbackUrl,omitempty"`
	EventsFilter         *string `form:"EventsFilter,omitempty"`
	MultitaskEnabled     *bool   `form:"MultitaskEnabled,omitempty"`
	Template             *string `form:"Template,omitempty"`
	PrioritizeQueueOrder *string `form:"PrioritizeQueueOrder,omitempty"`
}

// TaskRouterWorkspaceList struct to parse response of TaskRouterWorkspace read.
type TaskRouterWorkspaceList struct {
	Workspaces []*TaskRouterWorkspace `json:"workspaces"`
	Meta       *Meta                  `json:"meta,omitempty"`
}

// TaskRouterWorkspaceQueryParams query params to read TaskRouterWorkspaces.
type TaskRouterWorkspaceQueryParams struct {
	FriendlyName *string `form:"FriendlyName,omitempty"`
	PageSize     *int    `form:"PageSize,omitempty"`
}

// TaskRouterWorkspaceClient is the entrypoint for the TaskRouterWorkspace CRUD.
type TaskRouterWorkspaceClient struct {
	baseURL string
	client  *Twilio
}

// NewTaskRouterWorkspaceClient constructs a new TaskRouterWorkspace Client.
func NewTaskRouterWorkspaceClient(client *Twilio) *TaskRouterWorkspaceClient {
	c := new(TaskRouterWorkspaceClient)
	c.client = client
	c.baseURL = c.url(fmt.Sprintf("https://taskrouter.%s/v1", client.BaseURL))

	return c
}

// Create creates TaskRouterWorkspace with the given the config.
func (c *TaskRouterWorkspaceClient) Create(params *TaskRouterWorkspaceParams) (*TaskRouterWorkspace, error) {
	if len(*params.FriendlyName) == 0 {
		return nil, errors.New("friendlyname is required in params")
	}

	resp, err := c.client.Post(c.url("/Workspaces"), params)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	TaskRouterWorkspace := &TaskRouterWorkspace{}

	if err = json.NewDecoder(resp.Body).Decode(TaskRouterWorkspace); err != nil {
		return nil, err
	}

	return TaskRouterWorkspace, nil
}

// Fetch fetches TaskRouterWorkspace for the given TaskRouterWorkspace SID.
func (c *TaskRouterWorkspaceClient) Fetch(sid string) (*TaskRouterWorkspace, error) {
	resp, err := c.client.Get(c.url("/Workspaces/"+sid), nil)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	TaskRouterWorkspace := &TaskRouterWorkspace{}

	if err = json.NewDecoder(resp.Body).Decode(TaskRouterWorkspace); err != nil {
		return nil, err
	}

	return TaskRouterWorkspace, nil
}

// Read returns all existing TaskRouterWorkspaces.
func (c *TaskRouterWorkspaceClient) Read(params *TaskRouterWorkspaceQueryParams) (*TaskRouterWorkspaceList, error) {
	resp, err := c.client.Get(c.url("/Workspaces"), params)
	if err != nil {
		return nil, err
	}

	TaskRouterWorkspaces := &TaskRouterWorkspaceList{}

	if err = json.NewDecoder(resp.Body).Decode(&TaskRouterWorkspaces); err != nil {
		return nil, err
	}

	return TaskRouterWorkspaces, nil
}

// Update updates TaskRouterWorkspace with given config.
func (c *TaskRouterWorkspaceClient) Update(
	sid string,
	params *TaskRouterWorkspaceParams,
) (*TaskRouterWorkspace, error) {
	url := c.url("/Workspaces/" + sid)

	resp, err := c.client.Post(url, params)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	TaskRouterWorkspace := &TaskRouterWorkspace{}

	if err = json.NewDecoder(resp.Body).Decode(TaskRouterWorkspace); err != nil {
		return nil, err
	}

	return TaskRouterWorkspace, nil
}

// Delete deletes TaskRouterWorkspace for given SID.
func (c *TaskRouterWorkspaceClient) Delete(sid string) error {
	resp, err := c.client.Delete(c.url("/Workspaces/" + sid))

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}

func (c *TaskRouterWorkspaceClient) url(path string) string {
	if c.client.defaultbaseURL != nil {
		return *c.client.defaultbaseURL + path
	}

	return c.baseURL + path
}
