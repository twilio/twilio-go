# ServicesStreamsMessagesApi

All URIs are relative to *https://sync.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStreamMessage**](ServicesStreamsMessagesApi.md#CreateStreamMessage) | **Post** /v1/Services/{ServiceSid}/Streams/{StreamSid}/Messages | 



## CreateStreamMessage

> SyncV1StreamMessage CreateStreamMessage(ctx, ServiceSidStreamSidoptional)



Create a new Stream Message.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) to create the new Stream Message in.
**StreamSid** | **string** | The SID of the Sync Stream to create the new Stream Message resource for.

### Other Parameters

Other parameters are passed through a pointer to a CreateStreamMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**Data** | [**interface{}**](interface{}.md) | A JSON string that represents an arbitrary, schema-less object that makes up the Stream Message body. Can be up to 4 KiB in length.

### Return type

[**SyncV1StreamMessage**](SyncV1StreamMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

