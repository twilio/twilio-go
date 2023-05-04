# InsightsQualityManagementQuestionsApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInsightsQuestionnairesQuestion**](InsightsQualityManagementQuestionsApi.md#CreateInsightsQuestionnairesQuestion) | **Post** /v1/Insights/QualityManagement/Questions | 
[**DeleteInsightsQuestionnairesQuestion**](InsightsQualityManagementQuestionsApi.md#DeleteInsightsQuestionnairesQuestion) | **Delete** /v1/Insights/QualityManagement/Questions/{QuestionSid} | 
[**ListInsightsQuestionnairesQuestion**](InsightsQualityManagementQuestionsApi.md#ListInsightsQuestionnairesQuestion) | **Get** /v1/Insights/QualityManagement/Questions | 
[**UpdateInsightsQuestionnairesQuestion**](InsightsQualityManagementQuestionsApi.md#UpdateInsightsQuestionnairesQuestion) | **Post** /v1/Insights/QualityManagement/Questions/{QuestionSid} | 



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
**CategorySid** | **string** | The SID of the category
**Question** | **string** | The question.
**AnswerSetId** | **string** | The answer_set for the question.
**AllowNa** | **bool** | The flag to enable for disable NA for answer.
**Description** | **string** | The description for the question.

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

> DeleteInsightsQuestionnairesQuestion(ctx, QuestionSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**QuestionSid** | **string** | The SID of the question

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


## ListInsightsQuestionnairesQuestion

> []FlexV1InsightsQuestionnairesQuestion ListInsightsQuestionnairesQuestion(ctx, optional)



To get all the question for the given categories

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListInsightsQuestionnairesQuestionParams struct


Name | Type | Description
------------- | ------------- | -------------
**Token** | **string** | The Token HTTP request header
**CategorySid** | **[]string** | The list of category SIDs
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]FlexV1InsightsQuestionnairesQuestion**](FlexV1InsightsQuestionnairesQuestion.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInsightsQuestionnairesQuestion

> FlexV1InsightsQuestionnairesQuestion UpdateInsightsQuestionnairesQuestion(ctx, QuestionSidoptional)



To update the question

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**QuestionSid** | **string** | The SID of the question

### Other Parameters

Other parameters are passed through a pointer to a UpdateInsightsQuestionnairesQuestionParams struct


Name | Type | Description
------------- | ------------- | -------------
**Token** | **string** | The Token HTTP request header
**AllowNa** | **bool** | The flag to enable for disable NA for answer.
**CategorySid** | **string** | The SID of the category
**Question** | **string** | The question.
**Description** | **string** | The description for the question.
**AnswerSetId** | **string** | The answer_set for the question.

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

