# HostedNumberEligibilityBulkApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBulkEligibility**](HostedNumberEligibilityBulkApi.md#CreateBulkEligibility) | **Post** /v1/HostedNumber/Eligibility/Bulk | Create a bulk eligibility check for a set of numbers that you want to host in Twilio.
[**FetchBulkEligibility**](HostedNumberEligibilityBulkApi.md#FetchBulkEligibility) | **Get** /v1/HostedNumber/Eligibility/Bulk/{RequestId} | Fetch an eligibility bulk check that you requested to host in Twilio.



## CreateBulkEligibility

> NumbersV1BulkEligibility CreateBulkEligibility(ctx, optional)

Create a bulk eligibility check for a set of numbers that you want to host in Twilio.

Create a bulk eligibility check for a set of numbers that you want to host in Twilio.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateBulkEligibilityParams struct


Name | Type | Description
------------- | ------------- | -------------
**Body** | **map[string]interface{}** | 

### Return type

[**NumbersV1BulkEligibility**](NumbersV1BulkEligibility.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchBulkEligibility

> NumbersV1BulkEligibility FetchBulkEligibility(ctx, RequestId)

Fetch an eligibility bulk check that you requested to host in Twilio.

Fetch an eligibility bulk check that you requested to host in Twilio.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RequestId** | **string** | The SID of the bulk eligibility check that you want to know about.

### Other Parameters

Other parameters are passed through a pointer to a FetchBulkEligibilityParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NumbersV1BulkEligibility**](NumbersV1BulkEligibility.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

