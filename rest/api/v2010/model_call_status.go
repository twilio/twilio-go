/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// CallStatus the model 'CallStatus'
type CallStatus string

// List of call_status
const (
	CALLSTATUS_QUEUED      CallStatus = "queued"
	CALLSTATUS_RINGING     CallStatus = "ringing"
	CALLSTATUS_IN_PROGRESS CallStatus = "in-progress"
	CALLSTATUS_COMPLETED   CallStatus = "completed"
	CALLSTATUS_BUSY        CallStatus = "busy"
	CALLSTATUS_FAILED      CallStatus = "failed"
	CALLSTATUS_NO_ANSWER   CallStatus = "no-answer"
	CALLSTATUS_CANCELED    CallStatus = "canceled"
)
