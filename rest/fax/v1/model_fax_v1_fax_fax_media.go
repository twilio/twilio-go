/*
 * Twilio - Fax
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
// FaxV1FaxFaxMedia struct for FaxV1FaxFaxMedia
type FaxV1FaxFaxMedia struct {
	AccountSid string `json:"AccountSid,omitempty"`
	ContentType string `json:"ContentType,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	FaxSid string `json:"FaxSid,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Url string `json:"Url,omitempty"`
}
