# \DefaultApi

All URIs are relative to *https://studio.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExecution**](DefaultApi.md#CreateExecution) | **Post** /v2/Flows/{FlowSid}/Executions | 
[**CreateFlow**](DefaultApi.md#CreateFlow) | **Post** /v2/Flows | 
[**DeleteExecution**](DefaultApi.md#DeleteExecution) | **Delete** /v2/Flows/{FlowSid}/Executions/{Sid} | 
[**DeleteFlow**](DefaultApi.md#DeleteFlow) | **Delete** /v2/Flows/{Sid} | 
[**FetchExecution**](DefaultApi.md#FetchExecution) | **Get** /v2/Flows/{FlowSid}/Executions/{Sid} | 
[**FetchExecutionContext**](DefaultApi.md#FetchExecutionContext) | **Get** /v2/Flows/{FlowSid}/Executions/{ExecutionSid}/Context | 
[**FetchExecutionStep**](DefaultApi.md#FetchExecutionStep) | **Get** /v2/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps/{Sid} | 
[**FetchExecutionStepContext**](DefaultApi.md#FetchExecutionStepContext) | **Get** /v2/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps/{StepSid}/Context | 
[**FetchFlow**](DefaultApi.md#FetchFlow) | **Get** /v2/Flows/{Sid} | 
[**FetchFlowRevision**](DefaultApi.md#FetchFlowRevision) | **Get** /v2/Flows/{Sid}/Revisions/{Revision} | 
[**FetchTestUser**](DefaultApi.md#FetchTestUser) | **Get** /v2/Flows/{Sid}/TestUsers | 
[**ListExecution**](DefaultApi.md#ListExecution) | **Get** /v2/Flows/{FlowSid}/Executions | 
[**ListExecutionStep**](DefaultApi.md#ListExecutionStep) | **Get** /v2/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps | 
[**ListFlow**](DefaultApi.md#ListFlow) | **Get** /v2/Flows | 
[**ListFlowRevision**](DefaultApi.md#ListFlowRevision) | **Get** /v2/Flows/{Sid}/Revisions | 
[**UpdateExecution**](DefaultApi.md#UpdateExecution) | **Post** /v2/Flows/{FlowSid}/Executions/{Sid} | 
[**UpdateFlow**](DefaultApi.md#UpdateFlow) | **Post** /v2/Flows/{Sid} | 
[**UpdateFlowValidate**](DefaultApi.md#UpdateFlowValidate) | **Post** /v2/Flows/Validate | 
[**UpdateTestUser**](DefaultApi.md#UpdateTestUser) | **Post** /v2/Flows/{Sid}/TestUsers | 



## CreateExecution

> StudioV2FlowExecution CreateExecution(ctx, FlowSid).From(From).Parameters(Parameters).To(To).Execute()





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
    FlowSid := "FlowSid_example" // string | The SID of the Excecution's Flow.
    From := "From_example" // string | The Twilio phone number to send messages or initiate calls from during the Flow's Execution. Available as variable `{{flow.channel.address}}`. (optional)
    Parameters := TODO // map[string]interface{} | JSON data that will be added to the Flow's context and that can be accessed as variables inside your Flow. For example, if you pass in `Parameters={\\\"name\\\":\\\"Zeke\\\"}`, a widget in your Flow can reference the variable `{{flow.data.name}}`, which returns \\\"Zeke\\\". Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode the JSON string. (optional)
    To := "To_example" // string | The Contact phone number to start a Studio Flow Execution, available as variable `{{contact.channel.address}}`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateExecution(context.Background(), FlowSid).From(From).Parameters(Parameters).To(To).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateExecution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateExecution`: StudioV2FlowExecution
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateExecution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Excecution&#39;s Flow.

### Other Parameters

Other parameters are passed through a pointer to a CreateExecutionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **From** | **string** | The Twilio phone number to send messages or initiate calls from during the Flow&#39;s Execution. Available as variable &#x60;{{flow.channel.address}}&#x60;.
 **Parameters** | [**map[string]interface{}**](map[string]interface{}.md) | JSON data that will be added to the Flow&#39;s context and that can be accessed as variables inside your Flow. For example, if you pass in &#x60;Parameters&#x3D;{\\\&quot;name\\\&quot;:\\\&quot;Zeke\\\&quot;}&#x60;, a widget in your Flow can reference the variable &#x60;{{flow.data.name}}&#x60;, which returns \\\&quot;Zeke\\\&quot;. Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode the JSON string.
 **To** | **string** | The Contact phone number to start a Studio Flow Execution, available as variable &#x60;{{contact.channel.address}}&#x60;.

### Return type

[**StudioV2FlowExecution**](StudioV2FlowExecution.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFlow

> StudioV2Flow CreateFlow(ctx).CommitMessage(CommitMessage).Definition(Definition).FriendlyName(FriendlyName).Status(Status).Execute()





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
    CommitMessage := "CommitMessage_example" // string | Description of change made in the revision. (optional)
    Definition := TODO // map[string]interface{} | JSON representation of flow definition. (optional)
    FriendlyName := "FriendlyName_example" // string | The string that you assigned to describe the Flow. (optional)
    Status := "Status_example" // string | The status of the Flow. Can be: `draft` or `published`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateFlow(context.Background()).CommitMessage(CommitMessage).Definition(Definition).FriendlyName(FriendlyName).Status(Status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFlow`: StudioV2Flow
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateFlow`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateFlowParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **CommitMessage** | **string** | Description of change made in the revision.
 **Definition** | [**map[string]interface{}**](map[string]interface{}.md) | JSON representation of flow definition.
 **FriendlyName** | **string** | The string that you assigned to describe the Flow.
 **Status** | **string** | The status of the Flow. Can be: &#x60;draft&#x60; or &#x60;published&#x60;.

### Return type

[**StudioV2Flow**](StudioV2Flow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExecution

> DeleteExecution(ctx, FlowSid, Sid).Execute()





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
    FlowSid := "FlowSid_example" // string | The SID of the Flow with the Execution resources to delete.
    Sid := "Sid_example" // string | The SID of the Execution resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteExecution(context.Background(), FlowSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteExecution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Execution resources to delete.
**Sid** | **string** | The SID of the Execution resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteExecutionParams struct via the builder pattern


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


## DeleteFlow

> DeleteFlow(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | The SID of the Flow resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteFlow(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flow resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteFlowParams struct via the builder pattern


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


## FetchExecution

> StudioV2FlowExecution FetchExecution(ctx, FlowSid, Sid).Execute()





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
    FlowSid := "FlowSid_example" // string | The SID of the Flow with the Execution resource to fetch
    Sid := "Sid_example" // string | The SID of the Execution resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchExecution(context.Background(), FlowSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchExecution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchExecution`: StudioV2FlowExecution
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchExecution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Execution resource to fetch
**Sid** | **string** | The SID of the Execution resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchExecutionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**StudioV2FlowExecution**](StudioV2FlowExecution.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExecutionContext

> StudioV2FlowExecutionExecutionContext FetchExecutionContext(ctx, FlowSid, ExecutionSid).Execute()





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
    FlowSid := "FlowSid_example" // string | The SID of the Flow with the Execution context to fetch.
    ExecutionSid := "ExecutionSid_example" // string | The SID of the Execution context to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchExecutionContext(context.Background(), FlowSid, ExecutionSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchExecutionContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchExecutionContext`: StudioV2FlowExecutionExecutionContext
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchExecutionContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Execution context to fetch.
**ExecutionSid** | **string** | The SID of the Execution context to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchExecutionContextParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**StudioV2FlowExecutionExecutionContext**](StudioV2FlowExecutionExecutionContext.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExecutionStep

> StudioV2FlowExecutionExecutionStep FetchExecutionStep(ctx, FlowSid, ExecutionSid, Sid).Execute()





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
    FlowSid := "FlowSid_example" // string | The SID of the Flow with the Step to fetch.
    ExecutionSid := "ExecutionSid_example" // string | The SID of the Execution resource with the Step to fetch.
    Sid := "Sid_example" // string | The SID of the ExecutionStep resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchExecutionStep(context.Background(), FlowSid, ExecutionSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchExecutionStep``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchExecutionStep`: StudioV2FlowExecutionExecutionStep
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchExecutionStep`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Step to fetch.
**ExecutionSid** | **string** | The SID of the Execution resource with the Step to fetch.
**Sid** | **string** | The SID of the ExecutionStep resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchExecutionStepParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**StudioV2FlowExecutionExecutionStep**](StudioV2FlowExecutionExecutionStep.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExecutionStepContext

> StudioV2FlowExecutionExecutionStepExecutionStepContext FetchExecutionStepContext(ctx, FlowSid, ExecutionSid, StepSid).Execute()





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
    FlowSid := "FlowSid_example" // string | The SID of the Flow with the Step to fetch.
    ExecutionSid := "ExecutionSid_example" // string | The SID of the Execution resource with the Step to fetch.
    StepSid := "StepSid_example" // string | The SID of the Step to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchExecutionStepContext(context.Background(), FlowSid, ExecutionSid, StepSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchExecutionStepContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchExecutionStepContext`: StudioV2FlowExecutionExecutionStepExecutionStepContext
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchExecutionStepContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Step to fetch.
**ExecutionSid** | **string** | The SID of the Execution resource with the Step to fetch.
**StepSid** | **string** | The SID of the Step to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchExecutionStepContextParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**StudioV2FlowExecutionExecutionStepExecutionStepContext**](StudioV2FlowExecutionExecutionStepExecutionStepContext.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFlow

> StudioV2Flow FetchFlow(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | The SID of the Flow resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchFlow(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchFlow`: StudioV2Flow
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flow resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFlowParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**StudioV2Flow**](StudioV2Flow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFlowRevision

> StudioV2FlowFlowRevision FetchFlowRevision(ctx, Sid, Revision).Execute()





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
    Sid := "Sid_example" // string | The SID of the Flow resource to fetch.
    Revision := "Revision_example" // string | Specific Revision number or can be `LatestPublished` and `LatestRevision`.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchFlowRevision(context.Background(), Sid, Revision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchFlowRevision``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchFlowRevision`: StudioV2FlowFlowRevision
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchFlowRevision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flow resource to fetch.
**Revision** | **string** | Specific Revision number or can be &#x60;LatestPublished&#x60; and &#x60;LatestRevision&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a FetchFlowRevisionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**StudioV2FlowFlowRevision**](StudioV2FlowFlowRevision.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTestUser

> StudioV2FlowTestUser FetchTestUser(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | Unique identifier of the flow.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchTestUser(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchTestUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTestUser`: StudioV2FlowTestUser
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchTestUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | Unique identifier of the flow.

### Other Parameters

Other parameters are passed through a pointer to a FetchTestUserParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**StudioV2FlowTestUser**](StudioV2FlowTestUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExecution

> ListExecutionResponse ListExecution(ctx, FlowSid).DateCreatedFrom(DateCreatedFrom).DateCreatedTo(DateCreatedTo).PageSize(PageSize).Execute()





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
    FlowSid := "FlowSid_example" // string | The SID of the Flow with the Execution resources to read.
    DateCreatedFrom := time.Now() // time.Time | Only show Execution resources starting on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time, given as `YYYY-MM-DDThh:mm:ss-hh:mm`. (optional)
    DateCreatedTo := time.Now() // time.Time | Only show Execution resources starting before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time, given as `YYYY-MM-DDThh:mm:ss-hh:mm`. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListExecution(context.Background(), FlowSid).DateCreatedFrom(DateCreatedFrom).DateCreatedTo(DateCreatedTo).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListExecution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExecution`: ListExecutionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListExecution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Execution resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListExecutionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **DateCreatedFrom** | **time.Time** | Only show Execution resources starting on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time, given as &#x60;YYYY-MM-DDThh:mm:ss-hh:mm&#x60;.
 **DateCreatedTo** | **time.Time** | Only show Execution resources starting before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time, given as &#x60;YYYY-MM-DDThh:mm:ss-hh:mm&#x60;.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListExecutionResponse**](ListExecutionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExecutionStep

> ListExecutionStepResponse ListExecutionStep(ctx, FlowSid, ExecutionSid).PageSize(PageSize).Execute()





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
    FlowSid := "FlowSid_example" // string | The SID of the Flow with the Steps to read.
    ExecutionSid := "ExecutionSid_example" // string | The SID of the Execution with the Steps to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListExecutionStep(context.Background(), FlowSid, ExecutionSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListExecutionStep``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExecutionStep`: ListExecutionStepResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListExecutionStep`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Steps to read.
**ExecutionSid** | **string** | The SID of the Execution with the Steps to read.

### Other Parameters

Other parameters are passed through a pointer to a ListExecutionStepParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListExecutionStepResponse**](ListExecutionStepResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFlow

> ListFlowResponse ListFlow(ctx).PageSize(PageSize).Execute()





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
    resp, r, err := api_client.DefaultApi.ListFlow(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFlow`: ListFlowResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListFlow`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListFlowParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListFlowResponse**](ListFlowResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFlowRevision

> ListFlowRevisionResponse ListFlowRevision(ctx, Sid).PageSize(PageSize).Execute()





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
    Sid := "Sid_example" // string | The SID of the Flow resource to fetch.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListFlowRevision(context.Background(), Sid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListFlowRevision``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFlowRevision`: ListFlowRevisionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListFlowRevision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flow resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a ListFlowRevisionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListFlowRevisionResponse**](ListFlowRevisionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExecution

> StudioV2FlowExecution UpdateExecution(ctx, FlowSid, Sid).Status(Status).Execute()





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
    FlowSid := "FlowSid_example" // string | The SID of the Flow with the Execution resources to update.
    Sid := "Sid_example" // string | The SID of the Execution resource to update.
    Status := "Status_example" // string | The status of the Execution. Can only be `ended`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateExecution(context.Background(), FlowSid, Sid).Status(Status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateExecution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExecution`: StudioV2FlowExecution
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateExecution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Execution resources to update.
**Sid** | **string** | The SID of the Execution resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateExecutionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Status** | **string** | The status of the Execution. Can only be &#x60;ended&#x60;.

### Return type

[**StudioV2FlowExecution**](StudioV2FlowExecution.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFlow

> StudioV2Flow UpdateFlow(ctx, Sid).CommitMessage(CommitMessage).Definition(Definition).FriendlyName(FriendlyName).Status(Status).Execute()





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
    Sid := "Sid_example" // string | The SID of the Flow resource to fetch.
    CommitMessage := "CommitMessage_example" // string | Description of change made in the revision. (optional)
    Definition := TODO // map[string]interface{} | JSON representation of flow definition. (optional)
    FriendlyName := "FriendlyName_example" // string | The string that you assigned to describe the Flow. (optional)
    Status := "Status_example" // string | The status of the Flow. Can be: `draft` or `published`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateFlow(context.Background(), Sid).CommitMessage(CommitMessage).Definition(Definition).FriendlyName(FriendlyName).Status(Status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFlow`: StudioV2Flow
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flow resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a UpdateFlowParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **CommitMessage** | **string** | Description of change made in the revision.
 **Definition** | [**map[string]interface{}**](map[string]interface{}.md) | JSON representation of flow definition.
 **FriendlyName** | **string** | The string that you assigned to describe the Flow.
 **Status** | **string** | The status of the Flow. Can be: &#x60;draft&#x60; or &#x60;published&#x60;.

### Return type

[**StudioV2Flow**](StudioV2Flow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFlowValidate

> StudioV2FlowValidate UpdateFlowValidate(ctx).CommitMessage(CommitMessage).Definition(Definition).FriendlyName(FriendlyName).Status(Status).Execute()





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
    CommitMessage := "CommitMessage_example" // string | Description of change made in the revision. (optional)
    Definition := TODO // map[string]interface{} | JSON representation of flow definition. (optional)
    FriendlyName := "FriendlyName_example" // string | The string that you assigned to describe the Flow. (optional)
    Status := "Status_example" // string | The status of the Flow. Can be: `draft` or `published`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateFlowValidate(context.Background()).CommitMessage(CommitMessage).Definition(Definition).FriendlyName(FriendlyName).Status(Status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateFlowValidate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFlowValidate`: StudioV2FlowValidate
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateFlowValidate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateFlowValidateParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **CommitMessage** | **string** | Description of change made in the revision.
 **Definition** | [**map[string]interface{}**](map[string]interface{}.md) | JSON representation of flow definition.
 **FriendlyName** | **string** | The string that you assigned to describe the Flow.
 **Status** | **string** | The status of the Flow. Can be: &#x60;draft&#x60; or &#x60;published&#x60;.

### Return type

[**StudioV2FlowValidate**](StudioV2FlowValidate.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTestUser

> StudioV2FlowTestUser UpdateTestUser(ctx, Sid).TestUsers(TestUsers).Execute()





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
    Sid := "Sid_example" // string | Unique identifier of the flow.
    TestUsers := []string{"Inner_example"} // []string | List of test user identities that can test draft versions of the flow. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateTestUser(context.Background(), Sid).TestUsers(TestUsers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateTestUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTestUser`: StudioV2FlowTestUser
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateTestUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | Unique identifier of the flow.

### Other Parameters

Other parameters are passed through a pointer to a UpdateTestUserParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **TestUsers** | **[]string** | List of test user identities that can test draft versions of the flow.

### Return type

[**StudioV2FlowTestUser**](StudioV2FlowTestUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

