/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AccountSipSipIpAccessControlList struct for ApiV2010AccountSipSipIpAccessControlList
type ApiV2010AccountSipSipIpAccessControlList struct {
	AccountSid      *string                 `json:"AccountSid,omitempty"`
	DateCreated     *string                 `json:"DateCreated,omitempty"`
	DateUpdated     *string                 `json:"DateUpdated,omitempty"`
	FriendlyName    *string                 `json:"FriendlyName,omitempty"`
	Sid             *string                 `json:"Sid,omitempty"`
	SubresourceUris *map[string]interface{} `json:"SubresourceUris,omitempty"`
	Uri             *string                 `json:"Uri,omitempty"`
}
