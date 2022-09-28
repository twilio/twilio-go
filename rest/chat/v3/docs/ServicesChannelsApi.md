# ServicesChannelsApi

All URIs are relative to *https://chat.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateChannel**](ServicesChannelsApi.md#UpdateChannel) | **Post** /v3/Services/{ServiceSid}/Channels/{Sid} | 



## UpdateChannel

> ChatV3Channel UpdateChannel(ctx, ServiceSidSidoptional)



Update a specific Channel.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Sid** | **string** | A 34 character string that uniquely identifies this Channel.

### Other Parameters

Other parameters are passed through a pointer to a UpdateChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
**Type** | **string** | 
**MessagingServiceSid** | **string** | The unique ID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) this channel belongs to.

### Return type

[**ChatV3Channel**](ChatV3Channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

