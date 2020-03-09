package twilio

import (
	"encoding/json"
	"fmt"
	"time"
)

// SyncService is the top-level scope of all other resources in the Sync REST API.
// See: https://www.twilio.com/docs/sync/api/service
type SyncService struct {
	SID                           *string            `json:"sid"`
	AccountSID                    *string            `json:"account_sid"`
	UniqueName                    *string            `json:"unique_name"`
	FriendlyName                  *string            `json:"friendly_name"`
	DateCreated                   *time.Time         `json:"date_created"`
	DateUpdated                   *time.Time         `json:"date_updated"`
	URL                           *string            `json:"url"`
	Links                         map[string]*string `json:"links"`
	WebhookURL                    *string            `json:"webhook_url"`
	WebhooksFromRestEnabled       *bool              `json:"webhooks_from_rest_enabled"`
	ReachabilityWebhooksEnabled   *bool              `json:"reachability_webhooks_enabled"`
	ACLEnabled                    *bool              `json:"acl_enabled"`
	ReachabilityDebouncingEnabled *bool              `json:"reachability_debouncing_enabled"`
	ReachabilityDebouncingWindow  *int               `json:"reachability_debouncing_window"`
}

// SyncServiceList is the API response for reading multiple Sync Services.
type SyncServiceList struct {
	Service []*SyncService `json:"services"`
	Meta    *Meta          `json:"meta"`
}

// SyncServiceParams is the set of parameters that can be used when creating or updating a service.
type SyncServiceParams struct {
	FriendlyName                  *string `form:",omitempty"`
	WebhookURL                    *string `form:"WebhookUrl,omitempty"`
	ReachabilityWebhooksEnabled   *bool   `form:",omitempty"`
	ACLEnabled                    *bool   `form:"AclEnabled,omitempty"`
	ReachabilityDebouncingEnabled *bool   `form:",omitempty"`
	ReachabilityDebouncingWindow  *int    `form:",omitempty"`
	WebhooksFromRestEnabled       *bool   `form:",omitempty"`
}

// syncServiceClient is the entrypoint for the Sync Service API.
type syncServiceClient struct {
	client  *Twilio
	baseURL string
}

// newSyncServiceClient constructs a new SyncService Client.
func newSyncServiceClient(client *Twilio) *syncServiceClient {
	c := new(syncServiceClient)
	c.client = client
	c.baseURL = fmt.Sprintf("https://sync.%s/v1", c.client.BaseURL)

	return c
}

// Create creates a new SyncService.
func (c *syncServiceClient) Create(params *SyncServiceParams) (*SyncService, error) {
	resp, err := c.client.Post(c.url("/Services"), params)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	p := &SyncService{}
	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		return nil, err
	}

	return p, err
}

// Fetch returns the details of a SyncService.
func (c *syncServiceClient) Fetch(sid string) (*SyncService, error) {
	resp, err := c.client.Get(c.url("/Services/"+sid), nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	p := &SyncService{}
	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		return nil, err
	}

	return p, err
}

// Read returns the details of multiple SyncServices.
func (c *syncServiceClient) Read(sid string, params *SyncServiceParams) (*SyncServiceList, error) {
	resp, err := c.client.Get(c.url("/Services"), nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	p := &SyncServiceList{}
	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		return nil, err
	}

	return p, err
}

// Update updates a SyncService.
func (c *syncServiceClient) Update(sid string, params *SyncServiceParams) (*SyncService, error) {
	resp, err := c.client.Post(c.url("/Services/"+sid), params)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	p := &SyncService{}
	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		return nil, err
	}

	return p, err
}

// Delete deletes a SyncService.
func (c *syncServiceClient) Delete(sid string) error {
	resp, err := c.client.Delete(c.url("/Services/" + sid))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *syncServiceClient) url(path string) string {
	if c.client.defaultbaseURL != nil {
		return *c.client.defaultbaseURL + path
	}

	return c.baseURL + path
}
