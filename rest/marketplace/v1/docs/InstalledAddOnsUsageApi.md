# InstalledAddOnsUsageApi

All URIs are relative to *https://marketplace.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBillingUsage**](InstalledAddOnsUsageApi.md#CreateBillingUsage) | **Post** /v1/InstalledAddOns/{InstalledAddOnSid}/Usage | 



## CreateBillingUsage

> MarketplaceV1BillingUsageResponse CreateBillingUsage(ctx, InstalledAddOnSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InstalledAddOnSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateBillingUsageParams struct


Name | Type | Description
------------- | ------------- | -------------
**CreateBillingUsageRequest** | [**CreateBillingUsageRequest**](CreateBillingUsageRequest.md) | 

### Return type

[**MarketplaceV1BillingUsageResponse**](MarketplaceV1BillingUsageResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

