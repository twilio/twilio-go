# DefaultApi

All URIs are relative to *https://events.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSink**](DefaultApi.md#CreateSink) | **Post** /v1/Sinks | 
[**CreateSinkTest**](DefaultApi.md#CreateSinkTest) | **Post** /v1/Sinks/{Sid}/Test | 
[**CreateSinkValidate**](DefaultApi.md#CreateSinkValidate) | **Post** /v1/Sinks/{Sid}/Validate | 
[**CreateSubscribedEvent**](DefaultApi.md#CreateSubscribedEvent) | **Post** /v1/Subscriptions/{SubscriptionSid}/SubscribedEvents | 
[**CreateSubscription**](DefaultApi.md#CreateSubscription) | **Post** /v1/Subscriptions | 
[**DeleteSink**](DefaultApi.md#DeleteSink) | **Delete** /v1/Sinks/{Sid} | 
[**DeleteSubscribedEvent**](DefaultApi.md#DeleteSubscribedEvent) | **Delete** /v1/Subscriptions/{SubscriptionSid}/SubscribedEvents/{Type} | 
[**DeleteSubscription**](DefaultApi.md#DeleteSubscription) | **Delete** /v1/Subscriptions/{Sid} | 
[**FetchEventType**](DefaultApi.md#FetchEventType) | **Get** /v1/Types/{Type} | 
[**FetchSchema**](DefaultApi.md#FetchSchema) | **Get** /v1/Schemas/{Id} | 
[**FetchSchemaVersion**](DefaultApi.md#FetchSchemaVersion) | **Get** /v1/Schemas/{Id}/Versions/{SchemaVersion} | 
[**FetchSink**](DefaultApi.md#FetchSink) | **Get** /v1/Sinks/{Sid} | 
[**FetchSubscribedEvent**](DefaultApi.md#FetchSubscribedEvent) | **Get** /v1/Subscriptions/{SubscriptionSid}/SubscribedEvents/{Type} | 
[**FetchSubscription**](DefaultApi.md#FetchSubscription) | **Get** /v1/Subscriptions/{Sid} | 
[**ListEventType**](DefaultApi.md#ListEventType) | **Get** /v1/Types | 
[**ListSchemaVersion**](DefaultApi.md#ListSchemaVersion) | **Get** /v1/Schemas/{Id}/Versions | 
[**ListSink**](DefaultApi.md#ListSink) | **Get** /v1/Sinks | 
[**ListSubscribedEvent**](DefaultApi.md#ListSubscribedEvent) | **Get** /v1/Subscriptions/{SubscriptionSid}/SubscribedEvents | 
[**ListSubscription**](DefaultApi.md#ListSubscription) | **Get** /v1/Subscriptions | 
[**UpdateSubscribedEvent**](DefaultApi.md#UpdateSubscribedEvent) | **Post** /v1/Subscriptions/{SubscriptionSid}/SubscribedEvents/{Type} | 
[**UpdateSubscription**](DefaultApi.md#UpdateSubscription) | **Post** /v1/Subscriptions/{Sid} | 



## CreateSink

> EventsV1Sink CreateSink(ctx, optional)



Create a new Sink

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSinkParams struct


Name | Type | Description
------------- | ------------- | -------------
**Description** | **string** | A human readable description for the Sink **This value should not contain PII.**
**SinkConfiguration** | [**map[string]interface{}**](map[string]interface{}.md) | The information required for Twilio to connect to the provided Sink encoded as JSON.
**SinkType** | **string** | The Sink type. Can only be \\\&quot;kinesis\\\&quot; or \\\&quot;webhook\\\&quot; currently.

### Return type

[**EventsV1Sink**](EventsV1Sink.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSinkTest

> EventsV1SinkSinkTest CreateSinkTest(ctx, Sid)



Create a new Sink Test Event for the given Sink.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies the Sink to be Tested.

### Other Parameters

Other parameters are passed through a pointer to a CreateSinkTestParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**EventsV1SinkSinkTest**](EventsV1SinkSinkTest.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSinkValidate

> EventsV1SinkSinkValidate CreateSinkValidate(ctx, Sidoptional)



Validate that a test event for a Sink was received.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies the Sink being validated.

### Other Parameters

Other parameters are passed through a pointer to a CreateSinkValidateParams struct


Name | Type | Description
------------- | ------------- | -------------
**TestId** | **string** | A 34 character string that uniquely identifies the test event for a Sink being validated.

### Return type

[**EventsV1SinkSinkValidate**](EventsV1SinkSinkValidate.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubscribedEvent

> EventsV1SubscriptionSubscribedEvent CreateSubscribedEvent(ctx, SubscriptionSidoptional)



Create a new Subscribed Event type for the subscription

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SubscriptionSid** | **string** | The unique SID identifier of the Subscription.

### Other Parameters

Other parameters are passed through a pointer to a CreateSubscribedEventParams struct


Name | Type | Description
------------- | ------------- | -------------
**SchemaVersion** | **int32** | The schema version that the subscription should use.
**Type** | **string** | Type of event being subscribed to.

### Return type

[**EventsV1SubscriptionSubscribedEvent**](EventsV1SubscriptionSubscribedEvent.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
**Types** | **[]map[string]interface{}** | An array of objects containing the subscribed Event Types

### Return type

[**EventsV1Subscription**](EventsV1Subscription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSink

> DeleteSink(ctx, Sid)



Delete a specific Sink.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Sink.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSinkParams struct


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


## DeleteSubscribedEvent

> DeleteSubscribedEvent(ctx, SubscriptionSidType)



Remove an event type from a subscription.

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


## FetchEventType

> EventsV1EventType FetchEventType(ctx, Type)



Fetch a specific Event Type.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Type** | **string** | A string that uniquely identifies this Event Type.

### Other Parameters

Other parameters are passed through a pointer to a FetchEventTypeParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**EventsV1EventType**](EventsV1EventType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSchema

> EventsV1Schema FetchSchema(ctx, Id)



Fetch a specific schema with its nested versions.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The unique identifier of the schema. Each schema can have multiple versions, that share the same id.

### Other Parameters

Other parameters are passed through a pointer to a FetchSchemaParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**EventsV1Schema**](EventsV1Schema.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSchemaVersion

> EventsV1SchemaSchemaVersion FetchSchemaVersion(ctx, IdSchemaVersion)



Fetch a specific schema and version.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The unique identifier of the schema. Each schema can have multiple versions, that share the same id.
**SchemaVersion** | **int32** | The version of the schema

### Other Parameters

Other parameters are passed through a pointer to a FetchSchemaVersionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**EventsV1SchemaSchemaVersion**](EventsV1SchemaSchemaVersion.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSink

> EventsV1Sink FetchSink(ctx, Sid)



Fetch a specific Sink.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Sink.

### Other Parameters

Other parameters are passed through a pointer to a FetchSinkParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**EventsV1Sink**](EventsV1Sink.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSubscribedEvent

> EventsV1SubscriptionSubscribedEvent FetchSubscribedEvent(ctx, SubscriptionSidType)



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

[**EventsV1SubscriptionSubscribedEvent**](EventsV1SubscriptionSubscribedEvent.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

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
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEventType

> ListEventTypeResponse ListEventType(ctx, optional)



Retrieve a paginated list of all the available Event Types.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListEventTypeParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListEventTypeResponse**](ListEventTypeResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSchemaVersion

> ListSchemaVersionResponse ListSchemaVersion(ctx, Idoptional)



Retrieve a paginated list of versions of the schema.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The unique identifier of the schema. Each schema can have multiple versions, that share the same id.

### Other Parameters

Other parameters are passed through a pointer to a ListSchemaVersionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSchemaVersionResponse**](ListSchemaVersionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSink

> ListSinkResponse ListSink(ctx, optional)



Retrieve a paginated list of Sinks belonging to the account used to make the request.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSinkParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSinkResponse**](ListSinkResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscribedEvent

> ListSubscribedEventResponse ListSubscribedEvent(ctx, SubscriptionSidoptional)



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
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSubscribedEventResponse**](ListSubscribedEventResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscription

> ListSubscriptionResponse ListSubscription(ctx, optional)



Retrieve a paginated list of Subscriptions belonging to the account used to make the request.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSubscriptionParams struct


Name | Type | Description
------------- | ------------- | -------------
**SinkSid** | **string** | The SID of the sink that the list of Subscriptions should be filtered by.
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSubscriptionResponse**](ListSubscriptionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscribedEvent

> EventsV1SubscriptionSubscribedEvent UpdateSubscribedEvent(ctx, SubscriptionSidTypeoptional)



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
**SchemaVersion** | **int32** | The schema version that the subscription should use.

### Return type

[**EventsV1SubscriptionSubscribedEvent**](EventsV1SubscriptionSubscribedEvent.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

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

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

