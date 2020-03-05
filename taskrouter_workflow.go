package twilio

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// TaskRouterWorkflow controls how tasks will be prioritized and routed into Queues,
// and how Tasks should escalate in priority or move across queues over time.
// refer: https://www.twilio.com/docs/taskrouter/api/workflow.
type TaskRouterWorkflow struct {
	AssignmentCallbackURL         *string            `json:"assignment_callback_url"`
	Configuration                 *string            `json:"configuration"`
	AccountSID                    *string            `json:"account_sid"`
	DateCreated                   *time.Time         `json:"date_created"`
	DateUpdated                   *time.Time         `json:"date_updated"`
	DocumentContentType           *string            `json:"document_content_type"`
	FallbackAssignmentCallbackURL *string            `json:"fallback_assignment_callback_url"`
	FriendlyName                  *string            `json:"friendly_name"`
	SID                           *string            `json:"sid"`
	TaskReservationTimeout        *int               `json:"task_reservation_timeout"`
	WorkspaceSID                  *string            `json:"workspace_sid"`
	URL                           *string            `json:"url"`
	Links                         map[string]*string `json:"links"`
}

// TaskRouterWorkflowParams TaskRouterWorkflow parameters.
type TaskRouterWorkflowParams struct {
	FriendlyName                  *string `form:",omitempty"`
	Configuration                 *string `form:",omitempty"`
	AssignmentCallbackURL         *string `form:"AssignmentCallbackUrl,omitempty"`
	FallbackAssignmentCallbackURL *string `form:"FallbackAssignmentCallbackUrl,omitempty"`
	TaskReservationTimeout        *int    `form:",omitempty"`
}

// TaskRouterWorkflowList struct to parse response of workspace read.
type TaskRouterWorkflowList struct {
	Workflows []*TaskRouterWorkflow `json:"workflows"`
	Meta      *Meta                 `json:"meta,omitempty"`
}

// TaskRouterWorkflowQueryParams query params to read workspaces.
type TaskRouterWorkflowQueryParams struct {
	FriendlyName *string `form:",omitempty"`
	PageSize     *int    `form:",omitempty"`
}

// TaskRouterWorkflowClient is the entrypoint for the TaskRouterWorkflow CRUD.
type TaskRouterWorkflowClient struct {
	baseURL string
	client  *Twilio
}

// newTaskRouterWorkflowClient constructs a new TaskRouterWorkflow Client.
func newTaskRouterWorkflowClient(client *Twilio) *TaskRouterWorkflowClient {
	c := new(TaskRouterWorkflowClient)
	c.client = client
	c.baseURL = c.url(fmt.Sprintf("https://taskrouter.%s/v1", client.BaseURL))

	return c
}

// Create creates TaskRouterWorkflow with the given the config.
func (c *TaskRouterWorkflowClient) Create(
	workspaceSID string,
	params *TaskRouterWorkflowParams,
) (*TaskRouterWorkflow, error) {
	if params.FriendlyName == nil {
		return nil, errors.New("friendly name is required in params")
	}

	if params.Configuration == nil {
		return nil, errors.New("configuration is required in params")
	}

	url := c.url(fmt.Sprintf("/Workspaces/%s/Workflows", workspaceSID))

	resp, err := c.client.Post(url, params)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	TaskRouterWorkflow := &TaskRouterWorkflow{}

	if err = json.NewDecoder(resp.Body).Decode(TaskRouterWorkflow); err != nil {
		return nil, err
	}

	return TaskRouterWorkflow, nil
}

// Fetch fetches TaskRouterWorkflow for the given workspace SID.
func (c *TaskRouterWorkflowClient) Fetch(workspaceSID string, sid string) (*TaskRouterWorkflow, error) {
	url := c.url(fmt.Sprintf("/Workspaces/%s/Workflows/%s", workspaceSID, sid))
	resp, err := c.client.Get(url, nil)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	TaskRouterWorkflow := &TaskRouterWorkflow{}

	if err = json.NewDecoder(resp.Body).Decode(TaskRouterWorkflow); err != nil {
		return nil, err
	}

	return TaskRouterWorkflow, nil
}

// Read returns all existing Workflows.
func (c *TaskRouterWorkflowClient) Read(
	workspaceSID string,
	params *TaskRouterWorkflowQueryParams,
) (*TaskRouterWorkflowList, error) {
	url := c.url(fmt.Sprintf("/Workspaces/%s/Workflows", workspaceSID))

	resp, err := c.client.Get(url, params)

	if err != nil {
		return nil, err
	}

	w := &TaskRouterWorkflowList{}

	if err = json.NewDecoder(resp.Body).Decode(w); err != nil {
		return nil, err
	}

	return w, nil
}

// Update updates TaskRouterWorkflow with given config.
func (c *TaskRouterWorkflowClient) Update(
	workspaceSID string,
	sid string,
	params *TaskRouterWorkflowParams,
) (*TaskRouterWorkflow, error) {
	url := c.url(fmt.Sprintf("/Workspaces/%s/Workflows/%s", workspaceSID, sid))

	resp, err := c.client.Post(url, params)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	TaskRouterWorkflow := &TaskRouterWorkflow{}

	if err := json.NewDecoder(resp.Body).Decode(TaskRouterWorkflow); err != nil {
		return nil, err
	}

	return TaskRouterWorkflow, nil
}

// Delete deletes TaskRouterWorkflow for given SID.
func (c *TaskRouterWorkflowClient) Delete(workspaceSID string, sid string) error {
	url := c.url(fmt.Sprintf("/Workspaces/%s/Workflows/%s", workspaceSID, sid))

	resp, err := c.client.Delete(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}

func (c *TaskRouterWorkflowClient) url(path string) string {
	if c.client.defaultbaseURL != nil {
		return *c.client.defaultbaseURL + path
	}

	return c.baseURL + path
}
