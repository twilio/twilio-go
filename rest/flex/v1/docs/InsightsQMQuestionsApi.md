# InsightsQMQuestionsApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInsightsQuestionnairesQuestion**](InsightsQMQuestionsApi.md#CreateInsightsQuestionnairesQuestion) | **Post** /v1/Insights/QM/Questions | 
[**DeleteInsightsQuestionnairesQuestion**](InsightsQMQuestionsApi.md#DeleteInsightsQuestionnairesQuestion) | **Delete** /v1/Insights/QM/Questions/{QuestionId} | 
[**UpdateInsightsQuestionnairesQuestion**](InsightsQMQuestionsApi.md#UpdateInsightsQuestionnairesQuestion) | **Post** /v1/Insights/QM/Questions/{QuestionId} | 



## CreateInsightsQuestionnairesQuestion

> FlexV1InsightsQuestionnairesQuestion CreateInsightsQuestionnairesQuestion(ctx, optional)



To create a question for a Category

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateInsightsQuestionnairesQuestionParams struct


Name | Type | Description
------------- | ------------- | -------------
**Token** | **string** | The Token HTTP request header
**CategoryId** | **string** | The ID of the category
**Question** | **string** | The question.
**Description** | **string** | The description for the question.
**AnswerSetId** | **string** | The answer_set for the question.
**AllowNa** | **bool** | The flag to enable for disable NA for answer.

### Return type

[**FlexV1InsightsQuestionnairesQuestion**](FlexV1InsightsQuestionnairesQuestion.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInsightsQuestionnairesQuestion

> DeleteInsightsQuestionnairesQuestion(ctx, QuestionIdoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**QuestionId** | **string** | The unique ID of the question

### Other Parameters

Other parameters are passed through a pointer to a DeleteInsightsQuestionnairesQuestionParams struct


Name | Type | Description
------------- | ------------- | -------------
**Token** | **string** | The Token HTTP request header

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInsightsQuestionnairesQuestion

> FlexV1InsightsQuestionnairesQuestion UpdateInsightsQuestionnairesQuestion(ctx, QuestionIdoptional)



To update the question

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**QuestionId** | **string** | The unique ID of the question

### Other Parameters

Other parameters are passed through a pointer to a UpdateInsightsQuestionnairesQuestionParams struct


Name | Type | Description
------------- | ------------- | -------------
**Token** | **string** | The Token HTTP request header
**Question** | **string** | The question.
**Description** | **string** | The description for the question.
**AnswerSetId** | **string** | The answer_set for the question.
**AllowNa** | **bool** | The flag to enable for disable NA for answer.
**CategoryId** | **string** | The ID of the category

### Return type

[**FlexV1InsightsQuestionnairesQuestion**](FlexV1InsightsQuestionnairesQuestion.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

