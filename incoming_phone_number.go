package twilio

import (
	"encoding/json"
	"fmt"
)

// IncomingPhoneNumber represents a Twilio phone number provisioned from Twilio, ported or hosted to Twilio.
// See: https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource
type IncomingPhoneNumber struct {
	Beta                 *bool            `json:"beta"`
	VoiceCallerIDLookup  *bool            `json:"voice_caller_id_lookup"`
	AccountSID           *string          `json:"account_sid"`
	AddressSID           *string          `json:"address_sid"`
	AddressRequirements  *string          `json:"address_requirements"`
	APIVersion           *string          `json:"api_version"`
	Capabilities         map[string]*bool `json:"capabilities"`
	DateCreated          *string          `json:"date_created"`
	DateUpdated          *string          `json:"date_updated"`
	FriendlyName         *string          `json:"friendly_name"`
	IdentitySID          *string          `json:"identity_sid"`
	PhoneNumber          *string          `json:"phone_number"`
	Origin               *string          `json:"origin"`
	SID                  *string          `json:"sid"`
	SMSApplicationSID    *string          `json:"sms_application_sid"`
	SMSFallbackMethod    *string          `json:"sms_fallback_method"`
	SMSFallbackURL       *string          `json:"sms_fallback_url"`
	SMSMethod            *string          `json:"sms_method"`
	SMSURL               *string          `json:"sms_url"`
	StatusCallback       *string          `json:"status_callback"`
	StatusCallbackMethod *string          `json:"status_callback_method"`
	TrunkSID             *string          `json:"trunk_sid"`
	URI                  *string          `json:"uri"`
	VoiceApplicationSID  *string          `json:"voice_application_sid"`
	VoiceFallbackMethod  *string          `json:"voice_fallback_method"`
	VoiceFallbackURL     *string          `json:"voice_fallback_url"`
	VoiceMethod          *string          `json:"voice_method"`
	VoiceURL             *string          `json:"voice_url"`
	EmergencyStatus      *string          `json:"emergency_status"`
	EmergencyAddressSID  *string          `json:"emergency_address_sid"`
	BundleSID            *string          `json:"bundle_sid"`
}

// IncomingPhoneNumberParams is the set of parameters that can
// be used when creating or updating an Incoming Phone Number.
type IncomingPhoneNumberParams struct {
	APIVersion           *string `form:"ApiVersion,omitempty"`
	FriendlyName         *string `form:",omitempty"`
	SMSApplicationSID    *string `form:"SmsApplicationSid,omitempty"`
	SMSFallbackMethod    *string `form:"SmsFallbackMethod,omitempty"`
	PhoneNumber          *string `form:",omitempty"`
	AreaCode             *string `form:",omitempty"`
	SMSFallbackURL       *string `form:"SmsFallbackUrl,omitempty"`
	SMSMethod            *string `form:"SmsMethod,omitempty"`
	SMSURL               *string `form:"SmsUrl,omitempty"`
	StatusCallback       *string `form:",omitempty"`
	AddressSID           *string `form:"AddressSid,omitempty"`
	StatusCallbackMethod *string `form:",omitempty"`
	VoiceApplicationSID  *string `form:"VoiceApplicationSid,omitempty"`
	VoiceCallerIDLookup  *bool   `form:"VoiceCallerIdLookup,omitempty"`
	VoiceFallbackMethod  *string `form:",omitempty"`
	VoiceFallbackURL     *string `form:"VoiceFallbackUrl,omitempty"`
	VoiceMethod          *string `form:",omitempty"`
	VoiceURL             *string `form:"VoiceUrl,omitempty"`
	VoiceReceiveMode     *string `form:",omitempty"`
	EmergencyStatus      *string `form:",omitempty"`
	EmergencyAddressSID  *string `form:"EmergencyAddressSid,omitempty"`
	BundleSID            *string `form:"BundleSid,omitempty"`
	IdentitySID          *string `form:"IdentitySid,omitempty"`
	TrunkSID             *string `form:"TrunkSid,omitempty"`
}

// incomingPhoneNumberClient is the entrypoint for the Incoming Phone Number resource.
// See: https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource
type incomingPhoneNumberClient struct {
	client  *Twilio
	baseURL string
}

// newIncomingPhoneNumberClient constructs a new IncomingPhoneNumber client.
func newIncomingPhoneNumberClient(client *Twilio) *incomingPhoneNumberClient {
	pn := new(incomingPhoneNumberClient)
	pn.client = client
	pn.baseURL = fmt.Sprintf("https://api.%s/2010-04-01", pn.client.BaseURL)

	return pn
}

// Create creates a new IncomingPhoneNumber.
func (c *incomingPhoneNumberClient) Create(params *IncomingPhoneNumberParams) (*IncomingPhoneNumber, error) {
	uri := c.url(fmt.Sprintf("/Accounts/%s/IncomingPhoneNumbers.json", c.client.AccountSID))

	resp, err := c.client.Post(uri, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	p := &IncomingPhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		return nil, err
	}

	return p, err
}

// Read returns the details of an IncomingPhoneNumber.
func (c *incomingPhoneNumberClient) Read(sid string) (*IncomingPhoneNumber, error) {
	uri := c.url(fmt.Sprintf("/Accounts/%s/IncomingPhoneNumbers/%s.json", c.client.AccountSID, sid))

	resp, err := c.client.Get(uri, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	p := &IncomingPhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		return nil, err
	}

	return p, err
}

// Update updates an IncomingPhoneNumber.
func (c *incomingPhoneNumberClient) Update(
	sid string,
	params *IncomingPhoneNumberParams,
) (*IncomingPhoneNumber, error) {
	url := c.url(fmt.Sprintf("/Accounts/%s/IncomingPhoneNumbers/%s.json", c.client.AccountSID, sid))

	resp, err := c.client.Post(url, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	p := &IncomingPhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(p); err != nil {
		return nil, err
	}

	return p, err
}

// Delete releases an existing IncomingPhoneNumber.
func (c *incomingPhoneNumberClient) Delete(sid string) error {
	uri := c.url(fmt.Sprintf("/Accounts/%s/IncomingPhoneNumbers/%s.json", c.client.AccountSID, sid))

	resp, err := c.client.Delete(uri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}

func (c *incomingPhoneNumberClient) url(path string) string {
	if c.client.defaultbaseURL != nil {
		return *c.client.defaultbaseURL + path
	}

	return c.baseURL + path
}
