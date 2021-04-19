# \DefaultApi

All URIs are relative to *https://studio.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEngagement**](DefaultApi.md#CreateEngagement) | **Post** /v1/Flows/{FlowSid}/Engagements | 
[**CreateExecution**](DefaultApi.md#CreateExecution) | **Post** /v1/Flows/{FlowSid}/Executions | 
[**DeleteEngagement**](DefaultApi.md#DeleteEngagement) | **Delete** /v1/Flows/{FlowSid}/Engagements/{Sid} | 
[**DeleteExecution**](DefaultApi.md#DeleteExecution) | **Delete** /v1/Flows/{FlowSid}/Executions/{Sid} | 
[**DeleteFlow**](DefaultApi.md#DeleteFlow) | **Delete** /v1/Flows/{Sid} | 
[**FetchEngagement**](DefaultApi.md#FetchEngagement) | **Get** /v1/Flows/{FlowSid}/Engagements/{Sid} | 
[**FetchEngagementContext**](DefaultApi.md#FetchEngagementContext) | **Get** /v1/Flows/{FlowSid}/Engagements/{EngagementSid}/Context | 
[**FetchExecution**](DefaultApi.md#FetchExecution) | **Get** /v1/Flows/{FlowSid}/Executions/{Sid} | 
[**FetchExecutionContext**](DefaultApi.md#FetchExecutionContext) | **Get** /v1/Flows/{FlowSid}/Executions/{ExecutionSid}/Context | 
[**FetchExecutionStep**](DefaultApi.md#FetchExecutionStep) | **Get** /v1/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps/{Sid} | 
[**FetchExecutionStepContext**](DefaultApi.md#FetchExecutionStepContext) | **Get** /v1/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps/{StepSid}/Context | 
[**FetchFlow**](DefaultApi.md#FetchFlow) | **Get** /v1/Flows/{Sid} | 
[**FetchStep**](DefaultApi.md#FetchStep) | **Get** /v1/Flows/{FlowSid}/Engagements/{EngagementSid}/Steps/{Sid} | 
[**FetchStepContext**](DefaultApi.md#FetchStepContext) | **Get** /v1/Flows/{FlowSid}/Engagements/{EngagementSid}/Steps/{StepSid}/Context | 
[**ListEngagement**](DefaultApi.md#ListEngagement) | **Get** /v1/Flows/{FlowSid}/Engagements | 
[**ListExecution**](DefaultApi.md#ListExecution) | **Get** /v1/Flows/{FlowSid}/Executions | 
[**ListExecutionStep**](DefaultApi.md#ListExecutionStep) | **Get** /v1/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps | 
[**ListFlow**](DefaultApi.md#ListFlow) | **Get** /v1/Flows | 
[**ListStep**](DefaultApi.md#ListStep) | **Get** /v1/Flows/{FlowSid}/Engagements/{EngagementSid}/Steps | 
[**UpdateExecution**](DefaultApi.md#UpdateExecution) | **Post** /v1/Flows/{FlowSid}/Executions/{Sid} | 



## CreateEngagement

> StudioV1FlowEngagement CreateEngagement(ctx, FlowSid).From(From).Parameters(Parameters).To(To).Execute()





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
    FlowSid := "FlowSid_example" // string | The SID of the Flow.
    From := "From_example" // string | The Twilio phone number to send messages or initiate calls from during the Flow Engagement. Available as variable `{{flow.channel.address}}` (optional)
    Parameters := TODO // map[string]interface{} | A JSON string we will add to your flow's context and that you can access as variables inside your flow. For example, if you pass in `Parameters={'name':'Zeke'}` then inside a widget you can reference the variable `{{flow.data.name}}` which will return the string 'Zeke'. Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode your JSON string. (optional)
    To := "To_example" // string | The Contact phone number to start a Studio Flow Engagement, available as variable `{{contact.channel.address}}`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateEngagement(context.Background(), FlowSid).From(From).Parameters(Parameters).To(To).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateEngagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEngagement`: StudioV1FlowEngagement
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateEngagement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow.

### Other Parameters

Other parameters are passed through a pointer to a CreateEngagementParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **From** | **string** | The Twilio phone number to send messages or initiate calls from during the Flow Engagement. Available as variable &#x60;{{flow.channel.address}}&#x60;
 **Parameters** | [**map[string]interface{}**](map[string]interface{}.md) | A JSON string we will add to your flow&#39;s context and that you can access as variables inside your flow. For example, if you pass in &#x60;Parameters&#x3D;{&#39;name&#39;:&#39;Zeke&#39;}&#x60; then inside a widget you can reference the variable &#x60;{{flow.data.name}}&#x60; which will return the string &#39;Zeke&#39;. Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode your JSON string.
 **To** | **string** | The Contact phone number to start a Studio Flow Engagement, available as variable &#x60;{{contact.channel.address}}&#x60;.

### Return type

[**StudioV1FlowEngagement**](StudioV1FlowEngagement.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExecution

> StudioV1FlowExecution CreateExecution(ctx, FlowSid).From(From).Parameters(Parameters).To(To).Execute()





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
    // response from `CreateExecution`: StudioV1FlowExecution
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

[**StudioV1FlowExecution**](StudioV1FlowExecution.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEngagement

> DeleteEngagement(ctx, FlowSid, Sid).Execute()





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
    FlowSid := "FlowSid_example" // string | The SID of the Flow to delete Engagements from.
    Sid := "Sid_example" // string | The SID of the Engagement resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteEngagement(context.Background(), FlowSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteEngagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow to delete Engagements from.
**Sid** | **string** | The SID of the Engagement resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteEngagementParams struct via the builder pattern


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


## FetchEngagement

> StudioV1FlowEngagement FetchEngagement(ctx, FlowSid, Sid).Execute()





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
    FlowSid := "FlowSid_example" // string | The SID of the Flow.
    Sid := "Sid_example" // string | The SID of the Engagement resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchEngagement(context.Background(), FlowSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchEngagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchEngagement`: StudioV1FlowEngagement
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchEngagement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow.
**Sid** | **string** | The SID of the Engagement resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchEngagementParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**StudioV1FlowEngagement**](StudioV1FlowEngagement.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchEngagementContext

> StudioV1FlowEngagementEngagementContext FetchEngagementContext(ctx, FlowSid, EngagementSid).Execute()





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
    FlowSid := "FlowSid_example" // string | The SID of the Flow.
    EngagementSid := "EngagementSid_example" // string | The SID of the Engagement.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchEngagementContext(context.Background(), FlowSid, EngagementSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchEngagementContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchEngagementContext`: StudioV1FlowEngagementEngagementContext
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchEngagementContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow.
**EngagementSid** | **string** | The SID of the Engagement.

### Other Parameters

Other parameters are passed through a pointer to a FetchEngagementContextParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**StudioV1FlowEngagementEngagementContext**](StudioV1FlowEngagementEngagementContext.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExecution

> StudioV1FlowExecution FetchExecution(ctx, FlowSid, Sid).Execute()





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
    // response from `FetchExecution`: StudioV1FlowExecution
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

[**StudioV1FlowExecution**](StudioV1FlowExecution.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExecutionContext

> StudioV1FlowExecutionExecutionContext FetchExecutionContext(ctx, FlowSid, ExecutionSid).Execute()





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
    // response from `FetchExecutionContext`: StudioV1FlowExecutionExecutionContext
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

[**StudioV1FlowExecutionExecutionContext**](StudioV1FlowExecutionExecutionContext.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExecutionStep

> StudioV1FlowExecutionExecutionStep FetchExecutionStep(ctx, FlowSid, ExecutionSid, Sid).Execute()





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
    // response from `FetchExecutionStep`: StudioV1FlowExecutionExecutionStep
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

[**StudioV1FlowExecutionExecutionStep**](StudioV1FlowExecutionExecutionStep.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExecutionStepContext

> StudioV1FlowExecutionExecutionStepExecutionStepContext FetchExecutionStepContext(ctx, FlowSid, ExecutionSid, StepSid).Execute()





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
    // response from `FetchExecutionStepContext`: StudioV1FlowExecutionExecutionStepExecutionStepContext
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

[**StudioV1FlowExecutionExecutionStepExecutionStepContext**](StudioV1FlowExecutionExecutionStepExecutionStepContext.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFlow

> StudioV1Flow FetchFlow(ctx, Sid).Execute()





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
    // response from `FetchFlow`: StudioV1Flow
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

[**StudioV1Flow**](StudioV1Flow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchStep

> StudioV1FlowEngagementStep FetchStep(ctx, FlowSid, EngagementSid, Sid).Execute()





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
    EngagementSid := "EngagementSid_example" // string | The SID of the Engagement with the Step to fetch.
    Sid := "Sid_example" // string | The SID of the Step resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchStep(context.Background(), FlowSid, EngagementSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchStep``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchStep`: StudioV1FlowEngagementStep
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchStep`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Step to fetch.
**EngagementSid** | **string** | The SID of the Engagement with the Step to fetch.
**Sid** | **string** | The SID of the Step resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchStepParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**StudioV1FlowEngagementStep**](StudioV1FlowEngagementStep.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchStepContext

> StudioV1FlowEngagementStepStepContext FetchStepContext(ctx, FlowSid, EngagementSid, StepSid).Execute()





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
    EngagementSid := "EngagementSid_example" // string | The SID of the Engagement with the Step to fetch.
    StepSid := "StepSid_example" // string | The SID of the Step to fetch

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchStepContext(context.Background(), FlowSid, EngagementSid, StepSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchStepContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchStepContext`: StudioV1FlowEngagementStepStepContext
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchStepContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Step to fetch.
**EngagementSid** | **string** | The SID of the Engagement with the Step to fetch.
**StepSid** | **string** | The SID of the Step to fetch

### Other Parameters

Other parameters are passed through a pointer to a FetchStepContextParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**StudioV1FlowEngagementStepStepContext**](StudioV1FlowEngagementStepStepContext.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEngagement

> ListEngagementResponse ListEngagement(ctx, FlowSid).PageSize(PageSize).Execute()





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
    FlowSid := "FlowSid_example" // string | The SID of the Flow to read Engagements from.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListEngagement(context.Background(), FlowSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListEngagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEngagement`: ListEngagementResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListEngagement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow to read Engagements from.

### Other Parameters

Other parameters are passed through a pointer to a ListEngagementParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListEngagementResponse**](ListEngagementResponse.md)

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


## ListStep

> ListStepResponse ListStep(ctx, FlowSid, EngagementSid).PageSize(PageSize).Execute()





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
    FlowSid := "FlowSid_example" // string | The SID of the Flow with the Step to read.
    EngagementSid := "EngagementSid_example" // string | The SID of the Engagement with the Step to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListStep(context.Background(), FlowSid, EngagementSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListStep``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStep`: ListStepResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListStep`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Step to read.
**EngagementSid** | **string** | The SID of the Engagement with the Step to read.

### Other Parameters

Other parameters are passed through a pointer to a ListStepParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListStepResponse**](ListStepResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExecution

> StudioV1FlowExecution UpdateExecution(ctx, FlowSid, Sid).Status(Status).Execute()





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
    // response from `UpdateExecution`: StudioV1FlowExecution
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

[**StudioV1FlowExecution**](StudioV1FlowExecution.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

