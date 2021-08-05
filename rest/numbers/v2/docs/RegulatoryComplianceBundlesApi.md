# RegulatoryComplianceBundlesApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBundle**](RegulatoryComplianceBundlesApi.md#CreateBundle) | **Post** /v2/RegulatoryCompliance/Bundles | 
[**DeleteBundle**](RegulatoryComplianceBundlesApi.md#DeleteBundle) | **Delete** /v2/RegulatoryCompliance/Bundles/{Sid} | 
[**FetchBundle**](RegulatoryComplianceBundlesApi.md#FetchBundle) | **Get** /v2/RegulatoryCompliance/Bundles/{Sid} | 
[**ListBundle**](RegulatoryComplianceBundlesApi.md#ListBundle) | **Get** /v2/RegulatoryCompliance/Bundles | 
[**UpdateBundle**](RegulatoryComplianceBundlesApi.md#UpdateBundle) | **Post** /v2/RegulatoryCompliance/Bundles/{Sid} | 



## CreateBundle

> NumbersV2Bundle CreateBundle(ctx, optional)



Create a new Bundle.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateBundleParams struct


Name | Type | Description
------------- | ------------- | -------------
**Email** | **string** | The email address that will receive updates when the Bundle resource changes status.
**EndUserType** | **string** | The type of End User of the Bundle resource.
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**IsoCountry** | **string** | The ISO country code of the Bundle&#39;s phone number country ownership request.
**NumberType** | **string** | The type of phone number of the Bundle&#39;s ownership request.
**RegulationSid** | **string** | The unique string of a regulation that is associated to the Bundle resource.
**StatusCallback** | **string** | The URL we call to inform your application of status changes.

### Return type

[**NumbersV2Bundle**](NumbersV2Bundle.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBundle

> DeleteBundle(ctx, Sid)



Delete a specific Bundle.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Bundle resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteBundleParams struct


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


## FetchBundle

> NumbersV2Bundle FetchBundle(ctx, Sid)



Fetch a specific Bundle instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Bundle resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchBundleParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NumbersV2Bundle**](NumbersV2Bundle.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBundle

> ListBundleResponse ListBundle(ctx, optional)



Retrieve a list of all Bundles for an account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListBundleParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | **string** | The verification status of the Bundle resource.
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**RegulationSid** | **string** | The unique string of a regulation that is associated to the Bundle resource.
**IsoCountry** | **string** | The ISO country code of the Bundle&#39;s phone number country ownership request.
**NumberType** | **string** | The type of phone number of the Bundle&#39;s ownership request.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**ListBundleResponse**](ListBundleResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBundle

> NumbersV2Bundle UpdateBundle(ctx, Sidoptional)



Updates a Bundle in an account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Bundle resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateBundleParams struct


Name | Type | Description
------------- | ------------- | -------------
**Email** | **string** | The email address that will receive updates when the Bundle resource changes status.
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**Status** | **string** | The verification status of the Bundle resource.
**StatusCallback** | **string** | The URL we call to inform your application of status changes.

### Return type

[**NumbersV2Bundle**](NumbersV2Bundle.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

