# ControlPlaneOperatorsVersionsApi

All URIs are relative to *https://intelligence.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchOperatorVersion**](ControlPlaneOperatorsVersionsApi.md#FetchOperatorVersion) | **Get** /v3/ControlPlane/Operators/{id}/Versions/{version} | Retrieve a specific Operator Version
[**ListOperatorVersions**](ControlPlaneOperatorsVersionsApi.md#ListOperatorVersions) | **Get** /v3/ControlPlane/Operators/{id}/Versions | Retrieve a list of Operator Versions



## FetchOperatorVersion

> FetchOperatorVersionResponse FetchOperatorVersion(ctx, IdVersion)

Retrieve a specific Operator Version

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The unique identifier (TTID) of the Language Operator.
**Version** | **int** | The numeric version number of the Language Operator.

### Other Parameters

Other parameters are passed through a pointer to a FetchOperatorVersionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**FetchOperatorVersionResponse**](FetchOperatorVersion200Response.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperatorVersions

> []ListOperatorVersionsResponseItems ListOperatorVersions(ctx, Idoptional)

Retrieve a list of Operator Versions

Retrieve a paginated list of all versions for a given Language Operator.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The unique identifier (TTID) of the Language Operator.

### Other Parameters

Other parameters are passed through a pointer to a ListOperatorVersionsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | The maximum number of resources to return
**PageToken** | **string** | Token for pagination
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ListOperatorVersionsResponseItems**](ListOperatorVersionsResponseItems.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

