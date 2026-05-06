# ConversationsApi

All URIs are relative to *https://intelligence.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchConversation**](ConversationsApi.md#FetchConversation) | **Get** /v3/Conversations/{id} | Retrieve a Conversation
[**ListConversations**](ConversationsApi.md#ListConversations) | **Get** /v3/Conversations | Retrieve a list of Conversations



## FetchConversation

> FetchConversationResponse FetchConversation(ctx, Id)

Retrieve a Conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The unique identifier of the conversation.

### Other Parameters

Other parameters are passed through a pointer to a FetchConversationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**FetchConversationResponse**](FetchConversation200Response.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConversations

> []ConversationListItem ListConversations(ctx, optional)

Retrieve a list of Conversations

Retrieve a list of Conversations processed by an Intelligence Configuration.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListConversationsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | The maximum number of resources to return
**PageToken** | **string** | Token for pagination
**CreatedAtBefore** | **time.Time** | Filter by Conversations created before this timestamp.
**CreatedAtAfter** | **time.Time** | Filter by Conversations created after this timestamp.
**Status** | **string** | Filter by Conversation status.
**ChannelId** | **string** | Filters Conversations by the underlying channel resource ID, such as a Call ID or Message ID.
**Channels** | **[]string** | Filters Conversations that include one or more of the specified communication channels (`OR` match).
**ConversationConfigurationId** | **string** | The configuration `id` used to generate the Conversation.
**IntelligenceConfigurationIds** | **[]string** | Filters Conversations activated by one or more of the specified Intelligence Configuration IDs (`OR` match).
**OperatorIds** | **[]string** | Filters Conversations to those where at least one of the specified Language Operators was executed (`OR` match).
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ConversationListItem**](ConversationListItem.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

