# AccountsConferencesParticipantsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateParticipant**](AccountsConferencesParticipantsApi.md#CreateParticipant) | **Post** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Participants.json | 
[**DeleteParticipant**](AccountsConferencesParticipantsApi.md#DeleteParticipant) | **Delete** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Participants/{CallSid}.json | 
[**FetchParticipant**](AccountsConferencesParticipantsApi.md#FetchParticipant) | **Get** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Participants/{CallSid}.json | 
[**ListParticipant**](AccountsConferencesParticipantsApi.md#ListParticipant) | **Get** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Participants.json | 
[**UpdateParticipant**](AccountsConferencesParticipantsApi.md#UpdateParticipant) | **Post** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Participants/{CallSid}.json | 



## CreateParticipant

> ApiV2010Participant CreateParticipant(ctx, ConferenceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConferenceSid** | **string** | The SID of the participant&#39;s conference.

### Other Parameters

Other parameters are passed through a pointer to a CreateParticipantParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**Beep** | **string** | Whether to play a notification beep to the conference when the participant joins. Can be: &#x60;true&#x60;, &#x60;false&#x60;, &#x60;onEnter&#x60;, or &#x60;onExit&#x60;. The default value is &#x60;true&#x60;.
**Byoc** | **string** | The SID of a BYOC (Bring Your Own Carrier) trunk to route this call with. Note that &#x60;byoc&#x60; is only meaningful when &#x60;to&#x60; is a phone number; it will otherwise be ignored. (Beta)
**CallReason** | **string** | The Reason for the outgoing call. Use it to specify the purpose of the call that is presented on the called party&#39;s phone. (Branded Calls Beta)
**CallSidToCoach** | **string** | The SID of the participant who is being &#x60;coached&#x60;. The participant being coached is the only participant who can hear the participant who is &#x60;coaching&#x60;.
**CallerId** | **string** | The phone number, Client identifier, or username portion of SIP address that made this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). Client identifiers are formatted &#x60;client:name&#x60;. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the &#x60;to&#x60; parameter is a phone number, &#x60;callerId&#x60; must also be a phone number. If &#x60;to&#x60; is sip address, this value of &#x60;callerId&#x60; should be a username portion to be used to populate the From header that is passed to the SIP endpoint.
**Coaching** | **bool** | Whether the participant is coaching another call. Can be: &#x60;true&#x60; or &#x60;false&#x60;. If not present, defaults to &#x60;false&#x60; unless &#x60;call_sid_to_coach&#x60; is defined. If &#x60;true&#x60;, &#x60;call_sid_to_coach&#x60; must be defined.
**ConferenceRecord** | **string** | Whether to record the conference the participant is joining. Can be: &#x60;true&#x60;, &#x60;false&#x60;, &#x60;record-from-start&#x60;, and &#x60;do-not-record&#x60;. The default value is &#x60;false&#x60;.
**ConferenceRecordingStatusCallback** | **string** | The URL we should call using the &#x60;conference_recording_status_callback_method&#x60; when the conference recording is available.
**ConferenceRecordingStatusCallbackEvent** | **[]string** | The conference recording state changes that generate a call to &#x60;conference_recording_status_callback&#x60;. Can be: &#x60;in-progress&#x60;, &#x60;completed&#x60;, &#x60;failed&#x60;, and &#x60;absent&#x60;. Separate multiple values with a space, ex: &#x60;&#39;in-progress completed failed&#39;&#x60;
**ConferenceRecordingStatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;conference_recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**ConferenceStatusCallback** | **string** | The URL we should call using the &#x60;conference_status_callback_method&#x60; when the conference events in &#x60;conference_status_callback_event&#x60; occur. Only the value set by the first participant to join the conference is used. Subsequent &#x60;conference_status_callback&#x60; values are ignored.
**ConferenceStatusCallbackEvent** | **[]string** | The conference state changes that should generate a call to &#x60;conference_status_callback&#x60;. Can be: &#x60;start&#x60;, &#x60;end&#x60;, &#x60;join&#x60;, &#x60;leave&#x60;, &#x60;mute&#x60;, &#x60;hold&#x60;, &#x60;modify&#x60;, &#x60;speaker&#x60;, and &#x60;announcement&#x60;. Separate multiple values with a space. Defaults to &#x60;start end&#x60;.
**ConferenceStatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;conference_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**ConferenceTrim** | **string** | Whether to trim leading and trailing silence from your recorded conference audio files. Can be: &#x60;trim-silence&#x60; or &#x60;do-not-trim&#x60; and defaults to &#x60;trim-silence&#x60;.
**EarlyMedia** | **bool** | Whether to allow an agent to hear the state of the outbound call, including ringing or disconnect messages. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;true&#x60;.
**EndConferenceOnExit** | **bool** | Whether to end the conference when the participant leaves. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;.
**From** | **string** | The phone number, Client identifier, or username portion of SIP address that made this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). Client identifiers are formatted &#x60;client:name&#x60;. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the &#x60;to&#x60; parameter is a phone number, &#x60;from&#x60; must also be a phone number. If &#x60;to&#x60; is sip address, this value of &#x60;from&#x60; should be a username portion to be used to populate the P-Asserted-Identity header that is passed to the SIP endpoint.
**JitterBufferSize** | **string** | Jitter buffer size for the connecting participant. Twilio will use this setting to apply Jitter Buffer before participant&#39;s audio is mixed into the conference. Can be: &#x60;off&#x60;, &#x60;small&#x60;, &#x60;medium&#x60;, and &#x60;large&#x60;. Default to &#x60;large&#x60;.
**Label** | **string** | A label for this participant. If one is supplied, it may subsequently be used to fetch, update or delete the participant.
**MaxParticipants** | **int** | The maximum number of participants in the conference. Can be a positive integer from &#x60;2&#x60; to &#x60;250&#x60;. The default value is &#x60;250&#x60;.
**Muted** | **bool** | Whether the agent is muted in the conference. Can be &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
**Record** | **bool** | Whether to record the participant and their conferences, including the time between conferences. Can be &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
**RecordingChannels** | **string** | The recording channels for the final recording. Can be: &#x60;mono&#x60; or &#x60;dual&#x60; and the default is &#x60;mono&#x60;.
**RecordingStatusCallback** | **string** | The URL that we should call using the &#x60;recording_status_callback_method&#x60; when the recording status changes.
**RecordingStatusCallbackEvent** | **[]string** | The recording state changes that should generate a call to &#x60;recording_status_callback&#x60;. Can be: &#x60;started&#x60;, &#x60;in-progress&#x60;, &#x60;paused&#x60;, &#x60;resumed&#x60;, &#x60;stopped&#x60;, &#x60;completed&#x60;, &#x60;failed&#x60;, and &#x60;absent&#x60;. Separate multiple values with a space, ex: &#x60;&#39;in-progress completed failed&#39;&#x60;.
**RecordingStatusCallbackMethod** | **string** | The HTTP method we should use when we call &#x60;recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**RecordingTrack** | **string** | The audio track to record for the call. Can be: &#x60;inbound&#x60;, &#x60;outbound&#x60; or &#x60;both&#x60;. The default is &#x60;both&#x60;. &#x60;inbound&#x60; records the audio that is received by Twilio. &#x60;outbound&#x60; records the audio that is sent from Twilio. &#x60;both&#x60; records the audio that is received and sent by Twilio.
**Region** | **string** | The [region](https://support.twilio.com/hc/en-us/articles/223132167-How-global-low-latency-routing-and-region-selection-work-for-conferences-and-Client-calls) where we should mix the recorded audio. Can be:&#x60;us1&#x60;, &#x60;ie1&#x60;, &#x60;de1&#x60;, &#x60;sg1&#x60;, &#x60;br1&#x60;, &#x60;au1&#x60;, or &#x60;jp1&#x60;.
**SipAuthPassword** | **string** | The SIP password for authentication.
**SipAuthUsername** | **string** | The SIP username used for authentication.
**StartConferenceOnEnter** | **bool** | Whether to start the conference when the participant joins, if it has not already started. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. If &#x60;false&#x60; and the conference has not started, the participant is muted and hears background music until another participant starts the conference.
**StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application.
**StatusCallbackEvent** | **[]string** | The conference state changes that should generate a call to &#x60;status_callback&#x60;. Can be: &#x60;initiated&#x60;, &#x60;ringing&#x60;, &#x60;answered&#x60;, and &#x60;completed&#x60;. Separate multiple values with a space. The default value is &#x60;completed&#x60;.
**StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; and &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**TimeLimit** | **int** | The maximum duration of the call in seconds. Constraints depend on account and configuration.
**Timeout** | **int** | The number of seconds that we should allow the phone to ring before assuming there is no answer. Can be an integer between &#x60;5&#x60; and &#x60;600&#x60;, inclusive. The default value is &#x60;60&#x60;. We always add a 5-second timeout buffer to outgoing calls, so  value of 10 would result in an actual timeout that was closer to 15 seconds.
**To** | **string** | The phone number, SIP address, or Client identifier that received this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). SIP addresses are formatted as &#x60;sip:name@company.com&#x60;. Client identifiers are formatted &#x60;client:name&#x60;. [Custom parameters](https://www.twilio.com/docs/voice/api/conference-participant-resource#custom-parameters) may also be specified.
**WaitMethod** | **string** | The HTTP method we should use to call &#x60;wait_url&#x60;. Can be &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. When using a static audio file, this should be &#x60;GET&#x60; so that we can cache the file.
**WaitUrl** | **string** | The URL we should call using the &#x60;wait_method&#x60; for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic).

### Return type

[**ApiV2010Participant**](ApiV2010Participant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteParticipant

> DeleteParticipant(ctx, ConferenceSidCallSidoptional)



Kick a participant from a given conference

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConferenceSid** | **string** | The SID of the conference with the participants to delete.
**CallSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID or label of the participant to delete. Non URL safe characters in a label must be percent encoded, for example, a space character is represented as %20.

### Other Parameters

Other parameters are passed through a pointer to a DeleteParticipantParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resources to delete.

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchParticipant

> ApiV2010Participant FetchParticipant(ctx, ConferenceSidCallSidoptional)



Fetch an instance of a participant

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConferenceSid** | **string** | The SID of the conference with the participant to fetch.
**CallSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID or label of the participant to fetch. Non URL safe characters in a label must be percent encoded, for example, a space character is represented as %20.

### Other Parameters

Other parameters are passed through a pointer to a FetchParticipantParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resource to fetch.

### Return type

[**ApiV2010Participant**](ApiV2010Participant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListParticipant

> []ApiV2010Participant ListParticipant(ctx, ConferenceSidoptional)



Retrieve a list of participants belonging to the account used to make the request

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConferenceSid** | **string** | The SID of the conference with the participants to read.

### Other Parameters

Other parameters are passed through a pointer to a ListParticipantParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resources to read.
**Muted** | **bool** | Whether to return only participants that are muted. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
**Hold** | **bool** | Whether to return only participants that are on hold. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
**Coaching** | **bool** | Whether to return only participants who are coaching another call. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010Participant**](ApiV2010Participant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateParticipant

> ApiV2010Participant UpdateParticipant(ctx, ConferenceSidCallSidoptional)



Update the properties of the participant

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConferenceSid** | **string** | The SID of the conference with the participant to update.
**CallSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID or label of the participant to update. Non URL safe characters in a label must be percent encoded, for example, a space character is represented as %20.

### Other Parameters

Other parameters are passed through a pointer to a UpdateParticipantParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resources to update.
**AnnounceMethod** | **string** | The HTTP method we should use to call &#x60;announce_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**AnnounceUrl** | **string** | The URL we call using the &#x60;announce_method&#x60; for an announcement to the participant. The URL must return an MP3 file, a WAV file, or a TwiML document that contains &#x60;&lt;Play&gt;&#x60; or &#x60;&lt;Say&gt;&#x60; commands.
**BeepOnExit** | **bool** | Whether to play a notification beep to the conference when the participant exits. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
**CallSidToCoach** | **string** | The SID of the participant who is being &#x60;coached&#x60;. The participant being coached is the only participant who can hear the participant who is &#x60;coaching&#x60;.
**Coaching** | **bool** | Whether the participant is coaching another call. Can be: &#x60;true&#x60; or &#x60;false&#x60;. If not present, defaults to &#x60;false&#x60; unless &#x60;call_sid_to_coach&#x60; is defined. If &#x60;true&#x60;, &#x60;call_sid_to_coach&#x60; must be defined.
**EndConferenceOnExit** | **bool** | Whether to end the conference when the participant leaves. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;.
**Hold** | **bool** | Whether the participant should be on hold. Can be: &#x60;true&#x60; or &#x60;false&#x60;. &#x60;true&#x60; puts the participant on hold, and &#x60;false&#x60; lets them rejoin the conference.
**HoldMethod** | **string** | The HTTP method we should use to call &#x60;hold_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;GET&#x60;.
**HoldUrl** | **string** | The URL we call using the &#x60;hold_method&#x60; for  music that plays when the participant is on hold. The URL may return an MP3 file, a WAV file, or a TwiML document that contains the &#x60;&lt;Play&gt;&#x60;, &#x60;&lt;Say&gt;&#x60; or &#x60;&lt;Redirect&gt;&#x60; commands.
**Muted** | **bool** | Whether the participant should be muted. Can be &#x60;true&#x60; or &#x60;false&#x60;. &#x60;true&#x60; will mute the participant, and &#x60;false&#x60; will un-mute them. Anything value other than &#x60;true&#x60; or &#x60;false&#x60; is interpreted as &#x60;false&#x60;.
**WaitMethod** | **string** | The HTTP method we should use to call &#x60;wait_url&#x60;. Can be &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. When using a static audio file, this should be &#x60;GET&#x60; so that we can cache the file.
**WaitUrl** | **string** | The URL we should call using the &#x60;wait_method&#x60; for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic).

### Return type

[**ApiV2010Participant**](ApiV2010Participant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

