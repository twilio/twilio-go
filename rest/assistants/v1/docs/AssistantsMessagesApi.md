# AssistantsMessagesApi

All URIs are relative to *https://assistants.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessage**](AssistantsMessagesApi.md#CreateMessage) | **Post** /v1/Assistants/{id}/Messages | Send a message to the assistant



## CreateMessage

> AssistantsV1AssistantSendMessageResponse CreateMessage(ctx, Idoptional)

Send a message to the assistant

send a message

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | the Assistant ID.

### Other Parameters

Other parameters are passed through a pointer to a CreateMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**AssistantsV1AssistantSendMessageRequest** | [**AssistantsV1AssistantSendMessageRequest**](AssistantsV1AssistantSendMessageRequest.md) | 

### Return type

[**AssistantsV1AssistantSendMessageResponse**](AssistantsV1AssistantSendMessageResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

