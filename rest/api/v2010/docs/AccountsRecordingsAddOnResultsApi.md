# AccountsRecordingsAddOnResultsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRecordingAddOnResult**](AccountsRecordingsAddOnResultsApi.md#DeleteRecordingAddOnResult) | **Delete** /2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{Sid}.json | 
[**FetchRecordingAddOnResult**](AccountsRecordingsAddOnResultsApi.md#FetchRecordingAddOnResult) | **Get** /2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{Sid}.json | 
[**ListRecordingAddOnResult**](AccountsRecordingsAddOnResultsApi.md#ListRecordingAddOnResult) | **Get** /2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults.json | 



## DeleteRecordingAddOnResult

> DeleteRecordingAddOnResult(ctx, ReferenceSidSidoptional)



Delete a result and purge all associated Payloads

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ReferenceSid** | **string** | The SID of the recording to which the result to delete belongs.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Recording AddOnResult resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteRecordingAddOnResultParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult resources to delete.

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRecordingAddOnResult

> ApiV2010RecordingAddOnResult FetchRecordingAddOnResult(ctx, ReferenceSidSidoptional)



Fetch an instance of an AddOnResult

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ReferenceSid** | **string** | The SID of the recording to which the result to fetch belongs.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Recording AddOnResult resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRecordingAddOnResultParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult resource to fetch.

### Return type

[**ApiV2010RecordingAddOnResult**](ApiV2010RecordingAddOnResult.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecordingAddOnResult

> []ApiV2010RecordingAddOnResult ListRecordingAddOnResult(ctx, ReferenceSidoptional)



Retrieve a list of results belonging to the recording

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ReferenceSid** | **string** | The SID of the recording to which the result to read belongs.

### Other Parameters

Other parameters are passed through a pointer to a ListRecordingAddOnResultParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult resources to read.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010RecordingAddOnResult**](ApiV2010RecordingAddOnResult.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

