# InsightsQMQuestionnairesApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInsightsQuestionnaires**](InsightsQMQuestionnairesApi.md#CreateInsightsQuestionnaires) | **Post** /v1/Insights/QM/Questionnaires | 
[**ListInsightsQuestionnaires**](InsightsQMQuestionnairesApi.md#ListInsightsQuestionnaires) | **Get** /v1/Insights/QM/Questionnaires | 



## CreateInsightsQuestionnaires

> FlexV1InsightsQuestionnaires CreateInsightsQuestionnaires(ctx, optional)



To create a Questionnaire

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateInsightsQuestionnairesParams struct


Name | Type | Description
------------- | ------------- | -------------
**Token** | **string** | The Token HTTP request header
**Name** | **string** | The name of this questionnaire
**Description** | **string** | The description of this questionnaire
**Active** | **bool** | The flag to enable or disable questionnaire
**QuestionIds** | **[]string** | The list of questions ids under a questionnaire

### Return type

[**FlexV1InsightsQuestionnaires**](FlexV1InsightsQuestionnaires.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInsightsQuestionnaires

> []FlexV1InsightsQuestionnaires ListInsightsQuestionnaires(ctx, optional)



To get all questionnaires with questions

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListInsightsQuestionnairesParams struct


Name | Type | Description
------------- | ------------- | -------------
**Token** | **string** | The Token HTTP request header
**IncludeInactive** | **bool** | Flag indicating whether to include inactive questionnaires or not
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]FlexV1InsightsQuestionnaires**](FlexV1InsightsQuestionnaires.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

