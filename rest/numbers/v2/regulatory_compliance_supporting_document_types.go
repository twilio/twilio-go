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

// Fetch a specific Supporting Document Type Instance.
func (c *ApiService) FetchSupportingDocumentType(Sid string) (*NumbersV2RegulatoryComplianceSupportingDocumentType, error) {
	path := "/v2/RegulatoryCompliance/SupportingDocumentTypes/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2RegulatoryComplianceSupportingDocumentType{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSupportingDocumentType'
type ListSupportingDocumentTypeParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSupportingDocumentTypeParams) SetPageSize(PageSize int) *ListSupportingDocumentTypeParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSupportingDocumentTypeParams) SetLimit(Limit int) *ListSupportingDocumentTypeParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SupportingDocumentType records from the API. Request is executed immediately.
func (c *ApiService) PageSupportingDocumentType(params *ListSupportingDocumentTypeParams, pageToken string, pageNumber string) (*ListSupportingDocumentTypeResponse, error) {
	path := "/v2/RegulatoryCompliance/SupportingDocumentTypes"

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

	ps := &ListSupportingDocumentTypeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SupportingDocumentType records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSupportingDocumentType(params *ListSupportingDocumentTypeParams) ([]NumbersV2RegulatoryComplianceSupportingDocumentType, error) {
	if params == nil {
		params = &ListSupportingDocumentTypeParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSupportingDocumentType(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []NumbersV2RegulatoryComplianceSupportingDocumentType

	for response != nil {
		records = append(records, response.SupportingDocumentTypes...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListSupportingDocumentTypeResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListSupportingDocumentTypeResponse)
	}

	return records, err
}

// Streams SupportingDocumentType records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSupportingDocumentType(params *ListSupportingDocumentTypeParams) (chan NumbersV2RegulatoryComplianceSupportingDocumentType, error) {
	if params == nil {
		params = &ListSupportingDocumentTypeParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSupportingDocumentType(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan NumbersV2RegulatoryComplianceSupportingDocumentType, 1)

	go func() {
		for response != nil {
			for item := range response.SupportingDocumentTypes {
				channel <- response.SupportingDocumentTypes[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListSupportingDocumentTypeResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListSupportingDocumentTypeResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListSupportingDocumentTypeResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSupportingDocumentTypeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
