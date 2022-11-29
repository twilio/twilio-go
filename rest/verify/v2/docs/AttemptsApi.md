# AttemptsApi

All URIs are relative to *https://verify.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchVerificationAttempt**](AttemptsApi.md#FetchVerificationAttempt) | **Get** /v2/Attempts/{Sid} | 
[**ListVerificationAttempt**](AttemptsApi.md#ListVerificationAttempt) | **Get** /v2/Attempts | 



## FetchVerificationAttempt

> VerifyV2VerificationAttempt FetchVerificationAttempt(ctx, Sid)



Fetch a specific verification attempt.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique SID identifier of a Verification Attempt

### Other Parameters

Other parameters are passed through a pointer to a FetchVerificationAttemptParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VerifyV2VerificationAttempt**](VerifyV2VerificationAttempt.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVerificationAttempt

> []VerifyV2VerificationAttempt ListVerificationAttempt(ctx, optional)



List all the verification attempts for a given Account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListVerificationAttemptParams struct


Name | Type | Description
------------- | ------------- | -------------
**DateCreatedAfter** | **time.Time** | Datetime filter used to query Verification Attempts created after this datetime. Given as GMT in RFC 2822 format.
**DateCreatedBefore** | **time.Time** | Datetime filter used to query Verification Attempts created before this datetime. Given as GMT in RFC 2822 format.
**ChannelDataTo** | **string** | Destination of a verification. It is phone number in E.164 format.
**Country** | **string** | Filter used to query Verification Attempts sent to the specified destination country.
**Channel** | **string** | Filter used to query Verification Attempts by communication channel. Valid values are `SMS` and `CALL`
**VerifyServiceSid** | **string** | Filter used to query Verification Attempts by verify service. Only attempts of the provided SID will be returned.
**VerificationSid** | **string** | Filter used to return all the Verification Attempts of a single verification. Only attempts of the provided verification SID will be returned.
**Status** | **string** | Filter used to query Verification Attempts by conversion status. Valid values are `UNCONVERTED`, for attempts that were not converted, and `CONVERTED`, for attempts that were confirmed.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VerifyV2VerificationAttempt**](VerifyV2VerificationAttempt.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

