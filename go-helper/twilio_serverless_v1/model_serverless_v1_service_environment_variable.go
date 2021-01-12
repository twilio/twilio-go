/*
 * Twilio - Serverless
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
// ServerlessV1ServiceEnvironmentVariable struct for ServerlessV1ServiceEnvironmentVariable
type ServerlessV1ServiceEnvironmentVariable struct {
	AccountSid string `json:"account_sid,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
	DateUpdated time.Time `json:"date_updated,omitempty"`
	EnvironmentSid string `json:"environment_sid,omitempty"`
	Key string `json:"key,omitempty"`
	ServiceSid string `json:"service_sid,omitempty"`
	Sid string `json:"sid,omitempty"`
	Url string `json:"url,omitempty"`
	Value string `json:"value,omitempty"`
}
