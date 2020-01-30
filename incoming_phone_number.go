package twilio

import (
	"encoding/json"
	"fmt"

	twilio "github.com/twilio/twilio-go/internal"
)

// IncomingPhoneNumber represents a Twilio phone number provisioned from Twilio, ported or hosted to Twilio.
// See: https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource
type IncomingPhoneNumber struct {
	Beta                 bool            `json:"beta"`
	VoiceCallerIDLookup  bool            `json:"voice_caller_id_lookup"`
	AccountSid           string          `json:"account_sid"`
	AddressSid           string          `json:"address_sid"`
	AddressRequirements  string          `json:"address_requirements"`
	APIVersion           string          `json:"api_version"`
	Capabilities         map[string]bool `json:"capabilities"`
	DateCreated          string          `json:"date_created"`
	DateUpdated          string          `json:"date_updated"`
	FriendlyName         string          `json:"friendly_name"`
	IdentitySid          string          `json:"identity_sid"`
	PhoneNumber          string          `json:"phone_number"`
	Origin               string          `json:"origin"`
	Sid                  string          `json:"sid"`
	SMSApplicationSid    string          `json:"sms_application_sid"`
	SMSFallbackMethod    string          `json:"sms_fallback_method"`
	SMSFallbackURL       string          `json:"sms_fallback_url"`
	SMSMethod            string          `json:"sms_method"`
	SMSURL               string          `json:"sms_url"`
	StatusCallback       string          `json:"status_callback"`
	StatusCallbackMethod string          `json:"status_callback_method"`
	TrunkSid             string          `json:"trunk_sid"`
	URI                  string          `json:"uri"`
	VoiceApplicationSid  string          `json:"voice_application_sid"`
	VoiceFallbackMethod  string          `json:"voice_fallback_method"`
	VoiceFallbackURL     string          `json:"voice_fallback_url"`
	VoiceMethod          string          `json:"voice_method"`
	VoiceURL             string          `json:"voice_url"`
	EmergencyStatus      string          `json:"emergency_status"`
	EmergencyAddressSid  string          `json:"emergency_address_sid"`
	BundleSid            string          `json:"bundle_sid"`
}

// IncomingPhoneNumberParams is the set of parameters that can
// be used when creating or updating an Incoming Phone Number.
type IncomingPhoneNumberParams struct {
	AccountSid           string `url:"AccountSid,omitempty"`
	APIVersion           string `url:"ApiVersion,omitempty"`
	FriendlyName         string `url:"FriendlyName,omitempty"`
	SMSApplicationSid    string `url:"SmsApplicationSid,omitempty"`
	SMSFallbackMethod    string `url:"SmsApplicationMethod,omitempty"`
	PhoneNumber          string `url:"PhoneNumber,omitempty"`
	AreaCode             string `url:"AreaCode,omitempty"`
	SMSFallbackURL       string `url:"SmsFallbackUrl,omitempty"`
	SMSMethod            string `url:"SmsMethod,omitempty"`
	SMSURL               string `url:"SmsUrl,omitempty"`
	StatusCallback       string `url:"StatusCallback,omitempty"`
	AddressSid           string `url:"AddressSid,omitempty"`
	StatusCallbackMethod string `url:"StatusCallbackMethod,omitempty"`
	VoiceApplicationSid  string `url:"VoiceApplicationSid,omitempty"`
	VoiceCallerIDLookup  bool   `url:"VoiceCallerIdLookup,omitempty"`
	VoiceFallbackMethod  string `url:"VoiceFallbackMethod,omitempty"`
	VoiceFallbackURL     string `url:"VoiceFallbackUrl,omitempty"`
	VoiceMethod          string `url:"VoiceMethod,omitempty"`
	VoiceURL             string `url:"VoiceUrl,omitempty"`
	VoiceReceiveMode     string `url:"VoiceReceiveMode,omitempty"`
	EmergencyStatus      string `url:"EmergencyStatus,omitempty"`
	EmergencyAddressSid  string `url:"EmergencyAddressSid,omitempty"`
	BundleSid            string `url:"BundleSid,omitempty"`
	IdentitySid          string `url:"IdentitySid,omitempty"`
	TrunkSid             string `url:"TrunkSid,omitempty"`
}

// IncomingPhoneNumberClient is the entrypoint for the Incoming Phone Number resource.
// See: https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource
type IncomingPhoneNumberClient struct {
	client     *twilio.Client
	serviceURL string
}

// NewIncomingPhoneNumberClient constructs a new IncomingPhoneNumber client.
func NewIncomingPhoneNumberClient(client *twilio.Client) *IncomingPhoneNumberClient {
	pn := new(IncomingPhoneNumberClient)
	pn.client = client
	pn.serviceURL = fmt.Sprintf("https://api.%s/2010-04-01", pn.client.BaseURL)

	return pn
}

// Create creates a new IncomingPhoneNumber.
func (c IncomingPhoneNumberClient) Create(params *IncomingPhoneNumberParams) (*IncomingPhoneNumber, error) {
	uri := fmt.Sprintf("%s/Accounts/%s/IncomingPhoneNumbers.json", c.serviceURL, c.client.AccountSid)

	resp, err := c.client.Post(uri, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ipn := &IncomingPhoneNumber{}
	if decodeErr := json.NewDecoder(resp.Body).Decode(ipn); decodeErr != nil {
		return nil, decodeErr
	}

	return ipn, err
}

// Read returns the details of an IncomingPhoneNumber.
func (c IncomingPhoneNumberClient) Read(sid string) (*IncomingPhoneNumber, error) {
	uri := fmt.Sprintf("%s/Accounts/%s/IncomingPhoneNumbers/%s.json", c.serviceURL, c.client.AccountSid, sid)

	resp, err := c.client.Get(uri, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ipn := &IncomingPhoneNumber{}
	if decodeErr := json.NewDecoder(resp.Body).Decode(ipn); decodeErr != nil {
		return nil, decodeErr
	}

	return ipn, err
}

// Update updates an IncomingPhoneNumber.
func (c IncomingPhoneNumberClient) Update(sid string, params *IncomingPhoneNumberParams) (*IncomingPhoneNumber, error) {
	uri := fmt.Sprintf("%s/Accounts/%s/IncomingPhoneNumbers/%s.json", c.serviceURL, c.client.AccountSid, sid)

	resp, err := c.client.Post(uri, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ipn := &IncomingPhoneNumber{}
	if decodeErr := json.NewDecoder(resp.Body).Decode(ipn); decodeErr != nil {
		return nil, decodeErr
	}

	return ipn, err
}

// Delete releases an existing IncomingPhoneNumber.
func (c IncomingPhoneNumberClient) Delete(sid string) error {
	uri := fmt.Sprintf("%s/Accounts/%s/IncomingPhoneNumbers/%s.json", c.serviceURL, c.client.AccountSid, sid)

	resp, err := c.client.Delete(uri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
