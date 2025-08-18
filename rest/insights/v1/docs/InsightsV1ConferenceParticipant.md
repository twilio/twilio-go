# InsightsV1ConferenceParticipant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParticipantSid** | Pointer to **string** | SID for this participant. |
**Label** | Pointer to **string** | The user-specified label of this participant. |
**ConferenceSid** | Pointer to **string** | The unique SID identifier of the Conference. |
**CallSid** | Pointer to **string** | Unique SID identifier of the call that generated the Participant resource. |
**AccountSid** | Pointer to **string** | The unique SID identifier of the Account. |
**CallDirection** | Pointer to [**string**](ConferenceParticipantEnumCallDirection.md) |  |
**From** | Pointer to **string** | Caller ID of the calling party. |
**To** | Pointer to **string** | Called party. |
**CallStatus** | Pointer to [**string**](ConferenceParticipantEnumCallStatus.md) |  |
**CountryCode** | Pointer to **string** | ISO alpha-2 country code of the participant based on caller ID or called number. |
**IsModerator** | Pointer to **bool** | Boolean. Indicates whether participant had startConferenceOnEnter=true or endConferenceOnExit=true. |
**JoinTime** | Pointer to [**time.Time**](time.Time.md) | ISO 8601 timestamp of participant join event. |
**LeaveTime** | Pointer to [**time.Time**](time.Time.md) | ISO 8601 timestamp of participant leave event. |
**DurationSeconds** | Pointer to **int** | Participant durations in seconds. |
**OutboundQueueLength** | Pointer to **int** | Add Participant API only. Estimated time in queue at call creation. |
**OutboundTimeInQueue** | Pointer to **int** | Add Participant API only. Actual time in queue in seconds. |
**JitterBufferSize** | Pointer to [**string**](ConferenceParticipantEnumJitterBufferSize.md) |  |
**IsCoach** | Pointer to **bool** | Boolean. Indicated whether participant was a coach. |
**CoachedParticipants** | Pointer to **[]string** | Call SIDs coached by this participant. |
**ParticipantRegion** | Pointer to [**string**](ConferenceParticipantEnumRegion.md) |  |
**ConferenceRegion** | Pointer to [**string**](ConferenceParticipantEnumRegion.md) |  |
**CallType** | Pointer to [**string**](ConferenceParticipantEnumCallType.md) |  |
**ProcessingState** | Pointer to [**string**](ConferenceParticipantEnumProcessingState.md) |  |
**Properties** | Pointer to **interface{}** | Participant properties and metadata. |
**Events** | Pointer to **interface{}** | Object containing information of actions taken by participants. Contains a dictionary of URL links to nested resources of this Conference Participant. |
**Metrics** | Pointer to **interface{}** | Object. Contains participant call quality metrics. |
**Url** | Pointer to **string** | The URL of this resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


