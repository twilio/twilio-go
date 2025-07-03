# RateLimitsApi

All URIs are relative to *https://lookups.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchLookupAccountRateLimits**](RateLimitsApi.md#FetchLookupAccountRateLimits) | **Get** /v2/RateLimits | Get account rate limits



## FetchLookupAccountRateLimits

> RateLimitListResponse FetchLookupAccountRateLimits(ctx, optional)

Get account rate limits

Retrieve the list of rate limits for all fields (if any) It returns also the twilio rate limits.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchLookupAccountRateLimitsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Fields** | **[]string** | 

### Return type

[**RateLimitListResponse**](RateLimitListResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

