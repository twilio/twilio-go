/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AccountCallCallFeedback struct for ApiV2010AccountCallCallFeedback
type ApiV2010AccountCallCallFeedback struct {
	AccountSid   *string   `json:"account_sid,omitempty"`
	DateCreated  *string   `json:"date_created,omitempty"`
	DateUpdated  *string   `json:"date_updated,omitempty"`
	Issues       *[]string `json:"issues,omitempty"`
	QualityScore *int32    `json:"quality_score,omitempty"`
	Sid          *string   `json:"sid,omitempty"`
}
