/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Iam
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// IamV1IndependentUser struct for IamV1IndependentUser
type IamV1IndependentUser struct {
	// Unique identifier of the user-organization mapping
	Id string `json:"id,omitempty"`
	// Unique Twilio organization sid
	OrganizationSid string `json:"organizationSid,omitempty"`
	// Unique Twilio user sid
	UserSid string `json:"userSid,omitempty"`
	// User first name
	FirstName string `json:"firstName,omitempty"`
	// User last name
	LastName string `json:"lastName,omitempty"`
	// User email address
	Email  string `json:"email,omitempty"`
	Status string `json:"status,omitempty"`
}
