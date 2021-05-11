/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AccountCallCallRecording struct for ApiV2010AccountCallCallRecording
type ApiV2010AccountCallCallRecording struct {
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
	Price             *float32                `json:"price,omitempty"`
	PriceUnit         *string                 `json:"price_unit,omitempty"`
	Sid               *string                 `json:"sid,omitempty"`
	Source            *CallRecordingSource    `json:"source,omitempty"`
	StartTime         *string                 `json:"start_time,omitempty"`
	Status            *CallRecordingStatus    `json:"status,omitempty"`
	Track             *string                 `json:"track,omitempty"`
	Uri               *string                 `json:"uri,omitempty"`
}
