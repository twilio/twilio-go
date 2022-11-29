# ServicesBuildsApi

All URIs are relative to *https://serverless.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBuild**](ServicesBuildsApi.md#CreateBuild) | **Post** /v1/Services/{ServiceSid}/Builds | 
[**DeleteBuild**](ServicesBuildsApi.md#DeleteBuild) | **Delete** /v1/Services/{ServiceSid}/Builds/{Sid} | 
[**FetchBuild**](ServicesBuildsApi.md#FetchBuild) | **Get** /v1/Services/{ServiceSid}/Builds/{Sid} | 
[**ListBuild**](ServicesBuildsApi.md#ListBuild) | **Get** /v1/Services/{ServiceSid}/Builds | 



## CreateBuild

> ServerlessV1Build CreateBuild(ctx, ServiceSidoptional)



Create a new Build resource. At least one function version or asset version is required.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to create the Build resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateBuildParams struct


Name | Type | Description
------------- | ------------- | -------------
**AssetVersions** | **[]string** | The list of Asset Version resource SIDs to include in the Build.
**FunctionVersions** | **[]string** | The list of the Function Version resource SIDs to include in the Build.
**Dependencies** | **string** | A list of objects that describe the Dependencies included in the Build. Each object contains the `name` and `version` of the dependency.
**Runtime** | **string** | The Runtime version that will be used to run the Build resource when it is deployed.

### Return type

[**ServerlessV1Build**](ServerlessV1Build.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBuild

> DeleteBuild(ctx, ServiceSidSid)



Delete a Build resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to delete the Build resource from.
**Sid** | **string** | The SID of the Build resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteBuildParams struct


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


## FetchBuild

> ServerlessV1Build FetchBuild(ctx, ServiceSidSid)



Retrieve a specific Build resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Build resource from.
**Sid** | **string** | The SID of the Build resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchBuildParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ServerlessV1Build**](ServerlessV1Build.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuild

> []ServerlessV1Build ListBuild(ctx, ServiceSidoptional)



Retrieve a list of all Builds.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to read the Build resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListBuildParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ServerlessV1Build**](ServerlessV1Build.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

