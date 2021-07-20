# A2pBrandRegistrationsApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBrandRegistrations**](A2pBrandRegistrationsApi.md#CreateBrandRegistrations) | **Post** /v1/a2p/BrandRegistrations | 
[**FetchBrandRegistrations**](A2pBrandRegistrationsApi.md#FetchBrandRegistrations) | **Get** /v1/a2p/BrandRegistrations/{Sid} | 
[**ListBrandRegistrations**](A2pBrandRegistrationsApi.md#ListBrandRegistrations) | **Get** /v1/a2p/BrandRegistrations | 



## CreateBrandRegistrations

> MessagingV1BrandRegistrations CreateBrandRegistrations(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateBrandRegistrationsParams struct


Name | Type | Description
------------- | ------------- | -------------
**A2pProfileBundleSid** | **string** | A2P Messaging Profile Bundle Sid.
**CustomerProfileBundleSid** | **string** | Customer Profile Bundle Sid.

### Return type

[**MessagingV1BrandRegistrations**](MessagingV1BrandRegistrations.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchBrandRegistrations

> MessagingV1BrandRegistrations FetchBrandRegistrations(ctx, Sid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Brand Registration resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchBrandRegistrationsParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MessagingV1BrandRegistrations**](MessagingV1BrandRegistrations.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBrandRegistrations

> ListBrandRegistrationsResponse ListBrandRegistrations(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListBrandRegistrationsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**ListBrandRegistrationsResponse**](ListBrandRegistrationsResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

