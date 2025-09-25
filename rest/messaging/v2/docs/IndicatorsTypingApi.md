# IndicatorsTypingApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTypingIndicator**](IndicatorsTypingApi.md#CreateTypingIndicator) | **Post** /v2/Indicators/Typing.json | Send a typing indicator



## CreateTypingIndicator

> CreateTypingIndicatorResponse CreateTypingIndicator(ctx, optional)

Send a typing indicator

Send a typing indicator to notify the recipient that you are composing a message. Currently supported for whatsapp channel only. For WhatsApp, `messageId` is required. 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateTypingIndicatorParams struct


Name | Type | Description
------------- | ------------- | -------------
**Channel** | **string** | Shared channel identifier
**MessageId** | **string** | Message SID that identifies the conversation thread for the typing indicator. Must be a valid Twilio Message SID (SM*) or Media SID (MM*) from an existing WhatsApp conversation. 

### Return type

[**CreateTypingIndicatorResponse**](CreateTypingIndicator200Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

