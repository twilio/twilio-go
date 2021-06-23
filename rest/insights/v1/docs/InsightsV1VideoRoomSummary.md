# InsightsV1VideoRoomSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | Account SID associated with this room. |
**Codecs** | Pointer to **[]string** | Codecs used by participants in the room. |
**ConcurrentParticipants** | Pointer to **int** | Actual number of concurrent participants. |
**CreateTime** | Pointer to [**time.Time**](time.Time.md) | Creation time of the room. |
**CreatedMethod** | Pointer to **string** | How the room was created. |
**DurationSec** | Pointer to **int** | Total room duration from create time to end time. |
**EdgeLocation** | Pointer to **string** | Edge location of Twilio media servers for the room. |
**EndReason** | Pointer to **string** | Reason the room ended. |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | End time for the room. |
**Links** | Pointer to **map[string]interface{}** | Room subresources. |
**MaxConcurrentParticipants** | Pointer to **int** | Maximum number of participants allowed in the room at the same time allowed by the application settings. |
**MaxParticipants** | Pointer to **int** | Max number of total participants allowed by the application settings. |
**MediaRegion** | Pointer to **string** | Region of Twilio media servers for the room. |
**ProcessingState** | Pointer to **string** | Video Log Analyzer resource state. Will be either `in-progress` or `complete`. |
**RecordingEnabled** | Pointer to **bool** | Boolean indicating if recording is enabled for the room. |
**RoomName** | Pointer to **string** | room friendly name. |
**RoomSid** | Pointer to **string** | Unique identifier for the room. |
**RoomStatus** | Pointer to **string** | Status of the room. |
**RoomType** | Pointer to **string** | Type of room. |
**StatusCallback** | Pointer to **string** | Webhook provided for status callbacks. |
**StatusCallbackMethod** | Pointer to **string** | HTTP method provided for status callback URL. |
**TotalParticipantDurationSec** | Pointer to **int** | Combined amount of participant time in the room. |
**TotalRecordingDurationSec** | Pointer to **int** | Combined amount of recorded seconds for participants in the room. |
**UniqueParticipantIdentities** | Pointer to **int** | Unique number of participant identities. |
**UniqueParticipants** | Pointer to **int** | Number of participants. May include duplicate identities for participants who left and rejoined. |
**Url** | Pointer to **string** | URL for the room resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


