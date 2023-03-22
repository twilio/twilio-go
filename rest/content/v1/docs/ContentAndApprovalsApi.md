# ContentAndApprovalsApi

All URIs are relative to *https://content.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListContentAndApprovals**](ContentAndApprovalsApi.md#ListContentAndApprovals) | **Get** /v1/ContentAndApprovals | 



## ListContentAndApprovals

> []ContentV1ContentAndApprovals ListContentAndApprovals(ctx, optional)



Retrieve a list of Contents with approval statuses belonging to the account used to make the request

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListContentAndApprovalsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ContentV1ContentAndApprovals**](ContentV1ContentAndApprovals.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

