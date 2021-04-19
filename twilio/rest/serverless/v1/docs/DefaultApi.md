# \DefaultApi

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

> ServerlessV1ServiceAsset CreateAsset(ctx, ServiceSid).FriendlyName(FriendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to create the Asset resource under.
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the Asset resource. It can be a maximum of 255 characters. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateAsset(context.Background(), ServiceSid).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAsset`: ServerlessV1ServiceAsset
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to create the Asset resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateAssetParams struct via the builder pattern


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

> ServerlessV1ServiceBuild CreateBuild(ctx, ServiceSid).AssetVersions(AssetVersions).Dependencies(Dependencies).FunctionVersions(FunctionVersions).Runtime(Runtime).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to create the Build resource under.
    AssetVersions := []string{"Inner_example"} // []string | The list of Asset Version resource SIDs to include in the Build. (optional)
    Dependencies := "Dependencies_example" // string | A list of objects that describe the Dependencies included in the Build. Each object contains the `name` and `version` of the dependency. (optional)
    FunctionVersions := []string{"Inner_example"} // []string | The list of the Function Version resource SIDs to include in the Build. (optional)
    Runtime := "Runtime_example" // string | The Runtime version that will be used to run the Build resource when it is deployed. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateBuild(context.Background(), ServiceSid).AssetVersions(AssetVersions).Dependencies(Dependencies).FunctionVersions(FunctionVersions).Runtime(Runtime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuild`: ServerlessV1ServiceBuild
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateBuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to create the Build resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateBuildParams struct via the builder pattern


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

> ServerlessV1ServiceEnvironmentDeployment CreateDeployment(ctx, ServiceSid, EnvironmentSid).BuildSid(BuildSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to create the Deployment resource under.
    EnvironmentSid := "EnvironmentSid_example" // string | The SID of the Environment for the Deployment.
    BuildSid := "BuildSid_example" // string | The SID of the Build for the Deployment. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateDeployment(context.Background(), ServiceSid, EnvironmentSid).BuildSid(BuildSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeployment`: ServerlessV1ServiceEnvironmentDeployment
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to create the Deployment resource under.
**EnvironmentSid** | **string** | The SID of the Environment for the Deployment.

### Other Parameters

Other parameters are passed through a pointer to a CreateDeploymentParams struct via the builder pattern


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

> ServerlessV1ServiceEnvironment CreateEnvironment(ctx, ServiceSid).DomainSuffix(DomainSuffix).UniqueName(UniqueName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to create the Environment resource under.
    DomainSuffix := "DomainSuffix_example" // string | A URL-friendly name that represents the environment and forms part of the domain name. It can be a maximum of 16 characters. (optional)
    UniqueName := "UniqueName_example" // string | A user-defined string that uniquely identifies the Environment resource. It can be a maximum of 100 characters. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateEnvironment(context.Background(), ServiceSid).DomainSuffix(DomainSuffix).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnvironment`: ServerlessV1ServiceEnvironment
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to create the Environment resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateEnvironmentParams struct via the builder pattern


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

> ServerlessV1ServiceFunction CreateFunction(ctx, ServiceSid).FriendlyName(FriendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to create the Function resource under.
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the Function resource. It can be a maximum of 255 characters. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateFunction(context.Background(), ServiceSid).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateFunction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFunction`: ServerlessV1ServiceFunction
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to create the Function resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateFunctionParams struct via the builder pattern


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

> ServerlessV1Service CreateService(ctx).FriendlyName(FriendlyName).IncludeCredentials(IncludeCredentials).UiEditable(UiEditable).UniqueName(UniqueName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the Service resource. It can be a maximum of 255 characters. (optional)
    IncludeCredentials := true // bool | Whether to inject Account credentials into a function invocation context. The default value is `true`. (optional)
    UiEditable := true // bool | Whether the Service's properties and subresources can be edited via the UI. The default value is `false`. (optional)
    UniqueName := "UniqueName_example" // string | A user-defined string that uniquely identifies the Service resource. It can be used as an alternative to the `sid` in the URL path to address the Service resource. This value must be 50 characters or less in length and be unique. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateService(context.Background()).FriendlyName(FriendlyName).IncludeCredentials(IncludeCredentials).UiEditable(UiEditable).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateService`: ServerlessV1Service
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateService`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceParams struct via the builder pattern


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

> ServerlessV1ServiceEnvironmentVariable CreateVariable(ctx, ServiceSid, EnvironmentSid).Key(Key).Value(Value).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to create the Variable resource under.
    EnvironmentSid := "EnvironmentSid_example" // string | The SID of the Environment in which the Variable resource exists.
    Key := "Key_example" // string | A string by which the Variable resource can be referenced. It can be a maximum of 128 characters. (optional)
    Value := "Value_example" // string | A string that contains the actual value of the Variable. It can be a maximum of 450 bytes in size. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateVariable(context.Background(), ServiceSid, EnvironmentSid).Key(Key).Value(Value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVariable`: ServerlessV1ServiceEnvironmentVariable
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to create the Variable resource under.
**EnvironmentSid** | **string** | The SID of the Environment in which the Variable resource exists.

### Other Parameters

Other parameters are passed through a pointer to a CreateVariableParams struct via the builder pattern


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

> DeleteAsset(ctx, ServiceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to delete the Asset resource from.
    Sid := "Sid_example" // string | The SID that identifies the Asset resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteAsset(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to delete the Asset resource from.
**Sid** | **string** | The SID that identifies the Asset resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteAssetParams struct via the builder pattern


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

> DeleteBuild(ctx, ServiceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to delete the Build resource from.
    Sid := "Sid_example" // string | The SID of the Build resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteBuild(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to delete the Build resource from.
**Sid** | **string** | The SID of the Build resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteBuildParams struct via the builder pattern


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

> DeleteEnvironment(ctx, ServiceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to delete the Environment resource from.
    Sid := "Sid_example" // string | The SID of the Environment resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteEnvironment(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to delete the Environment resource from.
**Sid** | **string** | The SID of the Environment resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteEnvironmentParams struct via the builder pattern


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

> DeleteFunction(ctx, ServiceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to delete the Function resource from.
    Sid := "Sid_example" // string | The SID of the Function resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteFunction(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteFunction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to delete the Function resource from.
**Sid** | **string** | The SID of the Function resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteFunctionParams struct via the builder pattern


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

> DeleteService(ctx, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The `sid` or `unique_name` of the Service resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteService(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The &#x60;sid&#x60; or &#x60;unique_name&#x60; of the Service resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceParams struct via the builder pattern


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

> DeleteVariable(ctx, ServiceSid, EnvironmentSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to delete the Variable resource from.
    EnvironmentSid := "EnvironmentSid_example" // string | The SID of the Environment with the Variables to delete.
    Sid := "Sid_example" // string | The SID of the Variable resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteVariable(context.Background(), ServiceSid, EnvironmentSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to delete the Variable resource from.
**EnvironmentSid** | **string** | The SID of the Environment with the Variables to delete.
**Sid** | **string** | The SID of the Variable resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteVariableParams struct via the builder pattern


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

> ServerlessV1ServiceAsset FetchAsset(ctx, ServiceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to fetch the Asset resource from.
    Sid := "Sid_example" // string | The SID that identifies the Asset resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchAsset(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAsset`: ServerlessV1ServiceAsset
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Asset resource from.
**Sid** | **string** | The SID that identifies the Asset resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchAssetParams struct via the builder pattern


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

> ServerlessV1ServiceAssetAssetVersion FetchAssetVersion(ctx, ServiceSid, AssetSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to fetch the Asset Version resource from.
    AssetSid := "AssetSid_example" // string | The SID of the Asset resource that is the parent of the Asset Version resource to fetch.
    Sid := "Sid_example" // string | The SID of the Asset Version resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchAssetVersion(context.Background(), ServiceSid, AssetSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchAssetVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAssetVersion`: ServerlessV1ServiceAssetAssetVersion
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchAssetVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Asset Version resource from.
**AssetSid** | **string** | The SID of the Asset resource that is the parent of the Asset Version resource to fetch.
**Sid** | **string** | The SID of the Asset Version resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchAssetVersionParams struct via the builder pattern


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

> ServerlessV1ServiceBuild FetchBuild(ctx, ServiceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to fetch the Build resource from.
    Sid := "Sid_example" // string | The SID of the Build resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchBuild(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchBuild`: ServerlessV1ServiceBuild
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchBuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Build resource from.
**Sid** | **string** | The SID of the Build resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchBuildParams struct via the builder pattern


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

> ServerlessV1ServiceBuildBuildStatus FetchBuildStatus(ctx, ServiceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to fetch the Build resource from.
    Sid := "Sid_example" // string | The SID of the Build resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchBuildStatus(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchBuildStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchBuildStatus`: ServerlessV1ServiceBuildBuildStatus
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchBuildStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Build resource from.
**Sid** | **string** | The SID of the Build resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchBuildStatusParams struct via the builder pattern


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

> ServerlessV1ServiceEnvironmentDeployment FetchDeployment(ctx, ServiceSid, EnvironmentSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to fetch the Deployment resource from.
    EnvironmentSid := "EnvironmentSid_example" // string | The SID of the Environment used by the Deployment to fetch.
    Sid := "Sid_example" // string | The SID that identifies the Deployment resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchDeployment(context.Background(), ServiceSid, EnvironmentSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchDeployment`: ServerlessV1ServiceEnvironmentDeployment
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Deployment resource from.
**EnvironmentSid** | **string** | The SID of the Environment used by the Deployment to fetch.
**Sid** | **string** | The SID that identifies the Deployment resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchDeploymentParams struct via the builder pattern


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

> ServerlessV1ServiceEnvironment FetchEnvironment(ctx, ServiceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to fetch the Environment resource from.
    Sid := "Sid_example" // string | The SID of the Environment resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchEnvironment(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchEnvironment`: ServerlessV1ServiceEnvironment
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Environment resource from.
**Sid** | **string** | The SID of the Environment resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchEnvironmentParams struct via the builder pattern


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

> ServerlessV1ServiceFunction FetchFunction(ctx, ServiceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to fetch the Function resource from.
    Sid := "Sid_example" // string | The SID of the Function resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchFunction(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchFunction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchFunction`: ServerlessV1ServiceFunction
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Function resource from.
**Sid** | **string** | The SID of the Function resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFunctionParams struct via the builder pattern


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

> ServerlessV1ServiceFunctionFunctionVersion FetchFunctionVersion(ctx, ServiceSid, FunctionSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to fetch the Function Version resource from.
    FunctionSid := "FunctionSid_example" // string | The SID of the function that is the parent of the Function Version resource to fetch.
    Sid := "Sid_example" // string | The SID of the Function Version resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchFunctionVersion(context.Background(), ServiceSid, FunctionSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchFunctionVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchFunctionVersion`: ServerlessV1ServiceFunctionFunctionVersion
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchFunctionVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Function Version resource from.
**FunctionSid** | **string** | The SID of the function that is the parent of the Function Version resource to fetch.
**Sid** | **string** | The SID of the Function Version resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFunctionVersionParams struct via the builder pattern


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

> ServerlessV1ServiceFunctionFunctionVersionFunctionVersionContent FetchFunctionVersionContent(ctx, ServiceSid, FunctionSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to fetch the Function Version content from.
    FunctionSid := "FunctionSid_example" // string | The SID of the Function that is the parent of the Function Version content to fetch.
    Sid := "Sid_example" // string | The SID of the Function Version content to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchFunctionVersionContent(context.Background(), ServiceSid, FunctionSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchFunctionVersionContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchFunctionVersionContent`: ServerlessV1ServiceFunctionFunctionVersionFunctionVersionContent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchFunctionVersionContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Function Version content from.
**FunctionSid** | **string** | The SID of the Function that is the parent of the Function Version content to fetch.
**Sid** | **string** | The SID of the Function Version content to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFunctionVersionContentParams struct via the builder pattern


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

> ServerlessV1ServiceEnvironmentLog FetchLog(ctx, ServiceSid, EnvironmentSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to fetch the Log resource from.
    EnvironmentSid := "EnvironmentSid_example" // string | The SID of the environment with the Log resource to fetch.
    Sid := "Sid_example" // string | The SID of the Log resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchLog(context.Background(), ServiceSid, EnvironmentSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchLog`: ServerlessV1ServiceEnvironmentLog
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Log resource from.
**EnvironmentSid** | **string** | The SID of the environment with the Log resource to fetch.
**Sid** | **string** | The SID of the Log resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchLogParams struct via the builder pattern


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

> ServerlessV1Service FetchService(ctx, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The `sid` or `unique_name` of the Service resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchService(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchService`: ServerlessV1Service
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The &#x60;sid&#x60; or &#x60;unique_name&#x60; of the Service resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceParams struct via the builder pattern


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

> ServerlessV1ServiceEnvironmentVariable FetchVariable(ctx, ServiceSid, EnvironmentSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to fetch the Variable resource from.
    EnvironmentSid := "EnvironmentSid_example" // string | The SID of the Environment with the Variable resource to fetch.
    Sid := "Sid_example" // string | The SID of the Variable resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchVariable(context.Background(), ServiceSid, EnvironmentSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchVariable`: ServerlessV1ServiceEnvironmentVariable
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to fetch the Variable resource from.
**EnvironmentSid** | **string** | The SID of the Environment with the Variable resource to fetch.
**Sid** | **string** | The SID of the Variable resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchVariableParams struct via the builder pattern


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

> ListAssetResponse ListAsset(ctx, ServiceSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to read the Asset resources from.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListAsset(context.Background(), ServiceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAsset`: ListAssetResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to read the Asset resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListAssetParams struct via the builder pattern


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

> ListAssetVersionResponse ListAssetVersion(ctx, ServiceSid, AssetSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to read the Asset Version resource from.
    AssetSid := "AssetSid_example" // string | The SID of the Asset resource that is the parent of the Asset Version resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListAssetVersion(context.Background(), ServiceSid, AssetSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAssetVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAssetVersion`: ListAssetVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAssetVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to read the Asset Version resource from.
**AssetSid** | **string** | The SID of the Asset resource that is the parent of the Asset Version resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListAssetVersionParams struct via the builder pattern


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

> ListBuildResponse ListBuild(ctx, ServiceSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to read the Build resources from.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListBuild(context.Background(), ServiceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListBuild``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuild`: ListBuildResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListBuild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to read the Build resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListBuildParams struct via the builder pattern


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

> ListDeploymentResponse ListDeployment(ctx, ServiceSid, EnvironmentSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to read the Deployment resources from.
    EnvironmentSid := "EnvironmentSid_example" // string | The SID of the Environment used by the Deployment resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListDeployment(context.Background(), ServiceSid, EnvironmentSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeployment`: ListDeploymentResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to read the Deployment resources from.
**EnvironmentSid** | **string** | The SID of the Environment used by the Deployment resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListDeploymentParams struct via the builder pattern


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

> ListEnvironmentResponse ListEnvironment(ctx, ServiceSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to read the Environment resources from.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListEnvironment(context.Background(), ServiceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEnvironment`: ListEnvironmentResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to read the Environment resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListEnvironmentParams struct via the builder pattern


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

> ListFunctionResponse ListFunction(ctx, ServiceSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to read the Function resources from.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListFunction(context.Background(), ServiceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListFunction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFunction`: ListFunctionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to read the Function resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListFunctionParams struct via the builder pattern


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

> ListFunctionVersionResponse ListFunctionVersion(ctx, ServiceSid, FunctionSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to read the Function Version resources from.
    FunctionSid := "FunctionSid_example" // string | The SID of the function that is the parent of the Function Version resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListFunctionVersion(context.Background(), ServiceSid, FunctionSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListFunctionVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFunctionVersion`: ListFunctionVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListFunctionVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to read the Function Version resources from.
**FunctionSid** | **string** | The SID of the function that is the parent of the Function Version resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListFunctionVersionParams struct via the builder pattern


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

> ListLogResponse ListLog(ctx, ServiceSid, EnvironmentSid).FunctionSid(FunctionSid).StartDate(StartDate).EndDate(EndDate).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to read the Log resource from.
    EnvironmentSid := "EnvironmentSid_example" // string | The SID of the environment with the Log resources to read.
    FunctionSid := "FunctionSid_example" // string | The SID of the function whose invocation produced the Log resources to read. (optional)
    StartDate := time.Now() // time.Time | The date/time (in GMT, ISO 8601) after which the Log resources must have been created. Defaults to 1 day prior to current date/time. (optional)
    EndDate := time.Now() // time.Time | The date/time (in GMT, ISO 8601) before which the Log resources must have been created. Defaults to current date/time. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListLog(context.Background(), ServiceSid, EnvironmentSid).FunctionSid(FunctionSid).StartDate(StartDate).EndDate(EndDate).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLog`: ListLogResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to read the Log resource from.
**EnvironmentSid** | **string** | The SID of the environment with the Log resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListLogParams struct via the builder pattern


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

> ListServiceResponse ListService(ctx).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListService(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListService`: ListServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListService`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceParams struct via the builder pattern


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

> ListVariableResponse ListVariable(ctx, ServiceSid, EnvironmentSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to read the Variable resources from.
    EnvironmentSid := "EnvironmentSid_example" // string | The SID of the Environment with the Variable resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListVariable(context.Background(), ServiceSid, EnvironmentSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVariable`: ListVariableResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to read the Variable resources from.
**EnvironmentSid** | **string** | The SID of the Environment with the Variable resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListVariableParams struct via the builder pattern


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

> ServerlessV1ServiceAsset UpdateAsset(ctx, ServiceSid, Sid).FriendlyName(FriendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to update the Asset resource from.
    Sid := "Sid_example" // string | The SID that identifies the Asset resource to update.
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the Asset resource. It can be a maximum of 255 characters. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateAsset(context.Background(), ServiceSid, Sid).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAsset`: ServerlessV1ServiceAsset
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to update the Asset resource from.
**Sid** | **string** | The SID that identifies the Asset resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateAssetParams struct via the builder pattern


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

> ServerlessV1ServiceFunction UpdateFunction(ctx, ServiceSid, Sid).FriendlyName(FriendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to update the Function resource from.
    Sid := "Sid_example" // string | The SID of the Function resource to update.
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the Function resource. It can be a maximum of 255 characters. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateFunction(context.Background(), ServiceSid, Sid).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateFunction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFunction`: ServerlessV1ServiceFunction
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to update the Function resource from.
**Sid** | **string** | The SID of the Function resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateFunctionParams struct via the builder pattern


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

> ServerlessV1Service UpdateService(ctx, Sid).FriendlyName(FriendlyName).IncludeCredentials(IncludeCredentials).UiEditable(UiEditable).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The `sid` or `unique_name` of the Service resource to update.
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the Service resource. It can be a maximum of 255 characters. (optional)
    IncludeCredentials := true // bool | Whether to inject Account credentials into a function invocation context. (optional)
    UiEditable := true // bool | Whether the Service resource's properties and subresources can be edited via the UI. The default value is `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateService(context.Background(), Sid).FriendlyName(FriendlyName).IncludeCredentials(IncludeCredentials).UiEditable(UiEditable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateService`: ServerlessV1Service
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The &#x60;sid&#x60; or &#x60;unique_name&#x60; of the Service resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceParams struct via the builder pattern


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

> ServerlessV1ServiceEnvironmentVariable UpdateVariable(ctx, ServiceSid, EnvironmentSid, Sid).Key(Key).Value(Value).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the Service to update the Variable resource under.
    EnvironmentSid := "EnvironmentSid_example" // string | The SID of the Environment with the Variable resource to update.
    Sid := "Sid_example" // string | The SID of the Variable resource to update.
    Key := "Key_example" // string | A string by which the Variable resource can be referenced. It can be a maximum of 128 characters. (optional)
    Value := "Value_example" // string | A string that contains the actual value of the Variable. It can be a maximum of 450 bytes in size. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateVariable(context.Background(), ServiceSid, EnvironmentSid, Sid).Key(Key).Value(Value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVariable`: ServerlessV1ServiceEnvironmentVariable
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the Service to update the Variable resource under.
**EnvironmentSid** | **string** | The SID of the Environment with the Variable resource to update.
**Sid** | **string** | The SID of the Variable resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateVariableParams struct via the builder pattern


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

