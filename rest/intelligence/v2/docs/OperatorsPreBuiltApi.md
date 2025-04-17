# OperatorsPreBuiltApi

All URIs are relative to *https://intelligence.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchPrebuiltOperator**](OperatorsPreBuiltApi.md#FetchPrebuiltOperator) | **Get** /v2/Operators/PreBuilt/{Sid} | Fetch a specific Pre-built Operator.
[**ListPrebuiltOperator**](OperatorsPreBuiltApi.md#ListPrebuiltOperator) | **Get** /v2/Operators/PreBuilt | Retrieves a list of all Pre-built Operators for an account.



## FetchPrebuiltOperator

> IntelligenceV2PrebuiltOperator FetchPrebuiltOperator(ctx, Sid)

Fetch a specific Pre-built Operator.

Fetch a specific Pre-built Operator.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Pre-built Operator.

### Other Parameters

Other parameters are passed through a pointer to a FetchPrebuiltOperatorParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IntelligenceV2PrebuiltOperator**](IntelligenceV2PrebuiltOperator.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrebuiltOperator

> []IntelligenceV2PrebuiltOperator ListPrebuiltOperator(ctx, optional)

Retrieves a list of all Pre-built Operators for an account.

Retrieves a list of all Pre-built Operators for an account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListPrebuiltOperatorParams struct


Name | Type | Description
------------- | ------------- | -------------
**Availability** | [**string**](stringstring.md) | Returns Pre-built Operators with the provided availability type. Possible values: internal, beta, public, retired.
**LanguageCode** | **string** | Returns Pre-built Operators that support the provided language code.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]IntelligenceV2PrebuiltOperator**](IntelligenceV2PrebuiltOperator.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

