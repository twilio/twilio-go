# RegulatoryComplianceEndUsersApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEndUser**](RegulatoryComplianceEndUsersApi.md#CreateEndUser) | **Post** /v2/RegulatoryCompliance/EndUsers | 
[**DeleteEndUser**](RegulatoryComplianceEndUsersApi.md#DeleteEndUser) | **Delete** /v2/RegulatoryCompliance/EndUsers/{Sid} | 
[**FetchEndUser**](RegulatoryComplianceEndUsersApi.md#FetchEndUser) | **Get** /v2/RegulatoryCompliance/EndUsers/{Sid} | 
[**ListEndUser**](RegulatoryComplianceEndUsersApi.md#ListEndUser) | **Get** /v2/RegulatoryCompliance/EndUsers | 
[**UpdateEndUser**](RegulatoryComplianceEndUsersApi.md#UpdateEndUser) | **Post** /v2/RegulatoryCompliance/EndUsers/{Sid} | 



## CreateEndUser

> NumbersV2RegulatoryComplianceEndUser CreateEndUser(ctx, optional)



Create a new End User.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateEndUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**Attributes** | [**map[string]interface{}**](map[string]interface{}.md) | The set of parameters that are the attributes of the End User resource which are derived End User Types.
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**Type** | **string** | The type of end user of the Bundle resource - can be &#x60;individual&#x60; or &#x60;business&#x60;.

### Return type

[**NumbersV2RegulatoryComplianceEndUser**](NumbersV2RegulatoryComplianceEndUser.md)

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

> NumbersV2RegulatoryComplianceEndUser FetchEndUser(ctx, Sid)



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

[**NumbersV2RegulatoryComplianceEndUser**](NumbersV2RegulatoryComplianceEndUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEndUser

> ListEndUserResponse ListEndUser(ctx, optional)



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

[**ListEndUserResponse**](ListEndUserResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndUser

> NumbersV2RegulatoryComplianceEndUser UpdateEndUser(ctx, Sidoptional)



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
**Attributes** | [**map[string]interface{}**](map[string]interface{}.md) | The set of parameters that are the attributes of the End User resource which are derived End User Types.
**FriendlyName** | **string** | The string that you assigned to describe the resource.

### Return type

[**NumbersV2RegulatoryComplianceEndUser**](NumbersV2RegulatoryComplianceEndUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

