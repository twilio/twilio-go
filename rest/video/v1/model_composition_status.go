/*
 * Twilio - Video
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// CompositionStatus the model 'CompositionStatus'
type CompositionStatus string

// List of composition_status
const (
	COMPOSITIONSTATUS_ENQUEUED   CompositionStatus = "enqueued"
	COMPOSITIONSTATUS_PROCESSING CompositionStatus = "processing"
	COMPOSITIONSTATUS_COMPLETED  CompositionStatus = "completed"
	COMPOSITIONSTATUS_DELETED    CompositionStatus = "deleted"
	COMPOSITIONSTATUS_FAILED     CompositionStatus = "failed"
)
