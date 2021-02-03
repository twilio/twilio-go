/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ApiV2010AccountQueueMember struct for ApiV2010AccountQueueMember
type ApiV2010AccountQueueMember struct {
	CallSid string `json:"CallSid,omitempty"`
	DateEnqueued string `json:"DateEnqueued,omitempty"`
	Position int32 `json:"Position,omitempty"`
	QueueSid string `json:"QueueSid,omitempty"`
	Uri string `json:"Uri,omitempty"`
	WaitTime int32 `json:"WaitTime,omitempty"`
}
