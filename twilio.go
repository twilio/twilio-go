// Package twilio provides bindings for Twilio's REST APIs.
package twilio

import (
	"net/http"
	"time"

	twilio "github.com/twilio/twilio-go/internal"
)

// Twilio provides access to Twilio services.
type Twilio struct {
	Chat                  *ChatClient
	Proxy                 *ProxyClient
	TaskRouter            *TaskRouterClient
	AvailablePhoneNumbers *AvailablePhoneNumbersClient
	IncomingPhoneNumbers  *IncomingPhoneNumberClient
}

// ChatClient holds all chat related resources.
type ChatClient struct {
	Service *ChatServiceClient
	Role    *ChatRoleClient
}

// ProxyClient holds all proxy related resources.
type ProxyClient struct {
	Service *ProxyServiceClient
}

const interval = 10

// NewClient provides an initialized Twilio client.
func NewClient(accountSid string, authToken string) *Twilio {
	var httpClient = &http.Client{
		Timeout: time.Second * interval,
	}

	credentials := twilio.Credentials{AccountSid: accountSid, AuthToken: authToken}
	client := &twilio.Client{Credentials: credentials, BaseURL: "twilio.com", HTTPClient: httpClient}

	twilioClient := Twilio{}
	twilioClient.AvailablePhoneNumbers = NewAvailablePhoneNumbersClient(client)
	twilioClient.IncomingPhoneNumbers = NewIncomingPhoneNumberClient(client)
	twilioClient.Chat = &ChatClient{
		Service: NewChatServiceClient(client),
		Role:    NewChatRoleClient(client),
	}
	twilioClient.Proxy = &ProxyClient{Service: NewProxyServiceClient(client)}
	twilioClient.TaskRouter = NewTaskRouterClient(client)

	return &twilioClient
}
