/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// ApiV2010Member struct for ApiV2010Member
type ApiV2010Member struct {
	// The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) the Member resource is associated with.
	CallSid *string `json:"call_sid,omitempty"`
	// The date that the member was enqueued, given in RFC 2822 format.
	DateEnqueued *string `json:"date_enqueued,omitempty"`
	// This member's current position in the queue.
	Position int `json:"position,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`.
	Uri *string `json:"uri,omitempty"`
	// The number of seconds the member has been in the queue.
	WaitTime int `json:"wait_time,omitempty"`
	// The SID of the Queue the member is in.
	QueueSid *string `json:"queue_sid,omitempty"`
}
