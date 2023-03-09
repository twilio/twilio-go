# InsightsQMAssessmentsApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInsightsAssessments**](InsightsQMAssessmentsApi.md#CreateInsightsAssessments) | **Post** /v1/Insights/QM/Assessments | 
[**ListInsightsAssessments**](InsightsQMAssessmentsApi.md#ListInsightsAssessments) | **Get** /v1/Insights/QM/Assessments | 
[**UpdateInsightsAssessments**](InsightsQMAssessmentsApi.md#UpdateInsightsAssessments) | **Post** /v1/Insights/QM/Assessments/{AssessmentId} | 



## CreateInsightsAssessments

> FlexV1InsightsAssessments CreateInsightsAssessments(ctx, optional)



Add assessments against conversation to dynamo db. Used in assessments screen by user. Users can select the questionnaire and pick up answers for each and every question.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateInsightsAssessmentsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Token** | **string** | The Token HTTP request header
**CategoryId** | **string** | The id of the category 
**CategoryName** | **string** | The name of the category
**SegmentId** | **string** | Segment Id of the conversation
**UserName** | **string** | Name of the user assessing conversation
**UserEmail** | **string** | Email of the user assessing conversation
**AgentId** | **string** | The id of the Agent
**Offset** | **float32** | The offset of the conversation.
**MetricId** | **string** | The question Id selected for assessment
**MetricName** | **string** | The question name of the assessment
**AnswerText** | **string** | The answer text selected by user
**AnswerId** | **string** | The id of the answer selected by user
**QuestionnaireId** | **string** | Questionnaire Id of the associated question

### Return type

[**FlexV1InsightsAssessments**](FlexV1InsightsAssessments.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInsightsAssessments

> []FlexV1InsightsAssessments ListInsightsAssessments(ctx, optional)



Get assessments done for a conversation by logged in user

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListInsightsAssessmentsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Token** | **string** | The Token HTTP request header
**SegmentId** | **string** | The id of the segment.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]FlexV1InsightsAssessments**](FlexV1InsightsAssessments.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInsightsAssessments

> FlexV1InsightsAssessments UpdateInsightsAssessments(ctx, AssessmentIdoptional)



Update a specific Assessment assessed earlier

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssessmentId** | **string** | The id of the assessment to be modified

### Other Parameters

Other parameters are passed through a pointer to a UpdateInsightsAssessmentsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Token** | **string** | The Token HTTP request header
**Offset** | **float32** | The offset of the conversation
**AnswerText** | **string** | The answer text selected by user
**AnswerId** | **string** | The id of the answer selected by user

### Return type

[**FlexV1InsightsAssessments**](FlexV1InsightsAssessments.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

