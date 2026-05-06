# ControlPlaneStoresIdentityResolutionSettingsApi

All URIs are relative to *https://memory.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchIdentityResolutionSettings**](ControlPlaneStoresIdentityResolutionSettingsApi.md#FetchIdentityResolutionSettings) | **Get** /v1/ControlPlane/Stores/{storeId}/IdentityResolutionSettings | Retrieve Identity Resolution Settings.
[**UpdateIdentityResolutionSettings**](ControlPlaneStoresIdentityResolutionSettingsApi.md#UpdateIdentityResolutionSettings) | **Put** /v1/ControlPlane/Stores/{storeId}/IdentityResolutionSettings | Update Identity Resolution Settings.



## FetchIdentityResolutionSettings

> IdentityResolutionSettings FetchIdentityResolutionSettings(ctx, StoreId)

Retrieve Identity Resolution Settings.

These settings determine how Customer Memory will use identifiers for searching,  merging, and resolving profiles.  > Some settings can only take on their default values and are not yet available to change, but coming soon.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a FetchIdentityResolutionSettingsParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IdentityResolutionSettings**](IdentityResolutionSettings.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdentityResolutionSettings

> UpdateIdentityResolutionSettingsResponse UpdateIdentityResolutionSettings(ctx, StoreIdoptional)

Update Identity Resolution Settings.

Update the identity resolution settings for this memory store, which determine how Customer Memory will use identifiers for searching,  merging, and resolving profiles. 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a UpdateIdentityResolutionSettingsParams struct


Name | Type | Description
------------- | ------------- | -------------
**IfMatch** | **string** | Allows for optimistic concurrency control by making the request conditional. Server will only act if the resource's current Entity Tag (ETag) matches the one provided, preventing accidental overwrites.
**IdentityResolutionSettingsCore** | [**IdentityResolutionSettingsCore**](IdentityResolutionSettingsCore.md) | 

### Return type

[**UpdateIdentityResolutionSettingsResponse**](UpdateIdentityResolutionSettings202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

