package twilio

import (
	"encoding/json"
	"fmt"
	"time"
)

// RuntimeService is the top-level scope of all other resources in the Programmable Runtime REST API.
// All other Programmable Runtime resources belong to a specific Service.
// See: https://www.twilio.com/docs/runtime/functions-assets-api/api/service
type RuntimeService struct {
	SID                *string            `json:"sid"`
	AccountSID         *string            `json:"account_sid"`
	UniqueName         *string            `json:"unique_name"`
	FriendlyName       *string            `json:"friendly_name"`
	IncludeCredentials *bool              `json:"include_credentials"`
	DateCreated        *time.Time         `json:"date_created"`
	DateUpdated        *time.Time         `json:"date_updated"`
	URL                *string            `json:"url"`
	Links              map[string]*string `json:"links"`
}

// RuntimeServiceList is the API response for reading multiple Runtime Services.
type RuntimeServiceList struct {
	Service []*RuntimeService `json:"services"`
	Meta    *Meta             `json:"meta"`
}

// RuntimeServiceParams is the set of parameters that can be used when creating or updating a service.
type RuntimeServiceParams struct {
	UniqueName         *string `form:",omitempty"`
	FriendlyName       *string `form:",omitempty"`
	IncludeCredentials *bool   `form:",omitempty"`
}

// RuntimeServiceClient is the entrypoint for the Runtime Service API.
type RuntimeServiceClient struct {
	client  *Twilio
	baseURL string
}

// newRuntimeServiceClient constructs a new RuntimeService Client.
func newRuntimeServiceClient(client *Twilio) *RuntimeServiceClient {
	c := new(RuntimeServiceClient)
	c.client = client
	c.baseURL = fmt.Sprintf("https://serverless.%s/v1", c.client.BaseURL)

	return c
}

// Create creates a new RuntimeService.
func (c *RuntimeServiceClient) Create(params *RuntimeServiceParams) (*RuntimeService, error) {
	resp, err := c.client.Post(c.url("/Services"), params)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	p := &RuntimeService{}
	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		return nil, err
	}

	return p, err
}

// Fetch returns the details of a RuntimeService.
func (c *RuntimeServiceClient) Fetch(sid string) (*RuntimeService, error) {
	resp, err := c.client.Get(c.url("/Services/"+sid), nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	p := &RuntimeService{}
	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		return nil, err
	}

	return p, err
}

// Read returns the details of multiple Runtime Service.
func (c *RuntimeServiceClient) Read(sid string, params *RuntimeServiceParams) (*RuntimeServiceList, error) {
	resp, err := c.client.Get(c.url("/Services"), nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	p := &RuntimeServiceList{}
	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		return nil, err
	}

	return p, err
}

// Update updates a RuntimeService.
func (c *RuntimeServiceClient) Update(sid string, params *RuntimeServiceParams) (*RuntimeService, error) {
	resp, err := c.client.Post(c.url("/Services/"+sid), params)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	p := &RuntimeService{}
	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		return nil, err
	}

	return p, err
}

// Delete deletes a RuntimeService.
func (c *RuntimeServiceClient) Delete(sid string) error {
	resp, err := c.client.Delete(c.url("/Services/" + sid))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *RuntimeServiceClient) url(path string) string {
	if c.client.defaultbaseURL != nil {
		return *c.client.defaultbaseURL + path
	}

	return c.baseURL + path
}
