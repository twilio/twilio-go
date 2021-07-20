/*
 * Twilio - Pricing
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

// Fetch a specific Country.
func (c *ApiService) FetchVoiceCountry(IsoCountry string) (*PricingV2VoiceVoiceCountryInstance, error) {
	path := "/v2/Voice/Countries/{IsoCountry}"
	path = strings.Replace(path, "{"+"IsoCountry"+"}", IsoCountry, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &PricingV2VoiceVoiceCountryInstance{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListVoiceCountry'
type ListVoiceCountryParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListVoiceCountryParams) SetPageSize(PageSize int) *ListVoiceCountryParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListVoiceCountryParams) SetLimit(Limit int) *ListVoiceCountryParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of VoiceCountry records from the API. Request is executed immediately.
func (c *ApiService) PageVoiceCountry(params *ListVoiceCountryParams, pageToken string, pageNumber string) (*ListVoiceCountryResponse, error) {
	path := "/v2/Voice/Countries"

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

	ps := &ListVoiceCountryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists VoiceCountry records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListVoiceCountry(params *ListVoiceCountryParams) ([]PricingV2VoiceVoiceCountry, error) {
	if params == nil {
		params = &ListVoiceCountryParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageVoiceCountry(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []PricingV2VoiceVoiceCountry

	for response != nil {
		records = append(records, response.Countries...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListVoiceCountryResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListVoiceCountryResponse)
	}

	return records, err
}

// Streams VoiceCountry records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamVoiceCountry(params *ListVoiceCountryParams) (chan PricingV2VoiceVoiceCountry, error) {
	if params == nil {
		params = &ListVoiceCountryParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageVoiceCountry(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan PricingV2VoiceVoiceCountry, 1)

	go func() {
		for response != nil {
			for item := range response.Countries {
				channel <- response.Countries[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListVoiceCountryResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListVoiceCountryResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListVoiceCountryResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListVoiceCountryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
