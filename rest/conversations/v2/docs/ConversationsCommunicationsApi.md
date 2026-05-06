# ConversationsCommunicationsApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCommunicationInConversation**](ConversationsCommunicationsApi.md#CreateCommunicationInConversation) | **Post** /v2/Conversations/{ConversationSid}/Communications | Create Communication
[**FetchCommunication**](ConversationsCommunicationsApi.md#FetchCommunication) | **Get** /v2/Conversations/{ConversationSid}/Communications/{Sid} | Fetch Communication
[**ListCommunicationByConversation**](ConversationsCommunicationsApi.md#ListCommunicationByConversation) | **Get** /v2/Conversations/{ConversationSid}/Communications | List Communications



## CreateCommunicationInConversation

> ListCommunicationByConversationResponseCommunications CreateCommunicationInConversation(ctx, ConversationSidoptional)

Create Communication

Create a Communication.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateCommunicationInConversationParams struct


Name | Type | Description
------------- | ------------- | -------------
**CreateCommunicationInConversationRequest** | [**CreateCommunicationInConversationRequest**](CreateCommunicationInConversationRequest.md) | 

### Return type

[**ListCommunicationByConversationResponseCommunications**](ListCommunicationByConversation200ResponseCommunications.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCommunication

> ListCommunicationByConversationResponseCommunications FetchCommunication(ctx, ConversationSidSid)

Fetch Communication

Retrieve a Communication.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchCommunicationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ListCommunicationByConversationResponseCommunications**](ListCommunicationByConversation200ResponseCommunications.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommunicationByConversation

> []ListCommunicationByConversationResponseCommunications ListCommunicationByConversation(ctx, ConversationSidoptional)

List Communications

Retrieve a list of Communications in a Conversation.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListCommunicationByConversationParams struct


Name | Type | Description
------------- | ------------- | -------------
**ChannelId** | **string** | Resource identifier to filter communications
**PageSize** | **int** | Maximum number of items to return
**PageToken** | **string** | Page token for pagination
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ListCommunicationByConversationResponseCommunications**](ListCommunicationByConversationResponseCommunications.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

