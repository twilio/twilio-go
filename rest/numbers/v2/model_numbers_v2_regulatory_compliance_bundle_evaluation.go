/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// NumbersV2RegulatoryComplianceBundleEvaluation struct for NumbersV2RegulatoryComplianceBundleEvaluation
type NumbersV2RegulatoryComplianceBundleEvaluation struct {
	AccountSid string `json:"AccountSid,omitempty"`
	BundleSid string `json:"BundleSid,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	RegulationSid string `json:"RegulationSid,omitempty"`
	Results []map[string]interface{} `json:"Results,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Status string `json:"Status,omitempty"`
	Url string `json:"Url,omitempty"`
}
