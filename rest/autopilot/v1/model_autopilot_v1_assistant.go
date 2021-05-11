/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// AutopilotV1Assistant struct for AutopilotV1Assistant
type AutopilotV1Assistant struct {
	AccountSid          *string                 `json:"account_sid,omitempty"`
	CallbackEvents      *string                 `json:"callback_events,omitempty"`
	CallbackUrl         *string                 `json:"callback_url,omitempty"`
	DateCreated         *time.Time              `json:"date_created,omitempty"`
	DateUpdated         *time.Time              `json:"date_updated,omitempty"`
	DevelopmentStage    *string                 `json:"development_stage,omitempty"`
	FriendlyName        *string                 `json:"friendly_name,omitempty"`
	LatestModelBuildSid *string                 `json:"latest_model_build_sid,omitempty"`
	Links               *map[string]interface{} `json:"links,omitempty"`
	LogQueries          *bool                   `json:"log_queries,omitempty"`
	NeedsModelBuild     *bool                   `json:"needs_model_build,omitempty"`
	Sid                 *string                 `json:"sid,omitempty"`
	UniqueName          *string                 `json:"unique_name,omitempty"`
	Url                 *string                 `json:"url,omitempty"`
}
