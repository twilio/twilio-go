package chat

import (
	"encoding/json"
	"fmt"

	twilio "github.com/twilio/twilio-go"
)

// Client sends Chat Service based requests
type Client struct {
	Client twilio.Client
}

// Create creates a new Service.
func (c Client) Create(params *twilio.ChatServiceParams) (*twilio.ChatService, error) {
	resp, err := c.Client.Post("/Services", params)
	if err != nil {
		return nil, fmt.Errorf("error creating Chat Service: %s", err)
	}

	defer resp.Body.Close()

	cs := &twilio.ChatService{}
	if err := json.NewDecoder(resp.Body).Decode(cs); err != nil {
		return nil, err
	}

	return cs, err
}

// Read returns the details of a Service.
func (c Client) Read(sid string) (*twilio.ChatService, error) {
	resp, err := c.Client.Get(fmt.Sprintf("/Services/%s", sid))

	if err != nil {
		return nil, fmt.Errorf("error read Chat Service: %s", err)
	}

	defer resp.Body.Close()

	cs := &twilio.ChatService{}
	if err := json.NewDecoder(resp.Body).Decode(cs); err != nil {
		return nil, fmt.Errorf("error updating Chat Service: %s", err)
	}

	return cs, nil
}

// Update updates a Service.
func (c Client) Update(sid string, params *twilio.ChatServiceParams) (*twilio.ChatService, error) {
	resp, err := c.Client.Post(fmt.Sprintf("/Services/%s", sid), params)
	if err != nil {
		return nil, fmt.Errorf("error updating Chat Service: %s", err)
	}

	defer resp.Body.Close()

	cs := &twilio.ChatService{}
	if err := json.NewDecoder(resp.Body).Decode(cs); err != nil {
		return nil, fmt.Errorf("error decoding Chat Service JSON: %s", err)
	}

	return cs, nil
}

// Delete deletes a Service.
func (c Client) Delete(sid string) error {
	resp, err := c.Client.Delete(fmt.Sprintf("/Services/%s", sid))

	if err != nil {
		return fmt.Errorf("error deleting Chat Service: %s", err)
	}

	defer resp.Body.Close()

	return nil
}
