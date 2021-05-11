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

// ApiV2010AccountSipSipDomainSipIpAccessControlListMapping struct for ApiV2010AccountSipSipDomainSipIpAccessControlListMapping
type ApiV2010AccountSipSipDomainSipIpAccessControlListMapping struct {
	// The unique id of the Account that is responsible for this resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The date that this resource was created, given as GMT in RFC 2822 format.
	DateCreated *string `json:"date_created,omitempty"`
	// The date that this resource was last updated, given as GMT in RFC 2822 format.
	DateUpdated *string `json:"date_updated,omitempty"`
	// A human readable descriptive text for this resource, up to 64 characters long.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// A 34 character string that uniquely identifies this resource.
	Sid *string `json:"sid,omitempty"`
	// The list of IP addresses associated with this domain.
	SubresourceUris *map[string]interface{} `json:"subresource_uris,omitempty"`
	// The URI for this resource, relative to https://api.twilio.com
	Uri *string `json:"uri,omitempty"`
}
