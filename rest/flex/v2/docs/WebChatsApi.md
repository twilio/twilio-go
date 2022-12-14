# WebChatsApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebChannel**](WebChatsApi.md#CreateWebChannel) | **Post** /v2/WebChats | 



## CreateWebChannel

> FlexV2WebChannel CreateWebChannel(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateWebChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**AddressSid** | **string** | The SID of the Conversations Address. See [Address Configuration Resource](https://www.twilio.com/docs/conversations/api/address-configuration-resource) for configuration details. When a conversation is created on the Flex backend, the callback URL will be set to the corresponding Studio Flow SID or webhook URL in your address configuration.
**ChatFriendlyName** | **string** | The Conversation's friendly name. See the [Conversation resource](https://www.twilio.com/docs/conversations/api/conversation-resource) for an example.
**CustomerFriendlyName** | **string** | The Conversation participant's friendly name. See the [Conversation Participant Resource](https://www.twilio.com/docs/conversations/api/conversation-participant-resource) for an example.
**PreEngagementData** | **string** | The pre-engagement data.

### Return type

[**FlexV2WebChannel**](FlexV2WebChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

