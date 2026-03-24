# VoiceReportsPhoneNumbersInboundApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInboundPhoneNumbersReport**](VoiceReportsPhoneNumbersInboundApi.md#CreateInboundPhoneNumbersReport) | **Post** /v2/Voice/Reports/PhoneNumbers/Inbound | Create Inbound Phone Numbers Level Report
[**ListInboundPhoneNumbersReport**](VoiceReportsPhoneNumbersInboundApi.md#ListInboundPhoneNumbersReport) | **Get** /v2/Voice/Reports/PhoneNumbers/Inbound/{reportId} | List Inbound Phone Numbers Level Reports



## CreateInboundPhoneNumbersReport

> InsightsV2CreateReportResponse CreateInboundPhoneNumbersReport(ctx, optional)

Create Inbound Phone Numbers Level Report

Create Inbound specific Phone Numbers Report for a specific account with given time range.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateInboundPhoneNumbersReportParams struct


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


## ListInboundPhoneNumbersReport

> []InsightsV2InboundPhoneNumberReport ListInboundPhoneNumbersReport(ctx, ReportIdoptional)

List Inbound Phone Numbers Level Reports

Get Inbound Phone Numbers Level Reports for the given Report Id.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ReportId** | **string** | A unique Report Id.

### Other Parameters

Other parameters are passed through a pointer to a ListInboundPhoneNumbersReportParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]InsightsV2InboundPhoneNumberReport**](InsightsV2InboundPhoneNumberReport.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

