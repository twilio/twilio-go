/*
 * Twilio - Flex
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject struct for InlineObject
type InlineObject struct {
	// The chat channel's friendly name.
	ChatFriendlyName string `json:"ChatFriendlyName"`
	// The chat channel's unique name.
	ChatUniqueName string `json:"ChatUniqueName,omitempty"`
	// The chat participant's friendly name.
	ChatUserFriendlyName string `json:"ChatUserFriendlyName"`
	// The SID of the Flex Flow.
	FlexFlowSid string `json:"FlexFlowSid"`
	// The `identity` value that uniquely identifies the new resource's chat User.
	Identity string `json:"Identity"`
	// Whether to create the channel as long-lived.
	LongLived bool `json:"LongLived,omitempty"`
	// The pre-engagement data.
	PreEngagementData string `json:"PreEngagementData,omitempty"`
	// The Target Contact Identity, for example the phone number of an SMS.
	Target string `json:"Target,omitempty"`
	// The Task attributes to be added for the TaskRouter Task.
	TaskAttributes string `json:"TaskAttributes,omitempty"`
	// The SID of the TaskRouter Task. Only valid when integration type is `task`. `null` for integration types `studio` & `external`
	TaskSid string `json:"TaskSid,omitempty"`
}
