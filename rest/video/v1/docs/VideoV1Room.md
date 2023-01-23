# VideoV1Room

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**Status** | Pointer to [**string**](RoomEnumRoomStatus.md) |  |
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**EnableTurn** | **bool** | Enable Twilio's Network Traversal TURN service |[optional] 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource |[optional] 
**StatusCallback** | **string** | The URL to send status information to your application |[optional] 
**StatusCallbackMethod** | **string** | The HTTP method we use to call status_callback |[optional] 
**EndTime** | [**time.Time**](time.Time.md) | The UTC end time of the room in UTC ISO 8601 format |[optional] 
**Duration** | Pointer to **int** | The duration of the room in seconds |
**Type** | Pointer to [**string**](RoomEnumRoomType.md) |  |
**MaxParticipants** | **int** | The maximum number of concurrent Participants allowed in the room |[optional] 
**MaxParticipantDuration** | **int** | The maximum number of seconds a Participant can be connected to the room |[optional] 
**MaxConcurrentPublishedTracks** | Pointer to **int** | The maximum number of published tracks allowed in the room at the same time |
**RecordParticipantsOnConnect** | **bool** | Whether to start recording when Participants connect |[optional] 
**VideoCodecs** | Pointer to [**[]string**](RoomEnumVideoCodec.md) | An array of the video codecs that are supported when publishing a track in the room |
**MediaRegion** | **string** | The region for the media server in Group Rooms |[optional] 
**AudioOnly** | **bool** | Indicates whether the room will only contain audio track participants for group rooms. |[optional] 
**EmptyRoomTimeout** | **int** | The time a room will remain active after last participant leaves. |[optional] 
**UnusedRoomTimeout** | **int** | The time a room will remain active when no one joins. |[optional] 
**LargeRoom** | **bool** | Indicates if this is a large room. |[optional] 
**Url** | **string** | The absolute URL of the resource |[optional] 
**Links** | **map[string]interface{}** | The URLs of related resources |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


