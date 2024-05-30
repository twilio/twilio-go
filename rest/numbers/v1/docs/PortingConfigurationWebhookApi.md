# PortingConfigurationWebhookApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePortingWebhookConfiguration**](PortingConfigurationWebhookApi.md#CreatePortingWebhookConfiguration) | **Post** /v1/Porting/Configuration/Webhook | 
[**DeletePortingWebhookConfigurationDelete**](PortingConfigurationWebhookApi.md#DeletePortingWebhookConfigurationDelete) | **Delete** /v1/Porting/Configuration/Webhook/{WebhookType} | 
[**FetchPortingWebhookConfigurationFetch**](PortingConfigurationWebhookApi.md#FetchPortingWebhookConfigurationFetch) | **Get** /v1/Porting/Configuration/Webhook | 



## CreatePortingWebhookConfiguration

> NumbersV1PortingWebhookConfiguration CreatePortingWebhookConfiguration(ctx, optional)



Allows to create a new webhook configuration

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreatePortingWebhookConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**Body** | **map[string]interface{}** | 

### Return type

[**NumbersV1PortingWebhookConfiguration**](NumbersV1PortingWebhookConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePortingWebhookConfigurationDelete

> DeletePortingWebhookConfigurationDelete(ctx, WebhookType)



Allows the client to delete a webhook configuration.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WebhookType** | **string** | The of the webhook type of the configuration to be deleted

### Other Parameters

Other parameters are passed through a pointer to a DeletePortingWebhookConfigurationDeleteParams struct


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


## FetchPortingWebhookConfigurationFetch

> NumbersV1PortingWebhookConfigurationFetch FetchPortingWebhookConfigurationFetch(ctx, )



Allows to fetch the webhook configuration

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchPortingWebhookConfigurationFetchParams struct


### Return type

[**NumbersV1PortingWebhookConfigurationFetch**](NumbersV1PortingWebhookConfigurationFetch.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

