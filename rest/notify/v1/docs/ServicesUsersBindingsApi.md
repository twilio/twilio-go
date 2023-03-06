# ServicesUsersBindingsApi

All URIs are relative to *https://notify.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserBinding**](ServicesUsersBindingsApi.md#CreateUserBinding) | **Post** /v1/Services/{ServiceSid}/Users/{Identity}/Bindings | 
[**DeleteUserBinding**](ServicesUsersBindingsApi.md#DeleteUserBinding) | **Delete** /v1/Services/{ServiceSid}/Users/{Identity}/Bindings/{Sid} | 
[**FetchUserBinding**](ServicesUsersBindingsApi.md#FetchUserBinding) | **Get** /v1/Services/{ServiceSid}/Users/{Identity}/Bindings/{Sid} | 
[**ListUserBinding**](ServicesUsersBindingsApi.md#ListUserBinding) | **Get** /v1/Services/{ServiceSid}/Users/{Identity}/Bindings | 



## CreateUserBinding

> NotifyV1UserBinding CreateUserBinding(ctx, ServiceSidIdentityoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Identity** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateUserBindingParams struct


Name | Type | Description
------------- | ------------- | -------------
**BindingType** | **string** | 
**Address** | **string** | 
**Tag** | **[]string** | 
**NotificationProtocolVersion** | **string** | 
**CredentialSid** | **string** | 
**Endpoint** | **string** | 

### Return type

[**NotifyV1UserBinding**](NotifyV1UserBinding.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserBinding

> DeleteUserBinding(ctx, ServiceSidIdentitySid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Identity** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteUserBindingParams struct


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


## FetchUserBinding

> NotifyV1UserBinding FetchUserBinding(ctx, ServiceSidIdentitySid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Identity** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchUserBindingParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NotifyV1UserBinding**](NotifyV1UserBinding.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserBinding

> []NotifyV1UserBinding ListUserBinding(ctx, ServiceSidIdentityoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Identity** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListUserBindingParams struct


Name | Type | Description
------------- | ------------- | -------------
**StartDate** | **string** | 
**EndDate** | **string** | 
**Tag** | **[]string** | 
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]NotifyV1UserBinding**](NotifyV1UserBinding.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

