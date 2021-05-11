/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// TrusthubV1CustomerProfileCustomerProfileChannelEndpointAssignment struct for TrusthubV1CustomerProfileCustomerProfileChannelEndpointAssignment
type TrusthubV1CustomerProfileCustomerProfileChannelEndpointAssignment struct {
	AccountSid          *string    `json:"account_sid,omitempty"`
	ChannelEndpointSid  *string    `json:"channel_endpoint_sid,omitempty"`
	ChannelEndpointType *string    `json:"channel_endpoint_type,omitempty"`
	CustomerProfileSid  *string    `json:"customer_profile_sid,omitempty"`
	DateCreated         *time.Time `json:"date_created,omitempty"`
	Sid                 *string    `json:"sid,omitempty"`
	Url                 *string    `json:"url,omitempty"`
}
