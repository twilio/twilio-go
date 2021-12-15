# VideoV1Room

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**AudioOnly** | Pointer to **bool** | Indicates whether the room will only contain audio track participants for group rooms. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Duration** | Pointer to **int** | The duration of the room in seconds |
**EmptyRoomTimeout** | Pointer to **int** | The time a room will remain active after last participant leaves. |
**EnableTurn** | Pointer to **bool** | Enable Twilio's Network Traversal TURN service |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The UTC end time of the room in UTC ISO 8601 format |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources |
**MaxConcurrentPublishedTracks** | Pointer to **int** | The maximum number of published tracks allowed in the room at the same time |
**MaxParticipantDuration** | Pointer to **int** | The maximum number of seconds a Participant can be connected to the room |
**MaxParticipants** | Pointer to **int** | The maximum number of concurrent Participants allowed in the room |
**MediaRegion** | Pointer to **string** | The region for the media server in Group Rooms |
**RecordParticipantsOnConnect** | Pointer to **bool** | Whether to start recording when Participants connect |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Status** | Pointer to **string** | The status of the room |
**StatusCallback** | Pointer to **string** | The URL to send status information to your application |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method we use to call status_callback |
**Type** | Pointer to **string** | The type of room |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource |
**UnusedRoomTimeout** | Pointer to **int** | The time a room will remain active when no one joins. |
**Url** | Pointer to **string** | The absolute URL of the resource |
**VideoCodecs** | Pointer to **[]string** | An array of the video codecs that are supported when publishing a track in the room |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


