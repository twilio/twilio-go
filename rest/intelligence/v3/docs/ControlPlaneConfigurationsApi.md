# ControlPlaneConfigurationsApi

All URIs are relative to *https://intelligence.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfiguration**](ControlPlaneConfigurationsApi.md#CreateConfiguration) | **Post** /v3/ControlPlane/Configurations | Create an Intelligence Configuration
[**DeleteConfiguration**](ControlPlaneConfigurationsApi.md#DeleteConfiguration) | **Delete** /v3/ControlPlane/Configurations/{id} | Delete an Intelligence Configuration
[**FetchConfiguration**](ControlPlaneConfigurationsApi.md#FetchConfiguration) | **Get** /v3/ControlPlane/Configurations/{id} | Retrieve an Intelligence Configuration
[**ListConfigurations**](ControlPlaneConfigurationsApi.md#ListConfigurations) | **Get** /v3/ControlPlane/Configurations | Retrieve a list of Intelligence Configurations
[**UpdateConfiguration**](ControlPlaneConfigurationsApi.md#UpdateConfiguration) | **Put** /v3/ControlPlane/Configurations/{id} | Update an Intelligence Configuration



## CreateConfiguration

> CreateConfigurationResponse CreateConfiguration(ctx, optional)

Create an Intelligence Configuration

Create an Intelligence Configuration to control how and when language operators analyze conversations.  You must include the `rules` field in the request. You can pass an empty array (`\"rules\": []`), but an Intelligence Configuration with an empty `rules` array doesn't execute any language operators. To add rules later, make a `PUT /v3/ControlPlane/Configurations/{id}` request.  After creating an Intelligence Configuration, you must attach it to a Conversations Configuration by using the Conversations API. 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**CreateConfigurationRequest** | [**CreateConfigurationRequest**](CreateConfigurationRequest.md) | 

### Return type

[**CreateConfigurationResponse**](CreateConfiguration201Response.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfiguration

> DeleteConfiguration(ctx, Id)

Delete an Intelligence Configuration

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The unique identifier of the Intelligence Configuration.

### Other Parameters

Other parameters are passed through a pointer to a DeleteConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConfiguration

> CreateConfigurationResponse FetchConfiguration(ctx, Id)

Retrieve an Intelligence Configuration

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The unique identifier of the Intelligence Configuration.

### Other Parameters

Other parameters are passed through a pointer to a FetchConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**CreateConfigurationResponse**](CreateConfiguration201Response.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConfigurations

> []IntelligenceConfiguration ListConfigurations(ctx, optional)

Retrieve a list of Intelligence Configurations

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListConfigurationsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | The maximum number of resources to return
**PageToken** | **string** | Token for pagination
**Limit** | **int** | Max number of records to return.

### Return type

[**[]IntelligenceConfiguration**](IntelligenceConfiguration.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration

> CreateConfigurationResponse UpdateConfiguration(ctx, Idoptional)

Update an Intelligence Configuration

Update an Intelligence Configuration, including its metadata and rules. The request must include all required fields.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The unique identifier of the Intelligence Configuration.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**UpdateConfigurationRequest** | [**UpdateConfigurationRequest**](UpdateConfigurationRequest.md) | 

### Return type

[**CreateConfigurationResponse**](CreateConfiguration201Response.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

