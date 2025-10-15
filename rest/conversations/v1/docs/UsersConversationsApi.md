# UsersConversationsApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUserConversation**](UsersConversationsApi.md#DeleteUserConversation) | **Delete** /v1/Users/{UserSid}/Conversations/{ConversationSid} | Delete a specific User Conversation.
[**FetchUserConversation**](UsersConversationsApi.md#FetchUserConversation) | **Get** /v1/Users/{UserSid}/Conversations/{ConversationSid} | Fetch a specific User Conversation.
[**ListUserConversation**](UsersConversationsApi.md#ListUserConversation) | **Get** /v1/Users/{UserSid}/Conversations | Retrieve a list of all User Conversations for the User.
[**UpdateUserConversation**](UsersConversationsApi.md#UpdateUserConversation) | **Post** /v1/Users/{UserSid}/Conversations/{ConversationSid} | Update a specific User Conversation.



## DeleteUserConversation

> DeleteUserConversation(ctx, UserSidConversationSid)

Delete a specific User Conversation.

Delete a specific User Conversation.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**UserSid** | **string** | The unique SID identifier of the [User resource](https://www.twilio.com/docs/conversations/api/user-resource). This value can be either the `sid` or the `identity` of the User resource.
**ConversationSid** | **string** | The unique SID identifier of the Conversation. This value can be either the `sid` or the `unique_name` of the [Conversation resource](https://www.twilio.com/docs/conversations/api/conversation-resource).

### Other Parameters

Other parameters are passed through a pointer to a DeleteUserConversationParams struct


Name | Type | Description
------------- | ------------- | -------------

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


## FetchUserConversation

> ConversationsV1UserConversation FetchUserConversation(ctx, UserSidConversationSid)

Fetch a specific User Conversation.

Fetch a specific User Conversation.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**UserSid** | **string** | The unique SID identifier of the [User resource](https://www.twilio.com/docs/conversations/api/user-resource). This value can be either the `sid` or the `identity` of the User resource.
**ConversationSid** | **string** | The unique SID identifier of the Conversation. This value can be either the `sid` or the `unique_name` of the [Conversation resource](https://www.twilio.com/docs/conversations/api/conversation-resource).

### Other Parameters

Other parameters are passed through a pointer to a FetchUserConversationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1UserConversation**](ConversationsV1UserConversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserConversation

> []ConversationsV1UserConversation ListUserConversation(ctx, UserSidoptional)

Retrieve a list of all User Conversations for the User.

Retrieve a list of all User Conversations for the User.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**UserSid** | **string** | The unique SID identifier of the [User resource](https://www.twilio.com/docs/conversations/api/user-resource). This value can be either the `sid` or the `identity` of the User resource.

### Other Parameters

Other parameters are passed through a pointer to a ListUserConversationParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 50.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ConversationsV1UserConversation**](ConversationsV1UserConversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserConversation

> ConversationsV1UserConversation UpdateUserConversation(ctx, UserSidConversationSidoptional)

Update a specific User Conversation.

Update a specific User Conversation.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**UserSid** | **string** | The unique SID identifier of the [User resource](https://www.twilio.com/docs/conversations/api/user-resource). This value can be either the `sid` or the `identity` of the User resource.
**ConversationSid** | **string** | The unique SID identifier of the Conversation. This value can be either the `sid` or the `unique_name` of the [Conversation resource](https://www.twilio.com/docs/conversations/api/conversation-resource).

### Other Parameters

Other parameters are passed through a pointer to a UpdateUserConversationParams struct


Name | Type | Description
------------- | ------------- | -------------
**NotificationLevel** | [**string**](string.md) | 
**LastReadTimestamp** | **time.Time** | The date of the last message read in conversation by the user, given in ISO 8601 format.
**LastReadMessageIndex** | **int** | The index of the last Message in the Conversation that the Participant has read.

### Return type

[**ConversationsV1UserConversation**](ConversationsV1UserConversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

