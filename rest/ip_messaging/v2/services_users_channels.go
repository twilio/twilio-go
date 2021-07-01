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
	"time"

	"github.com/twilio/twilio-go/client"
)

func (c *ApiService) DeleteUserChannel(ServiceSid string, UserSid string, ChannelSid string) error {
	path := "/v2/Services/{ServiceSid}/Users/{UserSid}/Channels/{ChannelSid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ApiService) FetchUserChannel(ServiceSid string, UserSid string, ChannelSid string) (*IpMessagingV2ServiceUserUserChannel, error) {
	path := "/v2/Services/{ServiceSid}/Users/{UserSid}/Channels/{ChannelSid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV2ServiceUserUserChannel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListUserChannel'
type ListUserChannelParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListUserChannelParams) SetPageSize(PageSize int) *ListUserChannelParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListUserChannel(ServiceSid string, UserSid string, params *ListUserChannelParams) (*ListUserChannelResponse, error) {
	path := "/v2/Services/{ServiceSid}/Users/{UserSid}/Channels"
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
func (c *ApiService) UserChannelPage(ServiceSid string, UserSid string, params *ListUserChannelParams, pageToken string, pageNumber string, pageSize string) *client.Page {
	path := "/v2/Services/{ServiceSid}/Users/{UserSid}/Channels"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

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

//Streams UserChannel records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) UserChannelStream(ServiceSid string, UserSid string, params *ListUserChannelParams, meta client.PaginationData) chan map[string]interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.UserChannelPage(ServiceSid, UserSid, params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.Stream(page, limits.Limit, limits.PageLimit)
}

//Lists UserChannel records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) UserChannelList(ServiceSid string, UserSid string, params *ListUserChannelParams, meta client.PaginationData) []interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.UserChannelPage(ServiceSid, UserSid, params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.List(page, limits.Limit, limits.PageLimit)
}

// Optional parameters for the method 'UpdateUserChannel'
type UpdateUserChannelParams struct {
	//
	LastConsumedMessageIndex *int `json:"LastConsumedMessageIndex,omitempty"`
	//
	LastConsumptionTimestamp *time.Time `json:"LastConsumptionTimestamp,omitempty"`
	//
	NotificationLevel *string `json:"NotificationLevel,omitempty"`
}

func (params *UpdateUserChannelParams) SetLastConsumedMessageIndex(LastConsumedMessageIndex int) *UpdateUserChannelParams {
	params.LastConsumedMessageIndex = &LastConsumedMessageIndex
	return params
}
func (params *UpdateUserChannelParams) SetLastConsumptionTimestamp(LastConsumptionTimestamp time.Time) *UpdateUserChannelParams {
	params.LastConsumptionTimestamp = &LastConsumptionTimestamp
	return params
}
func (params *UpdateUserChannelParams) SetNotificationLevel(NotificationLevel string) *UpdateUserChannelParams {
	params.NotificationLevel = &NotificationLevel
	return params
}

func (c *ApiService) UpdateUserChannel(ServiceSid string, UserSid string, ChannelSid string, params *UpdateUserChannelParams) (*IpMessagingV2ServiceUserUserChannel, error) {
	path := "/v2/Services/{ServiceSid}/Users/{UserSid}/Channels/{ChannelSid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.LastConsumedMessageIndex != nil {
		data.Set("LastConsumedMessageIndex", fmt.Sprint(*params.LastConsumedMessageIndex))
	}
	if params != nil && params.LastConsumptionTimestamp != nil {
		data.Set("LastConsumptionTimestamp", fmt.Sprint((*params.LastConsumptionTimestamp).Format(time.RFC3339)))
	}
	if params != nil && params.NotificationLevel != nil {
		data.Set("NotificationLevel", *params.NotificationLevel)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV2ServiceUserUserChannel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
