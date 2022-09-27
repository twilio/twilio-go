# ServicesApi

All URIs are relative to *https://ip-messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateService**](ServicesApi.md#CreateService) | **Post** /v1/Services | 
[**DeleteService**](ServicesApi.md#DeleteService) | **Delete** /v1/Services/{Sid} | 
[**FetchService**](ServicesApi.md#FetchService) | **Get** /v1/Services/{Sid} | 
[**ListService**](ServicesApi.md#ListService) | **Get** /v1/Services | 
[**UpdateService**](ServicesApi.md#UpdateService) | **Post** /v1/Services/{Sid} | 



## CreateService

> IpMessagingV1Service CreateService(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | 

### Return type

[**IpMessagingV1Service**](IpMessagingV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteService

> DeleteService(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceParams struct


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


## FetchService

> IpMessagingV1Service FetchService(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IpMessagingV1Service**](IpMessagingV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListService

> []IpMessagingV1Service ListService(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]IpMessagingV1Service**](IpMessagingV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> IpMessagingV1Service UpdateService(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | 
**DefaultServiceRoleSid** | **string** | 
**DefaultChannelRoleSid** | **string** | 
**DefaultChannelCreatorRoleSid** | **string** | 
**ReadStatusEnabled** | **bool** | 
**ReachabilityEnabled** | **bool** | 
**TypingIndicatorTimeout** | **int** | 
**ConsumptionReportInterval** | **int** | 
**NotificationsNewMessageEnabled** | **bool** | 
**NotificationsNewMessageTemplate** | **string** | 
**NotificationsAddedToChannelEnabled** | **bool** | 
**NotificationsAddedToChannelTemplate** | **string** | 
**NotificationsRemovedFromChannelEnabled** | **bool** | 
**NotificationsRemovedFromChannelTemplate** | **string** | 
**NotificationsInvitedToChannelEnabled** | **bool** | 
**NotificationsInvitedToChannelTemplate** | **string** | 
**PreWebhookUrl** | **string** | 
**PostWebhookUrl** | **string** | 
**WebhookMethod** | **string** | 
**WebhookFilters** | **[]string** | 
**WebhooksOnMessageSendUrl** | **string** | 
**WebhooksOnMessageSendMethod** | **string** | 
**WebhooksOnMessageUpdateUrl** | **string** | 
**WebhooksOnMessageUpdateMethod** | **string** | 
**WebhooksOnMessageRemoveUrl** | **string** | 
**WebhooksOnMessageRemoveMethod** | **string** | 
**WebhooksOnChannelAddUrl** | **string** | 
**WebhooksOnChannelAddMethod** | **string** | 
**WebhooksOnChannelDestroyUrl** | **string** | 
**WebhooksOnChannelDestroyMethod** | **string** | 
**WebhooksOnChannelUpdateUrl** | **string** | 
**WebhooksOnChannelUpdateMethod** | **string** | 
**WebhooksOnMemberAddUrl** | **string** | 
**WebhooksOnMemberAddMethod** | **string** | 
**WebhooksOnMemberRemoveUrl** | **string** | 
**WebhooksOnMemberRemoveMethod** | **string** | 
**WebhooksOnMessageSentUrl** | **string** | 
**WebhooksOnMessageSentMethod** | **string** | 
**WebhooksOnMessageUpdatedUrl** | **string** | 
**WebhooksOnMessageUpdatedMethod** | **string** | 
**WebhooksOnMessageRemovedUrl** | **string** | 
**WebhooksOnMessageRemovedMethod** | **string** | 
**WebhooksOnChannelAddedUrl** | **string** | 
**WebhooksOnChannelAddedMethod** | **string** | 
**WebhooksOnChannelDestroyedUrl** | **string** | 
**WebhooksOnChannelDestroyedMethod** | **string** | 
**WebhooksOnChannelUpdatedUrl** | **string** | 
**WebhooksOnChannelUpdatedMethod** | **string** | 
**WebhooksOnMemberAddedUrl** | **string** | 
**WebhooksOnMemberAddedMethod** | **string** | 
**WebhooksOnMemberRemovedUrl** | **string** | 
**WebhooksOnMemberRemovedMethod** | **string** | 
**LimitsChannelMembers** | **int** | 
**LimitsUserChannels** | **int** | 

### Return type

[**IpMessagingV1Service**](IpMessagingV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

