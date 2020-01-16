package twilio

import (
	twilio "github.com/twilio/twilio-go"
	"github.com/twilio/twilio-go/chat"
)

// Twilio provides access to Twilio services.
type Twilio struct {
	Chat *chat.Client
}

// NewClient provides an initialized Twilio client.
func NewClient(AccountSid string, AuthToken string) *Twilio {
	client := &Twilio{}
	credentials := twilio.Credentials{AccountSid: AccountSid, AuthToken: AuthToken}

	client.Chat = &chat.Client{Request: twilio.Request{Credentials: credentials, BaseURL: "https://chat.twilio.com/v2"}}
	return client
}
