# ListingApi

All URIs are relative to *https://marketplace.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchModuleDataManagement**](ListingApi.md#FetchModuleDataManagement) | **Get** /v1/Listing/{Sid} | 
[**UpdateModuleDataManagement**](ListingApi.md#UpdateModuleDataManagement) | **Post** /v1/Listing/{Sid} | 



## FetchModuleDataManagement

> MarketplaceV1ModuleDataManagement FetchModuleDataManagement(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchModuleDataManagementParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MarketplaceV1ModuleDataManagement**](MarketplaceV1ModuleDataManagement.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateModuleDataManagement

> MarketplaceV1ModuleDataManagement UpdateModuleDataManagement(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateModuleDataManagementParams struct


Name | Type | Description
------------- | ------------- | -------------
**ModuleInfo** | **string** | 
**Description** | **string** | 
**Documentation** | **string** | 
**Policies** | **string** | 
**Support** | **string** | 

### Return type

[**MarketplaceV1ModuleDataManagement**](MarketplaceV1ModuleDataManagement.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

