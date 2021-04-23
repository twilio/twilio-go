# DefaultApi

All URIs are relative to *https://serverless.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAsset**](DefaultApi.md#CreateAsset) | **Post** /v1/Services/{ServiceSid}/Assets | 
[**CreateBuild**](DefaultApi.md#CreateBuild) | **Post** /v1/Services/{ServiceSid}/Builds | 
[**CreateDeployment**](DefaultApi.md#CreateDeployment) | **Post** /v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Deployments | 
[**CreateEnvironment**](DefaultApi.md#CreateEnvironment) | **Post** /v1/Services/{ServiceSid}/Environments | 
[**CreateFunction**](DefaultApi.md#CreateFunction) | **Post** /v1/Services/{ServiceSid}/Functions | 
[**CreateService**](DefaultApi.md#CreateService) | **Post** /v1/Services | 
[**CreateVariable**](DefaultApi.md#CreateVariable) | **Post** /v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Variables | 
[**DeleteAsset**](DefaultApi.md#DeleteAsset) | **Delete** /v1/Services/{ServiceSid}/Assets/{Sid} | 
[**DeleteBuild**](DefaultApi.md#DeleteBuild) | **Delete** /v1/Services/{ServiceSid}/Builds/{Sid} | 
[**DeleteEnvironment**](DefaultApi.md#DeleteEnvironment) | **Delete** /v1/Services/{ServiceSid}/Environments/{Sid} | 
[**DeleteFunction**](DefaultApi.md#DeleteFunction) | **Delete** /v1/Services/{ServiceSid}/Functions/{Sid} | 
[**DeleteService**](DefaultApi.md#DeleteService) | **Delete** /v1/Services/{Sid} | 
[**DeleteVariable**](DefaultApi.md#DeleteVariable) | **Delete** /v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Variables/{Sid} | 
[**FetchAsset**](DefaultApi.md#FetchAsset) | **Get** /v1/Services/{ServiceSid}/Assets/{Sid} | 
[**FetchAssetVersion**](DefaultApi.md#FetchAssetVersion) | **Get** /v1/Services/{ServiceSid}/Assets/{AssetSid}/Versions/{Sid} | 
[**FetchBuild**](DefaultApi.md#FetchBuild) | **Get** /v1/Services/{ServiceSid}/Builds/{Sid} | 
[**FetchBuildStatus**](DefaultApi.md#FetchBuildStatus) | **Get** /v1/Services/{ServiceSid}/Builds/{Sid}/Status | 
[**FetchDeployment**](DefaultApi.md#FetchDeployment) | **Get** /v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Deployments/{Sid} | 
[**FetchEnvironment**](DefaultApi.md#FetchEnvironment) | **Get** /v1/Services/{ServiceSid}/Environments/{Sid} | 
[**FetchFunction**](DefaultApi.md#FetchFunction) | **Get** /v1/Services/{ServiceSid}/Functions/{Sid} | 
[**FetchFunctionVersion**](DefaultApi.md#FetchFunctionVersion) | **Get** /v1/Services/{ServiceSid}/Functions/{FunctionSid}/Versions/{Sid} | 
[**FetchFunctionVersionContent**](DefaultApi.md#FetchFunctionVersionContent) | **Get** /v1/Services/{ServiceSid}/Functions/{FunctionSid}/Versions/{Sid}/Content | 
[**FetchLog**](DefaultApi.md#FetchLog) | **Get** /v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Logs/{Sid} | 
[**FetchService**](DefaultApi.md#FetchService) | **Get** /v1/Services/{Sid} | 
[**FetchVariable**](DefaultApi.md#FetchVariable) | **Get** /v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Variables/{Sid} | 
[**ListAsset**](DefaultApi.md#ListAsset) | **Get** /v1/Services/{ServiceSid}/Assets | 
[**ListAssetVersion**](DefaultApi.md#ListAssetVersion) | **Get** /v1/Services/{ServiceSid}/Assets/{AssetSid}/Versions | 
[**ListBuild**](DefaultApi.md#ListBuild) | **Get** /v1/Services/{ServiceSid}/Builds | 
[**ListDeployment**](DefaultApi.md#ListDeployment) | **Get** /v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Deployments | 
[**ListEnvironment**](DefaultApi.md#ListEnvironment) | **Get** /v1/Services/{ServiceSid}/Environments | 
[**ListFunction**](DefaultApi.md#ListFunction) | **Get** /v1/Services/{ServiceSid}/Functions | 
[**ListFunctionVersion**](DefaultApi.md#ListFunctionVersion) | **Get** /v1/Services/{ServiceSid}/Functions/{FunctionSid}/Versions | 
[**ListLog**](DefaultApi.md#ListLog) | **Get** /v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Logs | 
[**ListService**](DefaultApi.md#ListService) | **Get** /v1/Services | 
[**ListVariable**](DefaultApi.md#ListVariable) | **Get** /v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Variables | 
[**UpdateAsset**](DefaultApi.md#UpdateAsset) | **Post** /v1/Services/{ServiceSid}/Assets/{Sid} | 
[**UpdateFunction**](DefaultApi.md#UpdateFunction) | **Post** /v1/Services/{ServiceSid}/Functions/{Sid} | 
[**UpdateService**](DefaultApi.md#UpdateService) | **Post** /v1/Services/{Sid} | 
[**UpdateVariable**](DefaultApi.md#UpdateVariable) | **Post** /v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Variables/{Sid} | 



## CreateAsset

> ServerlessV1ServiceAsset CreateAsset(ctx, ServiceSidoptional)



Create a new Asset resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to create the Asset resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateAssetParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the Asset resource. It can be a maximum of 255 characters.

### Return type

[**ServerlessV1ServiceAsset**](ServerlessV1ServiceAsset.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBuild

> ServerlessV1ServiceBuild CreateBuild(ctx, ServiceSidoptional)



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
**Dependencies** | **string** | A list of objects that describe the Dependencies included in the Build. Each object contains the &#x60;name&#x60; and &#x60;version&#x60; of the dependency.
**FunctionVersions** | **[]string** | The list of the Function Version resource SIDs to include in the Build.
**Runtime** | **string** | The Runtime version that will be used to run the Build resource when it is deployed.

### Return type

[**ServerlessV1ServiceBuild**](ServerlessV1ServiceBuild.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeployment

> ServerlessV1ServiceEnvironmentDeployment CreateDeployment(ctx, ServiceSidEnvironmentSidoptional)



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

[**ServerlessV1ServiceEnvironmentDeployment**](ServerlessV1ServiceEnvironmentDeployment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEnvironment

> ServerlessV1ServiceEnvironment CreateEnvironment(ctx, ServiceSidoptional)



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
**DomainSuffix** | **string** | A URL-friendly name that represents the environment and forms part of the domain name. It can be a maximum of 16 characters.
**UniqueName** | **string** | A user-defined string that uniquely identifies the Environment resource. It can be a maximum of 100 characters.

### Return type

[**ServerlessV1ServiceEnvironment**](ServerlessV1ServiceEnvironment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFunction

> ServerlessV1ServiceFunction CreateFunction(ctx, ServiceSidoptional)



Create a new Function resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to create the Function resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateFunctionParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the Function resource. It can be a maximum of 255 characters.

### Return type

[**ServerlessV1ServiceFunction**](ServerlessV1ServiceFunction.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateService

> ServerlessV1Service CreateService(ctx, optional)



Create a new Service resource.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the Service resource. It can be a maximum of 255 characters.
**IncludeCredentials** | **bool** | Whether to inject Account credentials into a function invocation context. The default value is &#x60;true&#x60;.
**UiEditable** | **bool** | Whether the Service&#39;s properties and subresources can be edited via the UI. The default value is &#x60;false&#x60;.
**UniqueName** | **string** | A user-defined string that uniquely identifies the Service resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the Service resource. This value must be 50 characters or less in length and be unique.

### Return type

[**ServerlessV1Service**](ServerlessV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVariable

> ServerlessV1ServiceEnvironmentVariable CreateVariable(ctx, ServiceSidEnvironmentSidoptional)



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

[**ServerlessV1ServiceEnvironmentVariable**](ServerlessV1ServiceEnvironmentVariable.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAsset

> DeleteAsset(ctx, ServiceSidSid)



Delete an Asset resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to delete the Asset resource from.
**Sid** | **string** | The SID that identifies the Asset resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteAssetParams struct


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


## DeleteFunction

> DeleteFunction(ctx, ServiceSidSid)



Delete a Function resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to delete the Function resource from.
**Sid** | **string** | The SID of the Function resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteFunctionParams struct


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


## DeleteService

> DeleteService(ctx, Sid)



Delete a Service resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The &#x60;sid&#x60; or &#x60;unique_name&#x60; of the Service resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceParams struct


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


## FetchAsset

> ServerlessV1ServiceAsset FetchAsset(ctx, ServiceSidSid)



Retrieve a specific Asset resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Asset resource from.
**Sid** | **string** | The SID that identifies the Asset resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchAssetParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ServerlessV1ServiceAsset**](ServerlessV1ServiceAsset.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAssetVersion

> ServerlessV1ServiceAssetAssetVersion FetchAssetVersion(ctx, ServiceSidAssetSidSid)



Retrieve a specific Asset Version.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Asset Version resource from.
**AssetSid** | **string** | The SID of the Asset resource that is the parent of the Asset Version resource to fetch.
**Sid** | **string** | The SID of the Asset Version resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchAssetVersionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ServerlessV1ServiceAssetAssetVersion**](ServerlessV1ServiceAssetAssetVersion.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchBuild

> ServerlessV1ServiceBuild FetchBuild(ctx, ServiceSidSid)



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

[**ServerlessV1ServiceBuild**](ServerlessV1ServiceBuild.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchBuildStatus

> ServerlessV1ServiceBuildBuildStatus FetchBuildStatus(ctx, ServiceSidSid)



Retrieve a specific Build resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Build resource from.
**Sid** | **string** | The SID of the Build resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchBuildStatusParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ServerlessV1ServiceBuildBuildStatus**](ServerlessV1ServiceBuildBuildStatus.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDeployment

> ServerlessV1ServiceEnvironmentDeployment FetchDeployment(ctx, ServiceSidEnvironmentSidSid)



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

[**ServerlessV1ServiceEnvironmentDeployment**](ServerlessV1ServiceEnvironmentDeployment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchEnvironment

> ServerlessV1ServiceEnvironment FetchEnvironment(ctx, ServiceSidSid)



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

[**ServerlessV1ServiceEnvironment**](ServerlessV1ServiceEnvironment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFunction

> ServerlessV1ServiceFunction FetchFunction(ctx, ServiceSidSid)



Retrieve a specific Function resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Function resource from.
**Sid** | **string** | The SID of the Function resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFunctionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ServerlessV1ServiceFunction**](ServerlessV1ServiceFunction.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFunctionVersion

> ServerlessV1ServiceFunctionFunctionVersion FetchFunctionVersion(ctx, ServiceSidFunctionSidSid)



Retrieve a specific Function Version resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Function Version resource from.
**FunctionSid** | **string** | The SID of the function that is the parent of the Function Version resource to fetch.
**Sid** | **string** | The SID of the Function Version resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFunctionVersionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ServerlessV1ServiceFunctionFunctionVersion**](ServerlessV1ServiceFunctionFunctionVersion.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFunctionVersionContent

> ServerlessV1ServiceFunctionFunctionVersionFunctionVersionContent FetchFunctionVersionContent(ctx, ServiceSidFunctionSidSid)



Retrieve a the content of a specific Function Version resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Function Version content from.
**FunctionSid** | **string** | The SID of the Function that is the parent of the Function Version content to fetch.
**Sid** | **string** | The SID of the Function Version content to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFunctionVersionContentParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ServerlessV1ServiceFunctionFunctionVersionFunctionVersionContent**](ServerlessV1ServiceFunctionFunctionVersionFunctionVersionContent.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchLog

> ServerlessV1ServiceEnvironmentLog FetchLog(ctx, ServiceSidEnvironmentSidSid)



Retrieve a specific log.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Log resource from.
**EnvironmentSid** | **string** | The SID of the environment with the Log resource to fetch.
**Sid** | **string** | The SID of the Log resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchLogParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ServerlessV1ServiceEnvironmentLog**](ServerlessV1ServiceEnvironmentLog.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> ServerlessV1Service FetchService(ctx, Sid)



Retrieve a specific Service resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The &#x60;sid&#x60; or &#x60;unique_name&#x60; of the Service resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ServerlessV1Service**](ServerlessV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchVariable

> ServerlessV1ServiceEnvironmentVariable FetchVariable(ctx, ServiceSidEnvironmentSidSid)



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

[**ServerlessV1ServiceEnvironmentVariable**](ServerlessV1ServiceEnvironmentVariable.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAsset

> ListAssetResponse ListAsset(ctx, ServiceSidoptional)



Retrieve a list of all Assets.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to read the Asset resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListAssetParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListAssetResponse**](ListAssetResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssetVersion

> ListAssetVersionResponse ListAssetVersion(ctx, ServiceSidAssetSidoptional)



Retrieve a list of all Asset Versions.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to read the Asset Version resource from.
**AssetSid** | **string** | The SID of the Asset resource that is the parent of the Asset Version resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListAssetVersionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListAssetVersionResponse**](ListAssetVersionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuild

> ListBuildResponse ListBuild(ctx, ServiceSidoptional)



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
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListBuildResponse**](ListBuildResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeployment

> ListDeploymentResponse ListDeployment(ctx, ServiceSidEnvironmentSidoptional)



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
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListDeploymentResponse**](ListDeploymentResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnvironment

> ListEnvironmentResponse ListEnvironment(ctx, ServiceSidoptional)



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
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListEnvironmentResponse**](ListEnvironmentResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFunction

> ListFunctionResponse ListFunction(ctx, ServiceSidoptional)



Retrieve a list of all Functions.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to read the Function resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListFunctionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListFunctionResponse**](ListFunctionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFunctionVersion

> ListFunctionVersionResponse ListFunctionVersion(ctx, ServiceSidFunctionSidoptional)



Retrieve a list of all Function Version resources.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to read the Function Version resources from.
**FunctionSid** | **string** | The SID of the function that is the parent of the Function Version resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListFunctionVersionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListFunctionVersionResponse**](ListFunctionVersionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLog

> ListLogResponse ListLog(ctx, ServiceSidEnvironmentSidoptional)



Retrieve a list of all logs.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to read the Log resource from.
**EnvironmentSid** | **string** | The SID of the environment with the Log resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListLogParams struct


Name | Type | Description
------------- | ------------- | -------------
**FunctionSid** | **string** | The SID of the function whose invocation produced the Log resources to read.
**StartDate** | **time.Time** | The date/time (in GMT, ISO 8601) after which the Log resources must have been created. Defaults to 1 day prior to current date/time.
**EndDate** | **time.Time** | The date/time (in GMT, ISO 8601) before which the Log resources must have been created. Defaults to current date/time.
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListLogResponse**](ListLogResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListService

> ListServiceResponse ListService(ctx, optional)



Retrieve a list of all Services.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListServiceResponse**](ListServiceResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVariable

> ListVariableResponse ListVariable(ctx, ServiceSidEnvironmentSidoptional)



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
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListVariableResponse**](ListVariableResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAsset

> ServerlessV1ServiceAsset UpdateAsset(ctx, ServiceSidSidoptional)



Update a specific Asset resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to update the Asset resource from.
**Sid** | **string** | The SID that identifies the Asset resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateAssetParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the Asset resource. It can be a maximum of 255 characters.

### Return type

[**ServerlessV1ServiceAsset**](ServerlessV1ServiceAsset.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFunction

> ServerlessV1ServiceFunction UpdateFunction(ctx, ServiceSidSidoptional)



Update a specific Function resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to update the Function resource from.
**Sid** | **string** | The SID of the Function resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateFunctionParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the Function resource. It can be a maximum of 255 characters.

### Return type

[**ServerlessV1ServiceFunction**](ServerlessV1ServiceFunction.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> ServerlessV1Service UpdateService(ctx, Sidoptional)



Update a specific Service resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The &#x60;sid&#x60; or &#x60;unique_name&#x60; of the Service resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the Service resource. It can be a maximum of 255 characters.
**IncludeCredentials** | **bool** | Whether to inject Account credentials into a function invocation context.
**UiEditable** | **bool** | Whether the Service resource&#39;s properties and subresources can be edited via the UI. The default value is &#x60;false&#x60;.

### Return type

[**ServerlessV1Service**](ServerlessV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVariable

> ServerlessV1ServiceEnvironmentVariable UpdateVariable(ctx, ServiceSidEnvironmentSidSidoptional)



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

[**ServerlessV1ServiceEnvironmentVariable**](ServerlessV1ServiceEnvironmentVariable.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

