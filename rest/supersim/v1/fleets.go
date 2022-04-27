/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
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

// Optional parameters for the method 'CreateFleet'
type CreateFleetParams struct {
	// Defines whether SIMs in the Fleet are capable of using 2G/3G/4G/LTE/CAT-M data connectivity. Defaults to `true`.
	DataEnabled *bool `json:"DataEnabled,omitempty"`
	// The total data usage (download and upload combined) in Megabytes that each Sim resource assigned to the Fleet resource can consume during a billing period (normally one month). Value must be between 1MB (1) and 2TB (2,000,000). Defaults to 1GB (1,000).
	DataLimit *int `json:"DataLimit,omitempty"`
	// A string representing the HTTP method to use when making a request to `ip_commands_url`. Can be one of `POST` or `GET`. Defaults to `POST`.
	IpCommandsMethod *string `json:"IpCommandsMethod,omitempty"`
	// The URL that will receive a webhook when a Super SIM in the Fleet is used to send an IP Command from your device to a special IP address. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored.
	IpCommandsUrl *string `json:"IpCommandsUrl,omitempty"`
	// The SID or unique name of the Network Access Profile that will control which cellular networks the Fleet's SIMs can connect to.
	NetworkAccessProfile *string `json:"NetworkAccessProfile,omitempty"`
	// Defines whether SIMs in the Fleet are capable of sending and receiving machine-to-machine SMS via Commands. Defaults to `true`.
	SmsCommandsEnabled *bool `json:"SmsCommandsEnabled,omitempty"`
	// A string representing the HTTP method to use when making a request to `sms_commands_url`. Can be one of `POST` or `GET`. Defaults to `POST`.
	SmsCommandsMethod *string `json:"SmsCommandsMethod,omitempty"`
	// The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the SMS Commands number. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored.
	SmsCommandsUrl *string `json:"SmsCommandsUrl,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *CreateFleetParams) SetDataEnabled(DataEnabled bool) *CreateFleetParams {
	params.DataEnabled = &DataEnabled
	return params
}
func (params *CreateFleetParams) SetDataLimit(DataLimit int) *CreateFleetParams {
	params.DataLimit = &DataLimit
	return params
}
func (params *CreateFleetParams) SetIpCommandsMethod(IpCommandsMethod string) *CreateFleetParams {
	params.IpCommandsMethod = &IpCommandsMethod
	return params
}
func (params *CreateFleetParams) SetIpCommandsUrl(IpCommandsUrl string) *CreateFleetParams {
	params.IpCommandsUrl = &IpCommandsUrl
	return params
}
func (params *CreateFleetParams) SetNetworkAccessProfile(NetworkAccessProfile string) *CreateFleetParams {
	params.NetworkAccessProfile = &NetworkAccessProfile
	return params
}
func (params *CreateFleetParams) SetSmsCommandsEnabled(SmsCommandsEnabled bool) *CreateFleetParams {
	params.SmsCommandsEnabled = &SmsCommandsEnabled
	return params
}
func (params *CreateFleetParams) SetSmsCommandsMethod(SmsCommandsMethod string) *CreateFleetParams {
	params.SmsCommandsMethod = &SmsCommandsMethod
	return params
}
func (params *CreateFleetParams) SetSmsCommandsUrl(SmsCommandsUrl string) *CreateFleetParams {
	params.SmsCommandsUrl = &SmsCommandsUrl
	return params
}
func (params *CreateFleetParams) SetUniqueName(UniqueName string) *CreateFleetParams {
	params.UniqueName = &UniqueName
	return params
}

// Create a Fleet
func (c *ApiService) CreateFleet(params *CreateFleetParams) (*SupersimV1Fleet, error) {
	path := "/v1/Fleets"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.DataEnabled != nil {
		data.Set("DataEnabled", fmt.Sprint(*params.DataEnabled))
	}
	if params != nil && params.DataLimit != nil {
		data.Set("DataLimit", fmt.Sprint(*params.DataLimit))
	}
	if params != nil && params.IpCommandsMethod != nil {
		data.Set("IpCommandsMethod", *params.IpCommandsMethod)
	}
	if params != nil && params.IpCommandsUrl != nil {
		data.Set("IpCommandsUrl", *params.IpCommandsUrl)
	}
	if params != nil && params.NetworkAccessProfile != nil {
		data.Set("NetworkAccessProfile", *params.NetworkAccessProfile)
	}
	if params != nil && params.SmsCommandsEnabled != nil {
		data.Set("SmsCommandsEnabled", fmt.Sprint(*params.SmsCommandsEnabled))
	}
	if params != nil && params.SmsCommandsMethod != nil {
		data.Set("SmsCommandsMethod", *params.SmsCommandsMethod)
	}
	if params != nil && params.SmsCommandsUrl != nil {
		data.Set("SmsCommandsUrl", *params.SmsCommandsUrl)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1Fleet{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Fetch a Fleet instance from your account.
func (c *ApiService) FetchFleet(Sid string) (*SupersimV1Fleet, error) {
	path := "/v1/Fleets/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1Fleet{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListFleet'
type ListFleetParams struct {
	// The SID or unique name of the Network Access Profile that controls which cellular networks the Fleet's SIMs can connect to.
	NetworkAccessProfile *string `json:"NetworkAccessProfile,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListFleetParams) SetNetworkAccessProfile(NetworkAccessProfile string) *ListFleetParams {
	params.NetworkAccessProfile = &NetworkAccessProfile
	return params
}
func (params *ListFleetParams) SetPageSize(PageSize int) *ListFleetParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListFleetParams) SetLimit(Limit int) *ListFleetParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Fleet records from the API. Request is executed immediately.
func (c *ApiService) PageFleet(params *ListFleetParams, pageToken, pageNumber string) (*ListFleetResponse, error) {
	path := "/v1/Fleets"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.NetworkAccessProfile != nil {
		data.Set("NetworkAccessProfile", *params.NetworkAccessProfile)
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

	ps := &ListFleetResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Fleet records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListFleet(params *ListFleetParams) ([]SupersimV1Fleet, error) {
	response, errors := c.StreamFleet(params)

	records := make([]SupersimV1Fleet, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Fleet records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamFleet(params *ListFleetParams) (chan SupersimV1Fleet, chan error) {
	if params == nil {
		params = &ListFleetParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan SupersimV1Fleet, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageFleet(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamFleet(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamFleet(response *ListFleetResponse, params *ListFleetParams, recordChannel chan SupersimV1Fleet, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Fleets
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListFleetResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListFleetResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListFleetResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFleetResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateFleet'
type UpdateFleetParams struct {
	// A string representing the HTTP method to use when making a request to `ip_commands_url`. Can be one of `POST` or `GET`. Defaults to `POST`.
	IpCommandsMethod *string `json:"IpCommandsMethod,omitempty"`
	// The URL that will receive a webhook when a Super SIM in the Fleet is used to send an IP Command from your device to a special IP address. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored.
	IpCommandsUrl *string `json:"IpCommandsUrl,omitempty"`
	// The SID or unique name of the Network Access Profile that will control which cellular networks the Fleet's SIMs can connect to.
	NetworkAccessProfile *string `json:"NetworkAccessProfile,omitempty"`
	// A string representing the HTTP method to use when making a request to `sms_commands_url`. Can be one of `POST` or `GET`. Defaults to `POST`.
	SmsCommandsMethod *string `json:"SmsCommandsMethod,omitempty"`
	// The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the SMS Commands number. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored.
	SmsCommandsUrl *string `json:"SmsCommandsUrl,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *UpdateFleetParams) SetIpCommandsMethod(IpCommandsMethod string) *UpdateFleetParams {
	params.IpCommandsMethod = &IpCommandsMethod
	return params
}
func (params *UpdateFleetParams) SetIpCommandsUrl(IpCommandsUrl string) *UpdateFleetParams {
	params.IpCommandsUrl = &IpCommandsUrl
	return params
}
func (params *UpdateFleetParams) SetNetworkAccessProfile(NetworkAccessProfile string) *UpdateFleetParams {
	params.NetworkAccessProfile = &NetworkAccessProfile
	return params
}
func (params *UpdateFleetParams) SetSmsCommandsMethod(SmsCommandsMethod string) *UpdateFleetParams {
	params.SmsCommandsMethod = &SmsCommandsMethod
	return params
}
func (params *UpdateFleetParams) SetSmsCommandsUrl(SmsCommandsUrl string) *UpdateFleetParams {
	params.SmsCommandsUrl = &SmsCommandsUrl
	return params
}
func (params *UpdateFleetParams) SetUniqueName(UniqueName string) *UpdateFleetParams {
	params.UniqueName = &UniqueName
	return params
}

// Updates the given properties of a Super SIM Fleet instance from your account.
func (c *ApiService) UpdateFleet(Sid string, params *UpdateFleetParams) (*SupersimV1Fleet, error) {
	path := "/v1/Fleets/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.IpCommandsMethod != nil {
		data.Set("IpCommandsMethod", *params.IpCommandsMethod)
	}
	if params != nil && params.IpCommandsUrl != nil {
		data.Set("IpCommandsUrl", *params.IpCommandsUrl)
	}
	if params != nil && params.NetworkAccessProfile != nil {
		data.Set("NetworkAccessProfile", *params.NetworkAccessProfile)
	}
	if params != nil && params.SmsCommandsMethod != nil {
		data.Set("SmsCommandsMethod", *params.SmsCommandsMethod)
	}
	if params != nil && params.SmsCommandsUrl != nil {
		data.Set("SmsCommandsUrl", *params.SmsCommandsUrl)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1Fleet{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
