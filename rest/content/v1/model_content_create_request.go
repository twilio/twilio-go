/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Content
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// ContentCreateRequest Content creation request body
type ContentCreateRequest struct {
	// User defined name of the content
	FriendlyName string `json:"friendly_name,omitempty"`
	// Key value pairs of variable name to value
	Variables map[string]string `json:"variables,omitempty"`
	// Language code for the content
	Language string `json:"language"`
	Types    Types  `json:"types"`
}
