# \DefaultApi

All URIs are relative to *http://localhost*

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

> ServerlessV1ServiceAsset CreateAsset(ctx, serviceSid, optional)



Create a new Asset resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to create the Asset resource under. | 
 **optional** | ***CreateAssetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAssetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **friendlyName** | **optional.String**| A descriptive string that you create to describe the Asset resource. It can be a maximum of 255 characters. | 

### Return type

[**ServerlessV1ServiceAsset**](serverless.v1.service.asset.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBuild

> ServerlessV1ServiceBuild CreateBuild(ctx, serviceSid, optional)



Create a new Build resource. At least one function version or asset version is required.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to create the Build resource under. | 
 **optional** | ***CreateBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateBuildOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetVersions** | [**optional.Interface of []string**](string.md)| The list of Asset Version resource SIDs to include in the Build. | 
 **dependencies** | **optional.String**| A list of objects that describe the Dependencies included in the Build. Each object contains the &#x60;name&#x60; and &#x60;version&#x60; of the dependency. | 
 **functionVersions** | [**optional.Interface of []string**](string.md)| The list of the Function Version resource SIDs to include in the Build. | 

### Return type

[**ServerlessV1ServiceBuild**](serverless.v1.service.build.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeployment

> ServerlessV1ServiceEnvironmentDeployment CreateDeployment(ctx, serviceSid, environmentSid, optional)



Create a new Deployment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to create the Deployment resource under. | 
**environmentSid** | **string**| The SID of the Environment for the Deployment. | 
 **optional** | ***CreateDeploymentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDeploymentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **buildSid** | **optional.String**| The SID of the Build for the Deployment. | 

### Return type

[**ServerlessV1ServiceEnvironmentDeployment**](serverless.v1.service.environment.deployment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEnvironment

> ServerlessV1ServiceEnvironment CreateEnvironment(ctx, serviceSid, optional)



Create a new environment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to create the Environment resource under. | 
 **optional** | ***CreateEnvironmentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateEnvironmentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domainSuffix** | **optional.String**| A URL-friendly name that represents the environment and forms part of the domain name. It can be a maximum of 16 characters. | 
 **uniqueName** | **optional.String**| A user-defined string that uniquely identifies the Environment resource. It can be a maximum of 100 characters. | 

### Return type

[**ServerlessV1ServiceEnvironment**](serverless.v1.service.environment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFunction

> ServerlessV1ServiceFunction CreateFunction(ctx, serviceSid, optional)



Create a new Function resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to create the Function resource under. | 
 **optional** | ***CreateFunctionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFunctionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **friendlyName** | **optional.String**| A descriptive string that you create to describe the Function resource. It can be a maximum of 255 characters. | 

### Return type

[**ServerlessV1ServiceFunction**](serverless.v1.service.function.md)

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

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the Service resource. It can be a maximum of 255 characters. | 
 **includeCredentials** | **optional.Bool**| Whether to inject Account credentials into a function invocation context. The default value is &#x60;true&#x60;. | 
 **uiEditable** | **optional.Bool**| Whether the Service&#39;s properties and subresources can be edited via the UI. The default value is &#x60;false&#x60;. | 
 **uniqueName** | **optional.String**| A user-defined string that uniquely identifies the Service resource. It can be used as an alternative to the &#x60;sid&#x60; in the URL path to address the Service resource. This value must be 50 characters or less in length and be unique. | 

### Return type

[**ServerlessV1Service**](serverless.v1.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVariable

> ServerlessV1ServiceEnvironmentVariable CreateVariable(ctx, serviceSid, environmentSid, optional)



Create a new Variable.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to create the Variable resource under. | 
**environmentSid** | **string**| The SID of the Environment in which the Variable resource exists. | 
 **optional** | ***CreateVariableOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateVariableOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **key** | **optional.String**| A string by which the Variable resource can be referenced. It can be a maximum of 128 characters. | 
 **value** | **optional.String**| A string that contains the actual value of the Variable. It can be a maximum of 450 bytes in size. | 

### Return type

[**ServerlessV1ServiceEnvironmentVariable**](serverless.v1.service.environment.variable.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAsset

> DeleteAsset(ctx, serviceSid, sid)



Delete an Asset resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to delete the Asset resource from. | 
**sid** | **string**| The SID that identifies the Asset resource to delete. | 

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

> DeleteBuild(ctx, serviceSid, sid)



Delete a Build resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to delete the Build resource from. | 
**sid** | **string**| The SID of the Build resource to delete. | 

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

> DeleteEnvironment(ctx, serviceSid, sid)



Delete a specific environment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to delete the Environment resource from. | 
**sid** | **string**| The SID of the Environment resource to delete. | 

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

> DeleteFunction(ctx, serviceSid, sid)



Delete a Function resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to delete the Function resource from. | 
**sid** | **string**| The SID of the Function resource to delete. | 

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

> DeleteService(ctx, sid)



Delete a Service resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The &#x60;sid&#x60; or &#x60;unique_name&#x60; of the Service resource to delete. | 

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

> DeleteVariable(ctx, serviceSid, environmentSid, sid)



Delete a specific Variable.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to delete the Variable resource from. | 
**environmentSid** | **string**| The SID of the Environment with the Variables to delete. | 
**sid** | **string**| The SID of the Variable resource to delete. | 

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

> ServerlessV1ServiceAsset FetchAsset(ctx, serviceSid, sid)



Retrieve a specific Asset resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to fetch the Asset resource from. | 
**sid** | **string**| The SID that identifies the Asset resource to fetch. | 

### Return type

[**ServerlessV1ServiceAsset**](serverless.v1.service.asset.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAssetVersion

> ServerlessV1ServiceAssetAssetVersion FetchAssetVersion(ctx, serviceSid, assetSid, sid)



Retrieve a specific Asset Version.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to fetch the Asset Version resource from. | 
**assetSid** | **string**| The SID of the Asset resource that is the parent of the Asset Version resource to fetch. | 
**sid** | **string**| The SID of the Asset Version resource to fetch. | 

### Return type

[**ServerlessV1ServiceAssetAssetVersion**](serverless.v1.service.asset.asset_version.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchBuild

> ServerlessV1ServiceBuild FetchBuild(ctx, serviceSid, sid)



Retrieve a specific Build resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to fetch the Build resource from. | 
**sid** | **string**| The SID of the Build resource to fetch. | 

### Return type

[**ServerlessV1ServiceBuild**](serverless.v1.service.build.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchBuildStatus

> ServerlessV1ServiceBuildBuildStatus FetchBuildStatus(ctx, serviceSid, sid)



Retrieve a specific Build resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to fetch the Build resource from. | 
**sid** | **string**| The SID of the Build resource to fetch. | 

### Return type

[**ServerlessV1ServiceBuildBuildStatus**](serverless.v1.service.build.build_status.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDeployment

> ServerlessV1ServiceEnvironmentDeployment FetchDeployment(ctx, serviceSid, environmentSid, sid)



Retrieve a specific Deployment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to fetch the Deployment resource from. | 
**environmentSid** | **string**| The SID of the Environment used by the Deployment to fetch. | 
**sid** | **string**| The SID that identifies the Deployment resource to fetch. | 

### Return type

[**ServerlessV1ServiceEnvironmentDeployment**](serverless.v1.service.environment.deployment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchEnvironment

> ServerlessV1ServiceEnvironment FetchEnvironment(ctx, serviceSid, sid)



Retrieve a specific environment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to fetch the Environment resource from. | 
**sid** | **string**| The SID of the Environment resource to fetch. | 

### Return type

[**ServerlessV1ServiceEnvironment**](serverless.v1.service.environment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFunction

> ServerlessV1ServiceFunction FetchFunction(ctx, serviceSid, sid)



Retrieve a specific Function resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to fetch the Function resource from. | 
**sid** | **string**| The SID of the Function resource to fetch. | 

### Return type

[**ServerlessV1ServiceFunction**](serverless.v1.service.function.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFunctionVersion

> ServerlessV1ServiceFunctionFunctionVersion FetchFunctionVersion(ctx, serviceSid, functionSid, sid)



Retrieve a specific Function Version resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to fetch the Function Version resource from. | 
**functionSid** | **string**| The SID of the function that is the parent of the Function Version resource to fetch. | 
**sid** | **string**| The SID of the Function Version resource to fetch. | 

### Return type

[**ServerlessV1ServiceFunctionFunctionVersion**](serverless.v1.service.function.function_version.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFunctionVersionContent

> ServerlessV1ServiceFunctionFunctionVersionFunctionVersionContent FetchFunctionVersionContent(ctx, serviceSid, functionSid, sid)



Retrieve a the content of a specific Function Version resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to fetch the Function Version content from. | 
**functionSid** | **string**| The SID of the Function that is the parent of the Function Version content to fetch. | 
**sid** | **string**| The SID of the Function Version content to fetch. | 

### Return type

[**ServerlessV1ServiceFunctionFunctionVersionFunctionVersionContent**](serverless.v1.service.function.function_version.function_version_content.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchLog

> ServerlessV1ServiceEnvironmentLog FetchLog(ctx, serviceSid, environmentSid, sid)



Retrieve a specific log.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to fetch the Log resource from. | 
**environmentSid** | **string**| The SID of the environment with the Log resource to fetch. | 
**sid** | **string**| The SID of the Log resource to fetch. | 

### Return type

[**ServerlessV1ServiceEnvironmentLog**](serverless.v1.service.environment.log.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> ServerlessV1Service FetchService(ctx, sid)



Retrieve a specific Service resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The &#x60;sid&#x60; or &#x60;unique_name&#x60; of the Service resource to fetch. | 

### Return type

[**ServerlessV1Service**](serverless.v1.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchVariable

> ServerlessV1ServiceEnvironmentVariable FetchVariable(ctx, serviceSid, environmentSid, sid)



Retrieve a specific Variable.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to fetch the Variable resource from. | 
**environmentSid** | **string**| The SID of the Environment with the Variable resource to fetch. | 
**sid** | **string**| The SID of the Variable resource to fetch. | 

### Return type

[**ServerlessV1ServiceEnvironmentVariable**](serverless.v1.service.environment.variable.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAsset

> ServerlessV1ServiceAssetReadResponse ListAsset(ctx, serviceSid, optional)



Retrieve a list of all Assets.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to read the Asset resources from. | 
 **optional** | ***ListAssetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAssetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ServerlessV1ServiceAssetReadResponse**](serverless_v1_service_assetReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssetVersion

> ServerlessV1ServiceAssetAssetVersionReadResponse ListAssetVersion(ctx, serviceSid, assetSid, optional)



Retrieve a list of all Asset Versions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to read the Asset Version resource from. | 
**assetSid** | **string**| The SID of the Asset resource that is the parent of the Asset Version resources to read. | 
 **optional** | ***ListAssetVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAssetVersionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ServerlessV1ServiceAssetAssetVersionReadResponse**](serverless_v1_service_asset_asset_versionReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuild

> ServerlessV1ServiceBuildReadResponse ListBuild(ctx, serviceSid, optional)



Retrieve a list of all Builds.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to read the Build resources from. | 
 **optional** | ***ListBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListBuildOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ServerlessV1ServiceBuildReadResponse**](serverless_v1_service_buildReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeployment

> ServerlessV1ServiceEnvironmentDeploymentReadResponse ListDeployment(ctx, serviceSid, environmentSid, optional)



Retrieve a list of all Deployments.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to read the Deployment resources from. | 
**environmentSid** | **string**| The SID of the Environment used by the Deployment resources to read. | 
 **optional** | ***ListDeploymentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDeploymentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ServerlessV1ServiceEnvironmentDeploymentReadResponse**](serverless_v1_service_environment_deploymentReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnvironment

> ServerlessV1ServiceEnvironmentReadResponse ListEnvironment(ctx, serviceSid, optional)



Retrieve a list of all environments.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to read the Environment resources from. | 
 **optional** | ***ListEnvironmentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListEnvironmentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ServerlessV1ServiceEnvironmentReadResponse**](serverless_v1_service_environmentReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFunction

> ServerlessV1ServiceFunctionReadResponse ListFunction(ctx, serviceSid, optional)



Retrieve a list of all Functions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to read the Function resources from. | 
 **optional** | ***ListFunctionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFunctionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ServerlessV1ServiceFunctionReadResponse**](serverless_v1_service_functionReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFunctionVersion

> ServerlessV1ServiceFunctionFunctionVersionReadResponse ListFunctionVersion(ctx, serviceSid, functionSid, optional)



Retrieve a list of all Function Version resources.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to read the Function Version resources from. | 
**functionSid** | **string**| The SID of the function that is the parent of the Function Version resources to read. | 
 **optional** | ***ListFunctionVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFunctionVersionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ServerlessV1ServiceFunctionFunctionVersionReadResponse**](serverless_v1_service_function_function_versionReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLog

> ServerlessV1ServiceEnvironmentLogReadResponse ListLog(ctx, serviceSid, environmentSid, optional)



Retrieve a list of all logs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to read the Log resource from. | 
**environmentSid** | **string**| The SID of the environment with the Log resources to read. | 
 **optional** | ***ListLogOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListLogOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **functionSid** | **optional.String**| The SID of the function whose invocation produced the Log resources to read. | 
 **startDate** | **optional.Time**| The date/time (in GMT, ISO 8601) after which the Log resources must have been created. Defaults to 1 day prior to current date/time. | 
 **endDate** | **optional.Time**| The date/time (in GMT, ISO 8601) before which the Log resources must have been created. Defaults to current date/time. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ServerlessV1ServiceEnvironmentLogReadResponse**](serverless_v1_service_environment_logReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListService

> ServerlessV1ServiceReadResponse ListService(ctx, optional)



Retrieve a list of all Services.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ServerlessV1ServiceReadResponse**](serverless_v1_serviceReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVariable

> ServerlessV1ServiceEnvironmentVariableReadResponse ListVariable(ctx, serviceSid, environmentSid, optional)



Retrieve a list of all Variables.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to read the Variable resources from. | 
**environmentSid** | **string**| The SID of the Environment with the Variable resources to read. | 
 **optional** | ***ListVariableOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListVariableOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ServerlessV1ServiceEnvironmentVariableReadResponse**](serverless_v1_service_environment_variableReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAsset

> ServerlessV1ServiceAsset UpdateAsset(ctx, serviceSid, sid, optional)



Update a specific Asset resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to update the Asset resource from. | 
**sid** | **string**| The SID that identifies the Asset resource to update. | 
 **optional** | ***UpdateAssetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAssetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **friendlyName** | **optional.String**| A descriptive string that you create to describe the Asset resource. It can be a maximum of 255 characters. | 

### Return type

[**ServerlessV1ServiceAsset**](serverless.v1.service.asset.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFunction

> ServerlessV1ServiceFunction UpdateFunction(ctx, serviceSid, sid, optional)



Update a specific Function resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to update the Function resource from. | 
**sid** | **string**| The SID of the Function resource to update. | 
 **optional** | ***UpdateFunctionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateFunctionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **friendlyName** | **optional.String**| A descriptive string that you create to describe the Function resource. It can be a maximum of 255 characters. | 

### Return type

[**ServerlessV1ServiceFunction**](serverless.v1.service.function.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> ServerlessV1Service UpdateService(ctx, sid, optional)



Update a specific Service resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The &#x60;sid&#x60; or &#x60;unique_name&#x60; of the Service resource to update. | 
 **optional** | ***UpdateServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **friendlyName** | **optional.String**| A descriptive string that you create to describe the Service resource. It can be a maximum of 255 characters. | 
 **includeCredentials** | **optional.Bool**| Whether to inject Account credentials into a function invocation context. | 
 **uiEditable** | **optional.Bool**| Whether the Service resource&#39;s properties and subresources can be edited via the UI. The default value is &#x60;false&#x60;. | 

### Return type

[**ServerlessV1Service**](serverless.v1.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVariable

> ServerlessV1ServiceEnvironmentVariable UpdateVariable(ctx, serviceSid, environmentSid, sid, optional)



Update a specific Variable.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the Service to update the Variable resource under. | 
**environmentSid** | **string**| The SID of the Environment with the Variable resource to update. | 
**sid** | **string**| The SID of the Variable resource to update. | 
 **optional** | ***UpdateVariableOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateVariableOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **key** | **optional.String**| A string by which the Variable resource can be referenced. It can be a maximum of 128 characters. | 
 **value** | **optional.String**| A string that contains the actual value of the Variable. It can be a maximum of 450 bytes in size. | 

### Return type

[**ServerlessV1ServiceEnvironmentVariable**](serverless.v1.service.environment.variable.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

