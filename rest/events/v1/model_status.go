/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Status the model 'Status'
type Status string

// List of status
const (
	STATUS_INITIALIZED Status = "initialized"
	STATUS_VALIDATING Status = "validating"
	STATUS_ACTIVE Status = "active"
	STATUS_FAILED Status = "failed"
)
