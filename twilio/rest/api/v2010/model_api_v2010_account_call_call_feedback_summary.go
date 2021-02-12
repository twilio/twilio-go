/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// ApiV2010AccountCallCallFeedbackSummary struct for ApiV2010AccountCallCallFeedbackSummary
type ApiV2010AccountCallCallFeedbackSummary struct {
	AccountSid string `json:"AccountSid,omitempty"`
	CallCount int32 `json:"CallCount,omitempty"`
	CallFeedbackCount int32 `json:"CallFeedbackCount,omitempty"`
	DateCreated string `json:"DateCreated,omitempty"`
	DateUpdated string `json:"DateUpdated,omitempty"`
	EndDate time.Time `json:"EndDate,omitempty"`
	IncludeSubaccounts bool `json:"IncludeSubaccounts,omitempty"`
	Issues []map[string]interface{} `json:"Issues,omitempty"`
	QualityScoreAverage float32 `json:"QualityScoreAverage,omitempty"`
	QualityScoreMedian float32 `json:"QualityScoreMedian,omitempty"`
	QualityScoreStandardDeviation float32 `json:"QualityScoreStandardDeviation,omitempty"`
	Sid string `json:"Sid,omitempty"`
	StartDate time.Time `json:"StartDate,omitempty"`
	Status Status `json:"Status,omitempty"`
}
