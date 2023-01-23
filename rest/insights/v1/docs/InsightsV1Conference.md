# InsightsV1Conference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConferenceSid** | Pointer to **string** | The unique SID identifier of the Conference. |
**AccountSid** | Pointer to **string** | The unique SID identifier of the Account. |
**FriendlyName** | Pointer to **string** | Custom label for the conference resource, up to 64 characters. |
**CreateTime** | Pointer to [**time.Time**](time.Time.md) | Conference creation date and time in ISO 8601 format. |
**StartTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp in ISO 8601 format when the conference started. Conferences do not start until at least two participants join, at least one of whom has startConferenceOnEnter=true. |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | Conference end date and time in ISO 8601 format. |
**DurationSeconds** | Pointer to **int** | Conference duration in seconds. |
**ConnectDurationSeconds** | Pointer to **int** | Duration of the between conference start event and conference end event in seconds. |
**Status** | Pointer to [**string**](ConferenceEnumConferenceStatus.md) |  |
**MaxParticipants** | Pointer to **int** | Maximum number of concurrent participants as specified by the configuration. |
**MaxConcurrentParticipants** | Pointer to **int** | Actual maximum number of concurrent participants in the conference. |
**UniqueParticipants** | Pointer to **int** | Unique conference participants based on caller ID. |
**EndReason** | Pointer to [**string**](ConferenceEnumConferenceEndReason.md) |  |
**EndedBy** | Pointer to **string** | Call SID of the participant whose actions ended the conference. |
**MixerRegion** | Pointer to [**string**](ConferenceEnumRegion.md) |  |
**MixerRegionRequested** | Pointer to [**string**](ConferenceEnumRegion.md) |  |
**RecordingEnabled** | Pointer to **bool** | Boolean. Indicates whether recording was enabled at the conference mixer. |
**DetectedIssues** | Pointer to **interface{}** | Potential issues detected by Twilio during the conference. |
**Tags** | Pointer to [**[]string**](ConferenceEnumTag.md) | Tags for detected conference conditions and participant behaviors which may be of interest. |
**TagInfo** | Pointer to **interface{}** | Object. Contains details about conference tags including severity. |
**ProcessingState** | Pointer to [**string**](ConferenceEnumProcessingState.md) |  |
**Url** | Pointer to **string** | The URL of this resource. |
**Links** | Pointer to **map[string]interface{}** | Contains a dictionary of URL links to nested resources of this Conference. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


