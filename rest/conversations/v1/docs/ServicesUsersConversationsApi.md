# ServicesUsersConversationsApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServiceUserConversation**](ServicesUsersConversationsApi.md#DeleteServiceUserConversation) | **Delete** /v1/Services/{ChatServiceSid}/Users/{UserSid}/Conversations/{ConversationSid} | Delete a specific User Conversation.
[**FetchServiceUserConversation**](ServicesUsersConversationsApi.md#FetchServiceUserConversation) | **Get** /v1/Services/{ChatServiceSid}/Users/{UserSid}/Conversations/{ConversationSid} | Fetch a specific User Conversation.
[**ListServiceUserConversation**](ServicesUsersConversationsApi.md#ListServiceUserConversation) | **Get** /v1/Services/{ChatServiceSid}/Users/{UserSid}/Conversations | Retrieve a list of all User Conversations for the User.
[**UpdateServiceUserConversation**](ServicesUsersConversationsApi.md#UpdateServiceUserConversation) | **Post** /v1/Services/{ChatServiceSid}/Users/{UserSid}/Conversations/{ConversationSid} | Update a specific User Conversation.



## DeleteServiceUserConversation

> DeleteServiceUserConversation(ctx, ChatServiceSidUserSidConversationSid)

Delete a specific User Conversation.

Delete a specific User Conversation.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.
**UserSid** | **string** | The unique SID identifier of the [User resource](https://www.twilio.com/docs/conversations/api/user-resource). This value can be either the `sid` or the `identity` of the User resource.
**ConversationSid** | **string** | The unique SID identifier of the Conversation. This value can be either the `sid` or the `unique_name` of the [Conversation resource](https://www.twilio.com/docs/conversations/api/conversation-resource).

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceUserConversationParams struct


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


## FetchServiceUserConversation

> ConversationsV1ServiceUserConversation FetchServiceUserConversation(ctx, ChatServiceSidUserSidConversationSid)

Fetch a specific User Conversation.

Fetch a specific User Conversation.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.
**UserSid** | **string** | The unique SID identifier of the [User resource](https://www.twilio.com/docs/conversations/api/user-resource). This value can be either the `sid` or the `identity` of the User resource.
**ConversationSid** | **string** | The unique SID identifier of the Conversation. This value can be either the `sid` or the `unique_name` of the [Conversation resource](https://www.twilio.com/docs/conversations/api/conversation-resource).

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceUserConversationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ServiceUserConversation**](ConversationsV1ServiceUserConversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceUserConversation

> []ConversationsV1ServiceUserConversation ListServiceUserConversation(ctx, ChatServiceSidUserSidoptional)

Retrieve a list of all User Conversations for the User.

Retrieve a list of all User Conversations for the User.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.
**UserSid** | **string** | The unique SID identifier of the [User resource](https://www.twilio.com/docs/conversations/api/user-resource). This value can be either the `sid` or the `identity` of the User resource.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceUserConversationParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 50.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ConversationsV1ServiceUserConversation**](ConversationsV1ServiceUserConversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceUserConversation

> ConversationsV1ServiceUserConversation UpdateServiceUserConversation(ctx, ChatServiceSidUserSidConversationSidoptional)

Update a specific User Conversation.

Update a specific User Conversation.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.
**UserSid** | **string** | The unique SID identifier of the [User resource](https://www.twilio.com/docs/conversations/api/user-resource). This value can be either the `sid` or the `identity` of the User resource.
**ConversationSid** | **string** | The unique SID identifier of the Conversation. This value can be either the `sid` or the `unique_name` of the [Conversation resource](https://www.twilio.com/docs/conversations/api/conversation-resource).

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceUserConversationParams struct


Name | Type | Description
------------- | ------------- | -------------
**NotificationLevel** | [**string**](string.md) | 
**LastReadTimestamp** | **time.Time** | The date of the last message read in conversation by the user, given in ISO 8601 format.
**LastReadMessageIndex** | **int** | The index of the last Message in the Conversation that the Participant has read.

### Return type

[**ConversationsV1ServiceUserConversation**](ConversationsV1ServiceUserConversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

