# WebhooksSettingsApi

All URIs are relative to *https://webhooks.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSetting**](WebhooksSettingsApi.md#CreateSetting) | **Post** /v1/Webhooks/Settings | Create a new Setting
[**DeleteSetting**](WebhooksSettingsApi.md#DeleteSetting) | **Delete** /v1/Webhooks/Settings/{settingId} | Delete Setting (hard delete)
[**FetchSetting**](WebhooksSettingsApi.md#FetchSetting) | **Get** /v1/Webhooks/Settings/{settingId} | Get Setting by ID
[**ListSetting**](WebhooksSettingsApi.md#ListSetting) | **Get** /v1/Webhooks/Settings | List all Settings
[**PatchSetting**](WebhooksSettingsApi.md#PatchSetting) | **Patch** /v1/Webhooks/Settings/{settingId} | Update Setting friendlyName and description



## CreateSetting

> Setting CreateSetting(ctx, optional)

Create a new Setting

Creates an immutable webhook configuration.  ### Authentication: Inline vs Shared  **Inline (Recommended):** - Credentials directly in Setting (source: INLINE) - Use when: Each Setting has unique credentials  **Shared (Advanced):** - Multiple Settings share one AuthProfile resource (source: REFERENCE) - Use when: Centralizing credentials, managing one AuthProfile instead of duplicating across Settings  ### Immutability  Settings are immutable (except `friendlyName`, `description`).  **To change configuration:** Create new Setting → Update Rule → Delete old Setting.  This enables safe testing before production cutover. 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSettingParams struct


Name | Type | Description
------------- | ------------- | -------------
**IdempotencyKey** | **string** | Idempotency key for mutation operations. Must be a UUIDv7 (other UUID versions are rejected with a 400).
**PostSetting** | [**PostSetting**](PostSetting.md) | 

### Return type

[**Setting**](Setting.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSetting

> DeleteSetting(ctx, SettingId)

Delete Setting (hard delete)

Permanently deletes the Setting resource. Returns 409 Conflict if the Setting is still referenced by any Rule or if it's a system-created Setting.  **409 Resolution:** 1. List dependencies: `GET /Rules?settingId={id}` 2. Delete/update them 3. Retry deletion 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SettingId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteSettingParams struct


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


## FetchSetting

> Setting FetchSetting(ctx, SettingId)

Get Setting by ID

Returns a Setting by its ID, including full auth, signature, connection, and edge configuration.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SettingId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchSettingParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**Setting**](Setting.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSetting

> []Setting ListSetting(ctx, optional)

List all Settings

Returns a paginated list of Setting summaries. Use query parameters to filter by authProfileId or sharedKeyId, or to list Twilio-managed system Settings instead of your own.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSettingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | Number of items per page (max 100, reflecting the 50-resource account limit per resource type)
**PageToken** | **string** | Token for pagination
**AuthProfileId** | **string** | Filter Settings that reference this AuthProfile ID
**SharedKeyId** | **string** | Filter Settings that reference this SharedKey ID
**Type** | **string** | Whether to list your own Settings or Twilio-managed system Settings. - CUSTOMER (default): Settings created by you - SYSTEM: Twilio-managed Settings, such as the account's default legacy configuration 
**Limit** | **int** | Max number of records to return.

### Return type

[**[]Setting**](Setting.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSetting

> Setting PatchSetting(ctx, SettingIdoptional)

Update Setting friendlyName and description

Updates only the friendlyName and/or description fields. Settings are immutable except for these fields. To change configuration (auth, signature, connection, etc.), create a new Setting resource. Returns 409 Conflict if attempting to modify a system-created Setting. 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SettingId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a PatchSettingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PatchSetting** | [**PatchSetting**](PatchSetting.md) | 

### Return type

[**Setting**](Setting.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

