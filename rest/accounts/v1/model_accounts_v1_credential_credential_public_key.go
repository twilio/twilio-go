/*
 * Twilio - Accounts
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// AccountsV1CredentialCredentialPublicKey struct for AccountsV1CredentialCredentialPublicKey
type AccountsV1CredentialCredentialPublicKey struct {
	AccountSid string `json:"account_sid,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
	DateUpdated time.Time `json:"date_updated,omitempty"`
	FriendlyName string `json:"friendly_name,omitempty"`
	Sid string `json:"sid,omitempty"`
	Url string `json:"url,omitempty"`
}
