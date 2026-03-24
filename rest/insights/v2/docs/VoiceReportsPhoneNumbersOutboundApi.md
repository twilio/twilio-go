# VoiceReportsPhoneNumbersOutboundApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOutboundPhoneNumbersReport**](VoiceReportsPhoneNumbersOutboundApi.md#CreateOutboundPhoneNumbersReport) | **Post** /v2/Voice/Reports/PhoneNumbers/Outbound | Create Outbound Phone Numbers Level Report
[**ListOutboundPhoneNumbersReport**](VoiceReportsPhoneNumbersOutboundApi.md#ListOutboundPhoneNumbersReport) | **Get** /v2/Voice/Reports/PhoneNumbers/Outbound/{reportId} | List Outbound Phone Numbers Level Reports



## CreateOutboundPhoneNumbersReport

> InsightsV2CreateReportResponse CreateOutboundPhoneNumbersReport(ctx, optional)

Create Outbound Phone Numbers Level Report

Create Outbound specific Phone Numbers Report for a specific account with given time range.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateOutboundPhoneNumbersReportParams struct


Name | Type | Description
------------- | ------------- | -------------
**InsightsV2CreatePhoneNumbersReportRequest** | [**InsightsV2CreatePhoneNumbersReportRequest**](InsightsV2CreatePhoneNumbersReportRequest.md) | 

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


## ListOutboundPhoneNumbersReport

> []InsightsV2OutboundPhoneNumberReport ListOutboundPhoneNumbersReport(ctx, ReportIdoptional)

List Outbound Phone Numbers Level Reports

Get Outbound Phone Numbers Level Report for the given Report Id.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ReportId** | **string** | A unique Report Id.

### Other Parameters

Other parameters are passed through a pointer to a ListOutboundPhoneNumbersReportParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]InsightsV2OutboundPhoneNumberReport**](InsightsV2OutboundPhoneNumberReport.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

