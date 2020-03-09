package twilio

import (
	"encoding/json"
	"fmt"
	"time"
)

// ChatRole represents what a user can do within a Chat Service instance.
// See: https://www.twilio.com/docs/chat/rest/role-resource
type ChatRole struct {
	SID          *string    `json:"sid"`
	AccountSID   *string    `json:"account_sid"`
	ServiceSID   *string    `json:"service_sid"`
	FriendlyName *string    `json:"friendly_name"`
	Type         *string    `json:"type"`
	Permissions  []*string  `json:"permissions"`
	DateCreated  *time.Time `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated"`
	URL          *string    `json:"url"`
}

// ChatRoles represents a paginated set of Chat Role structs.
type ChatRoles struct {
	Meta  *Meta `json:"meta"`
	Roles []*ChatRole
}

// ChatRoleParams is the set of parameters that can be used when creating or updating a Chat Role.
type ChatRoleParams struct {
	FriendlyName *string   `form:",omitempty"`
	Type         *string   `form:",omitempty"`
	Permission   []*string `form:",omitempty"`
}

// chatRoleClient is the entrypoint for the Programmable Chat Role API.
type chatRoleClient struct {
	baseURL string
	client  *Twilio
}

// newChatRoleClient constructs a new Chat Role client.
func newChatRoleClient(client *Twilio) *chatRoleClient {
	c := new(chatRoleClient)
	c.client = client
	c.baseURL = fmt.Sprintf("https://chat.%s/v2", c.client.BaseURL)

	return c
}

// Create creates a new Chat Role attached to a particular service.
func (c *chatRoleClient) Create(serviceSID string, params *ChatRoleParams) (*ChatRole, error) {
	path := fmt.Sprintf("/Services/%s/Roles", serviceSID)
	resp, err := c.client.Post(c.url(path), params)

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
func (c *chatRoleClient) Fetch(serviceSID, sid string) (*ChatRole, error) {
	path := fmt.Sprintf("/Services/%s/Roles/%s", serviceSID, sid)
	resp, err := c.client.Get(c.url(path), nil)

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
func (c *chatRoleClient) Read(serviceSID string) (*ChatRoles, error) {
	path := fmt.Sprintf("/Services/%s/Roles", serviceSID)
	resp, err := c.client.Get(c.url(path), nil)

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
func (c *chatRoleClient) Update(serviceSID, sid string, params *ChatRoleParams) (*ChatRole, error) {
	path := fmt.Sprintf("/Services/%s/Roles/%s", serviceSID, sid)
	resp, err := c.client.Post(c.url(path), params)

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
func (c *chatRoleClient) Delete(serviceSID, sid string) error {
	path := fmt.Sprintf("/Services/%s/Roles/%s", serviceSID, sid)
	resp, err := c.client.Delete(c.url(path))

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return err
}

func (c *chatRoleClient) url(path string) string {
	if c.client.defaultbaseURL != nil {
		return *c.client.defaultbaseURL + path
	}

	return c.baseURL + path
}
