package twilio

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// TaskRouterWorkflow controls how tasks will be prioritized and routed into Queues,
// and how Tasks should escalate in priority or move across queues over time.
// refer: https://www.twilio.com/docs/taskrouter/api/TaskRouterworkflow.
type TaskRouterWorkflow struct {
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

// TaskRouterWorkflowParams TaskRouterworkflow parameters.
type TaskRouterWorkflowParams struct {
	FriendlyName                  *string `form:",omitempty"`
	Configuration                 *string `form:",omitempty"`
	AssignmentCallbackURL         *string `form:"AssignmentCallbackUrl,omitempty"`
	FallbackAssignmentCallbackURL *string `form:"FallbackAssignmentCallbackUrl,omitempty"`
	TaskReservationTimeout        *int    `form:",omitempty"`
}

// TaskRouterWorkflowList struct to parse response of workspace read.
type TaskRouterWorkflowList struct {
	TaskRouterWorkflows *[]TaskRouterWorkflow `json:"TaskRouterworkflows"`
	Meta                *Meta                 `json:"meta,omitempty"`
}

// TaskRouterWorkflowQueryParams query params to read workspaces.
type TaskRouterWorkflowQueryParams struct {
	FriendlyName *string `form:",omitempty"`
	PageSize     *int    `form:",omitempty"`
}

// TaskRouterWorkflowClient is the entrypoint for the TaskRouterworkflow CRUD.
type TaskRouterWorkflowClient struct {
	ServiceURL string
	client     *Twilio
}

// NewTaskRouterWorkflowClient constructs a new TaskRouterworkflow Client.
func NewTaskRouterWorkflowClient(client *Twilio) *TaskRouterWorkflowClient {
	c := new(TaskRouterWorkflowClient)
	c.client = client
	c.ServiceURL = fmt.Sprintf("https://taskrouter.%s/v1/Workspaces", client.BaseURL)

	return c
}

// Create creates TaskRouterworkflow with the given the config.
func (c *TaskRouterWorkflowClient) Create(workspaceSID string, TaskRouterworkflowParams *TaskRouterWorkflowParams) (*TaskRouterWorkflow, error) {
	if len(*TaskRouterworkflowParams.FriendlyName) == 0 {
		return nil, errors.New("friendly name is required in TaskRouterworkflowParams")
	}

	if len(*TaskRouterworkflowParams.Configuration) == 0 {
		return nil, errors.New("configuration is required in TaskRouterworkflowParams")
	}

	url := fmt.Sprintf("%s/%s/%s", c.ServiceURL, workspaceSID, "TaskRouterWorkflows")

	resp, err := c.client.Post(url, TaskRouterworkflowParams)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	TaskRouterworkflow := &TaskRouterWorkflow{}

	if err = json.NewDecoder(resp.Body).Decode(TaskRouterworkflow); err != nil {
		return nil, err
	}

	return TaskRouterworkflow, nil
}

// Fetch fetches TaskRouterworkflow for the given workspace SID.
func (c *TaskRouterWorkflowClient) Fetch(workspaceSID string, TaskRouterworkflowSID string) (*TaskRouterWorkflow, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", c.ServiceURL, workspaceSID, "TaskRouterWorkflows", TaskRouterworkflowSID)
	resp, err := c.client.Get(url, nil)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	TaskRouterworkflow := &TaskRouterWorkflow{}

	if err = json.NewDecoder(resp.Body).Decode(TaskRouterworkflow); err != nil {
		return nil, err
	}

	return TaskRouterworkflow, nil
}

// Read returns all existing TaskRouterworkflows.
func (c *TaskRouterWorkflowClient) Read(workspaceSID string, queryParams *TaskRouterWorkflowQueryParams) (*[]TaskRouterWorkflow, error) {
	url := fmt.Sprintf("%s/%s/%s", c.ServiceURL, workspaceSID, "TaskRouterWorkflows")

	resp, err := c.client.Get(url, queryParams)

	if err != nil {
		return nil, err
	}

	TaskRouterworkflows := make([]TaskRouterWorkflow, 0)

	if err = json.NewDecoder(resp.Body).Decode(&TaskRouterworkflows); err != nil {
		return nil, err
	}

	return &TaskRouterworkflows, nil
}

// Update updates TaskRouterworkflow with given config.
func (c *TaskRouterWorkflowClient) Update(workspaceSID string,
	TaskRouterworkflowSID string, TaskRouterworkflowParams *TaskRouterWorkflowParams) (*TaskRouterWorkflow, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", c.ServiceURL, workspaceSID, "TaskRouterWorkflows", TaskRouterworkflowSID)

	resp, err := c.client.Post(url, TaskRouterworkflowParams)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	TaskRouterworkflow := &TaskRouterWorkflow{}

	if err := json.NewDecoder(resp.Body).Decode(TaskRouterworkflow); err != nil {
		return nil, err
	}

	return TaskRouterworkflow, nil
}

// Delete deletes TaskRouterworkflow for given SID.
func (c *TaskRouterWorkflowClient) Delete(workspaceSID string, TaskRouterworkflowSID string) error {
	url := fmt.Sprintf("%s/%s/%s/%s", c.ServiceURL, workspaceSID, "TaskRouterWorkflows", TaskRouterworkflowSID)

	resp, err := c.client.Delete(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
