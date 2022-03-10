# CustomerProfilesEvaluationsApi

All URIs are relative to *https://trusthub.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomerProfileEvaluation**](CustomerProfilesEvaluationsApi.md#CreateCustomerProfileEvaluation) | **Post** /v1/CustomerProfiles/{CustomerProfileSid}/Evaluations | 
[**FetchCustomerProfileEvaluation**](CustomerProfilesEvaluationsApi.md#FetchCustomerProfileEvaluation) | **Get** /v1/CustomerProfiles/{CustomerProfileSid}/Evaluations/{Sid} | 
[**ListCustomerProfileEvaluation**](CustomerProfilesEvaluationsApi.md#ListCustomerProfileEvaluation) | **Get** /v1/CustomerProfiles/{CustomerProfileSid}/Evaluations | 



## CreateCustomerProfileEvaluation

> TrusthubV1CustomerProfileEvaluation CreateCustomerProfileEvaluation(ctx, CustomerProfileSidoptional)



Create a new Evaluation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CustomerProfileSid** | **string** | The unique string that we created to identify the CustomerProfile resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateCustomerProfileEvaluationParams struct


Name | Type | Description
------------- | ------------- | -------------
**PolicySid** | **string** | The unique string of a policy that is associated to the customer_profile resource.

### Return type

[**TrusthubV1CustomerProfileEvaluation**](TrusthubV1CustomerProfileEvaluation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCustomerProfileEvaluation

> TrusthubV1CustomerProfileEvaluation FetchCustomerProfileEvaluation(ctx, CustomerProfileSidSid)



Fetch specific Evaluation Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CustomerProfileSid** | **string** | The unique string that we created to identify the customer_profile resource.
**Sid** | **string** | The unique string that identifies the Evaluation resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchCustomerProfileEvaluationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrusthubV1CustomerProfileEvaluation**](TrusthubV1CustomerProfileEvaluation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomerProfileEvaluation

> []TrusthubV1CustomerProfileEvaluation ListCustomerProfileEvaluation(ctx, CustomerProfileSidoptional)



Retrieve a list of Evaluations associated to the customer_profile resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CustomerProfileSid** | **string** | The unique string that we created to identify the CustomerProfile resource.

### Other Parameters

Other parameters are passed through a pointer to a ListCustomerProfileEvaluationParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TrusthubV1CustomerProfileEvaluation**](TrusthubV1CustomerProfileEvaluation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

