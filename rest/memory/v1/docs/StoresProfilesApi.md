# StoresProfilesApi

All URIs are relative to *https://memory.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProfile**](StoresProfilesApi.md#CreateProfile) | **Post** /v1/Stores/{storeId}/Profiles | Create Profile
[**DeleteProfile**](StoresProfilesApi.md#DeleteProfile) | **Delete** /v1/Stores/{storeId}/Profiles/{profileId} | Delete Profile
[**FetchProfile**](StoresProfilesApi.md#FetchProfile) | **Get** /v1/Stores/{storeId}/Profiles/{profileId} | Get Profile
[**ListProfiles**](StoresProfilesApi.md#ListProfiles) | **Get** /v1/Stores/{storeId}/Profiles | List Profiles
[**PatchProfileTraits**](StoresProfilesApi.md#PatchProfileTraits) | **Patch** /v1/Stores/{storeId}/Profiles/{profileId} | Patch Profile Traits



## CreateProfile

> CreateProfileResponse CreateProfile(ctx, StoreIdoptional)

Create Profile

Create a new profile and set initial traits. The request synchronously resolves identity and either creates a new profile ID or retrieves the associated canonical profile ID based on any provided identifier traits present in the request. The request must contain at least one trait that is promoted to an identifier in its trait group settings. Any additional traits are queued for asynchronous processing.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a CreateProfileParams struct


Name | Type | Description
------------- | ------------- | -------------
**ProfileData** | [**ProfileData**](ProfileData.md) | 

### Return type

[**CreateProfileResponse**](CreateProfile202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProfile

> DeleteProfileResponse DeleteProfile(ctx, StoreIdProfileId)

Delete Profile

Asynchronously delete the profile permanently and all associated identifiers and traits. This operation is irreversible. Downstream caches or analytical stores may take time to reflect the deletion. Use cautiously and consider a soft‑delete strategy if regulatory or recovery requirements apply.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**ProfileId** | **string** | The unique identifier for the profile using Twilio Type ID (TTID) format.

### Other Parameters

Other parameters are passed through a pointer to a DeleteProfileParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**DeleteProfileResponse**](DeleteProfile202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchProfile

> Profile FetchProfile(ctx, StoreIdProfileIdoptional)

Get Profile

Retrieve profile traits by profile ID. Use the `traitGroups` query parameter to restrict results to a comma-separated allow list of trait group names. For large sets of traits, prefer using the dedicated `/Traits` endpoint for pagination.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**ProfileId** | **string** | The unique identifier for the profile using Twilio Type ID (TTID) format.

### Other Parameters

Other parameters are passed through a pointer to a FetchProfileParams struct


Name | Type | Description
------------- | ------------- | -------------
**TraitGroups** | **string** | Comma separated list of trait group names to include.

### Return type

[**Profile**](Profile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProfiles

> []string ListProfiles(ctx, StoreIdoptional)

List Profiles

Return a paginated list of profile IDs ordered by most recently created first. Use the optional paging parameters (`pageSize`, `pageToken`, `orderBy`) to control pagination and sorting. This endpoint is optimized for browsing newly created profiles and lightweight lookups where only the identifiers are needed before requesting full profile details.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a ListProfilesParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | The maximum number of items to return per page, maximum of 1000.
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


## PatchProfileTraits

> PatchProfileTraitsResponse PatchProfileTraits(ctx, StoreIdProfileIdoptional)

Patch Profile Traits

Merge one or more trait groups into an existing profile. Only the traits provided are added or updated; unspecified traits remain unchanged. Only pre-defined trait groups and traits configured for the memory store can be patched. To remove a trait entirely, set its value to null.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**ProfileId** | **string** | The unique identifier for the profile using Twilio Type ID (TTID) format.

### Other Parameters

Other parameters are passed through a pointer to a PatchProfileTraitsParams struct


Name | Type | Description
------------- | ------------- | -------------
**ProfilePatch** | [**ProfilePatch**](ProfilePatch.md) | 

### Return type

[**PatchProfileTraitsResponse**](PatchProfileTraits202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

