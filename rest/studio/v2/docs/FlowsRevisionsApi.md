# FlowsRevisionsApi

All URIs are relative to *https://studio.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchFlowRevision**](FlowsRevisionsApi.md#FetchFlowRevision) | **Get** /v2/Flows/{Sid}/Revisions/{Revision} | 
[**ListFlowRevision**](FlowsRevisionsApi.md#ListFlowRevision) | **Get** /v2/Flows/{Sid}/Revisions | 



## FetchFlowRevision

> StudioV2FlowRevision FetchFlowRevision(ctx, SidRevision)



Retrieve a specific Flow revision.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flow resource to fetch.
**Revision** | **string** | Specific Revision number or can be `LatestPublished` and `LatestRevision`.

### Other Parameters

Other parameters are passed through a pointer to a FetchFlowRevisionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**StudioV2FlowRevision**](StudioV2FlowRevision.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFlowRevision

> []StudioV2FlowRevision ListFlowRevision(ctx, Sidoptional)



Retrieve a list of all Flows revisions.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flow resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a ListFlowRevisionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]StudioV2FlowRevision**](StudioV2FlowRevision.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

