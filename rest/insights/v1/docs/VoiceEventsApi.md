# VoiceEventsApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEvent**](VoiceEventsApi.md#ListEvent) | **Get** /v1/Voice/{CallSid}/Events | Get a list of Call Insight Events for a Call.



## ListEvent

> []InsightsV1Event ListEvent(ctx, CallSidoptional)

Get a list of Call Insight Events for a Call.

Get a list of Call Insight Events for a Call.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | The unique SID identifier of the Call.

### Other Parameters

Other parameters are passed through a pointer to a ListEventParams struct


Name | Type | Description
------------- | ------------- | -------------
**Edge** | **string** | The Edge of this Event. One of `unknown_edge`, `carrier_edge`, `sip_edge`, `sdk_edge` or `client_edge`.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]InsightsV1Event**](InsightsV1Event.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

