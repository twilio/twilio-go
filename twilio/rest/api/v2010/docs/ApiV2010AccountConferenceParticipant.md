# ApiV2010AccountConferenceParticipant

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**CallSid** | Pointer to **string** | The SID of the Call the resource is associated with |
**CallSidToCoach** | Pointer to **string** | The SID of the participant who is being `coached` |
**Coaching** | Pointer to **bool** | Indicates if the participant changed to coach |
**ConferenceSid** | Pointer to **string** | The SID of the conference the participant is in |
**DateCreated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was created |
**DateUpdated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was last updated |
**EndConferenceOnExit** | Pointer to **bool** | Whether the conference ends when the participant leaves |
**Hold** | Pointer to **bool** | Whether the participant is on hold |
**Label** | Pointer to **string** | The label of this participant |
**Muted** | Pointer to **bool** | Whether the participant is muted |
**StartConferenceOnEnter** | Pointer to **bool** | Whether the conference starts when the participant joins the conference |
**Status** | Pointer to **string** | The status of the participant's call in a session |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com` |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


