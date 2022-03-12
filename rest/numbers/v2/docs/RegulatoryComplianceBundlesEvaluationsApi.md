# RegulatoryComplianceBundlesEvaluationsApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEvaluation**](RegulatoryComplianceBundlesEvaluationsApi.md#CreateEvaluation) | **Post** /v2/RegulatoryCompliance/Bundles/{BundleSid}/Evaluations | 
[**FetchEvaluation**](RegulatoryComplianceBundlesEvaluationsApi.md#FetchEvaluation) | **Get** /v2/RegulatoryCompliance/Bundles/{BundleSid}/Evaluations/{Sid} | 
[**ListEvaluation**](RegulatoryComplianceBundlesEvaluationsApi.md#ListEvaluation) | **Get** /v2/RegulatoryCompliance/Bundles/{BundleSid}/Evaluations | 



## CreateEvaluation

> NumbersV2Evaluation CreateEvaluation(ctx, BundleSid)



Creates an evaluation for a bundle

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string** | The unique string that identifies the Bundle resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateEvaluationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NumbersV2Evaluation**](NumbersV2Evaluation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchEvaluation

> NumbersV2Evaluation FetchEvaluation(ctx, BundleSidSid)



Fetch specific Evaluation Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string** | The unique string that we created to identify the Bundle resource.
**Sid** | **string** | The unique string that identifies the Evaluation resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchEvaluationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NumbersV2Evaluation**](NumbersV2Evaluation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvaluation

> []NumbersV2Evaluation ListEvaluation(ctx, BundleSidoptional)



Retrieve a list of Evaluations associated to the Bundle resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string** | The unique string that identifies the Bundle resource.

### Other Parameters

Other parameters are passed through a pointer to a ListEvaluationParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]NumbersV2Evaluation**](NumbersV2Evaluation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

