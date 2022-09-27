/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Trusthub
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// TrusthubV1Policies struct for TrusthubV1Policies
type TrusthubV1Policies struct {
	// The unique string that identifies the Policy resource
	Sid *string `json:"sid,omitempty"`
	// A human-readable description of the Policy resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The sid of a Policy object that dictates requirements
	Requirements *interface{} `json:"requirements,omitempty"`
	// The absolute URL of the Policy resource
	Url *string `json:"url,omitempty"`
}
