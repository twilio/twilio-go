# AccountsTokensApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateToken**](AccountsTokensApi.md#CreateToken) | **Post** /2010-04-01/Accounts/{AccountSid}/Tokens.json | 



## CreateToken

> ApiV2010Token CreateToken(ctx, optional)



Create a new token for ICE servers

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateTokenParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**Ttl** | **int** | The duration in seconds for which the generated credentials are valid. The default value is 86400 (24 hours).

### Return type

[**ApiV2010Token**](ApiV2010Token.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

