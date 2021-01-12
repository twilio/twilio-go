/*
 * Twilio - Trunking
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
// TrunkingV1Trunk struct for TrunkingV1Trunk
type TrunkingV1Trunk struct {
	AccountSid string `json:"account_sid,omitempty"`
	AuthType string `json:"auth_type,omitempty"`
	AuthTypeSet []string `json:"auth_type_set,omitempty"`
	CnamLookupEnabled bool `json:"cnam_lookup_enabled,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
	DateUpdated time.Time `json:"date_updated,omitempty"`
	DisasterRecoveryMethod string `json:"disaster_recovery_method,omitempty"`
	DisasterRecoveryUrl string `json:"disaster_recovery_url,omitempty"`
	DomainName string `json:"domain_name,omitempty"`
	FriendlyName string `json:"friendly_name,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
	Recording map[string]interface{} `json:"recording,omitempty"`
	Secure bool `json:"secure,omitempty"`
	Sid string `json:"sid,omitempty"`
	TransferMode string `json:"transfer_mode,omitempty"`
	Url string `json:"url,omitempty"`
}
