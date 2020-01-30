package twilio

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	twilio "github.com/twilio/twilio-go/internal"
)

// A Workspace is a container for your Tasks, Workers, TaskQueues, Workflows, and Activities.
// refer: https://www.twilio.com/docs/taskrouter/api/workspace.

// Workspace struct.
type Workspace struct {
	Sid                  string            `json:"sid"`
	AccountSid           string            `json:"account_sid"`
	DateCreated          time.Time         `json:"date_created"`
	DateUpdated          time.Time         `json:"date_updated"`
	DefaultActivityName  *string           `json:"default_activity_name"`
	DefaultActivitySid   *string           `json:"default_activity_sid"`
	TimeoutActivityName  *string           `json:"timeout_activity_name"`
	TimeoutActivitySid   *string           `json:"timeout_activity_sid"`
	URL                  *string           `json:"url"`
	Links                map[string]string `json:"links"`
	FriendlyName         string            `json:"friendly_name"`
	EventCallbackURL     *string           `json:"event_callback_url"`
	EventsFilter         *string           `json:"events_filter"`
	MultitaskEnabled     *bool             `json:"multi_task_enabled"`
	Template             *string           `json:"template"`
	PrioritizeQueueOrder *string           `json:"prioritize_queue_order"`
}

// WorkspaceParams workspace params for CRUD.
type WorkspaceParams struct {
	FriendlyName         string  `url:"FriendlyName,omitempty"`
	EventCallbackURL     *string `url:"EventCallbackUrl,omitempty"`
	EventsFilter         *string `url:"EventsFilter,omitempty"`
	MultitaskEnabled     *bool   `url:"MultitaskEnabled,omitempty"`
	Template             *string `url:"Template,omitempty"`
	PrioritizeQueueOrder *string `url:"PrioritizeQueueOrder,omitempty"`
}

// WorkspaceList struct to parse response of workspace read.
type WorkspaceList struct {
	Workspaces *[]Workspace `json:"workspaces"`
	Meta       *Meta        `json:"meta,omitempty"`
}

// WorkspaceQueryParams query params to read workspaces.
type WorkspaceQueryParams struct {
	FriendlyName *string `url:"FriendlyName,omitempty"`
	PageSize     *int    `url:"PageSize,omitempty"`
}

// WorkspaceClient is the entrypoint for the workspace CRUD.
type WorkspaceClient struct {
	ServiceURL string
	Client     *twilio.Client
}

// NewWorkspaceClient constructs a new workspace Client.
func NewWorkspaceClient(twilioClient *twilio.Client) *WorkspaceClient {
	c := new(WorkspaceClient)
	c.Client = twilioClient
	c.ServiceURL = fmt.Sprintf("https://taskrouter.%s/v1/Workspaces", twilioClient.BaseURL)

	return c
}

// Create creates workspace with the given the config.
func (ws *WorkspaceClient) Create(workspaceParams WorkspaceParams) (*Workspace, error) {
	if len(workspaceParams.FriendlyName) == 0 {
		return nil, errors.New("friendlyname is required in workspaceParams")
	}

	url := ws.ServiceURL

	resp, err := ws.Client.Post(url, workspaceParams)

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
func (ws *WorkspaceClient) Fetch(workspaceSID string) (*Workspace, error) {
	url := fmt.Sprintf("%s/%s", ws.ServiceURL, workspaceSID)
	resp, err := ws.Client.Get(url, nil)

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
func (ws *WorkspaceClient) Read(queryParams *WorkspaceQueryParams) (*WorkspaceList, error) {
	url := ws.ServiceURL

	resp, err := ws.Client.Get(url, queryParams)

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
func (ws *WorkspaceClient) Update(workspaceParams WorkspaceParams, workspaceSID string) (*Workspace, error) {
	url := fmt.Sprintf("%s/%s", ws.ServiceURL, workspaceSID)

	resp, err := ws.Client.Post(url, workspaceParams)

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
func (ws *WorkspaceClient) Delete(workspaceSID string) error {
	url := fmt.Sprintf("%s/%s", ws.ServiceURL, workspaceSID)

	resp, err := ws.Client.Delete(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
