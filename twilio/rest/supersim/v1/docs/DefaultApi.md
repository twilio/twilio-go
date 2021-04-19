# \DefaultApi

All URIs are relative to *https://supersim.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCommand**](DefaultApi.md#CreateCommand) | **Post** /v1/Commands | 
[**CreateFleet**](DefaultApi.md#CreateFleet) | **Post** /v1/Fleets | 
[**CreateNetworkAccessProfile**](DefaultApi.md#CreateNetworkAccessProfile) | **Post** /v1/NetworkAccessProfiles | 
[**CreateNetworkAccessProfileNetwork**](DefaultApi.md#CreateNetworkAccessProfileNetwork) | **Post** /v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks | 
[**CreateSim**](DefaultApi.md#CreateSim) | **Post** /v1/Sims | 
[**CreateSmsCommand**](DefaultApi.md#CreateSmsCommand) | **Post** /v1/SmsCommands | 
[**DeleteNetworkAccessProfileNetwork**](DefaultApi.md#DeleteNetworkAccessProfileNetwork) | **Delete** /v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks/{Sid} | 
[**FetchCommand**](DefaultApi.md#FetchCommand) | **Get** /v1/Commands/{Sid} | 
[**FetchFleet**](DefaultApi.md#FetchFleet) | **Get** /v1/Fleets/{Sid} | 
[**FetchNetwork**](DefaultApi.md#FetchNetwork) | **Get** /v1/Networks/{Sid} | 
[**FetchNetworkAccessProfile**](DefaultApi.md#FetchNetworkAccessProfile) | **Get** /v1/NetworkAccessProfiles/{Sid} | 
[**FetchNetworkAccessProfileNetwork**](DefaultApi.md#FetchNetworkAccessProfileNetwork) | **Get** /v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks/{Sid} | 
[**FetchSim**](DefaultApi.md#FetchSim) | **Get** /v1/Sims/{Sid} | 
[**FetchSmsCommand**](DefaultApi.md#FetchSmsCommand) | **Get** /v1/SmsCommands/{Sid} | 
[**ListCommand**](DefaultApi.md#ListCommand) | **Get** /v1/Commands | 
[**ListFleet**](DefaultApi.md#ListFleet) | **Get** /v1/Fleets | 
[**ListNetwork**](DefaultApi.md#ListNetwork) | **Get** /v1/Networks | 
[**ListNetworkAccessProfile**](DefaultApi.md#ListNetworkAccessProfile) | **Get** /v1/NetworkAccessProfiles | 
[**ListNetworkAccessProfileNetwork**](DefaultApi.md#ListNetworkAccessProfileNetwork) | **Get** /v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks | 
[**ListSim**](DefaultApi.md#ListSim) | **Get** /v1/Sims | 
[**ListSmsCommand**](DefaultApi.md#ListSmsCommand) | **Get** /v1/SmsCommands | 
[**ListUsageRecord**](DefaultApi.md#ListUsageRecord) | **Get** /v1/UsageRecords | 
[**UpdateFleet**](DefaultApi.md#UpdateFleet) | **Post** /v1/Fleets/{Sid} | 
[**UpdateNetworkAccessProfile**](DefaultApi.md#UpdateNetworkAccessProfile) | **Post** /v1/NetworkAccessProfiles/{Sid} | 
[**UpdateSim**](DefaultApi.md#UpdateSim) | **Post** /v1/Sims/{Sid} | 



## CreateCommand

> SupersimV1Command CreateCommand(ctx).CallbackMethod(CallbackMethod).CallbackUrl(CallbackUrl).Command(Command).Sim(Sim).Execute()





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
    CallbackMethod := "CallbackMethod_example" // string | The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is POST. (optional)
    CallbackUrl := "CallbackUrl_example" // string | The URL we should call using the `callback_method` after we have sent the command. (optional)
    Command := "Command_example" // string | The message body of the command. (optional)
    Sim := "Sim_example" // string | The `sid` or `unique_name` of the [SIM](https://www.twilio.com/docs/wireless/api/sim-resource) to send the Command to. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateCommand(context.Background()).CallbackMethod(CallbackMethod).CallbackUrl(CallbackUrl).Command(Command).Sim(Sim).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCommand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCommand`: SupersimV1Command
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCommand`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateCommandParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **CallbackMethod** | **string** | The HTTP method we should use to call &#x60;callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is POST.
 **CallbackUrl** | **string** | The URL we should call using the &#x60;callback_method&#x60; after we have sent the command.
 **Command** | **string** | The message body of the command.
 **Sim** | **string** | The &#x60;sid&#x60; or &#x60;unique_name&#x60; of the [SIM](https://www.twilio.com/docs/wireless/api/sim-resource) to send the Command to.

### Return type

[**SupersimV1Command**](SupersimV1Command.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFleet

> SupersimV1Fleet CreateFleet(ctx).CommandsEnabled(CommandsEnabled).CommandsMethod(CommandsMethod).CommandsUrl(CommandsUrl).DataEnabled(DataEnabled).DataLimit(DataLimit).NetworkAccessProfile(NetworkAccessProfile).SmsCommandsEnabled(SmsCommandsEnabled).SmsCommandsMethod(SmsCommandsMethod).SmsCommandsUrl(SmsCommandsUrl).UniqueName(UniqueName).Execute()





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
    CommandsEnabled := true // bool | Defines whether SIMs in the Fleet are capable of sending and receiving machine-to-machine SMS via Commands. Defaults to `true`. (optional)
    CommandsMethod := "CommandsMethod_example" // string | A string representing the HTTP method to use when making a request to `commands_url`. Can be one of `POST` or `GET`. Defaults to `POST`. (optional)
    CommandsUrl := "CommandsUrl_example" // string | The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the Commands number. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored. (optional)
    DataEnabled := true // bool | Defines whether SIMs in the Fleet are capable of using 2G/3G/4G/LTE/CAT-M data connectivity. Defaults to `true`. (optional)
    DataLimit := int32(56) // int32 | The total data usage (download and upload combined) in Megabytes that each Sim resource assigned to the Fleet resource can consume during a billing period (normally one month). Value must be between 1MB (1) and 2TB (2,000,000). Defaults to 1GB (1,000). (optional)
    NetworkAccessProfile := "NetworkAccessProfile_example" // string | The SID or unique name of the Network Access Profile that will control which cellular networks the Fleet's SIMs can connect to. (optional)
    SmsCommandsEnabled := true // bool | Defines whether SIMs in the Fleet are capable of sending and receiving machine-to-machine SMS via Commands. Defaults to `true`. (optional)
    SmsCommandsMethod := "SmsCommandsMethod_example" // string | A string representing the HTTP method to use when making a request to `sms_commands_url`. Can be one of `POST` or `GET`. Defaults to `POST`. (optional)
    SmsCommandsUrl := "SmsCommandsUrl_example" // string | The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the SMS Commands number. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateFleet(context.Background()).CommandsEnabled(CommandsEnabled).CommandsMethod(CommandsMethod).CommandsUrl(CommandsUrl).DataEnabled(DataEnabled).DataLimit(DataLimit).NetworkAccessProfile(NetworkAccessProfile).SmsCommandsEnabled(SmsCommandsEnabled).SmsCommandsMethod(SmsCommandsMethod).SmsCommandsUrl(SmsCommandsUrl).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateFleet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFleet`: SupersimV1Fleet
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateFleet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateFleetParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **CommandsEnabled** | **bool** | Defines whether SIMs in the Fleet are capable of sending and receiving machine-to-machine SMS via Commands. Defaults to &#x60;true&#x60;.
 **CommandsMethod** | **string** | A string representing the HTTP method to use when making a request to &#x60;commands_url&#x60;. Can be one of &#x60;POST&#x60; or &#x60;GET&#x60;. Defaults to &#x60;POST&#x60;.
 **CommandsUrl** | **string** | The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the Commands number. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored.
 **DataEnabled** | **bool** | Defines whether SIMs in the Fleet are capable of using 2G/3G/4G/LTE/CAT-M data connectivity. Defaults to &#x60;true&#x60;.
 **DataLimit** | **int32** | The total data usage (download and upload combined) in Megabytes that each Sim resource assigned to the Fleet resource can consume during a billing period (normally one month). Value must be between 1MB (1) and 2TB (2,000,000). Defaults to 1GB (1,000).
 **NetworkAccessProfile** | **string** | The SID or unique name of the Network Access Profile that will control which cellular networks the Fleet&#39;s SIMs can connect to.
 **SmsCommandsEnabled** | **bool** | Defines whether SIMs in the Fleet are capable of sending and receiving machine-to-machine SMS via Commands. Defaults to &#x60;true&#x60;.
 **SmsCommandsMethod** | **string** | A string representing the HTTP method to use when making a request to &#x60;sms_commands_url&#x60;. Can be one of &#x60;POST&#x60; or &#x60;GET&#x60;. Defaults to &#x60;POST&#x60;.
 **SmsCommandsUrl** | **string** | The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the SMS Commands number. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored.
 **UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource&#39;s &#x60;sid&#x60; in the URL to address the resource.

### Return type

[**SupersimV1Fleet**](SupersimV1Fleet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkAccessProfile

> SupersimV1NetworkAccessProfile CreateNetworkAccessProfile(ctx).Networks(Networks).UniqueName(UniqueName).Execute()





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
    Networks := []string{"Inner_example"} // []string | List of Network SIDs that this Network Access Profile will allow connections to. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateNetworkAccessProfile(context.Background()).Networks(Networks).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateNetworkAccessProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkAccessProfile`: SupersimV1NetworkAccessProfile
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateNetworkAccessProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateNetworkAccessProfileParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **Networks** | **[]string** | List of Network SIDs that this Network Access Profile will allow connections to.
 **UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource&#39;s &#x60;sid&#x60; in the URL to address the resource.

### Return type

[**SupersimV1NetworkAccessProfile**](SupersimV1NetworkAccessProfile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkAccessProfileNetwork

> SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork CreateNetworkAccessProfileNetwork(ctx, NetworkAccessProfileSid).Network(Network).Execute()





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
    NetworkAccessProfileSid := "NetworkAccessProfileSid_example" // string | The unique string that identifies the Network Access Profile resource.
    Network := "Network_example" // string | The SID of the Network resource to be added to the Network Access Profile resource. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateNetworkAccessProfileNetwork(context.Background(), NetworkAccessProfileSid).Network(Network).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateNetworkAccessProfileNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkAccessProfileNetwork`: SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateNetworkAccessProfileNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**NetworkAccessProfileSid** | **string** | The unique string that identifies the Network Access Profile resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateNetworkAccessProfileNetworkParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Network** | **string** | The SID of the Network resource to be added to the Network Access Profile resource.

### Return type

[**SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork**](SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSim

> SupersimV1Sim CreateSim(ctx).Iccid(Iccid).RegistrationCode(RegistrationCode).Execute()





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
    Iccid := "Iccid_example" // string | The [ICCID](https://en.wikipedia.org/wiki/Subscriber_identity_module#ICCID) of the Super SIM to be added to your Account. (optional)
    RegistrationCode := "RegistrationCode_example" // string | The 10 digit code required to claim the Super SIM for your Account. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSim(context.Background()).Iccid(Iccid).RegistrationCode(RegistrationCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSim`: SupersimV1Sim
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSim`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSimParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **Iccid** | **string** | The [ICCID](https://en.wikipedia.org/wiki/Subscriber_identity_module#ICCID) of the Super SIM to be added to your Account.
 **RegistrationCode** | **string** | The 10 digit code required to claim the Super SIM for your Account.

### Return type

[**SupersimV1Sim**](SupersimV1Sim.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSmsCommand

> SupersimV1SmsCommand CreateSmsCommand(ctx).CallbackMethod(CallbackMethod).CallbackUrl(CallbackUrl).Payload(Payload).Sim(Sim).Execute()





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
    CallbackMethod := "CallbackMethod_example" // string | The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is POST. (optional)
    CallbackUrl := "CallbackUrl_example" // string | The URL we should call using the `callback_method` after we have sent the command. (optional)
    Payload := "Payload_example" // string | The message body of the SMS Command. (optional)
    Sim := "Sim_example" // string | The `sid` or `unique_name` of the [SIM](https://www.twilio.com/docs/wireless/api/sim-resource) to send the SMS Command to. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSmsCommand(context.Background()).CallbackMethod(CallbackMethod).CallbackUrl(CallbackUrl).Payload(Payload).Sim(Sim).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSmsCommand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSmsCommand`: SupersimV1SmsCommand
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSmsCommand`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSmsCommandParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **CallbackMethod** | **string** | The HTTP method we should use to call &#x60;callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is POST.
 **CallbackUrl** | **string** | The URL we should call using the &#x60;callback_method&#x60; after we have sent the command.
 **Payload** | **string** | The message body of the SMS Command.
 **Sim** | **string** | The &#x60;sid&#x60; or &#x60;unique_name&#x60; of the [SIM](https://www.twilio.com/docs/wireless/api/sim-resource) to send the SMS Command to.

### Return type

[**SupersimV1SmsCommand**](SupersimV1SmsCommand.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkAccessProfileNetwork

> DeleteNetworkAccessProfileNetwork(ctx, NetworkAccessProfileSid, Sid).Execute()





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
    NetworkAccessProfileSid := "NetworkAccessProfileSid_example" // string | The unique string that identifies the Network Access Profile resource.
    Sid := "Sid_example" // string | The SID of the Network resource to be removed from the Network Access Profile resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteNetworkAccessProfileNetwork(context.Background(), NetworkAccessProfileSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteNetworkAccessProfileNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**NetworkAccessProfileSid** | **string** | The unique string that identifies the Network Access Profile resource.
**Sid** | **string** | The SID of the Network resource to be removed from the Network Access Profile resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteNetworkAccessProfileNetworkParams struct via the builder pattern


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


## FetchCommand

> SupersimV1Command FetchCommand(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | The SID of the Command resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchCommand(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchCommand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCommand`: SupersimV1Command
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchCommand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Command resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchCommandParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**SupersimV1Command**](SupersimV1Command.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFleet

> SupersimV1Fleet FetchFleet(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | The SID of the Fleet resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchFleet(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchFleet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchFleet`: SupersimV1Fleet
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchFleet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Fleet resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFleetParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**SupersimV1Fleet**](SupersimV1Fleet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchNetwork

> SupersimV1Network FetchNetwork(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | The SID of the Network resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchNetwork(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchNetwork`: SupersimV1Network
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Network resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchNetworkParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**SupersimV1Network**](SupersimV1Network.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchNetworkAccessProfile

> SupersimV1NetworkAccessProfile FetchNetworkAccessProfile(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | The SID of the Network Access Profile resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchNetworkAccessProfile(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchNetworkAccessProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchNetworkAccessProfile`: SupersimV1NetworkAccessProfile
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchNetworkAccessProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Network Access Profile resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchNetworkAccessProfileParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**SupersimV1NetworkAccessProfile**](SupersimV1NetworkAccessProfile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchNetworkAccessProfileNetwork

> SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork FetchNetworkAccessProfileNetwork(ctx, NetworkAccessProfileSid, Sid).Execute()





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
    NetworkAccessProfileSid := "NetworkAccessProfileSid_example" // string | The unique string that identifies the Network Access Profile resource.
    Sid := "Sid_example" // string | The SID of the Network resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchNetworkAccessProfileNetwork(context.Background(), NetworkAccessProfileSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchNetworkAccessProfileNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchNetworkAccessProfileNetwork`: SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchNetworkAccessProfileNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**NetworkAccessProfileSid** | **string** | The unique string that identifies the Network Access Profile resource.
**Sid** | **string** | The SID of the Network resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchNetworkAccessProfileNetworkParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork**](SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSim

> SupersimV1Sim FetchSim(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | The SID of the Sim resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSim(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSim`: SupersimV1Sim
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Sim resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSimParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**SupersimV1Sim**](SupersimV1Sim.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSmsCommand

> SupersimV1SmsCommand FetchSmsCommand(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | The SID of the SMS Command resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSmsCommand(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSmsCommand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSmsCommand`: SupersimV1SmsCommand
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSmsCommand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the SMS Command resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSmsCommandParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**SupersimV1SmsCommand**](SupersimV1SmsCommand.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommand

> ListCommandResponse ListCommand(ctx).Sim(Sim).Status(Status).Direction(Direction).PageSize(PageSize).Execute()





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
    Sim := "Sim_example" // string | The SID or unique name of the Sim that Command was sent to or from. (optional)
    Status := "Status_example" // string | The status of the Command. Can be: `queued`, `sent`, `delivered`, `received` or `failed`. See the [Command Status Values](https://www.twilio.com/docs/wireless/api/command-resource#status-values) for a description of each. (optional)
    Direction := "Direction_example" // string | The direction of the Command. Can be `to_sim` or `from_sim`. The value of `to_sim` is synonymous with the term `mobile terminated`, and `from_sim` is synonymous with the term `mobile originated`. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListCommand(context.Background()).Sim(Sim).Status(Status).Direction(Direction).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCommand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCommand`: ListCommandResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCommand`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListCommandParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **Sim** | **string** | The SID or unique name of the Sim that Command was sent to or from.
 **Status** | **string** | The status of the Command. Can be: &#x60;queued&#x60;, &#x60;sent&#x60;, &#x60;delivered&#x60;, &#x60;received&#x60; or &#x60;failed&#x60;. See the [Command Status Values](https://www.twilio.com/docs/wireless/api/command-resource#status-values) for a description of each.
 **Direction** | **string** | The direction of the Command. Can be &#x60;to_sim&#x60; or &#x60;from_sim&#x60;. The value of &#x60;to_sim&#x60; is synonymous with the term &#x60;mobile terminated&#x60;, and &#x60;from_sim&#x60; is synonymous with the term &#x60;mobile originated&#x60;.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListCommandResponse**](ListCommandResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFleet

> ListFleetResponse ListFleet(ctx).NetworkAccessProfile(NetworkAccessProfile).PageSize(PageSize).Execute()





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
    NetworkAccessProfile := "NetworkAccessProfile_example" // string | The SID or unique name of the Network Access Profile that controls which cellular networks the Fleet's SIMs can connect to. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListFleet(context.Background()).NetworkAccessProfile(NetworkAccessProfile).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListFleet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFleet`: ListFleetResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListFleet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListFleetParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **NetworkAccessProfile** | **string** | The SID or unique name of the Network Access Profile that controls which cellular networks the Fleet&#39;s SIMs can connect to.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListFleetResponse**](ListFleetResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetwork

> ListNetworkResponse ListNetwork(ctx).IsoCountry(IsoCountry).Mcc(Mcc).Mnc(Mnc).PageSize(PageSize).Execute()





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
    IsoCountry := "IsoCountry_example" // string | The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the Network resources to read. (optional)
    Mcc := "Mcc_example" // string | The 'mobile country code' of a country. Network resources with this `mcc` in their `identifiers` will be read. (optional)
    Mnc := "Mnc_example" // string | The 'mobile network code' of a mobile operator network. Network resources with this `mnc` in their `identifiers` will be read. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListNetwork(context.Background()).IsoCountry(IsoCountry).Mcc(Mcc).Mnc(Mnc).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetwork`: ListNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListNetwork`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListNetworkParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **IsoCountry** | **string** | The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the Network resources to read.
 **Mcc** | **string** | The &#39;mobile country code&#39; of a country. Network resources with this &#x60;mcc&#x60; in their &#x60;identifiers&#x60; will be read.
 **Mnc** | **string** | The &#39;mobile network code&#39; of a mobile operator network. Network resources with this &#x60;mnc&#x60; in their &#x60;identifiers&#x60; will be read.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListNetworkResponse**](ListNetworkResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkAccessProfile

> ListNetworkAccessProfileResponse ListNetworkAccessProfile(ctx).PageSize(PageSize).Execute()





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
    resp, r, err := api_client.DefaultApi.ListNetworkAccessProfile(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListNetworkAccessProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkAccessProfile`: ListNetworkAccessProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListNetworkAccessProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListNetworkAccessProfileParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListNetworkAccessProfileResponse**](ListNetworkAccessProfileResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkAccessProfileNetwork

> ListNetworkAccessProfileNetworkResponse ListNetworkAccessProfileNetwork(ctx, NetworkAccessProfileSid).PageSize(PageSize).Execute()





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
    NetworkAccessProfileSid := "NetworkAccessProfileSid_example" // string | The unique string that identifies the Network Access Profile resource.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListNetworkAccessProfileNetwork(context.Background(), NetworkAccessProfileSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListNetworkAccessProfileNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkAccessProfileNetwork`: ListNetworkAccessProfileNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListNetworkAccessProfileNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**NetworkAccessProfileSid** | **string** | The unique string that identifies the Network Access Profile resource.

### Other Parameters

Other parameters are passed through a pointer to a ListNetworkAccessProfileNetworkParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListNetworkAccessProfileNetworkResponse**](ListNetworkAccessProfileNetworkResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSim

> ListSimResponse ListSim(ctx).Status(Status).Fleet(Fleet).Iccid(Iccid).PageSize(PageSize).Execute()





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
    Status := "Status_example" // string | The status of the Sim resources to read. Can be `new`, `ready`, `active`, `inactive`, or `scheduled`. (optional)
    Fleet := "Fleet_example" // string | The SID or unique name of the Fleet to which a list of Sims are assigned. (optional)
    Iccid := "Iccid_example" // string | The [ICCID](https://en.wikipedia.org/wiki/Subscriber_identity_module#ICCID) associated with a Super SIM to filter the list by. Passing this parameter will always return a list containing zero or one SIMs. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSim(context.Background()).Status(Status).Fleet(Fleet).Iccid(Iccid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSim`: ListSimResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSim`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSimParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **Status** | **string** | The status of the Sim resources to read. Can be &#x60;new&#x60;, &#x60;ready&#x60;, &#x60;active&#x60;, &#x60;inactive&#x60;, or &#x60;scheduled&#x60;.
 **Fleet** | **string** | The SID or unique name of the Fleet to which a list of Sims are assigned.
 **Iccid** | **string** | The [ICCID](https://en.wikipedia.org/wiki/Subscriber_identity_module#ICCID) associated with a Super SIM to filter the list by. Passing this parameter will always return a list containing zero or one SIMs.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSimResponse**](ListSimResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSmsCommand

> ListSmsCommandResponse ListSmsCommand(ctx).Sim(Sim).Status(Status).Direction(Direction).PageSize(PageSize).Execute()





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
    Sim := "Sim_example" // string | The SID or unique name of the Sim that SMS Command was sent to or from. (optional)
    Status := "Status_example" // string | The status of the SMS Command. Can be: `queued`, `sent`, `delivered`, `received` or `failed`. See the [SMS Command Status Values](https://www.twilio.com/docs/wireless/api/smscommand-resource#status-values) for a description of each. (optional)
    Direction := "Direction_example" // string | The direction of the SMS Command. Can be `to_sim` or `from_sim`. The value of `to_sim` is synonymous with the term `mobile terminated`, and `from_sim` is synonymous with the term `mobile originated`. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSmsCommand(context.Background()).Sim(Sim).Status(Status).Direction(Direction).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSmsCommand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSmsCommand`: ListSmsCommandResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSmsCommand`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSmsCommandParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **Sim** | **string** | The SID or unique name of the Sim that SMS Command was sent to or from.
 **Status** | **string** | The status of the SMS Command. Can be: &#x60;queued&#x60;, &#x60;sent&#x60;, &#x60;delivered&#x60;, &#x60;received&#x60; or &#x60;failed&#x60;. See the [SMS Command Status Values](https://www.twilio.com/docs/wireless/api/smscommand-resource#status-values) for a description of each.
 **Direction** | **string** | The direction of the SMS Command. Can be &#x60;to_sim&#x60; or &#x60;from_sim&#x60;. The value of &#x60;to_sim&#x60; is synonymous with the term &#x60;mobile terminated&#x60;, and &#x60;from_sim&#x60; is synonymous with the term &#x60;mobile originated&#x60;.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSmsCommandResponse**](ListSmsCommandResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecord

> ListUsageRecordResponse ListUsageRecord(ctx).Sim(Sim).Fleet(Fleet).Network(Network).IsoCountry(IsoCountry).Group(Group).Granularity(Granularity).StartTime(StartTime).EndTime(EndTime).PageSize(PageSize).Execute()





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
    Sim := "Sim_example" // string | SID or unique name of a Sim resource. Only show UsageRecords representing usage incurred by this Super SIM. (optional)
    Fleet := "Fleet_example" // string | SID or unique name of a Fleet resource. Only show UsageRecords representing usage for Super SIMs belonging to this Fleet resource at the time the usage occurred. (optional)
    Network := "Network_example" // string | SID of a Network resource. Only show UsageRecords representing usage on this network. (optional)
    IsoCountry := "IsoCountry_example" // string | Alpha-2 ISO Country Code. Only show UsageRecords representing usage in this country. (optional)
    Group := "Group_example" // string | Dimension over which to aggregate usage records. Can be: `sim`, `fleet`, `network`, `isoCountry`. Default is to not aggregate across any of these dimensions, UsageRecords will be aggregated into the time buckets described by the `Granularity` parameter. (optional)
    Granularity := "Granularity_example" // string | Time-based grouping that UsageRecords should be aggregated by. Can be: `hour`, `day`, or `all`. Default is `all`. `all` returns one UsageRecord that describes the usage for the entire period. (optional)
    StartTime := time.Now() // time.Time | Only include usage that occurred at or after this time, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Default is one month before the `end_time`. (optional)
    EndTime := time.Now() // time.Time | Only include usage that occurred before this time, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Default is the current time. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListUsageRecord(context.Background()).Sim(Sim).Fleet(Fleet).Network(Network).IsoCountry(IsoCountry).Group(Group).Granularity(Granularity).StartTime(StartTime).EndTime(EndTime).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsageRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsageRecord`: ListUsageRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsageRecord`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListUsageRecordParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **Sim** | **string** | SID or unique name of a Sim resource. Only show UsageRecords representing usage incurred by this Super SIM.
 **Fleet** | **string** | SID or unique name of a Fleet resource. Only show UsageRecords representing usage for Super SIMs belonging to this Fleet resource at the time the usage occurred.
 **Network** | **string** | SID of a Network resource. Only show UsageRecords representing usage on this network.
 **IsoCountry** | **string** | Alpha-2 ISO Country Code. Only show UsageRecords representing usage in this country.
 **Group** | **string** | Dimension over which to aggregate usage records. Can be: &#x60;sim&#x60;, &#x60;fleet&#x60;, &#x60;network&#x60;, &#x60;isoCountry&#x60;. Default is to not aggregate across any of these dimensions, UsageRecords will be aggregated into the time buckets described by the &#x60;Granularity&#x60; parameter.
 **Granularity** | **string** | Time-based grouping that UsageRecords should be aggregated by. Can be: &#x60;hour&#x60;, &#x60;day&#x60;, or &#x60;all&#x60;. Default is &#x60;all&#x60;. &#x60;all&#x60; returns one UsageRecord that describes the usage for the entire period.
 **StartTime** | **time.Time** | Only include usage that occurred at or after this time, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Default is one month before the &#x60;end_time&#x60;.
 **EndTime** | **time.Time** | Only include usage that occurred before this time, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Default is the current time.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListUsageRecordResponse**](ListUsageRecordResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFleet

> SupersimV1Fleet UpdateFleet(ctx, Sid).CommandsMethod(CommandsMethod).CommandsUrl(CommandsUrl).NetworkAccessProfile(NetworkAccessProfile).SmsCommandsMethod(SmsCommandsMethod).SmsCommandsUrl(SmsCommandsUrl).UniqueName(UniqueName).Execute()





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
    Sid := "Sid_example" // string | The SID of the Fleet resource to update.
    CommandsMethod := "CommandsMethod_example" // string | A string representing the HTTP method to use when making a request to `commands_url`. Can be one of `POST` or `GET`. Defaults to `POST`. (optional)
    CommandsUrl := "CommandsUrl_example" // string | The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the Commands number. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored. (optional)
    NetworkAccessProfile := "NetworkAccessProfile_example" // string | The SID or unique name of the Network Access Profile that will control which cellular networks the Fleet's SIMs can connect to. (optional)
    SmsCommandsMethod := "SmsCommandsMethod_example" // string | A string representing the HTTP method to use when making a request to `sms_commands_url`. Can be one of `POST` or `GET`. Defaults to `POST`. (optional)
    SmsCommandsUrl := "SmsCommandsUrl_example" // string | The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the SMS Commands number. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateFleet(context.Background(), Sid).CommandsMethod(CommandsMethod).CommandsUrl(CommandsUrl).NetworkAccessProfile(NetworkAccessProfile).SmsCommandsMethod(SmsCommandsMethod).SmsCommandsUrl(SmsCommandsUrl).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateFleet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFleet`: SupersimV1Fleet
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateFleet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Fleet resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateFleetParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **CommandsMethod** | **string** | A string representing the HTTP method to use when making a request to &#x60;commands_url&#x60;. Can be one of &#x60;POST&#x60; or &#x60;GET&#x60;. Defaults to &#x60;POST&#x60;.
 **CommandsUrl** | **string** | The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the Commands number. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored.
 **NetworkAccessProfile** | **string** | The SID or unique name of the Network Access Profile that will control which cellular networks the Fleet&#39;s SIMs can connect to.
 **SmsCommandsMethod** | **string** | A string representing the HTTP method to use when making a request to &#x60;sms_commands_url&#x60;. Can be one of &#x60;POST&#x60; or &#x60;GET&#x60;. Defaults to &#x60;POST&#x60;.
 **SmsCommandsUrl** | **string** | The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the SMS Commands number. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored.
 **UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource&#39;s &#x60;sid&#x60; in the URL to address the resource.

### Return type

[**SupersimV1Fleet**](SupersimV1Fleet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkAccessProfile

> SupersimV1NetworkAccessProfile UpdateNetworkAccessProfile(ctx, Sid).UniqueName(UniqueName).Execute()





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
    Sid := "Sid_example" // string | The SID of the Network Access Profile to update.
    UniqueName := "UniqueName_example" // string | The new unique name of the Network Access Profile. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateNetworkAccessProfile(context.Background(), Sid).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateNetworkAccessProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkAccessProfile`: SupersimV1NetworkAccessProfile
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateNetworkAccessProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Network Access Profile to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateNetworkAccessProfileParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **UniqueName** | **string** | The new unique name of the Network Access Profile.

### Return type

[**SupersimV1NetworkAccessProfile**](SupersimV1NetworkAccessProfile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSim

> SupersimV1Sim UpdateSim(ctx, Sid).AccountSid(AccountSid).CallbackMethod(CallbackMethod).CallbackUrl(CallbackUrl).Fleet(Fleet).Status(Status).UniqueName(UniqueName).Execute()





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
    Sid := "Sid_example" // string | The SID of the Sim resource to update.
    AccountSid := "AccountSid_example" // string | The SID of the Account to which the Sim resource should belong. The Account SID can only be that of the requesting Account or that of a Subaccount of the requesting Account. Only valid when the Sim resource's status is new. (optional)
    CallbackMethod := "CallbackMethod_example" // string | The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is POST. (optional)
    CallbackUrl := "CallbackUrl_example" // string | The URL we should call using the `callback_method` after an asynchronous update has finished. (optional)
    Fleet := "Fleet_example" // string | The SID or unique name of the Fleet to which the SIM resource should be assigned. (optional)
    Status := "Status_example" // string | The new status of the resource. Can be: `ready`, `active`, or `inactive`. See the [Super SIM Status Values](https://www.twilio.com/docs/iot/supersim/api/sim-resource#status-values) for more info. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSim(context.Background(), Sid).AccountSid(AccountSid).CallbackMethod(CallbackMethod).CallbackUrl(CallbackUrl).Fleet(Fleet).Status(Status).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSim`: SupersimV1Sim
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Sim resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSimParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **AccountSid** | **string** | The SID of the Account to which the Sim resource should belong. The Account SID can only be that of the requesting Account or that of a Subaccount of the requesting Account. Only valid when the Sim resource&#39;s status is new.
 **CallbackMethod** | **string** | The HTTP method we should use to call &#x60;callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is POST.
 **CallbackUrl** | **string** | The URL we should call using the &#x60;callback_method&#x60; after an asynchronous update has finished.
 **Fleet** | **string** | The SID or unique name of the Fleet to which the SIM resource should be assigned.
 **Status** | **string** | The new status of the resource. Can be: &#x60;ready&#x60;, &#x60;active&#x60;, or &#x60;inactive&#x60;. See the [Super SIM Status Values](https://www.twilio.com/docs/iot/supersim/api/sim-resource#status-values) for more info.
 **UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource&#39;s &#x60;sid&#x60; in the URL to address the resource.

### Return type

[**SupersimV1Sim**](SupersimV1Sim.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

