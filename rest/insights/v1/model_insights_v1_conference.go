/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// InsightsV1Conference struct for InsightsV1Conference
type InsightsV1Conference struct {
	AccountSid                *string                 `json:"account_sid,omitempty"`
	ConferenceSid             *string                 `json:"conference_sid,omitempty"`
	ConnectDurationSeconds    *int                    `json:"connect_duration_seconds,omitempty"`
	CreateTime                *time.Time              `json:"create_time,omitempty"`
	DetectedIssues            *map[string]interface{} `json:"detected_issues,omitempty"`
	DurationSeconds           *int                    `json:"duration_seconds,omitempty"`
	EndReason                 *string                 `json:"end_reason,omitempty"`
	EndTime                   *time.Time              `json:"end_time,omitempty"`
	EndedBy                   *string                 `json:"ended_by,omitempty"`
	FriendlyName              *string                 `json:"friendly_name,omitempty"`
	Links                     *map[string]interface{} `json:"links,omitempty"`
	MaxConcurrentParticipants *int                    `json:"max_concurrent_participants,omitempty"`
	MaxParticipants           *int                    `json:"max_participants,omitempty"`
	MixerRegion               *string                 `json:"mixer_region,omitempty"`
	MixerRegionRequested      *string                 `json:"mixer_region_requested,omitempty"`
	RecordingEnabled          *bool                   `json:"recording_enabled,omitempty"`
	StartTime                 *time.Time              `json:"start_time,omitempty"`
	Status                    *string                 `json:"status,omitempty"`
	TagInfo                   *map[string]interface{} `json:"tag_info,omitempty"`
	Tags                      *[]string               `json:"tags,omitempty"`
	UniqueParticipants        *int                    `json:"unique_participants,omitempty"`
	Url                       *string                 `json:"url,omitempty"`
}