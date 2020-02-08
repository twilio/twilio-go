package twilio

import (
	"fmt"
)

// TaskRouterClient is the entrypoint for the TaskRouter API.
type TaskRouterClient struct {
	WorkflowClient  *WorkflowClient
	ActivityClient  *ActivityClient
	WorkspaceClient *WorkspaceClient
	TaskQueueClient *TaskQueueClient
}

// Meta pagination metadata.
type Meta struct {
	FirstPageURL    *string `json:"first_page_url"`
	Key             *string `json:"key"`
	LastPageURL     *string `json:"last_page_url"`
	NextPageURL     *string `json:"next_page_url"`
	Page            *int    `json:"page"`
	PageSize        *int    `json:"page_size"`
	PreviousPageURL *string `json:"previous_page_url"`
	URL             *string `json:"url"`
}

// constants
const (
	//WorkspaceDefaultFriendlyName = "Flex Task Assignment"
	WorkspaceDefaultFriendlyName = "Flex Task Assignment"
	WorkflowDefaultFriendlyName  = "Assign to Anyone"
	TaskQueueDefaultFriendlyName = "Everyone"
	ActivityDefaultFriendlyName  = "Break"
)

// NewTaskRouterClient constructs a new TaskRouter Client.
func NewTaskRouterClient(twilioClient *Twilio) *TaskRouterClient {
	c := new(TaskRouterClient)
	serviceURL := fmt.Sprintf("https://taskrouter.%s/v1/Workspaces", twilioClient.BaseURL)

	c.WorkflowClient = &WorkflowClient{
		client:     twilioClient,
		ServiceURL: serviceURL,
	}

	c.ActivityClient = &ActivityClient{
		client:     twilioClient,
		ServiceURL: serviceURL,
	}

	c.WorkspaceClient = &WorkspaceClient{
		client:     twilioClient,
		ServiceURL: serviceURL,
	}

	c.TaskQueueClient = &TaskQueueClient{
		client:     twilioClient,
		ServiceURL: serviceURL,
	}

	return c
}
