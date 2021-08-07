/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.1
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

// Optional parameters for the method 'CreateChannel'
type CreateChannelParams struct {
	//
	Attributes *string `json:"Attributes,omitempty"`
	//
	FriendlyName *string `json:"FriendlyName,omitempty"`
	//
	Type *string `json:"Type,omitempty"`
	//
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *CreateChannelParams) SetAttributes(Attributes string) *CreateChannelParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateChannelParams) SetFriendlyName(FriendlyName string) *CreateChannelParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateChannelParams) SetType(Type string) *CreateChannelParams {
	params.Type = &Type
	return params
}
func (params *CreateChannelParams) SetUniqueName(UniqueName string) *CreateChannelParams {
	params.UniqueName = &UniqueName
	return params
}

func (c *ApiService) CreateChannel(ServiceSid string, params *CreateChannelParams) (*IpMessagingV1ServiceChannel, error) {
	path := "/v1/Services/{ServiceSid}/Channels"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Type != nil {
		data.Set("Type", *params.Type)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV1ServiceChannel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteChannel(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Channels/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
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

func (c *ApiService) FetchChannel(ServiceSid string, Sid string) (*IpMessagingV1ServiceChannel, error) {
	path := "/v1/Services/{ServiceSid}/Channels/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV1ServiceChannel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListChannel'
type ListChannelParams struct {
	//
	Type *[]string `json:"Type,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListChannelParams) SetType(Type []string) *ListChannelParams {
	params.Type = &Type
	return params
}
func (params *ListChannelParams) SetPageSize(PageSize int) *ListChannelParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListChannelParams) SetLimit(Limit int) *ListChannelParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Channel records from the API. Request is executed immediately.
func (c *ApiService) PageChannel(ServiceSid string, params *ListChannelParams, pageToken string, pageNumber string) (*ListChannelResponse, error) {
	path := "/v1/Services/{ServiceSid}/Channels"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Type != nil {
		for _, item := range *params.Type {
			data.Add("Type", item)
		}
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

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

	ps := &ListChannelResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Channel records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListChannel(ServiceSid string, params *ListChannelParams) ([]IpMessagingV1ServiceChannel, error) {
	if params == nil {
		params = &ListChannelParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageChannel(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []IpMessagingV1ServiceChannel

	for response != nil {
		records = append(records, response.Channels...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListChannelResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListChannelResponse)
	}

	return records, err
}

// Streams Channel records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamChannel(ServiceSid string, params *ListChannelParams) (chan IpMessagingV1ServiceChannel, error) {
	if params == nil {
		params = &ListChannelParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageChannel(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan IpMessagingV1ServiceChannel, 1)

	go func() {
		for response != nil {
			for item := range response.Channels {
				channel <- response.Channels[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListChannelResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListChannelResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListChannelResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListChannelResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateChannel'
type UpdateChannelParams struct {
	//
	Attributes *string `json:"Attributes,omitempty"`
	//
	FriendlyName *string `json:"FriendlyName,omitempty"`
	//
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *UpdateChannelParams) SetAttributes(Attributes string) *UpdateChannelParams {
	params.Attributes = &Attributes
	return params
}
func (params *UpdateChannelParams) SetFriendlyName(FriendlyName string) *UpdateChannelParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateChannelParams) SetUniqueName(UniqueName string) *UpdateChannelParams {
	params.UniqueName = &UniqueName
	return params
}

func (c *ApiService) UpdateChannel(ServiceSid string, Sid string, params *UpdateChannelParams) (*IpMessagingV1ServiceChannel, error) {
	path := "/v1/Services/{ServiceSid}/Channels/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV1ServiceChannel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
