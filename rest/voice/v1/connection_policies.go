/*
 * Twilio - Voice
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

// Optional parameters for the method 'CreateConnectionPolicy'
type CreateConnectionPolicyParams struct {
	// A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateConnectionPolicyParams) SetFriendlyName(FriendlyName string) *CreateConnectionPolicyParams {
	params.FriendlyName = &FriendlyName
	return params
}

func (c *ApiService) CreateConnectionPolicy(params *CreateConnectionPolicyParams) (*VoiceV1ConnectionPolicy, error) {
	path := "/v1/ConnectionPolicies"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1ConnectionPolicy{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteConnectionPolicy(Sid string) error {
	path := "/v1/ConnectionPolicies/{Sid}"
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

func (c *ApiService) FetchConnectionPolicy(Sid string) (*VoiceV1ConnectionPolicy, error) {
	path := "/v1/ConnectionPolicies/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1ConnectionPolicy{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListConnectionPolicy'
type ListConnectionPolicyParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListConnectionPolicyParams) SetPageSize(PageSize int) *ListConnectionPolicyParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListConnectionPolicy(params *ListConnectionPolicyParams) (*ListConnectionPolicyResponse, error) {
	path := "/v1/ConnectionPolicies"

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

	ps := &ListConnectionPolicyResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of ConnectionPolicy records from the API. Request is executed immediately.
func (c *ApiService) ConnectionPolicyPage(params *ListConnectionPolicyParams, pageToken string, pageNumber string) (*client.Page, error) {
	path := "/v1/ConnectionPolicies"

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

//Streams ConnectionPolicy records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) ConnectionPolicyStream(params *ListConnectionPolicyParams, limit int) (chan map[string]interface{}, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	page, err := c.ConnectionPolicyPage(params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.Stream(page, limit, 0), nil
}

//Lists ConnectionPolicy records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) ConnectionPolicyList(params *ListConnectionPolicyParams, limit int) ([]interface{}, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	page, err := c.ConnectionPolicyPage(params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.List(page, limit, 0), nil
}

// Optional parameters for the method 'UpdateConnectionPolicy'
type UpdateConnectionPolicyParams struct {
	// A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateConnectionPolicyParams) SetFriendlyName(FriendlyName string) *UpdateConnectionPolicyParams {
	params.FriendlyName = &FriendlyName
	return params
}

func (c *ApiService) UpdateConnectionPolicy(Sid string, params *UpdateConnectionPolicyParams) (*VoiceV1ConnectionPolicy, error) {
	path := "/v1/ConnectionPolicies/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1ConnectionPolicy{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
