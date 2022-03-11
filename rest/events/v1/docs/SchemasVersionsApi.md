# SchemasVersionsApi

All URIs are relative to *https://events.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchSchemaVersion**](SchemasVersionsApi.md#FetchSchemaVersion) | **Get** /v1/Schemas/{Id}/Versions/{SchemaVersion} | 
[**ListSchemaVersion**](SchemasVersionsApi.md#ListSchemaVersion) | **Get** /v1/Schemas/{Id}/Versions | 



## FetchSchemaVersion

> EventsV1SchemaVersion FetchSchemaVersion(ctx, IdSchemaVersion)



Fetch a specific schema and version.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The unique identifier of the schema. Each schema can have multiple versions, that share the same id.
**SchemaVersion** | **int** | The version of the schema

### Other Parameters

Other parameters are passed through a pointer to a FetchSchemaVersionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**EventsV1SchemaVersion**](EventsV1SchemaVersion.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSchemaVersion

> []EventsV1SchemaVersion ListSchemaVersion(ctx, Idoptional)



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
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]EventsV1SchemaVersion**](EventsV1SchemaVersion.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

