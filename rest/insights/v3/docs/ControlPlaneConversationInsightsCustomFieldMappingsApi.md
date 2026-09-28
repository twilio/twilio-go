# ControlPlaneConversationInsightsCustomFieldMappingsApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomFieldMapping**](ControlPlaneConversationInsightsCustomFieldMappingsApi.md#CreateCustomFieldMapping) | **Post** /v3/ControlPlane/ConversationInsights/CustomFieldMappings | Register a custom measure field mapping
[**FetchCustomFieldMappings**](ControlPlaneConversationInsightsCustomFieldMappingsApi.md#FetchCustomFieldMappings) | **Get** /v3/ControlPlane/ConversationInsights/CustomFieldMappings | Fetch custom measure field mappings for an account



## CreateCustomFieldMapping

> CustomFieldMappingResponse CreateCustomFieldMapping(ctx, optional)

Register a custom measure field mapping

Registers a custom field for an account under a category, sourced from an upstream entity — selected via the `entityMetadata.sourceType` discriminator, which determines the required shape of `entityMetadata`. Only the `MEASURE` category is supported for now.  For sources with a resolvable field-path concept, validates that `entityMetadata.field` resolves on the entity's latest active version by calling that source's own resolution API (the Intelligence fetch-operator API for `sourceType: INTELLIGENCE_OPERATOR`), then performs an idempotent upsert keyed on (accountSid, category, name). `entityMetadata.sourceType` plus that variant's own identity field (`operatorId` for `INTELLIGENCE_OPERATOR`) must match the existing registration's owner for the upsert to succeed — a mismatch returns `409`. Enforces the per-account, per-category capacity limit. 

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


## FetchCustomFieldMappings

> CustomFieldMappingsList FetchCustomFieldMappings(ctx, )

Fetch custom measure field mappings for an account

Returns custom field mappings registered for the account. 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchCustomFieldMappingsParams struct


### Return type

[**CustomFieldMappingsList**](CustomFieldMappingsList.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

