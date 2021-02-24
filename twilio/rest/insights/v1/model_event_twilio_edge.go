/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// EventTwilioEdge the model 'EventTwilioEdge'
type EventTwilioEdge string

// List of event_twilio_edge
const (
	EVENTTWILIOEDGE_UNKNOWN_EDGE EventTwilioEdge = "unknown_edge"
	EVENTTWILIOEDGE_CARRIER_EDGE EventTwilioEdge = "carrier_edge"
	EVENTTWILIOEDGE_SIP_EDGE     EventTwilioEdge = "sip_edge"
	EVENTTWILIOEDGE_SDK_EDGE     EventTwilioEdge = "sdk_edge"
	EVENTTWILIOEDGE_CLIENT_EDGE  EventTwilioEdge = "client_edge"
)
