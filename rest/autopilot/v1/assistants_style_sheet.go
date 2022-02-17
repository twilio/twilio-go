/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"net/url"
	"strings"
)

// Returns Style sheet JSON object for the Assistant
func (c *ApiService) FetchStyleSheet(AssistantSid string) (*AutopilotV1StyleSheet, error) {
	path := "/v1/Assistants/{AssistantSid}/StyleSheet"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1StyleSheet{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateStyleSheet'
type UpdateStyleSheetParams struct {
	// The JSON string that describes the style sheet object.
	StyleSheet *map[string]interface{} `json:"StyleSheet,omitempty"`
}

func (params *UpdateStyleSheetParams) SetStyleSheet(StyleSheet map[string]interface{}) *UpdateStyleSheetParams {
	params.StyleSheet = &StyleSheet
	return params
}

// Updates the style sheet for an Assistant identified by &#x60;assistant_sid&#x60;.
func (c *ApiService) UpdateStyleSheet(AssistantSid string, params *UpdateStyleSheetParams) (*AutopilotV1StyleSheet, error) {
	path := "/v1/Assistants/{AssistantSid}/StyleSheet"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.StyleSheet != nil {
		v, err := json.Marshal(params.StyleSheet)

		if err != nil {
			return nil, err
		}

		data.Set("StyleSheet", string(v))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1StyleSheet{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
