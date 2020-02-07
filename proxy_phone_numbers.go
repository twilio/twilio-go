package twilio

import (
	"encoding/json"
	"fmt"
	"time"

	twilio "github.com/twilio/twilio-go/internal"
)

// ProxyPhoneNumber represents a Twilio phone number provisioned from Twilio, ported or hosted to Twilio.
// See: https://www.twilio.com/docs/proxy/api/phone-number
type ProxyPhoneNumber struct {
	SID             string          `json:"sid"`
	AccountSID      string          `json:"account_sid"`
	ProxyServiceSID string          `json:"service_sid"`
	DateCreated     time.Time       `json:"date_created"`
	DateUpdated     time.Time       `json:"date_updated"`
	PhoneNumber     string          `json:"phone_number"`
	FriendlyName    string          `json:"friendly_name"`
	ISOCountry      string          `json:"iso_country"`
	Capabilities    map[string]bool `json:"capabilities"`
	URL             string          `json:"url"`
	IsReserved      bool            `json:"is_reserved"`
	InUse           int             `json:"in_use"`
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
	PhoneNumberSID *string `url:"Sid,omitempty"`
	IsReserved     *bool   `url:"IsReserved,omitempty"`
	PhoneNumber    *string `url:"PhoneNumber,omitempty"`
}

// ProxyPhoneNumberClient is the entrypoint for the Proxy Phone Number resource.
// See: https://www.twilio.com/docs/proxy/api/phone-number
type ProxyPhoneNumberClient struct {
	client     *twilio.Client
	serviceURL func(*pathParams) string
}

type pathParams struct {
	serviceSid string
	sid        string
}

// NewProxyPhoneNumberClient constructs a new ProxyPhoneNumber client.
func NewProxyPhoneNumberClient(client *twilio.Client) *ProxyPhoneNumberClient {
	pn := new(ProxyPhoneNumberClient)
	pn.client = client
	pn.serviceURL = func(p *pathParams) string {
		return fmt.Sprintf("https://proxy.%s/v1/Services/%s/PhoneNumbers/%s", pn.client.BaseURL, p.serviceSid, p.sid)
	}

	return pn
}

// Create creates a new ProxyPhoneNumber.
func (c ProxyPhoneNumberClient) Create(
	proxyServiceSID string,
	params *ProxyPhoneNumberCreateParams,
) (*ProxyPhoneNumber, error) {
	uri := c.serviceURL(&pathParams{serviceSid: proxyServiceSID})
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

// Read returns the details of a ProxyPhoneNumber.
func (c ProxyPhoneNumberClient) Read(proxyServiceSID string, sid string) (*ProxyPhoneNumber, error) {
	uri := c.serviceURL(&pathParams{proxyServiceSID, sid})
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

// Update updates a ProxyPhoneNumber.
func (c ProxyPhoneNumberClient) Update(
	proxyServiceSID string,
	sid string,
	params *ProxyPhoneNumberUpdateParams,
) (*ProxyPhoneNumber, error) {
	uri := c.serviceURL(&pathParams{proxyServiceSID, sid})
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
func (c ProxyPhoneNumberClient) Delete(proxyServiceSID string, sid string) error {
	uri := c.serviceURL(&pathParams{proxyServiceSID, sid})
	resp, err := c.client.Delete(uri)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return err
}
