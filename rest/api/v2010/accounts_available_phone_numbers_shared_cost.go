/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.0
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

// Optional parameters for the method 'ListAvailablePhoneNumberSharedCost'
type ListAvailablePhoneNumberSharedCostParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada.
	AreaCode *int `json:"AreaCode,omitempty"`
	// The pattern on which to match phone numbers. Valid characters are `*`, `0-9`, `a-z`, and `A-Z`. The `*` character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters.
	Contains *string `json:"Contains,omitempty"`
	// Whether the phone numbers can receive text messages. Can be: `true` or `false`.
	SmsEnabled *bool `json:"SmsEnabled,omitempty"`
	// Whether the phone numbers can receive MMS messages. Can be: `true` or `false`.
	MmsEnabled *bool `json:"MmsEnabled,omitempty"`
	// Whether the phone numbers can receive calls. Can be: `true` or `false`.
	VoiceEnabled *bool `json:"VoiceEnabled,omitempty"`
	// Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`.
	ExcludeAllAddressRequired *bool `json:"ExcludeAllAddressRequired,omitempty"`
	// Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`.
	ExcludeLocalAddressRequired *bool `json:"ExcludeLocalAddressRequired,omitempty"`
	// Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`.
	ExcludeForeignAddressRequired *bool `json:"ExcludeForeignAddressRequired,omitempty"`
	// Whether to read phone numbers that are new to the Twilio platform. Can be: `true` or `false` and the default is `true`.
	Beta *bool `json:"Beta,omitempty"`
	// Given a phone number, find a geographically close number within `distance` miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada.
	NearNumber *string `json:"NearNumber,omitempty"`
	// Given a latitude/longitude pair `lat,long` find geographically close numbers within `distance` miles. Applies to only phone numbers in the US and Canada.
	NearLatLong *string `json:"NearLatLong,omitempty"`
	// The search radius, in miles, for a `near_` query.  Can be up to `500` and the default is `25`. Applies to only phone numbers in the US and Canada.
	Distance *int `json:"Distance,omitempty"`
	// Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada.
	InPostalCode *string `json:"InPostalCode,omitempty"`
	// Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada.
	InRegion *string `json:"InRegion,omitempty"`
	// Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires `in_lata` to be set as well. Applies to only phone numbers in the US and Canada.
	InRateCenter *string `json:"InRateCenter,omitempty"`
	// Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada.
	InLata *string `json:"InLata,omitempty"`
	// Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number.
	InLocality *string `json:"InLocality,omitempty"`
	// Whether the phone numbers can receive faxes. Can be: `true` or `false`.
	FaxEnabled *bool `json:"FaxEnabled,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListAvailablePhoneNumberSharedCostParams) SetPathAccountSid(PathAccountSid string) *ListAvailablePhoneNumberSharedCostParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListAvailablePhoneNumberSharedCostParams) SetAreaCode(AreaCode int) *ListAvailablePhoneNumberSharedCostParams {
	params.AreaCode = &AreaCode
	return params
}
func (params *ListAvailablePhoneNumberSharedCostParams) SetContains(Contains string) *ListAvailablePhoneNumberSharedCostParams {
	params.Contains = &Contains
	return params
}
func (params *ListAvailablePhoneNumberSharedCostParams) SetSmsEnabled(SmsEnabled bool) *ListAvailablePhoneNumberSharedCostParams {
	params.SmsEnabled = &SmsEnabled
	return params
}
func (params *ListAvailablePhoneNumberSharedCostParams) SetMmsEnabled(MmsEnabled bool) *ListAvailablePhoneNumberSharedCostParams {
	params.MmsEnabled = &MmsEnabled
	return params
}
func (params *ListAvailablePhoneNumberSharedCostParams) SetVoiceEnabled(VoiceEnabled bool) *ListAvailablePhoneNumberSharedCostParams {
	params.VoiceEnabled = &VoiceEnabled
	return params
}
func (params *ListAvailablePhoneNumberSharedCostParams) SetExcludeAllAddressRequired(ExcludeAllAddressRequired bool) *ListAvailablePhoneNumberSharedCostParams {
	params.ExcludeAllAddressRequired = &ExcludeAllAddressRequired
	return params
}
func (params *ListAvailablePhoneNumberSharedCostParams) SetExcludeLocalAddressRequired(ExcludeLocalAddressRequired bool) *ListAvailablePhoneNumberSharedCostParams {
	params.ExcludeLocalAddressRequired = &ExcludeLocalAddressRequired
	return params
}
func (params *ListAvailablePhoneNumberSharedCostParams) SetExcludeForeignAddressRequired(ExcludeForeignAddressRequired bool) *ListAvailablePhoneNumberSharedCostParams {
	params.ExcludeForeignAddressRequired = &ExcludeForeignAddressRequired
	return params
}
func (params *ListAvailablePhoneNumberSharedCostParams) SetBeta(Beta bool) *ListAvailablePhoneNumberSharedCostParams {
	params.Beta = &Beta
	return params
}
func (params *ListAvailablePhoneNumberSharedCostParams) SetNearNumber(NearNumber string) *ListAvailablePhoneNumberSharedCostParams {
	params.NearNumber = &NearNumber
	return params
}
func (params *ListAvailablePhoneNumberSharedCostParams) SetNearLatLong(NearLatLong string) *ListAvailablePhoneNumberSharedCostParams {
	params.NearLatLong = &NearLatLong
	return params
}
func (params *ListAvailablePhoneNumberSharedCostParams) SetDistance(Distance int) *ListAvailablePhoneNumberSharedCostParams {
	params.Distance = &Distance
	return params
}
func (params *ListAvailablePhoneNumberSharedCostParams) SetInPostalCode(InPostalCode string) *ListAvailablePhoneNumberSharedCostParams {
	params.InPostalCode = &InPostalCode
	return params
}
func (params *ListAvailablePhoneNumberSharedCostParams) SetInRegion(InRegion string) *ListAvailablePhoneNumberSharedCostParams {
	params.InRegion = &InRegion
	return params
}
func (params *ListAvailablePhoneNumberSharedCostParams) SetInRateCenter(InRateCenter string) *ListAvailablePhoneNumberSharedCostParams {
	params.InRateCenter = &InRateCenter
	return params
}
func (params *ListAvailablePhoneNumberSharedCostParams) SetInLata(InLata string) *ListAvailablePhoneNumberSharedCostParams {
	params.InLata = &InLata
	return params
}
func (params *ListAvailablePhoneNumberSharedCostParams) SetInLocality(InLocality string) *ListAvailablePhoneNumberSharedCostParams {
	params.InLocality = &InLocality
	return params
}
func (params *ListAvailablePhoneNumberSharedCostParams) SetFaxEnabled(FaxEnabled bool) *ListAvailablePhoneNumberSharedCostParams {
	params.FaxEnabled = &FaxEnabled
	return params
}
func (params *ListAvailablePhoneNumberSharedCostParams) SetPageSize(PageSize int) *ListAvailablePhoneNumberSharedCostParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListAvailablePhoneNumberSharedCostParams) SetLimit(Limit int) *ListAvailablePhoneNumberSharedCostParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of AvailablePhoneNumberSharedCost records from the API. Request is executed immediately.
func (c *ApiService) PageAvailablePhoneNumberSharedCost(CountryCode string, params *ListAvailablePhoneNumberSharedCostParams, pageToken string, pageNumber string) (*ListAvailablePhoneNumberSharedCostResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers/{CountryCode}/SharedCost.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CountryCode"+"}", CountryCode, -1)

	data := url.Values{}
	if params != nil && params.AreaCode != nil {
		data.Set("AreaCode", fmt.Sprint(*params.AreaCode))
	}
	if params != nil && params.Contains != nil {
		data.Set("Contains", *params.Contains)
	}
	if params != nil && params.SmsEnabled != nil {
		data.Set("SmsEnabled", fmt.Sprint(*params.SmsEnabled))
	}
	if params != nil && params.MmsEnabled != nil {
		data.Set("MmsEnabled", fmt.Sprint(*params.MmsEnabled))
	}
	if params != nil && params.VoiceEnabled != nil {
		data.Set("VoiceEnabled", fmt.Sprint(*params.VoiceEnabled))
	}
	if params != nil && params.ExcludeAllAddressRequired != nil {
		data.Set("ExcludeAllAddressRequired", fmt.Sprint(*params.ExcludeAllAddressRequired))
	}
	if params != nil && params.ExcludeLocalAddressRequired != nil {
		data.Set("ExcludeLocalAddressRequired", fmt.Sprint(*params.ExcludeLocalAddressRequired))
	}
	if params != nil && params.ExcludeForeignAddressRequired != nil {
		data.Set("ExcludeForeignAddressRequired", fmt.Sprint(*params.ExcludeForeignAddressRequired))
	}
	if params != nil && params.Beta != nil {
		data.Set("Beta", fmt.Sprint(*params.Beta))
	}
	if params != nil && params.NearNumber != nil {
		data.Set("NearNumber", *params.NearNumber)
	}
	if params != nil && params.NearLatLong != nil {
		data.Set("NearLatLong", *params.NearLatLong)
	}
	if params != nil && params.Distance != nil {
		data.Set("Distance", fmt.Sprint(*params.Distance))
	}
	if params != nil && params.InPostalCode != nil {
		data.Set("InPostalCode", *params.InPostalCode)
	}
	if params != nil && params.InRegion != nil {
		data.Set("InRegion", *params.InRegion)
	}
	if params != nil && params.InRateCenter != nil {
		data.Set("InRateCenter", *params.InRateCenter)
	}
	if params != nil && params.InLata != nil {
		data.Set("InLata", *params.InLata)
	}
	if params != nil && params.InLocality != nil {
		data.Set("InLocality", *params.InLocality)
	}
	if params != nil && params.FaxEnabled != nil {
		data.Set("FaxEnabled", fmt.Sprint(*params.FaxEnabled))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}
	headers := make(map[string]interface{})

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

	ps := &ListAvailablePhoneNumberSharedCostResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists AvailablePhoneNumberSharedCost records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListAvailablePhoneNumberSharedCost(CountryCode string, params *ListAvailablePhoneNumberSharedCostParams) ([]ApiV2010AvailablePhoneNumberSharedCost, error) {
	if params == nil {
		params = &ListAvailablePhoneNumberSharedCostParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageAvailablePhoneNumberSharedCost(CountryCode, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010AvailablePhoneNumberSharedCost

	for response != nil {
		records = append(records, response.AvailablePhoneNumbers...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListAvailablePhoneNumberSharedCostResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListAvailablePhoneNumberSharedCostResponse)
	}

	return records, err
}

// Streams AvailablePhoneNumberSharedCost records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamAvailablePhoneNumberSharedCost(CountryCode string, params *ListAvailablePhoneNumberSharedCostParams) (chan ApiV2010AvailablePhoneNumberSharedCost, error) {
	if params == nil {
		params = &ListAvailablePhoneNumberSharedCostParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageAvailablePhoneNumberSharedCost(CountryCode, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010AvailablePhoneNumberSharedCost, 1)

	go func() {
		for response != nil {
			for item := range response.AvailablePhoneNumbers {
				channel <- response.AvailablePhoneNumbers[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListAvailablePhoneNumberSharedCostResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListAvailablePhoneNumberSharedCostResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListAvailablePhoneNumberSharedCostResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAvailablePhoneNumberSharedCostResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
