# ServicesRateLimitsApi

All URIs are relative to *https://verify.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRateLimit**](ServicesRateLimitsApi.md#CreateRateLimit) | **Post** /v2/Services/{ServiceSid}/RateLimits | 
[**DeleteRateLimit**](ServicesRateLimitsApi.md#DeleteRateLimit) | **Delete** /v2/Services/{ServiceSid}/RateLimits/{Sid} | 
[**FetchRateLimit**](ServicesRateLimitsApi.md#FetchRateLimit) | **Get** /v2/Services/{ServiceSid}/RateLimits/{Sid} | 
[**ListRateLimit**](ServicesRateLimitsApi.md#ListRateLimit) | **Get** /v2/Services/{ServiceSid}/RateLimits | 
[**UpdateRateLimit**](ServicesRateLimitsApi.md#UpdateRateLimit) | **Post** /v2/Services/{ServiceSid}/RateLimits/{Sid} | 



## CreateRateLimit

> VerifyV2RateLimit CreateRateLimit(ctx, ServiceSidoptional)



Create a new Rate Limit for a Service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.

### Other Parameters

Other parameters are passed through a pointer to a CreateRateLimitParams struct


Name | Type | Description
------------- | ------------- | -------------
**UniqueName** | **string** | Provides a unique and addressable name to be assigned to this Rate Limit, assigned by the developer, to be optionally used in addition to SID. **This value should not contain PII.**
**Description** | **string** | Description of this Rate Limit

### Return type

[**VerifyV2RateLimit**](VerifyV2RateLimit.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRateLimit

> DeleteRateLimit(ctx, ServiceSidSid)



Delete a specific Rate Limit.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Rate Limit resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a DeleteRateLimitParams struct


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


## FetchRateLimit

> VerifyV2RateLimit FetchRateLimit(ctx, ServiceSidSid)



Fetch a specific Rate Limit.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Rate Limit resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRateLimitParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VerifyV2RateLimit**](VerifyV2RateLimit.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRateLimit

> []VerifyV2RateLimit ListRateLimit(ctx, ServiceSidoptional)



Retrieve a list of all Rate Limits for a service.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.

### Other Parameters

Other parameters are passed through a pointer to a ListRateLimitParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VerifyV2RateLimit**](VerifyV2RateLimit.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRateLimit

> VerifyV2RateLimit UpdateRateLimit(ctx, ServiceSidSidoptional)



Update a specific Rate Limit.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Rate Limit resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRateLimitParams struct


Name | Type | Description
------------- | ------------- | -------------
**Description** | **string** | Description of this Rate Limit

### Return type

[**VerifyV2RateLimit**](VerifyV2RateLimit.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

