# InlineObject7

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnounceMethod** | **string** | The HTTP method we should use to call &#x60;announce_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | [optional] 
**AnnounceUrl** | **string** | The URL we call using the &#x60;announce_method&#x60; for an announcement to the participant. The URL must return an MP3 file, a WAV file, or a TwiML document that contains &#x60;&lt;Play&gt;&#x60; or &#x60;&lt;Say&gt;&#x60; commands. | [optional] 
**BeepOnExit** | **bool** | Whether to play a notification beep to the conference when the participant exits. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | [optional] 
**CallSidToCoach** | **string** | The SID of the participant who is being &#x60;coached&#x60;. The participant being coached is the only participant who can hear the participant who is &#x60;coaching&#x60;. | [optional] 
**Coaching** | **bool** | Whether the participant is coaching another call. Can be: &#x60;true&#x60; or &#x60;false&#x60;. If not present, defaults to &#x60;false&#x60; unless &#x60;call_sid_to_coach&#x60; is defined. If &#x60;true&#x60;, &#x60;call_sid_to_coach&#x60; must be defined. | [optional] 
**EndConferenceOnExit** | **bool** | Whether to end the conference when the participant leaves. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;. | [optional] 
**Hold** | **bool** | Whether the participant should be on hold. Can be: &#x60;true&#x60; or &#x60;false&#x60;. &#x60;true&#x60; puts the participant on hold, and &#x60;false&#x60; lets them rejoin the conference. | [optional] 
**HoldMethod** | **string** | The HTTP method we should use to call &#x60;hold_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;GET&#x60;. | [optional] 
**HoldUrl** | **string** | The URL we call using the &#x60;hold_method&#x60; for  music that plays when the participant is on hold. The URL may return an MP3 file, a WAV file, or a TwiML document that contains the &#x60;&lt;Play&gt;&#x60;, &#x60;&lt;Say&gt;&#x60; or &#x60;&lt;Redirect&gt;&#x60; commands. | [optional] 
**Muted** | **bool** | Whether the participant should be muted. Can be &#x60;true&#x60; or &#x60;false&#x60;. &#x60;true&#x60; will mute the participant, and &#x60;false&#x60; will un-mute them. Anything value other than &#x60;true&#x60; or &#x60;false&#x60; is interpreted as &#x60;false&#x60;. | [optional] 
**WaitMethod** | **string** | The HTTP method we should use to call &#x60;wait_url&#x60;. Can be &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. When using a static audio file, this should be &#x60;GET&#x60; so that we can cache the file. | [optional] 
**WaitUrl** | **string** | The URL we should call using the &#x60;wait_method&#x60; for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


