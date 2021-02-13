/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateSipIpAccessControlListMappingRequest struct for CreateSipIpAccessControlListMappingRequest
type CreateSipIpAccessControlListMappingRequest struct {
	// The unique id of the IP access control list to map to the SIP domain.
	IpAccessControlListSid string `json:"IpAccessControlListSid"`
}
