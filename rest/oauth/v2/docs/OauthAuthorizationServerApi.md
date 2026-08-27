# OauthAuthorizationServerApi

All URIs are relative to *https://oauth.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchOauth2ServerMetadata**](OauthAuthorizationServerApi.md#FetchOauth2ServerMetadata) | **Get** /.well-known/oauth-authorization-server | OAuth 2.0 Authorization Server Metadata



## FetchOauth2ServerMetadata

> FetchOauth2ServerMetadataResponse FetchOauth2ServerMetadata(ctx, )

OAuth 2.0 Authorization Server Metadata

Returns RFC 8414 Authorization Server Metadata for oauth.twilio.com

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchOauth2ServerMetadataParams struct


### Return type

[**FetchOauth2ServerMetadataResponse**](FetchOauth2ServerMetadata200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

