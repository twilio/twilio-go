# UsageRecordsApi

All URIs are relative to *https://supersim.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListUsageRecord**](UsageRecordsApi.md#ListUsageRecord) | **Get** /v1/UsageRecords | 



## ListUsageRecord

> []SupersimV1UsageRecord ListUsageRecord(ctx, optional)



List UsageRecords

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListUsageRecordParams struct


Name | Type | Description
------------- | ------------- | -------------
**Sim** | **string** | SID or unique name of a Sim resource. Only show UsageRecords representing usage incurred by this Super SIM.
**Fleet** | **string** | SID or unique name of a Fleet resource. Only show UsageRecords representing usage for Super SIMs belonging to this Fleet resource at the time the usage occurred.
**Network** | **string** | SID of a Network resource. Only show UsageRecords representing usage on this network.
**IsoCountry** | **string** | Alpha-2 ISO Country Code. Only show UsageRecords representing usage in this country.
**Group** | **string** | Dimension over which to aggregate usage records. Can be: `sim`, `fleet`, `network`, `isoCountry`. Default is to not aggregate across any of these dimensions, UsageRecords will be aggregated into the time buckets described by the `Granularity` parameter.
**Granularity** | **string** | Time-based grouping that UsageRecords should be aggregated by. Can be: `hour`, `day`, or `all`. Default is `all`. `all` returns one UsageRecord that describes the usage for the entire period.
**StartTime** | **time.Time** | Only include usage that occurred at or after this time, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Default is one month before the `end_time`.
**EndTime** | **time.Time** | Only include usage that occurred before this time (exclusive), specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Default is the current time.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]SupersimV1UsageRecord**](SupersimV1UsageRecord.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

