# RegulatoryComplianceBundlesClonesApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBundleClone**](RegulatoryComplianceBundlesClonesApi.md#CreateBundleClone) | **Post** /v2/RegulatoryCompliance/Bundles/{BundleSid}/Clones | Creates a new clone of the Bundle in target Account. It will internally create clones of all the bundle items (identities and documents) of the original bundle



## CreateBundleClone

> NumbersV2BundleClone CreateBundleClone(ctx, BundleSidoptional)

Creates a new clone of the Bundle in target Account. It will internally create clones of all the bundle items (identities and documents) of the original bundle

Creates a new clone of the Bundle in target Account. It will internally create clones of all the bundle items (identities and documents) of the original bundle

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string** | The unique string that identifies the Bundle to be cloned.

### Other Parameters

Other parameters are passed through a pointer to a CreateBundleCloneParams struct


Name | Type | Description
------------- | ------------- | -------------
**TargetAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) where the bundle needs to be cloned.
**MoveToDraft** | **bool** | If set to true, the cloned bundle will be in the DRAFT state, else it will be twilio-approved
**FriendlyName** | **string** | The string that you assigned to describe the cloned bundle.

### Return type

[**NumbersV2BundleClone**](NumbersV2BundleClone.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

