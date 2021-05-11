/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// TrusthubV1TrustProductTrustProductEvaluation struct for TrusthubV1TrustProductTrustProductEvaluation
type TrusthubV1TrustProductTrustProductEvaluation struct {
	AccountSid      *string                       `json:"account_sid,omitempty"`
	DateCreated     *time.Time                    `json:"date_created,omitempty"`
	PolicySid       *string                       `json:"policy_sid,omitempty"`
	Results         *[]map[string]interface{}     `json:"results,omitempty"`
	Sid             *string                       `json:"sid,omitempty"`
	Status          *TrustProductEvaluationStatus `json:"status,omitempty"`
	TrustProductSid *string                       `json:"trust_product_sid,omitempty"`
	Url             *string                       `json:"url,omitempty"`
}
