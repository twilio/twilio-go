# A2pBrandRegistrationsApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBrandRegistrations**](A2pBrandRegistrationsApi.md#CreateBrandRegistrations) | **Post** /v1/a2p/BrandRegistrations | 
[**FetchBrandRegistrations**](A2pBrandRegistrationsApi.md#FetchBrandRegistrations) | **Get** /v1/a2p/BrandRegistrations/{Sid} | 
[**ListBrandRegistrations**](A2pBrandRegistrationsApi.md#ListBrandRegistrations) | **Get** /v1/a2p/BrandRegistrations | 
[**UpdateBrandRegistrations**](A2pBrandRegistrationsApi.md#UpdateBrandRegistrations) | **Post** /v1/a2p/BrandRegistrations/{Sid} | 



## CreateBrandRegistrations

> MessagingV1BrandRegistrations CreateBrandRegistrations(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateBrandRegistrationsParams struct


Name | Type | Description
------------- | ------------- | -------------
**A2PProfileBundleSid** | **string** | A2P Messaging Profile Bundle Sid.
**BrandType** | **string** | Type of brand being created. One of: \\\&quot;STANDARD\\\&quot;, \\\&quot;STARTER\\\&quot;. STARTER is for low volume, starter use cases. STANDARD is for all other use cases.
**CustomerProfileBundleSid** | **string** | Customer Profile Bundle Sid.
**Mock** | **bool** | A boolean that specifies whether brand should be a mock or not. If true, brand will be registered as a mock brand. Defaults to false if no value is provided.
**SkipAutomaticSecVet** | **bool** | A flag to disable automatic secondary vetting for brands which it would otherwise be done.

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


## UpdateBrandRegistrations

> MessagingV1BrandRegistrations UpdateBrandRegistrations(ctx, Sid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Brand Registration resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateBrandRegistrationsParams struct


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

