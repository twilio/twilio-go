# ServicesPasskeysApproveChallengeApi

All URIs are relative to *https://verify.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateChallengePasskeys**](ServicesPasskeysApproveChallengeApi.md#UpdateChallengePasskeys) | **Post** /v2/Services/{ServiceSid}/Passkeys/ApproveChallenge | Approve a Passkeys Challenge



## UpdateChallengePasskeys

> CreateChallengePasskeysResponse UpdateChallengePasskeys(ctx, ServiceSidoptional)

Approve a Passkeys Challenge

Approve a Passkeys challenge

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.

### Other Parameters

Other parameters are passed through a pointer to a UpdateChallengePasskeysParams struct


Name | Type | Description
------------- | ------------- | -------------
**ApprovePasskeysChallengeRequest** | [**ApprovePasskeysChallengeRequest**](ApprovePasskeysChallengeRequest.md) | 

### Return type

[**CreateChallengePasskeysResponse**](CreateChallengePasskeys201Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

