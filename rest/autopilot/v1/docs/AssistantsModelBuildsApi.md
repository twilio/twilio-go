# AssistantsModelBuildsApi

All URIs are relative to *https://autopilot.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateModelBuild**](AssistantsModelBuildsApi.md#CreateModelBuild) | **Post** /v1/Assistants/{AssistantSid}/ModelBuilds | 
[**DeleteModelBuild**](AssistantsModelBuildsApi.md#DeleteModelBuild) | **Delete** /v1/Assistants/{AssistantSid}/ModelBuilds/{Sid} | 
[**FetchModelBuild**](AssistantsModelBuildsApi.md#FetchModelBuild) | **Get** /v1/Assistants/{AssistantSid}/ModelBuilds/{Sid} | 
[**ListModelBuild**](AssistantsModelBuildsApi.md#ListModelBuild) | **Get** /v1/Assistants/{AssistantSid}/ModelBuilds | 
[**UpdateModelBuild**](AssistantsModelBuildsApi.md#UpdateModelBuild) | **Post** /v1/Assistants/{AssistantSid}/ModelBuilds/{Sid} | 



## CreateModelBuild

> AutopilotV1ModelBuild CreateModelBuild(ctx, AssistantSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the new resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateModelBuildParams struct


Name | Type | Description
------------- | ------------- | -------------
**StatusCallback** | **string** | The URL we should call using a POST method to send status information to your application.
**UniqueName** | **string** | An application-defined string that uniquely identifies the new resource. This value must be a unique string of no more than 64 characters. It can be used as an alternative to the `sid` in the URL path to address the resource.

### Return type

[**AutopilotV1ModelBuild**](AutopilotV1ModelBuild.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteModelBuild

> DeleteModelBuild(ctx, AssistantSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ModelBuild resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteModelBuildParams struct


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


## FetchModelBuild

> AutopilotV1ModelBuild FetchModelBuild(ctx, AssistantSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ModelBuild resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchModelBuildParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AutopilotV1ModelBuild**](AutopilotV1ModelBuild.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListModelBuild

> []AutopilotV1ModelBuild ListModelBuild(ctx, AssistantSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListModelBuildParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]AutopilotV1ModelBuild**](AutopilotV1ModelBuild.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateModelBuild

> AutopilotV1ModelBuild UpdateModelBuild(ctx, AssistantSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ModelBuild resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateModelBuildParams struct


Name | Type | Description
------------- | ------------- | -------------
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. This value must be a unique string of no more than 64 characters. It can be used as an alternative to the `sid` in the URL path to address the resource.

### Return type

[**AutopilotV1ModelBuild**](AutopilotV1ModelBuild.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

