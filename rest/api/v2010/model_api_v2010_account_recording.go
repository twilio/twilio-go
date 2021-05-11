/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AccountRecording struct for ApiV2010AccountRecording
type ApiV2010AccountRecording struct {
	AccountSid        *string                 `json:"account_sid,omitempty"`
	ApiVersion        *string                 `json:"api_version,omitempty"`
	CallSid           *string                 `json:"call_sid,omitempty"`
	Channels          *int32                  `json:"channels,omitempty"`
	ConferenceSid     *string                 `json:"conference_sid,omitempty"`
	DateCreated       *string                 `json:"date_created,omitempty"`
	DateUpdated       *string                 `json:"date_updated,omitempty"`
	Duration          *string                 `json:"duration,omitempty"`
	EncryptionDetails *map[string]interface{} `json:"encryption_details,omitempty"`
	ErrorCode         *int32                  `json:"error_code,omitempty"`
	Price             *string                 `json:"price,omitempty"`
	PriceUnit         *string                 `json:"price_unit,omitempty"`
	Sid               *string                 `json:"sid,omitempty"`
	Source            *RecordingSource        `json:"source,omitempty"`
	StartTime         *string                 `json:"start_time,omitempty"`
	Status            *RecordingStatus        `json:"status,omitempty"`
	SubresourceUris   *map[string]interface{} `json:"subresource_uris,omitempty"`
	Uri               *string                 `json:"uri,omitempty"`
}
