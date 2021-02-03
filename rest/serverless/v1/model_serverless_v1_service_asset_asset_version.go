/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// ServerlessV1ServiceAssetAssetVersion struct for ServerlessV1ServiceAssetAssetVersion
type ServerlessV1ServiceAssetAssetVersion struct {
	AccountSid string `json:"AccountSid,omitempty"`
	AssetSid string `json:"AssetSid,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	Path string `json:"Path,omitempty"`
	ServiceSid string `json:"ServiceSid,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Url string `json:"Url,omitempty"`
	Visibility Visibility `json:"Visibility,omitempty"`
}
