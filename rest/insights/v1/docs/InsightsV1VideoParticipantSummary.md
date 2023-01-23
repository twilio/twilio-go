# InsightsV1VideoParticipantSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParticipantSid** | Pointer to **string** | Unique identifier for the participant. |
**ParticipantIdentity** | Pointer to **string** | The application-defined string that uniquely identifies the participant within a Room. |
**JoinTime** | Pointer to [**time.Time**](time.Time.md) | When the participant joined the room. |
**LeaveTime** | Pointer to [**time.Time**](time.Time.md) | When the participant left the room. |
**DurationSec** | Pointer to **int64** | Amount of time in seconds the participant was in the room. |
**AccountSid** | Pointer to **string** | Account SID associated with the room. |
**RoomSid** | Pointer to **string** | Unique identifier for the room. |
**Status** | Pointer to [**string**](VideoParticipantSummaryEnumRoomStatus.md) |  |
**Codecs** | Pointer to [**[]string**](VideoParticipantSummaryEnumCodec.md) | Codecs detected from the participant. Can be `VP8`, `H264`, or `VP9`. |
**EndReason** | Pointer to **string** | Reason the participant left the room. See [the list of possible values here](https://www.twilio.com/docs/video/video-log-analyzer/video-log-analyzer-api#end_reason). |
**ErrorCode** | Pointer to **int** | Errors encountered by the participant. |
**ErrorCodeUrl** | Pointer to **string** | Twilio error code dictionary link. |
**MediaRegion** | Pointer to [**string**](VideoParticipantSummaryEnumTwilioRealm.md) |  |
**Properties** | Pointer to **interface{}** | Object containing information about the participant's data from the room. See [below](https://www.twilio.com/docs/video/video-log-analyzer/video-log-analyzer-api#properties) for more information. |
**EdgeLocation** | Pointer to [**string**](VideoParticipantSummaryEnumEdgeLocation.md) |  |
**PublisherInfo** | Pointer to **interface{}** | Object containing information about the SDK name and version. See [below](https://www.twilio.com/docs/video/video-log-analyzer/video-log-analyzer-api#publisher_info) for more information. |
**Url** | Pointer to **string** | URL of the participant resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


