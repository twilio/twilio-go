/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ApiV2010AccountRecording struct for ApiV2010AccountRecording
type ApiV2010AccountRecording struct {
	AccountSid string `json:"AccountSid,omitempty"`
	ApiVersion string `json:"ApiVersion,omitempty"`
	CallSid string `json:"CallSid,omitempty"`
	Channels int32 `json:"Channels,omitempty"`
	ConferenceSid string `json:"ConferenceSid,omitempty"`
	DateCreated string `json:"DateCreated,omitempty"`
	DateUpdated string `json:"DateUpdated,omitempty"`
	Duration string `json:"Duration,omitempty"`
	EncryptionDetails map[string]interface{} `json:"EncryptionDetails,omitempty"`
	ErrorCode int32 `json:"ErrorCode,omitempty"`
	Price string `json:"Price,omitempty"`
	PriceUnit string `json:"PriceUnit,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Source Source `json:"Source,omitempty"`
	StartTime string `json:"StartTime,omitempty"`
	Status Status `json:"Status,omitempty"`
	SubresourceUris map[string]interface{} `json:"SubresourceUris,omitempty"`
	Uri string `json:"Uri,omitempty"`
}
