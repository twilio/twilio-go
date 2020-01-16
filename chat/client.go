package chat

import (
	"encoding/json"
	"fmt"
	twilio "github.com/twilio/twilio-go"
)

type Client struct {
	Request twilio.Request
}

// Create creates a new Service.
func (c Client) Create(params *twilio.ChatServiceParams) (*twilio.ChatService, error) {
	resp, err := c.Request.Post("/Services", params)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	cs := &twilio.ChatService{}
	json.NewDecoder(resp.Body).Decode(cs)
	return cs, err
}

// Read returns the details of a Service.
func (c Client) Read(sid string, params *twilio.ChatServiceParams) (*twilio.ChatService, error) {
	resp, err := c.Request.Get(fmt.Sprintf("/Services/%s", sid))
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	cs := &twilio.ChatService{}
	json.NewDecoder(resp.Body).Decode(cs)
	return cs, err
}

// Update updates a Service.
func (c Client) Update(sid string, params *twilio.ChatServiceParams) (*twilio.ChatService, error) {
	resp, err := c.Request.Post(fmt.Sprintf("/Services/%s", sid), params)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	cs := &twilio.ChatService{}
	json.NewDecoder(resp.Body).Decode(cs)
	return cs, err
}

// Delete deletes a Service.
func (c Client) Delete(sid string, params *twilio.ChatServiceParams) (*twilio.ChatService, error) {
	resp, err := c.Request.Delete(fmt.Sprintf("/Services/%s", sid))
	if err != nil {
		return nil, err
	}

	cs := &twilio.ChatService{}
	json.NewDecoder(resp.Body).Decode(cs)
	return cs, err
}
