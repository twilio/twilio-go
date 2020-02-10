package twilio

import (
	"encoding/json"
	"fmt"
	"time"
)

// StudioFlow  are individual workflows that you create. Flow definitions are expressed as instances of a JSON schema.
// See: https://www.twilio.com/docs/studio/rest-api/v2/flow
type StudioFlow struct {
	Sid           *string            `json:"sid"`
	AccountSID    *string            `json:"account_sid"`
	FriendlyName  *string            `json:"friendly_name"`
	Definition    *interface{}       `json:"definition"`
	Status        *string            `json:"status"`
	Revision      *int               `json:"revision"`
	CommitMessage *string            `json:"commit_message"`
	Errors        *interface{}       `json:"errors"`
	DateCreated   *time.Time         `json:"date_created"`
	DateUpdated   *time.Time         `json:"date_updated"`
	URL           *string            `json:"url"`
	Valid         *bool              `json:"valid"`
	WebhookURL    *string            `json:"webhook_url"`
	Links         map[string]*string `json:"links"`
}

// StudioFlowParams is the set of parameters that can be used when creating or updating a service.
type StudioFlowParams struct {
	FriendlyName  *string `form:",omitempty"`
	Status        *string `form:",omitempty"`
	Definition    *string `form:",omitempty"`
	CommitMessage *string `form:",omitempty"`
}

// StudioFlowClient is the entrypoint for the Proxy Service API.
type StudioFlowClient struct {
	serviceURL string
	client     *Twilio
}

// NewStudioFlowClient constructs a new StudioFlow Client.
func NewStudioFlowClient(client *Twilio) *StudioFlowClient {
	c := new(StudioFlowClient)
	c.client = client
	c.serviceURL = fmt.Sprintf("https://studio.%s/v2/Flows", c.client.BaseURL)

	return c
}

// Create creates a new StudioFlow.
func (c *StudioFlowClient) Create(params *StudioFlowParams) (*StudioFlow, error) {
	resp, err := c.client.Post(c.serviceURL, params)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioFlow{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Read returns the details of a StudioFlow.
func (c *StudioFlowClient) Read(sid string) (*StudioFlow, error) {
	resp, err := c.client.Get(fmt.Sprintf("%s/%s", c.serviceURL, sid), nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioFlow{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Update updates a StudioFlow.
func (c *StudioFlowClient) Update(sid string, params *StudioFlowParams) (*StudioFlow, error) {
	resp, err := c.client.Post(fmt.Sprintf("%s/%s", c.serviceURL, sid), params)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioFlow{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete deletes a StudioFlow.
func (c *StudioFlowClient) Delete(sid string) error {
	resp, err := c.client.Delete(fmt.Sprintf("%s/%s", c.serviceURL, sid))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}