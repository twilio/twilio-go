# ServicesAssetsVersionsApi

All URIs are relative to *https://serverless.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAssetVersion**](ServicesAssetsVersionsApi.md#FetchAssetVersion) | **Get** /v1/Services/{ServiceSid}/Assets/{AssetSid}/Versions/{Sid} | 
[**ListAssetVersion**](ServicesAssetsVersionsApi.md#ListAssetVersion) | **Get** /v1/Services/{ServiceSid}/Assets/{AssetSid}/Versions | 



## FetchAssetVersion

> ServerlessV1AssetVersion FetchAssetVersion(ctx, ServiceSidAssetSidSid)



Retrieve a specific Asset Version.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Asset Version resource from.
**AssetSid** | **string** | The SID of the Asset resource that is the parent of the Asset Version resource to fetch.
**Sid** | **string** | The SID of the Asset Version resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchAssetVersionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ServerlessV1AssetVersion**](ServerlessV1AssetVersion.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssetVersion

> []ServerlessV1AssetVersion ListAssetVersion(ctx, ServiceSidAssetSidoptional)



Retrieve a list of all Asset Versions.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to read the Asset Version resource from.
**AssetSid** | **string** | The SID of the Asset resource that is the parent of the Asset Version resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListAssetVersionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ServerlessV1AssetVersion**](ServerlessV1AssetVersion.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

