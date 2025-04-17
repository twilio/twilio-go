/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Organization Public API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// ScimUser struct for ScimUser
type ScimUser struct {
	// Unique Twilio user sid
	Id string `json:"id,omitempty"`
	// External unique resource id defined by provisioning client
	ExternalId string `json:"externalId,omitempty"`
	// Unique username, MUST be same as primary email address
	UserName string `json:"userName"`
	// User friendly display name
	DisplayName string   `json:"displayName,omitempty"`
	Name        ScimName `json:"name,omitempty"`
	// Email address list of the user. Primary email must be defined if there are more than 1 email. Primary email must match the username.
	Emails []ScimEmailAddress `json:"emails,omitempty"`
	// Indicates whether the user is active
	Active bool `json:"active,omitempty"`
	// User's locale
	Locale string `json:"locale,omitempty"`
	// User's time zone
	Timezone string `json:"timezone,omitempty"`
	// An array of URIs that indicate the schemas supported for this user resource
	Schemas []string `json:"schemas,omitempty"`
	Meta    ScimMeta `json:"meta,omitempty"`
}
