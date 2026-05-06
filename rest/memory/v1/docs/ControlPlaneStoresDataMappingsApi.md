# ControlPlaneStoresDataMappingsApi

All URIs are relative to *https://memory.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDataMapping**](ControlPlaneStoresDataMappingsApi.md#CreateDataMapping) | **Post** /v1/ControlPlane/Stores/{storeId}/DataMappings | Create a Data Mapping
[**DeleteDataMapping**](ControlPlaneStoresDataMappingsApi.md#DeleteDataMapping) | **Delete** /v1/ControlPlane/Stores/{storeId}/DataMappings/{dataMappingId} | Delete a Data Mapping
[**FetchDataMapping**](ControlPlaneStoresDataMappingsApi.md#FetchDataMapping) | **Get** /v1/ControlPlane/Stores/{storeId}/DataMappings/{dataMappingId} | Retrieve a Data Mapping
[**ListDataMappings**](ControlPlaneStoresDataMappingsApi.md#ListDataMappings) | **Get** /v1/ControlPlane/Stores/{storeId}/DataMappings | List Data Mappings
[**PatchDataMapping**](ControlPlaneStoresDataMappingsApi.md#PatchDataMapping) | **Patch** /v1/ControlPlane/Stores/{storeId}/DataMappings/{dataMappingId} | Update a Data Mapping



## CreateDataMapping

> CreateDataMappingResponse CreateDataMapping(ctx, StoreIdoptional)

Create a Data Mapping

Create a new data mapping to connect an external data source to this Memory Store. For DATASET types, validates that all mapped Trait Groups and traits exist and that dataset field data types match their respective mapped trait data types.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a CreateDataMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**CreateDataMappingInput** | [**CreateDataMappingInput**](CreateDataMappingInput.md) | 

### Return type

[**CreateDataMappingResponse**](CreateDataMapping202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDataMapping

> DeleteDataMappingResponse DeleteDataMapping(ctx, StoreIdDataMappingId)

Delete a Data Mapping

Delete a data mapping from the Memory Store. This action cannot be undone.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**DataMappingId** | **string** | A unique DataMapping ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a DeleteDataMappingParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**DeleteDataMappingResponse**](DeleteDataMapping202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDataMapping

> DataMapping FetchDataMapping(ctx, StoreIdDataMappingId)

Retrieve a Data Mapping

Retrieve the details of a specific data mapping by its unique ID.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**DataMappingId** | **string** | A unique DataMapping ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a FetchDataMappingParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**DataMapping**](DataMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDataMappings

> []DataMapping ListDataMappings(ctx, StoreIdoptional)

List Data Mappings

Get a list of data mappings configured for this Memory Store.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a ListDataMappingsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | The maximum number of items to return per page, maximum of 100.
**PageToken** | **string** | The token for the page of results to retrieve.
**OrderBy** | **string** | Either 'ASC' or 'DESC' to sort results ascending or descending respectively.
**Type** | [**DataMappingType**](DataMappingTypeDataMappingType.md) | Filter data mappings by type.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]DataMapping**](DataMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDataMapping

> PatchDataMappingResponse PatchDataMapping(ctx, StoreIdDataMappingIdoptional)

Update a Data Mapping

Partially update a data mapping. Only the fields provided in the request body will be updated, including replacing the entire contents of the `mappings` array.  For DATASET types, validates that all mapped Trait Groups and traits exist and that dataset field data types match their respective mapped trait data types.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**DataMappingId** | **string** | A unique DataMapping ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a PatchDataMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**IfMatch** | **string** | Allows for optimistic concurrency control by making the request conditional. Server will only act if the resource's current Entity Tag (ETag) matches the one provided, preventing accidental overwrites.
**DataMappingCore** | [**DataMappingCore**](DataMappingCore.md) | 

### Return type

[**PatchDataMappingResponse**](PatchDataMapping202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

