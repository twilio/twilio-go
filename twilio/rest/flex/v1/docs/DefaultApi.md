# \DefaultApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChannel**](DefaultApi.md#CreateChannel) | **Post** /v1/Channels | 
[**CreateFlexFlow**](DefaultApi.md#CreateFlexFlow) | **Post** /v1/FlexFlows | 
[**CreateWebChannel**](DefaultApi.md#CreateWebChannel) | **Post** /v1/WebChannels | 
[**DeleteChannel**](DefaultApi.md#DeleteChannel) | **Delete** /v1/Channels/{Sid} | 
[**DeleteFlexFlow**](DefaultApi.md#DeleteFlexFlow) | **Delete** /v1/FlexFlows/{Sid} | 
[**DeleteWebChannel**](DefaultApi.md#DeleteWebChannel) | **Delete** /v1/WebChannels/{Sid} | 
[**FetchChannel**](DefaultApi.md#FetchChannel) | **Get** /v1/Channels/{Sid} | 
[**FetchConfiguration**](DefaultApi.md#FetchConfiguration) | **Get** /v1/Configuration | 
[**FetchFlexFlow**](DefaultApi.md#FetchFlexFlow) | **Get** /v1/FlexFlows/{Sid} | 
[**FetchWebChannel**](DefaultApi.md#FetchWebChannel) | **Get** /v1/WebChannels/{Sid} | 
[**ListChannel**](DefaultApi.md#ListChannel) | **Get** /v1/Channels | 
[**ListFlexFlow**](DefaultApi.md#ListFlexFlow) | **Get** /v1/FlexFlows | 
[**ListWebChannel**](DefaultApi.md#ListWebChannel) | **Get** /v1/WebChannels | 
[**UpdateConfiguration**](DefaultApi.md#UpdateConfiguration) | **Post** /v1/Configuration | 
[**UpdateFlexFlow**](DefaultApi.md#UpdateFlexFlow) | **Post** /v1/FlexFlows/{Sid} | 
[**UpdateWebChannel**](DefaultApi.md#UpdateWebChannel) | **Post** /v1/WebChannels/{Sid} | 



## CreateChannel

> FlexV1Channel CreateChannel(ctx).ChatFriendlyName(ChatFriendlyName).ChatUniqueName(ChatUniqueName).ChatUserFriendlyName(ChatUserFriendlyName).FlexFlowSid(FlexFlowSid).Identity(Identity).LongLived(LongLived).PreEngagementData(PreEngagementData).Target(Target).TaskAttributes(TaskAttributes).TaskSid(TaskSid).Execute()



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
    ChatFriendlyName := "ChatFriendlyName_example" // string | The chat channel's friendly name. (optional)
    ChatUniqueName := "ChatUniqueName_example" // string | The chat channel's unique name. (optional)
    ChatUserFriendlyName := "ChatUserFriendlyName_example" // string | The chat participant's friendly name. (optional)
    FlexFlowSid := "FlexFlowSid_example" // string | The SID of the Flex Flow. (optional)
    Identity := "Identity_example" // string | The `identity` value that uniquely identifies the new resource's chat User. (optional)
    LongLived := true // bool | Whether to create the channel as long-lived. (optional)
    PreEngagementData := "PreEngagementData_example" // string | The pre-engagement data. (optional)
    Target := "Target_example" // string | The Target Contact Identity, for example the phone number of an SMS. (optional)
    TaskAttributes := "TaskAttributes_example" // string | The Task attributes to be added for the TaskRouter Task. (optional)
    TaskSid := "TaskSid_example" // string | The SID of the TaskRouter Task. Only valid when integration type is `task`. `null` for integration types `studio` & `external` (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateChannel(context.Background()).ChatFriendlyName(ChatFriendlyName).ChatUniqueName(ChatUniqueName).ChatUserFriendlyName(ChatUserFriendlyName).FlexFlowSid(FlexFlowSid).Identity(Identity).LongLived(LongLived).PreEngagementData(PreEngagementData).Target(Target).TaskAttributes(TaskAttributes).TaskSid(TaskSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateChannel`: FlexV1Channel
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateChannel`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **ChatFriendlyName** | **string** | The chat channel&#39;s friendly name.
 **ChatUniqueName** | **string** | The chat channel&#39;s unique name.
 **ChatUserFriendlyName** | **string** | The chat participant&#39;s friendly name.
 **FlexFlowSid** | **string** | The SID of the Flex Flow.
 **Identity** | **string** | The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s chat User.
 **LongLived** | **bool** | Whether to create the channel as long-lived.
 **PreEngagementData** | **string** | The pre-engagement data.
 **Target** | **string** | The Target Contact Identity, for example the phone number of an SMS.
 **TaskAttributes** | **string** | The Task attributes to be added for the TaskRouter Task.
 **TaskSid** | **string** | The SID of the TaskRouter Task. Only valid when integration type is &#x60;task&#x60;. &#x60;null&#x60; for integration types &#x60;studio&#x60; &amp; &#x60;external&#x60;

### Return type

[**FlexV1Channel**](FlexV1Channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFlexFlow

> FlexV1FlexFlow CreateFlexFlow(ctx).ChannelType(ChannelType).ChatServiceSid(ChatServiceSid).ContactIdentity(ContactIdentity).Enabled(Enabled).FriendlyName(FriendlyName).IntegrationChannel(IntegrationChannel).IntegrationCreationOnMessage(IntegrationCreationOnMessage).IntegrationFlowSid(IntegrationFlowSid).IntegrationPriority(IntegrationPriority).IntegrationRetryCount(IntegrationRetryCount).IntegrationTimeout(IntegrationTimeout).IntegrationUrl(IntegrationUrl).IntegrationWorkflowSid(IntegrationWorkflowSid).IntegrationWorkspaceSid(IntegrationWorkspaceSid).IntegrationType(IntegrationType).JanitorEnabled(JanitorEnabled).LongLived(LongLived).Execute()



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
    ChannelType := "ChannelType_example" // string | The channel type. Can be: `web`, `facebook`, `sms`, `whatsapp`, `line` or `custom`. (optional)
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the chat service. (optional)
    ContactIdentity := "ContactIdentity_example" // string | The channel contact's Identity. (optional)
    Enabled := true // bool | Whether the new Flex Flow is enabled. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the Flex Flow resource. (optional)
    IntegrationChannel := "IntegrationChannel_example" // string | The Task Channel for the TaskRouter Task that will be created. Applicable and required when integrationType is `task`. Set to `sms` for SMS, and to `chat` otherwise. The default value is `default` (optional)
    IntegrationCreationOnMessage := true // bool | In the context of outbound messaging, defines whether to create a Task immediately (and therefore reserve the conversation to current agent), or delay Task creation until the customer sends the first response. Set to false to create immediately, true to delay Task creation. This setting is only applicable for outbound messaging. (optional)
    IntegrationFlowSid := "IntegrationFlowSid_example" // string | The SID of the Studio Flow. Required when `integrationType` is `studio`. (optional)
    IntegrationPriority := int32(56) // int32 | The Task priority of a new Task. The default priority is 0. Optional when `integrationType` is `task`, not applicable otherwise. (optional)
    IntegrationRetryCount := int32(56) // int32 | The number of times to retry the webhook if the first attempt fails. Can be an integer between 0 and 3 (included), default is 0. Optional when integrationType is `external`, not applicable otherwise. (optional)
    IntegrationTimeout := int32(56) // int32 | The Task timeout in seconds for a new Task. Default is 86,400 seconds (24 hours). Optional when `integrationType` is `task`, not applicable otherwise. (optional)
    IntegrationUrl := "IntegrationUrl_example" // string | The URL of the external webhook. Required when `integrationType` is `external`. (optional)
    IntegrationWorkflowSid := "IntegrationWorkflowSid_example" // string | The Workflow SID for a new Task. Required when `integrationType` is `task`. (optional)
    IntegrationWorkspaceSid := "IntegrationWorkspaceSid_example" // string | The Workspace SID for a new Task. Required when `integrationType` is `task`. (optional)
    IntegrationType := "IntegrationType_example" // string | The integration type. Can be: `studio`, `external`, or `task`. (optional)
    JanitorEnabled := true // bool | When enabled, the Messaging Channel Janitor will remove active Proxy sessions if the associated Task is deleted outside of the Flex UI. Defaults to `false`. (optional)
    LongLived := true // bool | When enabled, Flex will keep the chat channel active so that it may be used for subsequent interactions with a contact identity. Defaults to `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateFlexFlow(context.Background()).ChannelType(ChannelType).ChatServiceSid(ChatServiceSid).ContactIdentity(ContactIdentity).Enabled(Enabled).FriendlyName(FriendlyName).IntegrationChannel(IntegrationChannel).IntegrationCreationOnMessage(IntegrationCreationOnMessage).IntegrationFlowSid(IntegrationFlowSid).IntegrationPriority(IntegrationPriority).IntegrationRetryCount(IntegrationRetryCount).IntegrationTimeout(IntegrationTimeout).IntegrationUrl(IntegrationUrl).IntegrationWorkflowSid(IntegrationWorkflowSid).IntegrationWorkspaceSid(IntegrationWorkspaceSid).IntegrationType(IntegrationType).JanitorEnabled(JanitorEnabled).LongLived(LongLived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateFlexFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFlexFlow`: FlexV1FlexFlow
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateFlexFlow`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateFlexFlowParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **ChannelType** | **string** | The channel type. Can be: &#x60;web&#x60;, &#x60;facebook&#x60;, &#x60;sms&#x60;, &#x60;whatsapp&#x60;, &#x60;line&#x60; or &#x60;custom&#x60;.
 **ChatServiceSid** | **string** | The SID of the chat service.
 **ContactIdentity** | **string** | The channel contact&#39;s Identity.
 **Enabled** | **bool** | Whether the new Flex Flow is enabled.
 **FriendlyName** | **string** | A descriptive string that you create to describe the Flex Flow resource.
 **IntegrationChannel** | **string** | The Task Channel for the TaskRouter Task that will be created. Applicable and required when integrationType is &#x60;task&#x60;. Set to &#x60;sms&#x60; for SMS, and to &#x60;chat&#x60; otherwise. The default value is &#x60;default&#x60;
 **IntegrationCreationOnMessage** | **bool** | In the context of outbound messaging, defines whether to create a Task immediately (and therefore reserve the conversation to current agent), or delay Task creation until the customer sends the first response. Set to false to create immediately, true to delay Task creation. This setting is only applicable for outbound messaging.
 **IntegrationFlowSid** | **string** | The SID of the Studio Flow. Required when &#x60;integrationType&#x60; is &#x60;studio&#x60;.
 **IntegrationPriority** | **int32** | The Task priority of a new Task. The default priority is 0. Optional when &#x60;integrationType&#x60; is &#x60;task&#x60;, not applicable otherwise.
 **IntegrationRetryCount** | **int32** | The number of times to retry the webhook if the first attempt fails. Can be an integer between 0 and 3 (included), default is 0. Optional when integrationType is &#x60;external&#x60;, not applicable otherwise.
 **IntegrationTimeout** | **int32** | The Task timeout in seconds for a new Task. Default is 86,400 seconds (24 hours). Optional when &#x60;integrationType&#x60; is &#x60;task&#x60;, not applicable otherwise.
 **IntegrationUrl** | **string** | The URL of the external webhook. Required when &#x60;integrationType&#x60; is &#x60;external&#x60;.
 **IntegrationWorkflowSid** | **string** | The Workflow SID for a new Task. Required when &#x60;integrationType&#x60; is &#x60;task&#x60;.
 **IntegrationWorkspaceSid** | **string** | The Workspace SID for a new Task. Required when &#x60;integrationType&#x60; is &#x60;task&#x60;.
 **IntegrationType** | **string** | The integration type. Can be: &#x60;studio&#x60;, &#x60;external&#x60;, or &#x60;task&#x60;.
 **JanitorEnabled** | **bool** | When enabled, the Messaging Channel Janitor will remove active Proxy sessions if the associated Task is deleted outside of the Flex UI. Defaults to &#x60;false&#x60;.
 **LongLived** | **bool** | When enabled, Flex will keep the chat channel active so that it may be used for subsequent interactions with a contact identity. Defaults to &#x60;false&#x60;.

### Return type

[**FlexV1FlexFlow**](FlexV1FlexFlow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebChannel

> FlexV1WebChannel CreateWebChannel(ctx).ChatFriendlyName(ChatFriendlyName).ChatUniqueName(ChatUniqueName).CustomerFriendlyName(CustomerFriendlyName).FlexFlowSid(FlexFlowSid).Identity(Identity).PreEngagementData(PreEngagementData).Execute()



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
    ChatFriendlyName := "ChatFriendlyName_example" // string | The chat channel's friendly name. (optional)
    ChatUniqueName := "ChatUniqueName_example" // string | The chat channel's unique name. (optional)
    CustomerFriendlyName := "CustomerFriendlyName_example" // string | The chat participant's friendly name. (optional)
    FlexFlowSid := "FlexFlowSid_example" // string | The SID of the Flex Flow. (optional)
    Identity := "Identity_example" // string | The chat identity. (optional)
    PreEngagementData := "PreEngagementData_example" // string | The pre-engagement data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateWebChannel(context.Background()).ChatFriendlyName(ChatFriendlyName).ChatUniqueName(ChatUniqueName).CustomerFriendlyName(CustomerFriendlyName).FlexFlowSid(FlexFlowSid).Identity(Identity).PreEngagementData(PreEngagementData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateWebChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebChannel`: FlexV1WebChannel
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateWebChannel`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateWebChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **ChatFriendlyName** | **string** | The chat channel&#39;s friendly name.
 **ChatUniqueName** | **string** | The chat channel&#39;s unique name.
 **CustomerFriendlyName** | **string** | The chat participant&#39;s friendly name.
 **FlexFlowSid** | **string** | The SID of the Flex Flow.
 **Identity** | **string** | The chat identity.
 **PreEngagementData** | **string** | The pre-engagement data.

### Return type

[**FlexV1WebChannel**](FlexV1WebChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChannel

> DeleteChannel(ctx, Sid).Execute()



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
    Sid := "Sid_example" // string | The SID of the Flex chat channel resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteChannel(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flex chat channel resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteChannelParams struct via the builder pattern


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


## DeleteFlexFlow

> DeleteFlexFlow(ctx, Sid).Execute()



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
    Sid := "Sid_example" // string | The SID of the Flex Flow resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteFlexFlow(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteFlexFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flex Flow resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteFlexFlowParams struct via the builder pattern


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


## DeleteWebChannel

> DeleteWebChannel(ctx, Sid).Execute()



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
    Sid := "Sid_example" // string | The SID of the WebChannel resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteWebChannel(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteWebChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the WebChannel resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteWebChannelParams struct via the builder pattern


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


## FetchChannel

> FlexV1Channel FetchChannel(ctx, Sid).Execute()



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
    Sid := "Sid_example" // string | The SID of the Flex chat channel resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchChannel(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchChannel`: FlexV1Channel
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flex chat channel resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**FlexV1Channel**](FlexV1Channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConfiguration

> FlexV1Configuration FetchConfiguration(ctx).UiVersion(UiVersion).Execute()



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
    UiVersion := "UiVersion_example" // string | The Pinned UI version of the Configuration resource to fetch. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchConfiguration(context.Background()).UiVersion(UiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchConfiguration`: FlexV1Configuration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchConfigurationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **UiVersion** | **string** | The Pinned UI version of the Configuration resource to fetch.

### Return type

[**FlexV1Configuration**](FlexV1Configuration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFlexFlow

> FlexV1FlexFlow FetchFlexFlow(ctx, Sid).Execute()



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
    Sid := "Sid_example" // string | The SID of the Flex Flow resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchFlexFlow(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchFlexFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchFlexFlow`: FlexV1FlexFlow
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchFlexFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flex Flow resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFlexFlowParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**FlexV1FlexFlow**](FlexV1FlexFlow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWebChannel

> FlexV1WebChannel FetchWebChannel(ctx, Sid).Execute()



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
    Sid := "Sid_example" // string | The SID of the WebChannel resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchWebChannel(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchWebChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchWebChannel`: FlexV1WebChannel
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchWebChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the WebChannel resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWebChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**FlexV1WebChannel**](FlexV1WebChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChannel

> ListChannelResponse ListChannel(ctx).PageSize(PageSize).Execute()



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
    resp, r, err := api_client.DefaultApi.ListChannel(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListChannel`: ListChannelResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListChannel`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListChannelResponse**](ListChannelResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFlexFlow

> ListFlexFlowResponse ListFlexFlow(ctx).FriendlyName(FriendlyName).PageSize(PageSize).Execute()



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
    FriendlyName := "FriendlyName_example" // string | The `friendly_name` of the Flex Flow resources to read. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListFlexFlow(context.Background()).FriendlyName(FriendlyName).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListFlexFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFlexFlow`: ListFlexFlowResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListFlexFlow`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListFlexFlowParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **FriendlyName** | **string** | The &#x60;friendly_name&#x60; of the Flex Flow resources to read.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListFlexFlowResponse**](ListFlexFlowResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebChannel

> ListWebChannelResponse ListWebChannel(ctx).PageSize(PageSize).Execute()



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
    resp, r, err := api_client.DefaultApi.ListWebChannel(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListWebChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWebChannel`: ListWebChannelResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListWebChannel`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListWebChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListWebChannelResponse**](ListWebChannelResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration

> FlexV1Configuration UpdateConfiguration(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfiguration`: FlexV1Configuration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConfigurationParams struct via the builder pattern


### Return type

[**FlexV1Configuration**](FlexV1Configuration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFlexFlow

> FlexV1FlexFlow UpdateFlexFlow(ctx, Sid).ChannelType(ChannelType).ChatServiceSid(ChatServiceSid).ContactIdentity(ContactIdentity).Enabled(Enabled).FriendlyName(FriendlyName).IntegrationChannel(IntegrationChannel).IntegrationCreationOnMessage(IntegrationCreationOnMessage).IntegrationFlowSid(IntegrationFlowSid).IntegrationPriority(IntegrationPriority).IntegrationRetryCount(IntegrationRetryCount).IntegrationTimeout(IntegrationTimeout).IntegrationUrl(IntegrationUrl).IntegrationWorkflowSid(IntegrationWorkflowSid).IntegrationWorkspaceSid(IntegrationWorkspaceSid).IntegrationType(IntegrationType).JanitorEnabled(JanitorEnabled).LongLived(LongLived).Execute()



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
    Sid := "Sid_example" // string | The SID of the Flex Flow resource to update.
    ChannelType := "ChannelType_example" // string | The channel type. Can be: `web`, `facebook`, `sms`, `whatsapp`, `line` or `custom`. (optional)
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the chat service. (optional)
    ContactIdentity := "ContactIdentity_example" // string | The channel contact's Identity. (optional)
    Enabled := true // bool | Whether the new Flex Flow is enabled. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the Flex Flow resource. (optional)
    IntegrationChannel := "IntegrationChannel_example" // string | The Task Channel for the TaskRouter Task that will be created. Applicable and required when integrationType is `task`. Set to `sms` for SMS, and to `chat` otherwise. The default value is `default` (optional)
    IntegrationCreationOnMessage := true // bool | In the context of outbound messaging, defines whether to create a Task immediately (and therefore reserve the conversation to current agent), or delay Task creation until the customer sends the first response. Set to false to create immediately, true to delay Task creation. This setting is only applicable for outbound messaging. (optional)
    IntegrationFlowSid := "IntegrationFlowSid_example" // string | The SID of the Studio Flow. Required when `integrationType` is `studio`. (optional)
    IntegrationPriority := int32(56) // int32 | The Task priority of a new Task. The default priority is 0. Optional when `integrationType` is `task`, not applicable otherwise. (optional)
    IntegrationRetryCount := int32(56) // int32 | The number of times to retry the webhook if the first attempt fails. Can be an integer between 0 and 3 (included), default is 0. Optional when integrationType is `external`, not applicable otherwise. (optional)
    IntegrationTimeout := int32(56) // int32 | The Task timeout in seconds for a new Task. Default is 86,400 seconds (24 hours). Optional when `integrationType` is `task`, not applicable otherwise. (optional)
    IntegrationUrl := "IntegrationUrl_example" // string | The URL of the external webhook. Required when `integrationType` is `external`. (optional)
    IntegrationWorkflowSid := "IntegrationWorkflowSid_example" // string | The Workflow SID for a new Task. Required when `integrationType` is `task`. (optional)
    IntegrationWorkspaceSid := "IntegrationWorkspaceSid_example" // string | The Workspace SID for a new Task. Required when `integrationType` is `task`. (optional)
    IntegrationType := "IntegrationType_example" // string | The integration type. Can be: `studio`, `external`, or `task`. (optional)
    JanitorEnabled := true // bool | When enabled, the Messaging Channel Janitor will remove active Proxy sessions if the associated Task is deleted outside of the Flex UI. Defaults to `false`. (optional)
    LongLived := true // bool | When enabled, Flex will keep the chat channel active so that it may be used for subsequent interactions with a contact identity. Defaults to `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateFlexFlow(context.Background(), Sid).ChannelType(ChannelType).ChatServiceSid(ChatServiceSid).ContactIdentity(ContactIdentity).Enabled(Enabled).FriendlyName(FriendlyName).IntegrationChannel(IntegrationChannel).IntegrationCreationOnMessage(IntegrationCreationOnMessage).IntegrationFlowSid(IntegrationFlowSid).IntegrationPriority(IntegrationPriority).IntegrationRetryCount(IntegrationRetryCount).IntegrationTimeout(IntegrationTimeout).IntegrationUrl(IntegrationUrl).IntegrationWorkflowSid(IntegrationWorkflowSid).IntegrationWorkspaceSid(IntegrationWorkspaceSid).IntegrationType(IntegrationType).JanitorEnabled(JanitorEnabled).LongLived(LongLived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateFlexFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFlexFlow`: FlexV1FlexFlow
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateFlexFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flex Flow resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateFlexFlowParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **ChannelType** | **string** | The channel type. Can be: &#x60;web&#x60;, &#x60;facebook&#x60;, &#x60;sms&#x60;, &#x60;whatsapp&#x60;, &#x60;line&#x60; or &#x60;custom&#x60;.
 **ChatServiceSid** | **string** | The SID of the chat service.
 **ContactIdentity** | **string** | The channel contact&#39;s Identity.
 **Enabled** | **bool** | Whether the new Flex Flow is enabled.
 **FriendlyName** | **string** | A descriptive string that you create to describe the Flex Flow resource.
 **IntegrationChannel** | **string** | The Task Channel for the TaskRouter Task that will be created. Applicable and required when integrationType is &#x60;task&#x60;. Set to &#x60;sms&#x60; for SMS, and to &#x60;chat&#x60; otherwise. The default value is &#x60;default&#x60;
 **IntegrationCreationOnMessage** | **bool** | In the context of outbound messaging, defines whether to create a Task immediately (and therefore reserve the conversation to current agent), or delay Task creation until the customer sends the first response. Set to false to create immediately, true to delay Task creation. This setting is only applicable for outbound messaging.
 **IntegrationFlowSid** | **string** | The SID of the Studio Flow. Required when &#x60;integrationType&#x60; is &#x60;studio&#x60;.
 **IntegrationPriority** | **int32** | The Task priority of a new Task. The default priority is 0. Optional when &#x60;integrationType&#x60; is &#x60;task&#x60;, not applicable otherwise.
 **IntegrationRetryCount** | **int32** | The number of times to retry the webhook if the first attempt fails. Can be an integer between 0 and 3 (included), default is 0. Optional when integrationType is &#x60;external&#x60;, not applicable otherwise.
 **IntegrationTimeout** | **int32** | The Task timeout in seconds for a new Task. Default is 86,400 seconds (24 hours). Optional when &#x60;integrationType&#x60; is &#x60;task&#x60;, not applicable otherwise.
 **IntegrationUrl** | **string** | The URL of the external webhook. Required when &#x60;integrationType&#x60; is &#x60;external&#x60;.
 **IntegrationWorkflowSid** | **string** | The Workflow SID for a new Task. Required when &#x60;integrationType&#x60; is &#x60;task&#x60;.
 **IntegrationWorkspaceSid** | **string** | The Workspace SID for a new Task. Required when &#x60;integrationType&#x60; is &#x60;task&#x60;.
 **IntegrationType** | **string** | The integration type. Can be: &#x60;studio&#x60;, &#x60;external&#x60;, or &#x60;task&#x60;.
 **JanitorEnabled** | **bool** | When enabled, the Messaging Channel Janitor will remove active Proxy sessions if the associated Task is deleted outside of the Flex UI. Defaults to &#x60;false&#x60;.
 **LongLived** | **bool** | When enabled, Flex will keep the chat channel active so that it may be used for subsequent interactions with a contact identity. Defaults to &#x60;false&#x60;.

### Return type

[**FlexV1FlexFlow**](FlexV1FlexFlow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebChannel

> FlexV1WebChannel UpdateWebChannel(ctx, Sid).ChatStatus(ChatStatus).PostEngagementData(PostEngagementData).Execute()



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
    Sid := "Sid_example" // string | The SID of the WebChannel resource to update.
    ChatStatus := "ChatStatus_example" // string | The chat status. Can only be `inactive`. (optional)
    PostEngagementData := "PostEngagementData_example" // string | The post-engagement data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateWebChannel(context.Background(), Sid).ChatStatus(ChatStatus).PostEngagementData(PostEngagementData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateWebChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebChannel`: FlexV1WebChannel
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateWebChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the WebChannel resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateWebChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **ChatStatus** | **string** | The chat status. Can only be &#x60;inactive&#x60;.
 **PostEngagementData** | **string** | The post-engagement data.

### Return type

[**FlexV1WebChannel**](FlexV1WebChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

