# HostedNumberEligibilityApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEligibility**](HostedNumberEligibilityApi.md#CreateEligibility) | **Post** /v1/HostedNumber/Eligibility | 



## CreateEligibility

> NumbersV1Eligibility CreateEligibility(ctx, )



Create an eligibility check for a number that you want to host in Twilio.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateEligibilityParams struct


### Return type

[**NumbersV1Eligibility**](NumbersV1Eligibility.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

