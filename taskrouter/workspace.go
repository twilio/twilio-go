package taskrouter

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	twilio "github.com/twilio/twilio-go/internal"
)

// refer: https://www.twilio.com/docs/taskrouter/api/workspace.

// Workspace struct.
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

// WorkspaceParams workspace params for CRUD
type WorkspaceParams struct {
	FriendlyName         string `url:"FriendlyName,omitempty"`
	EventCallbackURL     string `url:"EventCallbackUrl,omitempty"`
	EventsFilter         string `url:"EventsFilter,omitempty"`
	MultitaskEnabled     bool   `url:"MultitaskEnabled,omitempty"`
	Template             string `url:"Template,omitempty"`
	PrioritizeQueueOrder string `url:"PrioritizeQueueOrder,omitempty"`
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
	url := ws.ServiceURL

	resp, err := ws.Client.Post(url, workspaceParams)

	if err != nil {
		log.Printf("error creating workspace: %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	workspace := &Workspace{}

	if decodeErr := json.NewDecoder(resp.Body).Decode(workspace); decodeErr != nil {
		log.Printf("error decoding the output of workspace create: %s", decodeErr)
		return nil, decodeErr
	}

	return workspace, nil
}

// Fetch fetches workspace for the given workspace SID.
func (ws *WorkspaceClient) Fetch(workspaceSID string) (*Workspace, error) {
	url := fmt.Sprintf("%s/%s", ws.ServiceURL, workspaceSID)
	resp, err := ws.Client.Get(url, nil)

	if err != nil {
		log.Printf("error fetching workspace: %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	workspace := &Workspace{}

	if decodeErr := json.NewDecoder(resp.Body).Decode(workspace); decodeErr != nil {
		log.Printf("error decoding the output of workspace fetch: %s", decodeErr)
		return nil, decodeErr
	}

	return workspace, nil
}

// Read returns all existing workspaces.
func (ws *WorkspaceClient) Read() (*[]Workspace, error) {
	url := ws.ServiceURL

	resp, err := ws.Client.Get(url, nil)

	if err != nil {
		log.Printf("error reading workspaces: %s", err)
		return nil, err
	}

	workspaces := make([]Workspace, 0)

	if decodeErr := json.NewDecoder(resp.Body).Decode(&workspaces); decodeErr != nil {
		log.Printf("error decoding the output of workspace read: %s", decodeErr)
		return nil, decodeErr
	}

	return &workspaces, nil
}

// Update updates workspace with given config.
func (ws *WorkspaceClient) Update(workspaceParams WorkspaceParams, workspaceSID string) (*Workspace, error) {
	url := fmt.Sprintf("%s/%s", ws.ServiceURL, workspaceSID)

	resp, err := ws.Client.Post(url, workspaceParams)

	if err != nil {
		log.Printf("error updating workspace: %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	workspace := &Workspace{}

	if decodeErr := json.NewDecoder(resp.Body).Decode(workspace); decodeErr != nil {
		log.Printf("error decoding the output of workspace update: %s", decodeErr)
		return nil, decodeErr
	}

	return workspace, nil
}

// Delete deletes workspace for given SID.
func (ws *WorkspaceClient) Delete(workspaceSID string) error {
	url := fmt.Sprintf("%s/%s", ws.ServiceURL, workspaceSID)

	_, err := ws.Client.Delete(url)

	return err
}
