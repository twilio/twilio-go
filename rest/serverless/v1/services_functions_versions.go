/*
 * Twilio - Serverless
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

// Retrieve a specific Function Version resource.
func (c *ApiService) FetchFunctionVersion(ServiceSid string, FunctionSid string, Sid string) (*ServerlessV1ServiceFunctionFunctionVersion, error) {
	path := "/v1/Services/{ServiceSid}/Functions/{FunctionSid}/Versions/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"FunctionSid"+"}", FunctionSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ServerlessV1ServiceFunctionFunctionVersion{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListFunctionVersion'
type ListFunctionVersionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListFunctionVersionParams) SetPageSize(PageSize int) *ListFunctionVersionParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all Function Version resources.
func (c *ApiService) ListFunctionVersion(ServiceSid string, FunctionSid string, params *ListFunctionVersionParams) (*ListFunctionVersionResponse, error) {
	path := "/v1/Services/{ServiceSid}/Functions/{FunctionSid}/Versions"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"FunctionSid"+"}", FunctionSid, -1)

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

	ps := &ListFunctionVersionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of FunctionVersion records from the API. Request is executed immediately.
func (c *ApiService) FunctionVersionPage(ServiceSid string, FunctionSid string, params *ListFunctionVersionParams, pageToken string, pageNumber string) (*client.Page, error) {
	path := "/v1/Services/{ServiceSid}/Functions/{FunctionSid}/Versions"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"FunctionSid"+"}", FunctionSid, -1)

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

//Streams FunctionVersion records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) FunctionVersionStream(ServiceSid string, FunctionSid string, params *ListFunctionVersionParams, limit int) (chan map[string]interface{}, error) {
	if params.PageSize == nil {
		params.SetPageSize(0)
	}
	params.SetPageSize(c.requestHandler.ReadLimits(*params.PageSize, limit))
	page, err := c.FunctionVersionPage(ServiceSid, FunctionSid, params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.Stream(page, limit, 0), nil
}

//Lists FunctionVersion records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) FunctionVersionList(ServiceSid string, FunctionSid string, params *ListFunctionVersionParams, limit int) ([]interface{}, error) {
	if params.PageSize == nil {
		params.SetPageSize(0)
	}
	params.SetPageSize(c.requestHandler.ReadLimits(*params.PageSize, limit))
	page, err := c.FunctionVersionPage(ServiceSid, FunctionSid, params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.List(page, limit, 0), nil
}
