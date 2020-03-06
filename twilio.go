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
	Chat                  *chatClient
	Proxy                 *proxyClient
	TaskRouter            *taskRouterClient
	Studio                *studioClient
	AvailablePhoneNumbers *availablePhoneNumbersClient
	IncomingPhoneNumbers  *incomingPhoneNumberClient
	FlexFlow              *flexFlowClient
	Sync                  *syncClient
	Serverless            *serverlessClient
}

type service struct {
	client *Twilio
}

// chatClient holds all chat related resources.
type chatClient struct {
	Service *chatServiceClient
	Role    *chatRoleClient
}

// proxyClient holds all proxy related resources.
type proxyClient struct {
	Service     *proxyServiceClient
	PhoneNumber *proxyPhoneNumberClient
}

// studioClient holds all studio related resources.
type studioClient struct {
	Flow *studioFlowClient
}

// taskRouterClient holds all studio related resources.
type taskRouterClient struct {
	Workflows  *taskRouterWorkflowClient
	Activities *taskRouterActivityClient
	Workspaces *taskRouterWorkspaceClient
	TaskQueues *taskRouterTaskQueueClient
}

// serverlessClient holds all runtime related resources.
type serverlessClient struct {
	Service     *runtimeServiceClient
	Environment *runtimeEnvironmentClient
}

// syncClient holds all sync related resources.
type syncClient struct {
	Service *syncServiceClient
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
	c.AvailablePhoneNumbers = newAvailablePhoneNumbersClient(c)
	c.IncomingPhoneNumbers = newIncomingPhoneNumberClient(c)
	c.FlexFlow = newFlexFlowClient(c)
	c.Chat = &chatClient{
		Service: newChatServiceClient(c),
		Role:    newChatRoleClient(c),
	}
	c.Proxy = &proxyClient{
		Service:     newProxyServiceClient(c),
		PhoneNumber: newProxyPhoneNumberClient(c),
	}
	c.Studio = &studioClient{
		Flow: newStudioFlowClient(c),
	}
	c.TaskRouter = &taskRouterClient{
		Activities: newTaskRouterActivityClient(c),
		TaskQueues: newTaskRouterTaskQueueClient(c),
		Workspaces: newTaskRouterWorkspaceClient(c),
		Workflows:  newTaskRouterWorkflowClient(c),
	}

	c.Serverless = &serverlessClient{
		Service:     newRuntimeServiceClient(c),
		Environment: newRuntimeEnvironmentClient(c),
	}

	c.Sync = &syncClient{
		Service: newSyncServiceClient(c),
	}

	return c
}
