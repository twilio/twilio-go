# OperatorsApi

All URIs are relative to *https://intelligence.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchOperator**](OperatorsApi.md#FetchOperator) | **Get** /v2/Operators/{Sid} | Fetch a specific Operator. Works for both Custom and Pre-built Operators.
[**ListOperator**](OperatorsApi.md#ListOperator) | **Get** /v2/Operators | Retrieves a list of all Operators for an Account, both Custom and Pre-built.



## FetchOperator

> IntelligenceV2Operator FetchOperator(ctx, Sid)

Fetch a specific Operator. Works for both Custom and Pre-built Operators.

Fetch a specific Operator. Works for both Custom and Pre-built Operators.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Operator.

### Other Parameters

Other parameters are passed through a pointer to a FetchOperatorParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IntelligenceV2Operator**](IntelligenceV2Operator.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperator

> []IntelligenceV2Operator ListOperator(ctx, optional)

Retrieves a list of all Operators for an Account, both Custom and Pre-built.

Retrieves a list of all Operators for an Account, both Custom and Pre-built.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListOperatorParams struct


Name | Type | Description
------------- | ------------- | -------------
**Availability** | [**string**](stringstring.md) | Returns Operators with the provided availability type. Possible values: internal, beta, public, retired.
**LanguageCode** | **string** | Returns Operators that support the provided language code.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]IntelligenceV2Operator**](IntelligenceV2Operator.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

