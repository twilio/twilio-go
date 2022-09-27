# VideoV1Room

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Status** | Pointer to [**string**](RoomEnumRoomStatus.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**EnableTurn** | Pointer to **bool** | Enable Twilio's Network Traversal TURN service |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource |
**StatusCallback** | Pointer to **string** | The URL to send status information to your application |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method we use to call status_callback |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The UTC end time of the room in UTC ISO 8601 format |
**Duration** | Pointer to **int** | The duration of the room in seconds |
**Type** | Pointer to [**string**](RoomEnumRoomType.md) |  |
**MaxParticipants** | Pointer to **int** | The maximum number of concurrent Participants allowed in the room |
**MaxParticipantDuration** | Pointer to **int** | The maximum number of seconds a Participant can be connected to the room |
**MaxConcurrentPublishedTracks** | Pointer to **int** | The maximum number of published tracks allowed in the room at the same time |
**RecordParticipantsOnConnect** | Pointer to **bool** | Whether to start recording when Participants connect |
**VideoCodecs** | Pointer to [**[]string**](RoomEnumVideoCodec.md) | An array of the video codecs that are supported when publishing a track in the room |
**MediaRegion** | Pointer to **string** | The region for the media server in Group Rooms |
**AudioOnly** | Pointer to **bool** | Indicates whether the room will only contain audio track participants for group rooms. |
**EmptyRoomTimeout** | Pointer to **int** | The time a room will remain active after last participant leaves. |
**UnusedRoomTimeout** | Pointer to **int** | The time a room will remain active when no one joins. |
**LargeRoom** | Pointer to **bool** | Indicates if this is a large room. |
**Url** | Pointer to **string** | The absolute URL of the resource |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


