# StoresProfilesLookupApi

All URIs are relative to *https://memory.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProfilesLookup**](StoresProfilesLookupApi.md#CreateProfilesLookup) | **Post** /v1/Stores/{storeId}/Profiles/Lookup | Lookup Profiles



## CreateProfilesLookup

> CreateProfilesLookupResponse CreateProfilesLookup(ctx, StoreIdoptional)

Lookup Profiles

Find profiles that contain a specific identifier value (for example a phone number or email address). Submit an identifier object specifying the `idType` and `value`. The value is normalized using the configured identity resolution settings (such as phone number formatting) prior to matching. Multiple matches are returned if more than one profile is associated with the identifier. Returns canonical profile IDs (the earliest ID if profiles have been merged) along with the normalized value actually searched.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format

### Other Parameters

Other parameters are passed through a pointer to a CreateProfilesLookupParams struct


Name | Type | Description
------------- | ------------- | -------------
**Identifier** | [**Identifier**](Identifier.md) | 

### Return type

[**CreateProfilesLookupResponse**](CreateProfilesLookup200Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

