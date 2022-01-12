# RegulatoryComplianceBundlesCopiesApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBundleCopy**](RegulatoryComplianceBundlesCopiesApi.md#CreateBundleCopy) | **Post** /v2/RegulatoryCompliance/Bundles/{BundleSid}/Copies | 
[**ListBundleCopy**](RegulatoryComplianceBundlesCopiesApi.md#ListBundleCopy) | **Get** /v2/RegulatoryCompliance/Bundles/{BundleSid}/Copies | 



## CreateBundleCopy

> NumbersV2BundleCopy CreateBundleCopy(ctx, BundleSidoptional)



Creates a new copy of a Bundle. It will internally create copies of all the bundle items (identities and documents) of the original bundle

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string** | The unique string that identifies the Bundle to be copied.

### Other Parameters

Other parameters are passed through a pointer to a CreateBundleCopyParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | The string that you assigned to describe the copied bundle.

### Return type

[**NumbersV2BundleCopy**](NumbersV2BundleCopy.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBundleCopy

> ListBundleCopyResponse ListBundleCopy(ctx, BundleSidoptional)



Retrieve a list of all Bundles Copies for a Bundle.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string** | The unique string that we created to identify the Bundle resource.

### Other Parameters

Other parameters are passed through a pointer to a ListBundleCopyParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**ListBundleCopyResponse**](ListBundleCopyResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

