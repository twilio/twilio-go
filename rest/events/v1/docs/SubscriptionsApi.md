# SubscriptionsApi

All URIs are relative to *https://events.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscription**](SubscriptionsApi.md#CreateSubscription) | **Post** /v1/Subscriptions | 
[**DeleteSubscription**](SubscriptionsApi.md#DeleteSubscription) | **Delete** /v1/Subscriptions/{Sid} | 
[**FetchSubscription**](SubscriptionsApi.md#FetchSubscription) | **Get** /v1/Subscriptions/{Sid} | 
[**ListSubscription**](SubscriptionsApi.md#ListSubscription) | **Get** /v1/Subscriptions | 
[**UpdateSubscription**](SubscriptionsApi.md#UpdateSubscription) | **Post** /v1/Subscriptions/{Sid} | 



## CreateSubscription

> EventsV1Subscription CreateSubscription(ctx, optional)



Create a new Subscription.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSubscriptionParams struct


Name | Type | Description
------------- | ------------- | -------------
**Description** | **string** | A human readable description for the Subscription **This value should not contain PII.**
**SinkSid** | **string** | The SID of the sink that events selected by this subscription should be sent to. Sink must be active for the subscription to be created.
**Types** | **[]interface{}** | An array of objects containing the subscribed Event Types

### Return type

[**EventsV1Subscription**](EventsV1Subscription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscription

> DeleteSubscription(ctx, Sid)



Delete a specific Subscription.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Subscription.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSubscriptionParams struct


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


## FetchSubscription

> EventsV1Subscription FetchSubscription(ctx, Sid)



Fetch a specific Subscription.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Subscription.

### Other Parameters

Other parameters are passed through a pointer to a FetchSubscriptionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**EventsV1Subscription**](EventsV1Subscription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscription

> []EventsV1Subscription ListSubscription(ctx, optional)



Retrieve a paginated list of Subscriptions belonging to the account used to make the request.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSubscriptionParams struct


Name | Type | Description
------------- | ------------- | -------------
**SinkSid** | **string** | The SID of the sink that the list of Subscriptions should be filtered by.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]EventsV1Subscription**](EventsV1Subscription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscription

> EventsV1Subscription UpdateSubscription(ctx, Sidoptional)



Update a Subscription.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Subscription.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSubscriptionParams struct


Name | Type | Description
------------- | ------------- | -------------
**Description** | **string** | A human readable description for the Subscription.
**SinkSid** | **string** | The SID of the sink that events selected by this subscription should be sent to. Sink must be active for the subscription to be created.

### Return type

[**EventsV1Subscription**](EventsV1Subscription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

