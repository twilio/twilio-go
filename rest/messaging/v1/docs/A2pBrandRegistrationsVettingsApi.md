# A2pBrandRegistrationsVettingsApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBrandVetting**](A2pBrandRegistrationsVettingsApi.md#CreateBrandVetting) | **Post** /v1/a2p/BrandRegistrations/{BrandSid}/Vettings | 
[**FetchBrandVetting**](A2pBrandRegistrationsVettingsApi.md#FetchBrandVetting) | **Get** /v1/a2p/BrandRegistrations/{BrandSid}/Vettings/{BrandVettingSid} | 
[**ListBrandVetting**](A2pBrandRegistrationsVettingsApi.md#ListBrandVetting) | **Get** /v1/a2p/BrandRegistrations/{BrandSid}/Vettings | 



## CreateBrandVetting

> MessagingV1BrandVetting CreateBrandVetting(ctx, BrandSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BrandSid** | **string** | The SID of the Brand Registration resource of the vettings to create .

### Other Parameters

Other parameters are passed through a pointer to a CreateBrandVettingParams struct


Name | Type | Description
------------- | ------------- | -------------
**VettingProvider** | **string** | 
**VettingId** | **string** | The unique ID of the vetting

### Return type

[**MessagingV1BrandVetting**](MessagingV1BrandVetting.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchBrandVetting

> MessagingV1BrandVetting FetchBrandVetting(ctx, BrandSidBrandVettingSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BrandSid** | **string** | The SID of the Brand Registration resource of the vettings to read .
**BrandVettingSid** | **string** | The Twilio SID of the third-party vetting record.

### Other Parameters

Other parameters are passed through a pointer to a FetchBrandVettingParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MessagingV1BrandVetting**](MessagingV1BrandVetting.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBrandVetting

> []MessagingV1BrandVetting ListBrandVetting(ctx, BrandSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BrandSid** | **string** | The SID of the Brand Registration resource of the vettings to read .

### Other Parameters

Other parameters are passed through a pointer to a ListBrandVettingParams struct


Name | Type | Description
------------- | ------------- | -------------
**VettingProvider** | **string** | The third-party provider of the vettings to read
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]MessagingV1BrandVetting**](MessagingV1BrandVetting.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

