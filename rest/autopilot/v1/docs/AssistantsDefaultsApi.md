# AssistantsDefaultsApi

All URIs are relative to *https://autopilot.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchDefaults**](AssistantsDefaultsApi.md#FetchDefaults) | **Get** /v1/Assistants/{AssistantSid}/Defaults | 
[**UpdateDefaults**](AssistantsDefaultsApi.md#UpdateDefaults) | **Post** /v1/Assistants/{AssistantSid}/Defaults | 



## FetchDefaults

> AutopilotV1Defaults FetchDefaults(ctx, AssistantSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchDefaultsParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AutopilotV1Defaults**](AutopilotV1Defaults.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDefaults

> AutopilotV1Defaults UpdateDefaults(ctx, AssistantSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateDefaultsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Defaults** | [**interface{}**](interface{}.md) | A JSON string that describes the default task links for the `assistant_initiation`, `collect`, and `fallback` situations.

### Return type

[**AutopilotV1Defaults**](AutopilotV1Defaults.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

