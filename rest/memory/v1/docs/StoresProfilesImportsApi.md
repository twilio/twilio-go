# StoresProfilesImportsApi

All URIs are relative to *https://memory.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProfilesImportV2**](StoresProfilesImportsApi.md#CreateProfilesImportV2) | **Post** /v1/Stores/{storeId}/Profiles/Imports | Import Profiles
[**FetchProfileImportV2**](StoresProfilesImportsApi.md#FetchProfileImportV2) | **Get** /v1/Stores/{storeId}/Profiles/Imports/{importId} | Get Import Status
[**ListProfileImportsV2**](StoresProfilesImportsApi.md#ListProfileImportsV2) | **Get** /v1/Stores/{storeId}/Profiles/Imports | Get a list of imports



## CreateProfilesImportV2

> CreateProfilesImportV1Response CreateProfilesImportV2(ctx, StoreIdoptional)

Import Profiles

Initiate a profile import by requesting a pre-signed upload URL and an associated `importId`. Upload your CSV to the returned URL (single PUT). Query the import status endpoint to track processing progress. This endpoint creates the import task and allocates resources for subsequent ingestion.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a CreateProfilesImportV2Params struct


Name | Type | Description
------------- | ------------- | -------------
**CreateProfilesImportV2Request** | [**CreateProfilesImportV2Request**](CreateProfilesImportV2Request.md) | 

### Return type

[**CreateProfilesImportV1Response**](CreateProfilesImportV2201Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchProfileImportV2

> FetchProfileImportV0Response FetchProfileImportV2(ctx, StoreIdImportId)

Get Import Status

Retrieve the current processing status of a previously submitted bulk import task. Query this endpoint using the `importId` returned from the upload URL request until a terminal state (COMPLETED or FAILED) is reached.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**ImportId** | **string** | The task identifier for the import process.

### Other Parameters

Other parameters are passed through a pointer to a FetchProfileImportV2Params struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**FetchProfileImportV0Response**](FetchProfileImportV2200Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProfileImportsV2

> []string ListProfileImportsV2(ctx, StoreIdoptional)

Get a list of imports

Retrieve a list of profile import task IDs that have been submitted for this service. Use these IDs to query individual import status details.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a ListProfileImportsV2Params struct


Name | Type | Description
------------- | ------------- | -------------
**Limit** | **int** | Max number of records to return.
**PageSize** | **int** | Max number of records to return in a page

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

