package chat

import (
	"encoding/json"
	"fmt"
	"log"

	twilio "github.com/twilio/twilio-go"
)

type Client struct {
	Request twilio.Request
}

// Create creates a new Service.
func (c Client) Create(params *twilio.ChatServiceParams) (*twilio.ChatService, error) {
	resp, err := c.Request.Post("/Services", params)
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
	resp, err := c.Request.Get(fmt.Sprintf("/Services/%s", sid))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cs := &twilio.ChatService{}
	if decodeErr := json.NewDecoder(resp.Body).Decode(cs); decodeErr != nil {
		log.Fatal(decodeErr)
		return nil, decodeErr
	}

	return cs, err
}

// Update updates a Service.
func (c Client) Update(sid string, params *twilio.ChatServiceParams) (*twilio.ChatService, error) {
	resp, err := c.Request.Post(fmt.Sprintf("/Services/%s", sid), params)
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
	resp, err := c.Request.Delete(fmt.Sprintf("/Services/%s", sid))
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
