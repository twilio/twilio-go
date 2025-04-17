# ListingApi

All URIs are relative to *https://marketplace.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchModuleDataManagement**](ListingApi.md#FetchModuleDataManagement) | **Get** /v1/Listing/{Sid} | 
[**UpdateModuleDataManagement**](ListingApi.md#UpdateModuleDataManagement) | **Post** /v1/Listing/{Sid} | 



## FetchModuleDataManagement

> MarketplaceV1ModuleDataManagement FetchModuleDataManagement(ctx, Sid)



This endpoint returns the data of a given Listing. To find a Listing's SID, use the [Available Add-ons resource](/docs/marketplace/api/available-add-ons) or view its Listing details page in the Console by visiting the [Catalog](https://console.twilio.com/us1/develop/add-ons/catalog) or the [My Listings tab](https://console.twilio.com/us1/develop/add-ons/publish/my-listings) and selecting the Listing.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique identifier of a Listing.

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



This endpoint updates the data of a given Listing. To find a Listing's SID, use the [Available Add-ons resource](https://www.twilio.com/docs/marketplace/api/available-add-ons) or view its Listing details page in the Console by visiting the [Catalog](https://console.twilio.com/us1/develop/add-ons/catalog) or the [My Listings tab](https://console.twilio.com/us1/develop/add-ons/publish/my-listings) and selecting the Listing. Only Listing owners are allowed to update the Listing.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | SID that uniquely identifies the Listing.

### Other Parameters

Other parameters are passed through a pointer to a UpdateModuleDataManagementParams struct


Name | Type | Description
------------- | ------------- | -------------
**ModuleInfo** | **string** | A JSON object containing essential attributes that define a Listing.
**Description** | **string** | A JSON object describing the Listing. You can define the main body of the description, highlight key features or aspects of the Listing, and provide code samples for developers if applicable.
**Documentation** | **string** | A JSON object for providing comprehensive information, instructions, and resources related to the Listing.
**Policies** | **string** | A JSON object describing the Listing's privacy and legal policies. The maximum file size for Policies is 5MB.
**Support** | **string** | A JSON object containing information on how Marketplace users can obtain support for the Listing. Use this parameter to provide details such as contact information and support description.
**Configuration** | **string** | A JSON object for providing Listing-specific configuration. Contains button setup, notification URL, and more.
**Pricing** | **string** | A JSON object for providing Listing's purchase options.

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

