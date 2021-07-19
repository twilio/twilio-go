/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateByocTrunk'
type CreateByocTrunkParams struct {
	// Whether Caller ID Name (CNAM) lookup is enabled for the trunk. If enabled, all inbound calls to the BYOC Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information.
	CnamLookupEnabled *bool `json:"CnamLookupEnabled,omitempty"`
	// The SID of the Connection Policy that Twilio will use when routing traffic to your communications infrastructure.
	ConnectionPolicySid *string `json:"ConnectionPolicySid,omitempty"`
	// A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The SID of the SIP Domain that should be used in the `From` header of originating calls sent to your SIP infrastructure. If your SIP infrastructure allows users to \\\"call back\\\" an incoming call, configure this with a [SIP Domain](https://www.twilio.com/docs/voice/api/sending-sip) to ensure proper routing. If not configured, the from domain will default to \\\"sip.twilio.com\\\".
	FromDomainSid *string `json:"FromDomainSid,omitempty"`
	// The HTTP method we should use to call `status_callback_url`. Can be: `GET` or `POST`.
	StatusCallbackMethod *string `json:"StatusCallbackMethod,omitempty"`
	// The URL that we should call to pass status parameters (such as call ended) to your application.
	StatusCallbackUrl *string `json:"StatusCallbackUrl,omitempty"`
	// The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`.
	VoiceFallbackMethod *string `json:"VoiceFallbackMethod,omitempty"`
	// The URL that we should call when an error occurs while retrieving or executing the TwiML from `voice_url`.
	VoiceFallbackUrl *string `json:"VoiceFallbackUrl,omitempty"`
	// The HTTP method we should use to call `voice_url`. Can be: `GET` or `POST`.
	VoiceMethod *string `json:"VoiceMethod,omitempty"`
	// The URL we should call when the BYOC Trunk receives a call.
	VoiceUrl *string `json:"VoiceUrl,omitempty"`
}

func (params *CreateByocTrunkParams) SetCnamLookupEnabled(CnamLookupEnabled bool) *CreateByocTrunkParams {
	params.CnamLookupEnabled = &CnamLookupEnabled
	return params
}
func (params *CreateByocTrunkParams) SetConnectionPolicySid(ConnectionPolicySid string) *CreateByocTrunkParams {
	params.ConnectionPolicySid = &ConnectionPolicySid
	return params
}
func (params *CreateByocTrunkParams) SetFriendlyName(FriendlyName string) *CreateByocTrunkParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateByocTrunkParams) SetFromDomainSid(FromDomainSid string) *CreateByocTrunkParams {
	params.FromDomainSid = &FromDomainSid
	return params
}
func (params *CreateByocTrunkParams) SetStatusCallbackMethod(StatusCallbackMethod string) *CreateByocTrunkParams {
	params.StatusCallbackMethod = &StatusCallbackMethod
	return params
}
func (params *CreateByocTrunkParams) SetStatusCallbackUrl(StatusCallbackUrl string) *CreateByocTrunkParams {
	params.StatusCallbackUrl = &StatusCallbackUrl
	return params
}
func (params *CreateByocTrunkParams) SetVoiceFallbackMethod(VoiceFallbackMethod string) *CreateByocTrunkParams {
	params.VoiceFallbackMethod = &VoiceFallbackMethod
	return params
}
func (params *CreateByocTrunkParams) SetVoiceFallbackUrl(VoiceFallbackUrl string) *CreateByocTrunkParams {
	params.VoiceFallbackUrl = &VoiceFallbackUrl
	return params
}
func (params *CreateByocTrunkParams) SetVoiceMethod(VoiceMethod string) *CreateByocTrunkParams {
	params.VoiceMethod = &VoiceMethod
	return params
}
func (params *CreateByocTrunkParams) SetVoiceUrl(VoiceUrl string) *CreateByocTrunkParams {
	params.VoiceUrl = &VoiceUrl
	return params
}

func (c *ApiService) CreateByocTrunk(params *CreateByocTrunkParams) (*VoiceV1ByocTrunk, error) {
	path := "/v1/ByocTrunks"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CnamLookupEnabled != nil {
		data.Set("CnamLookupEnabled", fmt.Sprint(*params.CnamLookupEnabled))
	}
	if params != nil && params.ConnectionPolicySid != nil {
		data.Set("ConnectionPolicySid", *params.ConnectionPolicySid)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.FromDomainSid != nil {
		data.Set("FromDomainSid", *params.FromDomainSid)
	}
	if params != nil && params.StatusCallbackMethod != nil {
		data.Set("StatusCallbackMethod", *params.StatusCallbackMethod)
	}
	if params != nil && params.StatusCallbackUrl != nil {
		data.Set("StatusCallbackUrl", *params.StatusCallbackUrl)
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

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1ByocTrunk{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteByocTrunk(Sid string) error {
	path := "/v1/ByocTrunks/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ApiService) FetchByocTrunk(Sid string) (*VoiceV1ByocTrunk, error) {
	path := "/v1/ByocTrunks/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1ByocTrunk{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListByocTrunk'
type ListByocTrunkParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListByocTrunkParams) SetPageSize(PageSize int) *ListByocTrunkParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a single page of ByocTrunk records from the API. Request is executed immediately.
func (c *ApiService) PageByocTrunk(params *ListByocTrunkParams, pageToken string, pageNumber string) (*ListByocTrunkResponse, error) {
	path := "/v1/ByocTrunks"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageToken != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListByocTrunkResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ByocTrunk records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListByocTrunk(params *ListByocTrunkParams, limit int) ([]VoiceV1ByocTrunk, error) {
	if params == nil {
		params = &ListByocTrunkParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageByocTrunk(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []VoiceV1ByocTrunk

	for response != nil {
		records = append(records, response.ByocTrunks...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, limit, c.getNextListByocTrunkResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListByocTrunkResponse)
	}

	return records, err
}

// Streams ByocTrunk records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamByocTrunk(params *ListByocTrunkParams, limit int) (chan VoiceV1ByocTrunk, error) {
	if params == nil {
		params = &ListByocTrunkParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageByocTrunk(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan VoiceV1ByocTrunk, 1)

	go func() {
		for response != nil {
			for item := range response.ByocTrunks {
				channel <- response.ByocTrunks[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, limit, c.getNextListByocTrunkResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListByocTrunkResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListByocTrunkResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListByocTrunkResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateByocTrunk'
type UpdateByocTrunkParams struct {
	// Whether Caller ID Name (CNAM) lookup is enabled for the trunk. If enabled, all inbound calls to the BYOC Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information.
	CnamLookupEnabled *bool `json:"CnamLookupEnabled,omitempty"`
	// The SID of the Connection Policy that Twilio will use when routing traffic to your communications infrastructure.
	ConnectionPolicySid *string `json:"ConnectionPolicySid,omitempty"`
	// A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The SID of the SIP Domain that should be used in the `From` header of originating calls sent to your SIP infrastructure. If your SIP infrastructure allows users to \\\"call back\\\" an incoming call, configure this with a [SIP Domain](https://www.twilio.com/docs/voice/api/sending-sip) to ensure proper routing. If not configured, the from domain will default to \\\"sip.twilio.com\\\".
	FromDomainSid *string `json:"FromDomainSid,omitempty"`
	// The HTTP method we should use to call `status_callback_url`. Can be: `GET` or `POST`.
	StatusCallbackMethod *string `json:"StatusCallbackMethod,omitempty"`
	// The URL that we should call to pass status parameters (such as call ended) to your application.
	StatusCallbackUrl *string `json:"StatusCallbackUrl,omitempty"`
	// The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`.
	VoiceFallbackMethod *string `json:"VoiceFallbackMethod,omitempty"`
	// The URL that we should call when an error occurs while retrieving or executing the TwiML requested by `voice_url`.
	VoiceFallbackUrl *string `json:"VoiceFallbackUrl,omitempty"`
	// The HTTP method we should use to call `voice_url`
	VoiceMethod *string `json:"VoiceMethod,omitempty"`
	// The URL we should call when the BYOC Trunk receives a call.
	VoiceUrl *string `json:"VoiceUrl,omitempty"`
}

func (params *UpdateByocTrunkParams) SetCnamLookupEnabled(CnamLookupEnabled bool) *UpdateByocTrunkParams {
	params.CnamLookupEnabled = &CnamLookupEnabled
	return params
}
func (params *UpdateByocTrunkParams) SetConnectionPolicySid(ConnectionPolicySid string) *UpdateByocTrunkParams {
	params.ConnectionPolicySid = &ConnectionPolicySid
	return params
}
func (params *UpdateByocTrunkParams) SetFriendlyName(FriendlyName string) *UpdateByocTrunkParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateByocTrunkParams) SetFromDomainSid(FromDomainSid string) *UpdateByocTrunkParams {
	params.FromDomainSid = &FromDomainSid
	return params
}
func (params *UpdateByocTrunkParams) SetStatusCallbackMethod(StatusCallbackMethod string) *UpdateByocTrunkParams {
	params.StatusCallbackMethod = &StatusCallbackMethod
	return params
}
func (params *UpdateByocTrunkParams) SetStatusCallbackUrl(StatusCallbackUrl string) *UpdateByocTrunkParams {
	params.StatusCallbackUrl = &StatusCallbackUrl
	return params
}
func (params *UpdateByocTrunkParams) SetVoiceFallbackMethod(VoiceFallbackMethod string) *UpdateByocTrunkParams {
	params.VoiceFallbackMethod = &VoiceFallbackMethod
	return params
}
func (params *UpdateByocTrunkParams) SetVoiceFallbackUrl(VoiceFallbackUrl string) *UpdateByocTrunkParams {
	params.VoiceFallbackUrl = &VoiceFallbackUrl
	return params
}
func (params *UpdateByocTrunkParams) SetVoiceMethod(VoiceMethod string) *UpdateByocTrunkParams {
	params.VoiceMethod = &VoiceMethod
	return params
}
func (params *UpdateByocTrunkParams) SetVoiceUrl(VoiceUrl string) *UpdateByocTrunkParams {
	params.VoiceUrl = &VoiceUrl
	return params
}

func (c *ApiService) UpdateByocTrunk(Sid string, params *UpdateByocTrunkParams) (*VoiceV1ByocTrunk, error) {
	path := "/v1/ByocTrunks/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CnamLookupEnabled != nil {
		data.Set("CnamLookupEnabled", fmt.Sprint(*params.CnamLookupEnabled))
	}
	if params != nil && params.ConnectionPolicySid != nil {
		data.Set("ConnectionPolicySid", *params.ConnectionPolicySid)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.FromDomainSid != nil {
		data.Set("FromDomainSid", *params.FromDomainSid)
	}
	if params != nil && params.StatusCallbackMethod != nil {
		data.Set("StatusCallbackMethod", *params.StatusCallbackMethod)
	}
	if params != nil && params.StatusCallbackUrl != nil {
		data.Set("StatusCallbackUrl", *params.StatusCallbackUrl)
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

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1ByocTrunk{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
