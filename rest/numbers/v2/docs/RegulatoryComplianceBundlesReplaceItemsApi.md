# RegulatoryComplianceBundlesReplaceItemsApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReplaceItems**](RegulatoryComplianceBundlesReplaceItemsApi.md#CreateReplaceItems) | **Post** /v2/RegulatoryCompliance/Bundles/{BundleSid}/ReplaceItems | 



## CreateReplaceItems

> NumbersV2ReplaceItems CreateReplaceItems(ctx, BundleSidoptional)



Replaces all bundle items in the target bundle (specified in the path) with all the bundle items of the source bundle (specified by the from_bundle_sid body param)

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string** | The unique string that identifies the Bundle where the item assignments are going to be replaced.

### Other Parameters

Other parameters are passed through a pointer to a CreateReplaceItemsParams struct


Name | Type | Description
------------- | ------------- | -------------
**FromBundleSid** | **string** | The source bundle sid to copy the item assignments from.

### Return type

[**NumbersV2ReplaceItems**](NumbersV2ReplaceItems.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

