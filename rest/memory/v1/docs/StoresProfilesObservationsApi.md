# StoresProfilesObservationsApi

All URIs are relative to *https://memory.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProfileObservation**](StoresProfilesObservationsApi.md#CreateProfileObservation) | **Post** /v1/Stores/{storeId}/Profiles/{profileId}/Observations | Create Observations
[**DeleteProfileObservation**](StoresProfilesObservationsApi.md#DeleteProfileObservation) | **Delete** /v1/Stores/{storeId}/Profiles/{profileId}/Observations/{observationId} | Delete Observation
[**FetchProfileObservation**](StoresProfilesObservationsApi.md#FetchProfileObservation) | **Get** /v1/Stores/{storeId}/Profiles/{profileId}/Observations/{observationId} | Retrieve Observation
[**ListProfileObservations**](StoresProfilesObservationsApi.md#ListProfileObservations) | **Get** /v1/Stores/{storeId}/Profiles/{profileId}/Observations | List Observations
[**PatchProfileObservation**](StoresProfilesObservationsApi.md#PatchProfileObservation) | **Patch** /v1/Stores/{storeId}/Profiles/{profileId}/Observations/{observationId} | Update Observation



## CreateProfileObservation

> ObservationCreatedResponse CreateProfileObservation(ctx, StoreIdProfileIdoptional)

Create Observations

Create one or more transient observations associated with the specified profile. Supports both single observation creation and batch creation of up to 10 observations. Supports request compression for large batch operations and response compression for the response. All observations will be automatically indexed for semantic search capabilities. The content can be up to 4KB in length and should contain relevant information about the profile. The createdAt and updatedAt timestamps will be automatically set to the current time. Each observation will be created with a unique ID in Twilio Type ID (TTID) format.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**ProfileId** | **string** | The unique identifier for the profile using Twilio Type ID (TTID) format.

### Other Parameters

Other parameters are passed through a pointer to a CreateProfileObservationParams struct


Name | Type | Description
------------- | ------------- | -------------
**AcceptEncoding** | **string** | Compression algorithms supported by the client (e.g., gzip, deflate, br)
**ContentEncoding** | **string** | Compression algorithm used for the request body (e.g., gzip, deflate, br)
**CreateObservationsRequest** | [**CreateObservationsRequest**](CreateObservationsRequest.md) | 

### Return type

[**ObservationCreatedResponse**](ObservationCreatedResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProfileObservation

> DeleteProfileObservationResponse DeleteProfileObservation(ctx, StoreIdProfileIdObservationId)

Delete Observation

Delete a specific transient observation by its ID. This action is irreversible.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**ProfileId** | **string** | The unique identifier for the profile using Twilio Type ID (TTID) format.
**ObservationId** | **string** | The observation ID.

### Other Parameters

Other parameters are passed through a pointer to a DeleteProfileObservationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**DeleteProfileObservationResponse**](DeleteProfileObservation202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchProfileObservation

> ObservationInfo FetchProfileObservation(ctx, StoreIdProfileIdObservationId)

Retrieve Observation

Retrieve a specific transient observation by its ID for the given profile.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**ProfileId** | **string** | The unique identifier for the profile using Twilio Type ID (TTID) format.
**ObservationId** | **string** | The observation ID.

### Other Parameters

Other parameters are passed through a pointer to a FetchProfileObservationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ObservationInfo**](ObservationInfo.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProfileObservations

> []ObservationInfo ListProfileObservations(ctx, StoreIdProfileIdoptional)

List Observations

Retrieve a paginated list of transient observations for a specific profile. Observations are sorted by creation time in descending order by default. Results can be filtered by source parameter. Supports response compression for large datasets.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**ProfileId** | **string** | The unique identifier for the profile using Twilio Type ID (TTID) format.

### Other Parameters

Other parameters are passed through a pointer to a ListProfileObservationsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | The maximum number of items to return per page, maximum of 1000.
**PageToken** | **string** | The token for the page of results to retrieve.
**OrderBy** | **string** | Either 'ASC' or 'DESC' to sort results ascending or descending respectively.
**Source** | **string** | Filter by source. Allows letters, numbers, spaces, and URL-safe symbols. Excludes URL-unsafe characters like quotes, angle brackets, and control characters.
**CreatedAfter** | **time.Time** | Filter observations created after this timestamp (inclusive).
**CreatedBefore** | **time.Time** | Filter observations created before this timestamp (exclusive).
**AcceptEncoding** | **string** | Compression algorithms supported by the client (e.g., gzip, deflate, br)
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ObservationInfo**](ObservationInfo.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchProfileObservation

> PatchProfileObservationResponse PatchProfileObservation(ctx, StoreIdProfileIdObservationIdoptional)

Update Observation

Update a specific transient observation by its ID. Only provided fields will be updated. The updated timestamp will be automatically set. This allows for selective updates without needing to provide all fields.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**ProfileId** | **string** | The unique identifier for the profile using Twilio Type ID (TTID) format.
**ObservationId** | **string** | The observation ID.

### Other Parameters

Other parameters are passed through a pointer to a PatchProfileObservationParams struct


Name | Type | Description
------------- | ------------- | -------------
**ObservationCore** | [**ObservationCore**](ObservationCore.md) | 

### Return type

[**PatchProfileObservationResponse**](PatchProfileObservation202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

