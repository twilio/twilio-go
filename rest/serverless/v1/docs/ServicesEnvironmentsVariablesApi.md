# ServicesEnvironmentsVariablesApi

All URIs are relative to *https://serverless.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVariable**](ServicesEnvironmentsVariablesApi.md#CreateVariable) | **Post** /v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Variables | 
[**DeleteVariable**](ServicesEnvironmentsVariablesApi.md#DeleteVariable) | **Delete** /v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Variables/{Sid} | 
[**FetchVariable**](ServicesEnvironmentsVariablesApi.md#FetchVariable) | **Get** /v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Variables/{Sid} | 
[**ListVariable**](ServicesEnvironmentsVariablesApi.md#ListVariable) | **Get** /v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Variables | 
[**UpdateVariable**](ServicesEnvironmentsVariablesApi.md#UpdateVariable) | **Post** /v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Variables/{Sid} | 



## CreateVariable

> ServerlessV1Variable CreateVariable(ctx, ServiceSidEnvironmentSidoptional)



Create a new Variable.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to create the Variable resource under.
**EnvironmentSid** | **string** | The SID of the Environment in which the Variable resource exists.

### Other Parameters

Other parameters are passed through a pointer to a CreateVariableParams struct


Name | Type | Description
------------- | ------------- | -------------
**Key** | **string** | A string by which the Variable resource can be referenced. It can be a maximum of 128 characters.
**Value** | **string** | A string that contains the actual value of the Variable. It can be a maximum of 450 bytes in size.

### Return type

[**ServerlessV1Variable**](ServerlessV1Variable.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVariable

> DeleteVariable(ctx, ServiceSidEnvironmentSidSid)



Delete a specific Variable.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to delete the Variable resource from.
**EnvironmentSid** | **string** | The SID of the Environment with the Variables to delete.
**Sid** | **string** | The SID of the Variable resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteVariableParams struct


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


## FetchVariable

> ServerlessV1Variable FetchVariable(ctx, ServiceSidEnvironmentSidSid)



Retrieve a specific Variable.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Variable resource from.
**EnvironmentSid** | **string** | The SID of the Environment with the Variable resource to fetch.
**Sid** | **string** | The SID of the Variable resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchVariableParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ServerlessV1Variable**](ServerlessV1Variable.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVariable

> []ServerlessV1Variable ListVariable(ctx, ServiceSidEnvironmentSidoptional)



Retrieve a list of all Variables.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to read the Variable resources from.
**EnvironmentSid** | **string** | The SID of the Environment with the Variable resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListVariableParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ServerlessV1Variable**](ServerlessV1Variable.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVariable

> ServerlessV1Variable UpdateVariable(ctx, ServiceSidEnvironmentSidSidoptional)



Update a specific Variable.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to update the Variable resource under.
**EnvironmentSid** | **string** | The SID of the Environment with the Variable resource to update.
**Sid** | **string** | The SID of the Variable resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateVariableParams struct


Name | Type | Description
------------- | ------------- | -------------
**Key** | **string** | A string by which the Variable resource can be referenced. It can be a maximum of 128 characters.
**Value** | **string** | A string that contains the actual value of the Variable. It can be a maximum of 450 bytes in size.

### Return type

[**ServerlessV1Variable**](ServerlessV1Variable.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

