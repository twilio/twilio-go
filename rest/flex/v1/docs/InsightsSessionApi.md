# InsightsSessionApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInsightsSession**](InsightsSessionApi.md#CreateInsightsSession) | **Post** /v1/Insights/Session | 



## CreateInsightsSession

> FlexV1InsightsSession CreateInsightsSession(ctx, optional)



To obtain session details for fetching reports and dashboards

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateInsightsSessionParams struct


Name | Type | Description
------------- | ------------- | -------------
**Authorization** | **string** | The Authorization HTTP request header

### Return type

[**FlexV1InsightsSession**](FlexV1InsightsSession.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

