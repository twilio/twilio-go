# AssistantsFieldTypesFieldValuesApi

All URIs are relative to *https://autopilot.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFieldValue**](AssistantsFieldTypesFieldValuesApi.md#CreateFieldValue) | **Post** /v1/Assistants/{AssistantSid}/FieldTypes/{FieldTypeSid}/FieldValues | 
[**DeleteFieldValue**](AssistantsFieldTypesFieldValuesApi.md#DeleteFieldValue) | **Delete** /v1/Assistants/{AssistantSid}/FieldTypes/{FieldTypeSid}/FieldValues/{Sid} | 
[**FetchFieldValue**](AssistantsFieldTypesFieldValuesApi.md#FetchFieldValue) | **Get** /v1/Assistants/{AssistantSid}/FieldTypes/{FieldTypeSid}/FieldValues/{Sid} | 
[**ListFieldValue**](AssistantsFieldTypesFieldValuesApi.md#ListFieldValue) | **Get** /v1/Assistants/{AssistantSid}/FieldTypes/{FieldTypeSid}/FieldValues | 



## CreateFieldValue

> AutopilotV1FieldValue CreateFieldValue(ctx, AssistantSidFieldTypeSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the FieldType associated with the new resource.
**FieldTypeSid** | **string** | The SID of the Field Type associated with the Field Value.

### Other Parameters

Other parameters are passed through a pointer to a CreateFieldValueParams struct


Name | Type | Description
------------- | ------------- | -------------
**Language** | **string** | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) tag that specifies the language of the value. Currently supported tags: `en-US`
**Value** | **string** | The Field Value data.
**SynonymOf** | **string** | The string value that indicates which word the field value is a synonym of.

### Return type

[**AutopilotV1FieldValue**](AutopilotV1FieldValue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFieldValue

> DeleteFieldValue(ctx, AssistantSidFieldTypeSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the FieldType associated with the resources to delete.
**FieldTypeSid** | **string** | The SID of the Field Type associated with the Field Value to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the FieldValue resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteFieldValueParams struct


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


## FetchFieldValue

> AutopilotV1FieldValue FetchFieldValue(ctx, AssistantSidFieldTypeSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the FieldType associated with the resource to fetch.
**FieldTypeSid** | **string** | The SID of the Field Type associated with the Field Value to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the FieldValue resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFieldValueParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AutopilotV1FieldValue**](AutopilotV1FieldValue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFieldValue

> []AutopilotV1FieldValue ListFieldValue(ctx, AssistantSidFieldTypeSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the FieldType associated with the resources to read.
**FieldTypeSid** | **string** | The SID of the Field Type associated with the Field Value to read.

### Other Parameters

Other parameters are passed through a pointer to a ListFieldValueParams struct


Name | Type | Description
------------- | ------------- | -------------
**Language** | **string** | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) tag that specifies the language of the value. Currently supported tags: `en-US`
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]AutopilotV1FieldValue**](AutopilotV1FieldValue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

