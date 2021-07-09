/*
 * Twilio - Studio
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
	"net/url"
)

// Optional parameters for the method 'UpdateFlowValidate'
type UpdateFlowValidateParams struct {
	// Description of change made in the revision.
	CommitMessage *string `json:"CommitMessage,omitempty"`
	// JSON representation of flow definition.
	Definition *map[string]interface{} `json:"Definition,omitempty"`
	// The string that you assigned to describe the Flow.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The status of the Flow. Can be: `draft` or `published`.
	Status *string `json:"Status,omitempty"`
}

func (params *UpdateFlowValidateParams) SetCommitMessage(CommitMessage string) *UpdateFlowValidateParams {
	params.CommitMessage = &CommitMessage
	return params
}
func (params *UpdateFlowValidateParams) SetDefinition(Definition map[string]interface{}) *UpdateFlowValidateParams {
	params.Definition = &Definition
	return params
}
func (params *UpdateFlowValidateParams) SetFriendlyName(FriendlyName string) *UpdateFlowValidateParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateFlowValidateParams) SetStatus(Status string) *UpdateFlowValidateParams {
	params.Status = &Status
	return params
}

// Validate flow JSON definition
func (c *ApiService) UpdateFlowValidate(params *UpdateFlowValidateParams) (*StudioV2FlowValidate, error) {
	path := "/v2/Flows/Validate"

	data := url.Values{}
	if params != nil && params.CommitMessage != nil {
		data.Set("CommitMessage", *params.CommitMessage)
	}
	if params != nil && params.Definition != nil {
		v, err := json.Marshal(params.Definition)

		if err != nil {
			return nil, err
		}

		data.Set("Definition", string(v))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV2FlowValidate{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
