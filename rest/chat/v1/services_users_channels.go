/*
 * Twilio - Chat
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

// Optional parameters for the method 'ListUserChannel'
type ListUserChannelParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListUserChannelParams) SetPageSize(PageSize int) *ListUserChannelParams {
	params.PageSize = &PageSize
	return params
}

// List all Channels for a given User.
func (c *ApiService) ListUserChannel(ServiceSid string, UserSid string, params *ListUserChannelParams) (*ListUserChannelResponse, error) {
	path := "/v1/Services/{ServiceSid}/Users/{UserSid}/Channels"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)

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

	ps := &ListUserChannelResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of UserChannel records from the API. Request is executed immediately.
func (c *ApiService) UserChannelPage(ServiceSid string, UserSid string, params *ListUserChannelParams, pageToken string, pageNumber string) (*client.Page, error) {
	path := "/v1/Services/{ServiceSid}/Users/{UserSid}/Channels"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)

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

//Streams UserChannel records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) UserChannelStream(ServiceSid string, UserSid string, params *ListUserChannelParams, limit int) (chan map[string]interface{}, error) {
	if params.PageSize == nil {
		params.SetPageSize(0)
	}
	params.SetPageSize(c.requestHandler.ReadLimits(*params.PageSize, limit))
	page, err := c.UserChannelPage(ServiceSid, UserSid, params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.Stream(page, limit, 0), nil
}

//Lists UserChannel records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) UserChannelList(ServiceSid string, UserSid string, params *ListUserChannelParams, limit int) ([]interface{}, error) {
	if params.PageSize == nil {
		params.SetPageSize(0)
	}
	params.SetPageSize(c.requestHandler.ReadLimits(*params.PageSize, limit))
	page, err := c.UserChannelPage(ServiceSid, UserSid, params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.List(page, limit, 0), nil
}
