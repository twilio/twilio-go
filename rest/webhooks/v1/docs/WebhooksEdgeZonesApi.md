# WebhooksEdgeZonesApi

All URIs are relative to *https://webhooks.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEdgeZone**](WebhooksEdgeZonesApi.md#ListEdgeZone) | **Get** /v1/Webhooks/EdgeZones | List available edge zones



## ListEdgeZone

> []EdgeZone ListEdgeZone(ctx, optional)

List available edge zones

Returns the edge zone catalog for webhook delivery. Supports pagination via `pageSize` and `pageToken` (see [Pagination](#section/Pagination)).  **Read-only discovery endpoint** - use this to fetch available zone IDs before creating Settings.  ### Usage  Edge zones control geographic webhook delivery. Twilio selects the best available PoP within the zone at delivery time.  1. Call this endpoint to get available edge zones and their constituent PoPs 2. Use the zone `id` values in `Settings.edgeZones` when creating Settings  If edge zones are not specified in Settings, webhooks use the home region (same as the processing region). 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListEdgeZoneParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | Number of items per page (max 100, reflecting the 50-resource account limit per resource type)
**PageToken** | **string** | Token for pagination
**Limit** | **int** | Max number of records to return.

### Return type

[**[]EdgeZone**](EdgeZone.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

