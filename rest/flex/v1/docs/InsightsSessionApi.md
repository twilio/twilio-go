# InsightsSessionApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGooddata**](InsightsSessionApi.md#CreateGooddata) | **Post** /v1/Insights/Session | 



## CreateGooddata

> FlexV1Gooddata CreateGooddata(ctx, optional)



To obtain session details for fetching reports and dashboards

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

