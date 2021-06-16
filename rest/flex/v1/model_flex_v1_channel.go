/*
 * Twilio - Flex
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// FlexV1Channel struct for FlexV1Channel
type FlexV1Channel struct {
	// The SID of the Account that created the resource and owns this Workflow
	AccountSid *string `json:"account_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the Flex chat channel was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the Flex chat channel was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The SID of the Flex Flow
	FlexFlowSid *string `json:"flex_flow_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The SID of the TaskRouter Task
	TaskSid *string `json:"task_sid,omitempty"`
	// The absolute URL of the Flex chat channel resource
	Url *string `json:"url,omitempty"`
	// The SID of the chat user
	UserSid *string `json:"user_sid,omitempty"`
}
