package twilio

import (
	"encoding/json"
	"fmt"
	"time"

	twilio "github.com/twilio/twilio-go/internal"
)

// ProxyService is the top-level scope of all other resources in the Programmable Proxy REST API.
// All other Programmable Proxy resources belong to a specific Service.
// See: https://www.twilio.com/docs/proxy/api/service
type ProxyService struct {
	Sid                     string            `json:"sid"`
	AccountSID              string            `json:"account_sid"`
	ChatInstanceSID         string            `json:"chat_instance_sid"`
	UniqueName              string            `json:"unique_name"`
	DefaultTTL              int               `json:"default_ttl"`
	CallbackURL             string            `json:"callback_url"`
	GeoMatchLevel           string            `json:"geo_match_level"`
	NumberSelectionBehavior string            `json:"number_selection_behavior"`
	InterceptCallbackURL    string            `json:"intercept_callback_url"`
	OutOfSessionCallbackURL string            `json:"out_of_session_callback_url"`
	DateCreated             time.Time         `json:"date_created"`
	DateUpdated             time.Time         `json:"date_updated"`
	URL                     string            `json:"url"`
	Links                   map[string]string `json:"links"`
}

// ProxyServiceParams is the set of parameters that can be used when creating or updating a service.
type ProxyServiceParams struct {
	ChatInstanceSID         string `url:"ChatInstanceSid,omitempty"`
	UniqueName              string `url:"UniqueName,omitempty"`
	DefaultTTL              int    `url:"DefaultTtl,omitempty"`
	CallbackURL             string `url:"CallbackUrl,omitempty"`
	GeoMatchLevel           string `url:"GeoMatchLevel,omitempty"`
	NumberSelectionBehavior string `url:"NumberSelectionBehavior,omitempty"`
	InterceptCallbackURL    string `url:"InterceptCallbackURL,omitempty"`
	OutOfSessionCallbackURL string `url:"OutOfSessionCallbackUrl,omitempty"`
}

// ProxyServiceClient is the entrypoint for the Proxy Service API.
type ProxyServiceClient struct {
	serviceURL string
	client     *twilio.Client
}

// NewProxyServiceClient constructs a new ProxyService Client.
func NewProxyServiceClient(request *twilio.Client) *ProxyServiceClient {
	c := new(ProxyServiceClient)
	c.client = request
	c.serviceURL = fmt.Sprintf("https://proxy.%s/v1", c.client.BaseURL)

	return c
}

// Create creates a new ProxyService.
func (c *ProxyServiceClient) Create(params *ProxyServiceParams) (*ProxyService, error) {
	resp, err := c.client.Post(fmt.Sprintf("%s/Services", c.serviceURL), params)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyService{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Read returns the details of a ProxyService.
func (c *ProxyServiceClient) Read(sid string, params *ProxyServiceParams) (*ProxyService, error) {
	resp, err := c.client.Get(fmt.Sprintf("%s/Services/%s", c.serviceURL, sid))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyService{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Update updates a ProxyService.
func (c *ProxyServiceClient) Update(sid string, params *ProxyServiceParams) (*ProxyService, error) {
	resp, err := c.client.Post(fmt.Sprintf("%s/Services/%s", c.serviceURL, sid), params)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyService{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete deletes a ProxyService.
func (c *ProxyServiceClient) Delete(sid string, params *ProxyServiceParams) (*ProxyService, error) {
	resp, err := c.client.Delete(fmt.Sprintf("%s/Services/%s", c.serviceURL, sid))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyService{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
