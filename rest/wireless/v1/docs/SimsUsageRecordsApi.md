# SimsUsageRecordsApi

All URIs are relative to *https://wireless.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListUsageRecord**](SimsUsageRecordsApi.md#ListUsageRecord) | **Get** /v1/Sims/{SimSid}/UsageRecords | 



## ListUsageRecord

> []WirelessV1UsageRecord ListUsageRecord(ctx, SimSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SimSid** | **string** | The SID of the [Sim resource](https://www.twilio.com/docs/wireless/api/sim-resource)  to read the usage from.

### Other Parameters

Other parameters are passed through a pointer to a ListUsageRecordParams struct


Name | Type | Description
------------- | ------------- | -------------
**End** | **time.Time** | Only include usage that occurred on or before this date, specified in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html). The default is the current time.
**Start** | **time.Time** | Only include usage that has occurred on or after this date, specified in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html). The default is one month before the `end` parameter value.
**Granularity** | **string** | How to summarize the usage by time. Can be: `daily`, `hourly`, or `all`. The default is `all`. A value of `all` returns one Usage Record that describes the usage for the entire period.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]WirelessV1UsageRecord**](WirelessV1UsageRecord.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

