# WebhooksAuthProfilesApi

All URIs are relative to *https://webhooks.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthProfile**](WebhooksAuthProfilesApi.md#CreateAuthProfile) | **Post** /v1/Webhooks/AuthProfiles | Create a new AuthProfile resource
[**DeleteAuthProfile**](WebhooksAuthProfilesApi.md#DeleteAuthProfile) | **Delete** /v1/Webhooks/AuthProfiles/{authProfileId} | Delete AuthProfile (hard delete)
[**FetchAuthProfile**](WebhooksAuthProfilesApi.md#FetchAuthProfile) | **Get** /v1/Webhooks/AuthProfiles/{authProfileId} | Get AuthProfile by ID
[**ListAuthProfile**](WebhooksAuthProfilesApi.md#ListAuthProfile) | **Get** /v1/Webhooks/AuthProfiles | List all AuthProfile resources
[**PatchAuthProfile**](WebhooksAuthProfilesApi.md#PatchAuthProfile) | **Patch** /v1/Webhooks/AuthProfiles/{authProfileId} | Update AuthProfile friendlyName



## CreateAuthProfile

> AuthProfile CreateAuthProfile(ctx, optional)

Create a new AuthProfile resource

Creates a new reusable AuthProfile resource for use across multiple Settings.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateAuthProfileParams struct


Name | Type | Description
------------- | ------------- | -------------
**IdempotencyKey** | **string** | Idempotency key for mutation operations. Must be a UUIDv7 (other UUID versions are rejected with a 400).
**PostAuthProfile** | [**PostAuthProfile**](PostAuthProfile.md) | 

### Return type

[**AuthProfile**](AuthProfile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthProfile

> DeleteAuthProfile(ctx, AuthProfileId)

Delete AuthProfile (hard delete)

Permanently deletes the AuthProfile resource. Returns 409 Conflict if the AuthProfile is still referenced by any Settings resource.  **409 Resolution:** 1. List dependencies: `GET /Settings?authProfileId={id}` 2. Delete/update them 3. Retry deletion 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AuthProfileId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteAuthProfileParams struct


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


## FetchAuthProfile

> AuthProfile FetchAuthProfile(ctx, AuthProfileId)

Get AuthProfile by ID

Returns an AuthProfile resource by its ID. Secrets are never returned in GET responses.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AuthProfileId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchAuthProfileParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AuthProfile**](AuthProfile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthProfile

> []AuthProfileSummary ListAuthProfile(ctx, optional)

List all AuthProfile resources

Returns a paginated list of AuthProfile summaries. Secrets are never included in list responses.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAuthProfileParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | Number of items per page (max 100, reflecting the 50-resource account limit per resource type)
**PageToken** | **string** | Token for pagination
**Limit** | **int** | Max number of records to return.

### Return type

[**[]AuthProfileSummary**](AuthProfileSummary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAuthProfile

> AuthProfile PatchAuthProfile(ctx, AuthProfileIdoptional)

Update AuthProfile friendlyName

Updates only the friendlyName field. AuthProfile resources are immutable except for friendlyName. To change auth configuration (clientId, tokenUrl, etc.), create a new AuthProfile resource. 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AuthProfileId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a PatchAuthProfileParams struct


Name | Type | Description
------------- | ------------- | -------------
**PatchAuthProfile** | [**PatchAuthProfile**](PatchAuthProfile.md) | 

### Return type

[**AuthProfile**](AuthProfile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

