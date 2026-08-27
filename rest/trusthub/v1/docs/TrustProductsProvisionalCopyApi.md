# TrustProductsProvisionalCopyApi

All URIs are relative to *https://trusthub.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTrustProductProvisionalCopy**](TrustProductsProvisionalCopyApi.md#CreateTrustProductProvisionalCopy) | **Post** /v1/TrustProducts/{TrustProductSid}/ProvisionalCopy | Create the provisional copy for a given trust product with status: TWILIO_APPROVED, this is useful for making updates to an existing immutable trust product instance.
[**FetchTrustProductProvisionalCopy**](TrustProductsProvisionalCopyApi.md#FetchTrustProductProvisionalCopy) | **Get** /v1/TrustProducts/{TrustProductSid}/ProvisionalCopy | Fetch the provisional copy of a given trust product.



## CreateTrustProductProvisionalCopy

> TrusthubV1TrustProductProvisionalCopy CreateTrustProductProvisionalCopy(ctx, TrustProductSid)

Create the provisional copy for a given trust product with status: TWILIO_APPROVED, this is useful for making updates to an existing immutable trust product instance.

Create the provisional copy for a given trust product with status: TWILIO_APPROVED, this is useful for making updates to an existing immutable trust product instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrustProductSid** | **string** | The unique SID identifier of the Trust Product.

### Other Parameters

Other parameters are passed through a pointer to a CreateTrustProductProvisionalCopyParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrusthubV1TrustProductProvisionalCopy**](TrusthubV1TrustProductProvisionalCopy.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTrustProductProvisionalCopy

> TrusthubV1TrustProductProvisionalCopy FetchTrustProductProvisionalCopy(ctx, TrustProductSid)

Fetch the provisional copy of a given trust product.

Fetch the provisional copy of a given trust product.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrustProductSid** | **string** | The unique SID identifier of the Trust Product.

### Other Parameters

Other parameters are passed through a pointer to a FetchTrustProductProvisionalCopyParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrusthubV1TrustProductProvisionalCopy**](TrusthubV1TrustProductProvisionalCopy.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

