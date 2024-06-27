# AccountsCallsTranscriptionsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRealtimeTranscription**](AccountsCallsTranscriptionsApi.md#CreateRealtimeTranscription) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Transcriptions.json | 
[**UpdateRealtimeTranscription**](AccountsCallsTranscriptionsApi.md#UpdateRealtimeTranscription) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Transcriptions/{Sid}.json | 



## CreateRealtimeTranscription

> ApiV2010RealtimeTranscription CreateRealtimeTranscription(ctx, CallSidoptional)



Create a Transcription

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) the Transcription resource is associated with.

### Other Parameters

Other parameters are passed through a pointer to a CreateRealtimeTranscriptionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created this Transcription resource.
**Name** | **string** | The user-specified name of this Transcription, if one was given when the Transcription was created. This may be used to stop the Transcription.
**Track** | **string** | 
**StatusCallbackUrl** | **string** | Absolute URL of the status callback.
**StatusCallbackMethod** | **string** | The http method for the status_callback (one of GET, POST).
**InboundTrackLabel** | **string** | Friendly name given to the Inbound Track
**OutboundTrackLabel** | **string** | Friendly name given to the Outbound Track
**PartialResults** | **bool** | Indicates if partial results are going to be send to the customer
**LanguageCode** | **string** | Language code used by the transcription engine, specified in [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) format
**TranscriptionEngine** | **string** | Definition of the transcription engine to be used, between those supported by Twilio
**ProfanityFilter** | **bool** | indicates if the server will attempt to filter out profanities, replacing all but the initial character in each filtered word with asterisks
**SpeechModel** | **string** | Recognition model used by the transcription engine, between those supported by the provider
**Hints** | **string** | A Phrase contains words and phrase \\\"hints\\\" so that the speech recognition engine is more likely to recognize them.
**EnableAutomaticPunctuation** | **bool** | The provider will adds punctuation to recognition result hypotheses
**Parameter1Name** | **string** | Parameter name
**Parameter1Value** | **string** | Parameter value
**Parameter2Name** | **string** | Parameter name
**Parameter2Value** | **string** | Parameter value
**Parameter3Name** | **string** | Parameter name
**Parameter3Value** | **string** | Parameter value
**Parameter4Name** | **string** | Parameter name
**Parameter4Value** | **string** | Parameter value
**Parameter5Name** | **string** | Parameter name
**Parameter5Value** | **string** | Parameter value
**Parameter6Name** | **string** | Parameter name
**Parameter6Value** | **string** | Parameter value
**Parameter7Name** | **string** | Parameter name
**Parameter7Value** | **string** | Parameter value
**Parameter8Name** | **string** | Parameter name
**Parameter8Value** | **string** | Parameter value
**Parameter9Name** | **string** | Parameter name
**Parameter9Value** | **string** | Parameter value
**Parameter10Name** | **string** | Parameter name
**Parameter10Value** | **string** | Parameter value
**Parameter11Name** | **string** | Parameter name
**Parameter11Value** | **string** | Parameter value
**Parameter12Name** | **string** | Parameter name
**Parameter12Value** | **string** | Parameter value
**Parameter13Name** | **string** | Parameter name
**Parameter13Value** | **string** | Parameter value
**Parameter14Name** | **string** | Parameter name
**Parameter14Value** | **string** | Parameter value
**Parameter15Name** | **string** | Parameter name
**Parameter15Value** | **string** | Parameter value
**Parameter16Name** | **string** | Parameter name
**Parameter16Value** | **string** | Parameter value
**Parameter17Name** | **string** | Parameter name
**Parameter17Value** | **string** | Parameter value
**Parameter18Name** | **string** | Parameter name
**Parameter18Value** | **string** | Parameter value
**Parameter19Name** | **string** | Parameter name
**Parameter19Value** | **string** | Parameter value
**Parameter20Name** | **string** | Parameter name
**Parameter20Value** | **string** | Parameter value
**Parameter21Name** | **string** | Parameter name
**Parameter21Value** | **string** | Parameter value
**Parameter22Name** | **string** | Parameter name
**Parameter22Value** | **string** | Parameter value
**Parameter23Name** | **string** | Parameter name
**Parameter23Value** | **string** | Parameter value
**Parameter24Name** | **string** | Parameter name
**Parameter24Value** | **string** | Parameter value
**Parameter25Name** | **string** | Parameter name
**Parameter25Value** | **string** | Parameter value
**Parameter26Name** | **string** | Parameter name
**Parameter26Value** | **string** | Parameter value
**Parameter27Name** | **string** | Parameter name
**Parameter27Value** | **string** | Parameter value
**Parameter28Name** | **string** | Parameter name
**Parameter28Value** | **string** | Parameter value
**Parameter29Name** | **string** | Parameter name
**Parameter29Value** | **string** | Parameter value
**Parameter30Name** | **string** | Parameter name
**Parameter30Value** | **string** | Parameter value
**Parameter31Name** | **string** | Parameter name
**Parameter31Value** | **string** | Parameter value
**Parameter32Name** | **string** | Parameter name
**Parameter32Value** | **string** | Parameter value
**Parameter33Name** | **string** | Parameter name
**Parameter33Value** | **string** | Parameter value
**Parameter34Name** | **string** | Parameter name
**Parameter34Value** | **string** | Parameter value
**Parameter35Name** | **string** | Parameter name
**Parameter35Value** | **string** | Parameter value
**Parameter36Name** | **string** | Parameter name
**Parameter36Value** | **string** | Parameter value
**Parameter37Name** | **string** | Parameter name
**Parameter37Value** | **string** | Parameter value
**Parameter38Name** | **string** | Parameter name
**Parameter38Value** | **string** | Parameter value
**Parameter39Name** | **string** | Parameter name
**Parameter39Value** | **string** | Parameter value
**Parameter40Name** | **string** | Parameter name
**Parameter40Value** | **string** | Parameter value
**Parameter41Name** | **string** | Parameter name
**Parameter41Value** | **string** | Parameter value
**Parameter42Name** | **string** | Parameter name
**Parameter42Value** | **string** | Parameter value
**Parameter43Name** | **string** | Parameter name
**Parameter43Value** | **string** | Parameter value
**Parameter44Name** | **string** | Parameter name
**Parameter44Value** | **string** | Parameter value
**Parameter45Name** | **string** | Parameter name
**Parameter45Value** | **string** | Parameter value
**Parameter46Name** | **string** | Parameter name
**Parameter46Value** | **string** | Parameter value
**Parameter47Name** | **string** | Parameter name
**Parameter47Value** | **string** | Parameter value
**Parameter48Name** | **string** | Parameter name
**Parameter48Value** | **string** | Parameter value
**Parameter49Name** | **string** | Parameter name
**Parameter49Value** | **string** | Parameter value
**Parameter50Name** | **string** | Parameter name
**Parameter50Value** | **string** | Parameter value
**Parameter51Name** | **string** | Parameter name
**Parameter51Value** | **string** | Parameter value
**Parameter52Name** | **string** | Parameter name
**Parameter52Value** | **string** | Parameter value
**Parameter53Name** | **string** | Parameter name
**Parameter53Value** | **string** | Parameter value
**Parameter54Name** | **string** | Parameter name
**Parameter54Value** | **string** | Parameter value
**Parameter55Name** | **string** | Parameter name
**Parameter55Value** | **string** | Parameter value
**Parameter56Name** | **string** | Parameter name
**Parameter56Value** | **string** | Parameter value
**Parameter57Name** | **string** | Parameter name
**Parameter57Value** | **string** | Parameter value
**Parameter58Name** | **string** | Parameter name
**Parameter58Value** | **string** | Parameter value
**Parameter59Name** | **string** | Parameter name
**Parameter59Value** | **string** | Parameter value
**Parameter60Name** | **string** | Parameter name
**Parameter60Value** | **string** | Parameter value
**Parameter61Name** | **string** | Parameter name
**Parameter61Value** | **string** | Parameter value
**Parameter62Name** | **string** | Parameter name
**Parameter62Value** | **string** | Parameter value
**Parameter63Name** | **string** | Parameter name
**Parameter63Value** | **string** | Parameter value
**Parameter64Name** | **string** | Parameter name
**Parameter64Value** | **string** | Parameter value
**Parameter65Name** | **string** | Parameter name
**Parameter65Value** | **string** | Parameter value
**Parameter66Name** | **string** | Parameter name
**Parameter66Value** | **string** | Parameter value
**Parameter67Name** | **string** | Parameter name
**Parameter67Value** | **string** | Parameter value
**Parameter68Name** | **string** | Parameter name
**Parameter68Value** | **string** | Parameter value
**Parameter69Name** | **string** | Parameter name
**Parameter69Value** | **string** | Parameter value
**Parameter70Name** | **string** | Parameter name
**Parameter70Value** | **string** | Parameter value
**Parameter71Name** | **string** | Parameter name
**Parameter71Value** | **string** | Parameter value
**Parameter72Name** | **string** | Parameter name
**Parameter72Value** | **string** | Parameter value
**Parameter73Name** | **string** | Parameter name
**Parameter73Value** | **string** | Parameter value
**Parameter74Name** | **string** | Parameter name
**Parameter74Value** | **string** | Parameter value
**Parameter75Name** | **string** | Parameter name
**Parameter75Value** | **string** | Parameter value
**Parameter76Name** | **string** | Parameter name
**Parameter76Value** | **string** | Parameter value
**Parameter77Name** | **string** | Parameter name
**Parameter77Value** | **string** | Parameter value
**Parameter78Name** | **string** | Parameter name
**Parameter78Value** | **string** | Parameter value
**Parameter79Name** | **string** | Parameter name
**Parameter79Value** | **string** | Parameter value
**Parameter80Name** | **string** | Parameter name
**Parameter80Value** | **string** | Parameter value
**Parameter81Name** | **string** | Parameter name
**Parameter81Value** | **string** | Parameter value
**Parameter82Name** | **string** | Parameter name
**Parameter82Value** | **string** | Parameter value
**Parameter83Name** | **string** | Parameter name
**Parameter83Value** | **string** | Parameter value
**Parameter84Name** | **string** | Parameter name
**Parameter84Value** | **string** | Parameter value
**Parameter85Name** | **string** | Parameter name
**Parameter85Value** | **string** | Parameter value
**Parameter86Name** | **string** | Parameter name
**Parameter86Value** | **string** | Parameter value
**Parameter87Name** | **string** | Parameter name
**Parameter87Value** | **string** | Parameter value
**Parameter88Name** | **string** | Parameter name
**Parameter88Value** | **string** | Parameter value
**Parameter89Name** | **string** | Parameter name
**Parameter89Value** | **string** | Parameter value
**Parameter90Name** | **string** | Parameter name
**Parameter90Value** | **string** | Parameter value
**Parameter91Name** | **string** | Parameter name
**Parameter91Value** | **string** | Parameter value
**Parameter92Name** | **string** | Parameter name
**Parameter92Value** | **string** | Parameter value
**Parameter93Name** | **string** | Parameter name
**Parameter93Value** | **string** | Parameter value
**Parameter94Name** | **string** | Parameter name
**Parameter94Value** | **string** | Parameter value
**Parameter95Name** | **string** | Parameter name
**Parameter95Value** | **string** | Parameter value
**Parameter96Name** | **string** | Parameter name
**Parameter96Value** | **string** | Parameter value
**Parameter97Name** | **string** | Parameter name
**Parameter97Value** | **string** | Parameter value
**Parameter98Name** | **string** | Parameter name
**Parameter98Value** | **string** | Parameter value
**Parameter99Name** | **string** | Parameter name
**Parameter99Value** | **string** | Parameter value

### Return type

[**ApiV2010RealtimeTranscription**](ApiV2010RealtimeTranscription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRealtimeTranscription

> ApiV2010RealtimeTranscription UpdateRealtimeTranscription(ctx, CallSidSidoptional)



Stop a Transcription using either the SID of the Transcription resource or the `name` used when creating the resource

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) the Transcription resource is associated with.
**Sid** | **string** | The SID of the Transcription resource, or the `name` used when creating the resource

### Other Parameters

Other parameters are passed through a pointer to a UpdateRealtimeTranscriptionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created this Transcription resource.
**Status** | **string** | 

### Return type

[**ApiV2010RealtimeTranscription**](ApiV2010RealtimeTranscription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

