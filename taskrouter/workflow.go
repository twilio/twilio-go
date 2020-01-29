package taskrouter

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	twilio "github.com/twilio/twilio-go/internal"
)

//refer: https://www.twilio.com/docs/taskrouter/api/workflow.

// Workflow workflow struct.
type Workflow struct {
	AssignmentCallbackURL         string            `json:"assignment_callback_url"`
	Configuration                 string            `json:"configuration"`
	AccountSid                    string            `json:"account_sid"`
	DateCreated                   time.Time         `json:"date_created"`
	DateUpdated                   time.Time         `json:"date_updated"`
	DocumentContentType           string            `json:"document_content_type"`
	FallbackAssignmentCallbackURL string            `json:"fallback_assignment_callback_url"`
	FriendlyName                  string            `json:"friendly_name"`
	Sid                           string            `json:"sid"`
	TaskReservationTimeout        int               `json:"task_reservation_timeout"`
	WorkspaceSid                  string            `json:"workspace_sid"`
	URL                           string            `json:"url"`
	Links                         map[string]string `json:"links"`
}

// WorkflowParams workflow parameters.
type WorkflowParams struct {
	WorkspaceSid                  string `url:"WorkspaceSid,omitempty"`
	FriendlyName                  string `url:"FriendlyName,omitempty"`
	Configuration                 string `url:"Configuration,omitempty"`
	AssignmentCallbackURL         string `url:"AssignmentCallbackURL,omitempty"`
	FallbackAssignmentCallbackURL string `url:"FallbackAssignmentCallbackURL,omitempty"`
	TaskReservationTimeout        int    `url:"TaskReservationTimeout,omitempty"`
}

// WorkflowClient is the entrypoint for the workflow CRUD.
type WorkflowClient struct {
	ServiceURL string
	Client     *twilio.Client
}

// NewWorkflowClient constructs a new workflow Client.
func NewWorkflowClient(twilioClient *twilio.Client) *WorkflowClient {
	c := new(WorkflowClient)
	c.Client = twilioClient
	c.ServiceURL = fmt.Sprintf("https://taskrouter.%s/v1/Workspaces", twilioClient.BaseURL)

	return c
}

// Create creates workflow with the given the config.
func (wf *WorkflowClient) Create(workflowParams WorkflowParams, workspaceSID string) (*Workflow, error) {
	url := fmt.Sprintf("%s/%s/%s", wf.ServiceURL, workspaceSID, "Workflows")

	resp, err := wf.Client.Post(url, workflowParams)

	if err != nil {
		log.Printf("error creating workflow: %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	workflow := &Workflow{}

	if decodeErr := json.NewDecoder(resp.Body).Decode(workflow); decodeErr != nil {
		log.Printf("error decoding the output of workflow create: %s", decodeErr)
		return nil, decodeErr
	}

	return workflow, nil
}

// Fetch fetches workflow for the given workspace SID.
func (wf *WorkflowClient) Fetch(workspaceSID string, workflowSID string) (*Workflow, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", wf.ServiceURL, workspaceSID, "Workflows", workflowSID)
	resp, err := wf.Client.Get(url, nil)

	if err != nil {
		log.Printf("error fetching workflow: %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	workflow := &Workflow{}

	if decodeErr := json.NewDecoder(resp.Body).Decode(workflow); decodeErr != nil {
		log.Printf("error decoding the output of workflow fetch: %s", decodeErr)
		return nil, decodeErr
	}

	return workflow, nil
}

// Read returns all existing workflows.
func (wf *WorkflowClient) Read(workspaceSID string) (*[]Workflow, error) {
	url := fmt.Sprintf("%s/%s/%s", wf.ServiceURL, workspaceSID, "Workflows")

	resp, err := wf.Client.Get(url, nil)

	if err != nil {
		log.Printf("error reading workflows: %s", err)
		return nil, err
	}

	workflows := make([]Workflow, 0)

	if decodeErr := json.NewDecoder(resp.Body).Decode(&workflows); decodeErr != nil {
		log.Printf("error decoding the output of workflow read: %s", decodeErr)
		return nil, decodeErr
	}

	return &workflows, nil
}

// Update updates workflow with given config.
func (wf *WorkflowClient) Update(workflowParams WorkflowParams, workspaceSID string,
	workflowSID string) (*Workflow, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", wf.ServiceURL, workspaceSID, "Workflows", workflowSID)

	resp, err := wf.Client.Post(url, workflowParams)

	if err != nil {
		log.Printf("error updating workflow: %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	workflow := &Workflow{}

	if decodeErr := json.NewDecoder(resp.Body).Decode(workflow); decodeErr != nil {
		log.Printf("error decoding the output of workflow fetch: %s", decodeErr)
		return nil, decodeErr
	}

	return workflow, nil
}

// Delete deletes workflow for given SID.
func (wf *WorkflowClient) Delete(workspaceSID string, workflowSID string) error {
	url := fmt.Sprintf("%s/%s/%s/%s", wf.ServiceURL, workspaceSID, "Workflows", workflowSID)

	_, err := wf.Client.Delete(url)

	return err
}
