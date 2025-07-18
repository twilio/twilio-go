# AccountsCallsTranscriptionsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRealtimeTranscription**](AccountsCallsTranscriptionsApi.md#CreateRealtimeTranscription) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Transcriptions.json | Create a Transcription
[**UpdateRealtimeTranscription**](AccountsCallsTranscriptionsApi.md#UpdateRealtimeTranscription) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Transcriptions/{Sid}.json | Stop a Transcription using either the SID of the Transcription resource or the &#x60;name&#x60; used when creating the resource



## CreateRealtimeTranscription

> ApiV2010RealtimeTranscription CreateRealtimeTranscription(ctx, CallSidoptional)

Create a Transcription

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
**Track** | [**string**](string.md) | 
**StatusCallbackUrl** | **string** | Absolute URL of the status callback.
**StatusCallbackMethod** | **string** | The http method for the status_callback (one of GET, POST).
**InboundTrackLabel** | **string** | Friendly name given to the Inbound Track
**OutboundTrackLabel** | **string** | Friendly name given to the Outbound Track
**PartialResults** | **bool** | Indicates if partial results are going to be sent to the customer
**LanguageCode** | **string** | Language code used by the transcription engine, specified in [BCP-47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt) format
**TranscriptionEngine** | **string** | Definition of the transcription engine to be used, among those supported by Twilio
**ProfanityFilter** | **bool** | indicates if the server will attempt to filter out profanities, replacing all but the initial character in each filtered word with asterisks
**SpeechModel** | **string** | Recognition model used by the transcription engine, among those supported by the provider
**Hints** | **string** | A Phrase contains words and phrase \\\"hints\\\" so that the speech recognition engine is more likely to recognize them.
**EnableAutomaticPunctuation** | **bool** | The provider will add punctuation to recognition result
**IntelligenceService** | **string** | The SID or unique name of the [Intelligence Service](https://www.twilio.com/docs/conversational-intelligence/api/service-resource) for persisting transcripts and running post-call Language Operators .

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
**Status** | [**string**](string.md) | 

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

