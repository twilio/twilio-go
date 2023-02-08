# InsightsUserRolesApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchInsightsUserRoles**](InsightsUserRolesApi.md#FetchInsightsUserRoles) | **Get** /v1/Insights/UserRoles | 



## FetchInsightsUserRoles

> FlexV1InsightsUserRoles FetchInsightsUserRoles(ctx, optional)



This is used by Flex UI and Quality Management to fetch the Flex Insights roles for the user

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchInsightsUserRolesParams struct


Name | Type | Description
------------- | ------------- | -------------
**Authorization** | **string** | The Authorization HTTP request header

### Return type

[**FlexV1InsightsUserRoles**](FlexV1InsightsUserRoles.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

