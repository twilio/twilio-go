/*
 * Twilio - Ip_messaging
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

func (c *ApiService) DeleteUserBinding(ServiceSid string, UserSid string, Sid string) error {
	path := "/v2/Services/{ServiceSid}/Users/{UserSid}/Bindings/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)
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

func (c *ApiService) FetchUserBinding(ServiceSid string, UserSid string, Sid string) (*IpMessagingV2ServiceUserUserBinding, error) {
	path := "/v2/Services/{ServiceSid}/Users/{UserSid}/Bindings/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV2ServiceUserUserBinding{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListUserBinding'
type ListUserBindingParams struct {
	//
	BindingType *[]string `json:"BindingType,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListUserBindingParams) SetBindingType(BindingType []string) *ListUserBindingParams {
	params.BindingType = &BindingType
	return params
}
func (params *ListUserBindingParams) SetPageSize(PageSize int) *ListUserBindingParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListUserBinding(ServiceSid string, UserSid string, params *ListUserBindingParams) (*ListUserBindingResponse, error) {
	path := "/v2/Services/{ServiceSid}/Users/{UserSid}/Bindings"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.BindingType != nil {
		for _, item := range *params.BindingType {
			data.Add("BindingType", item)
		}
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListUserBindingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of UserBinding records from the API. Request is executed immediately.
func (c *ApiService) UserBindingPage(ServiceSid string, UserSid string, params *ListUserBindingParams, pageToken string, pageNumber string, pageSize string) *client.Page {
	path := "/v2/Services/{ServiceSid}/Users/{UserSid}/Bindings"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.BindingType != nil {
		for _, item := range *params.BindingType {
			data.Add("BindingType", item)
		}
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	data.Set("PageToken", pageToken)
	data.Set("PageNumber", pageNumber)
	data.Set("PageSize", pageSize)

	response, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil
	}

	return client.NewPage(c.baseURL, response)
}

//Streams UserBinding records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) UserBindingStream(ServiceSid string, UserSid string, params *ListUserBindingParams, meta client.PaginationData) chan map[string]interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.UserBindingPage(ServiceSid, UserSid, params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.Stream(page, limits.Limit, limits.PageLimit)
}

//Lists UserBinding records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) UserBindingList(ServiceSid string, UserSid string, params *ListUserBindingParams, meta client.PaginationData) []interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.UserBindingPage(ServiceSid, UserSid, params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.List(page, limits.Limit, limits.PageLimit)
}
