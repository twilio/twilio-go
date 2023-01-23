# VideoV1RoomRecording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the RoomRecording resource. |
**Status** | Pointer to [**string**](RoomRecordingEnumStatus.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Sid** | Pointer to **string** | The unique string that we created to identify the RoomRecording resource. |
**SourceSid** | Pointer to **string** | The SID of the recording source. For a Room Recording, this value is a `track_sid`. |
**Size** | Pointer to **int64** | The size of the recorded track in bytes. |
**Url** | Pointer to **string** | The absolute URL of the resource. |
**Type** | Pointer to [**string**](RoomRecordingEnumType.md) |  |
**Duration** | Pointer to **int** | The duration of the recording rounded to the nearest second. Sub-second duration tracks have a `duration` of 1 second |
**ContainerFormat** | Pointer to [**string**](RoomRecordingEnumFormat.md) |  |
**Codec** | Pointer to [**string**](RoomRecordingEnumCodec.md) |  |
**GroupingSids** | Pointer to **interface{}** | A list of SIDs related to the Recording. Includes the `room_sid` and `participant_sid`. |
**TrackName** | Pointer to **string** | The name that was given to the source track of the recording. If no name is given, the `source_sid` is used. |
**Offset** | Pointer to **int64** | The time in milliseconds elapsed between an arbitrary point in time, common to all group rooms, and the moment when the source room of this track started. This information provides a synchronization mechanism for recordings belonging to the same room. |
**MediaExternalLocation** | Pointer to **string** | The URL of the media file associated with the recording when stored externally. See [External S3 Recordings](/docs/video/api/external-s3-recordings) for more details. |
**RoomSid** | Pointer to **string** | The SID of the Room resource the recording is associated with. |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


