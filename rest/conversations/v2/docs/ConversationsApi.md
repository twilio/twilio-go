# ConversationsApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConversationWithConfig**](ConversationsApi.md#CreateConversationWithConfig) | **Post** /v2/Conversations | Create a new Conversation
[**DeleteConversationAsync**](ConversationsApi.md#DeleteConversationAsync) | **Delete** /v2/Conversations/{Sid} | Delete a Conversation (async)
[**FetchConversation**](ConversationsApi.md#FetchConversation) | **Get** /v2/Conversations/{Sid} | Fetch Conversation
[**ListConversationByAccount**](ConversationsApi.md#ListConversationByAccount) | **Get** /v2/Conversations | List Conversations
[**PatchConversationById**](ConversationsApi.md#PatchConversationById) | **Patch** /v2/Conversations/{Sid} | Partially Update a Conversation
[**UpdateConversationById**](ConversationsApi.md#UpdateConversationById) | **Put** /v2/Conversations/{Sid} | Update a Conversation



## CreateConversationWithConfig

> ListConversationByAccountResponseConversations CreateConversationWithConfig(ctx, optional)

Create a new Conversation

Create a new conversation

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateConversationWithConfigParams struct


Name | Type | Description
------------- | ------------- | -------------
**CreateConversationWithConfigRequest** | [**CreateConversationWithConfigRequest**](CreateConversationWithConfigRequest.md) | 

### Return type

[**ListConversationByAccountResponseConversations**](ListConversationByAccount200ResponseConversations.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConversationAsync

> CreateConfigurationResponse DeleteConversationAsync(ctx, Sidoptional)

Delete a Conversation (async)

Asynchronously delete a conversation and all associated data. Returns 202 Accepted with an Operation-Id for status tracking via GET /v2/ControlPlane/Operations/{operationId}. 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteConversationAsyncParams struct


Name | Type | Description
------------- | ------------- | -------------
**IdempotencyKey** | **string** | Client-generated UUID key to ensure idempotent behavior. Submitting the same key returns the original response without creating a duplicate operation. Keys are scoped to account + region with a 24-hour TTL.

### Return type

[**CreateConfigurationResponse**](CreateConfiguration202Response.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConversation

> ListConversationByAccountResponseConversations FetchConversation(ctx, Sid)

Fetch Conversation

Retrieve a Conversation.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchConversationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ListConversationByAccountResponseConversations**](ListConversationByAccount200ResponseConversations.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConversationByAccount

> []ListConversationByAccountResponseConversations ListConversationByAccount(ctx, optional)

List Conversations

Retrieve a list of Conversations.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListConversationByAccountParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | **[]string** | Filters for specific statuses
**ChannelId** | **string** | The resource identifier (such as callSid or messageSid) to filter conversations.
**PageSize** | **int** | Maximum number of items to return in a single response
**PageToken** | **string** | A URL-safe, base64-encoded token representing the page of results to return
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ListConversationByAccountResponseConversations**](ListConversationByAccountResponseConversations.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchConversationById

> ListConversationByAccountResponseConversations PatchConversationById(ctx, Sidoptional)

Partially Update a Conversation

Partially update the details of an existing Conversation.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a PatchConversationByIdParams struct


Name | Type | Description
------------- | ------------- | -------------
**PatchConversationByIdRequest** | [**PatchConversationByIdRequest**](PatchConversationByIdRequest.md) | The conversation fields to update

### Return type

[**ListConversationByAccountResponseConversations**](ListConversationByAccount200ResponseConversations.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConversationById

> ListConversationByAccountResponseConversations UpdateConversationById(ctx, Sidoptional)

Update a Conversation

Update an existing conversation

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateConversationByIdParams struct


Name | Type | Description
------------- | ------------- | -------------
**UpdateConversationByIdRequest** | [**UpdateConversationByIdRequest**](UpdateConversationByIdRequest.md) | The conversation to update

### Return type

[**ListConversationByAccountResponseConversations**](ListConversationByAccount200ResponseConversations.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

