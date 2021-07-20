# ServicesApi

All URIs are relative to *https://ip-messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateService**](ServicesApi.md#CreateService) | **Post** /v2/Services | 
[**DeleteService**](ServicesApi.md#DeleteService) | **Delete** /v2/Services/{Sid} | 
[**FetchService**](ServicesApi.md#FetchService) | **Get** /v2/Services/{Sid} | 
[**ListService**](ServicesApi.md#ListService) | **Get** /v2/Services | 
[**UpdateService**](ServicesApi.md#UpdateService) | **Post** /v2/Services/{Sid} | 



## CreateService

> IpMessagingV2Service CreateService(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | 

### Return type

[**IpMessagingV2Service**](IpMessagingV2Service.md)

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

> IpMessagingV2Service FetchService(ctx, Sid)



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

[**IpMessagingV2Service**](IpMessagingV2Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListService

> ListServiceResponse ListService(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**ListServiceResponse**](ListServiceResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> IpMessagingV2Service UpdateService(ctx, Sidoptional)



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
**MediaCompatibilityMessage** | **string** | 
**NotificationsAddedToChannelEnabled** | **bool** | 
**NotificationsAddedToChannelSound** | **string** | 
**NotificationsAddedToChannelTemplate** | **string** | 
**NotificationsInvitedToChannelEnabled** | **bool** | 
**NotificationsInvitedToChannelSound** | **string** | 
**NotificationsInvitedToChannelTemplate** | **string** | 
**NotificationsLogEnabled** | **bool** | 
**NotificationsNewMessageBadgeCountEnabled** | **bool** | 
**NotificationsNewMessageEnabled** | **bool** | 
**NotificationsNewMessageSound** | **string** | 
**NotificationsNewMessageTemplate** | **string** | 
**NotificationsRemovedFromChannelEnabled** | **bool** | 
**NotificationsRemovedFromChannelSound** | **string** | 
**NotificationsRemovedFromChannelTemplate** | **string** | 
**PostWebhookRetryCount** | **int** | 
**PostWebhookUrl** | **string** | 
**PreWebhookRetryCount** | **int** | 
**PreWebhookUrl** | **string** | 
**ReachabilityEnabled** | **bool** | 
**ReadStatusEnabled** | **bool** | 
**TypingIndicatorTimeout** | **int** | 
**WebhookFilters** | **[]string** | 
**WebhookMethod** | **string** | 

### Return type

[**IpMessagingV2Service**](IpMessagingV2Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

