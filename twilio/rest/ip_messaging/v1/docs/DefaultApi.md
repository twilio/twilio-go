# \DefaultApi

All URIs are relative to *https://ip-messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChannel**](DefaultApi.md#CreateChannel) | **Post** /v1/Services/{ServiceSid}/Channels | 
[**CreateCredential**](DefaultApi.md#CreateCredential) | **Post** /v1/Credentials | 
[**CreateInvite**](DefaultApi.md#CreateInvite) | **Post** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Invites | 
[**CreateMember**](DefaultApi.md#CreateMember) | **Post** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Members | 
[**CreateMessage**](DefaultApi.md#CreateMessage) | **Post** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Messages | 
[**CreateRole**](DefaultApi.md#CreateRole) | **Post** /v1/Services/{ServiceSid}/Roles | 
[**CreateService**](DefaultApi.md#CreateService) | **Post** /v1/Services | 
[**CreateUser**](DefaultApi.md#CreateUser) | **Post** /v1/Services/{ServiceSid}/Users | 
[**DeleteChannel**](DefaultApi.md#DeleteChannel) | **Delete** /v1/Services/{ServiceSid}/Channels/{Sid} | 
[**DeleteCredential**](DefaultApi.md#DeleteCredential) | **Delete** /v1/Credentials/{Sid} | 
[**DeleteInvite**](DefaultApi.md#DeleteInvite) | **Delete** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Invites/{Sid} | 
[**DeleteMember**](DefaultApi.md#DeleteMember) | **Delete** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Members/{Sid} | 
[**DeleteMessage**](DefaultApi.md#DeleteMessage) | **Delete** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Messages/{Sid} | 
[**DeleteRole**](DefaultApi.md#DeleteRole) | **Delete** /v1/Services/{ServiceSid}/Roles/{Sid} | 
[**DeleteService**](DefaultApi.md#DeleteService) | **Delete** /v1/Services/{Sid} | 
[**DeleteUser**](DefaultApi.md#DeleteUser) | **Delete** /v1/Services/{ServiceSid}/Users/{Sid} | 
[**FetchChannel**](DefaultApi.md#FetchChannel) | **Get** /v1/Services/{ServiceSid}/Channels/{Sid} | 
[**FetchCredential**](DefaultApi.md#FetchCredential) | **Get** /v1/Credentials/{Sid} | 
[**FetchInvite**](DefaultApi.md#FetchInvite) | **Get** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Invites/{Sid} | 
[**FetchMember**](DefaultApi.md#FetchMember) | **Get** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Members/{Sid} | 
[**FetchMessage**](DefaultApi.md#FetchMessage) | **Get** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Messages/{Sid} | 
[**FetchRole**](DefaultApi.md#FetchRole) | **Get** /v1/Services/{ServiceSid}/Roles/{Sid} | 
[**FetchService**](DefaultApi.md#FetchService) | **Get** /v1/Services/{Sid} | 
[**FetchUser**](DefaultApi.md#FetchUser) | **Get** /v1/Services/{ServiceSid}/Users/{Sid} | 
[**ListChannel**](DefaultApi.md#ListChannel) | **Get** /v1/Services/{ServiceSid}/Channels | 
[**ListCredential**](DefaultApi.md#ListCredential) | **Get** /v1/Credentials | 
[**ListInvite**](DefaultApi.md#ListInvite) | **Get** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Invites | 
[**ListMember**](DefaultApi.md#ListMember) | **Get** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Members | 
[**ListMessage**](DefaultApi.md#ListMessage) | **Get** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Messages | 
[**ListRole**](DefaultApi.md#ListRole) | **Get** /v1/Services/{ServiceSid}/Roles | 
[**ListService**](DefaultApi.md#ListService) | **Get** /v1/Services | 
[**ListUser**](DefaultApi.md#ListUser) | **Get** /v1/Services/{ServiceSid}/Users | 
[**ListUserChannel**](DefaultApi.md#ListUserChannel) | **Get** /v1/Services/{ServiceSid}/Users/{UserSid}/Channels | 
[**UpdateChannel**](DefaultApi.md#UpdateChannel) | **Post** /v1/Services/{ServiceSid}/Channels/{Sid} | 
[**UpdateCredential**](DefaultApi.md#UpdateCredential) | **Post** /v1/Credentials/{Sid} | 
[**UpdateMember**](DefaultApi.md#UpdateMember) | **Post** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Members/{Sid} | 
[**UpdateMessage**](DefaultApi.md#UpdateMessage) | **Post** /v1/Services/{ServiceSid}/Channels/{ChannelSid}/Messages/{Sid} | 
[**UpdateRole**](DefaultApi.md#UpdateRole) | **Post** /v1/Services/{ServiceSid}/Roles/{Sid} | 
[**UpdateService**](DefaultApi.md#UpdateService) | **Post** /v1/Services/{Sid} | 
[**UpdateUser**](DefaultApi.md#UpdateUser) | **Post** /v1/Services/{ServiceSid}/Users/{Sid} | 



## CreateChannel

> IpMessagingV1ServiceChannel CreateChannel(ctx, ServiceSid).Attributes(Attributes).FriendlyName(FriendlyName).Type(Type).UniqueName(UniqueName).Execute()



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
    Attributes := "Attributes_example" // string |  (optional)
    FriendlyName := "FriendlyName_example" // string |  (optional)
    Type := "Type_example" // string |  (optional)
    UniqueName := "UniqueName_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateChannel(context.Background(), ServiceSid).Attributes(Attributes).FriendlyName(FriendlyName).Type(Type).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateChannel`: IpMessagingV1ServiceChannel
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

 **Attributes** | **string** | 
 **FriendlyName** | **string** | 
 **Type** | **string** | 
 **UniqueName** | **string** | 

### Return type

[**IpMessagingV1ServiceChannel**](IpMessagingV1ServiceChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCredential

> IpMessagingV1Credential CreateCredential(ctx).ApiKey(ApiKey).Certificate(Certificate).FriendlyName(FriendlyName).PrivateKey(PrivateKey).Sandbox(Sandbox).Secret(Secret).Type(Type).Execute()



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
    // response from `CreateCredential`: IpMessagingV1Credential
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

[**IpMessagingV1Credential**](IpMessagingV1Credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvite

> IpMessagingV1ServiceChannelInvite CreateInvite(ctx, ServiceSid, ChannelSid).Identity(Identity).RoleSid(RoleSid).Execute()



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
    // response from `CreateInvite`: IpMessagingV1ServiceChannelInvite
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

[**IpMessagingV1ServiceChannelInvite**](IpMessagingV1ServiceChannelInvite.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMember

> IpMessagingV1ServiceChannelMember CreateMember(ctx, ServiceSid, ChannelSid).Identity(Identity).RoleSid(RoleSid).Execute()



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
    resp, r, err := api_client.DefaultApi.CreateMember(context.Background(), ServiceSid, ChannelSid).Identity(Identity).RoleSid(RoleSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMember`: IpMessagingV1ServiceChannelMember
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


 **Identity** | **string** | 
 **RoleSid** | **string** | 

### Return type

[**IpMessagingV1ServiceChannelMember**](IpMessagingV1ServiceChannelMember.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMessage

> IpMessagingV1ServiceChannelMessage CreateMessage(ctx, ServiceSid, ChannelSid).Attributes(Attributes).Body(Body).From(From).Execute()



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
    Attributes := "Attributes_example" // string |  (optional)
    Body := "Body_example" // string |  (optional)
    From := "From_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateMessage(context.Background(), ServiceSid, ChannelSid).Attributes(Attributes).Body(Body).From(From).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMessage`: IpMessagingV1ServiceChannelMessage
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


 **Attributes** | **string** | 
 **Body** | **string** | 
 **From** | **string** | 

### Return type

[**IpMessagingV1ServiceChannelMessage**](IpMessagingV1ServiceChannelMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> IpMessagingV1ServiceRole CreateRole(ctx, ServiceSid).FriendlyName(FriendlyName).Permission(Permission).Type(Type).Execute()



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
    // response from `CreateRole`: IpMessagingV1ServiceRole
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

[**IpMessagingV1ServiceRole**](IpMessagingV1ServiceRole.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateService

> IpMessagingV1Service CreateService(ctx).FriendlyName(FriendlyName).Execute()



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
    // response from `CreateService`: IpMessagingV1Service
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

[**IpMessagingV1Service**](IpMessagingV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> IpMessagingV1ServiceUser CreateUser(ctx, ServiceSid).Attributes(Attributes).FriendlyName(FriendlyName).Identity(Identity).RoleSid(RoleSid).Execute()



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
    Attributes := "Attributes_example" // string |  (optional)
    FriendlyName := "FriendlyName_example" // string |  (optional)
    Identity := "Identity_example" // string |  (optional)
    RoleSid := "RoleSid_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateUser(context.Background(), ServiceSid).Attributes(Attributes).FriendlyName(FriendlyName).Identity(Identity).RoleSid(RoleSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: IpMessagingV1ServiceUser
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

 **Attributes** | **string** | 
 **FriendlyName** | **string** | 
 **Identity** | **string** | 
 **RoleSid** | **string** | 

### Return type

[**IpMessagingV1ServiceUser**](IpMessagingV1ServiceUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChannel

> DeleteChannel(ctx, ServiceSid, Sid).Execute()



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
    resp, r, err := api_client.DefaultApi.DeleteChannel(context.Background(), ServiceSid, Sid).Execute()
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

> DeleteMember(ctx, ServiceSid, ChannelSid, Sid).Execute()



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
    resp, r, err := api_client.DefaultApi.DeleteMember(context.Background(), ServiceSid, ChannelSid, Sid).Execute()
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

> DeleteMessage(ctx, ServiceSid, ChannelSid, Sid).Execute()



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
    resp, r, err := api_client.DefaultApi.DeleteMessage(context.Background(), ServiceSid, ChannelSid, Sid).Execute()
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


## FetchChannel

> IpMessagingV1ServiceChannel FetchChannel(ctx, ServiceSid, Sid).Execute()



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
    // response from `FetchChannel`: IpMessagingV1ServiceChannel
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

[**IpMessagingV1ServiceChannel**](IpMessagingV1ServiceChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCredential

> IpMessagingV1Credential FetchCredential(ctx, Sid).Execute()



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
    // response from `FetchCredential`: IpMessagingV1Credential
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

[**IpMessagingV1Credential**](IpMessagingV1Credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchInvite

> IpMessagingV1ServiceChannelInvite FetchInvite(ctx, ServiceSid, ChannelSid, Sid).Execute()



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
    // response from `FetchInvite`: IpMessagingV1ServiceChannelInvite
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

[**IpMessagingV1ServiceChannelInvite**](IpMessagingV1ServiceChannelInvite.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMember

> IpMessagingV1ServiceChannelMember FetchMember(ctx, ServiceSid, ChannelSid, Sid).Execute()



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
    // response from `FetchMember`: IpMessagingV1ServiceChannelMember
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

[**IpMessagingV1ServiceChannelMember**](IpMessagingV1ServiceChannelMember.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMessage

> IpMessagingV1ServiceChannelMessage FetchMessage(ctx, ServiceSid, ChannelSid, Sid).Execute()



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
    // response from `FetchMessage`: IpMessagingV1ServiceChannelMessage
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

[**IpMessagingV1ServiceChannelMessage**](IpMessagingV1ServiceChannelMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRole

> IpMessagingV1ServiceRole FetchRole(ctx, ServiceSid, Sid).Execute()



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
    // response from `FetchRole`: IpMessagingV1ServiceRole
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

[**IpMessagingV1ServiceRole**](IpMessagingV1ServiceRole.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> IpMessagingV1Service FetchService(ctx, Sid).Execute()



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
    // response from `FetchService`: IpMessagingV1Service
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

[**IpMessagingV1Service**](IpMessagingV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUser

> IpMessagingV1ServiceUser FetchUser(ctx, ServiceSid, Sid).Execute()



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
    // response from `FetchUser`: IpMessagingV1ServiceUser
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

[**IpMessagingV1ServiceUser**](IpMessagingV1ServiceUser.md)

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

> IpMessagingV1ServiceChannel UpdateChannel(ctx, ServiceSid, Sid).Attributes(Attributes).FriendlyName(FriendlyName).UniqueName(UniqueName).Execute()



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
    Attributes := "Attributes_example" // string |  (optional)
    FriendlyName := "FriendlyName_example" // string |  (optional)
    UniqueName := "UniqueName_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateChannel(context.Background(), ServiceSid, Sid).Attributes(Attributes).FriendlyName(FriendlyName).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateChannel`: IpMessagingV1ServiceChannel
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


 **Attributes** | **string** | 
 **FriendlyName** | **string** | 
 **UniqueName** | **string** | 

### Return type

[**IpMessagingV1ServiceChannel**](IpMessagingV1ServiceChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredential

> IpMessagingV1Credential UpdateCredential(ctx, Sid).ApiKey(ApiKey).Certificate(Certificate).FriendlyName(FriendlyName).PrivateKey(PrivateKey).Sandbox(Sandbox).Secret(Secret).Execute()



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
    // response from `UpdateCredential`: IpMessagingV1Credential
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

[**IpMessagingV1Credential**](IpMessagingV1Credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMember

> IpMessagingV1ServiceChannelMember UpdateMember(ctx, ServiceSid, ChannelSid, Sid).LastConsumedMessageIndex(LastConsumedMessageIndex).RoleSid(RoleSid).Execute()



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
    LastConsumedMessageIndex := int32(56) // int32 |  (optional)
    RoleSid := "RoleSid_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateMember(context.Background(), ServiceSid, ChannelSid, Sid).LastConsumedMessageIndex(LastConsumedMessageIndex).RoleSid(RoleSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMember`: IpMessagingV1ServiceChannelMember
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



 **LastConsumedMessageIndex** | **int32** | 
 **RoleSid** | **string** | 

### Return type

[**IpMessagingV1ServiceChannelMember**](IpMessagingV1ServiceChannelMember.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessage

> IpMessagingV1ServiceChannelMessage UpdateMessage(ctx, ServiceSid, ChannelSid, Sid).Attributes(Attributes).Body(Body).Execute()



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
    Attributes := "Attributes_example" // string |  (optional)
    Body := "Body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateMessage(context.Background(), ServiceSid, ChannelSid, Sid).Attributes(Attributes).Body(Body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMessage`: IpMessagingV1ServiceChannelMessage
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



 **Attributes** | **string** | 
 **Body** | **string** | 

### Return type

[**IpMessagingV1ServiceChannelMessage**](IpMessagingV1ServiceChannelMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> IpMessagingV1ServiceRole UpdateRole(ctx, ServiceSid, Sid).Permission(Permission).Execute()



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
    // response from `UpdateRole`: IpMessagingV1ServiceRole
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

[**IpMessagingV1ServiceRole**](IpMessagingV1ServiceRole.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> IpMessagingV1Service UpdateService(ctx, Sid).ConsumptionReportInterval(ConsumptionReportInterval).DefaultChannelCreatorRoleSid(DefaultChannelCreatorRoleSid).DefaultChannelRoleSid(DefaultChannelRoleSid).DefaultServiceRoleSid(DefaultServiceRoleSid).FriendlyName(FriendlyName).LimitsChannelMembers(LimitsChannelMembers).LimitsUserChannels(LimitsUserChannels).NotificationsAddedToChannelEnabled(NotificationsAddedToChannelEnabled).NotificationsAddedToChannelTemplate(NotificationsAddedToChannelTemplate).NotificationsInvitedToChannelEnabled(NotificationsInvitedToChannelEnabled).NotificationsInvitedToChannelTemplate(NotificationsInvitedToChannelTemplate).NotificationsNewMessageEnabled(NotificationsNewMessageEnabled).NotificationsNewMessageTemplate(NotificationsNewMessageTemplate).NotificationsRemovedFromChannelEnabled(NotificationsRemovedFromChannelEnabled).NotificationsRemovedFromChannelTemplate(NotificationsRemovedFromChannelTemplate).PostWebhookUrl(PostWebhookUrl).PreWebhookUrl(PreWebhookUrl).ReachabilityEnabled(ReachabilityEnabled).ReadStatusEnabled(ReadStatusEnabled).TypingIndicatorTimeout(TypingIndicatorTimeout).WebhookFilters(WebhookFilters).WebhookMethod(WebhookMethod).WebhooksOnChannelAddMethod(WebhooksOnChannelAddMethod).WebhooksOnChannelAddUrl(WebhooksOnChannelAddUrl).WebhooksOnChannelAddedMethod(WebhooksOnChannelAddedMethod).WebhooksOnChannelAddedUrl(WebhooksOnChannelAddedUrl).WebhooksOnChannelDestroyMethod(WebhooksOnChannelDestroyMethod).WebhooksOnChannelDestroyUrl(WebhooksOnChannelDestroyUrl).WebhooksOnChannelDestroyedMethod(WebhooksOnChannelDestroyedMethod).WebhooksOnChannelDestroyedUrl(WebhooksOnChannelDestroyedUrl).WebhooksOnChannelUpdateMethod(WebhooksOnChannelUpdateMethod).WebhooksOnChannelUpdateUrl(WebhooksOnChannelUpdateUrl).WebhooksOnChannelUpdatedMethod(WebhooksOnChannelUpdatedMethod).WebhooksOnChannelUpdatedUrl(WebhooksOnChannelUpdatedUrl).WebhooksOnMemberAddMethod(WebhooksOnMemberAddMethod).WebhooksOnMemberAddUrl(WebhooksOnMemberAddUrl).WebhooksOnMemberAddedMethod(WebhooksOnMemberAddedMethod).WebhooksOnMemberAddedUrl(WebhooksOnMemberAddedUrl).WebhooksOnMemberRemoveMethod(WebhooksOnMemberRemoveMethod).WebhooksOnMemberRemoveUrl(WebhooksOnMemberRemoveUrl).WebhooksOnMemberRemovedMethod(WebhooksOnMemberRemovedMethod).WebhooksOnMemberRemovedUrl(WebhooksOnMemberRemovedUrl).WebhooksOnMessageRemoveMethod(WebhooksOnMessageRemoveMethod).WebhooksOnMessageRemoveUrl(WebhooksOnMessageRemoveUrl).WebhooksOnMessageRemovedMethod(WebhooksOnMessageRemovedMethod).WebhooksOnMessageRemovedUrl(WebhooksOnMessageRemovedUrl).WebhooksOnMessageSendMethod(WebhooksOnMessageSendMethod).WebhooksOnMessageSendUrl(WebhooksOnMessageSendUrl).WebhooksOnMessageSentMethod(WebhooksOnMessageSentMethod).WebhooksOnMessageSentUrl(WebhooksOnMessageSentUrl).WebhooksOnMessageUpdateMethod(WebhooksOnMessageUpdateMethod).WebhooksOnMessageUpdateUrl(WebhooksOnMessageUpdateUrl).WebhooksOnMessageUpdatedMethod(WebhooksOnMessageUpdatedMethod).WebhooksOnMessageUpdatedUrl(WebhooksOnMessageUpdatedUrl).Execute()



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
    NotificationsAddedToChannelEnabled := true // bool |  (optional)
    NotificationsAddedToChannelTemplate := "NotificationsAddedToChannelTemplate_example" // string |  (optional)
    NotificationsInvitedToChannelEnabled := true // bool |  (optional)
    NotificationsInvitedToChannelTemplate := "NotificationsInvitedToChannelTemplate_example" // string |  (optional)
    NotificationsNewMessageEnabled := true // bool |  (optional)
    NotificationsNewMessageTemplate := "NotificationsNewMessageTemplate_example" // string |  (optional)
    NotificationsRemovedFromChannelEnabled := true // bool |  (optional)
    NotificationsRemovedFromChannelTemplate := "NotificationsRemovedFromChannelTemplate_example" // string |  (optional)
    PostWebhookUrl := "PostWebhookUrl_example" // string |  (optional)
    PreWebhookUrl := "PreWebhookUrl_example" // string |  (optional)
    ReachabilityEnabled := true // bool |  (optional)
    ReadStatusEnabled := true // bool |  (optional)
    TypingIndicatorTimeout := int32(56) // int32 |  (optional)
    WebhookFilters := []string{"Inner_example"} // []string |  (optional)
    WebhookMethod := "WebhookMethod_example" // string |  (optional)
    WebhooksOnChannelAddMethod := "WebhooksOnChannelAddMethod_example" // string |  (optional)
    WebhooksOnChannelAddUrl := "WebhooksOnChannelAddUrl_example" // string |  (optional)
    WebhooksOnChannelAddedMethod := "WebhooksOnChannelAddedMethod_example" // string |  (optional)
    WebhooksOnChannelAddedUrl := "WebhooksOnChannelAddedUrl_example" // string |  (optional)
    WebhooksOnChannelDestroyMethod := "WebhooksOnChannelDestroyMethod_example" // string |  (optional)
    WebhooksOnChannelDestroyUrl := "WebhooksOnChannelDestroyUrl_example" // string |  (optional)
    WebhooksOnChannelDestroyedMethod := "WebhooksOnChannelDestroyedMethod_example" // string |  (optional)
    WebhooksOnChannelDestroyedUrl := "WebhooksOnChannelDestroyedUrl_example" // string |  (optional)
    WebhooksOnChannelUpdateMethod := "WebhooksOnChannelUpdateMethod_example" // string |  (optional)
    WebhooksOnChannelUpdateUrl := "WebhooksOnChannelUpdateUrl_example" // string |  (optional)
    WebhooksOnChannelUpdatedMethod := "WebhooksOnChannelUpdatedMethod_example" // string |  (optional)
    WebhooksOnChannelUpdatedUrl := "WebhooksOnChannelUpdatedUrl_example" // string |  (optional)
    WebhooksOnMemberAddMethod := "WebhooksOnMemberAddMethod_example" // string |  (optional)
    WebhooksOnMemberAddUrl := "WebhooksOnMemberAddUrl_example" // string |  (optional)
    WebhooksOnMemberAddedMethod := "WebhooksOnMemberAddedMethod_example" // string |  (optional)
    WebhooksOnMemberAddedUrl := "WebhooksOnMemberAddedUrl_example" // string |  (optional)
    WebhooksOnMemberRemoveMethod := "WebhooksOnMemberRemoveMethod_example" // string |  (optional)
    WebhooksOnMemberRemoveUrl := "WebhooksOnMemberRemoveUrl_example" // string |  (optional)
    WebhooksOnMemberRemovedMethod := "WebhooksOnMemberRemovedMethod_example" // string |  (optional)
    WebhooksOnMemberRemovedUrl := "WebhooksOnMemberRemovedUrl_example" // string |  (optional)
    WebhooksOnMessageRemoveMethod := "WebhooksOnMessageRemoveMethod_example" // string |  (optional)
    WebhooksOnMessageRemoveUrl := "WebhooksOnMessageRemoveUrl_example" // string |  (optional)
    WebhooksOnMessageRemovedMethod := "WebhooksOnMessageRemovedMethod_example" // string |  (optional)
    WebhooksOnMessageRemovedUrl := "WebhooksOnMessageRemovedUrl_example" // string |  (optional)
    WebhooksOnMessageSendMethod := "WebhooksOnMessageSendMethod_example" // string |  (optional)
    WebhooksOnMessageSendUrl := "WebhooksOnMessageSendUrl_example" // string |  (optional)
    WebhooksOnMessageSentMethod := "WebhooksOnMessageSentMethod_example" // string |  (optional)
    WebhooksOnMessageSentUrl := "WebhooksOnMessageSentUrl_example" // string |  (optional)
    WebhooksOnMessageUpdateMethod := "WebhooksOnMessageUpdateMethod_example" // string |  (optional)
    WebhooksOnMessageUpdateUrl := "WebhooksOnMessageUpdateUrl_example" // string |  (optional)
    WebhooksOnMessageUpdatedMethod := "WebhooksOnMessageUpdatedMethod_example" // string |  (optional)
    WebhooksOnMessageUpdatedUrl := "WebhooksOnMessageUpdatedUrl_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateService(context.Background(), Sid).ConsumptionReportInterval(ConsumptionReportInterval).DefaultChannelCreatorRoleSid(DefaultChannelCreatorRoleSid).DefaultChannelRoleSid(DefaultChannelRoleSid).DefaultServiceRoleSid(DefaultServiceRoleSid).FriendlyName(FriendlyName).LimitsChannelMembers(LimitsChannelMembers).LimitsUserChannels(LimitsUserChannels).NotificationsAddedToChannelEnabled(NotificationsAddedToChannelEnabled).NotificationsAddedToChannelTemplate(NotificationsAddedToChannelTemplate).NotificationsInvitedToChannelEnabled(NotificationsInvitedToChannelEnabled).NotificationsInvitedToChannelTemplate(NotificationsInvitedToChannelTemplate).NotificationsNewMessageEnabled(NotificationsNewMessageEnabled).NotificationsNewMessageTemplate(NotificationsNewMessageTemplate).NotificationsRemovedFromChannelEnabled(NotificationsRemovedFromChannelEnabled).NotificationsRemovedFromChannelTemplate(NotificationsRemovedFromChannelTemplate).PostWebhookUrl(PostWebhookUrl).PreWebhookUrl(PreWebhookUrl).ReachabilityEnabled(ReachabilityEnabled).ReadStatusEnabled(ReadStatusEnabled).TypingIndicatorTimeout(TypingIndicatorTimeout).WebhookFilters(WebhookFilters).WebhookMethod(WebhookMethod).WebhooksOnChannelAddMethod(WebhooksOnChannelAddMethod).WebhooksOnChannelAddUrl(WebhooksOnChannelAddUrl).WebhooksOnChannelAddedMethod(WebhooksOnChannelAddedMethod).WebhooksOnChannelAddedUrl(WebhooksOnChannelAddedUrl).WebhooksOnChannelDestroyMethod(WebhooksOnChannelDestroyMethod).WebhooksOnChannelDestroyUrl(WebhooksOnChannelDestroyUrl).WebhooksOnChannelDestroyedMethod(WebhooksOnChannelDestroyedMethod).WebhooksOnChannelDestroyedUrl(WebhooksOnChannelDestroyedUrl).WebhooksOnChannelUpdateMethod(WebhooksOnChannelUpdateMethod).WebhooksOnChannelUpdateUrl(WebhooksOnChannelUpdateUrl).WebhooksOnChannelUpdatedMethod(WebhooksOnChannelUpdatedMethod).WebhooksOnChannelUpdatedUrl(WebhooksOnChannelUpdatedUrl).WebhooksOnMemberAddMethod(WebhooksOnMemberAddMethod).WebhooksOnMemberAddUrl(WebhooksOnMemberAddUrl).WebhooksOnMemberAddedMethod(WebhooksOnMemberAddedMethod).WebhooksOnMemberAddedUrl(WebhooksOnMemberAddedUrl).WebhooksOnMemberRemoveMethod(WebhooksOnMemberRemoveMethod).WebhooksOnMemberRemoveUrl(WebhooksOnMemberRemoveUrl).WebhooksOnMemberRemovedMethod(WebhooksOnMemberRemovedMethod).WebhooksOnMemberRemovedUrl(WebhooksOnMemberRemovedUrl).WebhooksOnMessageRemoveMethod(WebhooksOnMessageRemoveMethod).WebhooksOnMessageRemoveUrl(WebhooksOnMessageRemoveUrl).WebhooksOnMessageRemovedMethod(WebhooksOnMessageRemovedMethod).WebhooksOnMessageRemovedUrl(WebhooksOnMessageRemovedUrl).WebhooksOnMessageSendMethod(WebhooksOnMessageSendMethod).WebhooksOnMessageSendUrl(WebhooksOnMessageSendUrl).WebhooksOnMessageSentMethod(WebhooksOnMessageSentMethod).WebhooksOnMessageSentUrl(WebhooksOnMessageSentUrl).WebhooksOnMessageUpdateMethod(WebhooksOnMessageUpdateMethod).WebhooksOnMessageUpdateUrl(WebhooksOnMessageUpdateUrl).WebhooksOnMessageUpdatedMethod(WebhooksOnMessageUpdatedMethod).WebhooksOnMessageUpdatedUrl(WebhooksOnMessageUpdatedUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateService`: IpMessagingV1Service
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
 **NotificationsAddedToChannelEnabled** | **bool** | 
 **NotificationsAddedToChannelTemplate** | **string** | 
 **NotificationsInvitedToChannelEnabled** | **bool** | 
 **NotificationsInvitedToChannelTemplate** | **string** | 
 **NotificationsNewMessageEnabled** | **bool** | 
 **NotificationsNewMessageTemplate** | **string** | 
 **NotificationsRemovedFromChannelEnabled** | **bool** | 
 **NotificationsRemovedFromChannelTemplate** | **string** | 
 **PostWebhookUrl** | **string** | 
 **PreWebhookUrl** | **string** | 
 **ReachabilityEnabled** | **bool** | 
 **ReadStatusEnabled** | **bool** | 
 **TypingIndicatorTimeout** | **int32** | 
 **WebhookFilters** | **[]string** | 
 **WebhookMethod** | **string** | 
 **WebhooksOnChannelAddMethod** | **string** | 
 **WebhooksOnChannelAddUrl** | **string** | 
 **WebhooksOnChannelAddedMethod** | **string** | 
 **WebhooksOnChannelAddedUrl** | **string** | 
 **WebhooksOnChannelDestroyMethod** | **string** | 
 **WebhooksOnChannelDestroyUrl** | **string** | 
 **WebhooksOnChannelDestroyedMethod** | **string** | 
 **WebhooksOnChannelDestroyedUrl** | **string** | 
 **WebhooksOnChannelUpdateMethod** | **string** | 
 **WebhooksOnChannelUpdateUrl** | **string** | 
 **WebhooksOnChannelUpdatedMethod** | **string** | 
 **WebhooksOnChannelUpdatedUrl** | **string** | 
 **WebhooksOnMemberAddMethod** | **string** | 
 **WebhooksOnMemberAddUrl** | **string** | 
 **WebhooksOnMemberAddedMethod** | **string** | 
 **WebhooksOnMemberAddedUrl** | **string** | 
 **WebhooksOnMemberRemoveMethod** | **string** | 
 **WebhooksOnMemberRemoveUrl** | **string** | 
 **WebhooksOnMemberRemovedMethod** | **string** | 
 **WebhooksOnMemberRemovedUrl** | **string** | 
 **WebhooksOnMessageRemoveMethod** | **string** | 
 **WebhooksOnMessageRemoveUrl** | **string** | 
 **WebhooksOnMessageRemovedMethod** | **string** | 
 **WebhooksOnMessageRemovedUrl** | **string** | 
 **WebhooksOnMessageSendMethod** | **string** | 
 **WebhooksOnMessageSendUrl** | **string** | 
 **WebhooksOnMessageSentMethod** | **string** | 
 **WebhooksOnMessageSentUrl** | **string** | 
 **WebhooksOnMessageUpdateMethod** | **string** | 
 **WebhooksOnMessageUpdateUrl** | **string** | 
 **WebhooksOnMessageUpdatedMethod** | **string** | 
 **WebhooksOnMessageUpdatedUrl** | **string** | 

### Return type

[**IpMessagingV1Service**](IpMessagingV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> IpMessagingV1ServiceUser UpdateUser(ctx, ServiceSid, Sid).Attributes(Attributes).FriendlyName(FriendlyName).RoleSid(RoleSid).Execute()



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
    Attributes := "Attributes_example" // string |  (optional)
    FriendlyName := "FriendlyName_example" // string |  (optional)
    RoleSid := "RoleSid_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateUser(context.Background(), ServiceSid, Sid).Attributes(Attributes).FriendlyName(FriendlyName).RoleSid(RoleSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: IpMessagingV1ServiceUser
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


 **Attributes** | **string** | 
 **FriendlyName** | **string** | 
 **RoleSid** | **string** | 

### Return type

[**IpMessagingV1ServiceUser**](IpMessagingV1ServiceUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

