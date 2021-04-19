# \DefaultApi

All URIs are relative to *https://ip-messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChannel**](DefaultApi.md#CreateChannel) | **Post** /v2/Services/{ServiceSid}/Channels | 
[**CreateChannelWebhook**](DefaultApi.md#CreateChannelWebhook) | **Post** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Webhooks | 
[**CreateCredential**](DefaultApi.md#CreateCredential) | **Post** /v2/Credentials | 
[**CreateInvite**](DefaultApi.md#CreateInvite) | **Post** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Invites | 
[**CreateMember**](DefaultApi.md#CreateMember) | **Post** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Members | 
[**CreateMessage**](DefaultApi.md#CreateMessage) | **Post** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Messages | 
[**CreateRole**](DefaultApi.md#CreateRole) | **Post** /v2/Services/{ServiceSid}/Roles | 
[**CreateService**](DefaultApi.md#CreateService) | **Post** /v2/Services | 
[**CreateUser**](DefaultApi.md#CreateUser) | **Post** /v2/Services/{ServiceSid}/Users | 
[**DeleteBinding**](DefaultApi.md#DeleteBinding) | **Delete** /v2/Services/{ServiceSid}/Bindings/{Sid} | 
[**DeleteChannel**](DefaultApi.md#DeleteChannel) | **Delete** /v2/Services/{ServiceSid}/Channels/{Sid} | 
[**DeleteChannelWebhook**](DefaultApi.md#DeleteChannelWebhook) | **Delete** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Webhooks/{Sid} | 
[**DeleteCredential**](DefaultApi.md#DeleteCredential) | **Delete** /v2/Credentials/{Sid} | 
[**DeleteInvite**](DefaultApi.md#DeleteInvite) | **Delete** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Invites/{Sid} | 
[**DeleteMember**](DefaultApi.md#DeleteMember) | **Delete** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Members/{Sid} | 
[**DeleteMessage**](DefaultApi.md#DeleteMessage) | **Delete** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Messages/{Sid} | 
[**DeleteRole**](DefaultApi.md#DeleteRole) | **Delete** /v2/Services/{ServiceSid}/Roles/{Sid} | 
[**DeleteService**](DefaultApi.md#DeleteService) | **Delete** /v2/Services/{Sid} | 
[**DeleteUser**](DefaultApi.md#DeleteUser) | **Delete** /v2/Services/{ServiceSid}/Users/{Sid} | 
[**DeleteUserBinding**](DefaultApi.md#DeleteUserBinding) | **Delete** /v2/Services/{ServiceSid}/Users/{UserSid}/Bindings/{Sid} | 
[**DeleteUserChannel**](DefaultApi.md#DeleteUserChannel) | **Delete** /v2/Services/{ServiceSid}/Users/{UserSid}/Channels/{ChannelSid} | 
[**FetchBinding**](DefaultApi.md#FetchBinding) | **Get** /v2/Services/{ServiceSid}/Bindings/{Sid} | 
[**FetchChannel**](DefaultApi.md#FetchChannel) | **Get** /v2/Services/{ServiceSid}/Channels/{Sid} | 
[**FetchChannelWebhook**](DefaultApi.md#FetchChannelWebhook) | **Get** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Webhooks/{Sid} | 
[**FetchCredential**](DefaultApi.md#FetchCredential) | **Get** /v2/Credentials/{Sid} | 
[**FetchInvite**](DefaultApi.md#FetchInvite) | **Get** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Invites/{Sid} | 
[**FetchMember**](DefaultApi.md#FetchMember) | **Get** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Members/{Sid} | 
[**FetchMessage**](DefaultApi.md#FetchMessage) | **Get** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Messages/{Sid} | 
[**FetchRole**](DefaultApi.md#FetchRole) | **Get** /v2/Services/{ServiceSid}/Roles/{Sid} | 
[**FetchService**](DefaultApi.md#FetchService) | **Get** /v2/Services/{Sid} | 
[**FetchUser**](DefaultApi.md#FetchUser) | **Get** /v2/Services/{ServiceSid}/Users/{Sid} | 
[**FetchUserBinding**](DefaultApi.md#FetchUserBinding) | **Get** /v2/Services/{ServiceSid}/Users/{UserSid}/Bindings/{Sid} | 
[**FetchUserChannel**](DefaultApi.md#FetchUserChannel) | **Get** /v2/Services/{ServiceSid}/Users/{UserSid}/Channels/{ChannelSid} | 
[**ListBinding**](DefaultApi.md#ListBinding) | **Get** /v2/Services/{ServiceSid}/Bindings | 
[**ListChannel**](DefaultApi.md#ListChannel) | **Get** /v2/Services/{ServiceSid}/Channels | 
[**ListChannelWebhook**](DefaultApi.md#ListChannelWebhook) | **Get** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Webhooks | 
[**ListCredential**](DefaultApi.md#ListCredential) | **Get** /v2/Credentials | 
[**ListInvite**](DefaultApi.md#ListInvite) | **Get** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Invites | 
[**ListMember**](DefaultApi.md#ListMember) | **Get** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Members | 
[**ListMessage**](DefaultApi.md#ListMessage) | **Get** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Messages | 
[**ListRole**](DefaultApi.md#ListRole) | **Get** /v2/Services/{ServiceSid}/Roles | 
[**ListService**](DefaultApi.md#ListService) | **Get** /v2/Services | 
[**ListUser**](DefaultApi.md#ListUser) | **Get** /v2/Services/{ServiceSid}/Users | 
[**ListUserBinding**](DefaultApi.md#ListUserBinding) | **Get** /v2/Services/{ServiceSid}/Users/{UserSid}/Bindings | 
[**ListUserChannel**](DefaultApi.md#ListUserChannel) | **Get** /v2/Services/{ServiceSid}/Users/{UserSid}/Channels | 
[**UpdateChannel**](DefaultApi.md#UpdateChannel) | **Post** /v2/Services/{ServiceSid}/Channels/{Sid} | 
[**UpdateChannelWebhook**](DefaultApi.md#UpdateChannelWebhook) | **Post** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Webhooks/{Sid} | 
[**UpdateCredential**](DefaultApi.md#UpdateCredential) | **Post** /v2/Credentials/{Sid} | 
[**UpdateMember**](DefaultApi.md#UpdateMember) | **Post** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Members/{Sid} | 
[**UpdateMessage**](DefaultApi.md#UpdateMessage) | **Post** /v2/Services/{ServiceSid}/Channels/{ChannelSid}/Messages/{Sid} | 
[**UpdateRole**](DefaultApi.md#UpdateRole) | **Post** /v2/Services/{ServiceSid}/Roles/{Sid} | 
[**UpdateService**](DefaultApi.md#UpdateService) | **Post** /v2/Services/{Sid} | 
[**UpdateUser**](DefaultApi.md#UpdateUser) | **Post** /v2/Services/{ServiceSid}/Users/{Sid} | 
[**UpdateUserChannel**](DefaultApi.md#UpdateUserChannel) | **Post** /v2/Services/{ServiceSid}/Users/{UserSid}/Channels/{ChannelSid} | 



## CreateChannel

> IpMessagingV2ServiceChannel CreateChannel(ctx, ServiceSid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).CreatedBy(CreatedBy).DateCreated(DateCreated).DateUpdated(DateUpdated).FriendlyName(FriendlyName).Type(Type).UniqueName(UniqueName).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)
    Attributes := "Attributes_example" // string |  (optional)
    CreatedBy := "CreatedBy_example" // string |  (optional)
    DateCreated := time.Now() // time.Time |  (optional)
    DateUpdated := time.Now() // time.Time |  (optional)
    FriendlyName := "FriendlyName_example" // string |  (optional)
    Type := "Type_example" // string |  (optional)
    UniqueName := "UniqueName_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateChannel(context.Background(), ServiceSid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).CreatedBy(CreatedBy).DateCreated(DateCreated).DateUpdated(DateUpdated).FriendlyName(FriendlyName).Type(Type).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateChannel`: IpMessagingV2ServiceChannel
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
 **Attributes** | **string** | 
 **CreatedBy** | **string** | 
 **DateCreated** | **time.Time** | 
 **DateUpdated** | **time.Time** | 
 **FriendlyName** | **string** | 
 **Type** | **string** | 
 **UniqueName** | **string** | 

### Return type

[**IpMessagingV2ServiceChannel**](IpMessagingV2ServiceChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateChannelWebhook

> IpMessagingV2ServiceChannelChannelWebhook CreateChannelWebhook(ctx, ServiceSid, ChannelSid).ConfigurationFilters(ConfigurationFilters).ConfigurationFlowSid(ConfigurationFlowSid).ConfigurationMethod(ConfigurationMethod).ConfigurationRetryCount(ConfigurationRetryCount).ConfigurationTriggers(ConfigurationTriggers).ConfigurationUrl(ConfigurationUrl).Type(Type).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    ChannelSid := "ChannelSid_example" // string | 
    ConfigurationFilters := []string{"Inner_example"} // []string |  (optional)
    ConfigurationFlowSid := "ConfigurationFlowSid_example" // string |  (optional)
    ConfigurationMethod := "ConfigurationMethod_example" // string |  (optional)
    ConfigurationRetryCount := int32(56) // int32 |  (optional)
    ConfigurationTriggers := []string{"Inner_example"} // []string |  (optional)
    ConfigurationUrl := "ConfigurationUrl_example" // string |  (optional)
    Type := "Type_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateChannelWebhook(context.Background(), ServiceSid, ChannelSid).ConfigurationFilters(ConfigurationFilters).ConfigurationFlowSid(ConfigurationFlowSid).ConfigurationMethod(ConfigurationMethod).ConfigurationRetryCount(ConfigurationRetryCount).ConfigurationTriggers(ConfigurationTriggers).ConfigurationUrl(ConfigurationUrl).Type(Type).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateChannelWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateChannelWebhook`: IpMessagingV2ServiceChannelChannelWebhook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateChannelWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateChannelWebhookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **ConfigurationFilters** | **[]string** | 
 **ConfigurationFlowSid** | **string** | 
 **ConfigurationMethod** | **string** | 
 **ConfigurationRetryCount** | **int32** | 
 **ConfigurationTriggers** | **[]string** | 
 **ConfigurationUrl** | **string** | 
 **Type** | **string** | 

### Return type

[**IpMessagingV2ServiceChannelChannelWebhook**](IpMessagingV2ServiceChannelChannelWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCredential

> IpMessagingV2Credential CreateCredential(ctx).ApiKey(ApiKey).Certificate(Certificate).FriendlyName(FriendlyName).PrivateKey(PrivateKey).Sandbox(Sandbox).Secret(Secret).Type(Type).Execute()



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
    ApiKey := "ApiKey_example" // string |  (optional)
    Certificate := "Certificate_example" // string |  (optional)
    FriendlyName := "FriendlyName_example" // string |  (optional)
    PrivateKey := "PrivateKey_example" // string |  (optional)
    Sandbox := true // bool |  (optional)
    Secret := "Secret_example" // string |  (optional)
    Type := "Type_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateCredential(context.Background()).ApiKey(ApiKey).Certificate(Certificate).FriendlyName(FriendlyName).PrivateKey(PrivateKey).Sandbox(Sandbox).Secret(Secret).Type(Type).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCredential`: IpMessagingV2Credential
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCredential`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateCredentialParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **ApiKey** | **string** | 
 **Certificate** | **string** | 
 **FriendlyName** | **string** | 
 **PrivateKey** | **string** | 
 **Sandbox** | **bool** | 
 **Secret** | **string** | 
 **Type** | **string** | 

### Return type

[**IpMessagingV2Credential**](IpMessagingV2Credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvite

> IpMessagingV2ServiceChannelInvite CreateInvite(ctx, ServiceSid, ChannelSid).Identity(Identity).RoleSid(RoleSid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    ChannelSid := "ChannelSid_example" // string | 
    Identity := "Identity_example" // string |  (optional)
    RoleSid := "RoleSid_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateInvite(context.Background(), ServiceSid, ChannelSid).Identity(Identity).RoleSid(RoleSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInvite`: IpMessagingV2ServiceChannelInvite
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateInviteParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Identity** | **string** | 
 **RoleSid** | **string** | 

### Return type

[**IpMessagingV2ServiceChannelInvite**](IpMessagingV2ServiceChannelInvite.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMember

> IpMessagingV2ServiceChannelMember CreateMember(ctx, ServiceSid, ChannelSid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).DateCreated(DateCreated).DateUpdated(DateUpdated).Identity(Identity).LastConsumedMessageIndex(LastConsumedMessageIndex).LastConsumptionTimestamp(LastConsumptionTimestamp).RoleSid(RoleSid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    ChannelSid := "ChannelSid_example" // string | 
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)
    Attributes := "Attributes_example" // string |  (optional)
    DateCreated := time.Now() // time.Time |  (optional)
    DateUpdated := time.Now() // time.Time |  (optional)
    Identity := "Identity_example" // string |  (optional)
    LastConsumedMessageIndex := int32(56) // int32 |  (optional)
    LastConsumptionTimestamp := time.Now() // time.Time |  (optional)
    RoleSid := "RoleSid_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateMember(context.Background(), ServiceSid, ChannelSid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).DateCreated(DateCreated).DateUpdated(DateUpdated).Identity(Identity).LastConsumedMessageIndex(LastConsumedMessageIndex).LastConsumptionTimestamp(LastConsumptionTimestamp).RoleSid(RoleSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMember`: IpMessagingV2ServiceChannelMember
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateMemberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
 **Attributes** | **string** | 
 **DateCreated** | **time.Time** | 
 **DateUpdated** | **time.Time** | 
 **Identity** | **string** | 
 **LastConsumedMessageIndex** | **int32** | 
 **LastConsumptionTimestamp** | **time.Time** | 
 **RoleSid** | **string** | 

### Return type

[**IpMessagingV2ServiceChannelMember**](IpMessagingV2ServiceChannelMember.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMessage

> IpMessagingV2ServiceChannelMessage CreateMessage(ctx, ServiceSid, ChannelSid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).Body(Body).DateCreated(DateCreated).DateUpdated(DateUpdated).From(From).LastUpdatedBy(LastUpdatedBy).MediaSid(MediaSid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    ChannelSid := "ChannelSid_example" // string | 
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)
    Attributes := "Attributes_example" // string |  (optional)
    Body := "Body_example" // string |  (optional)
    DateCreated := time.Now() // time.Time |  (optional)
    DateUpdated := time.Now() // time.Time |  (optional)
    From := "From_example" // string |  (optional)
    LastUpdatedBy := "LastUpdatedBy_example" // string |  (optional)
    MediaSid := "MediaSid_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateMessage(context.Background(), ServiceSid, ChannelSid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).Body(Body).DateCreated(DateCreated).DateUpdated(DateUpdated).From(From).LastUpdatedBy(LastUpdatedBy).MediaSid(MediaSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMessage`: IpMessagingV2ServiceChannelMessage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateMessageParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
 **Attributes** | **string** | 
 **Body** | **string** | 
 **DateCreated** | **time.Time** | 
 **DateUpdated** | **time.Time** | 
 **From** | **string** | 
 **LastUpdatedBy** | **string** | 
 **MediaSid** | **string** | 

### Return type

[**IpMessagingV2ServiceChannelMessage**](IpMessagingV2ServiceChannelMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> IpMessagingV2ServiceRole CreateRole(ctx, ServiceSid).FriendlyName(FriendlyName).Permission(Permission).Type(Type).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    FriendlyName := "FriendlyName_example" // string |  (optional)
    Permission := []string{"Inner_example"} // []string |  (optional)
    Type := "Type_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateRole(context.Background(), ServiceSid).FriendlyName(FriendlyName).Permission(Permission).Type(Type).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRole`: IpMessagingV2ServiceRole
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateRoleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **FriendlyName** | **string** | 
 **Permission** | **[]string** | 
 **Type** | **string** | 

### Return type

[**IpMessagingV2ServiceRole**](IpMessagingV2ServiceRole.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateService

> IpMessagingV2Service CreateService(ctx).FriendlyName(FriendlyName).Execute()



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
    FriendlyName := "FriendlyName_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateService(context.Background()).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateService`: IpMessagingV2Service
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateService`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **FriendlyName** | **string** | 

### Return type

[**IpMessagingV2Service**](IpMessagingV2Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> IpMessagingV2ServiceUser CreateUser(ctx, ServiceSid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).FriendlyName(FriendlyName).Identity(Identity).RoleSid(RoleSid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)
    Attributes := "Attributes_example" // string |  (optional)
    FriendlyName := "FriendlyName_example" // string |  (optional)
    Identity := "Identity_example" // string |  (optional)
    RoleSid := "RoleSid_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateUser(context.Background(), ServiceSid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).FriendlyName(FriendlyName).Identity(Identity).RoleSid(RoleSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: IpMessagingV2ServiceUser
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateUserParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
 **Attributes** | **string** | 
 **FriendlyName** | **string** | 
 **Identity** | **string** | 
 **RoleSid** | **string** | 

### Return type

[**IpMessagingV2ServiceUser**](IpMessagingV2ServiceUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBinding

> DeleteBinding(ctx, ServiceSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    Sid := "Sid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteBinding(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteBindingParams struct via the builder pattern


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


## DeleteChannel

> DeleteChannel(ctx, ServiceSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    Sid := "Sid_example" // string | 
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteChannel(context.Background(), ServiceSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Execute()
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
**ServiceSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header

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


## DeleteChannelWebhook

> DeleteChannelWebhook(ctx, ServiceSid, ChannelSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    ChannelSid := "ChannelSid_example" // string | 
    Sid := "Sid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteChannelWebhook(context.Background(), ServiceSid, ChannelSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteChannelWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteChannelWebhookParams struct via the builder pattern


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


## DeleteCredential

> DeleteCredential(ctx, Sid).Execute()



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
    resp, r, err := api_client.DefaultApi.DeleteCredential(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteCredentialParams struct via the builder pattern


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


## DeleteInvite

> DeleteInvite(ctx, ServiceSid, ChannelSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    ChannelSid := "ChannelSid_example" // string | 
    Sid := "Sid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteInvite(context.Background(), ServiceSid, ChannelSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteInviteParams struct via the builder pattern


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


## DeleteMember

> DeleteMember(ctx, ServiceSid, ChannelSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    ChannelSid := "ChannelSid_example" // string | 
    Sid := "Sid_example" // string | 
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteMember(context.Background(), ServiceSid, ChannelSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteMemberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header

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


## DeleteMessage

> DeleteMessage(ctx, ServiceSid, ChannelSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    ChannelSid := "ChannelSid_example" // string | 
    Sid := "Sid_example" // string | 
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteMessage(context.Background(), ServiceSid, ChannelSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteMessageParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header

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


## DeleteRole

> DeleteRole(ctx, ServiceSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    Sid := "Sid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteRole(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteRoleParams struct via the builder pattern


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
    Sid := "Sid_example" // string | 

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
**Sid** | **string** | 

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


## DeleteUser

> DeleteUser(ctx, ServiceSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    Sid := "Sid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteUser(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteUserParams struct via the builder pattern


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


## DeleteUserBinding

> DeleteUserBinding(ctx, ServiceSid, UserSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    UserSid := "UserSid_example" // string | 
    Sid := "Sid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteUserBinding(context.Background(), ServiceSid, UserSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteUserBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**UserSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteUserBindingParams struct via the builder pattern


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


## DeleteUserChannel

> DeleteUserChannel(ctx, ServiceSid, UserSid, ChannelSid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    UserSid := "UserSid_example" // string | 
    ChannelSid := "ChannelSid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteUserChannel(context.Background(), ServiceSid, UserSid, ChannelSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteUserChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**UserSid** | **string** | 
**ChannelSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteUserChannelParams struct via the builder pattern


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


## FetchBinding

> IpMessagingV2ServiceBinding FetchBinding(ctx, ServiceSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    Sid := "Sid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchBinding(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchBinding`: IpMessagingV2ServiceBinding
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchBindingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**IpMessagingV2ServiceBinding**](IpMessagingV2ServiceBinding.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchChannel

> IpMessagingV2ServiceChannel FetchChannel(ctx, ServiceSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    Sid := "Sid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchChannel(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchChannel`: IpMessagingV2ServiceChannel
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**IpMessagingV2ServiceChannel**](IpMessagingV2ServiceChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchChannelWebhook

> IpMessagingV2ServiceChannelChannelWebhook FetchChannelWebhook(ctx, ServiceSid, ChannelSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    ChannelSid := "ChannelSid_example" // string | 
    Sid := "Sid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchChannelWebhook(context.Background(), ServiceSid, ChannelSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchChannelWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchChannelWebhook`: IpMessagingV2ServiceChannelChannelWebhook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchChannelWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchChannelWebhookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**IpMessagingV2ServiceChannelChannelWebhook**](IpMessagingV2ServiceChannelChannelWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCredential

> IpMessagingV2Credential FetchCredential(ctx, Sid).Execute()



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
    resp, r, err := api_client.DefaultApi.FetchCredential(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCredential`: IpMessagingV2Credential
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchCredentialParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**IpMessagingV2Credential**](IpMessagingV2Credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchInvite

> IpMessagingV2ServiceChannelInvite FetchInvite(ctx, ServiceSid, ChannelSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    ChannelSid := "ChannelSid_example" // string | 
    Sid := "Sid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchInvite(context.Background(), ServiceSid, ChannelSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchInvite`: IpMessagingV2ServiceChannelInvite
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchInviteParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**IpMessagingV2ServiceChannelInvite**](IpMessagingV2ServiceChannelInvite.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMember

> IpMessagingV2ServiceChannelMember FetchMember(ctx, ServiceSid, ChannelSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    ChannelSid := "ChannelSid_example" // string | 
    Sid := "Sid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchMember(context.Background(), ServiceSid, ChannelSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchMember`: IpMessagingV2ServiceChannelMember
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchMemberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**IpMessagingV2ServiceChannelMember**](IpMessagingV2ServiceChannelMember.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMessage

> IpMessagingV2ServiceChannelMessage FetchMessage(ctx, ServiceSid, ChannelSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    ChannelSid := "ChannelSid_example" // string | 
    Sid := "Sid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchMessage(context.Background(), ServiceSid, ChannelSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchMessage`: IpMessagingV2ServiceChannelMessage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchMessageParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**IpMessagingV2ServiceChannelMessage**](IpMessagingV2ServiceChannelMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRole

> IpMessagingV2ServiceRole FetchRole(ctx, ServiceSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    Sid := "Sid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchRole(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRole`: IpMessagingV2ServiceRole
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchRoleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**IpMessagingV2ServiceRole**](IpMessagingV2ServiceRole.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> IpMessagingV2Service FetchService(ctx, Sid).Execute()



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
    resp, r, err := api_client.DefaultApi.FetchService(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchService`: IpMessagingV2Service
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**IpMessagingV2Service**](IpMessagingV2Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUser

> IpMessagingV2ServiceUser FetchUser(ctx, ServiceSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    Sid := "Sid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchUser(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchUser`: IpMessagingV2ServiceUser
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchUserParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**IpMessagingV2ServiceUser**](IpMessagingV2ServiceUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUserBinding

> IpMessagingV2ServiceUserUserBinding FetchUserBinding(ctx, ServiceSid, UserSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    UserSid := "UserSid_example" // string | 
    Sid := "Sid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchUserBinding(context.Background(), ServiceSid, UserSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchUserBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchUserBinding`: IpMessagingV2ServiceUserUserBinding
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchUserBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**UserSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchUserBindingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**IpMessagingV2ServiceUserUserBinding**](IpMessagingV2ServiceUserUserBinding.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUserChannel

> IpMessagingV2ServiceUserUserChannel FetchUserChannel(ctx, ServiceSid, UserSid, ChannelSid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    UserSid := "UserSid_example" // string | 
    ChannelSid := "ChannelSid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchUserChannel(context.Background(), ServiceSid, UserSid, ChannelSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchUserChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchUserChannel`: IpMessagingV2ServiceUserUserChannel
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchUserChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**UserSid** | **string** | 
**ChannelSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchUserChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**IpMessagingV2ServiceUserUserChannel**](IpMessagingV2ServiceUserUserChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBinding

> ListBindingResponse ListBinding(ctx, ServiceSid).BindingType(BindingType).Identity(Identity).PageSize(PageSize).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    BindingType := []string{"BindingType_example"} // []string |  (optional)
    Identity := []string{"Inner_example"} // []string |  (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListBinding(context.Background(), ServiceSid).BindingType(BindingType).Identity(Identity).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBinding`: ListBindingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListBindingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **BindingType** | **[]string** | 
 **Identity** | **[]string** | 
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListBindingResponse**](ListBindingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChannel

> ListChannelResponse ListChannel(ctx, ServiceSid).Type(Type).PageSize(PageSize).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    Type := []string{"Type_example"} // []string |  (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListChannel(context.Background(), ServiceSid).Type(Type).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListChannel`: ListChannelResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Type** | **[]string** | 
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


## ListChannelWebhook

> ListChannelWebhookResponse ListChannelWebhook(ctx, ServiceSid, ChannelSid).PageSize(PageSize).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    ChannelSid := "ChannelSid_example" // string | 
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListChannelWebhook(context.Background(), ServiceSid, ChannelSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListChannelWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListChannelWebhook`: ListChannelWebhookResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListChannelWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListChannelWebhookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListChannelWebhookResponse**](ListChannelWebhookResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCredential

> ListCredentialResponse ListCredential(ctx).PageSize(PageSize).Execute()



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
    resp, r, err := api_client.DefaultApi.ListCredential(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCredential`: ListCredentialResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCredential`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListCredentialParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListCredentialResponse**](ListCredentialResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInvite

> ListInviteResponse ListInvite(ctx, ServiceSid, ChannelSid).Identity(Identity).PageSize(PageSize).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    ChannelSid := "ChannelSid_example" // string | 
    Identity := []string{"Inner_example"} // []string |  (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListInvite(context.Background(), ServiceSid, ChannelSid).Identity(Identity).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInvite`: ListInviteResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListInviteParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Identity** | **[]string** | 
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListInviteResponse**](ListInviteResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMember

> ListMemberResponse ListMember(ctx, ServiceSid, ChannelSid).Identity(Identity).PageSize(PageSize).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    ChannelSid := "ChannelSid_example" // string | 
    Identity := []string{"Inner_example"} // []string |  (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListMember(context.Background(), ServiceSid, ChannelSid).Identity(Identity).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMember`: ListMemberResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListMemberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Identity** | **[]string** | 
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListMemberResponse**](ListMemberResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessage

> ListMessageResponse ListMessage(ctx, ServiceSid, ChannelSid).Order(Order).PageSize(PageSize).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    ChannelSid := "ChannelSid_example" // string | 
    Order := "Order_example" // string |  (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListMessage(context.Background(), ServiceSid, ChannelSid).Order(Order).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMessage`: ListMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListMessageParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Order** | **string** | 
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListMessageResponse**](ListMessageResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRole

> ListRoleResponse ListRole(ctx, ServiceSid).PageSize(PageSize).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListRole(context.Background(), ServiceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRole`: ListRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListRoleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListRoleResponse**](ListRoleResponse.md)

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


## ListUser

> ListUserResponse ListUser(ctx, ServiceSid).PageSize(PageSize).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListUser(context.Background(), ServiceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUser`: ListUserResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListUserParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListUserResponse**](ListUserResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserBinding

> ListUserBindingResponse ListUserBinding(ctx, ServiceSid, UserSid).BindingType(BindingType).PageSize(PageSize).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    UserSid := "UserSid_example" // string | 
    BindingType := []string{"BindingType_example"} // []string |  (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListUserBinding(context.Background(), ServiceSid, UserSid).BindingType(BindingType).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUserBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserBinding`: ListUserBindingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUserBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**UserSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListUserBindingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **BindingType** | **[]string** | 
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListUserBindingResponse**](ListUserBindingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserChannel

> ListUserChannelResponse ListUserChannel(ctx, ServiceSid, UserSid).PageSize(PageSize).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    UserSid := "UserSid_example" // string | 
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListUserChannel(context.Background(), ServiceSid, UserSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUserChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserChannel`: ListUserChannelResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUserChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**UserSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListUserChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListUserChannelResponse**](ListUserChannelResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChannel

> IpMessagingV2ServiceChannel UpdateChannel(ctx, ServiceSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).CreatedBy(CreatedBy).DateCreated(DateCreated).DateUpdated(DateUpdated).FriendlyName(FriendlyName).UniqueName(UniqueName).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    Sid := "Sid_example" // string | 
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)
    Attributes := "Attributes_example" // string |  (optional)
    CreatedBy := "CreatedBy_example" // string |  (optional)
    DateCreated := time.Now() // time.Time |  (optional)
    DateUpdated := time.Now() // time.Time |  (optional)
    FriendlyName := "FriendlyName_example" // string |  (optional)
    UniqueName := "UniqueName_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateChannel(context.Background(), ServiceSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).CreatedBy(CreatedBy).DateCreated(DateCreated).DateUpdated(DateUpdated).FriendlyName(FriendlyName).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateChannel`: IpMessagingV2ServiceChannel
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
 **Attributes** | **string** | 
 **CreatedBy** | **string** | 
 **DateCreated** | **time.Time** | 
 **DateUpdated** | **time.Time** | 
 **FriendlyName** | **string** | 
 **UniqueName** | **string** | 

### Return type

[**IpMessagingV2ServiceChannel**](IpMessagingV2ServiceChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChannelWebhook

> IpMessagingV2ServiceChannelChannelWebhook UpdateChannelWebhook(ctx, ServiceSid, ChannelSid, Sid).ConfigurationFilters(ConfigurationFilters).ConfigurationFlowSid(ConfigurationFlowSid).ConfigurationMethod(ConfigurationMethod).ConfigurationRetryCount(ConfigurationRetryCount).ConfigurationTriggers(ConfigurationTriggers).ConfigurationUrl(ConfigurationUrl).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    ChannelSid := "ChannelSid_example" // string | 
    Sid := "Sid_example" // string | 
    ConfigurationFilters := []string{"Inner_example"} // []string |  (optional)
    ConfigurationFlowSid := "ConfigurationFlowSid_example" // string |  (optional)
    ConfigurationMethod := "ConfigurationMethod_example" // string |  (optional)
    ConfigurationRetryCount := int32(56) // int32 |  (optional)
    ConfigurationTriggers := []string{"Inner_example"} // []string |  (optional)
    ConfigurationUrl := "ConfigurationUrl_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateChannelWebhook(context.Background(), ServiceSid, ChannelSid, Sid).ConfigurationFilters(ConfigurationFilters).ConfigurationFlowSid(ConfigurationFlowSid).ConfigurationMethod(ConfigurationMethod).ConfigurationRetryCount(ConfigurationRetryCount).ConfigurationTriggers(ConfigurationTriggers).ConfigurationUrl(ConfigurationUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateChannelWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateChannelWebhook`: IpMessagingV2ServiceChannelChannelWebhook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateChannelWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateChannelWebhookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **ConfigurationFilters** | **[]string** | 
 **ConfigurationFlowSid** | **string** | 
 **ConfigurationMethod** | **string** | 
 **ConfigurationRetryCount** | **int32** | 
 **ConfigurationTriggers** | **[]string** | 
 **ConfigurationUrl** | **string** | 

### Return type

[**IpMessagingV2ServiceChannelChannelWebhook**](IpMessagingV2ServiceChannelChannelWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredential

> IpMessagingV2Credential UpdateCredential(ctx, Sid).ApiKey(ApiKey).Certificate(Certificate).FriendlyName(FriendlyName).PrivateKey(PrivateKey).Sandbox(Sandbox).Secret(Secret).Execute()



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
    ApiKey := "ApiKey_example" // string |  (optional)
    Certificate := "Certificate_example" // string |  (optional)
    FriendlyName := "FriendlyName_example" // string |  (optional)
    PrivateKey := "PrivateKey_example" // string |  (optional)
    Sandbox := true // bool |  (optional)
    Secret := "Secret_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateCredential(context.Background(), Sid).ApiKey(ApiKey).Certificate(Certificate).FriendlyName(FriendlyName).PrivateKey(PrivateKey).Sandbox(Sandbox).Secret(Secret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCredential`: IpMessagingV2Credential
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateCredentialParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **ApiKey** | **string** | 
 **Certificate** | **string** | 
 **FriendlyName** | **string** | 
 **PrivateKey** | **string** | 
 **Sandbox** | **bool** | 
 **Secret** | **string** | 

### Return type

[**IpMessagingV2Credential**](IpMessagingV2Credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMember

> IpMessagingV2ServiceChannelMember UpdateMember(ctx, ServiceSid, ChannelSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).DateCreated(DateCreated).DateUpdated(DateUpdated).LastConsumedMessageIndex(LastConsumedMessageIndex).LastConsumptionTimestamp(LastConsumptionTimestamp).RoleSid(RoleSid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    ChannelSid := "ChannelSid_example" // string | 
    Sid := "Sid_example" // string | 
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)
    Attributes := "Attributes_example" // string |  (optional)
    DateCreated := time.Now() // time.Time |  (optional)
    DateUpdated := time.Now() // time.Time |  (optional)
    LastConsumedMessageIndex := int32(56) // int32 |  (optional)
    LastConsumptionTimestamp := time.Now() // time.Time |  (optional)
    RoleSid := "RoleSid_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateMember(context.Background(), ServiceSid, ChannelSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).DateCreated(DateCreated).DateUpdated(DateUpdated).LastConsumedMessageIndex(LastConsumedMessageIndex).LastConsumptionTimestamp(LastConsumptionTimestamp).RoleSid(RoleSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMember`: IpMessagingV2ServiceChannelMember
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateMemberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
 **Attributes** | **string** | 
 **DateCreated** | **time.Time** | 
 **DateUpdated** | **time.Time** | 
 **LastConsumedMessageIndex** | **int32** | 
 **LastConsumptionTimestamp** | **time.Time** | 
 **RoleSid** | **string** | 

### Return type

[**IpMessagingV2ServiceChannelMember**](IpMessagingV2ServiceChannelMember.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessage

> IpMessagingV2ServiceChannelMessage UpdateMessage(ctx, ServiceSid, ChannelSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).Body(Body).DateCreated(DateCreated).DateUpdated(DateUpdated).From(From).LastUpdatedBy(LastUpdatedBy).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    ChannelSid := "ChannelSid_example" // string | 
    Sid := "Sid_example" // string | 
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)
    Attributes := "Attributes_example" // string |  (optional)
    Body := "Body_example" // string |  (optional)
    DateCreated := time.Now() // time.Time |  (optional)
    DateUpdated := time.Now() // time.Time |  (optional)
    From := "From_example" // string |  (optional)
    LastUpdatedBy := "LastUpdatedBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateMessage(context.Background(), ServiceSid, ChannelSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).Body(Body).DateCreated(DateCreated).DateUpdated(DateUpdated).From(From).LastUpdatedBy(LastUpdatedBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMessage`: IpMessagingV2ServiceChannelMessage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**ChannelSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateMessageParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
 **Attributes** | **string** | 
 **Body** | **string** | 
 **DateCreated** | **time.Time** | 
 **DateUpdated** | **time.Time** | 
 **From** | **string** | 
 **LastUpdatedBy** | **string** | 

### Return type

[**IpMessagingV2ServiceChannelMessage**](IpMessagingV2ServiceChannelMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> IpMessagingV2ServiceRole UpdateRole(ctx, ServiceSid, Sid).Permission(Permission).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    Sid := "Sid_example" // string | 
    Permission := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateRole(context.Background(), ServiceSid, Sid).Permission(Permission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRole`: IpMessagingV2ServiceRole
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateRoleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Permission** | **[]string** | 

### Return type

[**IpMessagingV2ServiceRole**](IpMessagingV2ServiceRole.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> IpMessagingV2Service UpdateService(ctx, Sid).ConsumptionReportInterval(ConsumptionReportInterval).DefaultChannelCreatorRoleSid(DefaultChannelCreatorRoleSid).DefaultChannelRoleSid(DefaultChannelRoleSid).DefaultServiceRoleSid(DefaultServiceRoleSid).FriendlyName(FriendlyName).LimitsChannelMembers(LimitsChannelMembers).LimitsUserChannels(LimitsUserChannels).MediaCompatibilityMessage(MediaCompatibilityMessage).NotificationsAddedToChannelEnabled(NotificationsAddedToChannelEnabled).NotificationsAddedToChannelSound(NotificationsAddedToChannelSound).NotificationsAddedToChannelTemplate(NotificationsAddedToChannelTemplate).NotificationsInvitedToChannelEnabled(NotificationsInvitedToChannelEnabled).NotificationsInvitedToChannelSound(NotificationsInvitedToChannelSound).NotificationsInvitedToChannelTemplate(NotificationsInvitedToChannelTemplate).NotificationsLogEnabled(NotificationsLogEnabled).NotificationsNewMessageBadgeCountEnabled(NotificationsNewMessageBadgeCountEnabled).NotificationsNewMessageEnabled(NotificationsNewMessageEnabled).NotificationsNewMessageSound(NotificationsNewMessageSound).NotificationsNewMessageTemplate(NotificationsNewMessageTemplate).NotificationsRemovedFromChannelEnabled(NotificationsRemovedFromChannelEnabled).NotificationsRemovedFromChannelSound(NotificationsRemovedFromChannelSound).NotificationsRemovedFromChannelTemplate(NotificationsRemovedFromChannelTemplate).PostWebhookRetryCount(PostWebhookRetryCount).PostWebhookUrl(PostWebhookUrl).PreWebhookRetryCount(PreWebhookRetryCount).PreWebhookUrl(PreWebhookUrl).ReachabilityEnabled(ReachabilityEnabled).ReadStatusEnabled(ReadStatusEnabled).TypingIndicatorTimeout(TypingIndicatorTimeout).WebhookFilters(WebhookFilters).WebhookMethod(WebhookMethod).Execute()



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
    ConsumptionReportInterval := int32(56) // int32 |  (optional)
    DefaultChannelCreatorRoleSid := "DefaultChannelCreatorRoleSid_example" // string |  (optional)
    DefaultChannelRoleSid := "DefaultChannelRoleSid_example" // string |  (optional)
    DefaultServiceRoleSid := "DefaultServiceRoleSid_example" // string |  (optional)
    FriendlyName := "FriendlyName_example" // string |  (optional)
    LimitsChannelMembers := int32(56) // int32 |  (optional)
    LimitsUserChannels := int32(56) // int32 |  (optional)
    MediaCompatibilityMessage := "MediaCompatibilityMessage_example" // string |  (optional)
    NotificationsAddedToChannelEnabled := true // bool |  (optional)
    NotificationsAddedToChannelSound := "NotificationsAddedToChannelSound_example" // string |  (optional)
    NotificationsAddedToChannelTemplate := "NotificationsAddedToChannelTemplate_example" // string |  (optional)
    NotificationsInvitedToChannelEnabled := true // bool |  (optional)
    NotificationsInvitedToChannelSound := "NotificationsInvitedToChannelSound_example" // string |  (optional)
    NotificationsInvitedToChannelTemplate := "NotificationsInvitedToChannelTemplate_example" // string |  (optional)
    NotificationsLogEnabled := true // bool |  (optional)
    NotificationsNewMessageBadgeCountEnabled := true // bool |  (optional)
    NotificationsNewMessageEnabled := true // bool |  (optional)
    NotificationsNewMessageSound := "NotificationsNewMessageSound_example" // string |  (optional)
    NotificationsNewMessageTemplate := "NotificationsNewMessageTemplate_example" // string |  (optional)
    NotificationsRemovedFromChannelEnabled := true // bool |  (optional)
    NotificationsRemovedFromChannelSound := "NotificationsRemovedFromChannelSound_example" // string |  (optional)
    NotificationsRemovedFromChannelTemplate := "NotificationsRemovedFromChannelTemplate_example" // string |  (optional)
    PostWebhookRetryCount := int32(56) // int32 |  (optional)
    PostWebhookUrl := "PostWebhookUrl_example" // string |  (optional)
    PreWebhookRetryCount := int32(56) // int32 |  (optional)
    PreWebhookUrl := "PreWebhookUrl_example" // string |  (optional)
    ReachabilityEnabled := true // bool |  (optional)
    ReadStatusEnabled := true // bool |  (optional)
    TypingIndicatorTimeout := int32(56) // int32 |  (optional)
    WebhookFilters := []string{"Inner_example"} // []string |  (optional)
    WebhookMethod := "WebhookMethod_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateService(context.Background(), Sid).ConsumptionReportInterval(ConsumptionReportInterval).DefaultChannelCreatorRoleSid(DefaultChannelCreatorRoleSid).DefaultChannelRoleSid(DefaultChannelRoleSid).DefaultServiceRoleSid(DefaultServiceRoleSid).FriendlyName(FriendlyName).LimitsChannelMembers(LimitsChannelMembers).LimitsUserChannels(LimitsUserChannels).MediaCompatibilityMessage(MediaCompatibilityMessage).NotificationsAddedToChannelEnabled(NotificationsAddedToChannelEnabled).NotificationsAddedToChannelSound(NotificationsAddedToChannelSound).NotificationsAddedToChannelTemplate(NotificationsAddedToChannelTemplate).NotificationsInvitedToChannelEnabled(NotificationsInvitedToChannelEnabled).NotificationsInvitedToChannelSound(NotificationsInvitedToChannelSound).NotificationsInvitedToChannelTemplate(NotificationsInvitedToChannelTemplate).NotificationsLogEnabled(NotificationsLogEnabled).NotificationsNewMessageBadgeCountEnabled(NotificationsNewMessageBadgeCountEnabled).NotificationsNewMessageEnabled(NotificationsNewMessageEnabled).NotificationsNewMessageSound(NotificationsNewMessageSound).NotificationsNewMessageTemplate(NotificationsNewMessageTemplate).NotificationsRemovedFromChannelEnabled(NotificationsRemovedFromChannelEnabled).NotificationsRemovedFromChannelSound(NotificationsRemovedFromChannelSound).NotificationsRemovedFromChannelTemplate(NotificationsRemovedFromChannelTemplate).PostWebhookRetryCount(PostWebhookRetryCount).PostWebhookUrl(PostWebhookUrl).PreWebhookRetryCount(PreWebhookRetryCount).PreWebhookUrl(PreWebhookUrl).ReachabilityEnabled(ReachabilityEnabled).ReadStatusEnabled(ReadStatusEnabled).TypingIndicatorTimeout(TypingIndicatorTimeout).WebhookFilters(WebhookFilters).WebhookMethod(WebhookMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateService`: IpMessagingV2Service
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **ConsumptionReportInterval** | **int32** | 
 **DefaultChannelCreatorRoleSid** | **string** | 
 **DefaultChannelRoleSid** | **string** | 
 **DefaultServiceRoleSid** | **string** | 
 **FriendlyName** | **string** | 
 **LimitsChannelMembers** | **int32** | 
 **LimitsUserChannels** | **int32** | 
 **MediaCompatibilityMessage** | **string** | 
 **NotificationsAddedToChannelEnabled** | **bool** | 
 **NotificationsAddedToChannelSound** | **string** | 
 **NotificationsAddedToChannelTemplate** | **string** | 
 **NotificationsInvitedToChannelEnabled** | **bool** | 
 **NotificationsInvitedToChannelSound** | **string** | 
 **NotificationsInvitedToChannelTemplate** | **string** | 
 **NotificationsLogEnabled** | **bool** | 
 **NotificationsNewMessageBadgeCountEnabled** | **bool** | 
 **NotificationsNewMessageEnabled** | **bool** | 
 **NotificationsNewMessageSound** | **string** | 
 **NotificationsNewMessageTemplate** | **string** | 
 **NotificationsRemovedFromChannelEnabled** | **bool** | 
 **NotificationsRemovedFromChannelSound** | **string** | 
 **NotificationsRemovedFromChannelTemplate** | **string** | 
 **PostWebhookRetryCount** | **int32** | 
 **PostWebhookUrl** | **string** | 
 **PreWebhookRetryCount** | **int32** | 
 **PreWebhookUrl** | **string** | 
 **ReachabilityEnabled** | **bool** | 
 **ReadStatusEnabled** | **bool** | 
 **TypingIndicatorTimeout** | **int32** | 
 **WebhookFilters** | **[]string** | 
 **WebhookMethod** | **string** | 

### Return type

[**IpMessagingV2Service**](IpMessagingV2Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> IpMessagingV2ServiceUser UpdateUser(ctx, ServiceSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).FriendlyName(FriendlyName).RoleSid(RoleSid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    Sid := "Sid_example" // string | 
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)
    Attributes := "Attributes_example" // string |  (optional)
    FriendlyName := "FriendlyName_example" // string |  (optional)
    RoleSid := "RoleSid_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateUser(context.Background(), ServiceSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).FriendlyName(FriendlyName).RoleSid(RoleSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: IpMessagingV2ServiceUser
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateUserParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
 **Attributes** | **string** | 
 **FriendlyName** | **string** | 
 **RoleSid** | **string** | 

### Return type

[**IpMessagingV2ServiceUser**](IpMessagingV2ServiceUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserChannel

> IpMessagingV2ServiceUserUserChannel UpdateUserChannel(ctx, ServiceSid, UserSid, ChannelSid).LastConsumedMessageIndex(LastConsumedMessageIndex).LastConsumptionTimestamp(LastConsumptionTimestamp).NotificationLevel(NotificationLevel).Execute()



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
    ServiceSid := "ServiceSid_example" // string | 
    UserSid := "UserSid_example" // string | 
    ChannelSid := "ChannelSid_example" // string | 
    LastConsumedMessageIndex := int32(56) // int32 |  (optional)
    LastConsumptionTimestamp := time.Now() // time.Time |  (optional)
    NotificationLevel := "NotificationLevel_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateUserChannel(context.Background(), ServiceSid, UserSid, ChannelSid).LastConsumedMessageIndex(LastConsumedMessageIndex).LastConsumptionTimestamp(LastConsumptionTimestamp).NotificationLevel(NotificationLevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateUserChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserChannel`: IpMessagingV2ServiceUserUserChannel
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateUserChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**UserSid** | **string** | 
**ChannelSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateUserChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **LastConsumedMessageIndex** | **int32** | 
 **LastConsumptionTimestamp** | **time.Time** | 
 **NotificationLevel** | **string** | 

### Return type

[**IpMessagingV2ServiceUserUserChannel**](IpMessagingV2ServiceUserUserChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

