# InsightsV1Conference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConferenceSid** | **string** | Conference SID. |[optional] 
**AccountSid** | **string** | Account SID. |[optional] 
**FriendlyName** | **string** | Custom label for the conference. |[optional] 
**CreateTime** | [**time.Time**](time.Time.md) | Conference creation date/time. |[optional] 
**StartTime** | [**time.Time**](time.Time.md) | Timestamp in ISO 8601 format when the conference started. |[optional] 
**EndTime** | [**time.Time**](time.Time.md) | Conference end date/time. |[optional] 
**DurationSeconds** | Pointer to **int** | Conference duration in seconds. |
**ConnectDurationSeconds** | Pointer to **int** | Duration of the conference in seconds. |
**Status** | Pointer to [**string**](ConferenceEnumConferenceStatus.md) |  |
**MaxParticipants** | Pointer to **int** | Max participants specified in config. |
**MaxConcurrentParticipants** | Pointer to **int** | Actual maximum concurrent participants. |
**UniqueParticipants** | Pointer to **int** | Unique conference participants. |
**EndReason** | Pointer to [**string**](ConferenceEnumConferenceEndReason.md) |  |
**EndedBy** | **string** | Call SID that ended the conference. |[optional] 
**MixerRegion** | Pointer to [**string**](ConferenceEnumRegion.md) |  |
**MixerRegionRequested** | Pointer to [**string**](ConferenceEnumRegion.md) |  |
**RecordingEnabled** | **bool** | Boolean. Indicates whether recording was enabled. |[optional] 
**DetectedIssues** | Pointer to **interface{}** | Potential issues detected during the conference. |
**Tags** | Pointer to [**[]string**](ConferenceEnumTag.md) | Tags for detected conference conditions and participant behaviors. |
**TagInfo** | Pointer to **interface{}** | Object. Contains details about conference tags. |
**ProcessingState** | Pointer to [**string**](ConferenceEnumProcessingState.md) |  |
**Url** | **string** | The URL of this resource. |[optional] 
**Links** | **map[string]interface{}** | Nested resource URLs. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


