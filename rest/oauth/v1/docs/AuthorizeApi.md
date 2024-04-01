# AuthorizeApi

All URIs are relative to *https://oauth.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAuthorize**](AuthorizeApi.md#FetchAuthorize) | **Get** /v1/authorize | 



## FetchAuthorize

> OauthV1Authorize FetchAuthorize(ctx, optional)



Retrieves authorize uri

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchAuthorizeParams struct


Name | Type | Description
------------- | ------------- | -------------
**ResponseType** | **string** | Response Type
**ClientId** | **string** | The Client Identifier
**RedirectUri** | **string** | The url to which response will be redirected to
**Scope** | **string** | The scope of the access request
**State** | **string** | An opaque value which can be used to maintain state between the request and callback

### Return type

[**OauthV1Authorize**](OauthV1Authorize.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

