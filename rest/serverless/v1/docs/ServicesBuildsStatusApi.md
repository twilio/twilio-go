# ServicesBuildsStatusApi

All URIs are relative to *https://serverless.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchBuildStatus**](ServicesBuildsStatusApi.md#FetchBuildStatus) | **Get** /v1/Services/{ServiceSid}/Builds/{Sid}/Status | 



## FetchBuildStatus

> ServerlessV1BuildStatus FetchBuildStatus(ctx, ServiceSidSid)



Retrieve a specific Build resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Build resource from.
**Sid** | **string** | The SID of the Build resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchBuildStatusParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ServerlessV1BuildStatus**](ServerlessV1BuildStatus.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

