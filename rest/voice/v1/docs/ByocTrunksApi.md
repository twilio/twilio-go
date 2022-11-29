# ByocTrunksApi

All URIs are relative to *https://voice.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateByocTrunk**](ByocTrunksApi.md#CreateByocTrunk) | **Post** /v1/ByocTrunks | 
[**DeleteByocTrunk**](ByocTrunksApi.md#DeleteByocTrunk) | **Delete** /v1/ByocTrunks/{Sid} | 
[**FetchByocTrunk**](ByocTrunksApi.md#FetchByocTrunk) | **Get** /v1/ByocTrunks/{Sid} | 
[**ListByocTrunk**](ByocTrunksApi.md#ListByocTrunk) | **Get** /v1/ByocTrunks | 
[**UpdateByocTrunk**](ByocTrunksApi.md#UpdateByocTrunk) | **Post** /v1/ByocTrunks/{Sid} | 



## CreateByocTrunk

> VoiceV1ByocTrunk CreateByocTrunk(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateByocTrunkParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
**VoiceUrl** | **string** | The URL we should call when the BYOC Trunk receives a call.
**VoiceMethod** | **string** | The HTTP method we should use to call `voice_url`. Can be: `GET` or `POST`.
**VoiceFallbackUrl** | **string** | The URL that we should call when an error occurs while retrieving or executing the TwiML from `voice_url`.
**VoiceFallbackMethod** | **string** | The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`.
**StatusCallbackUrl** | **string** | The URL that we should call to pass status parameters (such as call ended) to your application.
**StatusCallbackMethod** | **string** | The HTTP method we should use to call `status_callback_url`. Can be: `GET` or `POST`.
**CnamLookupEnabled** | **bool** | Whether Caller ID Name (CNAM) lookup is enabled for the trunk. If enabled, all inbound calls to the BYOC Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information.
**ConnectionPolicySid** | **string** | The SID of the Connection Policy that Twilio will use when routing traffic to your communications infrastructure.
**FromDomainSid** | **string** | The SID of the SIP Domain that should be used in the `From` header of originating calls sent to your SIP infrastructure. If your SIP infrastructure allows users to \\\"call back\\\" an incoming call, configure this with a [SIP Domain](https://www.twilio.com/docs/voice/api/sending-sip) to ensure proper routing. If not configured, the from domain will default to \\\"sip.twilio.com\\\".

### Return type

[**VoiceV1ByocTrunk**](VoiceV1ByocTrunk.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteByocTrunk

> DeleteByocTrunk(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the BYOC Trunk resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteByocTrunkParams struct


Name | Type | Description
------------- | ------------- | -------------

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


## FetchByocTrunk

> VoiceV1ByocTrunk FetchByocTrunk(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the BYOC Trunk resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchByocTrunkParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VoiceV1ByocTrunk**](VoiceV1ByocTrunk.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListByocTrunk

> []VoiceV1ByocTrunk ListByocTrunk(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListByocTrunkParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VoiceV1ByocTrunk**](VoiceV1ByocTrunk.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateByocTrunk

> VoiceV1ByocTrunk UpdateByocTrunk(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the BYOC Trunk resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateByocTrunkParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
**VoiceUrl** | **string** | The URL we should call when the BYOC Trunk receives a call.
**VoiceMethod** | **string** | The HTTP method we should use to call `voice_url`
**VoiceFallbackUrl** | **string** | The URL that we should call when an error occurs while retrieving or executing the TwiML requested by `voice_url`.
**VoiceFallbackMethod** | **string** | The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`.
**StatusCallbackUrl** | **string** | The URL that we should call to pass status parameters (such as call ended) to your application.
**StatusCallbackMethod** | **string** | The HTTP method we should use to call `status_callback_url`. Can be: `GET` or `POST`.
**CnamLookupEnabled** | **bool** | Whether Caller ID Name (CNAM) lookup is enabled for the trunk. If enabled, all inbound calls to the BYOC Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information.
**ConnectionPolicySid** | **string** | The SID of the Connection Policy that Twilio will use when routing traffic to your communications infrastructure.
**FromDomainSid** | **string** | The SID of the SIP Domain that should be used in the `From` header of originating calls sent to your SIP infrastructure. If your SIP infrastructure allows users to \\\"call back\\\" an incoming call, configure this with a [SIP Domain](https://www.twilio.com/docs/voice/api/sending-sip) to ensure proper routing. If not configured, the from domain will default to \\\"sip.twilio.com\\\".

### Return type

[**VoiceV1ByocTrunk**](VoiceV1ByocTrunk.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

