# \DefaultApi

All URIs are relative to *https://bulkexports.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExportCustomJob**](DefaultApi.md#CreateExportCustomJob) | **Post** /v1/Exports/{ResourceType}/Jobs | 
[**DeleteJob**](DefaultApi.md#DeleteJob) | **Delete** /v1/Exports/Jobs/{JobSid} | 
[**FetchDay**](DefaultApi.md#FetchDay) | **Get** /v1/Exports/{ResourceType}/Days/{Day} | 
[**FetchExport**](DefaultApi.md#FetchExport) | **Get** /v1/Exports/{ResourceType} | 
[**FetchExportConfiguration**](DefaultApi.md#FetchExportConfiguration) | **Get** /v1/Exports/{ResourceType}/Configuration | 
[**FetchJob**](DefaultApi.md#FetchJob) | **Get** /v1/Exports/Jobs/{JobSid} | 
[**ListDay**](DefaultApi.md#ListDay) | **Get** /v1/Exports/{ResourceType}/Days | 
[**ListExportCustomJob**](DefaultApi.md#ListExportCustomJob) | **Get** /v1/Exports/{ResourceType}/Jobs | 
[**UpdateExportConfiguration**](DefaultApi.md#UpdateExportConfiguration) | **Post** /v1/Exports/{ResourceType}/Configuration | 



## CreateExportCustomJob

> BulkexportsV1ExportExportCustomJob CreateExportCustomJob(ctx, ResourceType).Email(Email).EndDay(EndDay).FriendlyName(FriendlyName).StartDay(StartDay).WebhookMethod(WebhookMethod).WebhookUrl(WebhookUrl).Execute()



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
    ResourceType := "ResourceType_example" // string | The type of communication – Messages or Calls, Conferences, and Participants
    Email := "Email_example" // string | The optional email to send the completion notification to. You can set both webhook, and email, or one or the other. If you set neither, the job will run but you will have to query to determine your job's status. (optional)
    EndDay := "EndDay_example" // string | The end day for the custom export specified as a string in the format of yyyy-mm-dd. End day is inclusive and must be 2 days earlier than the current UTC day. (optional)
    FriendlyName := "FriendlyName_example" // string | The friendly name specified when creating the job (optional)
    StartDay := "StartDay_example" // string | The start day for the custom export specified as a string in the format of yyyy-mm-dd (optional)
    WebhookMethod := "WebhookMethod_example" // string | This is the method used to call the webhook on completion of the job. If this is supplied, `WebhookUrl` must also be supplied. (optional)
    WebhookUrl := "WebhookUrl_example" // string | The optional webhook url called on completion of the job. If this is supplied, `WebhookMethod` must also be supplied. If you set neither webhook nor email, you will have to check your job's status manually. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateExportCustomJob(context.Background(), ResourceType).Email(Email).EndDay(EndDay).FriendlyName(FriendlyName).StartDay(StartDay).WebhookMethod(WebhookMethod).WebhookUrl(WebhookUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateExportCustomJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateExportCustomJob`: BulkexportsV1ExportExportCustomJob
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateExportCustomJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages or Calls, Conferences, and Participants

### Other Parameters

Other parameters are passed through a pointer to a CreateExportCustomJobParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Email** | **string** | The optional email to send the completion notification to. You can set both webhook, and email, or one or the other. If you set neither, the job will run but you will have to query to determine your job&#39;s status.
 **EndDay** | **string** | The end day for the custom export specified as a string in the format of yyyy-mm-dd. End day is inclusive and must be 2 days earlier than the current UTC day.
 **FriendlyName** | **string** | The friendly name specified when creating the job
 **StartDay** | **string** | The start day for the custom export specified as a string in the format of yyyy-mm-dd
 **WebhookMethod** | **string** | This is the method used to call the webhook on completion of the job. If this is supplied, &#x60;WebhookUrl&#x60; must also be supplied.
 **WebhookUrl** | **string** | The optional webhook url called on completion of the job. If this is supplied, &#x60;WebhookMethod&#x60; must also be supplied. If you set neither webhook nor email, you will have to check your job&#39;s status manually.

### Return type

[**BulkexportsV1ExportExportCustomJob**](BulkexportsV1ExportExportCustomJob.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteJob

> DeleteJob(ctx, JobSid).Execute()



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
    JobSid := "JobSid_example" // string | The unique string that that we created to identify the Bulk Export job

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteJob(context.Background(), JobSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**JobSid** | **string** | The unique string that that we created to identify the Bulk Export job

### Other Parameters

Other parameters are passed through a pointer to a DeleteJobParams struct via the builder pattern


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


## FetchDay

> FetchDay(ctx, ResourceType, Day).Execute()





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
    ResourceType := "ResourceType_example" // string | The type of communication – Messages, Calls, Conferences, and Participants
    Day := "Day_example" // string | The ISO 8601 format date of the resources in the file, for a UTC day

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchDay(context.Background(), ResourceType, Day).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchDay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants
**Day** | **string** | The ISO 8601 format date of the resources in the file, for a UTC day

### Other Parameters

Other parameters are passed through a pointer to a FetchDayParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExport

> BulkexportsV1Export FetchExport(ctx, ResourceType).Execute()





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
    ResourceType := "ResourceType_example" // string | The type of communication – Messages, Calls, Conferences, and Participants

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchExport(context.Background(), ResourceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchExport`: BulkexportsV1Export
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants

### Other Parameters

Other parameters are passed through a pointer to a FetchExportParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**BulkexportsV1Export**](BulkexportsV1Export.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExportConfiguration

> BulkexportsV1ExportConfiguration FetchExportConfiguration(ctx, ResourceType).Execute()





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
    ResourceType := "ResourceType_example" // string | The type of communication – Messages, Calls, Conferences, and Participants

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchExportConfiguration(context.Background(), ResourceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchExportConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchExportConfiguration`: BulkexportsV1ExportConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchExportConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants

### Other Parameters

Other parameters are passed through a pointer to a FetchExportConfigurationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**BulkexportsV1ExportConfiguration**](BulkexportsV1ExportConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchJob

> BulkexportsV1ExportJob FetchJob(ctx, JobSid).Execute()



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
    JobSid := "JobSid_example" // string | The unique string that that we created to identify the Bulk Export job

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchJob(context.Background(), JobSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchJob`: BulkexportsV1ExportJob
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**JobSid** | **string** | The unique string that that we created to identify the Bulk Export job

### Other Parameters

Other parameters are passed through a pointer to a FetchJobParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**BulkexportsV1ExportJob**](BulkexportsV1ExportJob.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDay

> ListDayResponse ListDay(ctx, ResourceType).PageSize(PageSize).Execute()





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
    ResourceType := "ResourceType_example" // string | The type of communication – Messages, Calls, Conferences, and Participants
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListDay(context.Background(), ResourceType).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListDay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDay`: ListDayResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListDay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants

### Other Parameters

Other parameters are passed through a pointer to a ListDayParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListDayResponse**](ListDayResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExportCustomJob

> ListExportCustomJobResponse ListExportCustomJob(ctx, ResourceType).PageSize(PageSize).Execute()



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
    ResourceType := "ResourceType_example" // string | The type of communication – Messages, Calls, Conferences, and Participants
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListExportCustomJob(context.Background(), ResourceType).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListExportCustomJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExportCustomJob`: ListExportCustomJobResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListExportCustomJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants

### Other Parameters

Other parameters are passed through a pointer to a ListExportCustomJobParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListExportCustomJobResponse**](ListExportCustomJobResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExportConfiguration

> BulkexportsV1ExportConfiguration UpdateExportConfiguration(ctx, ResourceType).Enabled(Enabled).WebhookMethod(WebhookMethod).WebhookUrl(WebhookUrl).Execute()





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
    ResourceType := "ResourceType_example" // string | The type of communication – Messages, Calls, Conferences, and Participants
    Enabled := true // bool | If true, Twilio will automatically generate every day's file when the day is over. (optional)
    WebhookMethod := "WebhookMethod_example" // string | Sets whether Twilio should call a webhook URL when the automatic generation is complete, using GET or POST. The actual destination is set in the webhook_url (optional)
    WebhookUrl := "WebhookUrl_example" // string | Stores the URL destination for the method specified in webhook_method. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateExportConfiguration(context.Background(), ResourceType).Enabled(Enabled).WebhookMethod(WebhookMethod).WebhookUrl(WebhookUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateExportConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExportConfiguration`: BulkexportsV1ExportConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateExportConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants

### Other Parameters

Other parameters are passed through a pointer to a UpdateExportConfigurationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Enabled** | **bool** | If true, Twilio will automatically generate every day&#39;s file when the day is over.
 **WebhookMethod** | **string** | Sets whether Twilio should call a webhook URL when the automatic generation is complete, using GET or POST. The actual destination is set in the webhook_url
 **WebhookUrl** | **string** | Stores the URL destination for the method specified in webhook_method.

### Return type

[**BulkexportsV1ExportConfiguration**](BulkexportsV1ExportConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

