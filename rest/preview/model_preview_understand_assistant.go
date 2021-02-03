/*
 * Twilio - Preview
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
// PreviewUnderstandAssistant struct for PreviewUnderstandAssistant
type PreviewUnderstandAssistant struct {
	AccountSid string `json:"AccountSid,omitempty"`
	CallbackEvents string `json:"CallbackEvents,omitempty"`
	CallbackUrl string `json:"CallbackUrl,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	LatestModelBuildSid string `json:"LatestModelBuildSid,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	LogQueries bool `json:"LogQueries,omitempty"`
	Sid string `json:"Sid,omitempty"`
	UniqueName string `json:"UniqueName,omitempty"`
	Url string `json:"Url,omitempty"`
}
