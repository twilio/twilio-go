/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// NumbersV2RegulatoryComplianceBundle struct for NumbersV2RegulatoryComplianceBundle
type NumbersV2RegulatoryComplianceBundle struct {
	AccountSid string `json:"AccountSid,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	Email string `json:"Email,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	RegulationSid string `json:"RegulationSid,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Status Status `json:"Status,omitempty"`
	StatusCallback string `json:"StatusCallback,omitempty"`
	Url string `json:"Url,omitempty"`
	ValidUntil time.Time `json:"ValidUntil,omitempty"`
}
