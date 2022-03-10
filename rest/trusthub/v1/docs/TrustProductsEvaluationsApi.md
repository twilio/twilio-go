# TrustProductsEvaluationsApi

All URIs are relative to *https://trusthub.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTrustProductEvaluation**](TrustProductsEvaluationsApi.md#CreateTrustProductEvaluation) | **Post** /v1/TrustProducts/{TrustProductSid}/Evaluations | 
[**FetchTrustProductEvaluation**](TrustProductsEvaluationsApi.md#FetchTrustProductEvaluation) | **Get** /v1/TrustProducts/{TrustProductSid}/Evaluations/{Sid} | 
[**ListTrustProductEvaluation**](TrustProductsEvaluationsApi.md#ListTrustProductEvaluation) | **Get** /v1/TrustProducts/{TrustProductSid}/Evaluations | 



## CreateTrustProductEvaluation

> TrusthubV1TrustProductEvaluation CreateTrustProductEvaluation(ctx, TrustProductSidoptional)



Create a new Evaluation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrustProductSid** | **string** | The unique string that we created to identify the trust_product resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateTrustProductEvaluationParams struct


Name | Type | Description
------------- | ------------- | -------------
**PolicySid** | **string** | The unique string of a policy that is associated to the customer_profile resource.

### Return type

[**TrusthubV1TrustProductEvaluation**](TrusthubV1TrustProductEvaluation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTrustProductEvaluation

> TrusthubV1TrustProductEvaluation FetchTrustProductEvaluation(ctx, TrustProductSidSid)



Fetch specific Evaluation Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrustProductSid** | **string** | The unique string that we created to identify the trust_product resource.
**Sid** | **string** | The unique string that identifies the Evaluation resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchTrustProductEvaluationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrusthubV1TrustProductEvaluation**](TrusthubV1TrustProductEvaluation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrustProductEvaluation

> []TrusthubV1TrustProductEvaluation ListTrustProductEvaluation(ctx, TrustProductSidoptional)



Retrieve a list of Evaluations associated to the trust_product resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrustProductSid** | **string** | The unique string that we created to identify the trust_product resource.

### Other Parameters

Other parameters are passed through a pointer to a ListTrustProductEvaluationParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TrusthubV1TrustProductEvaluation**](TrusthubV1TrustProductEvaluation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

