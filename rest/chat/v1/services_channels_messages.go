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

// Optional parameters for the method 'CreateMessage'
type CreateMessageParams struct {
	// A valid JSON string that contains application-specific data.
	Attributes *string `json:"Attributes,omitempty"`
	// The message to send to the channel. Can also be an empty string or `null`, which sets the value as an empty string. You can send structured data in the body by serializing it as a string.
	Body *string `json:"Body,omitempty"`
	// The [identity](https://www.twilio.com/docs/api/chat/guides/identity) of the new message's author. The default value is `system`.
	From *string `json:"From,omitempty"`
}

func (params *CreateMessageParams) SetAttributes(Attributes string) *CreateMessageParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateMessageParams) SetBody(Body string) *CreateMessageParams {
	params.Body = &Body
	return params
}
func (params *CreateMessageParams) SetFrom(From string) *CreateMessageParams {
	params.From = &From
	return params
}

func (c *ApiService) CreateMessage(ServiceSid string, ChannelSid string, params *CreateMessageParams) (*ChatV1ServiceChannelMessage, error) {
	path := "/v1/Services/{ServiceSid}/Channels/{ChannelSid}/Messages"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.Body != nil {
		data.Set("Body", *params.Body)
	}
	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ChatV1ServiceChannelMessage{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteMessage(ServiceSid string, ChannelSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Channels/{ChannelSid}/Messages/{Sid}"
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

func (c *ApiService) FetchMessage(ServiceSid string, ChannelSid string, Sid string) (*ChatV1ServiceChannelMessage, error) {
	path := "/v1/Services/{ServiceSid}/Channels/{ChannelSid}/Messages/{Sid}"
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

	ps := &ChatV1ServiceChannelMessage{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListMessage'
type ListMessageParams struct {
	// The sort order of the returned messages. Can be: `asc` (ascending) or `desc` (descending) with `asc` as the default.
	Order *string `json:"Order,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListMessageParams) SetOrder(Order string) *ListMessageParams {
	params.Order = &Order
	return params
}
func (params *ListMessageParams) SetPageSize(PageSize int) *ListMessageParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListMessage(ServiceSid string, ChannelSid string, params *ListMessageParams) (*ListMessageResponse, error) {
	path := "/v1/Services/{ServiceSid}/Channels/{ChannelSid}/Messages"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Order != nil {
		data.Set("Order", *params.Order)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMessageResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of Message records from the API. Request is executed immediately.
func (c *ApiService) MessagePage(ServiceSid string, ChannelSid string, params *ListMessageParams, pageToken string, pageNumber string) (*client.Page, error) {
	path := "/v1/Services/{ServiceSid}/Channels/{ChannelSid}/Messages"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Order != nil {
		data.Set("Order", *params.Order)
	}
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

//Streams Message records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) MessageStream(ServiceSid string, ChannelSid string, params *ListMessageParams, limit int) (chan map[string]interface{}, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	page, err := c.MessagePage(ServiceSid, ChannelSid, params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.Stream(page, limit, 0), nil
}

//Lists Message records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) MessageList(ServiceSid string, ChannelSid string, params *ListMessageParams, limit int) ([]interface{}, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	page, err := c.MessagePage(ServiceSid, ChannelSid, params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.List(page, limit, 0), nil
}

// Optional parameters for the method 'UpdateMessage'
type UpdateMessageParams struct {
	// A valid JSON string that contains application-specific data.
	Attributes *string `json:"Attributes,omitempty"`
	// The message to send to the channel. Can also be an empty string or `null`, which sets the value as an empty string. You can send structured data in the body by serializing it as a string.
	Body *string `json:"Body,omitempty"`
}

func (params *UpdateMessageParams) SetAttributes(Attributes string) *UpdateMessageParams {
	params.Attributes = &Attributes
	return params
}
func (params *UpdateMessageParams) SetBody(Body string) *UpdateMessageParams {
	params.Body = &Body
	return params
}

func (c *ApiService) UpdateMessage(ServiceSid string, ChannelSid string, Sid string, params *UpdateMessageParams) (*ChatV1ServiceChannelMessage, error) {
	path := "/v1/Services/{ServiceSid}/Channels/{ChannelSid}/Messages/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.Body != nil {
		data.Set("Body", *params.Body)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ChatV1ServiceChannelMessage{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
