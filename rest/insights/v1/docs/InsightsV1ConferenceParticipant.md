# InsightsV1ConferenceParticipant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | Account SID. |
**CallDirection** | Pointer to **string** | Call direction of the participant. |
**CallSid** | Pointer to **string** | Unique SID identifier of the call. |
**CallStatus** | Pointer to **string** | Call status of the call that generated the participant. |
**CallType** | Pointer to **string** | The Call Type of this Conference Participant. |
**CoachedParticipants** | Pointer to **[]string** | Call SIDs coached by this participant. |
**ConferenceRegion** | Pointer to **string** | The Conference Region of this Conference Participant. |
**ConferenceSid** | Pointer to **string** | Conference SID. |
**CountryCode** | Pointer to **string** | ISO alpha-2 country code of the participant. |
**DurationSeconds** | Pointer to **int** | Participant durations in seconds. |
**Events** | Pointer to **interface{}** | Object containing information of actions taken by participants. Nested resource URLs. |
**From** | Pointer to **string** | Caller ID of the calling party. |
**IsCoach** | Pointer to **bool** | Boolean. Indicated whether participant was a coach. |
**IsModerator** | Pointer to **bool** | Boolean. Indicates whether participant had startConferenceOnEnter=true or endConferenceOnExit=true. |
**JitterBufferSize** | Pointer to **string** | The Jitter Buffer Size of this Conference Participant. |
**JoinTime** | Pointer to [**time.Time**](time.Time.md) | ISO 8601 timestamp of participant join event. |
**Label** | Pointer to **string** | The user-specified label of this participant. |
**LeaveTime** | Pointer to [**time.Time**](time.Time.md) | ISO 8601 timestamp of participant leave event. |
**Metrics** | Pointer to **interface{}** | Object. Contains participant quality metrics. |
**OutboundQueueLength** | Pointer to **int** | Estimated time in queue at call creation. |
**OutboundTimeInQueue** | Pointer to **int** | Actual time in queue (seconds). |
**ParticipantRegion** | Pointer to **string** | Twilio region where the participant media originates. |
**ParticipantSid** | Pointer to **string** | SID for this participant. |
**ProcessingState** | Pointer to **string** | Processing state of the Participant Summary. |
**Properties** | Pointer to **interface{}** | Participant properties and metadata. |
**To** | Pointer to **string** | Called party. |
**Url** | Pointer to **string** | The URL of this resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


