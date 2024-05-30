# ContentApi

All URIs are relative to *https://content.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListContent**](ContentApi.md#ListContent) | **Get** /v2/Content | 



## ListContent

> []ContentV1Content ListContent(ctx, optional)



Retrieve a list of Contents belonging to the account used to make the request

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListContentParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ContentV1Content**](ContentV1Content.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

