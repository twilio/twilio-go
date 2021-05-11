/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// SummaryCallState the model 'SummaryCallState'
type SummaryCallState string

// List of summary_call_state
const (
	SUMMARYCALLSTATE_RINGING   SummaryCallState = "ringing"
	SUMMARYCALLSTATE_COMPLETED SummaryCallState = "completed"
	SUMMARYCALLSTATE_BUSY      SummaryCallState = "busy"
	SUMMARYCALLSTATE_FAIL      SummaryCallState = "fail"
	SUMMARYCALLSTATE_NOANSWER  SummaryCallState = "noanswer"
	SUMMARYCALLSTATE_CANCELED  SummaryCallState = "canceled"
	SUMMARYCALLSTATE_ANSWERED  SummaryCallState = "answered"
	SUMMARYCALLSTATE_UNDIALED  SummaryCallState = "undialed"
)
