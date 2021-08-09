# TrunksRecordingApi

All URIs are relative to *https://trunking.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchRecording**](TrunksRecordingApi.md#FetchRecording) | **Get** /v1/Trunks/{TrunkSid}/Recording | 
[**UpdateRecording**](TrunksRecordingApi.md#UpdateRecording) | **Post** /v1/Trunks/{TrunkSid}/Recording | 



## FetchRecording

> TrunkingV1Recording FetchRecording(ctx, TrunkSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to fetch the recording settings.

### Other Parameters

Other parameters are passed through a pointer to a FetchRecordingParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrunkingV1Recording**](TrunkingV1Recording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRecording

> TrunkingV1Recording UpdateRecording(ctx, TrunkSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk that will have its recording settings updated.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRecordingParams struct


Name | Type | Description
------------- | ------------- | -------------
**Mode** | **string** | The recording mode for the trunk. Can be do-not-record (default), record-from-ringing, record-from-answer, record-from-ringing-dual, or record-from-answer-dual.
**Trim** | **string** | The recording trim setting for the trunk. Can be do-not-trim (default) or trim-silence.

### Return type

[**TrunkingV1Recording**](TrunkingV1Recording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

