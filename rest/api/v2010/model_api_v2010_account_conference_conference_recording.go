/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.14.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AccountConferenceConferenceRecording struct for ApiV2010AccountConferenceConferenceRecording
type ApiV2010AccountConferenceConferenceRecording struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The API version used to create the recording
	ApiVersion *string `json:"api_version,omitempty"`
	// The SID of the Call the resource is associated with
	CallSid *string `json:"call_sid,omitempty"`
	// The number of channels in the final recording file as an integer
	Channels *int32 `json:"channels,omitempty"`
	// The Conference SID that identifies the conference associated with the recording
	ConferenceSid *string `json:"conference_sid,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// The length of the recording in seconds
	Duration *string `json:"duration,omitempty"`
	// How to decrypt the recording.
	EncryptionDetails *map[string]interface{} `json:"encryption_details,omitempty"`
	// More information about why the recording is missing, if status is `absent`.
	ErrorCode *int32 `json:"error_code,omitempty"`
	// The one-time cost of creating the recording.
	Price *float32 `json:"price,omitempty"`
	// The currency used in the price property.
	PriceUnit *string `json:"price_unit,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// How the recording was created
	Source *string `json:"source,omitempty"`
	// The start time of the recording, given in RFC 2822 format
	StartTime *string `json:"start_time,omitempty"`
	// The status of the recording
	Status *string `json:"status,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
}
