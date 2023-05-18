# InsightsQualityManagementQuestionnairesApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInsightsQuestionnaires**](InsightsQualityManagementQuestionnairesApi.md#CreateInsightsQuestionnaires) | **Post** /v1/Insights/QualityManagement/Questionnaires | 
[**DeleteInsightsQuestionnaires**](InsightsQualityManagementQuestionnairesApi.md#DeleteInsightsQuestionnaires) | **Delete** /v1/Insights/QualityManagement/Questionnaires/{QuestionnaireSid} | 
[**FetchInsightsQuestionnaires**](InsightsQualityManagementQuestionnairesApi.md#FetchInsightsQuestionnaires) | **Get** /v1/Insights/QualityManagement/Questionnaires/{QuestionnaireSid} | 
[**ListInsightsQuestionnaires**](InsightsQualityManagementQuestionnairesApi.md#ListInsightsQuestionnaires) | **Get** /v1/Insights/QualityManagement/Questionnaires | 
[**UpdateInsightsQuestionnaires**](InsightsQualityManagementQuestionnairesApi.md#UpdateInsightsQuestionnaires) | **Post** /v1/Insights/QualityManagement/Questionnaires/{QuestionnaireSid} | 



## CreateInsightsQuestionnaires

> FlexV1InsightsQuestionnaires CreateInsightsQuestionnaires(ctx, optional)



To create a Questionnaire

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateInsightsQuestionnairesParams struct


Name | Type | Description
------------- | ------------- | -------------
**Authorization** | **string** | The Authorization HTTP request header
**Name** | **string** | The name of this questionnaire
**Description** | **string** | The description of this questionnaire
**Active** | **bool** | The flag to enable or disable questionnaire
**QuestionSids** | **[]string** | The list of questions sids under a questionnaire

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


## DeleteInsightsQuestionnaires

> DeleteInsightsQuestionnaires(ctx, QuestionnaireSidoptional)



To delete the questionnaire

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**QuestionnaireSid** | **string** | The SID of the questionnaire

### Other Parameters

Other parameters are passed through a pointer to a DeleteInsightsQuestionnairesParams struct


Name | Type | Description
------------- | ------------- | -------------
**Authorization** | **string** | The Authorization HTTP request header

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


## FetchInsightsQuestionnaires

> FlexV1InsightsQuestionnaires FetchInsightsQuestionnaires(ctx, QuestionnaireSidoptional)



To get the Questionnaire Detail

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**QuestionnaireSid** | **string** | The SID of the questionnaire

### Other Parameters

Other parameters are passed through a pointer to a FetchInsightsQuestionnairesParams struct


Name | Type | Description
------------- | ------------- | -------------
**Authorization** | **string** | The Authorization HTTP request header

### Return type

[**FlexV1InsightsQuestionnaires**](FlexV1InsightsQuestionnaires.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
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
**Authorization** | **string** | The Authorization HTTP request header
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


## UpdateInsightsQuestionnaires

> FlexV1InsightsQuestionnaires UpdateInsightsQuestionnaires(ctx, QuestionnaireSidoptional)



To update the questionnaire

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**QuestionnaireSid** | **string** | The SID of the questionnaire

### Other Parameters

Other parameters are passed through a pointer to a UpdateInsightsQuestionnairesParams struct


Name | Type | Description
------------- | ------------- | -------------
**Authorization** | **string** | The Authorization HTTP request header
**Active** | **bool** | The flag to enable or disable questionnaire
**Name** | **string** | The name of this questionnaire
**Description** | **string** | The description of this questionnaire
**QuestionSids** | **[]string** | The list of questions sids under a questionnaire

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

