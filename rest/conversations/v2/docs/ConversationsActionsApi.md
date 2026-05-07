# ConversationsActionsApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConversationAction**](ConversationsActionsApi.md#CreateConversationAction) | **Post** /v2/Conversations/{ConversationId}/Actions | Create an Action
[**FetchConversationAction**](ConversationsActionsApi.md#FetchConversationAction) | **Get** /v2/Conversations/{ConversationId}/Actions/{ActionId} | Get Action Status



## CreateConversationAction

> CreateConversationActionResponse CreateConversationAction(ctx, ConversationIdoptional)

Create an Action

Creates an Action within a Conversation. Currently supports SEND_MESSAGE, which sends a message to recipients via the configured channel.  Returns 202 Accepted with the Action in PENDING status. Poll `GET /v2/Conversations/{ConversationId}/Actions/{ActionId}` to check completion. 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateConversationActionParams struct


Name | Type | Description
------------- | ------------- | -------------
**CreateConversationActionRequest** | [**CreateConversationActionRequest**](CreateConversationActionRequest.md) | The action to perform.

### Return type

[**CreateConversationActionResponse**](CreateConversationAction202Response.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConversationAction

> CreateConversationActionResponse FetchConversationAction(ctx, ConversationIdActionId)

Get Action Status

Retrieve the current status of an Action.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationId** | **string** | 
**ActionId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchConversationActionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**CreateConversationActionResponse**](CreateConversationAction202Response.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

