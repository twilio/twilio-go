# InsightsSegmentsApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchInsightsSegments**](InsightsSegmentsApi.md#FetchInsightsSegments) | **Get** /v1/Insights/Segments/{SegmentId} | 



## FetchInsightsSegments

> FlexV1InsightsSegments FetchInsightsSegments(ctx, SegmentIdoptional)



To get the Segments of an Account

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SegmentId** | **string** | To unique id of the segment

### Other Parameters

Other parameters are passed through a pointer to a FetchInsightsSegmentsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Token** | **string** | The Token HTTP request header

### Return type

[**FlexV1InsightsSegments**](FlexV1InsightsSegments.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

