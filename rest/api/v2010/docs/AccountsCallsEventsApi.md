# AccountsCallsEventsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCallEvent**](AccountsCallsEventsApi.md#ListCallEvent) | **Get** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Events.json | 



## ListCallEvent

> []ApiV2010CallEvent ListCallEvent(ctx, CallSidoptional)



Retrieve a list of all events for a call.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | The unique SID identifier of the Call.

### Other Parameters

Other parameters are passed through a pointer to a ListCallEventParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique SID identifier of the Account.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010CallEvent**](ApiV2010CallEvent.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

