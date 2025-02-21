# SigningRequestConfigurationApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSigningRequestConfiguration**](SigningRequestConfigurationApi.md#ListSigningRequestConfiguration) | **Get** /v1/SigningRequest/Configuration | 



## ListSigningRequestConfiguration

> []NumbersV1SigningRequestConfiguration ListSigningRequestConfiguration(ctx, optional)



Synchronous operation to retrieve configurations for the customer.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSigningRequestConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**Country** | **string** | The country ISO code to apply this configuration, this is an optional field, Example: US, MX
**Product** | **string** | The product or service for which is requesting the signature, this is an optional field, Example: Porting, Hosting
**PageSize** | **int64** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int64** | Max number of records to return.

### Return type

[**[]NumbersV1SigningRequestConfiguration**](NumbersV1SigningRequestConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

