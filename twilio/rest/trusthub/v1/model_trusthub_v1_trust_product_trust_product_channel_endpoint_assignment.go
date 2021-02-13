/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// TrusthubV1TrustProductTrustProductChannelEndpointAssignment struct for TrusthubV1TrustProductTrustProductChannelEndpointAssignment
type TrusthubV1TrustProductTrustProductChannelEndpointAssignment struct {
	AccountSid          string    `json:"AccountSid,omitempty"`
	ChannelEndpointSid  string    `json:"ChannelEndpointSid,omitempty"`
	ChannelEndpointType string    `json:"ChannelEndpointType,omitempty"`
	DateCreated         time.Time `json:"DateCreated,omitempty"`
	Sid                 string    `json:"Sid,omitempty"`
	TrustProductSid     string    `json:"TrustProductSid,omitempty"`
	Url                 string    `json:"Url,omitempty"`
}
