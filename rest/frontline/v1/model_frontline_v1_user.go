/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Frontline
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// FrontlineV1User struct for FrontlineV1User
type FrontlineV1User struct {
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The string that identifies the resource's User
	Identity *string `json:"identity,omitempty"`
	// The string that you assigned to describe the User
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The avatar URL which will be shown in Frontline application
	Avatar *string `json:"avatar,omitempty"`
	State  *string `json:"state,omitempty"`
	// Whether the User is available for new conversations
	IsAvailable *bool `json:"is_available,omitempty"`
	// An absolute URL for this user.
	Url *string `json:"url,omitempty"`
}
