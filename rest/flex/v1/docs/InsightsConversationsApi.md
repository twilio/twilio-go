# InsightsConversationsApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListInsightsConversations**](InsightsConversationsApi.md#ListInsightsConversations) | **Get** /v1/Insights/Conversations | 



## ListInsightsConversations

> []FlexV1InsightsConversations ListInsightsConversations(ctx, optional)



To get conversation with segment id

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListInsightsConversationsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Token** | **string** | The Token HTTP request header
**SegmentId** | **string** | Unique Id of the segment for which conversation details needs to be fetched
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]FlexV1InsightsConversations**](FlexV1InsightsConversations.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

