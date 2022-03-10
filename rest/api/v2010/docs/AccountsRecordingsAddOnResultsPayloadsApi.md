# AccountsRecordingsAddOnResultsPayloadsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRecordingAddOnResultPayload**](AccountsRecordingsAddOnResultsPayloadsApi.md#DeleteRecordingAddOnResultPayload) | **Delete** /2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{AddOnResultSid}/Payloads/{Sid}.json | 
[**FetchRecordingAddOnResultPayload**](AccountsRecordingsAddOnResultsPayloadsApi.md#FetchRecordingAddOnResultPayload) | **Get** /2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{AddOnResultSid}/Payloads/{Sid}.json | 
[**ListRecordingAddOnResultPayload**](AccountsRecordingsAddOnResultsPayloadsApi.md#ListRecordingAddOnResultPayload) | **Get** /2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{AddOnResultSid}/Payloads.json | 



## DeleteRecordingAddOnResultPayload

> DeleteRecordingAddOnResultPayload(ctx, ReferenceSidAddOnResultSidSidoptional)



Delete a payload from the result along with all associated Data

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ReferenceSid** | **string** | The SID of the recording to which the AddOnResult resource that contains the payloads to delete belongs.
**AddOnResultSid** | **string** | The SID of the AddOnResult to which the payloads to delete belongs.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Recording AddOnResult Payload resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteRecordingAddOnResultPayloadParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resources to delete.

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


## FetchRecordingAddOnResultPayload

> ApiV2010RecordingAddOnResultPayload FetchRecordingAddOnResultPayload(ctx, ReferenceSidAddOnResultSidSidoptional)



Fetch an instance of a result payload

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ReferenceSid** | **string** | The SID of the recording to which the AddOnResult resource that contains the payload to fetch belongs.
**AddOnResultSid** | **string** | The SID of the AddOnResult to which the payload to fetch belongs.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Recording AddOnResult Payload resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRecordingAddOnResultPayloadParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resource to fetch.

### Return type

[**ApiV2010RecordingAddOnResultPayload**](ApiV2010RecordingAddOnResultPayload.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecordingAddOnResultPayload

> []ApiV2010RecordingAddOnResultPayload ListRecordingAddOnResultPayload(ctx, ReferenceSidAddOnResultSidoptional)



Retrieve a list of payloads belonging to the AddOnResult

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ReferenceSid** | **string** | The SID of the recording to which the AddOnResult resource that contains the payloads to read belongs.
**AddOnResultSid** | **string** | The SID of the AddOnResult to which the payloads to read belongs.

### Other Parameters

Other parameters are passed through a pointer to a ListRecordingAddOnResultPayloadParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resources to read.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010RecordingAddOnResultPayload**](ApiV2010RecordingAddOnResultPayload.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

