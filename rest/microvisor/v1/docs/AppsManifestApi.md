# AppsManifestApi

All URIs are relative to *https://microvisor.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAppManifest**](AppsManifestApi.md#FetchAppManifest) | **Get** /v1/Apps/{AppSid}/Manifest | 



## FetchAppManifest

> MicrovisorV1AppManifest FetchAppManifest(ctx, AppSid)



Retrieve the Manifest for an App.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AppSid** | **string** | A 34-character string that uniquely identifies this App.

### Other Parameters

Other parameters are passed through a pointer to a FetchAppManifestParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MicrovisorV1AppManifest**](MicrovisorV1AppManifest.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

