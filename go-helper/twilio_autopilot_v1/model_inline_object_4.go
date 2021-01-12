/*
 * Twilio - Autopilot
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
	// The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) tag that specifies the language of the value. Currently supported tags: `en-US`
	Language string `json:"Language"`
	// The string value that indicates which word the field value is a synonym of.
	SynonymOf string `json:"SynonymOf,omitempty"`
	// The Field Value data.
	Value string `json:"Value"`
}
