# ServicesChannelSendersApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChannelSender**](ServicesChannelSendersApi.md#CreateChannelSender) | **Post** /v1/Services/{MessagingServiceSid}/ChannelSenders | 
[**DeleteChannelSender**](ServicesChannelSendersApi.md#DeleteChannelSender) | **Delete** /v1/Services/{MessagingServiceSid}/ChannelSenders/{Sid} | 
[**FetchChannelSender**](ServicesChannelSendersApi.md#FetchChannelSender) | **Get** /v1/Services/{MessagingServiceSid}/ChannelSenders/{Sid} | 
[**ListChannelSender**](ServicesChannelSendersApi.md#ListChannelSender) | **Get** /v1/Services/{MessagingServiceSid}/ChannelSenders | 



## CreateChannelSender

> MessagingV1ChannelSender CreateChannelSender(ctx, MessagingServiceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessagingServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateChannelSenderParams struct


Name | Type | Description
------------- | ------------- | -------------
**Sid** | **string** | The SID of the Channel Sender being added to the Service.

### Return type

[**MessagingV1ChannelSender**](MessagingV1ChannelSender.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChannelSender

> DeleteChannelSender(ctx, MessagingServiceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessagingServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the resource from.
**Sid** | **string** | The SID of the Channel Sender resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteChannelSenderParams struct


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


## FetchChannelSender

> MessagingV1ChannelSender FetchChannelSender(ctx, MessagingServiceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessagingServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the resource from.
**Sid** | **string** | The SID of the ChannelSender resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchChannelSenderParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MessagingV1ChannelSender**](MessagingV1ChannelSender.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChannelSender

> []MessagingV1ChannelSender ListChannelSender(ctx, MessagingServiceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessagingServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListChannelSenderParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]MessagingV1ChannelSender**](MessagingV1ChannelSender.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

