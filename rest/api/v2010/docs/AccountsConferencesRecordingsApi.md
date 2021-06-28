# AccountsConferencesRecordingsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConferenceRecording**](AccountsConferencesRecordingsApi.md#DeleteConferenceRecording) | **Delete** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Recordings/{Sid}.json | 
[**FetchConferenceRecording**](AccountsConferencesRecordingsApi.md#FetchConferenceRecording) | **Get** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Recordings/{Sid}.json | 
[**ListConferenceRecording**](AccountsConferencesRecordingsApi.md#ListConferenceRecording) | **Get** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Recordings.json | 
[**UpdateConferenceRecording**](AccountsConferencesRecordingsApi.md#UpdateConferenceRecording) | **Post** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Recordings/{Sid}.json | 



## DeleteConferenceRecording

> DeleteConferenceRecording(ctx, ConferenceSidSidoptional)



Delete a recording from your account

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConferenceSid** | **string** | The Conference SID that identifies the conference associated with the recording to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Conference Recording resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteConferenceRecordingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resources to delete.

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


## FetchConferenceRecording

> ApiV2010AccountConferenceConferenceRecording FetchConferenceRecording(ctx, ConferenceSidSidoptional)



Fetch an instance of a recording for a call

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConferenceSid** | **string** | The Conference SID that identifies the conference associated with the recording to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Conference Recording resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchConferenceRecordingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resource to fetch.

### Return type

[**ApiV2010AccountConferenceConferenceRecording**](ApiV2010AccountConferenceConferenceRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConferenceRecording

> ListConferenceRecordingResponse ListConferenceRecording(ctx, ConferenceSidoptional)



Retrieve a list of recordings belonging to the call used to make the request

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConferenceSid** | **string** | The Conference SID that identifies the conference associated with the recording to read.

### Other Parameters

Other parameters are passed through a pointer to a ListConferenceRecordingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resources to read.
**DateCreated** | **string** | The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. You can also specify inequality: &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60; will return recordings generated at or before midnight on a given date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; returns recordings generated at or after midnight on a date.
**DateCreatedBefore** | **string** | The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. You can also specify inequality: &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60; will return recordings generated at or before midnight on a given date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; returns recordings generated at or after midnight on a date.
**DateCreatedAfter** | **string** | The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. You can also specify inequality: &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60; will return recordings generated at or before midnight on a given date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; returns recordings generated at or after midnight on a date.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListConferenceRecordingResponse**](ListConferenceRecordingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConferenceRecording

> ApiV2010AccountConferenceConferenceRecording UpdateConferenceRecording(ctx, ConferenceSidSidoptional)



Changes the status of the recording to paused, stopped, or in-progress. Note: To use `Twilio.CURRENT`, pass it as recording sid.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConferenceSid** | **string** | The Conference SID that identifies the conference associated with the recording to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Conference Recording resource to update. Use &#x60;Twilio.CURRENT&#x60; to reference the current active recording.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConferenceRecordingParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resource to update.
**PauseBehavior** | **string** | Whether to record during a pause. Can be: &#x60;skip&#x60; or &#x60;silence&#x60; and the default is &#x60;silence&#x60;. &#x60;skip&#x60; does not record during the pause period, while &#x60;silence&#x60; will replace the actual audio of the call with silence during the pause period. This parameter only applies when setting &#x60;status&#x60; is set to &#x60;paused&#x60;.
**Status** | **string** | The new status of the recording. Can be: &#x60;stopped&#x60;, &#x60;paused&#x60;, &#x60;in-progress&#x60;.

### Return type

[**ApiV2010AccountConferenceConferenceRecording**](ApiV2010AccountConferenceConferenceRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

