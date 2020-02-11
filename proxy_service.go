package twilio

import (
	"encoding/json"
	"time"
)

// ProxyService is the top-level scope of all other resources in the Programmable Proxy REST API.
// All other Programmable Proxy resources belong to a specific Service.
// See: https://www.twilio.com/docs/proxy/api/service
type ProxyService struct {
	SID                     *string            `json:"sid"`
	AccountSID              *string            `json:"account_sid"`
	ChatInstanceSID         *string            `json:"chat_instance_sid"`
	UniqueName              *string            `json:"unique_name"`
	DefaultTTL              *int               `json:"default_ttl"`
	CallbackURL             *string            `json:"callback_url"`
	GeoMatchLevel           *string            `json:"geo_match_level"`
	NumberSelectionBehavior *string            `json:"number_selection_behavior"`
	InterceptCallbackURL    *string            `json:"intercept_callback_url"`
	OutOfSessionCallbackURL *string            `json:"out_of_session_callback_url"`
	DateCreated             *time.Time         `json:"date_created"`
	DateUpdated             *time.Time         `json:"date_updated"`
	URL                     *string            `json:"url"`
	Links                   map[string]*string `json:"links"`
}

// ProxyServiceParams is the set of parameters that can be used when creating or updating a service.
type ProxyServiceParams struct {
	ChatInstanceSID         *string `form:"ChatInstanceSid,omitempty"`
	UniqueName              *string `form:",omitempty"`
	DefaultTTL              *int    `form:"DefaultTtl,omitempty"`
	CallbackURL             *string `form:"CallbackUrl,omitempty"`
	GeoMatchLevel           *string `form:",omitempty"`
	NumberSelectionBehavior *string `form:",omitempty"`
	InterceptCallbackURL    *string `form:"InterceptCallbackUrl,omitempty"`
	OutOfSessionCallbackURL *string `form:"OutOfSessionCallbackUrl,omitempty"`
}

// ProxyServiceClient is the entrypoint for the Proxy Service API.
type ProxyServiceClient struct {
	client  *Twilio
	baseURL string
}

// NewProxyServiceClient constructs a new ProxyService Client.
func NewProxyServiceClient(client *Twilio) *ProxyServiceClient {
	c := new(ProxyServiceClient)
	c.client = client
	c.baseURL = "https://proxy.twilio.com/v1"

	return c
}

// Create creates a new ProxyService.
func (c *ProxyServiceClient) Create(params *ProxyServiceParams) (*ProxyService, error) {
	resp, err := c.client.Post(c.url("/Services"), params)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	p := &ProxyService{}
	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		return nil, err
	}

	return p, err
}

// Read returns the details of a ProxyService.
func (c *ProxyServiceClient) Read(sid string, params *ProxyServiceParams) (*ProxyService, error) {
	resp, err := c.client.Get(c.url("/Services/"+sid), nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	p := &ProxyService{}
	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		return nil, err
	}

	return p, err
}

// Update updates a ProxyService.
func (c *ProxyServiceClient) Update(sid string, params *ProxyServiceParams) (*ProxyService, error) {
	resp, err := c.client.Post(c.url("/Services/"+sid), params)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	p := &ProxyService{}
	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		return nil, err
	}

	return p, err
}

// Delete deletes a ProxyService.
func (c *ProxyServiceClient) Delete(sid string, params *ProxyServiceParams) error {
	resp, err := c.client.Delete(c.url("/Services/" + sid))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ProxyServiceClient) url(path string) string {
	if c.client.defaultbaseURL != nil {
		return *c.client.defaultbaseURL + path
	}

	return c.baseURL + path
}
