package twilio

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// A Workspace is a container for your Tasks, Workers, TaskQueues, Workflows, and Activities.
// refer: https://www.twilio.com/docs/taskrouter/api/workspace.
type Workspace struct {
	Sid                  *string            `json:"sid"`
	AccountSid           *string            `json:"account_sid"`
	DateCreated          *time.Time         `json:"date_created"`
	DateUpdated          *time.Time         `json:"date_updated"`
	DefaultActivityName  *string            `json:"default_activity_name"`
	DefaultActivitySid   *string            `json:"default_activity_sid"`
	TimeoutActivityName  *string            `json:"timeout_activity_name"`
	TimeoutActivitySid   *string            `json:"timeout_activity_sid"`
	URL                  *string            `json:"url"`
	Links                map[string]*string `json:"links"`
	FriendlyName         *string            `json:"friendly_name"`
	EventCallbackURL     *string            `json:"event_callback_url"`
	EventsFilter         *string            `json:"events_filter"`
	MultitaskEnabled     *bool              `json:"multi_task_enabled"`
	Template             *string            `json:"template"`
	PrioritizeQueueOrder *string            `json:"prioritize_queue_order"`
}

// WorkspaceParams workspace params for CRUD.
type WorkspaceParams struct {
	FriendlyName         *string `form:"FriendlyName,omitempty"`
	EventCallbackURL     *string `form:"EventCallbackUrl,omitempty"`
	EventsFilter         *string `form:"EventsFilter,omitempty"`
	MultitaskEnabled     *bool   `form:"MultitaskEnabled,omitempty"`
	Template             *string `form:"Template,omitempty"`
	PrioritizeQueueOrder *string `form:"PrioritizeQueueOrder,omitempty"`
}

// WorkspaceList struct to parse response of workspace read.
type WorkspaceList struct {
	Workspaces *[]Workspace `json:"workspaces"`
	Meta       *Meta        `json:"meta,omitempty"`
}

// WorkspaceQueryParams query params to read workspaces.
type WorkspaceQueryParams struct {
	FriendlyName *string `form:"FriendlyName,omitempty"`
	PageSize     *int    `form:"PageSize,omitempty"`
}

// WorkspaceClient is the entrypoint for the workspace CRUD.
type WorkspaceClient struct {
	ServiceURL string
	client     *Twilio
}

// NewWorkspaceClient constructs a new workspace Client.
func NewWorkspaceClient(client *Twilio) *WorkspaceClient {
	c := new(WorkspaceClient)
	c.client = client
	c.ServiceURL = fmt.Sprintf("https://taskrouter.%s/v1/Workspaces", client.BaseURL)

	return c
}

// Create creates workspace with the given the config.
func (c *WorkspaceClient) Create(workspaceParams WorkspaceParams) (*Workspace, error) {
	if len(*workspaceParams.FriendlyName) == 0 {
		return nil, errors.New("friendlyname is required in workspaceParams")
	}

	url := c.ServiceURL

	resp, err := c.client.Post(url, workspaceParams)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	workspace := &Workspace{}

	if err = json.NewDecoder(resp.Body).Decode(workspace); err != nil {
		return nil, err
	}

	return workspace, nil
}

// Fetch fetches workspace for the given workspace SID.
func (c *WorkspaceClient) Fetch(workspaceSID string) (*Workspace, error) {
	url := fmt.Sprintf("%s/%s", c.ServiceURL, workspaceSID)
	resp, err := c.client.Get(url, nil)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	workspace := &Workspace{}

	if err = json.NewDecoder(resp.Body).Decode(workspace); err != nil {
		return nil, err
	}

	return workspace, nil
}

// Read returns all existing workspaces.
func (c *WorkspaceClient) Read(queryParams *WorkspaceQueryParams) (*WorkspaceList, error) {
	url := c.ServiceURL

	resp, err := c.client.Get(url, queryParams)

	if err != nil {
		return nil, err
	}

	workspaces := &WorkspaceList{}

	if err = json.NewDecoder(resp.Body).Decode(&workspaces); err != nil {
		return nil, err
	}

	return workspaces, nil
}

// Update updates workspace with given config.
func (c *WorkspaceClient) Update(workspaceSID string, workspaceParams *WorkspaceParams) (*Workspace, error) {
	url := fmt.Sprintf("%s/%s", c.ServiceURL, workspaceSID)

	resp, err := c.client.Post(url, workspaceParams)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	workspace := &Workspace{}

	if err = json.NewDecoder(resp.Body).Decode(workspace); err != nil {
		return nil, err
	}

	return workspace, nil
}

// Delete deletes workspace for given SID.
func (c *WorkspaceClient) Delete(workspaceSID string) error {
	url := fmt.Sprintf("%s/%s", c.ServiceURL, workspaceSID)

	resp, err := c.client.Delete(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}