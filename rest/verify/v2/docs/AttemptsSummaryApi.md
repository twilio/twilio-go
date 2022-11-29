# AttemptsSummaryApi

All URIs are relative to *https://verify.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchVerificationAttemptsSummary**](AttemptsSummaryApi.md#FetchVerificationAttemptsSummary) | **Get** /v2/Attempts/Summary | 



## FetchVerificationAttemptsSummary

> VerifyV2VerificationAttemptsSummary FetchVerificationAttemptsSummary(ctx, optional)



Get a summary of how many attempts were made and how many were converted.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchVerificationAttemptsSummaryParams struct


Name | Type | Description
------------- | ------------- | -------------
**VerifyServiceSid** | **string** | Filter used to consider only Verification Attempts of the given verify service on the summary aggregation.
**DateCreatedAfter** | **time.Time** | Datetime filter used to consider only Verification Attempts created after this datetime on the summary aggregation. Given as GMT in RFC 2822 format.
**DateCreatedBefore** | **time.Time** | Datetime filter used to consider only Verification Attempts created before this datetime on the summary aggregation. Given as GMT in RFC 2822 format.
**Country** | **string** | Filter used to consider only Verification Attempts sent to the specified destination country on the summary aggregation.
**Channel** | **string** | Filter Verification Attempts considered on the summary aggregation by communication channel. Valid values are `SMS` and `CALL`
**DestinationPrefix** | **string** | Filter the Verification Attempts considered on the summary aggregation by Destination prefix. It is the prefix of a phone number in E.164 format.

### Return type

[**VerifyV2VerificationAttemptsSummary**](VerifyV2VerificationAttemptsSummary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

