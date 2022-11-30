# ContentApprovalRequestsApi

All URIs are relative to *https://content.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchApprovalFetch**](ContentApprovalRequestsApi.md#FetchApprovalFetch) | **Get** /v1/Content/{Sid}/ApprovalRequests | 



## FetchApprovalFetch

> ContentV1ApprovalFetch FetchApprovalFetch(ctx, Sid)



Fetch a Content resource's approval status by its unique Content Sid

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Content resource whose approval information to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchApprovalFetchParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ContentV1ApprovalFetch**](ContentV1ApprovalFetch.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

