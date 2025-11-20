# ContentApi

All URIs are relative to *https://content.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContent**](ContentApi.md#CreateContent) | **Post** /v1/Content | Create Content
[**DeleteContent**](ContentApi.md#DeleteContent) | **Delete** /v1/Content/{Sid} | Delete Content
[**FetchContent**](ContentApi.md#FetchContent) | **Get** /v1/Content/{Sid} | Fetch Content
[**ListContent**](ContentApi.md#ListContent) | **Get** /v1/Content | List Contents
[**UpdateContent**](ContentApi.md#UpdateContent) | **Put** /v1/Content/{Sid} | Update Content



## CreateContent

> ContentV1Content CreateContent(ctx, optional)

Create Content

Create a Content resource

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateContentParams struct


Name | Type | Description
------------- | ------------- | -------------
**ContentCreateRequest** | [**ContentCreateRequest**](ContentCreateRequest.md) | 

### Return type

[**ContentV1Content**](ContentV1Content.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContent

> DeleteContent(ctx, Sid)

Delete Content

Deletes a Content resource

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Content resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a DeleteContentParams struct


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


## FetchContent

> ContentV1Content FetchContent(ctx, Sid)

Fetch Content

Fetch a Content resource by its unique Content Sid

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Content resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchContentParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ContentV1Content**](ContentV1Content.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContent

> []ContentV1Content ListContent(ctx, optional)

List Contents

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


## UpdateContent

> ContentV1Content UpdateContent(ctx, Sidoptional)

Update Content

Update a Content resource

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Content resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateContentParams struct


Name | Type | Description
------------- | ------------- | -------------
**ContentUpdateRequest** | [**ContentUpdateRequest**](ContentUpdateRequest.md) | 

### Return type

[**ContentV1Content**](ContentV1Content.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

