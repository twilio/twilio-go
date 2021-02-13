/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListAssetResponse struct for ListAssetResponse
type ListAssetResponse struct {
	Assets []ServerlessV1ServiceAsset `json:"Assets,omitempty"`
	Meta   ListServiceResponseMeta    `json:"Meta,omitempty"`
}
