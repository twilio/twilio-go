/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CallType the model 'CallType'
type CallType string

// List of call_type
const (
	CALLTYPE_CARRIER CallType = "carrier"
	CALLTYPE_SIP CallType = "sip"
	CALLTYPE_TRUNKING CallType = "trunking"
	CALLTYPE_CLIENT CallType = "client"
)
