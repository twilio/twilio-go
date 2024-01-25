# DeviceCodeApi

All URIs are relative to *https://oauth.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceCode**](DeviceCodeApi.md#CreateDeviceCode) | **Post** /v1/device/code | 



## CreateDeviceCode

> OauthV1DeviceCode CreateDeviceCode(ctx, optional)



Issues a new Access token (optionally identity_token & refresh_token) in exchange of Oauth grant

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateDeviceCodeParams struct


Name | Type | Description
------------- | ------------- | -------------
**ClientSid** | **string** | A 34 character string that uniquely identifies this OAuth App.
**Scopes** | **[]string** | An Array of scopes for authorization request
**Audiences** | **[]string** | An array of intended audiences for token requests

### Return type

[**OauthV1DeviceCode**](OauthV1DeviceCode.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

