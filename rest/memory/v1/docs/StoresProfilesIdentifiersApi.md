# StoresProfilesIdentifiersApi

All URIs are relative to *https://memory.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProfileIdentifier**](StoresProfilesIdentifiersApi.md#CreateProfileIdentifier) | **Post** /v1/Stores/{storeId}/Profiles/{profileId}/Identifiers | Add Profile Identifier
[**DeleteProfileIdentifier**](StoresProfilesIdentifiersApi.md#DeleteProfileIdentifier) | **Delete** /v1/Stores/{storeId}/Profiles/{profileId}/Identifiers/{idType} | Remove Profile Identifier
[**FetchProfileIdentifier**](StoresProfilesIdentifiersApi.md#FetchProfileIdentifier) | **Get** /v1/Stores/{storeId}/Profiles/{profileId}/Identifiers/{idType} | Retrieve Profile Identifier
[**ListProfileIdentifiers**](StoresProfilesIdentifiersApi.md#ListProfileIdentifiers) | **Get** /v1/Stores/{storeId}/Profiles/{profileId}/Identifiers | List Profile Identifiers
[**PatchProfileIdentifier**](StoresProfilesIdentifiersApi.md#PatchProfileIdentifier) | **Patch** /v1/Stores/{storeId}/Profiles/{profileId}/Identifiers/{idType} | Modify Profile Identifier



## CreateProfileIdentifier

> CreateProfileIdentifierResponse CreateProfileIdentifier(ctx, StoreIdProfileIdoptional)

Add Profile Identifier

Asynchronously attach a new identifier value to a profile. When multiple values exist for an `idType`, ordering, limits, and uniqueness checks are enforced by the corresponding identifier settings `limitPolicy` and  reflect any normalization. If the identifier already exists and points to another profile, resolution rules or merge processes may apply. Use the identifier specific endpoint to inspect, update, or remove existing values.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**ProfileId** | **string** | The unique identifier for the profile using Twilio Type ID (TTID) format.

### Other Parameters

Other parameters are passed through a pointer to a CreateProfileIdentifierParams struct


Name | Type | Description
------------- | ------------- | -------------
**Identifier** | [**Identifier**](Identifier.md) | 

### Return type

[**CreateProfileIdentifierResponse**](CreateProfileIdentifier202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProfileIdentifier

> DeleteProfileIdentifierResponse DeleteProfileIdentifier(ctx, StoreIdProfileIdIdTypeoptional)

Remove Profile Identifier

Asynchronously remove stored values for the identifier type. The optional `removeAll` query parameter allows bulk removal; otherwise the service selects a single value to delete using the identifier settings `limitPolicy`.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**ProfileId** | **string** | The unique identifier for the profile using Twilio Type ID (TTID) format.
**IdType** | **string** | Identifier type configured for the service (for example `email`, `phone`, or `externalId`).

### Other Parameters

Other parameters are passed through a pointer to a DeleteProfileIdentifierParams struct


Name | Type | Description
------------- | ------------- | -------------
**RemoveAll** | **bool** | When true, removes every stored value for the identifier type in a single operation. Defaults to false.

### Return type

[**DeleteProfileIdentifierResponse**](DeleteProfileIdentifier202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchProfileIdentifier

> IdentifierSet FetchProfileIdentifier(ctx, StoreIdProfileIdIdType)

Retrieve Profile Identifier

Retrieve all stored values for a specific identifier type on a profile. Values are normalized according to the configured identifier settings and returned in the order determined by its `limitPolicy`.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**ProfileId** | **string** | The unique identifier for the profile using Twilio Type ID (TTID) format.
**IdType** | **string** | Identifier type configured for the service (for example `email`, `phone`, or `externalId`).

### Other Parameters

Other parameters are passed through a pointer to a FetchProfileIdentifierParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IdentifierSet**](IdentifierSet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProfileIdentifiers

> []IdentifierSet ListProfileIdentifiers(ctx, StoreIdProfileIdoptional)

List Profile Identifiers

Retrieve all identifier types linked to a profile along with their stored values. Use this to audit which external keys (phone, email, externalId, etc.) are currently associated. Values are returned in the order enforced by each identifier settings `limitPolicy` and reflect any normalization applied by the service.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**ProfileId** | **string** | The unique identifier for the profile using Twilio Type ID (TTID) format.

### Other Parameters

Other parameters are passed through a pointer to a ListProfileIdentifiersParams struct


Name | Type | Description
------------- | ------------- | -------------
**Limit** | **int** | Max number of records to return.
**PageSize** | **int** | Max number of records to return in a page

### Return type

[**[]IdentifierSet**](IdentifierSet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchProfileIdentifier

> PatchProfileIdentifierResponse PatchProfileIdentifier(ctx, StoreIdProfileIdIdTypeoptional)

Modify Profile Identifier

Asynchronously replace one of the stored values associated with the identifier type on a profile. `newValue` is normalized according to the configured identifier settings.   If the new identifier already exists and points to another profile, merge processes may apply.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**ProfileId** | **string** | The unique identifier for the profile using Twilio Type ID (TTID) format.
**IdType** | **string** | Identifier type configured for the service (for example `email`, `phone`, or `externalId`).

### Other Parameters

Other parameters are passed through a pointer to a PatchProfileIdentifierParams struct


Name | Type | Description
------------- | ------------- | -------------
**IdentifierUpdate** | [**IdentifierUpdate**](IdentifierUpdate.md) | 

### Return type

[**PatchProfileIdentifierResponse**](PatchProfileIdentifier202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

