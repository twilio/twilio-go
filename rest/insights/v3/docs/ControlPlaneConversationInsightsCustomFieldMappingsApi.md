# ControlPlaneConversationInsightsCustomFieldMappingsApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomFieldMapping**](ControlPlaneConversationInsightsCustomFieldMappingsApi.md#CreateCustomFieldMapping) | **Post** /v3/ControlPlane/ConversationInsights/CustomFieldMappings | Register a custom measure field mapping
[**FetchCustomFieldMapping**](ControlPlaneConversationInsightsCustomFieldMappingsApi.md#FetchCustomFieldMapping) | **Get** /v3/ControlPlane/ConversationInsights/CustomFieldMappings/{mappingId} | Fetch a custom field mapping
[**ListCustomFieldMappings**](ControlPlaneConversationInsightsCustomFieldMappingsApi.md#ListCustomFieldMappings) | **Get** /v3/ControlPlane/ConversationInsights/CustomFieldMappings | List custom measure field mappings for an account
[**PatchCustomFieldMapping**](ControlPlaneConversationInsightsCustomFieldMappingsApi.md#PatchCustomFieldMapping) | **Patch** /v3/ControlPlane/ConversationInsights/CustomFieldMappings/{mappingId} | Update a custom field mapping



## CreateCustomFieldMapping

> CustomFieldMappingResponse CreateCustomFieldMapping(ctx, optional)

Register a custom measure field mapping

Registers a custom field for the account. `entityMetadata.sourceType` selects where the field's value comes from, and the matching property on `entityMetadata` (`intelligenceOperator` for `INTELLIGENCE_OPERATOR`) carries its details. Only the `MEASURE` category is supported for now.  The request is validated before the mapping is created: - `name` contains only letters, digits and `_`, and is not already   registered for the account (otherwise `409`); - `description`, when present, is at most 150 characters; - for `INTELLIGENCE_OPERATOR`, the operator has an active version with   `outputFormat` `JSON`, and `entityMetadata.intelligenceOperator.field`   is one or two `.`-separated segments that resolve to a `number` or   `integer` property in that version's output schema (otherwise `400`); - the account has a free slot in the category; by default an account can   register up to 3 custom measures (otherwise `403`). Use   `GET /v3/ControlPlane/ConversationInsights/CustomFieldMappingsCapacity`   to check the remaining slots.  Create-only: if a mapping with the same `name` already exists for the account, `409` is returned with the existing mapping's `id` and nothing is changed. To modify an existing mapping, use `PATCH /v3/ControlPlane/ConversationInsights/CustomFieldMappings/{mappingId}`. 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateCustomFieldMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**CustomFieldMappingRequest** | [**CustomFieldMappingRequest**](CustomFieldMappingRequest.md) | 

### Return type

[**CustomFieldMappingResponse**](CustomFieldMappingResponse.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCustomFieldMapping

> CustomFieldMappingResponse FetchCustomFieldMapping(ctx, MappingId)

Fetch a custom field mapping

Returns a single custom field mapping by its `mappingId`. The `ETag` response header identifies the mapping's current version; supply it as `If-Match` on `PATCH` to guard against concurrent updates. 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MappingId** | **string** | The unique identifier of the custom field mapping, as returned in `id` when the mapping was created.

### Other Parameters

Other parameters are passed through a pointer to a FetchCustomFieldMappingParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**CustomFieldMappingResponse**](CustomFieldMappingResponse.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomFieldMappings

> []CustomFieldMappingResponse ListCustomFieldMappings(ctx, optional)

List custom measure field mappings for an account

Returns the custom field mappings registered for the account, one page at a time. Use `pageSize` to limit the number of mappings returned, and pass `meta.nextToken` from a response as `pageToken` to fetch the next page. On the last page, `meta.nextToken` is `null`. 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListCustomFieldMappingsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | The maximum number of resources to return
**PageToken** | **string** | Token for pagination
**Limit** | **int** | Max number of records to return.

### Return type

[**[]CustomFieldMappingResponse**](CustomFieldMappingResponse.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCustomFieldMapping

> CustomFieldMappingResponse PatchCustomFieldMapping(ctx, MappingIdoptional)

Update a custom field mapping

Partially updates an existing custom field mapping. Only the properties present in the body are changed. The updatable properties depend on the mapping's `sourceType`; for `INTELLIGENCE_OPERATOR`, only `entityMetadata.intelligenceOperator.field` can be updated today.  The new `field` is validated the same way as at registration: it must be one or two `.`-separated segments of letters, digits and `_`, and resolve to a `number` or `integer` property in the output schema of the operator's latest active version.  The mapping's `id`, `name`, `category`, `createdAt` and owning entity never change. Sending a `field` value equal to the current one is a no-op that returns `200` without modifying the mapping, so the operation is idempotent. Supply `If-Match` with the `ETag` returned by `GET`, `POST` or a previous `PATCH` to guard against concurrent updates.  Values already ingested are not recomputed; results processed after the update use the new `field`. 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MappingId** | **string** | The unique identifier of the custom field mapping, as returned in `id` when the mapping was created.

### Other Parameters

Other parameters are passed through a pointer to a PatchCustomFieldMappingParams struct


Name | Type | Description
------------- | ------------- | -------------
**IfMatch** | **string** | ETag of the custom field mapping, as returned in the `ETag` response header. When supplied, the update is applied only if the mapping has not changed since that ETag was issued; otherwise `412` is returned. 
**CustomFieldMappingUpdateRequest** | [**CustomFieldMappingUpdateRequest**](CustomFieldMappingUpdateRequest.md) | 

### Return type

[**CustomFieldMappingResponse**](CustomFieldMappingResponse.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

