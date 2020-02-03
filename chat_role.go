package twilio

import (
	"encoding/json"
	"fmt"
	"time"

	twilio "github.com/twilio/twilio-go/internal"
)

// ChatRole represents what a user can do within a Chat Service instance.
// See: https://www.twilio.com/docs/chat/rest/role-resource
type ChatRole struct {
	Sid          string    `json:"sid"`
	AccountSid   string    `json:"account_sid"`
	ServiceSid   string    `json:"service_sid"`
	FriendlyName string    `json:"friendly_name"`
	Type         string    `json:"type"`
	Permissions  []string  `json:"permissions"`
	DateCreated  time.Time `json:"date_created"`
	DateUpdated  time.Time `json:"date_updated"`
	URL          string    `json:"url"`
}

// ChatRoles represents a paginated set of Chat Role structs.
type ChatRoles struct {
	Meta  ChatRolesMeta `json:"meta"`
	Roles []*ChatRole
}

// ChatRolesMeta represents pagination metadata for ChatRoles.
type ChatRolesMeta struct {
	Page            int    `json:"page"`
	PageSize        int    `json:"page_size"`
	FirstPageURL    string `json:"first_page_url"`
	PreviousPageURL string `json:"previous_page_url"`
	URL             string `json:"url"`
	NextPageURL     string `json:"next_page_url"`
	Key             string `json:"key"`
}

// ChatRoleParams is the set of parameters that can be used when creating or updating a Chat Role.
type ChatRoleParams struct {
	FriendlyName string   `url:"FriendlyName,omitempty"`
	Type         string   `url:"Type,omitempty"`
	Permission   []string `url:"Permission,omitempty"`
}

// ChatRoleClient is the entrypoint for the Programmable Chat Role API.
type ChatRoleClient struct {
	serviceURL string
	client     *twilio.Client
}

// NewChatRoleClient constructs a new Chat Role client.
func NewChatRoleClient(request *twilio.Client) *ChatRoleClient {
	c := new(ChatRoleClient)
	c.client = request
	c.serviceURL = fmt.Sprintf("https://chat.%s/v2/Services", c.client.BaseURL)

	return c
}

// Create creates a new Chat Role attached to a particular service.
func (c *ChatRoleClient) Create(serviceSid string, params *ChatRoleParams) (*ChatRole, error) {
	resp, err := c.client.Post(fmt.Sprintf("%s/%s/Roles", c.serviceURL, serviceSid), params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cr := &ChatRole{}
	if err := json.NewDecoder(resp.Body).Decode(cr); err != nil {
		return nil, err
	}

	return cr, err
}

// Fetch returns the details of a Chat Role attached to a particular Service.
func (c *ChatRoleClient) Fetch(serviceSid, sid string) (*ChatRole, error) {
	resp, err := c.client.Get(fmt.Sprintf("%s/%s/Roles/%s", c.serviceURL, serviceSid, sid), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cr := &ChatRole{}
	if err := json.NewDecoder(resp.Body).Decode(cr); err != nil {
		return nil, err
	}

	return cr, err
}

// Read returns a paginated list of the roles attached to a particular Service.
func (c *ChatRoleClient) Read(serviceSid string) (*ChatRoles, error) {
	resp, err := c.client.Get(fmt.Sprintf("%s/%s/Roles", c.serviceURL, serviceSid), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cr := &ChatRoles{}
	if err := json.NewDecoder(resp.Body).Decode(cr); err != nil {
		return nil, err
	}

	return cr, err
}

// Update updates a Chat Role attached to a particular Service.
func (c *ChatRoleClient) Update(serviceSid, sid string, params *ChatRoleParams) (*ChatRole, error) {
	resp, err := c.client.Post(fmt.Sprintf("%s/%s/Roles/%s", c.serviceURL, serviceSid, sid), params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cr := &ChatRole{}
	if err := json.NewDecoder(resp.Body).Decode(cr); err != nil {
		return nil, err
	}

	return cr, err
}

// Delete deletes a Chat Role attached to a particular Service.
func (c *ChatRoleClient) Delete(serviceSid, sid string) error {
	resp, err := c.client.Delete(fmt.Sprintf("%s/%s/Roles/%s", c.serviceURL, serviceSid, sid))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return err
}
