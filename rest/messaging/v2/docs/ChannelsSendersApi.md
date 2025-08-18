# ChannelsSendersApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChannelsSender**](ChannelsSendersApi.md#CreateChannelsSender) | **Post** /v2/Channels/Senders | Create a new sender of WhatsApp.
[**DeleteChannelsSender**](ChannelsSendersApi.md#DeleteChannelsSender) | **Delete** /v2/Channels/Senders/{Sid} | Delete a specific sender by its unique identifier.
[**FetchChannelsSender**](ChannelsSendersApi.md#FetchChannelsSender) | **Get** /v2/Channels/Senders/{Sid} | Retrieve details of a specific sender by its unique identifier.
[**ListChannelsSender**](ChannelsSendersApi.md#ListChannelsSender) | **Get** /v2/Channels/Senders | Get a list of Senders for an account.
[**UpdateChannelsSender**](ChannelsSendersApi.md#UpdateChannelsSender) | **Post** /v2/Channels/Senders/{Sid} | Update a specific sender information like OTP Code, Webhook, Profile information.



## CreateChannelsSender

> MessagingV2ChannelsSenderResponse CreateChannelsSender(ctx, optional)

Create a new sender of WhatsApp.

Create a new sender of WhatsApp.

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

Delete a specific sender by its unique identifier.

Delete a specific sender by its unique identifier.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Sender.

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

Retrieve details of a specific sender by its unique identifier.

Retrieve details of a specific sender by its unique identifier.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Sender.

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

Get a list of Senders for an account.

Get a list of Senders for an account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListChannelsSenderParams struct


Name | Type | Description
------------- | ------------- | -------------
**Channel** | **string** | 
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
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

Update a specific sender information like OTP Code, Webhook, Profile information.

Update a specific sender information like OTP Code, Webhook, Profile information.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Sender.

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

