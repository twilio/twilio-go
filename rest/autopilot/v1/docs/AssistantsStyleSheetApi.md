# AssistantsStyleSheetApi

All URIs are relative to *https://autopilot.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchStyleSheet**](AssistantsStyleSheetApi.md#FetchStyleSheet) | **Get** /v1/Assistants/{AssistantSid}/StyleSheet | 
[**UpdateStyleSheet**](AssistantsStyleSheetApi.md#UpdateStyleSheet) | **Post** /v1/Assistants/{AssistantSid}/StyleSheet | 



## FetchStyleSheet

> AutopilotV1StyleSheet FetchStyleSheet(ctx, AssistantSid)



Returns Style sheet JSON object for the Assistant

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchStyleSheetParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AutopilotV1StyleSheet**](AutopilotV1StyleSheet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStyleSheet

> AutopilotV1StyleSheet UpdateStyleSheet(ctx, AssistantSidoptional)



Updates the style sheet for an Assistant identified by `assistant_sid`.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateStyleSheetParams struct


Name | Type | Description
------------- | ------------- | -------------
**StyleSheet** | [**interface{}**](interface{}.md) | The JSON string that describes the style sheet object.

### Return type

[**AutopilotV1StyleSheet**](AutopilotV1StyleSheet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

