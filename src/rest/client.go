package twilio

import (
	. "../base"
	"./chat/v2"
)

type Client struct {
	Chat *chat.ChatService
}

func NewClient(AccountSid, AuthToken string) *Client {
	client := &Client{
		Chat: &chat.ChatService{Request: Request{AccountSid: AccountSid,
			AuthToken: AuthToken,
			BaseURL:   "https://chat.twilio.com/v2",
		}},
	}
	return client
}
