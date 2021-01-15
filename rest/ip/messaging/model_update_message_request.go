/*
 * Twilio - Ip_messaging
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
// UpdateMessageRequest struct for UpdateMessageRequest
type UpdateMessageRequest struct {
	Attributes string `json:"Attributes,omitempty"`
	Body string `json:"Body,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	From string `json:"From,omitempty"`
	LastUpdatedBy string `json:"LastUpdatedBy,omitempty"`
}
