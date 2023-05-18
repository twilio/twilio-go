# InsightsSegmentsApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListInsightsSegments**](InsightsSegmentsApi.md#ListInsightsSegments) | **Get** /v1/Insights/Segments | 



## ListInsightsSegments

> []FlexV1InsightsSegments ListInsightsSegments(ctx, optional)



To get segments for given reservation Ids

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListInsightsSegmentsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Authorization** | **string** | The Authorization HTTP request header
**SegmentId** | **string** | To unique id of the segment
**ReservationId** | **[]string** | The list of reservation Ids
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]FlexV1InsightsSegments**](FlexV1InsightsSegments.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

