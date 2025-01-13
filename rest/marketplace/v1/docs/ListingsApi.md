# ListingsApi

All URIs are relative to *https://marketplace.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateModuleData**](ListingsApi.md#CreateModuleData) | **Post** /v1/Listings | 
[**FetchModuleDataForListingOwner**](ListingsApi.md#FetchModuleDataForListingOwner) | **Get** /v1/Listings | 



## CreateModuleData

> MarketplaceV1ModuleDataManagement CreateModuleData(ctx, optional)



This endpoint creates a Listing based on the given data.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateModuleDataParams struct


Name | Type | Description
------------- | ------------- | -------------
**ModuleInfo** | **string** | A JSON object containing essential attributes that define a Listing.
**Configuration** | **string** | A JSON object for providing Listing-specific configuration. Contains button setup, notification URL, and more.

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


## FetchModuleDataForListingOwner

> FetchModuleDataForListingOwnerResponse FetchModuleDataForListingOwner(ctx, )



This endpoint returns the Listings owned by the authenticated Account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchModuleDataForListingOwnerParams struct


### Return type

[**FetchModuleDataForListingOwnerResponse**](FetchModuleDataForListingOwner200Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

