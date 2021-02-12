/*
 * Twilio - Wireless
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
	STATUS_NEW Status = "new"
	STATUS_READY Status = "ready"
	STATUS_ACTIVE Status = "active"
	STATUS_SUSPENDED Status = "suspended"
	STATUS_DEACTIVATED Status = "deactivated"
	STATUS_CANCELED Status = "canceled"
	STATUS_SCHEDULED Status = "scheduled"
	STATUS_UPDATING Status = "updating"
)
