# InstalledAddOnsUsageApi

All URIs are relative to *https://marketplace.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMarketplaceBillingUsage**](InstalledAddOnsUsageApi.md#CreateMarketplaceBillingUsage) | **Post** /v1/InstalledAddOns/{InstalledAddOnSid}/Usage | 



## CreateMarketplaceBillingUsage

> MarketplaceInstalledAddOnBillingUsageResponse CreateMarketplaceBillingUsage(ctx, InstalledAddOnSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InstalledAddOnSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateMarketplaceBillingUsageParams struct


Name | Type | Description
------------- | ------------- | -------------
**CreateMarketplaceBillingUsageRequest** | [**CreateMarketplaceBillingUsageRequest**](CreateMarketplaceBillingUsageRequest.md) | 

### Return type

[**MarketplaceInstalledAddOnBillingUsageResponse**](MarketplaceInstalledAddOnBillingUsageResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

