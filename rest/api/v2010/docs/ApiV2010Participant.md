# ApiV2010Participant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**CallSid** | **string** | The SID of the Call the resource is associated with |[optional] 
**Label** | **string** | The label of this participant |[optional] 
**CallSidToCoach** | **string** | The SID of the participant who is being `coached` |[optional] 
**Coaching** | **bool** | Indicates if the participant changed to coach |[optional] 
**ConferenceSid** | **string** | The SID of the conference the participant is in |[optional] 
**DateCreated** | **string** | The RFC 2822 date and time in GMT that the resource was created |[optional] 
**DateUpdated** | **string** | The RFC 2822 date and time in GMT that the resource was last updated |[optional] 
**EndConferenceOnExit** | **bool** | Whether the conference ends when the participant leaves |[optional] 
**Muted** | **bool** | Whether the participant is muted |[optional] 
**Hold** | **bool** | Whether the participant is on hold |[optional] 
**StartConferenceOnEnter** | **bool** | Whether the conference starts when the participant joins the conference |[optional] 
**Status** | Pointer to [**string**](ParticipantEnumStatus.md) |  |
**Uri** | **string** | The URI of the resource, relative to `https://api.twilio.com` |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


