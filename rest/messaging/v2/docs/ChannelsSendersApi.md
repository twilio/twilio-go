# ChannelsSendersApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChannelsSender**](ChannelsSendersApi.md#CreateChannelsSender) | **Post** /v2/Channels/Senders | Create a Sender
[**DeleteChannelsSender**](ChannelsSendersApi.md#DeleteChannelsSender) | **Delete** /v2/Channels/Senders/{Sid} | Delete a Sender
[**FetchChannelsSender**](ChannelsSendersApi.md#FetchChannelsSender) | **Get** /v2/Channels/Senders/{Sid} | Retrieve a Sender
[**ListChannelsSender**](ChannelsSendersApi.md#ListChannelsSender) | **Get** /v2/Channels/Senders | Retrieve a list of Senders
[**UpdateChannelsSender**](ChannelsSendersApi.md#UpdateChannelsSender) | **Post** /v2/Channels/Senders/{Sid} | Update a sender&#39;s information, including &#x60;profile&#x60;, &#x60;webhook&#x60;, and &#x60;configuration&#x60;.



## CreateChannelsSender

> MessagingV2ChannelsSenderResponse CreateChannelsSender(ctx, optional)

Create a Sender

Create a Sender.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateChannelsSenderParams struct


Name | Type | Description
------------- | ------------- | -------------
**MessagingV2Create** | [**MessagingV2Create**](MessagingV2Create.md) | 

### Return type

[**MessagingV2ChannelsSenderResponse**](MessagingV2ChannelsSenderResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChannelsSender

> DeleteChannelsSender(ctx, Sid)

Delete a Sender

(WhatsApp only) Delete a Sender.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the sender.

### Other Parameters

Other parameters are passed through a pointer to a DeleteChannelsSenderParams struct


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


## FetchChannelsSender

> MessagingV2ChannelsSenderResponse FetchChannelsSender(ctx, Sid)

Retrieve a Sender

Retrieve a Sender.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the sender.

### Other Parameters

Other parameters are passed through a pointer to a FetchChannelsSenderParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MessagingV2ChannelsSenderResponse**](MessagingV2ChannelsSenderResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChannelsSender

> []MessagingV2ChannelsSenderResponse ListChannelsSender(ctx, optional)

Retrieve a list of Senders

Retrieve a list of Senders for an account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListChannelsSenderParams struct


Name | Type | Description
------------- | ------------- | -------------
**Channel** | **string** | 
**PageSize** | **int** | The number of items to return per page. For WhatsApp, the default is `20`.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]MessagingV2ChannelsSenderResponse**](MessagingV2ChannelsSenderResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChannelsSender

> MessagingV2ChannelsSenderResponse UpdateChannelsSender(ctx, Sidoptional)

Update a sender's information, including `profile`, `webhook`, and `configuration`.

(WhatsApp only) Update a Sender. You can update a sender's information, including `profile`, `webhook`, and `configuration`. To verify a phone number, set `configuration.verification_code` to the One-time Password (OTP) that you received.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the sender.

### Other Parameters

Other parameters are passed through a pointer to a UpdateChannelsSenderParams struct


Name | Type | Description
------------- | ------------- | -------------
**MessagingV2Update** | [**MessagingV2Update**](MessagingV2Update.md) | 

### Return type

[**MessagingV2ChannelsSenderResponse**](MessagingV2ChannelsSenderResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

