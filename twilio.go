// Package twilio provides bindings for Twilio's REST APIs.
package twilio

import (
	"net/http"

	twilio "github.com/twilio/twilio-go/internal"
)

// Twilio provides access to Twilio services.
type Twilio struct {
	*twilio.Credentials
	*twilio.Client
	defaultBaseURL        *string
	common                service
	Chat                  *ChatClient
	Proxy                 *ProxyClient
	TaskRouter            *TaskRouterClient
	AvailablePhoneNumbers *AvailablePhoneNumbersClient
	IncomingPhoneNumbers  *IncomingPhoneNumberClient
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

const interval = 10

// NewClient provides an initialized Twilio client.
func NewClient(accountSid string, authToken string) *Twilio {
	var httpClient = http.DefaultClient

	credentials := &twilio.Credentials{AccountSid: accountSid, AuthToken: authToken}

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
	c.Chat = &ChatClient{
		Service: NewChatServiceClient(c),
		Role:    NewChatRoleClient(c),
	}
	c.Proxy = &ProxyClient{
		Service:     NewProxyServiceClient(c),
		PhoneNumber: NewProxyPhoneNumberClient(c),
	}

	c.TaskRouter = NewTaskRouterClient(c)

	return c
}
