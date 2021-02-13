/*
 * Twilio - Ip_messaging
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

// UpdateChannelRequest struct for UpdateChannelRequest
type UpdateChannelRequest struct {
	Attributes   string    `json:"Attributes,omitempty"`
	CreatedBy    string    `json:"CreatedBy,omitempty"`
	DateCreated  time.Time `json:"DateCreated,omitempty"`
	DateUpdated  time.Time `json:"DateUpdated,omitempty"`
	FriendlyName string    `json:"FriendlyName,omitempty"`
	UniqueName   string    `json:"UniqueName,omitempty"`
}
