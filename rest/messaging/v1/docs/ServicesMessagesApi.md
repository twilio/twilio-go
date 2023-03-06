# ServicesMessagesApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessage**](ServicesMessagesApi.md#CreateMessage) | **Post** /v1/Services/{MessagingServiceSid}/Messages | 
[**DeleteMessage**](ServicesMessagesApi.md#DeleteMessage) | **Delete** /v1/Services/{MessagingServiceSid}/Messages/{Sid} | 
[**FetchMessage**](ServicesMessagesApi.md#FetchMessage) | **Get** /v1/Services/{MessagingServiceSid}/Messages/{Sid} | 
[**ListMessage**](ServicesMessagesApi.md#ListMessage) | **Get** /v1/Services/{MessagingServiceSid}/Messages | 
[**UpdateMessage**](ServicesMessagesApi.md#UpdateMessage) | **Post** /v1/Services/{MessagingServiceSid}/Messages/{Sid} | 



## CreateMessage

> MessagingV1Message CreateMessage(ctx, MessagingServiceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessagingServiceSid** | **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) the message belongs to.

### Other Parameters

Other parameters are passed through a pointer to a CreateMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**Recipient** | **[]string** | The list of message recipients, formatted as `channel:number`.
**From** | **string** | A single colon-delimited address.
**Ttl** | **int** | The message's time-to-live in seconds. Can be an integer between 1 and 14,400.
**Template** | **string** | Template for message body.
**TemplateSid** | **string** | Template sid to get template for message body.
**TemplateLanguage** | **string** | Template language which complements template sid to get template for message body.
**TemplateArgs** | [**interface{}**](interface{}.md) | The dictionary of arguments to construct message body from template.
**Body** | **string** | The message body.
**MediaUrl** | **[]string** | The list of media URLs.

### Return type

[**MessagingV1Message**](MessagingV1Message.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMessage

> DeleteMessage(ctx, MessagingServiceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessagingServiceSid** | **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) with the Message resource to delete.
**Sid** | **string** | The SID of the Message resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteMessageParams struct


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


## FetchMessage

> MessagingV1Message FetchMessage(ctx, MessagingServiceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessagingServiceSid** | **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) with the Message resource to fetch.
**Sid** | **string** | The SID of the Message resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchMessageParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MessagingV1Message**](MessagingV1Message.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessage

> []MessagingV1Message ListMessage(ctx, MessagingServiceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessagingServiceSid** | **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) with the Message resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**DateSent** | **string** | The date when the resources to read were created.
**DateSentBefore** | **string** | The date when the resources to read were created.
**DateSentAfter** | **string** | The date when the resources to read were created.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]MessagingV1Message**](MessagingV1Message.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessage

> MessagingV1Message UpdateMessage(ctx, MessagingServiceSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessagingServiceSid** | **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) with the Message resource to update.
**Sid** | **string** | The SID of the Message resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**Body** | **string** | The message body.

### Return type

[**MessagingV1Message**](MessagingV1Message.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

