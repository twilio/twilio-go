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
// InlineObject4 struct for InlineObject4
type InlineObject4 struct {
	// How long, in seconds, before the Sync List expires (time-to-live) and is deleted.  Can be an integer from 0 to 31,536,000 (1 year). The default value is `0`, which means the Sync List does not expire. The Sync List will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources's deletion.
	CollectionTtl int32 `json:"CollectionTtl,omitempty"`
	// Alias for collection_ttl. If both are provided, this value is ignored.
	Ttl int32 `json:"Ttl,omitempty"`
	// An application-defined string that uniquely identifies the resource. This value must be unique within its Service and it can be up to 320 characters long. The `unique_name` value can be used as an alternative to the `sid` in the URL path to address the resource.
	UniqueName string `json:"UniqueName,omitempty"`
}
