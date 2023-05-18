# InsightsQualityManagementAssessmentsCommentsApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInsightsAssessmentsComment**](InsightsQualityManagementAssessmentsCommentsApi.md#CreateInsightsAssessmentsComment) | **Post** /v1/Insights/QualityManagement/Assessments/Comments | 
[**ListInsightsAssessmentsComment**](InsightsQualityManagementAssessmentsCommentsApi.md#ListInsightsAssessmentsComment) | **Get** /v1/Insights/QualityManagement/Assessments/Comments | 



## CreateInsightsAssessmentsComment

> FlexV1InsightsAssessmentsComment CreateInsightsAssessmentsComment(ctx, optional)



To create a comment assessment for a conversation

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateInsightsAssessmentsCommentParams struct


Name | Type | Description
------------- | ------------- | -------------
**Authorization** | **string** | The Authorization HTTP request header
**CategoryId** | **string** | The ID of the category
**CategoryName** | **string** | The name of the category
**Comment** | **string** | The Assessment comment.
**SegmentId** | **string** | The id of the segment.
**AgentId** | **string** | The id of the agent.
**Offset** | **float32** | The offset

### Return type

[**FlexV1InsightsAssessmentsComment**](FlexV1InsightsAssessmentsComment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInsightsAssessmentsComment

> []FlexV1InsightsAssessmentsComment ListInsightsAssessmentsComment(ctx, optional)



To create a comment assessment for a conversation

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListInsightsAssessmentsCommentParams struct


Name | Type | Description
------------- | ------------- | -------------
**Authorization** | **string** | The Authorization HTTP request header
**SegmentId** | **string** | The id of the segment.
**AgentId** | **string** | The id of the agent.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]FlexV1InsightsAssessmentsComment**](FlexV1InsightsAssessmentsComment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

