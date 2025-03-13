# AccountsRecordingsAddOnResultsPayloadsDataApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchRecordingAddOnResultPayloadData**](AccountsRecordingsAddOnResultsPayloadsDataApi.md#FetchRecordingAddOnResultPayloadData) | **Get** /2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{AddOnResultSid}/Payloads/{PayloadSid}/Data.json | Fetch an instance of a result payload



## FetchRecordingAddOnResultPayloadData

> ApiV2010RecordingAddOnResultPayloadData FetchRecordingAddOnResultPayloadData(ctx, ReferenceSidAddOnResultSidPayloadSidoptional)

Fetch an instance of a result payload

Fetch an instance of a result payload

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ReferenceSid** | **string** | The SID of the recording to which the AddOnResult resource that contains the payload to fetch belongs.
**AddOnResultSid** | **string** | The SID of the AddOnResult to which the payload to fetch belongs.
**PayloadSid** | **string** | The Twilio-provided string that uniquely identifies the Recording AddOnResult Payload resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRecordingAddOnResultPayloadDataParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resource to fetch.

### Return type

[**ApiV2010RecordingAddOnResultPayloadData**](ApiV2010RecordingAddOnResultPayloadData.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

