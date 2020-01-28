package twilio

import (
	"encoding/json"
	"fmt"
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

// CreateIncomingPhoneNumber creates a new IncomingPhoneNumber.
func (pn PhoneNumberClient) CreateIncomingPhoneNumber(params *IncomingPhoneNumberParams) (*IncomingPhoneNumber, error) {
	uri := fmt.Sprintf("%s/Accounts/%s/IncomingPhoneNumbers.json", pn.serviceURL, pn.client.AccountSid)

	resp, err := pn.client.Post(uri, params)
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

// ReadIncomingPhoneNumber returns the details of an IncomingPhoneNumber.
func (pn PhoneNumberClient) ReadIncomingPhoneNumber(sid string) (*IncomingPhoneNumber, error) {
	uri := fmt.Sprintf("%s/Accounts/%s/IncomingPhoneNumbers/%s.json", pn.serviceURL, pn.client.AccountSid, sid)

	resp, err := pn.client.Get(uri, nil)
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

// UpdateIncomingPhoneNumber updates an IncomingPhoneNumber.
func (pn PhoneNumberClient) UpdateIncomingPhoneNumber(sid string,
	data IncomingPhoneNumberParams) (*IncomingPhoneNumber, error) {
	uri := fmt.Sprintf("%s/Accounts/%s/IncomingPhoneNumbers/%s.json", pn.serviceURL, pn.client.AccountSid, sid)

	resp, err := pn.client.Post(uri, data)
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

// DeleteIncomingPhoneNumber releases an existing IncomingPhoneNumber.
func (pn PhoneNumberClient) DeleteIncomingPhoneNumber(sid string) error {
	uri := fmt.Sprintf("%s/Accounts/%s/IncomingPhoneNumbers/%s.json", pn.serviceURL, pn.client.AccountSid, sid)

	resp, err := pn.client.Delete(uri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
