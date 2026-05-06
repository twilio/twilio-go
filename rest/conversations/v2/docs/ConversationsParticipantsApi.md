# ConversationsParticipantsApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateParticipantInConversation**](ConversationsParticipantsApi.md#CreateParticipantInConversation) | **Post** /v2/Conversations/{ConversationSid}/Participants | Create Participant
[**FetchParticipant**](ConversationsParticipantsApi.md#FetchParticipant) | **Get** /v2/Conversations/{ConversationSid}/Participants/{Sid} | Fetch Participant
[**ListParticipantByConversation**](ConversationsParticipantsApi.md#ListParticipantByConversation) | **Get** /v2/Conversations/{ConversationSid}/Participants | List Participants
[**UpdateParticipantInConversation**](ConversationsParticipantsApi.md#UpdateParticipantInConversation) | **Put** /v2/Conversations/{ConversationSid}/Participants/{Sid} | Update a Participant



## CreateParticipantInConversation

> ListParticipantByConversationResponseParticipants CreateParticipantInConversation(ctx, ConversationSidoptional)

Create Participant

Create a Participant.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateParticipantInConversationParams struct


Name | Type | Description
------------- | ------------- | -------------
**CreateParticipantInConversationRequest** | [**CreateParticipantInConversationRequest**](CreateParticipantInConversationRequest.md) | 

### Return type

[**ListParticipantByConversationResponseParticipants**](ListParticipantByConversation200ResponseParticipants.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchParticipant

> ListParticipantByConversationResponseParticipants FetchParticipant(ctx, ConversationSidSid)

Fetch Participant

Retrieve a Participant.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchParticipantParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ListParticipantByConversationResponseParticipants**](ListParticipantByConversation200ResponseParticipants.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListParticipantByConversation

> []ListParticipantByConversationResponseParticipants ListParticipantByConversation(ctx, ConversationSidoptional)

List Participants

Retrieve a list of Participants in a Conversation.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListParticipantByConversationParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | Maximum number of items to return
**PageToken** | **string** | Page token for pagination
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ListParticipantByConversationResponseParticipants**](ListParticipantByConversationResponseParticipants.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateParticipantInConversation

> ListParticipantByConversationResponseParticipants UpdateParticipantInConversation(ctx, ConversationSidSidoptional)

Update a Participant

Update an existing Participant

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateParticipantInConversationParams struct


Name | Type | Description
------------- | ------------- | -------------
**UpdateParticipantInConversationRequest** | [**UpdateParticipantInConversationRequest**](UpdateParticipantInConversationRequest.md) | 

### Return type

[**ListParticipantByConversationResponseParticipants**](ListParticipantByConversation200ResponseParticipants.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

