/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ParticipantStatus the model 'ParticipantStatus'
type ParticipantStatus string

// List of participant_status
const (
	PARTICIPANTSTATUS_QUEUED     ParticipantStatus = "queued"
	PARTICIPANTSTATUS_CONNECTING ParticipantStatus = "connecting"
	PARTICIPANTSTATUS_RINGING    ParticipantStatus = "ringing"
	PARTICIPANTSTATUS_CONNECTED  ParticipantStatus = "connected"
	PARTICIPANTSTATUS_COMPLETE   ParticipantStatus = "complete"
	PARTICIPANTSTATUS_FAILED     ParticipantStatus = "failed"
)
