/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ApiV2010Account struct for ApiV2010Account
type ApiV2010Account struct {
	AuthToken string `json:"auth_token,omitempty"`
	DateCreated string `json:"date_created,omitempty"`
	DateUpdated string `json:"date_updated,omitempty"`
	FriendlyName string `json:"friendly_name,omitempty"`
	OwnerAccountSid string `json:"owner_account_sid,omitempty"`
	Sid string `json:"sid,omitempty"`
	Status string `json:"status,omitempty"`
	SubresourceUris map[string]interface{} `json:"subresource_uris,omitempty"`
	Type string `json:"type,omitempty"`
	Uri string `json:"uri,omitempty"`
}
