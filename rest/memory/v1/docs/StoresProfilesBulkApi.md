# StoresProfilesBulkApi

All URIs are relative to *https://memory.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateProfilesBulk**](StoresProfilesBulkApi.md#UpdateProfilesBulk) | **Put** /v1/Stores/{storeId}/Profiles/Bulk | Bulk Upsert Profiles



## UpdateProfilesBulk

> UpdateProfilesBulkResponse UpdateProfilesBulk(ctx, StoreIdoptional)

Bulk Upsert Profiles

Create or update up to 1000 profiles in a single asynchronous batch. Each profile body follows the same structure as single profile creation. If a profile already exists its traits are merged (new keys added, existing keys overwritten). Large batches may take time to process; the 202 response indicates the batch has been accepted. Monitor downstream telemetry or audit logs to confirm completion.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a UpdateProfilesBulkParams struct


Name | Type | Description
------------- | ------------- | -------------
**UpdateProfilesBulkRequest** | [**UpdateProfilesBulkRequest**](UpdateProfilesBulkRequest.md) | 

### Return type

[**UpdateProfilesBulkResponse**](UpdateProfilesBulk202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

