# InsightsV1Conference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConferenceSid** | Pointer to **string** | Conference SID. |
**AccountSid** | Pointer to **string** | Account SID. |
**FriendlyName** | Pointer to **string** | Custom label for the conference. |
**CreateTime** | Pointer to [**time.Time**](time.Time.md) | Conference creation date/time. |
**StartTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp in ISO 8601 format when the conference started. |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | Conference end date/time. |
**DurationSeconds** | Pointer to **int** | Conference duration in seconds. |
**ConnectDurationSeconds** | Pointer to **int** | Duration of the conference in seconds. |
**Status** | Pointer to [**string**](ConferenceEnumConferenceStatus.md) |  |
**MaxParticipants** | Pointer to **int** | Max participants specified in config. |
**MaxConcurrentParticipants** | Pointer to **int** | Actual maximum concurrent participants. |
**UniqueParticipants** | Pointer to **int** | Unique conference participants. |
**EndReason** | Pointer to [**string**](ConferenceEnumConferenceEndReason.md) |  |
**EndedBy** | Pointer to **string** | Call SID that ended the conference. |
**MixerRegion** | Pointer to [**string**](ConferenceEnumRegion.md) |  |
**MixerRegionRequested** | Pointer to [**string**](ConferenceEnumRegion.md) |  |
**RecordingEnabled** | Pointer to **bool** | Boolean. Indicates whether recording was enabled. |
**DetectedIssues** | Pointer to **interface{}** | Potential issues detected during the conference. |
**Tags** | Pointer to [**[]string**](ConferenceEnumTag.md) | Tags for detected conference conditions and participant behaviors. |
**TagInfo** | Pointer to **interface{}** | Object. Contains details about conference tags. |
**ProcessingState** | Pointer to [**string**](ConferenceEnumProcessingState.md) |  |
**Url** | Pointer to **string** | The URL of this resource. |
**Links** | Pointer to **map[string]interface{}** | Nested resource URLs. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


