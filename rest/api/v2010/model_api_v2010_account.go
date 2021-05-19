/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010Account struct for ApiV2010Account
type ApiV2010Account struct {
	// The authorization token for this account
	AuthToken *string `json:"auth_token,omitempty"`
	// The date this account was created
	DateCreated *string `json:"date_created,omitempty"`
	// The date this account was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// A human readable description of this account
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The unique 34 character id representing the parent of this account
	OwnerAccountSid *string `json:"owner_account_sid,omitempty"`
	// A 34 character string that uniquely identifies this resource.
	Sid *string `json:"sid,omitempty"`
	// The status of this account
	Status *string `json:"status,omitempty"`
	// Account Instance Subresources
	SubresourceUris *map[string]interface{} `json:"subresource_uris,omitempty"`
	// The type of this account
	Type *string `json:"type,omitempty"`
	// The URI for this resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
}
