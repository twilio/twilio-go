# AuthorizeApi

All URIs are relative to *https://oauth.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchOauth2Authorize**](AuthorizeApi.md#FetchOauth2Authorize) | **Get** /v2/authorize | Retrieves authorize uri



## FetchOauth2Authorize

> OauthV2Authorize FetchOauth2Authorize(ctx, optional)

Retrieves authorize uri

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchOauth2AuthorizeParams struct


Name | Type | Description
------------- | ------------- | -------------
**ResponseType** | **string** | 
**ClientId** | **string** | 
**RedirectUri** | **string** | 
**Scope** | **string** | 
**State** | **string** | 

### Return type

[**OauthV2Authorize**](OauthV2Authorize.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

