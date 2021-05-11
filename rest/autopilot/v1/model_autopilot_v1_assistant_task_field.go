/*
 * Twilio - Autopilot
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

// AutopilotV1AssistantTaskField struct for AutopilotV1AssistantTaskField
type AutopilotV1AssistantTaskField struct {
	AccountSid   *string    `json:"account_sid,omitempty"`
	AssistantSid *string    `json:"assistant_sid,omitempty"`
	DateCreated  *time.Time `json:"date_created,omitempty"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FieldType    *string    `json:"field_type,omitempty"`
	Sid          *string    `json:"sid,omitempty"`
	TaskSid      *string    `json:"task_sid,omitempty"`
	UniqueName   *string    `json:"unique_name,omitempty"`
	Url          *string    `json:"url,omitempty"`
}
