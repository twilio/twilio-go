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
// CreateDocumentRequest struct for CreateDocumentRequest
type CreateDocumentRequest struct {
	// A JSON string that represents an arbitrary, schema-less object that the Sync Document stores. Can be up to 16KB in length.
	Data map[string]interface{} `json:"Data,omitempty"`
	// How long, in seconds, before the Sync Document expires and is deleted (the Sync Document's time-to-live). Can be an integer from 0 to 31,536,000 (1 year). The default value is `0`, which means the Sync Document does not expire. The Sync Document will be deleted automatically after it expires, but there can be a delay between the expiration time and the resources's deletion.
	Ttl int32 `json:"Ttl,omitempty"`
	// An application-defined string that uniquely identifies the Sync Document
	UniqueName string `json:"UniqueName,omitempty"`
}
