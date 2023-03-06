# RegulatoryComplianceIdentitiesApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdentity**](RegulatoryComplianceIdentitiesApi.md#CreateIdentity) | **Post** /v1/RegulatoryCompliance/Identities | 
[**FetchIdentity**](RegulatoryComplianceIdentitiesApi.md#FetchIdentity) | **Get** /v1/RegulatoryCompliance/Identities/{Sid} | 
[**ListIdentity**](RegulatoryComplianceIdentitiesApi.md#ListIdentity) | **Get** /v1/RegulatoryCompliance/Identities | 



## CreateIdentity

> NumbersV1Identity CreateIdentity(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateIdentityParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 128 characters long.
**Email** | **string** | The email address.
**Type** | **string** | The type of Identity resource to create.
**PurchaseIntentIsoCountry** | **string** | The ISO country code of the country.
**PurchaseIntentNumberType** | **string** | The type of phone number.
**StatusCallbackUrl** | **string** | The URL we should call using the `status_callback_method` to inform your application of status changes.
**StatusCallbackMethod** | **string** | The HTTP method we should use when calling `status_callback_url`.

### Return type

[**NumbersV1Identity**](NumbersV1Identity.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchIdentity

> NumbersV1Identity FetchIdentity(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Identity resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchIdentityParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NumbersV1Identity**](NumbersV1Identity.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentity

> []NumbersV1Identity ListIdentity(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListIdentityParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | **string** | Show only the Identity resources that have this status.
**PurchaseIntentIsoCountry** | **string** | Show only the Identity resources with this ISO country code.
**PurchaseIntentNumberType** | **string** | Show only the Identity resources with this phone number type.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]NumbersV1Identity**](NumbersV1Identity.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

