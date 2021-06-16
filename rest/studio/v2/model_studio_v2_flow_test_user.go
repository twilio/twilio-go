/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// StudioV2FlowTestUser struct for StudioV2FlowTestUser
type StudioV2FlowTestUser struct {
	// Unique identifier of the flow.
	Sid *string `json:"sid,omitempty"`
	// List of test user identities that can test draft versions of the flow.
	TestUsers *[]string `json:"test_users,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
}
