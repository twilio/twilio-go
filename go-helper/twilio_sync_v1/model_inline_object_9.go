/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject9 struct for InlineObject9
type InlineObject9 struct {
	// How long, in seconds, before the Sync Map expires (time-to-live) and is deleted. Can be an integer from 0 to 31,536,000 (1 year). The default value is `0`, which means the Sync Map does not expire. The Sync Map will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources's deletion.
	CollectionTtl int32 `json:"CollectionTtl,omitempty"`
	// An alias for `collection_ttl`. If both parameters are provided, this value is ignored.
	Ttl int32 `json:"Ttl,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used as an alternative to the `sid` in the URL path to address the resource.
	UniqueName string `json:"UniqueName,omitempty"`
}
