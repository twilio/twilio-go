/*
 * Twilio - Supersim
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

// SupersimV1SmsCommand struct for SupersimV1SmsCommand
type SupersimV1SmsCommand struct {
	AccountSid  *string              `json:"account_sid,omitempty"`
	DateCreated *time.Time           `json:"date_created,omitempty"`
	DateUpdated *time.Time           `json:"date_updated,omitempty"`
	Direction   *SmsCommandDirection `json:"direction,omitempty"`
	Payload     *string              `json:"payload,omitempty"`
	Sid         *string              `json:"sid,omitempty"`
	SimSid      *string              `json:"sim_sid,omitempty"`
	Status      *SmsCommandStatus    `json:"status,omitempty"`
	Url         *string              `json:"url,omitempty"`
}
