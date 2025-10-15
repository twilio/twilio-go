# OperatorTypesApi

All URIs are relative to *https://intelligence.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchOperatorType**](OperatorTypesApi.md#FetchOperatorType) | **Get** /v2/OperatorTypes/{Sid} | Fetch a specific Operator Type.
[**ListOperatorType**](OperatorTypesApi.md#ListOperatorType) | **Get** /v2/OperatorTypes | Retrieves a list of all Operator Types for an Account.



## FetchOperatorType

> IntelligenceV2OperatorType FetchOperatorType(ctx, Sid)

Fetch a specific Operator Type.

Fetch a specific Operator Type.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | Either a 34 character string that uniquely identifies this Operator Type or the unique name that references an Operator Type.

### Other Parameters

Other parameters are passed through a pointer to a FetchOperatorTypeParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IntelligenceV2OperatorType**](IntelligenceV2OperatorType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperatorType

> []IntelligenceV2OperatorType ListOperatorType(ctx, optional)

Retrieves a list of all Operator Types for an Account.

Retrieves a list of all Operator Types for an Account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListOperatorTypeParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]IntelligenceV2OperatorType**](IntelligenceV2OperatorType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

