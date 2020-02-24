package twilio

import (
	"encoding/json"
	"fmt"
	"time"
)

// RuntimeEnvironment represents a Twilio phone number provisioned from Twilio, ported or hosted to Twilio.
// See: https://www.twilio.com/docs/runtime/functions-assets-api/api/environment
type RuntimeEnvironment struct {
	SID          *string            `json:"sid"`
	AccountSID   *string            `json:"account_sid"`
	ServiceSID   *string            `json:"service_sid"`
	BuildSID     *string            `json:"build_sid"`
	UniqueName   *string            `json:"unique_name"`
	DomainName   *string            `json:"domain_name"`
	DomainSuffix *string            `json:"domain_suffix"`
	DateCreated  *time.Time         `json:"date_created"`
	DateUpdated  *time.Time         `json:"date_updated"`
	URL          *string            `json:"url"`
	Links        map[string]*string `json:"links"`
}

// RuntimeEnvironmentList is the API response for reading multiple Runtime Environments.
type RuntimeEnvironmentList struct {
	Environments []*RuntimeEnvironment `json:"environments"`
	Meta         *Meta                 `json:"meta"`
}

// RuntimeEnvironmentParams is the set of parameters that can
// be used when creating a Runtime Environment.
type RuntimeEnvironmentParams struct {
	UniqueName   *string `form:",omitempty"`
	DomainSuffix *string `form:",omitempty"`
}

// RuntimeEnvironmentClient is the entrypoint for the Runtime Environment resource.
// See: https://www.twilio.com/docs/runtime/functions-assets-api/api/environment
type RuntimeEnvironmentClient struct {
	client  *Twilio
	baseURL string
}

// NewRuntimeEnvironmentClient constructs a new RuntimeEnvironment client.
func NewRuntimeEnvironmentClient(client *Twilio) *RuntimeEnvironmentClient {
	c := new(RuntimeEnvironmentClient)
	c.client = client
	c.baseURL = fmt.Sprintf("https://serverless.%s/v1", c.client.BaseURL)

	return c
}

// Create creates a new RuntimeEnvironment.
func (c *RuntimeEnvironmentClient) Create(
	serviceSID string,
	params *RuntimeEnvironmentParams,
) (*RuntimeEnvironment, error) {
	uri := c.url(fmt.Sprintf("/Services/%s/Environments", serviceSID))
	resp, err := c.client.Post(uri, params)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	p := &RuntimeEnvironment{}
	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		return nil, err
	}

	return p, err
}

// Fetch returns the details of a RuntimeEnvironment.
func (c *RuntimeEnvironmentClient) Fetch(serviceSID string, sid string) (*RuntimeEnvironment, error) {
	uri := c.url(fmt.Sprintf("/Services/%s/Environments/%s", serviceSID, sid))
	resp, err := c.client.Get(uri, nil)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	p := &RuntimeEnvironment{}
	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		return nil, err
	}

	return p, err
}

// Read returns the details of a RuntimeEnvironment.
func (c *RuntimeEnvironmentClient) Read(serviceSID string) (*RuntimeEnvironmentList, error) {
	uri := c.url(fmt.Sprintf("/Services/%s/Environments", serviceSID))
	resp, err := c.client.Get(uri, nil)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	p := &RuntimeEnvironmentList{}
	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		return nil, err
	}

	return p, err
}

// Delete releases an existing RuntimeEnvironment.
func (c *RuntimeEnvironmentClient) Delete(serviceSID string, sid string) error {
	uri := c.url(fmt.Sprintf("/Services/%s/Environments/%s", serviceSID, sid))
	resp, err := c.client.Delete(uri)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return err
}

func (c *RuntimeEnvironmentClient) url(path string) string {
	if c.client.defaultbaseURL != nil {
		return *c.client.defaultbaseURL + path
	}

	return "https://serverless.twilio.com/v1" + path
}
