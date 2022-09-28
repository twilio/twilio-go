# ServicesRateLimitsBucketsApi

All URIs are relative to *https://verify.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBucket**](ServicesRateLimitsBucketsApi.md#CreateBucket) | **Post** /v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets | 
[**DeleteBucket**](ServicesRateLimitsBucketsApi.md#DeleteBucket) | **Delete** /v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets/{Sid} | 
[**FetchBucket**](ServicesRateLimitsBucketsApi.md#FetchBucket) | **Get** /v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets/{Sid} | 
[**ListBucket**](ServicesRateLimitsBucketsApi.md#ListBucket) | **Get** /v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets | 
[**UpdateBucket**](ServicesRateLimitsBucketsApi.md#UpdateBucket) | **Post** /v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets/{Sid} | 



## CreateBucket

> VerifyV2Bucket CreateBucket(ctx, ServiceSidRateLimitSidoptional)



Create a new Bucket for a Rate Limit

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
**RateLimitSid** | **string** | The Twilio-provided string that uniquely identifies the Rate Limit resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateBucketParams struct


Name | Type | Description
------------- | ------------- | -------------
**Max** | **int** | Maximum number of requests permitted in during the interval.
**Interval** | **int** | Number of seconds that the rate limit will be enforced over.

### Return type

[**VerifyV2Bucket**](VerifyV2Bucket.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBucket

> DeleteBucket(ctx, ServiceSidRateLimitSidSid)



Delete a specific Bucket.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
**RateLimitSid** | **string** | The Twilio-provided string that uniquely identifies the Rate Limit resource.
**Sid** | **string** | A 34 character string that uniquely identifies this Bucket.

### Other Parameters

Other parameters are passed through a pointer to a DeleteBucketParams struct


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


## FetchBucket

> VerifyV2Bucket FetchBucket(ctx, ServiceSidRateLimitSidSid)



Fetch a specific Bucket.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
**RateLimitSid** | **string** | The Twilio-provided string that uniquely identifies the Rate Limit resource.
**Sid** | **string** | A 34 character string that uniquely identifies this Bucket.

### Other Parameters

Other parameters are passed through a pointer to a FetchBucketParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VerifyV2Bucket**](VerifyV2Bucket.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBucket

> []VerifyV2Bucket ListBucket(ctx, ServiceSidRateLimitSidoptional)



Retrieve a list of all Buckets for a Rate Limit.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
**RateLimitSid** | **string** | The Twilio-provided string that uniquely identifies the Rate Limit resource.

### Other Parameters

Other parameters are passed through a pointer to a ListBucketParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VerifyV2Bucket**](VerifyV2Bucket.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBucket

> VerifyV2Bucket UpdateBucket(ctx, ServiceSidRateLimitSidSidoptional)



Update a specific Bucket.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
**RateLimitSid** | **string** | The Twilio-provided string that uniquely identifies the Rate Limit resource.
**Sid** | **string** | A 34 character string that uniquely identifies this Bucket.

### Other Parameters

Other parameters are passed through a pointer to a UpdateBucketParams struct


Name | Type | Description
------------- | ------------- | -------------
**Max** | **int** | Maximum number of requests permitted in during the interval.
**Interval** | **int** | Number of seconds that the rate limit will be enforced over.

### Return type

[**VerifyV2Bucket**](VerifyV2Bucket.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

