package chat

import (
	"encoding/json"
	"fmt"

	twilio "github.com/twilio/twilio-go"
)

// Client twilio chat client wrapper
type Client struct {
	Client twilio.Client
}

// Create creates a new Service.
func (c Client) Create(params *twilio.ChatServiceParams) (*twilio.ChatService, error) {
	resp, err := c.Client.Post("/Services", params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cs := &twilio.ChatService{}
	if decodeErr := json.NewDecoder(resp.Body).Decode(cs); decodeErr != nil {
		return nil, decodeErr
	}

	return cs, err
}

// Read returns the details of a Service.
func (c Client) Read(sid string, params *twilio.ChatServiceParams) (*twilio.ChatService, error) {
	resp, err := c.Client.Get(fmt.Sprintf("/Services/%s", sid))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cs := &twilio.ChatService{}
	if decodeErr := json.NewDecoder(resp.Body).Decode(cs); decodeErr != nil {
		return nil, decodeErr
	}

	return cs, err
}

// Update updates a Service.
func (c Client) Update(sid string, params *twilio.ChatServiceParams) (*twilio.ChatService, error) {
	resp, err := c.Client.Post(fmt.Sprintf("/Services/%s", sid), params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cs := &twilio.ChatService{}
	if decodeErr := json.NewDecoder(resp.Body).Decode(cs); decodeErr != nil {
		return nil, decodeErr
	}

	return cs, err
}

// Delete deletes a Service.
func (c Client) Delete(sid string, params *twilio.ChatServiceParams) (*twilio.ChatService, error) {
	resp, err := c.Client.Delete(fmt.Sprintf("/Services/%s", sid))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cs := &twilio.ChatService{}
	if decodeErr := json.NewDecoder(resp.Body).Decode(cs); decodeErr != nil {
		return nil, decodeErr
	}

	return cs, err
}
