# CustomerProfilesApi

All URIs are relative to *https://trusthub.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomerProfile**](CustomerProfilesApi.md#CreateCustomerProfile) | **Post** /v1/CustomerProfiles | 
[**DeleteCustomerProfile**](CustomerProfilesApi.md#DeleteCustomerProfile) | **Delete** /v1/CustomerProfiles/{Sid} | 
[**FetchCustomerProfile**](CustomerProfilesApi.md#FetchCustomerProfile) | **Get** /v1/CustomerProfiles/{Sid} | 
[**ListCustomerProfile**](CustomerProfilesApi.md#ListCustomerProfile) | **Get** /v1/CustomerProfiles | 
[**UpdateCustomerProfile**](CustomerProfilesApi.md#UpdateCustomerProfile) | **Post** /v1/CustomerProfiles/{Sid} | 



## CreateCustomerProfile

> TrusthubV1CustomerProfile CreateCustomerProfile(ctx, optional)



Create a new Customer-Profile.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateCustomerProfileParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**Email** | **string** | The email address that will receive updates when the Customer-Profile resource changes status.
**PolicySid** | **string** | The unique string of a policy that is associated to the Customer-Profile resource.
**StatusCallback** | **string** | The URL we call to inform your application of status changes.

### Return type

[**TrusthubV1CustomerProfile**](TrusthubV1CustomerProfile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomerProfile

> DeleteCustomerProfile(ctx, Sid)



Delete a specific Customer-Profile.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Customer-Profile resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteCustomerProfileParams struct


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


## FetchCustomerProfile

> TrusthubV1CustomerProfile FetchCustomerProfile(ctx, Sid)



Fetch a specific Customer-Profile instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Customer-Profile resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchCustomerProfileParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrusthubV1CustomerProfile**](TrusthubV1CustomerProfile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomerProfile

> []TrusthubV1CustomerProfile ListCustomerProfile(ctx, optional)



Retrieve a list of all Customer-Profiles for an account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListCustomerProfileParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | **string** | The verification status of the Customer-Profile resource.
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**PolicySid** | **string** | The unique string of a policy that is associated to the Customer-Profile resource.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TrusthubV1CustomerProfile**](TrusthubV1CustomerProfile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomerProfile

> TrusthubV1CustomerProfile UpdateCustomerProfile(ctx, Sidoptional)



Updates a Customer-Profile in an account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Customer-Profile resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateCustomerProfileParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | **string** | 
**StatusCallback** | **string** | The URL we call to inform your application of status changes.
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**Email** | **string** | The email address that will receive updates when the Customer-Profile resource changes status.

### Return type

[**TrusthubV1CustomerProfile**](TrusthubV1CustomerProfile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

