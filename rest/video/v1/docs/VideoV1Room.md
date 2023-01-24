# VideoV1Room

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Room resource. |
**Status** | Pointer to [**string**](RoomEnumRoomStatus.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Room resource. |
**EnableTurn** | Pointer to **bool** | Deprecated, now always considered to be true. |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource. It can be used as a `room_sid` in place of the resource's `sid` in the URL to address the resource, assuming it does not contain any [reserved characters](https://tools.ietf.org/html/rfc3986#section-2.2) that would need to be URL encoded. This value is unique for `in-progress` rooms. SDK clients can use this name to connect to the room. REST API clients can use this name in place of the Room SID to interact with the room as long as the room is `in-progress`. |
**StatusCallback** | Pointer to **string** | The URL we call using the `status_callback_method` to send status information to your application on every room event. See [Status Callbacks](https://www.twilio.com/docs/video/api/status-callbacks) for more info. |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method we use to call `status_callback`. Can be `POST` or `GET` and defaults to `POST`. |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The UTC end time of the room in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format. |
**Duration** | Pointer to **int** | The duration of the room in seconds. |
**Type** | Pointer to [**string**](RoomEnumRoomType.md) |  |
**MaxParticipants** | Pointer to **int** | The maximum number of concurrent Participants allowed in the room.  |
**MaxParticipantDuration** | Pointer to **int** | The maximum number of seconds a Participant can be connected to the room. The maximum possible value is 86400 seconds (24 hours). The default is 14400 seconds (4 hours). |
**MaxConcurrentPublishedTracks** | Pointer to **int** | The maximum number of published audio, video, and data tracks all participants combined are allowed to publish in the room at the same time. Check [Programmable Video Limits](https://www.twilio.com/docs/video/programmable-video-limits) for more details. If it is set to 0 it means unconstrained. |
**RecordParticipantsOnConnect** | Pointer to **bool** | Whether to start recording when Participants connect. ***This feature is not available in `peer-to-peer` rooms.*** |
**VideoCodecs** | Pointer to [**[]string**](RoomEnumVideoCodec.md) | An array of the video codecs that are supported when publishing a track in the room.  Can be: `VP8` and `H264`.  ***This feature is not available in `peer-to-peer` rooms*** |
**MediaRegion** | Pointer to **string** | The region for the media server in Group Rooms.  Can be: one of the [available Media Regions](https://www.twilio.com/docs/video/ip-address-whitelisting#media-servers). ***This feature is not available in `peer-to-peer` rooms.*** |
**AudioOnly** | Pointer to **bool** | When set to true, indicates that the participants in the room will only publish audio. No video tracks will be allowed. Group rooms only. |
**EmptyRoomTimeout** | Pointer to **int** | Specifies how long (in minutes) a room will remain active after last participant leaves. Can be configured when creating a room via REST API. For Ad-Hoc rooms this value cannot be changed. |
**UnusedRoomTimeout** | Pointer to **int** | Specifies how long (in minutes) a room will remain active if no one joins. Can be configured when creating a room via REST API. For Ad-Hoc rooms this value cannot be changed. |
**LargeRoom** | Pointer to **bool** | Indicates if this is a large room. |
**Url** | Pointer to **string** | The absolute URL of the resource. |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


