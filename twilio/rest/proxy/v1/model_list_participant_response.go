/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListParticipantResponse struct for ListParticipantResponse
type ListParticipantResponse struct {
	Meta         ListServiceResponseMeta            `json:"Meta,omitempty"`
	Participants []ProxyV1ServiceSessionParticipant `json:"Participants,omitempty"`
}
