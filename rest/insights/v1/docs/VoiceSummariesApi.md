# VoiceSummariesApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCallSummaries**](VoiceSummariesApi.md#ListCallSummaries) | **Get** /v1/Voice/Summaries | 



## ListCallSummaries

> []InsightsV1CallSummaries ListCallSummaries(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListCallSummariesParams struct


Name | Type | Description
------------- | ------------- | -------------
**From** | **string** | 
**To** | **string** | 
**FromCarrier** | **string** | 
**ToCarrier** | **string** | 
**FromCountryCode** | **string** | 
**ToCountryCode** | **string** | 
**Branded** | **bool** | 
**VerifiedCaller** | **bool** | 
**HasTag** | **bool** | 
**StartTime** | **string** | 
**EndTime** | **string** | 
**CallType** | **string** | 
**CallState** | **string** | 
**Direction** | **string** | 
**ProcessingState** | **string** | 
**SortBy** | **string** | 
**Subaccount** | **string** | 
**AbnormalSession** | **bool** | 
**AnsweredBy** | **string** | 
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]InsightsV1CallSummaries**](InsightsV1CallSummaries.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

