# RegulatoryComplianceRegulationsApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchRegulation**](RegulatoryComplianceRegulationsApi.md#FetchRegulation) | **Get** /v2/RegulatoryCompliance/Regulations/{Sid} | 
[**ListRegulation**](RegulatoryComplianceRegulationsApi.md#ListRegulation) | **Get** /v2/RegulatoryCompliance/Regulations | 



## FetchRegulation

> NumbersV2Regulation FetchRegulation(ctx, Sidoptional)



Fetch specific Regulation Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that identifies the Regulation resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchRegulationParams struct


Name | Type | Description
------------- | ------------- | -------------
**IncludeConstraints** | **bool** | A boolean parameter indicating whether to include constraints or not for supporting end user, documents and their fields

### Return type

[**NumbersV2Regulation**](NumbersV2Regulation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRegulation

> []NumbersV2Regulation ListRegulation(ctx, optional)



Retrieve a list of all Regulations.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListRegulationParams struct


Name | Type | Description
------------- | ------------- | -------------
**EndUserType** | **string** | The type of End User the regulation requires - can be `individual` or `business`.
**IsoCountry** | **string** | The ISO country code of the phone number's country.
**NumberType** | **string** | The type of phone number that the regulatory requiremnt is restricting.
**IncludeConstraints** | **bool** | A boolean parameter indicating whether to include constraints or not for supporting end user, documents and their fields
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]NumbersV2Regulation**](NumbersV2Regulation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

