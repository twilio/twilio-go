/*
 * Twilio - Video
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
	STATUS_PROCESSING Status = "processing"
	STATUS_COMPLETED Status = "completed"
	STATUS_DELETED Status = "deleted"
	STATUS_FAILED Status = "failed"
)
