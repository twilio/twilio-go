# ControlPlaneStoresTraitGroupsApi

All URIs are relative to *https://memory.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTraitGroup**](ControlPlaneStoresTraitGroupsApi.md#CreateTraitGroup) | **Post** /v1/ControlPlane/Stores/{storeId}/TraitGroups | Create a Trait Group
[**DeleteTraitGroup**](ControlPlaneStoresTraitGroupsApi.md#DeleteTraitGroup) | **Delete** /v1/ControlPlane/Stores/{storeId}/TraitGroups/{traitGroupName} | Delete a Trait Group
[**FetchTraitGroup**](ControlPlaneStoresTraitGroupsApi.md#FetchTraitGroup) | **Get** /v1/ControlPlane/Stores/{storeId}/TraitGroups/{traitGroupName} | Retrieve a Trait Group
[**ListTraitGroups**](ControlPlaneStoresTraitGroupsApi.md#ListTraitGroups) | **Get** /v1/ControlPlane/Stores/{storeId}/TraitGroups | List Trait Groups
[**PatchTraitGroup**](ControlPlaneStoresTraitGroupsApi.md#PatchTraitGroup) | **Patch** /v1/ControlPlane/Stores/{storeId}/TraitGroups/{traitGroupName} | Update Trait Group



## CreateTraitGroup

> CreateTraitGroupResponse CreateTraitGroup(ctx, StoreIdoptional)

Create a Trait Group

Creates a Trait Group with declarative trait schemas for this Memory Store. 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a CreateTraitGroupParams struct


Name | Type | Description
------------- | ------------- | -------------
**TraitGroupRequest** | [**TraitGroupRequest**](TraitGroupRequest.md) | 

### Return type

[**CreateTraitGroupResponse**](CreateTraitGroup202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTraitGroup

> DeleteTraitGroupResponse DeleteTraitGroup(ctx, StoreIdTraitGroupName)

Delete a Trait Group

Deletes the Trait Group and all associated Traits from the Memory Store. 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**TraitGroupName** | **string** | A unique Name of the TraitGroup

### Other Parameters

Other parameters are passed through a pointer to a DeleteTraitGroupParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**DeleteTraitGroupResponse**](DeleteTraitGroup202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTraitGroup

> FetchTraitGroupResponse FetchTraitGroup(ctx, StoreIdTraitGroupNameoptional)

Retrieve a Trait Group

Retrieve a specific Trait Group by its unique name for this Memory Store, with optional traits information.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**TraitGroupName** | **string** | A unique Name of the TraitGroup

### Other Parameters

Other parameters are passed through a pointer to a FetchTraitGroupParams struct


Name | Type | Description
------------- | ------------- | -------------
**IncludeTraits** | **bool** | Whether to include trait definitions in the response
**PageSize** | **int** | The maximum number of items to return per page, maximum of 100.
**PageToken** | **string** | The token for the page of results to retrieve.
**OrderBy** | **string** | Either 'ASC' or 'DESC' to sort results ascending or descending respectively.

### Return type

[**FetchTraitGroupResponse**](FetchTraitGroup200Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTraitGroups

> []TraitGroup ListTraitGroups(ctx, StoreIdoptional)

List Trait Groups

Returns a list of Trait Groups for this Memory Store, with optional traits information.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a ListTraitGroupsParams struct


Name | Type | Description
------------- | ------------- | -------------
**IncludeTraits** | **bool** | Whether to include trait definitions in the response
**PageSize** | **int** | The maximum number of items to return per page, maximum of 100.
**PageToken** | **string** | The token for the page of results to retrieve.
**OrderBy** | **string** | Either 'ASC' or 'DESC' to sort results ascending or descending respectively.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TraitGroup**](TraitGroup.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTraitGroup

> PatchTraitGroupResponse PatchTraitGroup(ctx, StoreIdTraitGroupNameoptional)

Update Trait Group

Partially updates a Trait Group. Only the fields provided in the request body will be updated. For traits: existing traits with matching keys will be updated, new keys will be added. To remove a trait, set its dataType to an empty string (\"\") - this serves as an explicit deletion marker. 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**TraitGroupName** | **string** | A unique Name of the TraitGroup

### Other Parameters

Other parameters are passed through a pointer to a PatchTraitGroupParams struct


Name | Type | Description
------------- | ------------- | -------------
**IfMatch** | **string** | Allows for optimistic concurrency control by making the request conditional. Server will only act if the resource's current Entity Tag (ETag) matches the one provided, preventing accidental overwrites.
**PatchTraitGroupRequest** | [**PatchTraitGroupRequest**](PatchTraitGroupRequest.md) | 

### Return type

[**PatchTraitGroupResponse**](PatchTraitGroup202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

