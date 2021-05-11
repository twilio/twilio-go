/*
 * Twilio - Flex
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// FlexV1Channel struct for FlexV1Channel
type FlexV1Channel struct {
	AccountSid  *string    `json:"account_sid,omitempty"`
	DateCreated *time.Time `json:"date_created,omitempty"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	FlexFlowSid *string    `json:"flex_flow_sid,omitempty"`
	Sid         *string    `json:"sid,omitempty"`
	TaskSid     *string    `json:"task_sid,omitempty"`
	Url         *string    `json:"url,omitempty"`
	UserSid     *string    `json:"user_sid,omitempty"`
}
