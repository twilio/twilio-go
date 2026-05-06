# StoresProfilesTraitsApi

All URIs are relative to *https://memory.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListProfileTraits**](StoresProfilesTraitsApi.md#ListProfileTraits) | **Get** /v1/Stores/{storeId}/Profiles/{profileId}/Traits | List Profile Traits



## ListProfileTraits

> []FullTrait ListProfileTraits(ctx, StoreIdProfileIdoptional)

List Profile Traits

Return a paginated, optionally filtered list of individual traits for a profile (flattened view). This is useful for UIs or tools that need to browse trait metadata (names, groups, timestamps). Use the optional `traitGroups` filter to scope results and the paging parameters to iterate through larger sets.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**ProfileId** | **string** | The unique identifier for the profile using Twilio Type ID (TTID) format.

### Other Parameters

Other parameters are passed through a pointer to a ListProfileTraitsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | The maximum number of items to return per page, maximum of 1000.
**PageToken** | **string** | The token for the page of results to retrieve.
**OrderBy** | **string** | Either 'ASC' or 'DESC' to sort results ascending or descending respectively.
**TraitGroups** | **string** | Comma separated list of trait group names to include.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]FullTrait**](FullTrait.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

