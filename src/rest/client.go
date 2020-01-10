package twilio

import (
	. "../base"
	"./chat/v2"
)

type Client struct {
	Chat *chat.ChatService
}
type Twilio struct {
	Chat *chat.Chat
}

func NewClient(AccountSid string, AuthToken string) *Twilio {
	client := &Twilio{}
	credentials := Credentials{AccountSid: AccountSid, AuthToken: AuthToken}

	client.Chat = &chat.Chat{Request: Request{Credentials: credentials, BaseURL: "https://chat.twilio.com/v2"}}

	return client
}
