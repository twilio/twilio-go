/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// UpdateSipIpAddressRequest struct for UpdateSipIpAddressRequest
type UpdateSipIpAddressRequest struct {
	// An integer representing the length of the CIDR prefix to use with this IP address when accepting traffic. By default the entire IP address is used.
	CidrPrefixLength int32 `json:"CidrPrefixLength,omitempty"`
	// A human readable descriptive text for this resource, up to 64 characters long.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// An IP address in dotted decimal notation from which you want to accept traffic. Any SIP requests from this IP address will be allowed by Twilio. IPv4 only supported today.
	IpAddress string `json:"IpAddress,omitempty"`
}
