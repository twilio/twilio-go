# TokenApi

All URIs are relative to *https://oauth.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateToken**](TokenApi.md#CreateToken) | **Post** /v1/token | 



## CreateToken

> OauthV1Token CreateToken(ctx, optional)



Issues a new Access token (optionally identity_token & refresh_token) in exchange of Oauth grant

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateTokenParams struct


Name | Type | Description
------------- | ------------- | -------------
**GrantType** | **string** | Grant type is a credential representing resource owner's authorization which can be used by client to obtain access token.
**ClientSid** | **string** | A 34 character string that uniquely identifies this OAuth App.
**ClientSecret** | **string** | The credential for confidential OAuth App.
**Code** | **string** | JWT token related to the authorization code grant type.
**CodeVerifier** | **string** | A code which is generation cryptographically.
**DeviceCode** | **string** | JWT token related to the device code grant type.
**RefreshToken** | **string** | JWT token related to the refresh token grant type.
**DeviceId** | **string** | The Id of the device associated with the token (refresh token).

### Return type

[**OauthV1Token**](OauthV1Token.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

