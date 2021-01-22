/*
 * Twilio - Preview
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
// PreviewMarketplaceInstalledAddOn struct for PreviewMarketplaceInstalledAddOn
type PreviewMarketplaceInstalledAddOn struct {
	AccountSid string `json:"AccountSid,omitempty"`
	Configuration map[string]interface{} `json:"Configuration,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	Description string `json:"Description,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	Sid string `json:"Sid,omitempty"`
	UniqueName string `json:"UniqueName,omitempty"`
	Url string `json:"Url,omitempty"`
}
