/*
 * Twilio - Trunking
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// AddressRequirement the model 'AddressRequirement'
type AddressRequirement string

// List of address_requirement
const (
	ADDRESSREQUIREMENT_NONE    AddressRequirement = "none"
	ADDRESSREQUIREMENT_ANY     AddressRequirement = "any"
	ADDRESSREQUIREMENT_LOCAL   AddressRequirement = "local"
	ADDRESSREQUIREMENT_FOREIGN AddressRequirement = "foreign"
)
