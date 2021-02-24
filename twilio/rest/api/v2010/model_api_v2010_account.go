/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ApiV2010Account struct for ApiV2010Account
type ApiV2010Account struct {
	AuthToken       *string                 `json:"AuthToken,omitempty"`
	DateCreated     *string                 `json:"DateCreated,omitempty"`
	DateUpdated     *string                 `json:"DateUpdated,omitempty"`
	FriendlyName    *string                 `json:"FriendlyName,omitempty"`
	OwnerAccountSid *string                 `json:"OwnerAccountSid,omitempty"`
	Sid             *string                 `json:"Sid,omitempty"`
	Status          *AccountStatus          `json:"Status,omitempty"`
	SubresourceUris *map[string]interface{} `json:"SubresourceUris,omitempty"`
	Type            *AccountType            `json:"Type,omitempty"`
	Uri             *string                 `json:"Uri,omitempty"`
}
