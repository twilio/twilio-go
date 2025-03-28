/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Video
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// VideoV1CompositionHook struct for VideoV1CompositionHook
type VideoV1CompositionHook struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CompositionHook resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The string that you assigned to describe the resource. Can be up to 100 characters long and must be unique within the account.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// Whether the CompositionHook is active. When `true`, the CompositionHook is triggered for every completed Group Room on the account. When `false`, the CompositionHook is never triggered.
	Enabled *bool `json:"enabled,omitempty"`
	// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The unique string that we created to identify the CompositionHook resource.
	Sid *string `json:"sid,omitempty"`
	// The array of track names to include in the compositions created by the composition hook. A composition triggered by the composition hook includes all audio sources specified in `audio_sources` except those specified in `audio_sources_excluded`. The track names in this property can include an asterisk as a wild card character, which matches zero or more characters in a track name. For example, `student*` includes tracks named `student` as well as `studentTeam`. Please, be aware that either video_layout or audio_sources have to be provided to get a valid creation request
	AudioSources *[]string `json:"audio_sources,omitempty"`
	// The array of track names to exclude from the compositions created by the composition hook. A composition triggered by the composition hook includes all audio sources specified in `audio_sources` except for those specified in `audio_sources_excluded`. The track names in this property can include an asterisk as a wild card character, which matches zero or more characters in a track name. For example, `student*` excludes `student` as well as `studentTeam`. This parameter can also be empty.
	AudioSourcesExcluded *[]string `json:"audio_sources_excluded,omitempty"`
	// A JSON object that describes the video layout of the composition in terms of regions as specified in the HTTP POST request that created the CompositionHook resource. See [POST Parameters](https://www.twilio.com/docs/video/api/compositions-resource#http-post-parameters) for more information. Please, be aware that either video_layout or audio_sources have to be provided to get a valid creation request
	VideoLayout *map[string]interface{} `json:"video_layout,omitempty"`
	// The dimensions of the video image in pixels expressed as columns (width) and rows (height). The string's format is `{width}x{height}`, such as `640x480`.
	Resolution *string `json:"resolution,omitempty"`
	// Whether intervals with no media are clipped, as specified in the POST request that created the CompositionHook resource. Compositions with `trim` enabled are shorter when the Room is created and no Participant joins for a while as well as if all the Participants leave the room and join later, because those gaps will be removed. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info.
	Trim   *bool   `json:"trim,omitempty"`
	Format *string `json:"format,omitempty"`
	// The URL we call using the `status_callback_method` to send status information to your application.
	StatusCallback *string `json:"status_callback,omitempty"`
	// The HTTP method we should use to call `status_callback`. Can be `POST` or `GET` and defaults to `POST`.
	StatusCallbackMethod *string `json:"status_callback_method,omitempty"`
	// The absolute URL of the resource.
	Url *string `json:"url,omitempty"`
}
