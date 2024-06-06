# ListingApi

All URIs are relative to *https://marketplace.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchMarketplaceModuleDataManagement**](ListingApi.md#FetchMarketplaceModuleDataManagement) | **Get** /v1/Listing/{Sid} | 
[**UpdateMarketplaceModuleDataManagement**](ListingApi.md#UpdateMarketplaceModuleDataManagement) | **Post** /v1/Listing/{Sid} | 



## FetchMarketplaceModuleDataManagement

> MarketplaceModuleDataManagement FetchMarketplaceModuleDataManagement(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchMarketplaceModuleDataManagementParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MarketplaceModuleDataManagement**](MarketplaceModuleDataManagement.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMarketplaceModuleDataManagement

> MarketplaceModuleDataManagement UpdateMarketplaceModuleDataManagement(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateMarketplaceModuleDataManagementParams struct


Name | Type | Description
------------- | ------------- | -------------
**ModuleInfo** | **string** | 
**Description** | **string** | 
**Documentation** | **string** | 
**Policies** | **string** | 
**Support** | **string** | 

### Return type

[**MarketplaceModuleDataManagement**](MarketplaceModuleDataManagement.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

