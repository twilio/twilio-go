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

// ApiV2010AccountSipSipCredentialList struct for ApiV2010AccountSipSipCredentialList
type ApiV2010AccountSipSipCredentialList struct {
	AccountSid      *string                 `json:"account_sid,omitempty"`
	DateCreated     *string                 `json:"date_created,omitempty"`
	DateUpdated     *string                 `json:"date_updated,omitempty"`
	FriendlyName    *string                 `json:"friendly_name,omitempty"`
	Sid             *string                 `json:"sid,omitempty"`
	SubresourceUris *map[string]interface{} `json:"subresource_uris,omitempty"`
	Uri             *string                 `json:"uri,omitempty"`
}
