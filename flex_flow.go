package twilio

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// A FlexFlow is the logic linking a Messaging Channel, like SMS, to Flex.
// See: https://www.twilio.com/docs/flex/flow
type FlexFlow struct {
	AccountSID      *string      `json:"account_sid"`
	DateCreated     *time.Time   `json:"date_created"`
	DateUpdated     *time.Time   `json:"date_updated"`
	SID             *string      `json:"sid"`
	FriendlyName    *string      `json:"friendly_name"`
	ChatServiceSID  *string      `json:"chat_service_sid"`
	ChannelType     *string      `json:"channel_type"`
	ContactIdentity *string      `json:"contact_identity"`
	Enabled         *bool        `json:"enabled"`
	IntegrationType *string      `json:"integration_type"`
	Integration     *Integration `json:"integration"`
	LongLived       *bool        `json:"long_lived"`
	JanitorEnabled  *bool        `json:"janitor_enabled"`
	URL             *string      `json:"url"`
}

// An Integration represents a connection to Studio flow, TaskRouter task, or external webhook.
type Integration struct {
	WorkspaceSID      *string `json:"workspace_sid" form:"WorkspaceSid,omitempty"`
	WorkflowSID       *string `json:"workflow_sid" form:"WorkflowSid,omitempty"`
	Channel           *string `json:"channel" form:",omitempty"`
	Timeout           *int    `json:"timeout" form:",omitempty"`
	FlowSID           *string `json:"flow_sid" form:"FlowSid,omitempty"`
	RetryCount        *int    `json:"retry_count" form:",omitempty"`
	Priority          *int    `json:"priority" form:",omitempty"`
	CreationOnMessage *bool   `json:"creation_on_message" form:",omitempty"`
}

// FlexFlowList is the API response for reading multiple Proxy Services.
type FlexFlowList struct {
	FlexFlows []*FlexFlow `json:"flex_flows"`
	Meta      *Meta       `json:"meta"`
}

// FlexFlowParams is the set of parameters that can be used when creating or updating a Flex Flow.
type FlexFlowParams struct {
	FriendlyName    *string      `form:",omitempty"`
	Priority        *int         `form:",omitempty"`
	ChatServiceSID  *string      `form:"ChatServiceSid,omitempty"`
	ChannelType     *string      `form:",omitempty"`
	ContactIdentity *string      `form:",omitempty"`
	Enabled         *bool        `form:",omitempty"`
	IntegrationType *string      `form:",omitempty"`
	Integration     *Integration `form:",omitempty"`
	LongLived       *bool        `form:",omitempty"`
	JanitorEnabled  *bool        `form:",omitempty"`
	URL             *string      `form:"Url,omitempty"`
}

// flexFlowClient is the entrypoint for the Flex Flow API.
type flexFlowClient struct {
	baseURL string
	client  *Twilio
}

// newFlexFlowClient constructs a new Flex Flow client.
func newFlexFlowClient(client *Twilio) *flexFlowClient {
	c := new(flexFlowClient)
	c.client = client
	c.baseURL = fmt.Sprintf("https://flex-api.%s/v1/FlexFlows", c.client.BaseURL)

	return c
}

// Create creates a new Flex Flow.
func (c *flexFlowClient) Create(params *FlexFlowParams) (*FlexFlow, error) {
	resp, err := c.client.Post(c.url("/"), params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	f := &FlexFlow{}
	if err := json.NewDecoder(resp.Body).Decode(f); err != nil {
		return nil, err
	}

	return f, err
}

// Fetch returns the details of a Flex Flow.
func (c *flexFlowClient) Fetch(sid string) (*FlexFlow, error) {
	resp, err := c.client.Get(c.url("/"+sid), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	f := &FlexFlow{}
	if err := json.NewDecoder(resp.Body).Decode(f); err != nil {
		return nil, err
	}

	return f, err
}

// Read returns a paginated list of Flex Flows.
func (c *flexFlowClient) Read(friendlyName string) (*FlexFlowList, error) {
	v := url.Values{"FriendlyName": []string{friendlyName}}

	resp, err := c.client.Get(c.url("/"), v)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	f := &FlexFlowList{}
	if err := json.NewDecoder(resp.Body).Decode(f); err != nil {
		return nil, err
	}

	return f, err
}

// Update updates a Flex Flow.
func (c *flexFlowClient) Update(sid string, params *FlexFlowParams) (*FlexFlow, error) {
	resp, err := c.client.Post(c.url("/"+sid), params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	f := &FlexFlow{}
	if err := json.NewDecoder(resp.Body).Decode(f); err != nil {
		return nil, err
	}

	return f, err
}

// Delete deletes a Flex Flow.
func (c *flexFlowClient) Delete(sid string) error {
	resp, err := c.client.Delete(c.url("/" + sid))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *flexFlowClient) url(path string) string {
	if c.client.defaultbaseURL != nil {
		return *c.client.defaultbaseURL + path
	}

	return c.baseURL + path
}
