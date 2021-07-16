/*
 * Twilio - Messaging
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
)

// Optional parameters for the method 'CreateExternalCampaign'
type CreateExternalCampaignParams struct {
	// ID of the preregistered campaign.
	CampaignId *string `json:"CampaignId,omitempty"`
	// The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) that the resource is associated with.
	MessagingServiceSid *string `json:"MessagingServiceSid,omitempty"`
}

func (params *CreateExternalCampaignParams) SetCampaignId(CampaignId string) *CreateExternalCampaignParams {
	params.CampaignId = &CampaignId
	return params
}
func (params *CreateExternalCampaignParams) SetMessagingServiceSid(MessagingServiceSid string) *CreateExternalCampaignParams {
	params.MessagingServiceSid = &MessagingServiceSid
	return params
}

func (c *ApiService) CreateExternalCampaign(params *CreateExternalCampaignParams) (*MessagingV1ExternalCampaign, error) {
	path := "/v1/Services/PreregisteredUsa2p"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CampaignId != nil {
		data.Set("CampaignId", *params.CampaignId)
	}
	if params != nil && params.MessagingServiceSid != nil {
		data.Set("MessagingServiceSid", *params.MessagingServiceSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1ExternalCampaign{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
