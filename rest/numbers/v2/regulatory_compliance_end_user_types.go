/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.18.0
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
}

func (params *ListEndUserTypeParams) SetPageSize(PageSize int) *ListEndUserTypeParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all End-User Types.
func (c *ApiService) ListEndUserType(params *ListEndUserTypeParams) (*ListEndUserTypeResponse, error) {
	path := "/v2/RegulatoryCompliance/EndUserTypes"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
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

//Retrieve a single page of  records from the API. Request is executed immediately.
func (c *ApiService) RegulatoryComplianceEndUserTypesPage(params *ListEndUserTypeParams, pageToken string, pageNumber string) (*ListEndUserTypeResponse, error) {
	path := "/v2/RegulatoryCompliance/EndUserTypes"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	data.Set("PageToken", pageToken)
	data.Set("PageNumber", pageNumber)

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

//Lists RegulatoryComplianceEndUserTypes records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) RegulatoryComplianceEndUserTypesList(params *ListEndUserTypeParams, limit int) ([]ListEndUserTypeResponse, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	response, err := c.ListEndUserType(params)
	if err != nil {
		return nil, err
	}

	page := client.NewPage(c.baseURL, response)

	resp := c.requestHandler.List(page, limit, 0)
	ret := make([]ListEndUserTypeResponse, len(resp))

	for i := range resp {
		jsonStr, _ := json.Marshal(resp[i])
		ps := ListEndUserTypeResponse{}
		if err := json.Unmarshal(jsonStr, &ps); err != nil {
			return ret, err
		}

		ret[i] = ps
	}

	return ret, nil
}

//Streams RegulatoryComplianceEndUserTypes records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) RegulatoryComplianceEndUserTypesStream(params *ListEndUserTypeParams, limit int) (chan interface{}, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	response, err := c.ListEndUserType(params)
	if err != nil {
		return nil, err
	}

	page := client.NewPage(c.baseURL, response)

	ps := ListEndUserTypeResponse{}
	return c.requestHandler.Stream(page, limit, 0, ps), nil
}
