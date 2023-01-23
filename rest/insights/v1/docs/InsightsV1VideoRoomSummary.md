# InsightsV1VideoRoomSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | Account SID associated with this room. |[optional] 
**RoomSid** | **string** | Unique identifier for the room. |[optional] 
**RoomName** | **string** | room friendly name. |[optional] 
**CreateTime** | [**time.Time**](time.Time.md) | Creation time of the room. |[optional] 
**EndTime** | [**time.Time**](time.Time.md) | End time for the room. |[optional] 
**RoomType** | Pointer to [**string**](VideoRoomSummaryEnumRoomType.md) |  |
**RoomStatus** | Pointer to [**string**](VideoRoomSummaryEnumRoomStatus.md) |  |
**StatusCallback** | **string** | Webhook provided for status callbacks. |[optional] 
**StatusCallbackMethod** | **string** | HTTP method provided for status callback URL. |[optional] 
**CreatedMethod** | Pointer to [**string**](VideoRoomSummaryEnumCreatedMethod.md) |  |
**EndReason** | Pointer to [**string**](VideoRoomSummaryEnumEndReason.md) |  |
**MaxParticipants** | Pointer to **int** | Max number of total participants allowed by the application settings. |
**UniqueParticipants** | Pointer to **int** | Number of participants. May include duplicate identities for participants who left and rejoined. |
**UniqueParticipantIdentities** | Pointer to **int** | Unique number of participant identities. |
**ConcurrentParticipants** | Pointer to **int** | Actual number of concurrent participants. |
**MaxConcurrentParticipants** | Pointer to **int** | Maximum number of participants allowed in the room at the same time allowed by the application settings. |
**Codecs** | Pointer to [**[]string**](VideoRoomSummaryEnumCodec.md) | Codecs used by participants in the room. |
**MediaRegion** | Pointer to [**string**](VideoRoomSummaryEnumTwilioRealm.md) |  |
**DurationSec** | **int64** | Total room duration from create time to end time. |[optional] 
**TotalParticipantDurationSec** | **int64** | Combined amount of participant time in the room. |[optional] 
**TotalRecordingDurationSec** | **int64** | Combined amount of recorded seconds for participants in the room. |[optional] 
**ProcessingState** | Pointer to [**string**](VideoRoomSummaryEnumProcessingState.md) |  |
**RecordingEnabled** | **bool** | Boolean indicating if recording is enabled for the room. |[optional] 
**EdgeLocation** | Pointer to [**string**](VideoRoomSummaryEnumEdgeLocation.md) |  |
**Url** | **string** | URL for the room resource. |[optional] 
**Links** | **map[string]interface{}** | Room subresources. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


