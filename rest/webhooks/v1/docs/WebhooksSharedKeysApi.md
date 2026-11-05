# WebhooksSharedKeysApi

All URIs are relative to *https://webhooks.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSharedKey**](WebhooksSharedKeysApi.md#CreateSharedKey) | **Post** /v1/Webhooks/SharedKeys | Create a new SharedKey
[**DeleteSharedKey**](WebhooksSharedKeysApi.md#DeleteSharedKey) | **Delete** /v1/Webhooks/SharedKeys/{sharedKeyId} | Delete SharedKey (hard delete)
[**FetchSharedKey**](WebhooksSharedKeysApi.md#FetchSharedKey) | **Get** /v1/Webhooks/SharedKeys/{sharedKeyId} | Get a SharedKey by ID
[**ListSharedKey**](WebhooksSharedKeysApi.md#ListSharedKey) | **Get** /v1/Webhooks/SharedKeys | List all SharedKeys
[**PatchSharedKey**](WebhooksSharedKeysApi.md#PatchSharedKey) | **Patch** /v1/Webhooks/SharedKeys/{sharedKeyId} | Update SharedKey friendlyName



## CreateSharedKey

> SharedKeyWithSecret CreateSharedKey(ctx, optional)

Create a new SharedKey

Creates a new Shared Key for webhook signature validation.  **IMPORTANT**: The `secret` field is ONLY returned in this response and cannot be retrieved later. Store it securely in your secrets management system. 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSharedKeyParams struct


Name | Type | Description
------------- | ------------- | -------------
**IdempotencyKey** | **string** | Idempotency key for mutation operations. Must be a UUIDv7 (other UUID versions are rejected with a 400).
**PostSharedKey** | [**PostSharedKey**](PostSharedKey.md) | 

### Return type

[**SharedKeyWithSecret**](SharedKeyWithSecret.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSharedKey

> DeleteSharedKey(ctx, SharedKeyId)

Delete SharedKey (hard delete)

Permanently deletes a Shared Key.  The API prevents deletion if any Settings reference this key. Delete the dependent Settings first, then delete the SharedKey. 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SharedKeyId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteSharedKeyParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSharedKey

> SharedKey FetchSharedKey(ctx, SharedKeyId)

Get a SharedKey by ID

Returns a SharedKey by its ID. The secret is never included in GET responses.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SharedKeyId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchSharedKeyParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SharedKey**](SharedKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSharedKey

> []SharedKeySummary ListSharedKey(ctx, optional)

List all SharedKeys

Returns paginated list of SharedKey summaries (secrets never included)

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSharedKeyParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | Number of items per page (max 100, reflecting the 50-resource account limit per resource type)
**PageToken** | **string** | Token for pagination
**Limit** | **int** | Max number of records to return.

### Return type

[**[]SharedKeySummary**](SharedKeySummary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSharedKey

> SharedKey PatchSharedKey(ctx, SharedKeyIdoptional)

Update SharedKey friendlyName

Update SharedKey's friendlyName only. SharedKeys are otherwise immutable.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SharedKeyId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a PatchSharedKeyParams struct


Name | Type | Description
------------- | ------------- | -------------
**PatchSharedKey** | [**PatchSharedKey**](PatchSharedKey.md) | 

### Return type

[**SharedKey**](SharedKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

