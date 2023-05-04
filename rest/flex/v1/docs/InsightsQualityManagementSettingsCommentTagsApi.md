# InsightsQualityManagementSettingsCommentTagsApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchInsightsSettingsComment**](InsightsQualityManagementSettingsCommentTagsApi.md#FetchInsightsSettingsComment) | **Get** /v1/Insights/QualityManagement/Settings/CommentTags | 



## FetchInsightsSettingsComment

> FlexV1InsightsSettingsComment FetchInsightsSettingsComment(ctx, optional)



To get the Comment Settings for an Account

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchInsightsSettingsCommentParams struct


Name | Type | Description
------------- | ------------- | -------------
**Token** | **string** | The Token HTTP request header

### Return type

[**FlexV1InsightsSettingsComment**](FlexV1InsightsSettingsComment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

