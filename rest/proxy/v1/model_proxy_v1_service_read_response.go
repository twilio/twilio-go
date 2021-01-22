/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ProxyV1ServiceReadResponse struct for ProxyV1ServiceReadResponse
type ProxyV1ServiceReadResponse struct {
	Meta ProxyV1ServiceReadResponseMeta `json:"Meta,omitempty"`
	Services []ProxyV1Service `json:"Services,omitempty"`
}
