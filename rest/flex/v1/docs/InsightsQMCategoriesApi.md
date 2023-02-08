# InsightsQMCategoriesApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInsightsQuestionnairesCategory**](InsightsQMCategoriesApi.md#CreateInsightsQuestionnairesCategory) | **Post** /v1/Insights/QM/Categories | 
[**DeleteInsightsQuestionnairesCategory**](InsightsQMCategoriesApi.md#DeleteInsightsQuestionnairesCategory) | **Delete** /v1/Insights/QM/Categories/{CategoryId} | 
[**ListInsightsQuestionnairesCategory**](InsightsQMCategoriesApi.md#ListInsightsQuestionnairesCategory) | **Get** /v1/Insights/QM/Categories | 
[**UpdateInsightsQuestionnairesCategory**](InsightsQMCategoriesApi.md#UpdateInsightsQuestionnairesCategory) | **Post** /v1/Insights/QM/Categories/{CategoryId} | 



## CreateInsightsQuestionnairesCategory

> FlexV1InsightsQuestionnairesCategory CreateInsightsQuestionnairesCategory(ctx, optional)



To create a category for Questions

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateInsightsQuestionnairesCategoryParams struct


Name | Type | Description
------------- | ------------- | -------------
**Token** | **string** | The Token HTTP request header
**Name** | **string** | The name of this category.

### Return type

[**FlexV1InsightsQuestionnairesCategory**](FlexV1InsightsQuestionnairesCategory.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInsightsQuestionnairesCategory

> DeleteInsightsQuestionnairesCategory(ctx, CategoryIdoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CategoryId** | **string** | The ID of the category to be deleted

### Other Parameters

Other parameters are passed through a pointer to a DeleteInsightsQuestionnairesCategoryParams struct


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


## ListInsightsQuestionnairesCategory

> []FlexV1InsightsQuestionnairesCategory ListInsightsQuestionnairesCategory(ctx, optional)



To get all the categories

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListInsightsQuestionnairesCategoryParams struct


Name | Type | Description
------------- | ------------- | -------------
**Token** | **string** | The Token HTTP request header
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]FlexV1InsightsQuestionnairesCategory**](FlexV1InsightsQuestionnairesCategory.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInsightsQuestionnairesCategory

> FlexV1InsightsQuestionnairesCategory UpdateInsightsQuestionnairesCategory(ctx, CategoryIdoptional)



To update the category for Questions

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CategoryId** | **string** | The ID of the category to be update

### Other Parameters

Other parameters are passed through a pointer to a UpdateInsightsQuestionnairesCategoryParams struct


Name | Type | Description
------------- | ------------- | -------------
**Token** | **string** | The Token HTTP request header
**Name** | **string** | The name of this category.

### Return type

[**FlexV1InsightsQuestionnairesCategory**](FlexV1InsightsQuestionnairesCategory.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

