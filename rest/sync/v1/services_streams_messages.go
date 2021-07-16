/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"net/url"

	"strings"
)

// Optional parameters for the method 'CreateStreamMessage'
type CreateStreamMessageParams struct {
	// A JSON string that represents an arbitrary, schema-less object that makes up the Stream Message body. Can be up to 4 KiB in length.
	Data *map[string]interface{} `json:"Data,omitempty"`
}

func (params *CreateStreamMessageParams) SetData(Data map[string]interface{}) *CreateStreamMessageParams {
	params.Data = &Data
	return params
}

// Create a new Stream Message.
func (c *ApiService) CreateStreamMessage(ServiceSid string, StreamSid string, params *CreateStreamMessageParams) (*SyncV1ServiceSyncStreamStreamMessage, error) {
	path := "/v1/Services/{ServiceSid}/Streams/{StreamSid}/Messages"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"StreamSid"+"}", StreamSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Data != nil {
		v, err := json.Marshal(params.Data)

		if err != nil {
			return nil, err
		}

		data.Set("Data", string(v))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1ServiceSyncStreamStreamMessage{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
