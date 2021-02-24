/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// SinkStatus the model 'SinkStatus'
type SinkStatus string

// List of sink_status
const (
	SINKSTATUS_INITIALIZED SinkStatus = "initialized"
	SINKSTATUS_VALIDATING  SinkStatus = "validating"
	SINKSTATUS_ACTIVE      SinkStatus = "active"
	SINKSTATUS_FAILED      SinkStatus = "failed"
)
