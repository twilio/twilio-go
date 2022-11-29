# RegulatoryComplianceBundlesApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBundle**](RegulatoryComplianceBundlesApi.md#CreateBundle) | **Post** /v2/RegulatoryCompliance/Bundles | 
[**DeleteBundle**](RegulatoryComplianceBundlesApi.md#DeleteBundle) | **Delete** /v2/RegulatoryCompliance/Bundles/{Sid} | 
[**FetchBundle**](RegulatoryComplianceBundlesApi.md#FetchBundle) | **Get** /v2/RegulatoryCompliance/Bundles/{Sid} | 
[**ListBundle**](RegulatoryComplianceBundlesApi.md#ListBundle) | **Get** /v2/RegulatoryCompliance/Bundles | 
[**UpdateBundle**](RegulatoryComplianceBundlesApi.md#UpdateBundle) | **Post** /v2/RegulatoryCompliance/Bundles/{Sid} | 



## CreateBundle

> NumbersV2Bundle CreateBundle(ctx, optional)



Create a new Bundle.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateBundleParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**Email** | **string** | The email address that will receive updates when the Bundle resource changes status.
**StatusCallback** | **string** | The URL we call to inform your application of status changes.
**RegulationSid** | **string** | The unique string of a regulation that is associated to the Bundle resource.
**IsoCountry** | **string** | The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the Bundle's phone number country ownership request.
**EndUserType** | **string** | 
**NumberType** | **string** | The type of phone number of the Bundle's ownership request. Can be `local`, `mobile`, `national`, or `toll free`.

### Return type

[**NumbersV2Bundle**](NumbersV2Bundle.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBundle

> DeleteBundle(ctx, Sid)



Delete a specific Bundle.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Bundle resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteBundleParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchBundle

> NumbersV2Bundle FetchBundle(ctx, Sid)



Fetch a specific Bundle instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Bundle resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchBundleParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NumbersV2Bundle**](NumbersV2Bundle.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBundle

> []NumbersV2Bundle ListBundle(ctx, optional)



Retrieve a list of all Bundles for an account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListBundleParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | **string** | The verification status of the Bundle resource. Please refer to [Bundle Statuses](https://www.twilio.com/docs/phone-numbers/regulatory/api/bundles#bundle-statuses) for more details.
**FriendlyName** | **string** | The string that you assigned to describe the resource. The column can contain 255 variable characters.
**RegulationSid** | **string** | The unique string of a [Regulation resource](https://www.twilio.com/docs/phone-numbers/regulatory/api/regulations) that is associated to the Bundle resource.
**IsoCountry** | **string** | The 2-digit [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the Bundle's phone number country ownership request.
**NumberType** | **string** | The type of phone number of the Bundle's ownership request. Can be `local`, `mobile`, `national`, or `tollfree`.
**HasValidUntilDate** | **bool** | Indicates that the Bundle is a valid Bundle until a specified expiration date.
**SortBy** | **string** | Can be `valid-until` or `date-updated`. Defaults to `date-created`.
**SortDirection** | **string** | Default is `DESC`. Can be `ASC` or `DESC`.
**ValidUntilDate** | **time.Time** | Date to filter Bundles having their `valid_until_date` before or after the specified date. Can be `ValidUntilDate>=` or `ValidUntilDate<=`. Both can be used in conjunction as well. [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) is the acceptable date format.
**ValidUntilDateBefore** | **time.Time** | Date to filter Bundles having their `valid_until_date` before or after the specified date. Can be `ValidUntilDate>=` or `ValidUntilDate<=`. Both can be used in conjunction as well. [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) is the acceptable date format.
**ValidUntilDateAfter** | **time.Time** | Date to filter Bundles having their `valid_until_date` before or after the specified date. Can be `ValidUntilDate>=` or `ValidUntilDate<=`. Both can be used in conjunction as well. [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) is the acceptable date format.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]NumbersV2Bundle**](NumbersV2Bundle.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBundle

> NumbersV2Bundle UpdateBundle(ctx, Sidoptional)



Updates a Bundle in an account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Bundle resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateBundleParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | **string** | 
**StatusCallback** | **string** | The URL we call to inform your application of status changes.
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**Email** | **string** | The email address that will receive updates when the Bundle resource changes status.

### Return type

[**NumbersV2Bundle**](NumbersV2Bundle.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

