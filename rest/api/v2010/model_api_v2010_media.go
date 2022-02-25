/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010Media struct for ApiV2010Media
type ApiV2010Media struct {
	// The SID of the Account that created this resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The default mime-type of the media
	ContentType *string `json:"content_type,omitempty"`
	// The RFC 2822 date and time in GMT that this resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that this resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// The SID of the resource that created the media
	ParentSid *string `json:"parent_sid,omitempty"`
	// The unique string that identifies this resource
	Sid *string `json:"sid,omitempty"`
	// The URI of this resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
}
