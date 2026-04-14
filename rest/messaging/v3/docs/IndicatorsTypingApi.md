# IndicatorsTypingApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateV3TypingIndicator**](IndicatorsTypingApi.md#CreateV3TypingIndicator) | **Post** /v3/Indicators/Typing.json | Send a typing indicator



## CreateV3TypingIndicator

> CreateV3TypingIndicatorResponse CreateV3TypingIndicator(ctx, optional)

Send a typing indicator

Send a typing indicator to notify the recipient that you are composing a message. Supported channels: WhatsApp, Apple Messages for Business. The request body varies by channel — use the `channel` field as the discriminator. 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateV3TypingIndicatorParams struct


Name | Type | Description
------------- | ------------- | -------------
**TypingIndicatorRequest** | [**TypingIndicatorRequest**](TypingIndicatorRequest.md) | 

### Return type

[**CreateV3TypingIndicatorResponse**](CreateV3TypingIndicator200Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

