package twilio

import (
	"net/http"
	"time"

	twilio "github.com/twilio/twilio-go"
	"github.com/twilio/twilio-go/chat"
)

// Twilio provides access to Twilio services.
type Twilio struct {
	Chat *chat.Client
}

const interval = 10

// NewClient provides an initialized Twilio client.
func NewClient(accountSid string, authToken string) *Twilio {
	var httpClient = &http.Client{
		Timeout: time.Second * interval,
	}

	credentials := twilio.Credentials{AccountSID: accountSid, AuthToken: authToken}

	client := &Twilio{}
	client.Chat = &chat.Client{Request: twilio.Client{Credentials: credentials,
		BaseURL: "https://chat.twilio.com/v2",
		Client:  httpClient}}

	return client
}
