# InsightsV1VideoRoomSummaryVideoParticipantSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | Account SID associated with the room. |
**Codecs** | Pointer to **[]string** | Codecs detected from the participant. |
**DurationSec** | Pointer to **int** | Amount of time in seconds the participant was in the room. |
**EdgeLocation** | Pointer to **string** | Name of the edge location the participant connected to. |
**EndReason** | Pointer to **string** | Reason the participant left the room. |
**ErrorCode** | Pointer to **int** | Errors encountered by the participant. |
**ErrorCodeUrl** | Pointer to **string** | Twilio error code dictionary link. |
**JoinTime** | Pointer to [**time.Time**](time.Time.md) | When the participant joined the room. |
**LeaveTime** | Pointer to [**time.Time**](time.Time.md) | When the participant left the room |
**MediaRegion** | Pointer to **string** | Twilio media region the participant connected to. |
**ParticipantIdentity** | Pointer to **string** | The application-defined string that uniquely identifies the participant within a Room. |
**ParticipantSid** | Pointer to **string** | Unique identifier for the participant. |
**Properties** | Pointer to **map[string]interface{}** | Object containing information about the participant's data from the room. |
**PublisherInfo** | Pointer to **map[string]interface{}** | Object containing information about the SDK name and version. |
**RoomSid** | Pointer to **string** | Unique identifier for the room. |
**Status** | Pointer to **string** | Status of the room. |
**Url** | Pointer to **string** | URL of the participant resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


