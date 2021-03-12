/*
    * Twilio - Accounts
    *
    * This is the public Twilio REST API.
    *
    * API version: 1.10.0
    * Contact: support@twilio.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi
import (
	"time"
)
// AccountsV1SecondaryAuthToken struct for AccountsV1SecondaryAuthToken
type AccountsV1SecondaryAuthToken struct {
	// The SID of the Account that the secondary Auth Token was created for
	AccountSid *string `json:"AccountSid,omitempty"`
	// The ISO 8601 formatted date and time in UTC when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The ISO 8601 formatted date and time in UTC when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The generated secondary Auth Token
	SecondaryAuthToken *string `json:"SecondaryAuthToken,omitempty"`
	// The URI for this resource, relative to `https://accounts.twilio.com`
	Url *string `json:"Url,omitempty"`
}
