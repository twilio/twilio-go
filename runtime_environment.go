package twilio

import (
	"encoding/json"
	"fmt"
	"time"
)

// RuntimeEnvironment defines the different domains your Functions and Assets are available under.
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

// runtimeEnvironmentClient is the entrypoint for the Runtime Environment resource.
// See: https://www.twilio.com/docs/runtime/functions-assets-api/api/environment
type runtimeEnvironmentClient struct {
	client  *Twilio
	baseURL string
}

// newRuntimeEnvironmentClient constructs a new RuntimeEnvironment client.
func newRuntimeEnvironmentClient(client *Twilio) *runtimeEnvironmentClient {
	c := new(runtimeEnvironmentClient)
	c.client = client
	c.baseURL = fmt.Sprintf("https://serverless.%s/v1", c.client.BaseURL)

	return c
}

// Create creates a new RuntimeEnvironment.
func (c *runtimeEnvironmentClient) Create(
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
func (c *runtimeEnvironmentClient) Fetch(serviceSID string, sid string) (*RuntimeEnvironment, error) {
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
func (c *runtimeEnvironmentClient) Read(serviceSID string) (*RuntimeEnvironmentList, error) {
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
func (c *runtimeEnvironmentClient) Delete(serviceSID string, sid string) error {
	uri := c.url(fmt.Sprintf("/Services/%s/Environments/%s", serviceSID, sid))
	resp, err := c.client.Delete(uri)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return err
}

func (c *runtimeEnvironmentClient) url(path string) string {
	if c.client.defaultbaseURL != nil {
		return *c.client.defaultbaseURL + path
	}

	return "https://serverless.twilio.com/v1" + path
}
