# ServicesEnvironmentsApi

All URIs are relative to *https://serverless.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvironment**](ServicesEnvironmentsApi.md#CreateEnvironment) | **Post** /v1/Services/{ServiceSid}/Environments | 
[**DeleteEnvironment**](ServicesEnvironmentsApi.md#DeleteEnvironment) | **Delete** /v1/Services/{ServiceSid}/Environments/{Sid} | 
[**FetchEnvironment**](ServicesEnvironmentsApi.md#FetchEnvironment) | **Get** /v1/Services/{ServiceSid}/Environments/{Sid} | 
[**ListEnvironment**](ServicesEnvironmentsApi.md#ListEnvironment) | **Get** /v1/Services/{ServiceSid}/Environments | 



## CreateEnvironment

> ServerlessV1Environment CreateEnvironment(ctx, ServiceSidoptional)



Create a new environment.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to create the Environment resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateEnvironmentParams struct


Name | Type | Description
------------- | ------------- | -------------
**UniqueName** | **string** | A user-defined string that uniquely identifies the Environment resource. It can be a maximum of 100 characters.
**DomainSuffix** | **string** | A URL-friendly name that represents the environment and forms part of the domain name. It can be a maximum of 16 characters.

### Return type

[**ServerlessV1Environment**](ServerlessV1Environment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironment

> DeleteEnvironment(ctx, ServiceSidSid)



Delete a specific environment.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to delete the Environment resource from.
**Sid** | **string** | The SID of the Environment resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteEnvironmentParams struct


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


## FetchEnvironment

> ServerlessV1Environment FetchEnvironment(ctx, ServiceSidSid)



Retrieve a specific environment.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Environment resource from.
**Sid** | **string** | The SID of the Environment resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchEnvironmentParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ServerlessV1Environment**](ServerlessV1Environment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnvironment

> []ServerlessV1Environment ListEnvironment(ctx, ServiceSidoptional)



Retrieve a list of all environments.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to read the Environment resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListEnvironmentParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ServerlessV1Environment**](ServerlessV1Environment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

