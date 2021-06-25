# AuthTokensPromoteApi

All URIs are relative to *https://accounts.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateAuthTokenPromotion**](AuthTokensPromoteApi.md#UpdateAuthTokenPromotion) | **Post** /v1/AuthTokens/Promote | 



## UpdateAuthTokenPromotion

> AccountsV1AuthTokenPromotion UpdateAuthTokenPromotion(ctx, )



Promote the secondary Auth Token to primary. After promoting the new token, all requests to Twilio using your old primary Auth Token will result in an error.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateAuthTokenPromotionParams struct


### Return type

[**AccountsV1AuthTokenPromotion**](AccountsV1AuthTokenPromotion.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

