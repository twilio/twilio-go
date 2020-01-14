package twilio

import (
	"../base/backend"
	"./chat/v2"
)

// Twilio provides access to Twilio services.
type Twilio struct {
	Chat *chat.Chat
}

// NewClient provides an initialized Twilio client.
func NewClient(AccountSid string, AuthToken string) *Twilio {
	client := &Twilio{}
	credentials := backend.Credentials{AccountSid: AccountSid, AuthToken: AuthToken}

	client.Chat = &chat.Chat{Request: backend.Request{Credentials: credentials, BaseURL: "https://chat.twilio.com/v2"}}

	return client
}
