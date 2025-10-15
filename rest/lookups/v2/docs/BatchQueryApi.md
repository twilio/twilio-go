# BatchQueryApi

All URIs are relative to *https://lookups.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBulkLookup**](BatchQueryApi.md#CreateBulkLookup) | **Post** /v2/batch/query | In Request Bulk



## CreateBulkLookup

> LookupResponse1 CreateBulkLookup(ctx, optional)

In Request Bulk

Discussions made regarding how to help the customer to correlation request and response objects: - Respecting the natural order (requests vs. response) - Using phone numbers as unique key - Adding a correlation_id key

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateBulkLookupParams struct


Name | Type | Description
------------- | ------------- | -------------
**LookupRequest** | [**LookupRequest**](LookupRequest.md) | 

### Return type

[**LookupResponse1**](LookupResponse1.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

