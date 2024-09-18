# AssistantsFeedbacksApi

All URIs are relative to *https://assistants.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFeedback**](AssistantsFeedbacksApi.md#CreateFeedback) | **Post** /v1/Assistants/{id}/Feedbacks | Create feedback
[**ListFeedback**](AssistantsFeedbacksApi.md#ListFeedback) | **Get** /v1/Assistants/{id}/Feedbacks | List feedbacks



## CreateFeedback

> AssistantsV1Feedback CreateFeedback(ctx, Idoptional)

Create feedback

Create feedback

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The assistant ID.

### Other Parameters

Other parameters are passed through a pointer to a CreateFeedbackParams struct


Name | Type | Description
------------- | ------------- | -------------
**AssistantsV1CreateFeedbackRequest** | [**AssistantsV1CreateFeedbackRequest**](AssistantsV1CreateFeedbackRequest.md) | 

### Return type

[**AssistantsV1Feedback**](AssistantsV1Feedback.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFeedback

> []AssistantsV1Feedback ListFeedback(ctx, Idoptional)

List feedbacks

List feedbacks

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The assistant ID.

### Other Parameters

Other parameters are passed through a pointer to a ListFeedbackParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]AssistantsV1Feedback**](AssistantsV1Feedback.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

