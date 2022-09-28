# ConnectionPoliciesApi

All URIs are relative to *https://voice.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnectionPolicy**](ConnectionPoliciesApi.md#CreateConnectionPolicy) | **Post** /v1/ConnectionPolicies | 
[**DeleteConnectionPolicy**](ConnectionPoliciesApi.md#DeleteConnectionPolicy) | **Delete** /v1/ConnectionPolicies/{Sid} | 
[**FetchConnectionPolicy**](ConnectionPoliciesApi.md#FetchConnectionPolicy) | **Get** /v1/ConnectionPolicies/{Sid} | 
[**ListConnectionPolicy**](ConnectionPoliciesApi.md#ListConnectionPolicy) | **Get** /v1/ConnectionPolicies | 
[**UpdateConnectionPolicy**](ConnectionPoliciesApi.md#UpdateConnectionPolicy) | **Post** /v1/ConnectionPolicies/{Sid} | 



## CreateConnectionPolicy

> VoiceV1ConnectionPolicy CreateConnectionPolicy(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateConnectionPolicyParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.

### Return type

[**VoiceV1ConnectionPolicy**](VoiceV1ConnectionPolicy.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnectionPolicy

> DeleteConnectionPolicy(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Connection Policy resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteConnectionPolicyParams struct


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


## FetchConnectionPolicy

> VoiceV1ConnectionPolicy FetchConnectionPolicy(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Connection Policy resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchConnectionPolicyParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VoiceV1ConnectionPolicy**](VoiceV1ConnectionPolicy.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectionPolicy

> []VoiceV1ConnectionPolicy ListConnectionPolicy(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListConnectionPolicyParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VoiceV1ConnectionPolicy**](VoiceV1ConnectionPolicy.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectionPolicy

> VoiceV1ConnectionPolicy UpdateConnectionPolicy(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Connection Policy resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConnectionPolicyParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.

### Return type

[**VoiceV1ConnectionPolicy**](VoiceV1ConnectionPolicy.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

