# ControlPlaneConfigurationsApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfiguration**](ControlPlaneConfigurationsApi.md#CreateConfiguration) | **Post** /v2/ControlPlane/Configurations | Create a Configuration
[**DeleteConfiguration**](ControlPlaneConfigurationsApi.md#DeleteConfiguration) | **Delete** /v2/ControlPlane/Configurations/{Sid} | Delete Configuration
[**FetchConfiguration**](ControlPlaneConfigurationsApi.md#FetchConfiguration) | **Get** /v2/ControlPlane/Configurations/{Sid} | Fetch Configuration
[**ListConfiguration**](ControlPlaneConfigurationsApi.md#ListConfiguration) | **Get** /v2/ControlPlane/Configurations | List Configurations
[**UpdateConfiguration**](ControlPlaneConfigurationsApi.md#UpdateConfiguration) | **Put** /v2/ControlPlane/Configurations/{Sid} | Update Configuration



## CreateConfiguration

> CreateConfigurationResponse CreateConfiguration(ctx, optional)

Create a Configuration

Create a new Configuration

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**IdempotencyKey** | **string** | Client-generated UUID key to ensure idempotent behavior. Submitting the same key returns the original response without creating a duplicate operation. Keys are scoped to account + region with a 24-hour TTL.
**CreateConfigurationRequest** | [**CreateConfigurationRequest**](CreateConfigurationRequest.md) | The configuration to create

### Return type

[**CreateConfigurationResponse**](CreateConfiguration202Response.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfiguration

> CreateConfigurationResponse DeleteConfiguration(ctx, Sidoptional)

Delete Configuration

Delete a Configuration

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**IdempotencyKey** | **string** | Client-generated UUID key to ensure idempotent behavior. Submitting the same key returns the original response without creating a duplicate operation. Keys are scoped to account + region with a 24-hour TTL.

### Return type

[**CreateConfigurationResponse**](CreateConfiguration202Response.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConfiguration

> ListConfigurationResponseConfigurations FetchConfiguration(ctx, Sid)

Fetch Configuration

Retrieve a Configuration.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ListConfigurationResponseConfigurations**](ListConfiguration200ResponseConfigurations.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConfiguration

> []ListConfigurationResponseConfigurations ListConfiguration(ctx, optional)

List Configurations

Retrieve a list of Configurations.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | Maximum number of items to return in a single response
**PageToken** | **string** | A URL-safe, base64-encoded token representing the page of results to return
**MemoryStoreId** | **string** | Filter configurations by Memory Store ID
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ListConfigurationResponseConfigurations**](ListConfigurationResponseConfigurations.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration

> CreateConfigurationResponse UpdateConfiguration(ctx, Sidoptional)

Update Configuration

Update an existing Configuration

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**IdempotencyKey** | **string** | Client-generated UUID key to ensure idempotent behavior. Submitting the same key returns the original response without creating a duplicate operation. Keys are scoped to account + region with a 24-hour TTL.
**UpdateConfigurationRequest** | [**UpdateConfigurationRequest**](UpdateConfigurationRequest.md) | The configuration to update

### Return type

[**CreateConfigurationResponse**](CreateConfiguration202Response.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

