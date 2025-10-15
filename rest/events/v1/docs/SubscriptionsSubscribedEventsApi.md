# SubscriptionsSubscribedEventsApi

All URIs are relative to *https://events.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscribedEvent**](SubscriptionsSubscribedEventsApi.md#CreateSubscribedEvent) | **Post** /v1/Subscriptions/{SubscriptionSid}/SubscribedEvents | Add an event type to a Subscription.
[**DeleteSubscribedEvent**](SubscriptionsSubscribedEventsApi.md#DeleteSubscribedEvent) | **Delete** /v1/Subscriptions/{SubscriptionSid}/SubscribedEvents/{Type} | Remove an event type from a Subscription.
[**FetchSubscribedEvent**](SubscriptionsSubscribedEventsApi.md#FetchSubscribedEvent) | **Get** /v1/Subscriptions/{SubscriptionSid}/SubscribedEvents/{Type} | Read an Event for a Subscription.
[**ListSubscribedEvent**](SubscriptionsSubscribedEventsApi.md#ListSubscribedEvent) | **Get** /v1/Subscriptions/{SubscriptionSid}/SubscribedEvents | Retrieve a list of all Subscribed Event types for a Subscription.
[**UpdateSubscribedEvent**](SubscriptionsSubscribedEventsApi.md#UpdateSubscribedEvent) | **Post** /v1/Subscriptions/{SubscriptionSid}/SubscribedEvents/{Type} | Update an Event for a Subscription.



## CreateSubscribedEvent

> EventsV1SubscribedEvent CreateSubscribedEvent(ctx, SubscriptionSidoptional)

Add an event type to a Subscription.

Add an event type to a Subscription.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SubscriptionSid** | **string** | The unique SID identifier of the Subscription.

### Other Parameters

Other parameters are passed through a pointer to a CreateSubscribedEventParams struct


Name | Type | Description
------------- | ------------- | -------------
**Type** | **string** | Type of event being subscribed to.
**SchemaVersion** | **int** | The schema version that the Subscription should use.

### Return type

[**EventsV1SubscribedEvent**](EventsV1SubscribedEvent.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscribedEvent

> DeleteSubscribedEvent(ctx, SubscriptionSidType)

Remove an event type from a Subscription.

Remove an event type from a Subscription.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SubscriptionSid** | **string** | The unique SID identifier of the Subscription.
**Type** | **string** | Type of event being subscribed to.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSubscribedEventParams struct


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


## FetchSubscribedEvent

> EventsV1SubscribedEvent FetchSubscribedEvent(ctx, SubscriptionSidType)

Read an Event for a Subscription.

Read an Event for a Subscription.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SubscriptionSid** | **string** | The unique SID identifier of the Subscription.
**Type** | **string** | Type of event being subscribed to.

### Other Parameters

Other parameters are passed through a pointer to a FetchSubscribedEventParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**EventsV1SubscribedEvent**](EventsV1SubscribedEvent.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscribedEvent

> []EventsV1SubscribedEvent ListSubscribedEvent(ctx, SubscriptionSidoptional)

Retrieve a list of all Subscribed Event types for a Subscription.

Retrieve a list of all Subscribed Event types for a Subscription.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SubscriptionSid** | **string** | The unique SID identifier of the Subscription.

### Other Parameters

Other parameters are passed through a pointer to a ListSubscribedEventParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]EventsV1SubscribedEvent**](EventsV1SubscribedEvent.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscribedEvent

> EventsV1SubscribedEvent UpdateSubscribedEvent(ctx, SubscriptionSidTypeoptional)

Update an Event for a Subscription.

Update an Event for a Subscription.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SubscriptionSid** | **string** | The unique SID identifier of the Subscription.
**Type** | **string** | Type of event being subscribed to.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSubscribedEventParams struct


Name | Type | Description
------------- | ------------- | -------------
**SchemaVersion** | **int** | The schema version that the Subscription should use.

### Return type

[**EventsV1SubscribedEvent**](EventsV1SubscribedEvent.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

