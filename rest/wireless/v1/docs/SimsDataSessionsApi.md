# SimsDataSessionsApi

All URIs are relative to *https://wireless.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDataSession**](SimsDataSessionsApi.md#ListDataSession) | **Get** /v1/Sims/{SimSid}/DataSessions | 



## ListDataSession

> []WirelessV1DataSession ListDataSession(ctx, SimSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SimSid** | **string** | The SID of the [Sim resource](https://www.twilio.com/docs/wireless/api/sim-resource) with the Data Sessions to read.

### Other Parameters

Other parameters are passed through a pointer to a ListDataSessionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]WirelessV1DataSession**](WirelessV1DataSession.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

