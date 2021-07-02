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
}

func (params *ListSupportingDocumentTypeParams) SetPageSize(PageSize int) *ListSupportingDocumentTypeParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all Supporting Document Types.
func (c *ApiService) ListSupportingDocumentType(params *ListSupportingDocumentTypeParams) (*ListSupportingDocumentTypeResponse, error) {
	path := "/v2/RegulatoryCompliance/SupportingDocumentTypes"

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

	ps := &ListSupportingDocumentTypeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of SupportingDocumentType records from the API. Request is executed immediately.
func (c *ApiService) SupportingDocumentTypePage(params *ListSupportingDocumentTypeParams, pageToken string, pageNumber string) (*client.Page, error) {
	path := "/v2/RegulatoryCompliance/SupportingDocumentTypes"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	data.Set("PageToken", pageToken)
	data.Set("PageNumber", pageNumber)

	response, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	return client.NewPage(c.baseURL, response), nil
}

//Streams SupportingDocumentType records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) SupportingDocumentTypeStream(params *ListSupportingDocumentTypeParams, limit int) (chan map[string]interface{}, error) {
	if params.PageSize == nil {
		params.SetPageSize(0)
	}
	params.SetPageSize(c.requestHandler.ReadLimits(*params.PageSize, limit))
	page, err := c.SupportingDocumentTypePage(params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.Stream(page, limit, 0), nil
}

//Lists SupportingDocumentType records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) SupportingDocumentTypeList(params *ListSupportingDocumentTypeParams, limit int) ([]interface{}, error) {
	if params.PageSize == nil {
		params.SetPageSize(0)
	}
	params.SetPageSize(c.requestHandler.ReadLimits(*params.PageSize, limit))
	page, err := c.SupportingDocumentTypePage(params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.List(page, limit, 0), nil
}
