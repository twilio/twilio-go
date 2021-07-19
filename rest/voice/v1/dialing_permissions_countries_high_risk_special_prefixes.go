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

// Optional parameters for the method 'ListDialingPermissionsHrsPrefixes'
type ListDialingPermissionsHrsPrefixesParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListDialingPermissionsHrsPrefixesParams) SetPageSize(PageSize int) *ListDialingPermissionsHrsPrefixesParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a single page of DialingPermissionsHrsPrefixes records from the API. Request is executed immediately.
func (c *ApiService) PageDialingPermissionsHrsPrefixes(IsoCode string, params *ListDialingPermissionsHrsPrefixesParams, pageToken string, pageNumber string) (*ListDialingPermissionsHrsPrefixesResponse, error) {
	path := "/v1/DialingPermissions/Countries/{IsoCode}/HighRiskSpecialPrefixes"

	path = strings.Replace(path, "{"+"IsoCode"+"}", IsoCode, -1)

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

	ps := &ListDialingPermissionsHrsPrefixesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists DialingPermissionsHrsPrefixes records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListDialingPermissionsHrsPrefixes(IsoCode string, params *ListDialingPermissionsHrsPrefixesParams, limit int) ([]VoiceV1DialingPermissionsDialingPermissionsCountryDialingPermissionsHrsPrefixes, error) {
	if params == nil {
		params = &ListDialingPermissionsHrsPrefixesParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageDialingPermissionsHrsPrefixes(IsoCode, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []VoiceV1DialingPermissionsDialingPermissionsCountryDialingPermissionsHrsPrefixes

	for response != nil {
		records = append(records, response.Content...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, limit, c.getNextListDialingPermissionsHrsPrefixesResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListDialingPermissionsHrsPrefixesResponse)
	}

	return records, err
}

// Streams DialingPermissionsHrsPrefixes records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamDialingPermissionsHrsPrefixes(IsoCode string, params *ListDialingPermissionsHrsPrefixesParams, limit int) (chan VoiceV1DialingPermissionsDialingPermissionsCountryDialingPermissionsHrsPrefixes, error) {
	if params == nil {
		params = &ListDialingPermissionsHrsPrefixesParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageDialingPermissionsHrsPrefixes(IsoCode, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan VoiceV1DialingPermissionsDialingPermissionsCountryDialingPermissionsHrsPrefixes, 1)

	go func() {
		for response != nil {
			for item := range response.Content {
				channel <- response.Content[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, limit, c.getNextListDialingPermissionsHrsPrefixesResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListDialingPermissionsHrsPrefixesResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListDialingPermissionsHrsPrefixesResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListDialingPermissionsHrsPrefixesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
