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
// InlineObject5 struct for InlineObject5
type InlineObject5 struct {
	// How long, in seconds, before the List Item's parent Sync List expires (time-to-live) and is deleted.  Can be an integer from 0 to 31,536,000 (1 year). The default value is `0`, which means the parent Sync List does not expire. The Sync List will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources's deletion.
	CollectionTtl int32 `json:"CollectionTtl,omitempty"`
	// A JSON string that represents an arbitrary, schema-less object that the List Item stores. Can be up to 16KB in length.
	Data map[string]interface{} `json:"Data"`
	// How long, in seconds, before the List Item expires (time-to-live) and is deleted.  Can be an integer from 0 to 31,536,000 (1 year). The default value is `0`, which means the List Item does not expire. The List Item will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources's deletion.
	ItemTtl int32 `json:"ItemTtl,omitempty"`
	// An alias for `item_ttl`. If both parameters are provided, this value is ignored.
	Ttl int32 `json:"Ttl,omitempty"`
}
