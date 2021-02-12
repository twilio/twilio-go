/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Status the model 'Status'
type Status string

// List of status
const (
	STATUS_OPEN Status = "open"
	STATUS_IN_PROGRESS Status = "in-progress"
	STATUS_CLOSED Status = "closed"
	STATUS_FAILED Status = "failed"
	STATUS_UNKNOWN Status = "unknown"
)
