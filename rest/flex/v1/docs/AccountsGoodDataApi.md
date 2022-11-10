# AccountsGoodDataApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGooddata**](AccountsGoodDataApi.md#CreateGooddata) | **Post** /v1/Accounts/GoodData | 



## CreateGooddata

> FlexV1Gooddata CreateGooddata(ctx, optional)



To create a GoodData Session id to access GoodData dashboards

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateGooddataParams struct


Name | Type | Description
------------- | ------------- | -------------
**Token** | **string** | The Token HTTP request header

### Return type

[**FlexV1Gooddata**](FlexV1Gooddata.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

