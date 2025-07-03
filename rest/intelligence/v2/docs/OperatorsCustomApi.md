# OperatorsCustomApi

All URIs are relative to *https://intelligence.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomOperator**](OperatorsCustomApi.md#CreateCustomOperator) | **Post** /v2/Operators/Custom | Create a new Custom Operator for the given Account
[**DeleteCustomOperator**](OperatorsCustomApi.md#DeleteCustomOperator) | **Delete** /v2/Operators/Custom/{Sid} | Delete a specific Custom Operator.
[**FetchCustomOperator**](OperatorsCustomApi.md#FetchCustomOperator) | **Get** /v2/Operators/Custom/{Sid} | Fetch a specific Custom Operator.
[**ListCustomOperator**](OperatorsCustomApi.md#ListCustomOperator) | **Get** /v2/Operators/Custom | Retrieves a list of all Custom Operators for an Account.
[**UpdateCustomOperator**](OperatorsCustomApi.md#UpdateCustomOperator) | **Post** /v2/Operators/Custom/{Sid} | Update a specific Custom Operator.



## CreateCustomOperator

> IntelligenceV2CustomOperator CreateCustomOperator(ctx, optional)

Create a new Custom Operator for the given Account

Create a new Custom Operator for the given Account

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateCustomOperatorParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A human readable description of the new Operator, up to 64 characters.
**OperatorType** | **string** | Operator Type for this Operator. References an existing Operator Type resource.
**Config** | [**interface{}**](interface{}.md) | Operator configuration, following the schema defined by the Operator Type.

### Return type

[**IntelligenceV2CustomOperator**](IntelligenceV2CustomOperator.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomOperator

> DeleteCustomOperator(ctx, Sid)

Delete a specific Custom Operator.

Delete a specific Custom Operator.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Custom Operator.

### Other Parameters

Other parameters are passed through a pointer to a DeleteCustomOperatorParams struct


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


## FetchCustomOperator

> IntelligenceV2CustomOperator FetchCustomOperator(ctx, Sid)

Fetch a specific Custom Operator.

Fetch a specific Custom Operator.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Custom Operator.

### Other Parameters

Other parameters are passed through a pointer to a FetchCustomOperatorParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IntelligenceV2CustomOperator**](IntelligenceV2CustomOperator.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomOperator

> []IntelligenceV2CustomOperator ListCustomOperator(ctx, optional)

Retrieves a list of all Custom Operators for an Account.

Retrieves a list of all Custom Operators for an Account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListCustomOperatorParams struct


Name | Type | Description
------------- | ------------- | -------------
**Availability** | [**string**](stringstring.md) | Returns Custom Operators with the provided availability type. Possible values: internal, beta, public, retired.
**LanguageCode** | **string** | Returns Custom Operators that support the provided language code.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]IntelligenceV2CustomOperator**](IntelligenceV2CustomOperator.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomOperator

> IntelligenceV2CustomOperator UpdateCustomOperator(ctx, Sidoptional)

Update a specific Custom Operator.

Update a specific Custom Operator.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Custom Operator.

### Other Parameters

Other parameters are passed through a pointer to a UpdateCustomOperatorParams struct


Name | Type | Description
------------- | ------------- | -------------
**IfMatch** | **string** | The If-Match HTTP request header
**FriendlyName** | **string** | A human-readable name of this resource, up to 64 characters.
**Config** | [**interface{}**](interface{}.md) | Operator configuration, following the schema defined by the Operator Type.

### Return type

[**IntelligenceV2CustomOperator**](IntelligenceV2CustomOperator.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

