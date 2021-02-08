/*
 * Twilio-Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateMessageRequest struct for CreateMessageRequest
type CreateMessageRequest struct {
	Attributes string `json:"Attributes,omitempty"`
	Body string `json:"Body"`
	From string `json:"From,omitempty"`
}
