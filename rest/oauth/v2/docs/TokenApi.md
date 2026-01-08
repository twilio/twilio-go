# TokenApi

All URIs are relative to *https://oauth.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOauth2Token**](TokenApi.md#CreateOauth2Token) | **Post** /v2/token | Issues a new Access token (optionally identity_token &amp; refresh_token) in exchange of Oauth grant



## CreateOauth2Token

> V2Oauth2TokenResponse CreateOauth2Token(ctx, optional)

Issues a new Access token (optionally identity_token & refresh_token) in exchange of Oauth grant

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateOauth2TokenParams struct


Name | Type | Description
------------- | ------------- | -------------
**AccountSid** | **string** | Optional Account SID to perform on behalf of requests.
**GrantType** | **string** | Grant type is a credential representing resource owner's authorization which can be used by client to obtain access token.
**ClientId** | **string** | A 34 character string that uniquely identifies this OAuth App.
**ClientSecret** | **string** | The credential for confidential OAuth App.
**Code** | **string** | JWT token related to the authorization code grant type.
**RedirectUri** | **string** | The redirect uri
**Audience** | **string** | The targeted audience uri
**RefreshToken** | **string** | JWT token related to refresh access token.
**Scope** | **string** | The scope of token

### Return type

[**V2Oauth2TokenResponse**](V2Oauth2TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

