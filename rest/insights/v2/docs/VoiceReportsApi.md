# VoiceReportsApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountReport**](VoiceReportsApi.md#CreateAccountReport) | **Post** /v2/Voice/Reports | Create Account Level Report
[**FetchAccountReport**](VoiceReportsApi.md#FetchAccountReport) | **Get** /v2/Voice/Reports/{reportId} | Fetch Voice Insights Account Level Report



## CreateAccountReport

> InsightsV2CreateReportResponse CreateAccountReport(ctx, optional)

Create Account Level Report

Create a Report for a specific account with given time range and filters.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateAccountReportParams struct


Name | Type | Description
------------- | ------------- | -------------
**InsightsV2CreateAccountReportRequest** | [**InsightsV2CreateAccountReportRequest**](InsightsV2CreateAccountReportRequest.md) | 

### Return type

[**InsightsV2CreateReportResponse**](InsightsV2CreateReportResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAccountReport

> InsightsV2AccountReport FetchAccountReport(ctx, ReportId)

Fetch Voice Insights Account Level Report

Get Account Level Report for the given Report Id.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ReportId** | **string** | A unique request id.

### Other Parameters

Other parameters are passed through a pointer to a FetchAccountReportParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**InsightsV2AccountReport**](InsightsV2AccountReport.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

