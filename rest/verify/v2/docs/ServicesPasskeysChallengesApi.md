# ServicesPasskeysChallengesApi

All URIs are relative to *https://verify.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChallengePasskeys**](ServicesPasskeysChallengesApi.md#CreateChallengePasskeys) | **Post** /v2/Services/{ServiceSid}/Passkeys/Challenges | Create a Passkeys Challenge



## CreateChallengePasskeys

> CreateChallengePasskeysResponse CreateChallengePasskeys(ctx, ServiceSidoptional)

Create a Passkeys Challenge

Create a Passkeys Challenge

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.

### Other Parameters

Other parameters are passed through a pointer to a CreateChallengePasskeysParams struct


Name | Type | Description
------------- | ------------- | -------------
**CreatePasskeysChallengeRequest** | [**CreatePasskeysChallengeRequest**](CreatePasskeysChallengeRequest.md) | 

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

