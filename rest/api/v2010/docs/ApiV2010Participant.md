# ApiV2010Participant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resource. |
**CallSid** | Pointer to **string** | The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) the Participant resource is associated with. |
**Label** | Pointer to **string** | The user-specified label of this participant, if one was given when the participant was created. This may be used to fetch, update or delete the participant. |
**CallSidToCoach** | Pointer to **string** | The SID of the participant who is being `coached`. The participant being coached is the only participant who can hear the participant who is `coaching`. |
**Coaching** | Pointer to **bool** | Whether the participant is coaching another call. Can be: `true` or `false`. If not present, defaults to `false` unless `call_sid_to_coach` is defined. If `true`, `call_sid_to_coach` must be defined. |
**ConferenceSid** | Pointer to **string** | The SID of the conference the participant is in. |
**DateCreated** | Pointer to **string** | The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to **string** | The date and time in GMT that the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**EndConferenceOnExit** | Pointer to **bool** | Whether the conference ends when the participant leaves. Can be: `true` or `false` and the default is `false`. If `true`, the conference ends and all other participants drop out when the participant leaves. |
**Muted** | Pointer to **bool** | Whether the participant is muted. Can be `true` or `false`. |
**Hold** | Pointer to **bool** | Whether the participant is on hold. Can be `true` or `false`. |
**StartConferenceOnEnter** | Pointer to **bool** | Whether the conference starts when the participant joins the conference, if it has not already started. Can be: `true` or `false` and the default is `true`. If `false` and the conference has not started, the participant is muted and hears background music until another participant starts the conference. |
**Status** | Pointer to [**string**](ParticipantEnumStatus.md) |  |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com`. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


