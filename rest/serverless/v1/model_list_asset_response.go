/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListAssetResponse struct for ListAssetResponse
type ListAssetResponse struct {
	Assets []ServerlessV1ServiceAsset `json:"assets,omitempty"`
	Meta   ListServiceResponseMeta    `json:"meta,omitempty"`
}
