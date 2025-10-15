# ServicesPasskeysFactorsApi

All URIs are relative to *https://verify.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewFactorPasskey**](ServicesPasskeysFactorsApi.md#CreateNewFactorPasskey) | **Post** /v2/Services/{ServiceSid}/Passkeys/Factors | Create a new Passkeys Factor for the Entity



## CreateNewFactorPasskey

> CreateNewFactorPasskeyResponse CreateNewFactorPasskey(ctx, ServiceSidoptional)

Create a new Passkeys Factor for the Entity

Create a new Passkeys Factor for the Entity

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.

### Other Parameters

Other parameters are passed through a pointer to a CreateNewFactorPasskeyParams struct


Name | Type | Description
------------- | ------------- | -------------
**CreateNewPasskeysFactorRequest** | [**CreateNewPasskeysFactorRequest**](CreateNewPasskeysFactorRequest.md) | 

### Return type

[**CreateNewFactorPasskeyResponse**](CreateNewFactorPasskey201Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

