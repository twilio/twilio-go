# ControlPlaneOperatorsApi

All URIs are relative to *https://intelligence.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOperator**](ControlPlaneOperatorsApi.md#CreateOperator) | **Post** /v3/ControlPlane/Operators | Create an Operator
[**DeleteOperator**](ControlPlaneOperatorsApi.md#DeleteOperator) | **Delete** /v3/ControlPlane/Operators/{id} | Delete an Operator
[**FetchOperator**](ControlPlaneOperatorsApi.md#FetchOperator) | **Get** /v3/ControlPlane/Operators/{id} | Retrieve an Operator
[**ListOperators**](ControlPlaneOperatorsApi.md#ListOperators) | **Get** /v3/ControlPlane/Operators | Retrieve a list of Operators
[**UpdateOperator**](ControlPlaneOperatorsApi.md#UpdateOperator) | **Put** /v3/ControlPlane/Operators/{id} | Update an Operator



## CreateOperator

> CreateOperatorResponse CreateOperator(ctx, optional)

Create an Operator

Create a custom language operator. You can define a reusable, programmable conversational analysis task tailored to your business needs.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateOperatorParams struct


Name | Type | Description
------------- | ------------- | -------------
**LanguageOperator** | [**LanguageOperator**](LanguageOperator.md) | Create Language Operator request

### Return type

[**CreateOperatorResponse**](CreateOperator201Response.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOperator

> DeleteOperator(ctx, Id)

Delete an Operator

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The unique identifier (TTID) of the Language Operator.

### Other Parameters

Other parameters are passed through a pointer to a DeleteOperatorParams struct


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


## FetchOperator

> CreateOperatorResponse FetchOperator(ctx, Id)

Retrieve an Operator

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The unique identifier (TTID) of the Language Operator.

### Other Parameters

Other parameters are passed through a pointer to a FetchOperatorParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**CreateOperatorResponse**](CreateOperator201Response.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperators

> []LanguageOperator ListOperators(ctx, optional)

Retrieve a list of Operators

Retrieve a list of Operators for the account, including those not attached to an Intelligence Configuration. This request returns both Twilio-authored and custom language operators.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListOperatorsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | The maximum number of resources to return
**PageToken** | **string** | Token for pagination
**Limit** | **int** | Max number of records to return.

### Return type

[**[]LanguageOperator**](LanguageOperator.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOperator

> CreateOperatorResponse UpdateOperator(ctx, Idoptional)

Update an Operator

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The unique identifier (TTID) of the Language Operator.

### Other Parameters

Other parameters are passed through a pointer to a UpdateOperatorParams struct


Name | Type | Description
------------- | ------------- | -------------
**IfMatch** | **string** | Current ETag/version required for conditional update
**LanguageOperator** | [**LanguageOperator**](LanguageOperator.md) | Full replacement of a Language Operator configuration

### Return type

[**CreateOperatorResponse**](CreateOperator201Response.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

