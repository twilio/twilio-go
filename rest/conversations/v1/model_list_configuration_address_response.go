/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListConfigurationAddressResponse struct for ListConfigurationAddressResponse
type ListConfigurationAddressResponse struct {
	AddressConfigurations []ConversationsV1ConfigurationAddress `json:"address_configurations,omitempty"`
	Meta                  ListConfigurationAddressResponseMeta  `json:"meta,omitempty"`
}
