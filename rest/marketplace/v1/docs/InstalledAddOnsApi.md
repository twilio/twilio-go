# InstalledAddOnsApi

All URIs are relative to *https://marketplace.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInstalledAddOn**](InstalledAddOnsApi.md#CreateInstalledAddOn) | **Post** /v1/InstalledAddOns | 
[**DeleteInstalledAddOn**](InstalledAddOnsApi.md#DeleteInstalledAddOn) | **Delete** /v1/InstalledAddOns/{Sid} | 
[**FetchInstalledAddOn**](InstalledAddOnsApi.md#FetchInstalledAddOn) | **Get** /v1/InstalledAddOns/{Sid} | 
[**ListInstalledAddOn**](InstalledAddOnsApi.md#ListInstalledAddOn) | **Get** /v1/InstalledAddOns | 
[**UpdateInstalledAddOn**](InstalledAddOnsApi.md#UpdateInstalledAddOn) | **Post** /v1/InstalledAddOns/{Sid} | 



## CreateInstalledAddOn

> MarketplaceV1InstalledAddOn CreateInstalledAddOn(ctx, optional)



Install an Add-on for the Account specified.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateInstalledAddOnParams struct


Name | Type | Description
------------- | ------------- | -------------
**AvailableAddOnSid** | **string** | The SID of the AvaliableAddOn to install.
**AcceptTermsOfService** | **bool** | Whether the Terms of Service were accepted.
**Configuration** | [**interface{}**](interface{}.md) | The JSON object that represents the configuration of the new Add-on being installed.
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. This value must be unique within the Account.

### Return type

[**MarketplaceV1InstalledAddOn**](MarketplaceV1InstalledAddOn.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstalledAddOn

> DeleteInstalledAddOn(ctx, Sid)



Remove an Add-on installation from your account

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the InstalledAddOn resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteInstalledAddOnParams struct


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


## FetchInstalledAddOn

> MarketplaceV1InstalledAddOn FetchInstalledAddOn(ctx, Sid)



Fetch an instance of an Add-on currently installed on this Account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the InstalledAddOn resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchInstalledAddOnParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MarketplaceV1InstalledAddOn**](MarketplaceV1InstalledAddOn.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstalledAddOn

> []MarketplaceV1InstalledAddOn ListInstalledAddOn(ctx, optional)



Retrieve a list of Add-ons currently installed on this Account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListInstalledAddOnParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]MarketplaceV1InstalledAddOn**](MarketplaceV1InstalledAddOn.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInstalledAddOn

> MarketplaceV1InstalledAddOn UpdateInstalledAddOn(ctx, Sidoptional)



Update an Add-on installation for the Account specified.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the InstalledAddOn resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateInstalledAddOnParams struct


Name | Type | Description
------------- | ------------- | -------------
**Configuration** | [**interface{}**](interface{}.md) | Valid JSON object that conform to the configuration schema exposed by the associated AvailableAddOn resource. This is only required by Add-ons that need to be configured
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. This value must be unique within the Account.

### Return type

[**MarketplaceV1InstalledAddOn**](MarketplaceV1InstalledAddOn.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

