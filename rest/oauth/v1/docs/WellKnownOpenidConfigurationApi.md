# WellKnownOpenidConfigurationApi

All URIs are relative to *https://oauth.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchOpenidDiscovery**](WellKnownOpenidConfigurationApi.md#FetchOpenidDiscovery) | **Get** /v1/.well-known/openid-configuration | 



## FetchOpenidDiscovery

> OauthV1OpenidDiscovery FetchOpenidDiscovery(ctx, )



Fetch configuration details about the OpenID Connect Authorization Server

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchOpenidDiscoveryParams struct


### Return type

[**OauthV1OpenidDiscovery**](OauthV1OpenidDiscovery.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

