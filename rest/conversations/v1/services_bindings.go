/*
 * Twilio - Conversations
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

// Remove a push notification binding from the conversation service
func (c *ApiService) DeleteServiceBinding(ChatServiceSid string, Sid string) error {
	path := "/v1/Services/{ChatServiceSid}/Bindings/{Sid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
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

// Fetch a push notification binding from the conversation service
func (c *ApiService) FetchServiceBinding(ChatServiceSid string, Sid string) (*ConversationsV1ServiceServiceBinding, error) {
	path := "/v1/Services/{ChatServiceSid}/Bindings/{Sid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceServiceBinding{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListServiceBinding'
type ListServiceBindingParams struct {
	// The push technology used by the Binding resources to read.  Can be: `apn`, `gcm`, or `fcm`.  See [push notification configuration](https://www.twilio.com/docs/chat/push-notification-configuration) for more info.
	BindingType *[]string `json:"BindingType,omitempty"`
	// The identity of a [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource) this binding belongs to. See [access tokens](https://www.twilio.com/docs/conversations/create-tokens) for more details.
	Identity *[]string `json:"Identity,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListServiceBindingParams) SetBindingType(BindingType []string) *ListServiceBindingParams {
	params.BindingType = &BindingType
	return params
}
func (params *ListServiceBindingParams) SetIdentity(Identity []string) *ListServiceBindingParams {
	params.Identity = &Identity
	return params
}
func (params *ListServiceBindingParams) SetPageSize(PageSize int) *ListServiceBindingParams {
	params.PageSize = &PageSize
	return params
}

//Retrieve a single page of ServiceBinding records from the API. Request is executed immediately.
func (c *ApiService) PageServiceBinding(ChatServiceSid string, params *ListServiceBindingParams, pageToken string, pageNumber string) (*ListServiceBindingResponse, error) {
	path := "/v1/Services/{ChatServiceSid}/Bindings"

	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.BindingType != nil {
		for _, item := range *params.BindingType {
			data.Add("BindingType", item)
		}
	}
	if params != nil && params.Identity != nil {
		for _, item := range *params.Identity {
			data.Add("Identity", item)
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

	ps := &ListServiceBindingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Lists ServiceBinding records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListServiceBinding(ChatServiceSid string, params *ListServiceBindingParams, limit *int) ([]*ListServiceBindingResponse, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageServiceBinding(ChatServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []*ListServiceBindingResponse

	for response != nil {
		records = append(records, response)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, limit, c.getNextListServiceBindingResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListServiceBindingResponse)
	}

	return records, err
}

//Streams ServiceBinding records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamServiceBinding(ChatServiceSid string, params *ListServiceBindingParams, limit *int) (chan *ListServiceBindingResponse, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageServiceBinding(ChatServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan *ListServiceBindingResponse, 1)

	go func() {
		for response != nil {
			channel <- response

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, limit, c.getNextListServiceBindingResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListServiceBindingResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListServiceBindingResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListServiceBindingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
