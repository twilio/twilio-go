# InsightsV1Conference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | Account SID. |
**ConferenceSid** | Pointer to **string** | Conference SID. |
**ConnectDurationSeconds** | Pointer to **int** | Duration of the conference in seconds. |
**CreateTime** | Pointer to [**time.Time**](time.Time.md) | Conference creation date/time. |
**DetectedIssues** | Pointer to **interface{}** | Potential issues detected during the conference. |
**DurationSeconds** | Pointer to **int** | Conference duration in seconds. |
**EndReason** | Pointer to **string** | Conference end reason. |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | Conference end date/time. |
**EndedBy** | Pointer to **string** | Call SID that ended the conference. |
**FriendlyName** | Pointer to **string** | Custom label for the conference. |
**Links** | Pointer to **map[string]interface{}** | Nested resource URLs. |
**MaxConcurrentParticipants** | Pointer to **int** | Actual maximum concurrent participants. |
**MaxParticipants** | Pointer to **int** | Max participants specified in config. |
**MixerRegion** | Pointer to **string** | Region where the conference was mixed. |
**MixerRegionRequested** | Pointer to **string** | Configuration-requested conference mixer region. |
**ProcessingState** | Pointer to **string** | Processing state for the Conference Summary resource. |
**RecordingEnabled** | Pointer to **bool** | Boolean. Indicates whether recording was enabled. |
**StartTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp in ISO 8601 format when the conference started. |
**Status** | Pointer to **string** | Status of conference |
**TagInfo** | Pointer to **interface{}** | Object. Contains details about conference tags. |
**Tags** | Pointer to **[]string** | Tags for detected conference conditions and participant behaviors. |
**UniqueParticipants** | Pointer to **int** | Unique conference participants. |
**Url** | Pointer to **string** | The URL of this resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


