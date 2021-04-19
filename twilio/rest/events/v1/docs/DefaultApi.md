# \DefaultApi

All URIs are relative to *https://events.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSink**](DefaultApi.md#CreateSink) | **Post** /v1/Sinks | 
[**CreateSinkTest**](DefaultApi.md#CreateSinkTest) | **Post** /v1/Sinks/{Sid}/Test | 
[**CreateSinkValidate**](DefaultApi.md#CreateSinkValidate) | **Post** /v1/Sinks/{Sid}/Validate | 
[**CreateSubscribedEvent**](DefaultApi.md#CreateSubscribedEvent) | **Post** /v1/Subscriptions/{SubscriptionSid}/SubscribedEvents | 
[**CreateSubscription**](DefaultApi.md#CreateSubscription) | **Post** /v1/Subscriptions | 
[**DeleteSink**](DefaultApi.md#DeleteSink) | **Delete** /v1/Sinks/{Sid} | 
[**DeleteSubscribedEvent**](DefaultApi.md#DeleteSubscribedEvent) | **Delete** /v1/Subscriptions/{SubscriptionSid}/SubscribedEvents/{Type} | 
[**DeleteSubscription**](DefaultApi.md#DeleteSubscription) | **Delete** /v1/Subscriptions/{Sid} | 
[**FetchEventType**](DefaultApi.md#FetchEventType) | **Get** /v1/Types/{Type} | 
[**FetchSchema**](DefaultApi.md#FetchSchema) | **Get** /v1/Schemas/{Id} | 
[**FetchSink**](DefaultApi.md#FetchSink) | **Get** /v1/Sinks/{Sid} | 
[**FetchSubscribedEvent**](DefaultApi.md#FetchSubscribedEvent) | **Get** /v1/Subscriptions/{SubscriptionSid}/SubscribedEvents/{Type} | 
[**FetchSubscription**](DefaultApi.md#FetchSubscription) | **Get** /v1/Subscriptions/{Sid} | 
[**FetchVersion**](DefaultApi.md#FetchVersion) | **Get** /v1/Schemas/{Id}/Versions/{SchemaVersion} | 
[**ListEventType**](DefaultApi.md#ListEventType) | **Get** /v1/Types | 
[**ListSink**](DefaultApi.md#ListSink) | **Get** /v1/Sinks | 
[**ListSubscribedEvent**](DefaultApi.md#ListSubscribedEvent) | **Get** /v1/Subscriptions/{SubscriptionSid}/SubscribedEvents | 
[**ListSubscription**](DefaultApi.md#ListSubscription) | **Get** /v1/Subscriptions | 
[**ListVersion**](DefaultApi.md#ListVersion) | **Get** /v1/Schemas/{Id}/Versions | 
[**UpdateSubscribedEvent**](DefaultApi.md#UpdateSubscribedEvent) | **Post** /v1/Subscriptions/{SubscriptionSid}/SubscribedEvents/{Type} | 
[**UpdateSubscription**](DefaultApi.md#UpdateSubscription) | **Post** /v1/Subscriptions/{Sid} | 



## CreateSink

> EventsV1Sink CreateSink(ctx).Description(Description).SinkConfiguration(SinkConfiguration).SinkType(SinkType).Execute()





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
    Description := "Description_example" // string | A human readable description for the Sink **This value should not contain PII.** (optional)
    SinkConfiguration := TODO // map[string]interface{} | The information required for Twilio to connect to the provided Sink encoded as JSON. (optional)
    SinkType := "SinkType_example" // string | The Sink type. Can only be \\\"kinesis\\\" or \\\"webhook\\\" currently. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSink(context.Background()).Description(Description).SinkConfiguration(SinkConfiguration).SinkType(SinkType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSink`: EventsV1Sink
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSink`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSinkParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **Description** | **string** | A human readable description for the Sink **This value should not contain PII.**
 **SinkConfiguration** | [**map[string]interface{}**](map[string]interface{}.md) | The information required for Twilio to connect to the provided Sink encoded as JSON.
 **SinkType** | **string** | The Sink type. Can only be \\\&quot;kinesis\\\&quot; or \\\&quot;webhook\\\&quot; currently.

### Return type

[**EventsV1Sink**](EventsV1Sink.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSinkTest

> EventsV1SinkSinkTest CreateSinkTest(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies the Sink to be Tested.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSinkTest(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSinkTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSinkTest`: EventsV1SinkSinkTest
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSinkTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies the Sink to be Tested.

### Other Parameters

Other parameters are passed through a pointer to a CreateSinkTestParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**EventsV1SinkSinkTest**](EventsV1SinkSinkTest.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSinkValidate

> EventsV1SinkSinkValidate CreateSinkValidate(ctx, Sid).TestId(TestId).Execute()





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
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies the Sink being validated.
    TestId := "TestId_example" // string | A 34 character string that uniquely identifies the test event for a Sink being validated. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSinkValidate(context.Background(), Sid).TestId(TestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSinkValidate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSinkValidate`: EventsV1SinkSinkValidate
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSinkValidate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies the Sink being validated.

### Other Parameters

Other parameters are passed through a pointer to a CreateSinkValidateParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **TestId** | **string** | A 34 character string that uniquely identifies the test event for a Sink being validated.

### Return type

[**EventsV1SinkSinkValidate**](EventsV1SinkSinkValidate.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubscribedEvent

> EventsV1SubscriptionSubscribedEvent CreateSubscribedEvent(ctx, SubscriptionSid).Type(Type).Version(Version).Execute()





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
    SubscriptionSid := "SubscriptionSid_example" // string | The unique SID identifier of the Subscription.
    Type := "Type_example" // string | Type of event being subscribed to. (optional)
    Version := int32(56) // int32 | The schema version that the subscription should use. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSubscribedEvent(context.Background(), SubscriptionSid).Type(Type).Version(Version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSubscribedEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubscribedEvent`: EventsV1SubscriptionSubscribedEvent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSubscribedEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SubscriptionSid** | **string** | The unique SID identifier of the Subscription.

### Other Parameters

Other parameters are passed through a pointer to a CreateSubscribedEventParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Type** | **string** | Type of event being subscribed to.
 **Version** | **int32** | The schema version that the subscription should use.

### Return type

[**EventsV1SubscriptionSubscribedEvent**](EventsV1SubscriptionSubscribedEvent.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubscription

> EventsV1Subscription CreateSubscription(ctx).Description(Description).SinkSid(SinkSid).Types(Types).Execute()





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
    Description := "Description_example" // string | A human readable description for the Subscription **This value should not contain PII.** (optional)
    SinkSid := "SinkSid_example" // string | The SID of the sink that events selected by this subscription should be sent to. Sink must be active for the subscription to be created. (optional)
    Types := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | An array of objects containing the subscribed Event Types (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSubscription(context.Background()).Description(Description).SinkSid(SinkSid).Types(Types).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubscription`: EventsV1Subscription
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSubscription`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSubscriptionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **Description** | **string** | A human readable description for the Subscription **This value should not contain PII.**
 **SinkSid** | **string** | The SID of the sink that events selected by this subscription should be sent to. Sink must be active for the subscription to be created.
 **Types** | **[]map[string]interface{}** | An array of objects containing the subscribed Event Types

### Return type

[**EventsV1Subscription**](EventsV1Subscription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSink

> DeleteSink(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this Sink.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSink(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Sink.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSinkParams struct via the builder pattern


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


## DeleteSubscribedEvent

> DeleteSubscribedEvent(ctx, SubscriptionSid, Type).Execute()





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
    SubscriptionSid := "SubscriptionSid_example" // string | The unique SID identifier of the Subscription.
    Type := "Type_example" // string | Type of event being subscribed to.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSubscribedEvent(context.Background(), SubscriptionSid, Type).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSubscribedEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SubscriptionSid** | **string** | The unique SID identifier of the Subscription.
**Type** | **string** | Type of event being subscribed to.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSubscribedEventParams struct via the builder pattern


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


## DeleteSubscription

> DeleteSubscription(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this Subscription.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSubscription(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Subscription.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSubscriptionParams struct via the builder pattern


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


## FetchEventType

> EventsV1EventType FetchEventType(ctx, Type).Execute()





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
    Type := "Type_example" // string | A string that uniquely identifies this Event Type.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchEventType(context.Background(), Type).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchEventType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchEventType`: EventsV1EventType
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchEventType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Type** | **string** | A string that uniquely identifies this Event Type.

### Other Parameters

Other parameters are passed through a pointer to a FetchEventTypeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**EventsV1EventType**](EventsV1EventType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSchema

> EventsV1Schema FetchSchema(ctx, Id).Execute()





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
    Id := "Id_example" // string | The unique identifier of the schema. Each schema can have multiple versions, that share the same id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSchema(context.Background(), Id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSchema`: EventsV1Schema
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The unique identifier of the schema. Each schema can have multiple versions, that share the same id.

### Other Parameters

Other parameters are passed through a pointer to a FetchSchemaParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**EventsV1Schema**](EventsV1Schema.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSink

> EventsV1Sink FetchSink(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this Sink.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSink(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSink`: EventsV1Sink
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Sink.

### Other Parameters

Other parameters are passed through a pointer to a FetchSinkParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**EventsV1Sink**](EventsV1Sink.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSubscribedEvent

> EventsV1SubscriptionSubscribedEvent FetchSubscribedEvent(ctx, SubscriptionSid, Type).Execute()





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
    SubscriptionSid := "SubscriptionSid_example" // string | The unique SID identifier of the Subscription.
    Type := "Type_example" // string | Type of event being subscribed to.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSubscribedEvent(context.Background(), SubscriptionSid, Type).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSubscribedEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSubscribedEvent`: EventsV1SubscriptionSubscribedEvent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSubscribedEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SubscriptionSid** | **string** | The unique SID identifier of the Subscription.
**Type** | **string** | Type of event being subscribed to.

### Other Parameters

Other parameters are passed through a pointer to a FetchSubscribedEventParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**EventsV1SubscriptionSubscribedEvent**](EventsV1SubscriptionSubscribedEvent.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSubscription

> EventsV1Subscription FetchSubscription(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this Subscription.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSubscription(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSubscription`: EventsV1Subscription
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Subscription.

### Other Parameters

Other parameters are passed through a pointer to a FetchSubscriptionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**EventsV1Subscription**](EventsV1Subscription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchVersion

> EventsV1SchemaVersion FetchVersion(ctx, Id, SchemaVersion).Execute()





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
    Id := "Id_example" // string | The unique identifier of the schema. Each schema can have multiple versions, that share the same id.
    SchemaVersion := int32(56) // int32 | The version of the schema

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchVersion(context.Background(), Id, SchemaVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchVersion`: EventsV1SchemaVersion
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The unique identifier of the schema. Each schema can have multiple versions, that share the same id.
**SchemaVersion** | **int32** | The version of the schema

### Other Parameters

Other parameters are passed through a pointer to a FetchVersionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**EventsV1SchemaVersion**](EventsV1SchemaVersion.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEventType

> ListEventTypeResponse ListEventType(ctx).PageSize(PageSize).Execute()





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
    resp, r, err := api_client.DefaultApi.ListEventType(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListEventType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEventType`: ListEventTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListEventType`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListEventTypeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListEventTypeResponse**](ListEventTypeResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSink

> ListSinkResponse ListSink(ctx).PageSize(PageSize).Execute()





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
    resp, r, err := api_client.DefaultApi.ListSink(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSink`: ListSinkResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSink`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSinkParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSinkResponse**](ListSinkResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscribedEvent

> ListSubscribedEventResponse ListSubscribedEvent(ctx, SubscriptionSid).PageSize(PageSize).Execute()





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
    SubscriptionSid := "SubscriptionSid_example" // string | The unique SID identifier of the Subscription.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSubscribedEvent(context.Background(), SubscriptionSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSubscribedEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSubscribedEvent`: ListSubscribedEventResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSubscribedEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SubscriptionSid** | **string** | The unique SID identifier of the Subscription.

### Other Parameters

Other parameters are passed through a pointer to a ListSubscribedEventParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSubscribedEventResponse**](ListSubscribedEventResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscription

> ListSubscriptionResponse ListSubscription(ctx).SinkSid(SinkSid).PageSize(PageSize).Execute()





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
    SinkSid := "SinkSid_example" // string | The SID of the sink that the list of Subscriptions should be filtered by. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSubscription(context.Background()).SinkSid(SinkSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSubscription`: ListSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSubscription`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSubscriptionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **SinkSid** | **string** | The SID of the sink that the list of Subscriptions should be filtered by.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSubscriptionResponse**](ListSubscriptionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVersion

> ListVersionResponse ListVersion(ctx, Id).PageSize(PageSize).Execute()





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
    Id := "Id_example" // string | The unique identifier of the schema. Each schema can have multiple versions, that share the same id.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListVersion(context.Background(), Id).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVersion`: ListVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The unique identifier of the schema. Each schema can have multiple versions, that share the same id.

### Other Parameters

Other parameters are passed through a pointer to a ListVersionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListVersionResponse**](ListVersionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscribedEvent

> EventsV1SubscriptionSubscribedEvent UpdateSubscribedEvent(ctx, SubscriptionSid, Type).Version(Version).Execute()





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
    SubscriptionSid := "SubscriptionSid_example" // string | The unique SID identifier of the Subscription.
    Type := "Type_example" // string | Type of event being subscribed to.
    Version := int32(56) // int32 | The schema version that the subscription should use. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSubscribedEvent(context.Background(), SubscriptionSid, Type).Version(Version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSubscribedEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSubscribedEvent`: EventsV1SubscriptionSubscribedEvent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSubscribedEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SubscriptionSid** | **string** | The unique SID identifier of the Subscription.
**Type** | **string** | Type of event being subscribed to.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSubscribedEventParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Version** | **int32** | The schema version that the subscription should use.

### Return type

[**EventsV1SubscriptionSubscribedEvent**](EventsV1SubscriptionSubscribedEvent.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscription

> EventsV1Subscription UpdateSubscription(ctx, Sid).Description(Description).SinkSid(SinkSid).Execute()





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
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this Subscription.
    Description := "Description_example" // string | A human readable description for the Subscription. (optional)
    SinkSid := "SinkSid_example" // string | The SID of the sink that events selected by this subscription should be sent to. Sink must be active for the subscription to be created. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSubscription(context.Background(), Sid).Description(Description).SinkSid(SinkSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSubscription`: EventsV1Subscription
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Subscription.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSubscriptionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Description** | **string** | A human readable description for the Subscription.
 **SinkSid** | **string** | The SID of the sink that events selected by this subscription should be sent to. Sink must be active for the subscription to be created.

### Return type

[**EventsV1Subscription**](EventsV1Subscription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

