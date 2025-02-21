/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateIncomingPhoneNumber'
type CreateIncomingPhoneNumberParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The API version to use for incoming calls made to the new phone number. The default is `2010-04-01`.
	ApiVersion *string `json:"ApiVersion,omitempty"`
	// A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, this is a formatted version of the new phone number.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The SID of the application that should handle SMS messages sent to the new phone number. If an `sms_application_sid` is present, we ignore all of the `sms_*_url` urls and use those set on the application.
	SmsApplicationSid *string `json:"SmsApplicationSid,omitempty"`
	// The HTTP method that we should use to call `sms_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`.
	SmsFallbackMethod *string `json:"SmsFallbackMethod,omitempty"`
	// The URL that we should call when an error occurs while requesting or executing the TwiML defined by `sms_url`.
	SmsFallbackUrl *string `json:"SmsFallbackUrl,omitempty"`
	// The HTTP method that we should use to call `sms_url`. Can be: `GET` or `POST` and defaults to `POST`.
	SmsMethod *string `json:"SmsMethod,omitempty"`
	// The URL we should call when the new phone number receives an incoming SMS message.
	SmsUrl *string `json:"SmsUrl,omitempty"`
	// The URL we should call using the `status_callback_method` to send status information to your application.
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST` and defaults to `POST`.
	StatusCallbackMethod *string `json:"StatusCallbackMethod,omitempty"`
	// The SID of the application we should use to handle calls to the new phone number. If a `voice_application_sid` is present, we ignore all of the voice urls and use only those set on the application. Setting a `voice_application_sid` will automatically delete your `trunk_sid` and vice versa.
	VoiceApplicationSid *string `json:"VoiceApplicationSid,omitempty"`
	// Whether to lookup the caller's name from the CNAM database and post it to your app. Can be: `true` or `false` and defaults to `false`.
	VoiceCallerIdLookup *bool `json:"VoiceCallerIdLookup,omitempty"`
	// The HTTP method that we should use to call `voice_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`.
	VoiceFallbackMethod *string `json:"VoiceFallbackMethod,omitempty"`
	// The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`.
	VoiceFallbackUrl *string `json:"VoiceFallbackUrl,omitempty"`
	// The HTTP method that we should use to call `voice_url`. Can be: `GET` or `POST` and defaults to `POST`.
	VoiceMethod *string `json:"VoiceMethod,omitempty"`
	// The URL that we should call to answer a call to the new phone number. The `voice_url` will not be called if a `voice_application_sid` or a `trunk_sid` is set.
	VoiceUrl *string `json:"VoiceUrl,omitempty"`
	//
	EmergencyStatus *string `json:"EmergencyStatus,omitempty"`
	// The SID of the emergency address configuration to use for emergency calling from the new phone number.
	EmergencyAddressSid *string `json:"EmergencyAddressSid,omitempty"`
	// The SID of the Trunk we should use to handle calls to the new phone number. If a `trunk_sid` is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a `trunk_sid` will automatically delete your `voice_application_sid` and vice versa.
	TrunkSid *string `json:"TrunkSid,omitempty"`
	// The SID of the Identity resource that we should associate with the new phone number. Some regions require an identity to meet local regulations.
	IdentitySid *string `json:"IdentitySid,omitempty"`
	// The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations.
	AddressSid *string `json:"AddressSid,omitempty"`
	//
	VoiceReceiveMode *string `json:"VoiceReceiveMode,omitempty"`
	// The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations.
	BundleSid *string `json:"BundleSid,omitempty"`
	// The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234.
	PhoneNumber *string `json:"PhoneNumber,omitempty"`
	// The desired area code for your new incoming phone number. Can be any three-digit, US or Canada area code. We will provision an available phone number within this area code for you. **You must provide an `area_code` or a `phone_number`.** (US and Canada only).
	AreaCode *string `json:"AreaCode,omitempty"`
}

func (params *CreateIncomingPhoneNumberParams) SetPathAccountSid(PathAccountSid string) *CreateIncomingPhoneNumberParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateIncomingPhoneNumberParams) SetApiVersion(ApiVersion string) *CreateIncomingPhoneNumberParams {
	params.ApiVersion = &ApiVersion
	return params
}
func (params *CreateIncomingPhoneNumberParams) SetFriendlyName(FriendlyName string) *CreateIncomingPhoneNumberParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateIncomingPhoneNumberParams) SetSmsApplicationSid(SmsApplicationSid string) *CreateIncomingPhoneNumberParams {
	params.SmsApplicationSid = &SmsApplicationSid
	return params
}
func (params *CreateIncomingPhoneNumberParams) SetSmsFallbackMethod(SmsFallbackMethod string) *CreateIncomingPhoneNumberParams {
	params.SmsFallbackMethod = &SmsFallbackMethod
	return params
}
func (params *CreateIncomingPhoneNumberParams) SetSmsFallbackUrl(SmsFallbackUrl string) *CreateIncomingPhoneNumberParams {
	params.SmsFallbackUrl = &SmsFallbackUrl
	return params
}
func (params *CreateIncomingPhoneNumberParams) SetSmsMethod(SmsMethod string) *CreateIncomingPhoneNumberParams {
	params.SmsMethod = &SmsMethod
	return params
}
func (params *CreateIncomingPhoneNumberParams) SetSmsUrl(SmsUrl string) *CreateIncomingPhoneNumberParams {
	params.SmsUrl = &SmsUrl
	return params
}
func (params *CreateIncomingPhoneNumberParams) SetStatusCallback(StatusCallback string) *CreateIncomingPhoneNumberParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *CreateIncomingPhoneNumberParams) SetStatusCallbackMethod(StatusCallbackMethod string) *CreateIncomingPhoneNumberParams {
	params.StatusCallbackMethod = &StatusCallbackMethod
	return params
}
func (params *CreateIncomingPhoneNumberParams) SetVoiceApplicationSid(VoiceApplicationSid string) *CreateIncomingPhoneNumberParams {
	params.VoiceApplicationSid = &VoiceApplicationSid
	return params
}
func (params *CreateIncomingPhoneNumberParams) SetVoiceCallerIdLookup(VoiceCallerIdLookup bool) *CreateIncomingPhoneNumberParams {
	params.VoiceCallerIdLookup = &VoiceCallerIdLookup
	return params
}
func (params *CreateIncomingPhoneNumberParams) SetVoiceFallbackMethod(VoiceFallbackMethod string) *CreateIncomingPhoneNumberParams {
	params.VoiceFallbackMethod = &VoiceFallbackMethod
	return params
}
func (params *CreateIncomingPhoneNumberParams) SetVoiceFallbackUrl(VoiceFallbackUrl string) *CreateIncomingPhoneNumberParams {
	params.VoiceFallbackUrl = &VoiceFallbackUrl
	return params
}
func (params *CreateIncomingPhoneNumberParams) SetVoiceMethod(VoiceMethod string) *CreateIncomingPhoneNumberParams {
	params.VoiceMethod = &VoiceMethod
	return params
}
func (params *CreateIncomingPhoneNumberParams) SetVoiceUrl(VoiceUrl string) *CreateIncomingPhoneNumberParams {
	params.VoiceUrl = &VoiceUrl
	return params
}
func (params *CreateIncomingPhoneNumberParams) SetEmergencyStatus(EmergencyStatus string) *CreateIncomingPhoneNumberParams {
	params.EmergencyStatus = &EmergencyStatus
	return params
}
func (params *CreateIncomingPhoneNumberParams) SetEmergencyAddressSid(EmergencyAddressSid string) *CreateIncomingPhoneNumberParams {
	params.EmergencyAddressSid = &EmergencyAddressSid
	return params
}
func (params *CreateIncomingPhoneNumberParams) SetTrunkSid(TrunkSid string) *CreateIncomingPhoneNumberParams {
	params.TrunkSid = &TrunkSid
	return params
}
func (params *CreateIncomingPhoneNumberParams) SetIdentitySid(IdentitySid string) *CreateIncomingPhoneNumberParams {
	params.IdentitySid = &IdentitySid
	return params
}
func (params *CreateIncomingPhoneNumberParams) SetAddressSid(AddressSid string) *CreateIncomingPhoneNumberParams {
	params.AddressSid = &AddressSid
	return params
}
func (params *CreateIncomingPhoneNumberParams) SetVoiceReceiveMode(VoiceReceiveMode string) *CreateIncomingPhoneNumberParams {
	params.VoiceReceiveMode = &VoiceReceiveMode
	return params
}
func (params *CreateIncomingPhoneNumberParams) SetBundleSid(BundleSid string) *CreateIncomingPhoneNumberParams {
	params.BundleSid = &BundleSid
	return params
}
func (params *CreateIncomingPhoneNumberParams) SetPhoneNumber(PhoneNumber string) *CreateIncomingPhoneNumberParams {
	params.PhoneNumber = &PhoneNumber
	return params
}
func (params *CreateIncomingPhoneNumberParams) SetAreaCode(AreaCode string) *CreateIncomingPhoneNumberParams {
	params.AreaCode = &AreaCode
	return params
}

// Purchase a phone-number for the account.
func (c *ApiService) CreateIncomingPhoneNumber(params *CreateIncomingPhoneNumberParams) (*ApiV2010IncomingPhoneNumber, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.ApiVersion != nil {
		data.Set("ApiVersion", *params.ApiVersion)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.SmsApplicationSid != nil {
		data.Set("SmsApplicationSid", *params.SmsApplicationSid)
	}
	if params != nil && params.SmsFallbackMethod != nil {
		data.Set("SmsFallbackMethod", *params.SmsFallbackMethod)
	}
	if params != nil && params.SmsFallbackUrl != nil {
		data.Set("SmsFallbackUrl", *params.SmsFallbackUrl)
	}
	if params != nil && params.SmsMethod != nil {
		data.Set("SmsMethod", *params.SmsMethod)
	}
	if params != nil && params.SmsUrl != nil {
		data.Set("SmsUrl", *params.SmsUrl)
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.StatusCallbackMethod != nil {
		data.Set("StatusCallbackMethod", *params.StatusCallbackMethod)
	}
	if params != nil && params.VoiceApplicationSid != nil {
		data.Set("VoiceApplicationSid", *params.VoiceApplicationSid)
	}
	if params != nil && params.VoiceCallerIdLookup != nil {
		data.Set("VoiceCallerIdLookup", fmt.Sprint(*params.VoiceCallerIdLookup))
	}
	if params != nil && params.VoiceFallbackMethod != nil {
		data.Set("VoiceFallbackMethod", *params.VoiceFallbackMethod)
	}
	if params != nil && params.VoiceFallbackUrl != nil {
		data.Set("VoiceFallbackUrl", *params.VoiceFallbackUrl)
	}
	if params != nil && params.VoiceMethod != nil {
		data.Set("VoiceMethod", *params.VoiceMethod)
	}
	if params != nil && params.VoiceUrl != nil {
		data.Set("VoiceUrl", *params.VoiceUrl)
	}
	if params != nil && params.EmergencyStatus != nil {
		data.Set("EmergencyStatus", *params.EmergencyStatus)
	}
	if params != nil && params.EmergencyAddressSid != nil {
		data.Set("EmergencyAddressSid", *params.EmergencyAddressSid)
	}
	if params != nil && params.TrunkSid != nil {
		data.Set("TrunkSid", *params.TrunkSid)
	}
	if params != nil && params.IdentitySid != nil {
		data.Set("IdentitySid", *params.IdentitySid)
	}
	if params != nil && params.AddressSid != nil {
		data.Set("AddressSid", *params.AddressSid)
	}
	if params != nil && params.VoiceReceiveMode != nil {
		data.Set("VoiceReceiveMode", *params.VoiceReceiveMode)
	}
	if params != nil && params.BundleSid != nil {
		data.Set("BundleSid", *params.BundleSid)
	}
	if params != nil && params.PhoneNumber != nil {
		data.Set("PhoneNumber", *params.PhoneNumber)
	}
	if params != nil && params.AreaCode != nil {
		data.Set("AreaCode", *params.AreaCode)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010IncomingPhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteIncomingPhoneNumber'
type DeleteIncomingPhoneNumberParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resources to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteIncomingPhoneNumberParams) SetPathAccountSid(PathAccountSid string) *DeleteIncomingPhoneNumberParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete a phone-numbers belonging to the account used to make the request.
func (c *ApiService) DeleteIncomingPhoneNumber(Sid string, params *DeleteIncomingPhoneNumberParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Optional parameters for the method 'FetchIncomingPhoneNumber'
type FetchIncomingPhoneNumberParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchIncomingPhoneNumberParams) SetPathAccountSid(PathAccountSid string) *FetchIncomingPhoneNumberParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch an incoming-phone-number belonging to the account used to make the request.
func (c *ApiService) FetchIncomingPhoneNumber(Sid string, params *FetchIncomingPhoneNumberParams) (*ApiV2010IncomingPhoneNumber, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010IncomingPhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListIncomingPhoneNumber'
type ListIncomingPhoneNumberParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// Whether to include phone numbers new to the Twilio platform. Can be: `true` or `false` and the default is `true`.
	Beta *bool `json:"Beta,omitempty"`
	// A string that identifies the IncomingPhoneNumber resources to read.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The phone numbers of the IncomingPhoneNumber resources to read. You can specify partial numbers and use '*' as a wildcard for any digit.
	PhoneNumber *string `json:"PhoneNumber,omitempty"`
	// Whether to include phone numbers based on their origin. Can be: `twilio` or `hosted`. By default, phone numbers of all origin are included.
	Origin *string `json:"Origin,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int64 `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int64 `json:"limit,omitempty"`
}

func (params *ListIncomingPhoneNumberParams) SetPathAccountSid(PathAccountSid string) *ListIncomingPhoneNumberParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListIncomingPhoneNumberParams) SetBeta(Beta bool) *ListIncomingPhoneNumberParams {
	params.Beta = &Beta
	return params
}
func (params *ListIncomingPhoneNumberParams) SetFriendlyName(FriendlyName string) *ListIncomingPhoneNumberParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListIncomingPhoneNumberParams) SetPhoneNumber(PhoneNumber string) *ListIncomingPhoneNumberParams {
	params.PhoneNumber = &PhoneNumber
	return params
}
func (params *ListIncomingPhoneNumberParams) SetOrigin(Origin string) *ListIncomingPhoneNumberParams {
	params.Origin = &Origin
	return params
}
func (params *ListIncomingPhoneNumberParams) SetPageSize(PageSize int64) *ListIncomingPhoneNumberParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListIncomingPhoneNumberParams) SetLimit(Limit int64) *ListIncomingPhoneNumberParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of IncomingPhoneNumber records from the API. Request is executed immediately.
func (c *ApiService) PageIncomingPhoneNumber(params *ListIncomingPhoneNumberParams, pageToken, pageNumber string) (*ListIncomingPhoneNumberResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Beta != nil {
		data.Set("Beta", fmt.Sprint(*params.Beta))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.PhoneNumber != nil {
		data.Set("PhoneNumber", *params.PhoneNumber)
	}
	if params != nil && params.Origin != nil {
		data.Set("Origin", *params.Origin)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListIncomingPhoneNumberResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists IncomingPhoneNumber records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListIncomingPhoneNumber(params *ListIncomingPhoneNumberParams) ([]ApiV2010IncomingPhoneNumber, error) {
	response, errors := c.StreamIncomingPhoneNumber(params)

	records := make([]ApiV2010IncomingPhoneNumber, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams IncomingPhoneNumber records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamIncomingPhoneNumber(params *ListIncomingPhoneNumberParams) (chan ApiV2010IncomingPhoneNumber, chan error) {
	if params == nil {
		params = &ListIncomingPhoneNumberParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ApiV2010IncomingPhoneNumber, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageIncomingPhoneNumber(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamIncomingPhoneNumber(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamIncomingPhoneNumber(response *ListIncomingPhoneNumberResponse, params *ListIncomingPhoneNumberParams, recordChannel chan ApiV2010IncomingPhoneNumber, errorChannel chan error) {
	var curRecord int64 = 1

	for response != nil {
		responseRecords := response.IncomingPhoneNumbers
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListIncomingPhoneNumberResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListIncomingPhoneNumberResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListIncomingPhoneNumberResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListIncomingPhoneNumberResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateIncomingPhoneNumber'
type UpdateIncomingPhoneNumberParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resource to update.  For more information, see [Exchanging Numbers Between Subaccounts](https://www.twilio.com/docs/iam/api/subaccounts#exchanging-numbers).
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resource to update.  For more information, see [Exchanging Numbers Between Subaccounts](https://www.twilio.com/docs/iam/api/subaccounts#exchanging-numbers).
	AccountSid *string `json:"AccountSid,omitempty"`
	// The API version to use for incoming calls made to the phone number. The default is `2010-04-01`.
	ApiVersion *string `json:"ApiVersion,omitempty"`
	// A descriptive string that you created to describe this phone number. It can be up to 64 characters long. By default, this is a formatted version of the phone number.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The SID of the application that should handle SMS messages sent to the number. If an `sms_application_sid` is present, we ignore all of the `sms_*_url` urls and use those set on the application.
	SmsApplicationSid *string `json:"SmsApplicationSid,omitempty"`
	// The HTTP method that we should use to call `sms_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`.
	SmsFallbackMethod *string `json:"SmsFallbackMethod,omitempty"`
	// The URL that we should call when an error occurs while requesting or executing the TwiML defined by `sms_url`.
	SmsFallbackUrl *string `json:"SmsFallbackUrl,omitempty"`
	// The HTTP method that we should use to call `sms_url`. Can be: `GET` or `POST` and defaults to `POST`.
	SmsMethod *string `json:"SmsMethod,omitempty"`
	// The URL we should call when the phone number receives an incoming SMS message.
	SmsUrl *string `json:"SmsUrl,omitempty"`
	// The URL we should call using the `status_callback_method` to send status information to your application.
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST` and defaults to `POST`.
	StatusCallbackMethod *string `json:"StatusCallbackMethod,omitempty"`
	// The SID of the application we should use to handle phone calls to the phone number. If a `voice_application_sid` is present, we ignore all of the voice urls and use only those set on the application. Setting a `voice_application_sid` will automatically delete your `trunk_sid` and vice versa.
	VoiceApplicationSid *string `json:"VoiceApplicationSid,omitempty"`
	// Whether to lookup the caller's name from the CNAM database and post it to your app. Can be: `true` or `false` and defaults to `false`.
	VoiceCallerIdLookup *bool `json:"VoiceCallerIdLookup,omitempty"`
	// The HTTP method that we should use to call `voice_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`.
	VoiceFallbackMethod *string `json:"VoiceFallbackMethod,omitempty"`
	// The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`.
	VoiceFallbackUrl *string `json:"VoiceFallbackUrl,omitempty"`
	// The HTTP method that we should use to call `voice_url`. Can be: `GET` or `POST` and defaults to `POST`.
	VoiceMethod *string `json:"VoiceMethod,omitempty"`
	// The URL that we should call to answer a call to the phone number. The `voice_url` will not be called if a `voice_application_sid` or a `trunk_sid` is set.
	VoiceUrl *string `json:"VoiceUrl,omitempty"`
	//
	EmergencyStatus *string `json:"EmergencyStatus,omitempty"`
	// The SID of the emergency address configuration to use for emergency calling from this phone number.
	EmergencyAddressSid *string `json:"EmergencyAddressSid,omitempty"`
	// The SID of the Trunk we should use to handle phone calls to the phone number. If a `trunk_sid` is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a `trunk_sid` will automatically delete your `voice_application_sid` and vice versa.
	TrunkSid *string `json:"TrunkSid,omitempty"`
	//
	VoiceReceiveMode *string `json:"VoiceReceiveMode,omitempty"`
	// The SID of the Identity resource that we should associate with the phone number. Some regions require an identity to meet local regulations.
	IdentitySid *string `json:"IdentitySid,omitempty"`
	// The SID of the Address resource we should associate with the phone number. Some regions require addresses to meet local regulations.
	AddressSid *string `json:"AddressSid,omitempty"`
	// The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations.
	BundleSid *string `json:"BundleSid,omitempty"`
}

func (params *UpdateIncomingPhoneNumberParams) SetPathAccountSid(PathAccountSid string) *UpdateIncomingPhoneNumberParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateIncomingPhoneNumberParams) SetAccountSid(AccountSid string) *UpdateIncomingPhoneNumberParams {
	params.AccountSid = &AccountSid
	return params
}
func (params *UpdateIncomingPhoneNumberParams) SetApiVersion(ApiVersion string) *UpdateIncomingPhoneNumberParams {
	params.ApiVersion = &ApiVersion
	return params
}
func (params *UpdateIncomingPhoneNumberParams) SetFriendlyName(FriendlyName string) *UpdateIncomingPhoneNumberParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateIncomingPhoneNumberParams) SetSmsApplicationSid(SmsApplicationSid string) *UpdateIncomingPhoneNumberParams {
	params.SmsApplicationSid = &SmsApplicationSid
	return params
}
func (params *UpdateIncomingPhoneNumberParams) SetSmsFallbackMethod(SmsFallbackMethod string) *UpdateIncomingPhoneNumberParams {
	params.SmsFallbackMethod = &SmsFallbackMethod
	return params
}
func (params *UpdateIncomingPhoneNumberParams) SetSmsFallbackUrl(SmsFallbackUrl string) *UpdateIncomingPhoneNumberParams {
	params.SmsFallbackUrl = &SmsFallbackUrl
	return params
}
func (params *UpdateIncomingPhoneNumberParams) SetSmsMethod(SmsMethod string) *UpdateIncomingPhoneNumberParams {
	params.SmsMethod = &SmsMethod
	return params
}
func (params *UpdateIncomingPhoneNumberParams) SetSmsUrl(SmsUrl string) *UpdateIncomingPhoneNumberParams {
	params.SmsUrl = &SmsUrl
	return params
}
func (params *UpdateIncomingPhoneNumberParams) SetStatusCallback(StatusCallback string) *UpdateIncomingPhoneNumberParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *UpdateIncomingPhoneNumberParams) SetStatusCallbackMethod(StatusCallbackMethod string) *UpdateIncomingPhoneNumberParams {
	params.StatusCallbackMethod = &StatusCallbackMethod
	return params
}
func (params *UpdateIncomingPhoneNumberParams) SetVoiceApplicationSid(VoiceApplicationSid string) *UpdateIncomingPhoneNumberParams {
	params.VoiceApplicationSid = &VoiceApplicationSid
	return params
}
func (params *UpdateIncomingPhoneNumberParams) SetVoiceCallerIdLookup(VoiceCallerIdLookup bool) *UpdateIncomingPhoneNumberParams {
	params.VoiceCallerIdLookup = &VoiceCallerIdLookup
	return params
}
func (params *UpdateIncomingPhoneNumberParams) SetVoiceFallbackMethod(VoiceFallbackMethod string) *UpdateIncomingPhoneNumberParams {
	params.VoiceFallbackMethod = &VoiceFallbackMethod
	return params
}
func (params *UpdateIncomingPhoneNumberParams) SetVoiceFallbackUrl(VoiceFallbackUrl string) *UpdateIncomingPhoneNumberParams {
	params.VoiceFallbackUrl = &VoiceFallbackUrl
	return params
}
func (params *UpdateIncomingPhoneNumberParams) SetVoiceMethod(VoiceMethod string) *UpdateIncomingPhoneNumberParams {
	params.VoiceMethod = &VoiceMethod
	return params
}
func (params *UpdateIncomingPhoneNumberParams) SetVoiceUrl(VoiceUrl string) *UpdateIncomingPhoneNumberParams {
	params.VoiceUrl = &VoiceUrl
	return params
}
func (params *UpdateIncomingPhoneNumberParams) SetEmergencyStatus(EmergencyStatus string) *UpdateIncomingPhoneNumberParams {
	params.EmergencyStatus = &EmergencyStatus
	return params
}
func (params *UpdateIncomingPhoneNumberParams) SetEmergencyAddressSid(EmergencyAddressSid string) *UpdateIncomingPhoneNumberParams {
	params.EmergencyAddressSid = &EmergencyAddressSid
	return params
}
func (params *UpdateIncomingPhoneNumberParams) SetTrunkSid(TrunkSid string) *UpdateIncomingPhoneNumberParams {
	params.TrunkSid = &TrunkSid
	return params
}
func (params *UpdateIncomingPhoneNumberParams) SetVoiceReceiveMode(VoiceReceiveMode string) *UpdateIncomingPhoneNumberParams {
	params.VoiceReceiveMode = &VoiceReceiveMode
	return params
}
func (params *UpdateIncomingPhoneNumberParams) SetIdentitySid(IdentitySid string) *UpdateIncomingPhoneNumberParams {
	params.IdentitySid = &IdentitySid
	return params
}
func (params *UpdateIncomingPhoneNumberParams) SetAddressSid(AddressSid string) *UpdateIncomingPhoneNumberParams {
	params.AddressSid = &AddressSid
	return params
}
func (params *UpdateIncomingPhoneNumberParams) SetBundleSid(BundleSid string) *UpdateIncomingPhoneNumberParams {
	params.BundleSid = &BundleSid
	return params
}

// Update an incoming-phone-number instance.
func (c *ApiService) UpdateIncomingPhoneNumber(Sid string, params *UpdateIncomingPhoneNumberParams) (*ApiV2010IncomingPhoneNumber, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.AccountSid != nil {
		data.Set("AccountSid", *params.AccountSid)
	}
	if params != nil && params.ApiVersion != nil {
		data.Set("ApiVersion", *params.ApiVersion)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.SmsApplicationSid != nil {
		data.Set("SmsApplicationSid", *params.SmsApplicationSid)
	}
	if params != nil && params.SmsFallbackMethod != nil {
		data.Set("SmsFallbackMethod", *params.SmsFallbackMethod)
	}
	if params != nil && params.SmsFallbackUrl != nil {
		data.Set("SmsFallbackUrl", *params.SmsFallbackUrl)
	}
	if params != nil && params.SmsMethod != nil {
		data.Set("SmsMethod", *params.SmsMethod)
	}
	if params != nil && params.SmsUrl != nil {
		data.Set("SmsUrl", *params.SmsUrl)
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.StatusCallbackMethod != nil {
		data.Set("StatusCallbackMethod", *params.StatusCallbackMethod)
	}
	if params != nil && params.VoiceApplicationSid != nil {
		data.Set("VoiceApplicationSid", *params.VoiceApplicationSid)
	}
	if params != nil && params.VoiceCallerIdLookup != nil {
		data.Set("VoiceCallerIdLookup", fmt.Sprint(*params.VoiceCallerIdLookup))
	}
	if params != nil && params.VoiceFallbackMethod != nil {
		data.Set("VoiceFallbackMethod", *params.VoiceFallbackMethod)
	}
	if params != nil && params.VoiceFallbackUrl != nil {
		data.Set("VoiceFallbackUrl", *params.VoiceFallbackUrl)
	}
	if params != nil && params.VoiceMethod != nil {
		data.Set("VoiceMethod", *params.VoiceMethod)
	}
	if params != nil && params.VoiceUrl != nil {
		data.Set("VoiceUrl", *params.VoiceUrl)
	}
	if params != nil && params.EmergencyStatus != nil {
		data.Set("EmergencyStatus", *params.EmergencyStatus)
	}
	if params != nil && params.EmergencyAddressSid != nil {
		data.Set("EmergencyAddressSid", *params.EmergencyAddressSid)
	}
	if params != nil && params.TrunkSid != nil {
		data.Set("TrunkSid", *params.TrunkSid)
	}
	if params != nil && params.VoiceReceiveMode != nil {
		data.Set("VoiceReceiveMode", *params.VoiceReceiveMode)
	}
	if params != nil && params.IdentitySid != nil {
		data.Set("IdentitySid", *params.IdentitySid)
	}
	if params != nil && params.AddressSid != nil {
		data.Set("AddressSid", *params.AddressSid)
	}
	if params != nil && params.BundleSid != nil {
		data.Set("BundleSid", *params.BundleSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010IncomingPhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
