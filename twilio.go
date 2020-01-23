// Package twilio provides bindings for Twilio's REST APIs.
package twilio

import (
	"net/http"
	"reflect"
	"time"

	twilio "github.com/twilio/twilio-go/internal"
)

// Twilio provides access to Twilio services.
type Twilio struct {
	Client     *twilio.Client
	Chat       *Chat
	TaskRouter *TaskRouter
}

type service interface {
	Create(*twilio.Client)
}

const interval = 10

// NewClient provides an initialized Twilio client.
func NewClient(accountSid string, authToken string) *Twilio {
	var httpClient = &http.Client{
		Timeout: time.Second * interval,
	}

	credentials := twilio.Credentials{AccountSid: accountSid, AuthToken: authToken}

	twilioClient := &Twilio{}
	client := &twilio.Client{Credentials: credentials, BaseURL: "twilio.com", HTTPClient: httpClient}
	twilioClient.Chat = &Chat{}
	twilioClient.TaskRouter = &TaskRouter{}

	cRef := reflect.ValueOf(client)
	for i := 0; i < cRef.NumField(); i++ {
		switch v := cRef.Field(i).Interface().(type) { //nolint // Avoids "one condition switch" complaint, which we ignore because (type) can only be used within switch statements.
		case service:
			v.Create(client)
		}
	}

	return twilioClient
}
