# ParticipantConversationsApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListParticipantConversation**](ParticipantConversationsApi.md#ListParticipantConversation) | **Get** /v1/ParticipantConversations | 



## ListParticipantConversation

> []ConversationsV1ParticipantConversation ListParticipantConversation(ctx, optional)



Retrieve a list of all Conversations that this Participant belongs to by identity or by address. Only one parameter should be specified.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListParticipantConversationParams struct


Name | Type | Description
------------- | ------------- | -------------
**Identity** | **string** | A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversations SDK to communicate. Limited to 256 characters.
**Address** | **string** | A unique string identifier for the conversation participant who's not a Conversation User. This parameter could be found in messaging_binding.address field of Participant resource. It should be url-encoded.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ConversationsV1ParticipantConversation**](ConversationsV1ParticipantConversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

