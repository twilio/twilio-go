/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject4 struct for InlineObject4
type InlineObject4 struct {
	// Whether the Target is enabled.
	Enabled bool `json:"Enabled,omitempty"`
	// A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The relative importance of the target. Can be an integer from 0 to 65535, inclusive. The lowest number represents the most important target.
	Priority int32 `json:"Priority,omitempty"`
	// The SIP address you want Twilio to route your calls to. This must be a `sip:` schema. `sips` is NOT supported.
	Target string `json:"Target,omitempty"`
	// The value that determines the relative share of the load the Target should receive compared to other Targets with the same priority. Can be an integer from 1 to 65535, inclusive. Targets with higher values receive more load than those with lower ones with the same priority.
	Weight int32 `json:"Weight,omitempty"`
}
