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

// SummaryCallType the model 'SummaryCallType'
type SummaryCallType string

// List of summary_call_type
const (
	SUMMARYCALLTYPE_CARRIER  SummaryCallType = "carrier"
	SUMMARYCALLTYPE_SIP      SummaryCallType = "sip"
	SUMMARYCALLTYPE_TRUNKING SummaryCallType = "trunking"
	SUMMARYCALLTYPE_CLIENT   SummaryCallType = "client"
)
