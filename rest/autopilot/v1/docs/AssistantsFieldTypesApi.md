# AssistantsFieldTypesApi

All URIs are relative to *https://autopilot.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFieldType**](AssistantsFieldTypesApi.md#CreateFieldType) | **Post** /v1/Assistants/{AssistantSid}/FieldTypes | 
[**DeleteFieldType**](AssistantsFieldTypesApi.md#DeleteFieldType) | **Delete** /v1/Assistants/{AssistantSid}/FieldTypes/{Sid} | 
[**FetchFieldType**](AssistantsFieldTypesApi.md#FetchFieldType) | **Get** /v1/Assistants/{AssistantSid}/FieldTypes/{Sid} | 
[**ListFieldType**](AssistantsFieldTypesApi.md#ListFieldType) | **Get** /v1/Assistants/{AssistantSid}/FieldTypes | 
[**UpdateFieldType**](AssistantsFieldTypesApi.md#UpdateFieldType) | **Post** /v1/Assistants/{AssistantSid}/FieldTypes/{Sid} | 



## CreateFieldType

> AutopilotV1FieldType CreateFieldType(ctx, AssistantSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the new resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateFieldTypeParams struct


Name | Type | Description
------------- | ------------- | -------------
**UniqueName** | **string** | An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the `sid` in the URL path to address the resource. The first 64 characters must be unique.
**FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It is not unique and can be up to 255 characters long.

### Return type

[**AutopilotV1FieldType**](AutopilotV1FieldType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFieldType

> DeleteFieldType(ctx, AssistantSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the FieldType resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteFieldTypeParams struct


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


## FetchFieldType

> AutopilotV1FieldType FetchFieldType(ctx, AssistantSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the FieldType resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFieldTypeParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AutopilotV1FieldType**](AutopilotV1FieldType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFieldType

> []AutopilotV1FieldType ListFieldType(ctx, AssistantSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListFieldTypeParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]AutopilotV1FieldType**](AutopilotV1FieldType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFieldType

> AutopilotV1FieldType UpdateFieldType(ctx, AssistantSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the FieldType resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateFieldTypeParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used as an alternative to the `sid` in the URL path to address the resource. The first 64 characters must be unique.

### Return type

[**AutopilotV1FieldType**](AutopilotV1FieldType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

