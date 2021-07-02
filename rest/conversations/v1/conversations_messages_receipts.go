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

// Fetch the delivery and read receipts of the conversation message
func (c *ApiService) FetchConversationMessageReceipt(ConversationSid string, MessageSid string, Sid string) (*ConversationsV1ConversationConversationMessageConversationMessageReceipt, error) {
	path := "/v1/Conversations/{ConversationSid}/Messages/{MessageSid}/Receipts/{Sid}"
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"MessageSid"+"}", MessageSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ConversationConversationMessageConversationMessageReceipt{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListConversationMessageReceipt'
type ListConversationMessageReceiptParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListConversationMessageReceiptParams) SetPageSize(PageSize int) *ListConversationMessageReceiptParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all delivery and read receipts of the conversation message
func (c *ApiService) ListConversationMessageReceipt(ConversationSid string, MessageSid string, params *ListConversationMessageReceiptParams) (*ListConversationMessageReceiptResponse, error) {
	path := "/v1/Conversations/{ConversationSid}/Messages/{MessageSid}/Receipts"
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"MessageSid"+"}", MessageSid, -1)

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

	ps := &ListConversationMessageReceiptResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of ConversationMessageReceipt records from the API. Request is executed immediately.
func (c *ApiService) ConversationMessageReceiptPage(ConversationSid string, MessageSid string, params *ListConversationMessageReceiptParams, pageToken string, pageNumber string) (*client.Page, error) {
	path := "/v1/Conversations/{ConversationSid}/Messages/{MessageSid}/Receipts"
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"MessageSid"+"}", MessageSid, -1)

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

//Streams ConversationMessageReceipt records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) ConversationMessageReceiptStream(ConversationSid string, MessageSid string, params *ListConversationMessageReceiptParams, limit int) (chan map[string]interface{}, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	page, err := c.ConversationMessageReceiptPage(ConversationSid, MessageSid, params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.Stream(page, limit, 0), nil
}

//Lists ConversationMessageReceipt records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) ConversationMessageReceiptList(ConversationSid string, MessageSid string, params *ListConversationMessageReceiptParams, limit int) ([]interface{}, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	page, err := c.ConversationMessageReceiptPage(ConversationSid, MessageSid, params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.List(page, limit, 0), nil
}
