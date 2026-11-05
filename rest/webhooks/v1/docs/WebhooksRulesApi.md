# WebhooksRulesApi

All URIs are relative to *https://webhooks.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRule**](WebhooksRulesApi.md#CreateRule) | **Post** /v1/Webhooks/Rules | Create a new Rule (async)
[**DeleteRule**](WebhooksRulesApi.md#DeleteRule) | **Delete** /v1/Webhooks/Rules/{ruleId} | Delete Rule (hard delete, async)
[**FetchRule**](WebhooksRulesApi.md#FetchRule) | **Get** /v1/Webhooks/Rules/{ruleId} | Get a Rule by ID
[**ListRule**](WebhooksRulesApi.md#ListRule) | **Get** /v1/Webhooks/Rules | List all Rules
[**PatchRule**](WebhooksRulesApi.md#PatchRule) | **Patch** /v1/Webhooks/Rules/{ruleId} | Update Rule (async)



## CreateRule

> RuleOperationResponse CreateRule(ctx, optional)

Create a new Rule (async)

Creates a URL matching rule. Returns immediately after validation, then propagates globally.  ### How It Works  **Immediate:** Request validated → 202 Accepted (with `Operation-Id` header) or 4xx error  **Async (30-60s):** Propagates globally. Poll `/Operations/{id}` for status (`pending` → `completed`)  ### During Propagation  During propagation, matching webhooks may use other rules (e.g., default) until this rule is active.  ### Security Restrictions  Default rules (matchType: default) cannot use Basic Authentication. Use OAuth2, Digest, or create specific URL patterns instead. 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateRuleParams struct


Name | Type | Description
------------- | ------------- | -------------
**IdempotencyKey** | **string** | Idempotency key for mutation operations. Must be a UUIDv7 (other UUID versions are rejected with a 400).
**PostRule** | [**PostRule**](PostRule.md) | 

### Return type

[**RuleOperationResponse**](RuleOperationResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRule

> RuleOperationResponse DeleteRule(ctx, RuleIdoptional)

Delete Rule (hard delete, async)

Permanently deletes a Rule. Returns 202 Accepted; the deletion propagates globally in 30-60 seconds. Poll the returned Operation-Id for completion.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RuleId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteRuleParams struct


Name | Type | Description
------------- | ------------- | -------------
**IdempotencyKey** | **string** | Idempotency key for mutation operations. Must be a UUIDv7 (other UUID versions are rejected with a 400).

### Return type

[**RuleOperationResponse**](RuleOperationResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRule

> Rule FetchRule(ctx, RuleId)

Get a Rule by ID

Returns a Rule by its ID, including matchType, match value, and current traffic distribution settings.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RuleId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchRuleParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**Rule**](Rule.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRule

> []Rule ListRule(ctx, optional)

List all Rules

Returns paginated list of Rules sorted by evaluation order: 1. exact-match rules (evaluated first, longest match first) 2. prefix-match rules (longest prefix first) 3. default rule (matches all incoming requests when no other rule matches) 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListRuleParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | Number of items per page (max 100, reflecting the 50-resource account limit per resource type)
**PageToken** | **string** | Token for pagination
**SettingId** | **string** | Filter Rules that reference this Setting ID
**Limit** | **int** | Max number of records to return.

### Return type

[**[]Rule**](Rule.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRule

> RuleOperationResponse PatchRule(ctx, RuleIdoptional)

Update Rule (async)

Updates rule configuration. Returns 202 Accepted, propagates in 30-60 seconds.  ### Gradual Rollout Pattern  Test configuration changes with production traffic before full deployment.  **Strategy:** Gradually shift traffic to new Setting, monitor, increase to 100%.  **Flow:** 90/10 → 50/50 → 0/100 (rollback to 100/0 if issues)  ### During Propagation  During propagation (~30-60s), webhooks may use previous or updated config until change is active.  ### Use Cases  - **Credential rotation**: Verify new auth works before full cutover - **Timeout tuning**: Test connection settings with real traffic  ### What Can Change  - `friendlyName`, `description`: Metadata - `settings`: Traffic distribution (percentages must sum to 100) - `matchType` + its companion value (`prefix` or `value`): changes the rule's match type and value together (async, same 30-60s propagation as any other patch). Returns `409` if the new match key collides with another Rule, same check `POST` runs on create. Patching a Rule to the match config it already has is not a collision. 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RuleId** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a PatchRuleParams struct


Name | Type | Description
------------- | ------------- | -------------
**IdempotencyKey** | **string** | Idempotency key for mutation operations. Must be a UUIDv7 (other UUID versions are rejected with a 400).
**PatchRule** | [**PatchRule**](PatchRule.md) | 

### Return type

[**RuleOperationResponse**](RuleOperationResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

