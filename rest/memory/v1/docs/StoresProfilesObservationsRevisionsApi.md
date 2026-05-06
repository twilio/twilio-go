# StoresProfilesObservationsRevisionsApi

All URIs are relative to *https://memory.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListObservationRevisions**](StoresProfilesObservationsRevisionsApi.md#ListObservationRevisions) | **Get** /v1/Stores/{storeId}/Profiles/{profileId}/Observations/{observationId}/Revisions | List Observation Revisions



## ListObservationRevisions

> []ObservationInfo ListObservationRevisions(ctx, StoreIdProfileIdObservationIdoptional)

List Observation Revisions

Retrieve a chronologically ordered list of all past revisions of a specific observation by its ID. Revisions represent the complete history of updates and modifications made to an observation, with each revision capturing the full state at the time of change. Revisions are sorted by update time in descending order (newest first), allowing you to track how an observation has evolved through updates over time.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**ProfileId** | **string** | The unique identifier for the profile using Twilio Type ID (TTID) format.
**ObservationId** | **string** | The observation ID.

### Other Parameters

Other parameters are passed through a pointer to a ListObservationRevisionsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | The maximum number of items to return per page, maximum of 1000.
**PageToken** | **string** | The token for the page of results to retrieve.
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

