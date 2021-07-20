/*
 * Twilio - Numbers
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

// Fetch a specific End-User Type Instance.
func (c *ApiService) FetchEndUserType(Sid string) (*NumbersV2RegulatoryComplianceEndUserType, error) {
	path := "/v2/RegulatoryCompliance/EndUserTypes/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2RegulatoryComplianceEndUserType{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListEndUserType'
type ListEndUserTypeParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListEndUserTypeParams) SetPageSize(PageSize int) *ListEndUserTypeParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListEndUserTypeParams) SetLimit(Limit int) *ListEndUserTypeParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of EndUserType records from the API. Request is executed immediately.
func (c *ApiService) PageEndUserType(params *ListEndUserTypeParams, pageToken string, pageNumber string) (*ListEndUserTypeResponse, error) {
	path := "/v2/RegulatoryCompliance/EndUserTypes"

	data := url.Values{}
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

	ps := &ListEndUserTypeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists EndUserType records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListEndUserType(params *ListEndUserTypeParams) ([]NumbersV2RegulatoryComplianceEndUserType, error) {
	if params == nil {
		params = &ListEndUserTypeParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageEndUserType(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []NumbersV2RegulatoryComplianceEndUserType

	for response != nil {
		records = append(records, response.EndUserTypes...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListEndUserTypeResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListEndUserTypeResponse)
	}

	return records, err
}

// Streams EndUserType records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamEndUserType(params *ListEndUserTypeParams) (chan NumbersV2RegulatoryComplianceEndUserType, error) {
	if params == nil {
		params = &ListEndUserTypeParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageEndUserType(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan NumbersV2RegulatoryComplianceEndUserType, 1)

	go func() {
		for response != nil {
			for item := range response.EndUserTypes {
				channel <- response.EndUserTypes[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListEndUserTypeResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListEndUserTypeResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListEndUserTypeResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListEndUserTypeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
