/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Intelligence
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// IntelligenceV2Media struct for IntelligenceV2Media
type IntelligenceV2Media struct {
	// The unique SID identifier of the Account.
	AccountSid *string `json:"account_sid,omitempty"`
	// Downloadable URL for media, if stored in Twilio AI.
	MediaUrl *string `json:"media_url,omitempty"`
	// The unique SID identifier of the Service.
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique SID identifier of the Transcript.
	Sid *string `json:"sid,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
}