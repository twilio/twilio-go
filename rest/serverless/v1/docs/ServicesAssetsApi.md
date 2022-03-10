# ServicesAssetsApi

All URIs are relative to *https://serverless.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAsset**](ServicesAssetsApi.md#CreateAsset) | **Post** /v1/Services/{ServiceSid}/Assets | 
[**DeleteAsset**](ServicesAssetsApi.md#DeleteAsset) | **Delete** /v1/Services/{ServiceSid}/Assets/{Sid} | 
[**FetchAsset**](ServicesAssetsApi.md#FetchAsset) | **Get** /v1/Services/{ServiceSid}/Assets/{Sid} | 
[**ListAsset**](ServicesAssetsApi.md#ListAsset) | **Get** /v1/Services/{ServiceSid}/Assets | 
[**UpdateAsset**](ServicesAssetsApi.md#UpdateAsset) | **Post** /v1/Services/{ServiceSid}/Assets/{Sid} | 



## CreateAsset

> ServerlessV1Asset CreateAsset(ctx, ServiceSidoptional)



Create a new Asset resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to create the Asset resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateAssetParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the Asset resource. It can be a maximum of 255 characters.

### Return type

[**ServerlessV1Asset**](ServerlessV1Asset.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAsset

> DeleteAsset(ctx, ServiceSidSid)



Delete an Asset resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to delete the Asset resource from.
**Sid** | **string** | The SID that identifies the Asset resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteAssetParams struct


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


## FetchAsset

> ServerlessV1Asset FetchAsset(ctx, ServiceSidSid)



Retrieve a specific Asset resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Asset resource from.
**Sid** | **string** | The SID that identifies the Asset resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchAssetParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ServerlessV1Asset**](ServerlessV1Asset.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAsset

> []ServerlessV1Asset ListAsset(ctx, ServiceSidoptional)



Retrieve a list of all Assets.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to read the Asset resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListAssetParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ServerlessV1Asset**](ServerlessV1Asset.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAsset

> ServerlessV1Asset UpdateAsset(ctx, ServiceSidSidoptional)



Update a specific Asset resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to update the Asset resource from.
**Sid** | **string** | The SID that identifies the Asset resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateAssetParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the Asset resource. It can be a maximum of 255 characters.

### Return type

[**ServerlessV1Asset**](ServerlessV1Asset.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

