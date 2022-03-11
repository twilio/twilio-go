# PoliciesApi

All URIs are relative to *https://trusthub.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchPolicies**](PoliciesApi.md#FetchPolicies) | **Get** /v1/Policies/{Sid} | 
[**ListPolicies**](PoliciesApi.md#ListPolicies) | **Get** /v1/Policies | 



## FetchPolicies

> TrusthubV1Policies FetchPolicies(ctx, Sid)



Fetch specific Policy Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that identifies the Policy resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchPoliciesParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrusthubV1Policies**](TrusthubV1Policies.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicies

> []TrusthubV1Policies ListPolicies(ctx, optional)



Retrieve a list of all Policys.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListPoliciesParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TrusthubV1Policies**](TrusthubV1Policies.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

