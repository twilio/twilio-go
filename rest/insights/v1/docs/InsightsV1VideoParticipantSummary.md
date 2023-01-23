# InsightsV1VideoParticipantSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParticipantSid** | **string** | Unique identifier for the participant. |[optional] 
**ParticipantIdentity** | **string** | The application-defined string that uniquely identifies the participant within a Room. |[optional] 
**JoinTime** | [**time.Time**](time.Time.md) | When the participant joined the room. |[optional] 
**LeaveTime** | [**time.Time**](time.Time.md) | When the participant left the room |[optional] 
**DurationSec** | **int64** | Amount of time in seconds the participant was in the room. |[optional] 
**AccountSid** | **string** | Account SID associated with the room. |[optional] 
**RoomSid** | **string** | Unique identifier for the room. |[optional] 
**Status** | Pointer to [**string**](VideoParticipantSummaryEnumRoomStatus.md) |  |
**Codecs** | Pointer to [**[]string**](VideoParticipantSummaryEnumCodec.md) | Codecs detected from the participant. |
**EndReason** | **string** | Reason the participant left the room. |[optional] 
**ErrorCode** | Pointer to **int** | Errors encountered by the participant. |
**ErrorCodeUrl** | **string** | Twilio error code dictionary link. |[optional] 
**MediaRegion** | Pointer to [**string**](VideoParticipantSummaryEnumTwilioRealm.md) |  |
**Properties** | Pointer to **interface{}** | Object containing information about the participant's data from the room. |
**EdgeLocation** | Pointer to [**string**](VideoParticipantSummaryEnumEdgeLocation.md) |  |
**PublisherInfo** | Pointer to **interface{}** | Object containing information about the SDK name and version. |
**Url** | **string** | URL of the participant resource. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


