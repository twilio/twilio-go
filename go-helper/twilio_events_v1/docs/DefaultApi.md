# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSink**](DefaultApi.md#CreateSink) | **Post** /v1/Sinks | 
[**CreateSinkTest**](DefaultApi.md#CreateSinkTest) | **Post** /v1/Sinks/{Sid}/Test | 
[**CreateSinkValidate**](DefaultApi.md#CreateSinkValidate) | **Post** /v1/Sinks/{Sid}/Validate | 
[**CreateSubscription**](DefaultApi.md#CreateSubscription) | **Post** /v1/Subscriptions | 
[**DeleteSink**](DefaultApi.md#DeleteSink) | **Delete** /v1/Sinks/{Sid} | 
[**DeleteSubscription**](DefaultApi.md#DeleteSubscription) | **Delete** /v1/Subscriptions/{Sid} | 
[**FetchEventType**](DefaultApi.md#FetchEventType) | **Get** /v1/Types/{Type} | 
[**FetchSchema**](DefaultApi.md#FetchSchema) | **Get** /v1/Schemas/{Id} | 
[**FetchSink**](DefaultApi.md#FetchSink) | **Get** /v1/Sinks/{Sid} | 
[**FetchSubscription**](DefaultApi.md#FetchSubscription) | **Get** /v1/Subscriptions/{Sid} | 
[**FetchVersion**](DefaultApi.md#FetchVersion) | **Get** /v1/Schemas/{Id}/Versions/{SchemaVersion} | 
[**ListEventType**](DefaultApi.md#ListEventType) | **Get** /v1/Types | 
[**ListSink**](DefaultApi.md#ListSink) | **Get** /v1/Sinks | 
[**ListSubscribedEvent**](DefaultApi.md#ListSubscribedEvent) | **Get** /v1/Subscriptions/{SubscriptionSid}/SubscribedEvents | 
[**ListSubscription**](DefaultApi.md#ListSubscription) | **Get** /v1/Subscriptions | 
[**ListVersion**](DefaultApi.md#ListVersion) | **Get** /v1/Schemas/{Id}/Versions | 
[**UpdateSubscription**](DefaultApi.md#UpdateSubscription) | **Post** /v1/Subscriptions/{Sid} | 



## CreateSink

> EventsV1Sink CreateSink(ctx, optional)



Create a new Sink

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateSinkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSinkOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **optional.String**| A human readable description for the Sink | 
 **sinkConfiguration** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The information required for Twilio to connect to the provided Sink encoded as JSON. | 
 **sinkType** | **optional.String**| The Sink type. Can only be \\\&quot;kinesis\\\&quot; currently. | 

### Return type

[**EventsV1Sink**](events.v1.sink.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSinkTest

> EventsV1SinkSinkTest CreateSinkTest(ctx, sid)



Create a new Sink Test Event for the given Sink.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| A 34 character string that uniquely identifies the Sink to be Tested. | 

### Return type

[**EventsV1SinkSinkTest**](events.v1.sink.sink_test.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSinkValidate

> EventsV1SinkSinkValidate CreateSinkValidate(ctx, sid, optional)



Validate that a test event for a Sink was received.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**|  | 
 **optional** | ***CreateSinkValidateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSinkValidateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testId** | **optional.String**| A 34 character string that uniquely identifies the test event for a Sink being validated. | 

### Return type

[**EventsV1SinkSinkValidate**](events.v1.sink.sink_validate.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubscription

> EventsV1Subscription CreateSubscription(ctx, optional)



Create a new Subscription.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSubscriptionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **optional.String**| A human readable description for the Subscription | 
 **sinkSid** | **optional.String**| The SID of the sink that events selected by this subscription should be sent to. Sink must be active for the subscription to be created. | 
 **types** | [**optional.Interface of []map[string]interface{}**](map[string]interface{}.md)| Contains a dictionary of URL links to nested resources of this Subscription. | 

### Return type

[**EventsV1Subscription**](events.v1.subscription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSink

> DeleteSink(ctx, sid)



Delete a specific Sink.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| A 34 character string that uniquely identifies this Sink. | 

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


## DeleteSubscription

> DeleteSubscription(ctx, sid)



Delete a specific Subscription.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| A 34 character string that uniquely identifies this Subscription. | 

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


## FetchEventType

> EventsV1EventType FetchEventType(ctx, type_)



Fetch a specific Event Type.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string**| A string that uniquely identifies this Event Type. | 

### Return type

[**EventsV1EventType**](events.v1.event_type.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSchema

> EventsV1Schema FetchSchema(ctx, id)



Fetch a specific schema with its nested versions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The unique identifier of the schema. Each schema can have multiple versions, that share the same id. | 

### Return type

[**EventsV1Schema**](events.v1.schema.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSink

> EventsV1Sink FetchSink(ctx, sid)



Fetch a specific Sink.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| A 34 character string that uniquely identifies this Sink. | 

### Return type

[**EventsV1Sink**](events.v1.sink.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSubscription

> EventsV1Subscription FetchSubscription(ctx, sid)



Fetch a specific Subscription.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| A 34 character string that uniquely identifies this Subscription. | 

### Return type

[**EventsV1Subscription**](events.v1.subscription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchVersion

> EventsV1SchemaVersion FetchVersion(ctx, id, schemaVersion)



Fetch a specific schema and version.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The unique identifier of the schema. Each schema can have multiple versions, that share the same id. | 
**schemaVersion** | **int32**| The version of the schema | 

### Return type

[**EventsV1SchemaVersion**](events.v1.schema.version.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEventType

> InlineResponse2004 ListEventType(ctx, optional)



Retrieve a paginated list of all the available Event Types.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListEventTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListEventTypeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSink

> InlineResponse2001 ListSink(ctx, optional)



Retrieve a paginated list of Sinks belonging to the account used to make the request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListSinkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSinkOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscribedEvent

> InlineResponse2003 ListSubscribedEvent(ctx, subscriptionSid, optional)



Retrieve a list of all Subscribed Event types for a Subscription.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionSid** | **string**| The unique SID identifier of the Subscription. | 
 **optional** | ***ListSubscribedEventOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSubscribedEventOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscription

> InlineResponse2002 ListSubscription(ctx, optional)



Retrieve a paginated list of Subscriptions belonging to the account used to make the request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSubscriptionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sinkSid** | **optional.String**| The SID of the sink that the list of Subscriptions should be filtered by. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVersion

> InlineResponse200 ListVersion(ctx, id, optional)



Retrieve a paginated list of versions of the schema.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The unique identifier of the schema. Each schema can have multiple versions, that share the same id. | 
 **optional** | ***ListVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListVersionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscription

> EventsV1Subscription UpdateSubscription(ctx, sid, optional)



Update a Subscription.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| A 34 character string that uniquely identifies this Subscription. | 
 **optional** | ***UpdateSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSubscriptionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **description** | **optional.String**| A human readable description for the Subscription. | 
 **sinkSid** | **optional.String**| The SID of the sink that events selected by this subscription should be sent to. Sink must be active for the subscription to be created. | 

### Return type

[**EventsV1Subscription**](events.v1.subscription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

