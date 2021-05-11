/*
 * Twilio - Video
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// VideoV1CompositionSettings struct for VideoV1CompositionSettings
type VideoV1CompositionSettings struct {
	AccountSid        *string `json:"account_sid,omitempty"`
	AwsCredentialsSid *string `json:"aws_credentials_sid,omitempty"`
	AwsS3Url          *string `json:"aws_s3_url,omitempty"`
	AwsStorageEnabled *bool   `json:"aws_storage_enabled,omitempty"`
	EncryptionEnabled *bool   `json:"encryption_enabled,omitempty"`
	EncryptionKeySid  *string `json:"encryption_key_sid,omitempty"`
	FriendlyName      *string `json:"friendly_name,omitempty"`
	Url               *string `json:"url,omitempty"`
}
