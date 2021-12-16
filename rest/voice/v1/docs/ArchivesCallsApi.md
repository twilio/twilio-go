# ArchivesCallsApi

All URIs are relative to *https://voice.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteArchivedCall**](ArchivesCallsApi.md#DeleteArchivedCall) | **Delete** /v1/Archives/{Date}/Calls/{Sid} | 



## DeleteArchivedCall

> DeleteArchivedCall(ctx, DateSid)



Delete an archived call record from Bulk Export. Note: this does not also delete the record from the Voice API.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Date** | **string** | The date of the Call in UTC.
**Sid** | **string** | The Twilio-provided Call SID that uniquely identifies the Call resource to delete

### Other Parameters

Other parameters are passed through a pointer to a DeleteArchivedCallParams struct


Name | Type | Description
------------- | ------------- | -------------

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

