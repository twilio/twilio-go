# SigningRequestConfigurationApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSigningRequestConfiguration**](SigningRequestConfigurationApi.md#CreateSigningRequestConfiguration) | **Post** /v1/SigningRequest/Configuration | 



## CreateSigningRequestConfiguration

> NumbersV1SigningRequestConfiguration CreateSigningRequestConfiguration(ctx, optional)



Synchronous operation to insert or update a configuration for the customer.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSigningRequestConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**Body** | **map[string]interface{}** | 

### Return type

[**NumbersV1SigningRequestConfiguration**](NumbersV1SigningRequestConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

