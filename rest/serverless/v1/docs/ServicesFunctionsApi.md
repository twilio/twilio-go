# ServicesFunctionsApi

All URIs are relative to *https://serverless.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFunction**](ServicesFunctionsApi.md#CreateFunction) | **Post** /v1/Services/{ServiceSid}/Functions | 
[**DeleteFunction**](ServicesFunctionsApi.md#DeleteFunction) | **Delete** /v1/Services/{ServiceSid}/Functions/{Sid} | 
[**FetchFunction**](ServicesFunctionsApi.md#FetchFunction) | **Get** /v1/Services/{ServiceSid}/Functions/{Sid} | 
[**ListFunction**](ServicesFunctionsApi.md#ListFunction) | **Get** /v1/Services/{ServiceSid}/Functions | 
[**UpdateFunction**](ServicesFunctionsApi.md#UpdateFunction) | **Post** /v1/Services/{ServiceSid}/Functions/{Sid} | 



## CreateFunction

> ServerlessV1Function CreateFunction(ctx, ServiceSidoptional)



Create a new Function resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to create the Function resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateFunctionParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the Function resource. It can be a maximum of 255 characters.

### Return type

[**ServerlessV1Function**](ServerlessV1Function.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFunction

> DeleteFunction(ctx, ServiceSidSid)



Delete a Function resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to delete the Function resource from.
**Sid** | **string** | The SID of the Function resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteFunctionParams struct


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


## FetchFunction

> ServerlessV1Function FetchFunction(ctx, ServiceSidSid)



Retrieve a specific Function resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Function resource from.
**Sid** | **string** | The SID of the Function resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFunctionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ServerlessV1Function**](ServerlessV1Function.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFunction

> []ServerlessV1Function ListFunction(ctx, ServiceSidoptional)



Retrieve a list of all Functions.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to read the Function resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListFunctionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ServerlessV1Function**](ServerlessV1Function.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFunction

> ServerlessV1Function UpdateFunction(ctx, ServiceSidSidoptional)



Update a specific Function resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to update the Function resource from.
**Sid** | **string** | The SID of the Function resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateFunctionParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the Function resource. It can be a maximum of 255 characters.

### Return type

[**ServerlessV1Function**](ServerlessV1Function.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

