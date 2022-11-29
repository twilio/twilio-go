# WorkspacesActivitiesApi

All URIs are relative to *https://taskrouter.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateActivity**](WorkspacesActivitiesApi.md#CreateActivity) | **Post** /v1/Workspaces/{WorkspaceSid}/Activities | 
[**DeleteActivity**](WorkspacesActivitiesApi.md#DeleteActivity) | **Delete** /v1/Workspaces/{WorkspaceSid}/Activities/{Sid} | 
[**FetchActivity**](WorkspacesActivitiesApi.md#FetchActivity) | **Get** /v1/Workspaces/{WorkspaceSid}/Activities/{Sid} | 
[**ListActivity**](WorkspacesActivitiesApi.md#ListActivity) | **Get** /v1/Workspaces/{WorkspaceSid}/Activities | 
[**UpdateActivity**](WorkspacesActivitiesApi.md#UpdateActivity) | **Post** /v1/Workspaces/{WorkspaceSid}/Activities/{Sid} | 



## CreateActivity

> TaskrouterV1Activity CreateActivity(ctx, WorkspaceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace that the new Activity belongs to.

### Other Parameters

Other parameters are passed through a pointer to a CreateActivityParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the Activity resource. It can be up to 64 characters long. These names are used to calculate and expose statistics about Workers, and provide visibility into the state of each Worker. Examples of friendly names include: `on-call`, `break`, and `email`.
**Available** | **bool** | Whether the Worker should be eligible to receive a Task when it occupies the Activity. A value of `true`, `1`, or `yes` specifies the Activity is available. All other values specify that it is not. The value cannot be changed after the Activity is created.

### Return type

[**TaskrouterV1Activity**](TaskrouterV1Activity.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteActivity

> DeleteActivity(ctx, WorkspaceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Activity resources to delete.
**Sid** | **string** | The SID of the Activity resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteActivityParams struct


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


## FetchActivity

> TaskrouterV1Activity FetchActivity(ctx, WorkspaceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Activity resources to fetch.
**Sid** | **string** | The SID of the Activity resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchActivityParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TaskrouterV1Activity**](TaskrouterV1Activity.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActivity

> []TaskrouterV1Activity ListActivity(ctx, WorkspaceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Activity resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListActivityParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | The `friendly_name` of the Activity resources to read.
**Available** | **string** | Whether return only Activity resources that are available or unavailable. A value of `true` returns only available activities. Values of '1' or `yes` also indicate `true`. All other values represent `false` and return activities that are unavailable.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TaskrouterV1Activity**](TaskrouterV1Activity.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateActivity

> TaskrouterV1Activity UpdateActivity(ctx, WorkspaceSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Activity resources to update.
**Sid** | **string** | The SID of the Activity resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateActivityParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the Activity resource. It can be up to 64 characters long. These names are used to calculate and expose statistics about Workers, and provide visibility into the state of each Worker. Examples of friendly names include: `on-call`, `break`, and `email`.

### Return type

[**TaskrouterV1Activity**](TaskrouterV1Activity.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

