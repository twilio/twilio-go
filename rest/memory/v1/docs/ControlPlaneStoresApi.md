# ControlPlaneStoresApi

All URIs are relative to *https://memory.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStore**](ControlPlaneStoresApi.md#CreateStore) | **Post** /v1/ControlPlane/Stores | Create a Memory Store
[**DeleteStore**](ControlPlaneStoresApi.md#DeleteStore) | **Delete** /v1/ControlPlane/Stores/{storeId} | Delete a Memory Store
[**FetchStore**](ControlPlaneStoresApi.md#FetchStore) | **Get** /v1/ControlPlane/Stores/{storeId} | Retrieve a Memory Store.
[**ListStores**](ControlPlaneStoresApi.md#ListStores) | **Get** /v1/ControlPlane/Stores | List Memory Stores
[**PatchStore**](ControlPlaneStoresApi.md#PatchStore) | **Patch** /v1/ControlPlane/Stores/{storeId} | Update a Memory Store



## CreateStore

> CreateStoreResponse CreateStore(ctx, optional)

Create a Memory Store

Create a new Memory Store for the Twilio account. Accounts can have multiple memory stores, up to a maximum limit of 15. Each memory store will automatically provision other dependent resources as needed, including Conversational Intelligence capabilities.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateStoreParams struct


Name | Type | Description
------------- | ------------- | -------------
**ServiceRequest** | [**ServiceRequest**](ServiceRequest.md) | 

### Return type

[**CreateStoreResponse**](CreateStore202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStore

> DeleteStoreResponse DeleteStore(ctx, StoreId)

Delete a Memory Store

Deletes the Memory Store and all associated resources including identity resolution settings, trait groups, profiles, traits, observations, and summaries.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a DeleteStoreParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**DeleteStoreResponse**](DeleteStore202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchStore

> Store FetchStore(ctx, StoreId)

Retrieve a Memory Store.

Retrieve the details of a specific Memory Store by its unique ID.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a FetchStoreParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**Store**](Store.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStores

> []string ListStores(ctx, optional)

List Memory Stores

Get a list of memory stores for the Twilio account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListStoresParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | The maximum number of items to return per page, maximum of 100.
**PageToken** | **string** | The token for the page of results to retrieve.
**OrderBy** | **string** | Either 'ASC' or 'DESC' to sort results ascending or descending respectively.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]string**](string.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchStore

> PatchStoreResponse PatchStore(ctx, StoreIdoptional)

Update a Memory Store

Partially update a Memory Store. Only the fields provided in the request body will be updated.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a PatchStoreParams struct


Name | Type | Description
------------- | ------------- | -------------
**IfMatch** | **string** | Allows for optimistic concurrency control by making the request conditional. Server will only act if the resource's current Entity Tag (ETag) matches the one provided, preventing accidental overwrites.
**PatchStoreRequest** | [**PatchStoreRequest**](PatchStoreRequest.md) | 

### Return type

[**PatchStoreResponse**](PatchStore202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

