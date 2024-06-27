# AvailableAddOnsExtensionsApi

All URIs are relative to *https://marketplace.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAvailableAddOnExtension**](AvailableAddOnsExtensionsApi.md#FetchAvailableAddOnExtension) | **Get** /v1/AvailableAddOns/{AvailableAddOnSid}/Extensions/{Sid} | 
[**ListAvailableAddOnExtension**](AvailableAddOnsExtensionsApi.md#ListAvailableAddOnExtension) | **Get** /v1/AvailableAddOns/{AvailableAddOnSid}/Extensions | 



## FetchAvailableAddOnExtension

> MarketplaceV1AvailableAddOnExtension FetchAvailableAddOnExtension(ctx, AvailableAddOnSidSid)



Fetch an instance of an Extension for the Available Add-on.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AvailableAddOnSid** | **string** | The SID of the AvailableAddOn resource with the extension to fetch.
**Sid** | **string** | The SID of the AvailableAddOn Extension resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchAvailableAddOnExtensionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MarketplaceV1AvailableAddOnExtension**](MarketplaceV1AvailableAddOnExtension.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailableAddOnExtension

> []MarketplaceV1AvailableAddOnExtension ListAvailableAddOnExtension(ctx, AvailableAddOnSidoptional)



Retrieve a list of Extensions for the Available Add-on.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AvailableAddOnSid** | **string** | The SID of the AvailableAddOn resource with the extensions to read.

### Other Parameters

Other parameters are passed through a pointer to a ListAvailableAddOnExtensionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]MarketplaceV1AvailableAddOnExtension**](MarketplaceV1AvailableAddOnExtension.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

