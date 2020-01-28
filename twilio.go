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
	Chat        *Chat
	TaskRouter  *TaskRouter
	PhoneNumber *PhoneNumberClient
}

type service interface {
	Initialize(*twilio.Client)
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
	twilioClient.Chat = new(Chat)
	twilioClient.TaskRouter = new(TaskRouter)
	twilioClient.PhoneNumber = new(PhoneNumberClient)

	tcRef := reflect.ValueOf(twilioClient)
	for i := 0; i < tcRef.NumField(); i++ {
		s := tcRef.Field(i).Interface().(service)
		s.Initialize(client)
	}

	return &twilioClient
}
