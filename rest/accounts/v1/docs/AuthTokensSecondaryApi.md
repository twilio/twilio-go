# AuthTokensSecondaryApi

All URIs are relative to *https://accounts.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecondaryAuthToken**](AuthTokensSecondaryApi.md#CreateSecondaryAuthToken) | **Post** /v1/AuthTokens/Secondary | 
[**DeleteSecondaryAuthToken**](AuthTokensSecondaryApi.md#DeleteSecondaryAuthToken) | **Delete** /v1/AuthTokens/Secondary | 



## CreateSecondaryAuthToken

> AccountsV1SecondaryAuthToken CreateSecondaryAuthToken(ctx, )



Create a new secondary Auth Token

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSecondaryAuthTokenParams struct


### Return type

[**AccountsV1SecondaryAuthToken**](AccountsV1SecondaryAuthToken.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecondaryAuthToken

> DeleteSecondaryAuthToken(ctx, )



Delete the secondary Auth Token from your account

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSecondaryAuthTokenParams struct


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

