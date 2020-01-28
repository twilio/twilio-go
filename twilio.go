// Package twilio provides bindings for Twilio's REST APIs.
package twilio

import (
	"net/http"
	"time"

	twilio "github.com/twilio/twilio-go/internal"
)

// Twilio provides access to Twilio services.
type Twilio struct {
	Chat       *ChatClient
	TaskRouter *TaskRouterClient
}

type service interface {
	Initialize(*twilio.Client)
}

// ChatClient holds all chat related resources.
type ChatClient struct {
	Service *ChatServiceClient
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
	twilioClient.Chat = &ChatClient{
		Service: NewChatServiceClient(client),
	}
	twilioClient.TaskRouter = NewTaskRouterClient(client)

	return &twilioClient
}
