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

// VideoV1RoomRecording struct for VideoV1RoomRecording
type VideoV1RoomRecording struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the RoomRecording resource.
	AccountSid *string `json:"account_sid,omitempty"`
	Status     *string `json:"status,omitempty"`
	// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The unique string that we created to identify the RoomRecording resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the recording source. For a Room Recording, this value is a `track_sid`.
	SourceSid *string `json:"source_sid,omitempty"`
	// The size of the recorded track in bytes.
	Size *int64 `json:"size,omitempty"`
	// The absolute URL of the resource.
	Url  *string `json:"url,omitempty"`
	Type *string `json:"type,omitempty"`
	// The duration of the recording rounded to the nearest second. Sub-second duration tracks have a `duration` of 1 second
	Duration        *int    `json:"duration,omitempty"`
	ContainerFormat *string `json:"container_format,omitempty"`
	Codec           *string `json:"codec,omitempty"`
	// A list of SIDs related to the Recording. Includes the `room_sid` and `participant_sid`.
	GroupingSids *map[string]interface{} `json:"grouping_sids,omitempty"`
	// The name that was given to the source track of the recording. If no name is given, the `source_sid` is used.
	TrackName *string `json:"track_name,omitempty"`
	// The time in milliseconds elapsed between an arbitrary point in time, common to all group rooms, and the moment when the source room of this track started. This information provides a synchronization mechanism for recordings belonging to the same room.
	Offset *int64 `json:"offset,omitempty"`
	// The URL of the media file associated with the recording when stored externally. See [External S3 Recordings](/docs/video/api/external-s3-recordings) for more details.
	MediaExternalLocation *string `json:"media_external_location,omitempty"`
	// The SID of the Room resource the recording is associated with.
	RoomSid *string `json:"room_sid,omitempty"`
	// The URLs of related resources.
	Links *map[string]interface{} `json:"links,omitempty"`
}
