# SchemasApi

All URIs are relative to *https://events.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchSchema**](SchemasApi.md#FetchSchema) | **Get** /v1/Schemas/{Id} | 



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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

