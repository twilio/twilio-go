# InsightsV1ConferenceParticipant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParticipantSid** | **string** | SID for this participant. |[optional] 
**Label** | **string** | The user-specified label of this participant. |[optional] 
**ConferenceSid** | **string** | Conference SID. |[optional] 
**CallSid** | **string** | Unique SID identifier of the call. |[optional] 
**AccountSid** | **string** | Account SID. |[optional] 
**CallDirection** | Pointer to [**string**](ConferenceParticipantEnumCallDirection.md) |  |
**From** | **string** | Caller ID of the calling party. |[optional] 
**To** | **string** | Called party. |[optional] 
**CallStatus** | Pointer to [**string**](ConferenceParticipantEnumCallStatus.md) |  |
**CountryCode** | **string** | ISO alpha-2 country code of the participant. |[optional] 
**IsModerator** | **bool** | Boolean. Indicates whether participant had startConferenceOnEnter=true or endConferenceOnExit=true. |[optional] 
**JoinTime** | [**time.Time**](time.Time.md) | ISO 8601 timestamp of participant join event. |[optional] 
**LeaveTime** | [**time.Time**](time.Time.md) | ISO 8601 timestamp of participant leave event. |[optional] 
**DurationSeconds** | Pointer to **int** | Participant durations in seconds. |
**OutboundQueueLength** | Pointer to **int** | Estimated time in queue at call creation. |
**OutboundTimeInQueue** | Pointer to **int** | Actual time in queue (seconds). |
**JitterBufferSize** | Pointer to [**string**](ConferenceParticipantEnumJitterBufferSize.md) |  |
**IsCoach** | **bool** | Boolean. Indicated whether participant was a coach. |[optional] 
**CoachedParticipants** | **[]string** | Call SIDs coached by this participant. |[optional] 
**ParticipantRegion** | Pointer to [**string**](ConferenceParticipantEnumRegion.md) |  |
**ConferenceRegion** | Pointer to [**string**](ConferenceParticipantEnumRegion.md) |  |
**CallType** | Pointer to [**string**](ConferenceParticipantEnumCallType.md) |  |
**ProcessingState** | Pointer to [**string**](ConferenceParticipantEnumProcessingState.md) |  |
**Properties** | Pointer to **interface{}** | Participant properties and metadata. |
**Events** | Pointer to **interface{}** | Object containing information of actions taken by participants. Nested resource URLs. |
**Metrics** | Pointer to **interface{}** | Object. Contains participant quality metrics. |
**Url** | **string** | The URL of this resource. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


