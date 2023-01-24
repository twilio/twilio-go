# InsightsV1VideoRoomSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | Account SID associated with this room. |
**RoomSid** | Pointer to **string** | Unique identifier for the room. |
**RoomName** | Pointer to **string** | Room friendly name. |
**CreateTime** | Pointer to [**time.Time**](time.Time.md) | Creation time of the room. |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | End time for the room. |
**RoomType** | Pointer to [**string**](VideoRoomSummaryEnumRoomType.md) |  |
**RoomStatus** | Pointer to [**string**](VideoRoomSummaryEnumRoomStatus.md) |  |
**StatusCallback** | Pointer to **string** | Webhook provided for status callbacks. |
**StatusCallbackMethod** | Pointer to **string** | HTTP method provided for status callback URL. |
**CreatedMethod** | Pointer to [**string**](VideoRoomSummaryEnumCreatedMethod.md) |  |
**EndReason** | Pointer to [**string**](VideoRoomSummaryEnumEndReason.md) |  |
**MaxParticipants** | Pointer to **int** | Max number of total participants allowed by the application settings. |
**UniqueParticipants** | Pointer to **int** | Number of participants. May include duplicate identities for participants who left and rejoined. |
**UniqueParticipantIdentities** | Pointer to **int** | Unique number of participant identities. |
**ConcurrentParticipants** | Pointer to **int** | Actual number of concurrent participants. |
**MaxConcurrentParticipants** | Pointer to **int** | Maximum number of participants allowed in the room at the same time allowed by the application settings. |
**Codecs** | Pointer to [**[]string**](VideoRoomSummaryEnumCodec.md) | Codecs used by participants in the room. Can be `VP8`, `H264`, or `VP9`. |
**MediaRegion** | Pointer to [**string**](VideoRoomSummaryEnumTwilioRealm.md) |  |
**DurationSec** | Pointer to **int64** | Total room duration from create time to end time. |
**TotalParticipantDurationSec** | Pointer to **int64** | Combined amount of participant time in the room. |
**TotalRecordingDurationSec** | Pointer to **int64** | Combined amount of recorded seconds for participants in the room. |
**ProcessingState** | Pointer to [**string**](VideoRoomSummaryEnumProcessingState.md) |  |
**RecordingEnabled** | Pointer to **bool** | Boolean indicating if recording is enabled for the room. |
**EdgeLocation** | Pointer to [**string**](VideoRoomSummaryEnumEdgeLocation.md) |  |
**Url** | Pointer to **string** | URL for the room resource. |
**Links** | Pointer to **map[string]interface{}** | Room subresources. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


