# InstalledAddOnsExtensionsApi

All URIs are relative to *https://marketplace.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchMarketplaceInstalledAddOnExtension**](InstalledAddOnsExtensionsApi.md#FetchMarketplaceInstalledAddOnExtension) | **Get** /v1/InstalledAddOns/{InstalledAddOnSid}/Extensions/{Sid} | 
[**ListMarketplaceInstalledAddOnExtension**](InstalledAddOnsExtensionsApi.md#ListMarketplaceInstalledAddOnExtension) | **Get** /v1/InstalledAddOns/{InstalledAddOnSid}/Extensions | 
[**UpdateMarketplaceInstalledAddOnExtension**](InstalledAddOnsExtensionsApi.md#UpdateMarketplaceInstalledAddOnExtension) | **Post** /v1/InstalledAddOns/{InstalledAddOnSid}/Extensions/{Sid} | 



## FetchMarketplaceInstalledAddOnExtension

> MarketplaceInstalledAddOnInstalledAddOnExtension FetchMarketplaceInstalledAddOnExtension(ctx, InstalledAddOnSidSid)



Fetch an instance of an Extension for the Installed Add-on.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InstalledAddOnSid** | **string** | The SID of the InstalledAddOn resource with the extension to fetch.
**Sid** | **string** | The SID of the InstalledAddOn Extension resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchMarketplaceInstalledAddOnExtensionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MarketplaceInstalledAddOnInstalledAddOnExtension**](MarketplaceInstalledAddOnInstalledAddOnExtension.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMarketplaceInstalledAddOnExtension

> []MarketplaceInstalledAddOnInstalledAddOnExtension ListMarketplaceInstalledAddOnExtension(ctx, InstalledAddOnSidoptional)



Retrieve a list of Extensions for the Installed Add-on.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InstalledAddOnSid** | **string** | The SID of the InstalledAddOn resource with the extensions to read.

### Other Parameters

Other parameters are passed through a pointer to a ListMarketplaceInstalledAddOnExtensionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]MarketplaceInstalledAddOnInstalledAddOnExtension**](MarketplaceInstalledAddOnInstalledAddOnExtension.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMarketplaceInstalledAddOnExtension

> MarketplaceInstalledAddOnInstalledAddOnExtension UpdateMarketplaceInstalledAddOnExtension(ctx, InstalledAddOnSidSidoptional)



Update an Extension for an Add-on installation.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InstalledAddOnSid** | **string** | The SID of the InstalledAddOn resource with the extension to update.
**Sid** | **string** | The SID of the InstalledAddOn Extension resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateMarketplaceInstalledAddOnExtensionParams struct


Name | Type | Description
------------- | ------------- | -------------
**Enabled** | **bool** | Whether the Extension should be invoked.

### Return type

[**MarketplaceInstalledAddOnInstalledAddOnExtension**](MarketplaceInstalledAddOnInstalledAddOnExtension.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

