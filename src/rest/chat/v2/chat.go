package chat

import (
	"../../../base"
	"net/http"
	"net/url"
)

type ChatService struct {
	Request twilio.Request
}

func (chat *ChatService) Create(friendlyName string) (*http.Response, error) {
	values := url.Values{}
	values.Set("FriendlyName", friendlyName)
	resp, err := chat.Request.Post("/Services", values)
	return resp, err
}
