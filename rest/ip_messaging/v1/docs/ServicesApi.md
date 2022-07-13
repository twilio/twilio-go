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
**ConsumptionReportInterval** | **int** | 
**DefaultChannelCreatorRoleSid** | **string** | 
**DefaultChannelRoleSid** | **string** | 
**DefaultServiceRoleSid** | **string** | 
**FriendlyName** | **string** | 
**LimitsChannelMembers** | **int** | 
**LimitsUserChannels** | **int** | 
**NotificationsAddedToChannelEnabled** | **bool** | 
**NotificationsAddedToChannelTemplate** | **string** | 
**NotificationsInvitedToChannelEnabled** | **bool** | 
**NotificationsInvitedToChannelTemplate** | **string** | 
**NotificationsNewMessageEnabled** | **bool** | 
**NotificationsNewMessageTemplate** | **string** | 
**NotificationsRemovedFromChannelEnabled** | **bool** | 
**NotificationsRemovedFromChannelTemplate** | **string** | 
**PostWebhookUrl** | **string** | 
**PreWebhookUrl** | **string** | 
**ReachabilityEnabled** | **bool** | 
**ReadStatusEnabled** | **bool** | 
**TypingIndicatorTimeout** | **int** | 
**WebhookFilters** | **[]string** | 
**WebhookMethod** | **string** | 
**WebhooksOnChannelAddMethod** | **string** | 
**WebhooksOnChannelAddUrl** | **string** | 
**WebhooksOnChannelAddedMethod** | **string** | 
**WebhooksOnChannelAddedUrl** | **string** | 
**WebhooksOnChannelDestroyMethod** | **string** | 
**WebhooksOnChannelDestroyUrl** | **string** | 
**WebhooksOnChannelDestroyedMethod** | **string** | 
**WebhooksOnChannelDestroyedUrl** | **string** | 
**WebhooksOnChannelUpdateMethod** | **string** | 
**WebhooksOnChannelUpdateUrl** | **string** | 
**WebhooksOnChannelUpdatedMethod** | **string** | 
**WebhooksOnChannelUpdatedUrl** | **string** | 
**WebhooksOnMemberAddMethod** | **string** | 
**WebhooksOnMemberAddUrl** | **string** | 
**WebhooksOnMemberAddedMethod** | **string** | 
**WebhooksOnMemberAddedUrl** | **string** | 
**WebhooksOnMemberRemoveMethod** | **string** | 
**WebhooksOnMemberRemoveUrl** | **string** | 
**WebhooksOnMemberRemovedMethod** | **string** | 
**WebhooksOnMemberRemovedUrl** | **string** | 
**WebhooksOnMessageRemoveMethod** | **string** | 
**WebhooksOnMessageRemoveUrl** | **string** | 
**WebhooksOnMessageRemovedMethod** | **string** | 
**WebhooksOnMessageRemovedUrl** | **string** | 
**WebhooksOnMessageSendMethod** | **string** | 
**WebhooksOnMessageSendUrl** | **string** | 
**WebhooksOnMessageSentMethod** | **string** | 
**WebhooksOnMessageSentUrl** | **string** | 
**WebhooksOnMessageUpdateMethod** | **string** | 
**WebhooksOnMessageUpdateUrl** | **string** | 
**WebhooksOnMessageUpdatedMethod** | **string** | 
**WebhooksOnMessageUpdatedUrl** | **string** | 

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

