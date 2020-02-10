package twilio

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// Workflow controls how tasks will be prioritized and routed into Queues,
// and how Tasks should escalate in priority or move across queues over time.
// refer: https://www.twilio.com/docs/taskrouter/api/workflow.
type Workflow struct {
	AssignmentCallbackURL         *string            `json:"assignment_callback_url"`
	Configuration                 *string            `json:"configuration"`
	AccountSid                    *string            `json:"account_sid"`
	DateCreated                   *time.Time         `json:"date_created"`
	DateUpdated                   *time.Time         `json:"date_updated"`
	DocumentContentType           *string            `json:"document_content_type"`
	FallbackAssignmentCallbackURL *string            `json:"fallback_assignment_callback_url"`
	FriendlyName                  *string            `json:"friendly_name"`
	Sid                           *string            `json:"sid"`
	TaskReservationTimeout        *int               `json:"task_reservation_timeout"`
	WorkspaceSid                  *string            `json:"workspace_sid"`
	URL                           *string            `json:"url"`
	Links                         map[string]*string `json:"links"`
}

// WorkflowParams workflow parameters.
type WorkflowParams struct {
	FriendlyName                  *string `form:",omitempty"`
	Configuration                 *string `form:",omitempty"`
	AssignmentCallbackURL         *string `form:"AssignmentCallbackUrl,omitempty"`
	FallbackAssignmentCallbackURL *string `form:"FallbackAssignmentCallbackUrl,omitempty"`
	TaskReservationTimeout        *int    `form:",omitempty"`
}

// WorkflowList struct to parse response of workspace read.
type WorkflowList struct {
	Workflows *[]Workflow `json:"workflows"`
	Meta      *Meta       `json:"meta,omitempty"`
}

// WorkflowQueryParams query params to read workspaces.
type WorkflowQueryParams struct {
	FriendlyName *string `form:",omitempty"`
	PageSize     *int    `form:",omitempty"`
}

// WorkflowClient is the entrypoint for the workflow CRUD.
type WorkflowClient struct {
	ServiceURL string
	client     *Twilio
}

// NewWorkflowClient constructs a new workflow Client.
func NewWorkflowClient(client *Twilio) *WorkflowClient {
	c := new(WorkflowClient)
	c.client = client
	c.ServiceURL = fmt.Sprintf("https://taskrouter.%s/v1/Workspaces", client.BaseURL)

	return c
}

// Create creates workflow with the given the config.
func (c *WorkflowClient) Create(workspaceSID string, workflowParams *WorkflowParams) (*Workflow, error) {
	if len(*workflowParams.FriendlyName) == 0 {
		return nil, errors.New("friendly name is required in workflowParams")
	}

	if len(*workflowParams.Configuration) == 0 {
		return nil, errors.New("configuration is required in workflowParams")
	}

	url := fmt.Sprintf("%s/%s/%s", c.ServiceURL, workspaceSID, "Workflows")

	resp, err := c.client.Post(url, workflowParams)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	workflow := &Workflow{}

	if err = json.NewDecoder(resp.Body).Decode(workflow); err != nil {
		return nil, err
	}

	return workflow, nil
}

// Fetch fetches workflow for the given workspace SID.
func (c *WorkflowClient) Fetch(workspaceSID string, workflowSID string) (*Workflow, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", c.ServiceURL, workspaceSID, "Workflows", workflowSID)
	resp, err := c.client.Get(url, nil)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	workflow := &Workflow{}

	if err = json.NewDecoder(resp.Body).Decode(workflow); err != nil {
		return nil, err
	}

	return workflow, nil
}

// Read returns all existing workflows.
func (c *WorkflowClient) Read(workspaceSID string, queryParams *WorkflowQueryParams) (*[]Workflow, error) {
	url := fmt.Sprintf("%s/%s/%s", c.ServiceURL, workspaceSID, "Workflows")

	resp, err := c.client.Get(url, queryParams)

	if err != nil {
		return nil, err
	}

	workflows := make([]Workflow, 0)

	if err = json.NewDecoder(resp.Body).Decode(&workflows); err != nil {
		return nil, err
	}

	return &workflows, nil
}

// Update updates workflow with given config.
func (c *WorkflowClient) Update(workspaceSID string,
	workflowSID string, workflowParams *WorkflowParams) (*Workflow, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", c.ServiceURL, workspaceSID, "Workflows", workflowSID)

	resp, err := c.client.Post(url, workflowParams)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	workflow := &Workflow{}

	if err := json.NewDecoder(resp.Body).Decode(workflow); err != nil {
		return nil, err
	}

	return workflow, nil
}

// Delete deletes workflow for given SID.
func (c *WorkflowClient) Delete(workspaceSID string, workflowSID string) error {
	url := fmt.Sprintf("%s/%s/%s/%s", c.ServiceURL, workspaceSID, "Workflows", workflowSID)

	resp, err := c.client.Delete(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
