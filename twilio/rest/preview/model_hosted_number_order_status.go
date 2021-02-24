/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// HostedNumberOrderStatus the model 'HostedNumberOrderStatus'
type HostedNumberOrderStatus string

// List of hosted_number_order_status
const (
	HOSTEDNUMBERORDERSTATUS_RECEIVED             HostedNumberOrderStatus = "received"
	HOSTEDNUMBERORDERSTATUS_PENDING_VERIFICATION HostedNumberOrderStatus = "pending-verification"
	HOSTEDNUMBERORDERSTATUS_VERIFIED             HostedNumberOrderStatus = "verified"
	HOSTEDNUMBERORDERSTATUS_PENDING_LOA          HostedNumberOrderStatus = "pending-loa"
	HOSTEDNUMBERORDERSTATUS_CARRIER_PROCESSING   HostedNumberOrderStatus = "carrier-processing"
	HOSTEDNUMBERORDERSTATUS_TESTING              HostedNumberOrderStatus = "testing"
	HOSTEDNUMBERORDERSTATUS_COMPLETED            HostedNumberOrderStatus = "completed"
	HOSTEDNUMBERORDERSTATUS_FAILED               HostedNumberOrderStatus = "failed"
	HOSTEDNUMBERORDERSTATUS_ACTION_REQUIRED      HostedNumberOrderStatus = "action-required"
)
