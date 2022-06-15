/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListParticipantResponse struct for ListParticipantResponse
type ListParticipantResponse struct {
	Meta         ListServiceResponseMeta `json:"meta,omitempty"`
	Participants []ProxyV1Participant    `json:"participants,omitempty"`
}
