package twilio

import (
	"encoding/json"
	"fmt"
	"time"
)

// StudioFlow  are individual workflows that you create. Flow definitions are expressed as instances of a JSON schema.
// See: https://www.twilio.com/docs/studio/rest-api/v2/flow
type StudioFlow struct {
	SID           *string            `json:"sid"`
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

// studioFlowClient is the entrypoint for the Proxy Service API.
type studioFlowClient struct {
	baseURL string
	client  *Twilio
}

// newStudioFlowClient constructs a new StudioFlow Client.
func newStudioFlowClient(client *Twilio) *studioFlowClient {
	c := new(studioFlowClient)
	c.client = client
	c.baseURL = fmt.Sprintf("https://studio.%s/v2", c.client.BaseURL)

	return c
}

// Create creates a new StudioFlow.
func (c *studioFlowClient) Create(params *StudioFlowParams) (*StudioFlow, error) {
	resp, err := c.client.Post(c.url("/Flows"), params)

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
func (c *studioFlowClient) Read(sid string) (*StudioFlow, error) {
	resp, err := c.client.Get(c.url("/Flows/"+sid), nil)
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
func (c *studioFlowClient) Update(sid string, params *StudioFlowParams) (*StudioFlow, error) {
	resp, err := c.client.Post(c.url("/Flows/"+sid), params)
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
func (c *studioFlowClient) Delete(sid string) error {
	resp, err := c.client.Delete(c.url("/Flows/" + sid))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *studioFlowClient) url(path string) string {
	if c.client.defaultbaseURL != nil {
		return *c.client.defaultbaseURL + path
	}

	return c.baseURL + path
}
