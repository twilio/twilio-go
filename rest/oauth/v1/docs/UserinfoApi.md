# UserinfoApi

All URIs are relative to *https://oauth.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchUserInfo**](UserinfoApi.md#FetchUserInfo) | **Get** /v1/userinfo | 



## FetchUserInfo

> OauthV1UserInfo FetchUserInfo(ctx, )



Retrieves the consented UserInfo and other claims about the logged-in subject (end-user).

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchUserInfoParams struct


### Return type

[**OauthV1UserInfo**](OauthV1UserInfo.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

