# AssistantsApi

All URIs are relative to *https://autopilot.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAssistant**](AssistantsApi.md#CreateAssistant) | **Post** /v1/Assistants | 
[**DeleteAssistant**](AssistantsApi.md#DeleteAssistant) | **Delete** /v1/Assistants/{Sid} | 
[**FetchAssistant**](AssistantsApi.md#FetchAssistant) | **Get** /v1/Assistants/{Sid} | 
[**ListAssistant**](AssistantsApi.md#ListAssistant) | **Get** /v1/Assistants | 
[**UpdateAssistant**](AssistantsApi.md#UpdateAssistant) | **Post** /v1/Assistants/{Sid} | 



## CreateAssistant

> AutopilotV1Assistant CreateAssistant(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateAssistantParams struct


Name | Type | Description
------------- | ------------- | -------------
**CallbackEvents** | **string** | Reserved.
**CallbackUrl** | **string** | Reserved.
**Defaults** | [**interface{}**](interface{}.md) | A JSON object that defines the Assistant&#39;s [default tasks](https://www.twilio.com/docs/autopilot/api/assistant/defaults) for various scenarios, including initiation actions and fallback tasks.
**FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It is not unique and can be up to 255 characters long.
**LogQueries** | **bool** | Whether queries should be logged and kept after training. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;true&#x60;. If &#x60;true&#x60;, queries are stored for 30 days, and then deleted. If &#x60;false&#x60;, no queries are stored.
**StyleSheet** | [**interface{}**](interface{}.md) | The JSON string that defines the Assistant&#39;s [style sheet](https://www.twilio.com/docs/autopilot/api/assistant/stylesheet)
**UniqueName** | **string** | An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. The first 64 characters must be unique.

### Return type

[**AutopilotV1Assistant**](AutopilotV1Assistant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssistant

> DeleteAssistant(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Assistant resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteAssistantParams struct


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


## FetchAssistant

> AutopilotV1Assistant FetchAssistant(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Assistant resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchAssistantParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AutopilotV1Assistant**](AutopilotV1Assistant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssistant

> []AutopilotV1Assistant ListAssistant(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAssistantParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]AutopilotV1Assistant**](AutopilotV1Assistant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssistant

> AutopilotV1Assistant UpdateAssistant(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Assistant resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateAssistantParams struct


Name | Type | Description
------------- | ------------- | -------------
**CallbackEvents** | **string** | Reserved.
**CallbackUrl** | **string** | Reserved.
**Defaults** | [**interface{}**](interface{}.md) | A JSON object that defines the Assistant&#39;s [default tasks](https://www.twilio.com/docs/autopilot/api/assistant/defaults) for various scenarios, including initiation actions and fallback tasks.
**DevelopmentStage** | **string** | A string describing the state of the assistant.
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
**LogQueries** | **bool** | Whether queries should be logged and kept after training. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;true&#x60;. If &#x60;true&#x60;, queries are stored for 30 days, and then deleted. If &#x60;false&#x60;, no queries are stored.
**StyleSheet** | [**interface{}**](interface{}.md) | The JSON string that defines the Assistant&#39;s [style sheet](https://www.twilio.com/docs/autopilot/api/assistant/stylesheet)
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the resource. The first 64 characters must be unique.

### Return type

[**AutopilotV1Assistant**](AutopilotV1Assistant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

