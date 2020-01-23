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
	Request    *twilio.Request
	Chat       *Chat
	TaskRouter *TaskRouter
}

type service interface {
	Create(*twilio.Request)
}

const interval = 10

// NewClient provides an initialized Twilio client.
func NewClient(accountSid string, authToken string) *Twilio {
	var httpClient = &http.Client{
		Timeout: time.Second * interval,
	}

	credentials := twilio.Credentials{AccountSid: accountSid, AuthToken: authToken}

	client := &Twilio{}
	request := &twilio.Request{Credentials: credentials, BaseURL: "twilio.com", Client: httpClient}
	client.Chat = &Chat{}
	client.TaskRouter = &TaskRouter{}

	cRef := reflect.ValueOf(client)
	for i := 0; i < cRef.NumField(); i++ {
		switch v := cRef.Field(i).Interface().(type) { //nolint
		case service:
			v.Create(request)
		}
	}

	return client
}
