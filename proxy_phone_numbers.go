package twilio

import (
	"encoding/json"
	"fmt"
	"time"
)

// ProxyPhoneNumber represents a Twilio phone number provisioned from Twilio, ported or hosted to Twilio.
// See: https://www.twilio.com/docs/proxy/api/phone-number
type ProxyPhoneNumber struct {
	SID             *string          `json:"sid"`
	AccountSID      *string          `json:"account_sid"`
	ProxyServiceSID *string          `json:"service_sid"`
	DateCreated     *time.Time       `json:"date_created"`
	DateUpdated     *time.Time       `json:"date_updated"`
	PhoneNumber     *string          `json:"phone_number"`
	FriendlyName    *string          `json:"friendly_name"`
	ISOCountry      *string          `json:"iso_country"`
	Capabilities    map[string]*bool `json:"capabilities"`
	URL             *string          `json:"url"`
	IsReserved      *bool            `json:"is_reserved"`
	InUse           *int             `json:"in_use"`
}

// ProxyPhoneNumberList is the API response for reading multiple Proxy Phone Numbers
type ProxyPhoneNumberList struct {
	PhoneNumbers []*ProxyPhoneNumber `json:"phone_numbers"`
	Meta         *Meta               `json:"meta"`
}

// ProxyPhoneNumberUpdateParams is the set of parameters that can
// be used when updating updating a Proxy Phone Number.
type ProxyPhoneNumberUpdateParams struct {
	PhoneNumberSID *string `url:"Sid,omitempty"`
	IsReserved     *bool   `url:"IsReserved,omitempty"`
}

// ProxyPhoneNumberCreateParams is the set of parameters that can
// be used when creating a Proxy Phone Number.
type ProxyPhoneNumberCreateParams struct {
	PhoneNumberSID *string `form:"Sid,omitempty"`
	IsReserved     *bool   `form:"IsReserved,omitempty"`
	PhoneNumber    *string `form:"PhoneNumber,omitempty"`
}

// proxyPhoneNumberClient is the entrypoint for the Proxy Phone Number resource.
// See: https://www.twilio.com/docs/proxy/api/phone-number
type proxyPhoneNumberClient struct {
	client  *Twilio
	baseURL string
}

// newProxyPhoneNumberClient constructs a new ProxyPhoneNumber client.
func newProxyPhoneNumberClient(client *Twilio) *proxyPhoneNumberClient {
	c := new(proxyPhoneNumberClient)
	c.client = client
	c.baseURL = fmt.Sprintf("https://proxy.%s/v1", c.client.BaseURL)

	return c
}

// Create creates a new ProxyPhoneNumber.
func (c *proxyPhoneNumberClient) Create(
	proxyServiceSID string,
	params *ProxyPhoneNumberCreateParams,
) (*ProxyPhoneNumber, error) {
	uri := c.url(fmt.Sprintf("/Services/%s/PhoneNumbers", proxyServiceSID))
	resp, err := c.client.Post(uri, params)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	p := &ProxyPhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		return nil, err
	}

	return p, err
}

// Fetch returns the details of a ProxyPhoneNumber.
func (c *proxyPhoneNumberClient) Fetch(proxyServiceSID string, sid string) (*ProxyPhoneNumber, error) {
	uri := c.url(fmt.Sprintf("/Services/%s/PhoneNumbers/%s", proxyServiceSID, sid))
	resp, err := c.client.Get(uri, nil)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	p := &ProxyPhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		return nil, err
	}

	return p, err
}

// Read returns the details of a ProxyPhoneNumber.
func (c *proxyPhoneNumberClient) Read(proxyServiceSID string) (*ProxyPhoneNumberList, error) {
	uri := c.url(fmt.Sprintf("/Services/%s/PhoneNumbers", proxyServiceSID))
	resp, err := c.client.Get(uri, nil)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	p := &ProxyPhoneNumberList{}
	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		return nil, err
	}

	return p, err
}

// Update updates a ProxyPhoneNumber.
func (c *proxyPhoneNumberClient) Update(
	proxyServiceSID string,
	sid string,
	params *ProxyPhoneNumberUpdateParams,
) (*ProxyPhoneNumber, error) {
	uri := c.url(fmt.Sprintf("/Services/%s/PhoneNumbers/%s", proxyServiceSID, sid))
	resp, err := c.client.Post(uri, params)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	p := &ProxyPhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		return nil, err
	}

	return p, err
}

// Delete releases an existing ProxyPhoneNumber.
func (c *proxyPhoneNumberClient) Delete(proxyServiceSID string, sid string) error {
	uri := c.url(fmt.Sprintf("/Services/%s/PhoneNumbers/%s", proxyServiceSID, sid))
	resp, err := c.client.Delete(uri)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return err
}

func (c *proxyPhoneNumberClient) url(path string) string {
	if c.client.defaultbaseURL != nil {
		return *c.client.defaultbaseURL + path
	}

	return "https://proxy.twilio.com/v1" + path
}
