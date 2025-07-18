# RateLimitsFieldsBucketApi

All URIs are relative to *https://lookups.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLookupRateLimit**](RateLimitsFieldsBucketApi.md#DeleteLookupRateLimit) | **Delete** /v2/RateLimits/Fields/{Field}/Bucket/{Bucket} | Delete rate limit
[**FetchLookupRateLimit**](RateLimitsFieldsBucketApi.md#FetchLookupRateLimit) | **Get** /v2/RateLimits/Fields/{Field}/Bucket/{Bucket} | Get rate limit
[**UpdateLookupRateLimit**](RateLimitsFieldsBucketApi.md#UpdateLookupRateLimit) | **Put** /v2/RateLimits/Fields/{Field}/Bucket/{Bucket} | Upsert rate limit



## DeleteLookupRateLimit

> DeleteLookupRateLimit(ctx, FieldBucket)

Delete rate limit

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Field** | **string** | bucket name
**Bucket** | **string** | bucket name

### Other Parameters

Other parameters are passed through a pointer to a DeleteLookupRateLimitParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchLookupRateLimit

> RateLimitResponse FetchLookupRateLimit(ctx, FieldBucket)

Get rate limit

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Field** | **string** | bucket name
**Bucket** | **string** | bucket name

### Other Parameters

Other parameters are passed through a pointer to a FetchLookupRateLimitParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**RateLimitResponse**](RateLimitResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLookupRateLimit

> RateLimitResponse UpdateLookupRateLimit(ctx, FieldBucketoptional)

Upsert rate limit

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Field** | **string** | field name
**Bucket** | **string** | bucket name

### Other Parameters

Other parameters are passed through a pointer to a UpdateLookupRateLimitParams struct


Name | Type | Description
------------- | ------------- | -------------
**RateLimitRequest** | [**RateLimitRequest**](RateLimitRequest.md) | 

### Return type

[**RateLimitResponse**](RateLimitResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

