# ServicesEnvironmentsDeploymentsApi

All URIs are relative to *https://serverless.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeployment**](ServicesEnvironmentsDeploymentsApi.md#CreateDeployment) | **Post** /v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Deployments | 
[**FetchDeployment**](ServicesEnvironmentsDeploymentsApi.md#FetchDeployment) | **Get** /v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Deployments/{Sid} | 
[**ListDeployment**](ServicesEnvironmentsDeploymentsApi.md#ListDeployment) | **Get** /v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Deployments | 



## CreateDeployment

> ServerlessV1Deployment CreateDeployment(ctx, ServiceSidEnvironmentSidoptional)



Create a new Deployment.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to create the Deployment resource under.
**EnvironmentSid** | **string** | The SID of the Environment for the Deployment.

### Other Parameters

Other parameters are passed through a pointer to a CreateDeploymentParams struct


Name | Type | Description
------------- | ------------- | -------------
**BuildSid** | **string** | The SID of the Build for the Deployment.

### Return type

[**ServerlessV1Deployment**](ServerlessV1Deployment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDeployment

> ServerlessV1Deployment FetchDeployment(ctx, ServiceSidEnvironmentSidSid)



Retrieve a specific Deployment.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Deployment resource from.
**EnvironmentSid** | **string** | The SID of the Environment used by the Deployment to fetch.
**Sid** | **string** | The SID that identifies the Deployment resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchDeploymentParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ServerlessV1Deployment**](ServerlessV1Deployment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeployment

> []ServerlessV1Deployment ListDeployment(ctx, ServiceSidEnvironmentSidoptional)



Retrieve a list of all Deployments.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to read the Deployment resources from.
**EnvironmentSid** | **string** | The SID of the Environment used by the Deployment resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListDeploymentParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ServerlessV1Deployment**](ServerlessV1Deployment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

