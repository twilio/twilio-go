# MessagesApi

All URIs are relative to *https://preview.messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessages**](MessagesApi.md#CreateMessages) | **Post** /v1/Messages | 



## CreateMessages

> MessagingV1CreateMessagesResult CreateMessages(ctx, optional)



Send messages to multiple recipients

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateMessagesParams struct


Name | Type | Description
------------- | ------------- | -------------
**CreateMessagesRequest** | [**CreateMessagesRequest**](CreateMessagesRequest.md) | 

### Return type

[**MessagingV1CreateMessagesResult**](MessagingV1CreateMessagesResult.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

