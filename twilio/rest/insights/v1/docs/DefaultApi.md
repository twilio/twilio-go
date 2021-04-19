# \DefaultApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchCall**](DefaultApi.md#FetchCall) | **Get** /v1/Voice/{Sid} | 
[**FetchSummary**](DefaultApi.md#FetchSummary) | **Get** /v1/Voice/{CallSid}/Summary | 
[**FetchVideoParticipantSummary**](DefaultApi.md#FetchVideoParticipantSummary) | **Get** /v1/Video/Rooms/{RoomSid}/Participants/{ParticipantSid} | 
[**FetchVideoRoomSummary**](DefaultApi.md#FetchVideoRoomSummary) | **Get** /v1/Video/Rooms/{RoomSid} | 
[**ListEvent**](DefaultApi.md#ListEvent) | **Get** /v1/Voice/{CallSid}/Events | 
[**ListMetric**](DefaultApi.md#ListMetric) | **Get** /v1/Voice/{CallSid}/Metrics | 
[**ListVideoParticipantSummary**](DefaultApi.md#ListVideoParticipantSummary) | **Get** /v1/Video/Rooms/{RoomSid}/Participants | 
[**ListVideoRoomSummary**](DefaultApi.md#ListVideoRoomSummary) | **Get** /v1/Video/Rooms | 



## FetchCall

> InsightsV1Call FetchCall(ctx, Sid).Execute()



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
    Sid := "Sid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchCall(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchCall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCall`: InsightsV1Call
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchCallParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**InsightsV1Call**](InsightsV1Call.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSummary

> InsightsV1CallSummary FetchSummary(ctx, CallSid).ProcessingState(ProcessingState).Execute()



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
    CallSid := "CallSid_example" // string | 
    ProcessingState := "ProcessingState_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSummary(context.Background(), CallSid).ProcessingState(ProcessingState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSummary`: InsightsV1CallSummary
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchSummaryParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **ProcessingState** | **string** | 

### Return type

[**InsightsV1CallSummary**](InsightsV1CallSummary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchVideoParticipantSummary

> InsightsV1VideoRoomSummaryVideoParticipantSummary FetchVideoParticipantSummary(ctx, RoomSid, ParticipantSid).Execute()





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
    RoomSid := "RoomSid_example" // string | The SID of the Room resource.
    ParticipantSid := "ParticipantSid_example" // string | The SID of the Participant resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchVideoParticipantSummary(context.Background(), RoomSid, ParticipantSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchVideoParticipantSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchVideoParticipantSummary`: InsightsV1VideoRoomSummaryVideoParticipantSummary
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchVideoParticipantSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room resource.
**ParticipantSid** | **string** | The SID of the Participant resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchVideoParticipantSummaryParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**InsightsV1VideoRoomSummaryVideoParticipantSummary**](InsightsV1VideoRoomSummaryVideoParticipantSummary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchVideoRoomSummary

> InsightsV1VideoRoomSummary FetchVideoRoomSummary(ctx, RoomSid).Execute()





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
    RoomSid := "RoomSid_example" // string | The SID of the Room resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchVideoRoomSummary(context.Background(), RoomSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchVideoRoomSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchVideoRoomSummary`: InsightsV1VideoRoomSummary
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchVideoRoomSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchVideoRoomSummaryParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**InsightsV1VideoRoomSummary**](InsightsV1VideoRoomSummary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvent

> ListEventResponse ListEvent(ctx, CallSid).Edge(Edge).PageSize(PageSize).Execute()



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
    CallSid := "CallSid_example" // string | 
    Edge := "Edge_example" // string |  (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListEvent(context.Background(), CallSid).Edge(Edge).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEvent`: ListEventResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListEventParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Edge** | **string** | 
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListEventResponse**](ListEventResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetric

> ListMetricResponse ListMetric(ctx, CallSid).Edge(Edge).Direction(Direction).PageSize(PageSize).Execute()



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
    CallSid := "CallSid_example" // string | 
    Edge := "Edge_example" // string |  (optional)
    Direction := "Direction_example" // string |  (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListMetric(context.Background(), CallSid).Edge(Edge).Direction(Direction).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMetric`: ListMetricResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListMetricParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Edge** | **string** | 
 **Direction** | **string** | 
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListMetricResponse**](ListMetricResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVideoParticipantSummary

> ListVideoParticipantSummaryResponse ListVideoParticipantSummary(ctx, RoomSid).PageSize(PageSize).Execute()





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
    RoomSid := "RoomSid_example" // string | The SID of the Room resource.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListVideoParticipantSummary(context.Background(), RoomSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListVideoParticipantSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVideoParticipantSummary`: ListVideoParticipantSummaryResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListVideoParticipantSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room resource.

### Other Parameters

Other parameters are passed through a pointer to a ListVideoParticipantSummaryParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListVideoParticipantSummaryResponse**](ListVideoParticipantSummaryResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVideoRoomSummary

> ListVideoRoomSummaryResponse ListVideoRoomSummary(ctx).RoomType(RoomType).Codec(Codec).RoomName(RoomName).CreatedAfter(CreatedAfter).CreatedBefore(CreatedBefore).PageSize(PageSize).Execute()





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
    RoomType := []string{"RoomType_example"} // []string | Type of room. Can be `go`, `peer_to_peer`, `group`, or `group_small`. (optional)
    Codec := []string{"Codec_example"} // []string | Codecs used by participants in the room. Can be `VP8`, `H264`, or `VP9`. (optional)
    RoomName := "RoomName_example" // string | Room friendly name. (optional)
    CreatedAfter := time.Now() // time.Time | Only read rooms that started on or after this ISO 8601 timestamp. (optional)
    CreatedBefore := time.Now() // time.Time | Only read rooms that started before this ISO 8601 timestamp. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListVideoRoomSummary(context.Background()).RoomType(RoomType).Codec(Codec).RoomName(RoomName).CreatedAfter(CreatedAfter).CreatedBefore(CreatedBefore).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListVideoRoomSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVideoRoomSummary`: ListVideoRoomSummaryResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListVideoRoomSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListVideoRoomSummaryParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **RoomType** | **[]string** | Type of room. Can be &#x60;go&#x60;, &#x60;peer_to_peer&#x60;, &#x60;group&#x60;, or &#x60;group_small&#x60;.
 **Codec** | **[]string** | Codecs used by participants in the room. Can be &#x60;VP8&#x60;, &#x60;H264&#x60;, or &#x60;VP9&#x60;.
 **RoomName** | **string** | Room friendly name.
 **CreatedAfter** | **time.Time** | Only read rooms that started on or after this ISO 8601 timestamp.
 **CreatedBefore** | **time.Time** | Only read rooms that started before this ISO 8601 timestamp.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListVideoRoomSummaryResponse**](ListVideoRoomSummaryResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

