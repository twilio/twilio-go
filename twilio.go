// Package twilio provides bindings for Twilio's REST APIs.
package twilio

import (
	"net/http"
	"time"

	twilio "github.com/twilio/twilio-go/internal"
)

// Twilio provides access to Twilio services.
type Twilio struct {
	*twilio.Credentials
	*twilio.Client
	defaultbaseURL        *string
	common                service
	Chat                  *ChatClient
	Proxy                 *ProxyClient
	TaskRouter            *TaskRouterClient
	Studio                *StudioClient
	AvailablePhoneNumbers *AvailablePhoneNumbersClient
	IncomingPhoneNumbers  *IncomingPhoneNumberClient
	FlexFlow              *FlexFlowClient
	Sync                  *SyncClient
	Serverless            *ServerlessClient
}

type service struct {
	client *Twilio
}

// ChatClient holds all chat related resources.
type ChatClient struct {
	Service *ChatServiceClient
	Role    *ChatRoleClient
}

// ProxyClient holds all proxy related resources.
type ProxyClient struct {
	Service     *ProxyServiceClient
	PhoneNumber *ProxyPhoneNumberClient
}

// StudioClient holds all studio related resources.
type StudioClient struct {
	Flow *StudioFlowClient
}

// TaskRouterClient holds all studio related resources.
type TaskRouterClient struct {
	Workflows  *TaskRouterWorkflowClient
	Activities *TaskRouterActivityClient
	Workspaces *TaskRouterWorkspaceClient
	TaskQueues *TaskRouterTaskQueueClient
}

// ServerlessClient holds all runtime related resources.
type ServerlessClient struct {
	Service     *RuntimeServiceClient
	Environment *RuntimeEnvironmentClient
}

// SyncClient holds all sync related resources.
type SyncClient struct {
	Service *SyncServiceClient
}

// Meta holds relevant pagination resources.
type Meta struct {
	FirstPageURL    *string `json:"first_page_url"`
	Key             *string `json:"key"`
	LastPageURL     *string `json:"last_page_url,omitempty"`
	NextPageURL     *string `json:"next_page_url"`
	Page            *int    `json:"page"`
	PageSize        *int    `json:"page_size"`
	PreviousPageURL *string `json:"previous_page_url"`
	URL             *string `json:"url"`
}

// NewClient provides an initialized Twilio client.
func NewClient(accountSID string, authToken string) *Twilio {
	var httpClient = &http.Client{
		Timeout: time.Second * 10,
	}

	credentials := &twilio.Credentials{AccountSID: accountSID, AuthToken: authToken}

	c := &Twilio{
		Credentials: credentials,
		Client: &twilio.Client{
			Credentials: credentials,
			HTTPClient:  httpClient,
			BaseURL:     "twilio.com",
		},
	}

	c.common.client = c
	c.AvailablePhoneNumbers = NewAvailablePhoneNumbersClient(c)
	c.IncomingPhoneNumbers = NewIncomingPhoneNumberClient(c)
	c.FlexFlow = NewFlexFlowClient(c)
	c.Chat = &ChatClient{
		Service: NewChatServiceClient(c),
		Role:    NewChatRoleClient(c),
	}
	c.Proxy = &ProxyClient{
		Service:     NewProxyServiceClient(c),
		PhoneNumber: NewProxyPhoneNumberClient(c),
	}
	c.Studio = &StudioClient{
		Flow: NewStudioFlowClient(c),
	}
	c.TaskRouter = &TaskRouterClient{
		Activities: NewTaskRouterActivityClient(c),
		TaskQueues: NewTaskRouterTaskQueueClient(c),
		Workspaces: NewTaskRouterWorkspaceClient(c),
		Workflows:  NewTaskRouterWorkflowClient(c),
	}

	c.Serverless = &ServerlessClient{
		Service:     NewRuntimeServiceClient(c),
		Environment: NewRuntimeEnvironmentClient(c),
	}

	c.Sync = &SyncClient{
		Service: NewSyncServiceClient(c),
	}

	return c
}
