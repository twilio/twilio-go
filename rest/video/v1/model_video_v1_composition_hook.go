/*
 * Twilio - Video
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// VideoV1CompositionHook struct for VideoV1CompositionHook
type VideoV1CompositionHook struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The array of track names to include in the compositions created by the composition hook
	AudioSources *[]string `json:"audio_sources,omitempty"`
	// The array of track names to exclude from the compositions created by the composition hook
	AudioSourcesExcluded *[]string `json:"audio_sources_excluded,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Whether the CompositionHook is active
	Enabled *bool `json:"enabled,omitempty"`
	// The container format of the media files used by the compositions created by the composition hook
	Format *string `json:"format,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The dimensions of the video image in pixels expressed as columns (width) and rows (height)
	Resolution *string `json:"resolution,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The URL to send status information to your application
	StatusCallback *string `json:"status_callback,omitempty"`
	// The HTTP method we should use to call status_callback
	StatusCallbackMethod *string `json:"status_callback_method,omitempty"`
	// Whether intervals with no media are clipped
	Trim *bool `json:"trim,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
	// A JSON object that describes the video layout of the Composition
	VideoLayout *map[string]interface{} `json:"video_layout,omitempty"`
}
