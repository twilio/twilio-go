/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AccountQueue struct for ApiV2010AccountQueue
type ApiV2010AccountQueue struct {
	AccountSid      *string `json:"account_sid,omitempty"`
	AverageWaitTime *int32  `json:"average_wait_time,omitempty"`
	CurrentSize     *int32  `json:"current_size,omitempty"`
	DateCreated     *string `json:"date_created,omitempty"`
	DateUpdated     *string `json:"date_updated,omitempty"`
	FriendlyName    *string `json:"friendly_name,omitempty"`
	MaxSize         *int32  `json:"max_size,omitempty"`
	Sid             *string `json:"sid,omitempty"`
	Uri             *string `json:"uri,omitempty"`
}
