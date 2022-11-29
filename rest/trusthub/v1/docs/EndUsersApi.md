# EndUsersApi

All URIs are relative to *https://trusthub.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEndUser**](EndUsersApi.md#CreateEndUser) | **Post** /v1/EndUsers | 
[**DeleteEndUser**](EndUsersApi.md#DeleteEndUser) | **Delete** /v1/EndUsers/{Sid} | 
[**FetchEndUser**](EndUsersApi.md#FetchEndUser) | **Get** /v1/EndUsers/{Sid} | 
[**ListEndUser**](EndUsersApi.md#ListEndUser) | **Get** /v1/EndUsers | 
[**UpdateEndUser**](EndUsersApi.md#UpdateEndUser) | **Post** /v1/EndUsers/{Sid} | 



## CreateEndUser

> TrusthubV1EndUser CreateEndUser(ctx, optional)



Create a new End User.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateEndUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**Type** | **string** | The type of end user of the Bundle resource - can be `individual` or `business`.
**Attributes** | [**interface{}**](interface{}.md) | The set of parameters that are the attributes of the End User resource which are derived End User Types.

### Return type

[**TrusthubV1EndUser**](TrusthubV1EndUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEndUser

> DeleteEndUser(ctx, Sid)



Delete a specific End User.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string created by Twilio to identify the End User resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteEndUserParams struct


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


## FetchEndUser

> TrusthubV1EndUser FetchEndUser(ctx, Sid)



Fetch specific End User Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string created by Twilio to identify the End User resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchEndUserParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrusthubV1EndUser**](TrusthubV1EndUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEndUser

> []TrusthubV1EndUser ListEndUser(ctx, optional)



Retrieve a list of all End User for an account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListEndUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TrusthubV1EndUser**](TrusthubV1EndUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndUser

> TrusthubV1EndUser UpdateEndUser(ctx, Sidoptional)



Update an existing End User.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string created by Twilio to identify the End User resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateEndUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**Attributes** | [**interface{}**](interface{}.md) | The set of parameters that are the attributes of the End User resource which are derived End User Types.

### Return type

[**TrusthubV1EndUser**](TrusthubV1EndUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

