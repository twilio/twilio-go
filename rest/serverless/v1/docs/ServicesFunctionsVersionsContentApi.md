# ServicesFunctionsVersionsContentApi

All URIs are relative to *https://serverless.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchFunctionVersionContent**](ServicesFunctionsVersionsContentApi.md#FetchFunctionVersionContent) | **Get** /v1/Services/{ServiceSid}/Functions/{FunctionSid}/Versions/{Sid}/Content | 



## FetchFunctionVersionContent

> ServerlessV1FunctionVersionContent FetchFunctionVersionContent(ctx, ServiceSidFunctionSidSid)



Retrieve a the content of a specific Function Version resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Function Version content from.
**FunctionSid** | **string** | The SID of the Function that is the parent of the Function Version content to fetch.
**Sid** | **string** | The SID of the Function Version content to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFunctionVersionContentParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ServerlessV1FunctionVersionContent**](ServerlessV1FunctionVersionContent.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

