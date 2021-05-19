/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListNetworkAccessProfileNetworkResponse struct for ListNetworkAccessProfileNetworkResponse
type ListNetworkAccessProfileNetworkResponse struct {
	Meta     ListCommandResponseMeta                                     `json:"meta,omitempty"`
	Networks []SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork `json:"networks,omitempty"`
}
