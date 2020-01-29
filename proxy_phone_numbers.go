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
	SID          string          `json:"sid"`
	AccountSID   string          `json:"account_sid"`
	ServiceSID   string          `json:"service_sid"`
	DateCreated  time.Time       `json:"date_created"`
	DateUpdated  time.Time       `json:"date_updated"`
	PhoneNumber  string          `json:"phone_number"`
	FriendlyName string          `json:"friendly_name"`
	ISOCountry   string          `json:"iso_country"`
	Capabilities map[string]bool `json:"capabilities"`
	URL          string          `json:"url"`
	IsReserved   bool            `json:"is_reserved"`
	InUse        int             `json:"in_use"`
}

// ProxyPhoneNumberUpdateParams is the set of parameters that can
// be used when updating updating an Proxy Phone Number.

type ProxyPhoneNumberUpdateParams struct {
	ServiceSID *string `url:"-"`
	SID        *string `url:"Sid"`
	IsReserved *bool   `url:"IsReserved"`
}

// ProxyPhoneNumberCreateParams is the set of parameters that can
// be used when creating an Proxy Phone Number.
type ProxyPhoneNumberCreateParams struct {
	ProxyPhoneNumberUpdateParams
	PhoneNumber *string `url:"PhoneNumber"`
}

// ProxyPhoneNumberClient is the entrypoint for the Proxy Phone Number resource.
// See: https://www.twilio.com/docs/phone-numbers/api/ProxyPhoneNumber-resource
type ProxyPhoneNumberClient struct {
	client     *twilio.Client
	serviceURL string
}

// NewProxyPhoneNumberClient constructs a new ProxyPhoneNumber client.
func NewProxyPhoneNumberClient(client *twilio.Client) *ProxyPhoneNumberClient {
	pn := new(ProxyPhoneNumberClient)
	pn.client = client
	pn.serviceURL = fmt.Sprintf("https://proxy.%s/v1/Services", pn.client.BaseURL)

	return pn
}

// Create creates a new ProxyPhoneNumber.
func (c ProxyPhoneNumberClient) Create(params *ProxyPhoneNumberCreateParams) (*ProxyPhoneNumber, error) {
	resp, err := c.client.Post(fmt.Sprintf("/%v", params.ServiceSID), params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ppn := &ProxyPhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ppn); err != nil {
		return nil, err
	}

	return ppn, err
}

// Read returns the details of an ProxyPhoneNumber.
func (c ProxyPhoneNumberClient) Read(serviceSid string, sid string) (*ProxyPhoneNumber, error) {
	resp, err := c.client.Get(fmt.Sprintf("/%s/%s", serviceSid, sid), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ppn := &ProxyPhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ppn); err != nil {
		return nil, err
	}

	return ppn, err
}

// Update updates an ProxyPhoneNumber.
func (c ProxyPhoneNumberClient) Update(sid string, data ProxyPhoneNumberUpdateParams) (*ProxyPhoneNumber, error) {
	resp, err := c.client.Post(fmt.Sprintf("/%v/%s", data.ServiceSID, sid), data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ppn := &ProxyPhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ppn); err != nil {
		return nil, err
	}

	return ppn, err
}

// Delete releases an existing ProxyPhoneNumber.
func (c ProxyPhoneNumberClient) Delete(serviceSid, string, sid string) error {
	resp, err := c.client.Delete(fmt.Sprintf("/%s/%s", serviceSid, sid))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
