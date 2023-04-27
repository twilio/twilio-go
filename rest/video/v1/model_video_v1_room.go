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
	"encoding/json"
	"github.com/twilio/twilio-go/client"
	"time"
)
// VideoV1Room struct for VideoV1Room
type VideoV1Room struct {
		// The unique string that we created to identify the Room resource.
	Sid *string `json:"sid,omitempty"`
	Status *string `json:"status,omitempty"`
		// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
		// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
		// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Room resource.
	AccountSid *string `json:"account_sid,omitempty"`
		// Deprecated, now always considered to be true.
	EnableTurn *bool `json:"enable_turn,omitempty"`
		// An application-defined string that uniquely identifies the resource. It can be used as a `room_sid` in place of the resource's `sid` in the URL to address the resource, assuming it does not contain any [reserved characters](https://tools.ietf.org/html/rfc3986#section-2.2) that would need to be URL encoded. This value is unique for `in-progress` rooms. SDK clients can use this name to connect to the room. REST API clients can use this name in place of the Room SID to interact with the room as long as the room is `in-progress`.
	UniqueName *string `json:"unique_name,omitempty"`
		// The URL we call using the `status_callback_method` to send status information to your application on every room event. See [Status Callbacks](https://www.twilio.com/docs/video/api/status-callbacks) for more info.
	StatusCallback *string `json:"status_callback,omitempty"`
		// The HTTP method we use to call `status_callback`. Can be `POST` or `GET` and defaults to `POST`.
	StatusCallbackMethod *string `json:"status_callback_method,omitempty"`
		// The UTC end time of the room in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format.
	EndTime *time.Time `json:"end_time,omitempty"`
		// The duration of the room in seconds.
	Duration *int `json:"duration,omitempty"`
	Type *string `json:"type,omitempty"`
		// The maximum number of concurrent Participants allowed in the room. 
	MaxParticipants *int `json:"max_participants,omitempty"`
		// The maximum number of seconds a Participant can be connected to the room. The maximum possible value is 86400 seconds (24 hours). The default is 14400 seconds (4 hours).
	MaxParticipantDuration *int `json:"max_participant_duration,omitempty"`
		// The maximum number of published audio, video, and data tracks all participants combined are allowed to publish in the room at the same time. Check [Programmable Video Limits](https://www.twilio.com/docs/video/programmable-video-limits) for more details. If it is set to 0 it means unconstrained.
	MaxConcurrentPublishedTracks *int `json:"max_concurrent_published_tracks,omitempty"`
		// Whether to start recording when Participants connect. ***This feature is not available in `peer-to-peer` rooms.***
	RecordParticipantsOnConnect *bool `json:"record_participants_on_connect,omitempty"`
		// An array of the video codecs that are supported when publishing a track in the room.  Can be: `VP8` and `H264`.  ***This feature is not available in `peer-to-peer` rooms***
	VideoCodecs *[]string `json:"video_codecs,omitempty"`
		// The region for the media server in Group Rooms.  Can be: one of the [available Media Regions](https://www.twilio.com/docs/video/ip-address-whitelisting#media-servers). ***This feature is not available in `peer-to-peer` rooms.***
	MediaRegion *string `json:"media_region,omitempty"`
		// When set to true, indicates that the participants in the room will only publish audio. No video tracks will be allowed. Group rooms only.
	AudioOnly *bool `json:"audio_only,omitempty"`
		// Specifies how long (in minutes) a room will remain active after last participant leaves. Can be configured when creating a room via REST API. For Ad-Hoc rooms this value cannot be changed.
	EmptyRoomTimeout *int `json:"empty_room_timeout,omitempty"`
		// Specifies how long (in minutes) a room will remain active if no one joins. Can be configured when creating a room via REST API. For Ad-Hoc rooms this value cannot be changed.
	UnusedRoomTimeout *int `json:"unused_room_timeout,omitempty"`
		// Indicates if this is a large room.
	LargeRoom *bool `json:"large_room,omitempty"`
		// The absolute URL of the resource.
	Url *string `json:"url,omitempty"`
		// The URLs of related resources.
	Links *map[string]interface{} `json:"links,omitempty"`
}


