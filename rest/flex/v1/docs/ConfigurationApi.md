# ConfigurationApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchConfiguration**](ConfigurationApi.md#FetchConfiguration) | **Get** /v1/Configuration | 



## FetchConfiguration

> FlexV1Configuration FetchConfiguration(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**UiVersion** | **string** | The Pinned UI version of the Configuration resource to fetch.

### Return type

[**FlexV1Configuration**](FlexV1Configuration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

