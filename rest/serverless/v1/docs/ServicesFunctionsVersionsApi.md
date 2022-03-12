# ServicesFunctionsVersionsApi

All URIs are relative to *https://serverless.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchFunctionVersion**](ServicesFunctionsVersionsApi.md#FetchFunctionVersion) | **Get** /v1/Services/{ServiceSid}/Functions/{FunctionSid}/Versions/{Sid} | 
[**ListFunctionVersion**](ServicesFunctionsVersionsApi.md#ListFunctionVersion) | **Get** /v1/Services/{ServiceSid}/Functions/{FunctionSid}/Versions | 



## FetchFunctionVersion

> ServerlessV1FunctionVersion FetchFunctionVersion(ctx, ServiceSidFunctionSidSid)



Retrieve a specific Function Version resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Function Version resource from.
**FunctionSid** | **string** | The SID of the function that is the parent of the Function Version resource to fetch.
**Sid** | **string** | The SID of the Function Version resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFunctionVersionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ServerlessV1FunctionVersion**](ServerlessV1FunctionVersion.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFunctionVersion

> []ServerlessV1FunctionVersion ListFunctionVersion(ctx, ServiceSidFunctionSidoptional)



Retrieve a list of all Function Version resources.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to read the Function Version resources from.
**FunctionSid** | **string** | The SID of the function that is the parent of the Function Version resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListFunctionVersionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ServerlessV1FunctionVersion**](ServerlessV1FunctionVersion.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

