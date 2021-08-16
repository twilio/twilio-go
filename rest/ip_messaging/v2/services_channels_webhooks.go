/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.0
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

// Optional parameters for the method 'CreateChannelWebhook'
type CreateChannelWebhookParams struct {
	//
	ConfigurationFilters *[]string `json:"Configuration.Filters,omitempty"`
	//
	ConfigurationFlowSid *string `json:"Configuration.FlowSid,omitempty"`
	//
	ConfigurationMethod *string `json:"Configuration.Method,omitempty"`
	//
	ConfigurationRetryCount *int `json:"Configuration.RetryCount,omitempty"`
	//
	ConfigurationTriggers *[]string `json:"Configuration.Triggers,omitempty"`
	//
	ConfigurationUrl *string `json:"Configuration.Url,omitempty"`
	//
	Type *string `json:"Type,omitempty"`
}

func (params *CreateChannelWebhookParams) SetConfigurationFilters(ConfigurationFilters []string) *CreateChannelWebhookParams {
	params.ConfigurationFilters = &ConfigurationFilters
	return params
}
func (params *CreateChannelWebhookParams) SetConfigurationFlowSid(ConfigurationFlowSid string) *CreateChannelWebhookParams {
	params.ConfigurationFlowSid = &ConfigurationFlowSid
	return params
}
func (params *CreateChannelWebhookParams) SetConfigurationMethod(ConfigurationMethod string) *CreateChannelWebhookParams {
	params.ConfigurationMethod = &ConfigurationMethod
	return params
}
func (params *CreateChannelWebhookParams) SetConfigurationRetryCount(ConfigurationRetryCount int) *CreateChannelWebhookParams {
	params.ConfigurationRetryCount = &ConfigurationRetryCount
	return params
}
func (params *CreateChannelWebhookParams) SetConfigurationTriggers(ConfigurationTriggers []string) *CreateChannelWebhookParams {
	params.ConfigurationTriggers = &ConfigurationTriggers
	return params
}
func (params *CreateChannelWebhookParams) SetConfigurationUrl(ConfigurationUrl string) *CreateChannelWebhookParams {
	params.ConfigurationUrl = &ConfigurationUrl
	return params
}
func (params *CreateChannelWebhookParams) SetType(Type string) *CreateChannelWebhookParams {
	params.Type = &Type
	return params
}

func (c *ApiService) CreateChannelWebhook(ServiceSid string, ChannelSid string, params *CreateChannelWebhookParams) (*IpMessagingV2ChannelWebhook, error) {
	path := "/v2/Services/{ServiceSid}/Channels/{ChannelSid}/Webhooks"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

	data := url.Values{}
	if params != nil && params.ConfigurationFilters != nil {
		for _, item := range *params.ConfigurationFilters {
			data.Add("Configuration.Filters", item)
		}
	}
	if params != nil && params.ConfigurationFlowSid != nil {
		data.Set("Configuration.FlowSid", *params.ConfigurationFlowSid)
	}
	if params != nil && params.ConfigurationMethod != nil {
		data.Set("Configuration.Method", *params.ConfigurationMethod)
	}
	if params != nil && params.ConfigurationRetryCount != nil {
		data.Set("Configuration.RetryCount", fmt.Sprint(*params.ConfigurationRetryCount))
	}
	if params != nil && params.ConfigurationTriggers != nil {
		for _, item := range *params.ConfigurationTriggers {
			data.Add("Configuration.Triggers", item)
		}
	}
	if params != nil && params.ConfigurationUrl != nil {
		data.Set("Configuration.Url", *params.ConfigurationUrl)
	}
	if params != nil && params.Type != nil {
		data.Set("Type", *params.Type)
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV2ChannelWebhook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteChannelWebhook(ServiceSid string, ChannelSid string, Sid string) error {
	path := "/v2/Services/{ServiceSid}/Channels/{ChannelSid}/Webhooks/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)
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

func (c *ApiService) FetchChannelWebhook(ServiceSid string, ChannelSid string, Sid string) (*IpMessagingV2ChannelWebhook, error) {
	path := "/v2/Services/{ServiceSid}/Channels/{ChannelSid}/Webhooks/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV2ChannelWebhook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListChannelWebhook'
type ListChannelWebhookParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListChannelWebhookParams) SetPageSize(PageSize int) *ListChannelWebhookParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListChannelWebhookParams) SetLimit(Limit int) *ListChannelWebhookParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ChannelWebhook records from the API. Request is executed immediately.
func (c *ApiService) PageChannelWebhook(ServiceSid string, ChannelSid string, params *ListChannelWebhookParams, pageToken string, pageNumber string) (*ListChannelWebhookResponse, error) {
	path := "/v2/Services/{ServiceSid}/Channels/{ChannelSid}/Webhooks"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

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

	ps := &ListChannelWebhookResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ChannelWebhook records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListChannelWebhook(ServiceSid string, ChannelSid string, params *ListChannelWebhookParams) ([]IpMessagingV2ChannelWebhook, error) {
	if params == nil {
		params = &ListChannelWebhookParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageChannelWebhook(ServiceSid, ChannelSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []IpMessagingV2ChannelWebhook

	for response != nil {
		records = append(records, response.Webhooks...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListChannelWebhookResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListChannelWebhookResponse)
	}

	return records, err
}

// Streams ChannelWebhook records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamChannelWebhook(ServiceSid string, ChannelSid string, params *ListChannelWebhookParams) (chan IpMessagingV2ChannelWebhook, error) {
	if params == nil {
		params = &ListChannelWebhookParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageChannelWebhook(ServiceSid, ChannelSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan IpMessagingV2ChannelWebhook, 1)

	go func() {
		for response != nil {
			for item := range response.Webhooks {
				channel <- response.Webhooks[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListChannelWebhookResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListChannelWebhookResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListChannelWebhookResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListChannelWebhookResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateChannelWebhook'
type UpdateChannelWebhookParams struct {
	//
	ConfigurationFilters *[]string `json:"Configuration.Filters,omitempty"`
	//
	ConfigurationFlowSid *string `json:"Configuration.FlowSid,omitempty"`
	//
	ConfigurationMethod *string `json:"Configuration.Method,omitempty"`
	//
	ConfigurationRetryCount *int `json:"Configuration.RetryCount,omitempty"`
	//
	ConfigurationTriggers *[]string `json:"Configuration.Triggers,omitempty"`
	//
	ConfigurationUrl *string `json:"Configuration.Url,omitempty"`
}

func (params *UpdateChannelWebhookParams) SetConfigurationFilters(ConfigurationFilters []string) *UpdateChannelWebhookParams {
	params.ConfigurationFilters = &ConfigurationFilters
	return params
}
func (params *UpdateChannelWebhookParams) SetConfigurationFlowSid(ConfigurationFlowSid string) *UpdateChannelWebhookParams {
	params.ConfigurationFlowSid = &ConfigurationFlowSid
	return params
}
func (params *UpdateChannelWebhookParams) SetConfigurationMethod(ConfigurationMethod string) *UpdateChannelWebhookParams {
	params.ConfigurationMethod = &ConfigurationMethod
	return params
}
func (params *UpdateChannelWebhookParams) SetConfigurationRetryCount(ConfigurationRetryCount int) *UpdateChannelWebhookParams {
	params.ConfigurationRetryCount = &ConfigurationRetryCount
	return params
}
func (params *UpdateChannelWebhookParams) SetConfigurationTriggers(ConfigurationTriggers []string) *UpdateChannelWebhookParams {
	params.ConfigurationTriggers = &ConfigurationTriggers
	return params
}
func (params *UpdateChannelWebhookParams) SetConfigurationUrl(ConfigurationUrl string) *UpdateChannelWebhookParams {
	params.ConfigurationUrl = &ConfigurationUrl
	return params
}

func (c *ApiService) UpdateChannelWebhook(ServiceSid string, ChannelSid string, Sid string, params *UpdateChannelWebhookParams) (*IpMessagingV2ChannelWebhook, error) {
	path := "/v2/Services/{ServiceSid}/Channels/{ChannelSid}/Webhooks/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	if params != nil && params.ConfigurationFilters != nil {
		for _, item := range *params.ConfigurationFilters {
			data.Add("Configuration.Filters", item)
		}
	}
	if params != nil && params.ConfigurationFlowSid != nil {
		data.Set("Configuration.FlowSid", *params.ConfigurationFlowSid)
	}
	if params != nil && params.ConfigurationMethod != nil {
		data.Set("Configuration.Method", *params.ConfigurationMethod)
	}
	if params != nil && params.ConfigurationRetryCount != nil {
		data.Set("Configuration.RetryCount", fmt.Sprint(*params.ConfigurationRetryCount))
	}
	if params != nil && params.ConfigurationTriggers != nil {
		for _, item := range *params.ConfigurationTriggers {
			data.Add("Configuration.Triggers", item)
		}
	}
	if params != nil && params.ConfigurationUrl != nil {
		data.Set("Configuration.Url", *params.ConfigurationUrl)
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV2ChannelWebhook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
