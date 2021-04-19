# \DefaultApi

All URIs are relative to *https://chat.twilio.com*

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

> ChatV1ServiceChannel CreateChannel(ctx, ServiceSid).Attributes(Attributes).FriendlyName(FriendlyName).Type(Type).UniqueName(UniqueName).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under.
    Attributes := "Attributes_example" // string | A valid JSON string that contains application-specific data. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the new resource. It can be up to 64 characters long. (optional)
    Type := "Type_example" // string | The visibility of the channel. Can be: `public` or `private` and defaults to `public`. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL. This value must be 64 characters or less in length and be unique within the Service. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateChannel(context.Background(), ServiceSid).Attributes(Attributes).FriendlyName(FriendlyName).Type(Type).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateChannel`: ChatV1ServiceChannel
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Attributes** | **string** | A valid JSON string that contains application-specific data.
 **FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
 **Type** | **string** | The visibility of the channel. Can be: &#x60;public&#x60; or &#x60;private&#x60; and defaults to &#x60;public&#x60;.
 **UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource&#39;s &#x60;sid&#x60; in the URL. This value must be 64 characters or less in length and be unique within the Service.

### Return type

[**ChatV1ServiceChannel**](ChatV1ServiceChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCredential

> ChatV1Credential CreateCredential(ctx).ApiKey(ApiKey).Certificate(Certificate).FriendlyName(FriendlyName).PrivateKey(PrivateKey).Sandbox(Sandbox).Secret(Secret).Type(Type).Execute()



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
    ApiKey := "ApiKey_example" // string | [GCM only] The API key for the project that was obtained from the Google Developer console for your GCM Service application credential. (optional)
    Certificate := "Certificate_example" // string | [APN only] The URL encoded representation of the certificate. For example,  `-----BEGIN CERTIFICATE----- MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEFBQAwgZYxCzAJBgNV.....A== -----END CERTIFICATE-----` (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the new resource. It can be up to 64 characters long. (optional)
    PrivateKey := "PrivateKey_example" // string | [APN only] The URL encoded representation of the private key. For example, `-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fGgvCI1l9s+cmBY3WIz+cUDqmxiieR. -----END RSA PRIVATE KEY-----` (optional)
    Sandbox := true // bool | [APN only] Whether to send the credential to sandbox APNs. Can be `true` to send to sandbox APNs or `false` to send to production. (optional)
    Secret := "Secret_example" // string | [FCM only] The **Server key** of your project from the Firebase console, found under Settings / Cloud messaging. (optional)
    Type := "Type_example" // string | The type of push-notification service the credential is for. Can be: `gcm`, `fcm`, or `apn`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateCredential(context.Background()).ApiKey(ApiKey).Certificate(Certificate).FriendlyName(FriendlyName).PrivateKey(PrivateKey).Sandbox(Sandbox).Secret(Secret).Type(Type).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCredential`: ChatV1Credential
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCredential`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateCredentialParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **ApiKey** | **string** | [GCM only] The API key for the project that was obtained from the Google Developer console for your GCM Service application credential.
 **Certificate** | **string** | [APN only] The URL encoded representation of the certificate. For example,  &#x60;-----BEGIN CERTIFICATE----- MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEFBQAwgZYxCzAJBgNV.....A&#x3D;&#x3D; -----END CERTIFICATE-----&#x60;
 **FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
 **PrivateKey** | **string** | [APN only] The URL encoded representation of the private key. For example, &#x60;-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fGgvCI1l9s+cmBY3WIz+cUDqmxiieR. -----END RSA PRIVATE KEY-----&#x60;
 **Sandbox** | **bool** | [APN only] Whether to send the credential to sandbox APNs. Can be &#x60;true&#x60; to send to sandbox APNs or &#x60;false&#x60; to send to production.
 **Secret** | **string** | [FCM only] The **Server key** of your project from the Firebase console, found under Settings / Cloud messaging.
 **Type** | **string** | The type of push-notification service the credential is for. Can be: &#x60;gcm&#x60;, &#x60;fcm&#x60;, or &#x60;apn&#x60;.

### Return type

[**ChatV1Credential**](ChatV1Credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvite

> ChatV1ServiceChannelInvite CreateInvite(ctx, ServiceSid, ChannelSid).Identity(Identity).RoleSid(RoleSid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under.
    ChannelSid := "ChannelSid_example" // string | The SID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the new resource belongs to.
    Identity := "Identity_example" // string | The `identity` value that uniquely identifies the new resource's [User](https://www.twilio.com/docs/api/chat/rest/v1/user) within the [Service](https://www.twilio.com/docs/api/chat/rest/v1/service). See [access tokens](https://www.twilio.com/docs/api/chat/guides/create-tokens) for more info. (optional)
    RoleSid := "RoleSid_example" // string | The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) assigned to the new member. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateInvite(context.Background(), ServiceSid, ChannelSid).Identity(Identity).RoleSid(RoleSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInvite`: ChatV1ServiceChannelInvite
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the new resource belongs to.

### Other Parameters

Other parameters are passed through a pointer to a CreateInviteParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Identity** | **string** | The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s [User](https://www.twilio.com/docs/api/chat/rest/v1/user) within the [Service](https://www.twilio.com/docs/api/chat/rest/v1/service). See [access tokens](https://www.twilio.com/docs/api/chat/guides/create-tokens) for more info.
 **RoleSid** | **string** | The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) assigned to the new member.

### Return type

[**ChatV1ServiceChannelInvite**](ChatV1ServiceChannelInvite.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMember

> ChatV1ServiceChannelMember CreateMember(ctx, ServiceSid, ChannelSid).Identity(Identity).RoleSid(RoleSid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under.
    ChannelSid := "ChannelSid_example" // string | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the new member belongs to. Can be the Channel resource's `sid` or `unique_name`.
    Identity := "Identity_example" // string | The `identity` value that uniquely identifies the new resource's [User](https://www.twilio.com/docs/api/chat/rest/v1/user) within the [Service](https://www.twilio.com/docs/api/chat/rest/services). See [access tokens](https://www.twilio.com/docs/api/chat/guides/create-tokens) for more details. (optional)
    RoleSid := "RoleSid_example" // string | The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) to assign to the member. The default roles are those specified on the [Service](https://www.twilio.com/docs/chat/api/services). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateMember(context.Background(), ServiceSid, ChannelSid).Identity(Identity).RoleSid(RoleSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMember`: ChatV1ServiceChannelMember
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under.
**ChannelSid** | **string** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the new member belongs to. Can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a CreateMemberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Identity** | **string** | The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s [User](https://www.twilio.com/docs/api/chat/rest/v1/user) within the [Service](https://www.twilio.com/docs/api/chat/rest/services). See [access tokens](https://www.twilio.com/docs/api/chat/guides/create-tokens) for more details.
 **RoleSid** | **string** | The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) to assign to the member. The default roles are those specified on the [Service](https://www.twilio.com/docs/chat/api/services).

### Return type

[**ChatV1ServiceChannelMember**](ChatV1ServiceChannelMember.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMessage

> ChatV1ServiceChannelMessage CreateMessage(ctx, ServiceSid, ChannelSid).Attributes(Attributes).Body(Body).From(From).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under.
    ChannelSid := "ChannelSid_example" // string | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the new resource belongs to. Can be the Channel resource's `sid` or `unique_name`.
    Attributes := "Attributes_example" // string | A valid JSON string that contains application-specific data. (optional)
    Body := "Body_example" // string | The message to send to the channel. Can also be an empty string or `null`, which sets the value as an empty string. You can send structured data in the body by serializing it as a string. (optional)
    From := "From_example" // string | The [identity](https://www.twilio.com/docs/api/chat/guides/identity) of the new message's author. The default value is `system`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateMessage(context.Background(), ServiceSid, ChannelSid).Attributes(Attributes).Body(Body).From(From).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMessage`: ChatV1ServiceChannelMessage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under.
**ChannelSid** | **string** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the new resource belongs to. Can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a CreateMessageParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Attributes** | **string** | A valid JSON string that contains application-specific data.
 **Body** | **string** | The message to send to the channel. Can also be an empty string or &#x60;null&#x60;, which sets the value as an empty string. You can send structured data in the body by serializing it as a string.
 **From** | **string** | The [identity](https://www.twilio.com/docs/api/chat/guides/identity) of the new message&#39;s author. The default value is &#x60;system&#x60;.

### Return type

[**ChatV1ServiceChannelMessage**](ChatV1ServiceChannelMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> ChatV1ServiceRole CreateRole(ctx, ServiceSid).FriendlyName(FriendlyName).Permission(Permission).Type(Type).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under.
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the new resource. It can be up to 64 characters long. (optional)
    Permission := []string{"Inner_example"} // []string | A permission that you grant to the new role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. The values for this parameter depend on the role's `type` and are described in the documentation. (optional)
    Type := "Type_example" // string | The type of role. Can be: `channel` for [Channel](https://www.twilio.com/docs/chat/api/channels) roles or `deployment` for [Service](https://www.twilio.com/docs/chat/api/services) roles. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateRole(context.Background(), ServiceSid).FriendlyName(FriendlyName).Permission(Permission).Type(Type).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRole`: ChatV1ServiceRole
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateRoleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
 **Permission** | **[]string** | A permission that you grant to the new role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. The values for this parameter depend on the role&#39;s &#x60;type&#x60; and are described in the documentation.
 **Type** | **string** | The type of role. Can be: &#x60;channel&#x60; for [Channel](https://www.twilio.com/docs/chat/api/channels) roles or &#x60;deployment&#x60; for [Service](https://www.twilio.com/docs/chat/api/services) roles.

### Return type

[**ChatV1ServiceRole**](ChatV1ServiceRole.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateService

> ChatV1Service CreateService(ctx).FriendlyName(FriendlyName).Execute()



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
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateService(context.Background()).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateService`: ChatV1Service
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateService`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.

### Return type

[**ChatV1Service**](ChatV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> ChatV1ServiceUser CreateUser(ctx, ServiceSid).Attributes(Attributes).FriendlyName(FriendlyName).Identity(Identity).RoleSid(RoleSid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under.
    Attributes := "Attributes_example" // string | A valid JSON string that contains application-specific data. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the new resource. This value is often used for display purposes. (optional)
    Identity := "Identity_example" // string | The `identity` value that uniquely identifies the new resource's [User](https://www.twilio.com/docs/api/chat/rest/v1/user) within the [Service](https://www.twilio.com/docs/api/chat/rest/v1/service). This value is often a username or email address. See the Identity documentation for more details. (optional)
    RoleSid := "RoleSid_example" // string | The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) assigned to the new User. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateUser(context.Background(), ServiceSid).Attributes(Attributes).FriendlyName(FriendlyName).Identity(Identity).RoleSid(RoleSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: ChatV1ServiceUser
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to create the resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateUserParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Attributes** | **string** | A valid JSON string that contains application-specific data.
 **FriendlyName** | **string** | A descriptive string that you create to describe the new resource. This value is often used for display purposes.
 **Identity** | **string** | The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s [User](https://www.twilio.com/docs/api/chat/rest/v1/user) within the [Service](https://www.twilio.com/docs/api/chat/rest/v1/service). This value is often a username or email address. See the Identity documentation for more details.
 **RoleSid** | **string** | The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) assigned to the new User.

### Return type

[**ChatV1ServiceUser**](ChatV1ServiceUser.md)

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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to delete the resource from.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Channel resource to delete.

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
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to delete the resource from.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Channel resource to delete.

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
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Credential resource to delete.

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
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Credential resource to delete.

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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to delete the resource from.
    ChannelSid := "ChannelSid_example" // string | The SID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the resource to delete belongs to.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Invite resource to delete.

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
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to delete the resource from.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the resource to delete belongs to.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Invite resource to delete.

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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to delete the resource from.
    ChannelSid := "ChannelSid_example" // string | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the message belongs to.  Can be the Channel's `sid` or `unique_name`.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Member resource to delete.

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
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to delete the resource from.
**ChannelSid** | **string** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the message belongs to.  Can be the Channel&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Member resource to delete.

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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to delete the resource from.
    ChannelSid := "ChannelSid_example" // string | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the message to delete belongs to.  Can be the Channel resource's `sid` or `unique_name`.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Message resource to delete.

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
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to delete the resource from.
**ChannelSid** | **string** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the message to delete belongs to.  Can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Message resource to delete.

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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to delete the resource from.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Role resource to delete.

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
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to delete the resource from.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Role resource to delete.

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
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Service resource to delete.

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
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Service resource to delete.

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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to delete the resource from.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the User resource to delete.

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
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to delete the resource from.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the User resource to delete.

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

> ChatV1ServiceChannel FetchChannel(ctx, ServiceSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to fetch the resource from.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Channel resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchChannel(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchChannel`: ChatV1ServiceChannel
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to fetch the resource from.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Channel resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ChatV1ServiceChannel**](ChatV1ServiceChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCredential

> ChatV1Credential FetchCredential(ctx, Sid).Execute()



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
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Credential resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchCredential(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCredential`: ChatV1Credential
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Credential resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchCredentialParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**ChatV1Credential**](ChatV1Credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchInvite

> ChatV1ServiceChannelInvite FetchInvite(ctx, ServiceSid, ChannelSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to fetch the resource from.
    ChannelSid := "ChannelSid_example" // string | The SID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the resource to fetch belongs to.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Invite resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchInvite(context.Background(), ServiceSid, ChannelSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchInvite`: ChatV1ServiceChannelInvite
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to fetch the resource from.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the resource to fetch belongs to.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Invite resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchInviteParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ChatV1ServiceChannelInvite**](ChatV1ServiceChannelInvite.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMember

> ChatV1ServiceChannelMember FetchMember(ctx, ServiceSid, ChannelSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to fetch the resource from.
    ChannelSid := "ChannelSid_example" // string | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the member to fetch belongs to. Can be the Channel resource's `sid` or `unique_name` value.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Member resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchMember(context.Background(), ServiceSid, ChannelSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchMember`: ChatV1ServiceChannelMember
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to fetch the resource from.
**ChannelSid** | **string** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the member to fetch belongs to. Can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60; value.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Member resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchMemberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ChatV1ServiceChannelMember**](ChatV1ServiceChannelMember.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMessage

> ChatV1ServiceChannelMessage FetchMessage(ctx, ServiceSid, ChannelSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to fetch the resource from.
    ChannelSid := "ChannelSid_example" // string | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the message to fetch belongs to. Can be the Channel's `sid` or `unique_name`.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Message resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchMessage(context.Background(), ServiceSid, ChannelSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchMessage`: ChatV1ServiceChannelMessage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to fetch the resource from.
**ChannelSid** | **string** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the message to fetch belongs to. Can be the Channel&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Message resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchMessageParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ChatV1ServiceChannelMessage**](ChatV1ServiceChannelMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRole

> ChatV1ServiceRole FetchRole(ctx, ServiceSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to fetch the resource from.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Role resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchRole(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRole`: ChatV1ServiceRole
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to fetch the resource from.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Role resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRoleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ChatV1ServiceRole**](ChatV1ServiceRole.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> ChatV1Service FetchService(ctx, Sid).Execute()



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
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Service resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchService(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchService`: ChatV1Service
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Service resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**ChatV1Service**](ChatV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUser

> ChatV1ServiceUser FetchUser(ctx, ServiceSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to fetch the resource from.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the User resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchUser(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchUser`: ChatV1ServiceUser
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to fetch the resource from.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the User resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchUserParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ChatV1ServiceUser**](ChatV1ServiceUser.md)

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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from.
    Type := []string{"Type_example"} // []string | The visibility of the Channels to read. Can be: `public` or `private` and defaults to `public`. (optional)
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
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Type** | **[]string** | The visibility of the Channels to read. Can be: &#x60;public&#x60; or &#x60;private&#x60; and defaults to &#x60;public&#x60;.
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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from.
    ChannelSid := "ChannelSid_example" // string | The SID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the resources to read belong to.
    Identity := []string{"Inner_example"} // []string | The [User](https://www.twilio.com/docs/api/chat/rest/v1/user)'s `identity` value of the resources to read. See [access tokens](https://www.twilio.com/docs/api/chat/guides/create-tokens) for more details. (optional)
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
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from.
**ChannelSid** | **string** | The SID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the resources to read belong to.

### Other Parameters

Other parameters are passed through a pointer to a ListInviteParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Identity** | **[]string** | The [User](https://www.twilio.com/docs/api/chat/rest/v1/user)&#39;s &#x60;identity&#x60; value of the resources to read. See [access tokens](https://www.twilio.com/docs/api/chat/guides/create-tokens) for more details.
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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from.
    ChannelSid := "ChannelSid_example" // string | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the members to read belong to. Can be the Channel resource's `sid` or `unique_name` value.
    Identity := []string{"Inner_example"} // []string | The [User](https://www.twilio.com/docs/api/chat/rest/v1/user)'s `identity` value of the resources to read. See [access tokens](https://www.twilio.com/docs/api/chat/guides/create-tokens) for more details. (optional)
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
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from.
**ChannelSid** | **string** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the members to read belong to. Can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60; value.

### Other Parameters

Other parameters are passed through a pointer to a ListMemberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Identity** | **[]string** | The [User](https://www.twilio.com/docs/api/chat/rest/v1/user)&#39;s &#x60;identity&#x60; value of the resources to read. See [access tokens](https://www.twilio.com/docs/api/chat/guides/create-tokens) for more details.
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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from.
    ChannelSid := "ChannelSid_example" // string | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the message to read belongs to. Can be the Channel's `sid` or `unique_name`.
    Order := "Order_example" // string | The sort order of the returned messages. Can be: `asc` (ascending) or `desc` (descending) with `asc` as the default. (optional)
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
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from.
**ChannelSid** | **string** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the message to read belongs to. Can be the Channel&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a ListMessageParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Order** | **string** | The sort order of the returned messages. Can be: &#x60;asc&#x60; (ascending) or &#x60;desc&#x60; (descending) with &#x60;asc&#x60; as the default.
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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from.
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
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from.

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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from.
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
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from.

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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from.
    UserSid := "UserSid_example" // string | The SID of the [User](https://www.twilio.com/docs/api/chat/rest/users) to read the User Channel resources from.
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
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to read the resources from.
**UserSid** | **string** | The SID of the [User](https://www.twilio.com/docs/api/chat/rest/users) to read the User Channel resources from.

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

> ChatV1ServiceChannel UpdateChannel(ctx, ServiceSid, Sid).Attributes(Attributes).FriendlyName(FriendlyName).UniqueName(UniqueName).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to update the resource from.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Channel resource to update.
    Attributes := "Attributes_example" // string | A valid JSON string that contains application-specific data. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL. This value must be 64 characters or less in length and be unique within the Service. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateChannel(context.Background(), ServiceSid, Sid).Attributes(Attributes).FriendlyName(FriendlyName).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateChannel`: ChatV1ServiceChannel
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to update the resource from.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Channel resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Attributes** | **string** | A valid JSON string that contains application-specific data.
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
 **UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource&#39;s &#x60;sid&#x60; in the URL. This value must be 64 characters or less in length and be unique within the Service.

### Return type

[**ChatV1ServiceChannel**](ChatV1ServiceChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredential

> ChatV1Credential UpdateCredential(ctx, Sid).ApiKey(ApiKey).Certificate(Certificate).FriendlyName(FriendlyName).PrivateKey(PrivateKey).Sandbox(Sandbox).Secret(Secret).Execute()



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
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Credential resource to update.
    ApiKey := "ApiKey_example" // string | [GCM only] The API key for the project that was obtained from the Google Developer console for your GCM Service application credential. (optional)
    Certificate := "Certificate_example" // string | [APN only] The URL encoded representation of the certificate. For example,  `-----BEGIN CERTIFICATE----- MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEFBQAwgZYxCzAJBgNV.....A== -----END CERTIFICATE-----` (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)
    PrivateKey := "PrivateKey_example" // string | [APN only] The URL encoded representation of the private key. For example, `-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fGgvCI1l9s+cmBY3WIz+cUDqmxiieR. -----END RSA PRIVATE KEY-----` (optional)
    Sandbox := true // bool | [APN only] Whether to send the credential to sandbox APNs. Can be `true` to send to sandbox APNs or `false` to send to production. (optional)
    Secret := "Secret_example" // string | [FCM only] The **Server key** of your project from the Firebase console, found under Settings / Cloud messaging. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateCredential(context.Background(), Sid).ApiKey(ApiKey).Certificate(Certificate).FriendlyName(FriendlyName).PrivateKey(PrivateKey).Sandbox(Sandbox).Secret(Secret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCredential`: ChatV1Credential
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Credential resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateCredentialParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **ApiKey** | **string** | [GCM only] The API key for the project that was obtained from the Google Developer console for your GCM Service application credential.
 **Certificate** | **string** | [APN only] The URL encoded representation of the certificate. For example,  &#x60;-----BEGIN CERTIFICATE----- MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEFBQAwgZYxCzAJBgNV.....A&#x3D;&#x3D; -----END CERTIFICATE-----&#x60;
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
 **PrivateKey** | **string** | [APN only] The URL encoded representation of the private key. For example, &#x60;-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fGgvCI1l9s+cmBY3WIz+cUDqmxiieR. -----END RSA PRIVATE KEY-----&#x60;
 **Sandbox** | **bool** | [APN only] Whether to send the credential to sandbox APNs. Can be &#x60;true&#x60; to send to sandbox APNs or &#x60;false&#x60; to send to production.
 **Secret** | **string** | [FCM only] The **Server key** of your project from the Firebase console, found under Settings / Cloud messaging.

### Return type

[**ChatV1Credential**](ChatV1Credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMember

> ChatV1ServiceChannelMember UpdateMember(ctx, ServiceSid, ChannelSid, Sid).LastConsumedMessageIndex(LastConsumedMessageIndex).RoleSid(RoleSid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to update the resource from.
    ChannelSid := "ChannelSid_example" // string | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the member to update belongs to. Can be the Channel resource's `sid` or `unique_name`.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Member resource to update.
    LastConsumedMessageIndex := int32(56) // int32 | The index of the last [Message](https://www.twilio.com/docs/api/chat/rest/messages) that the Member has read within the [Channel](https://www.twilio.com/docs/api/chat/rest/channels). (optional)
    RoleSid := "RoleSid_example" // string | The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) to assign to the member. The default roles are those specified on the [Service](https://www.twilio.com/docs/chat/api/services). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateMember(context.Background(), ServiceSid, ChannelSid, Sid).LastConsumedMessageIndex(LastConsumedMessageIndex).RoleSid(RoleSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMember`: ChatV1ServiceChannelMember
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to update the resource from.
**ChannelSid** | **string** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the member to update belongs to. Can be the Channel resource&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Member resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateMemberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **LastConsumedMessageIndex** | **int32** | The index of the last [Message](https://www.twilio.com/docs/api/chat/rest/messages) that the Member has read within the [Channel](https://www.twilio.com/docs/api/chat/rest/channels).
 **RoleSid** | **string** | The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) to assign to the member. The default roles are those specified on the [Service](https://www.twilio.com/docs/chat/api/services).

### Return type

[**ChatV1ServiceChannelMember**](ChatV1ServiceChannelMember.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessage

> ChatV1ServiceChannelMessage UpdateMessage(ctx, ServiceSid, ChannelSid, Sid).Attributes(Attributes).Body(Body).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to update the resource from.
    ChannelSid := "ChannelSid_example" // string | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the message belongs to. Can be the Channel's `sid` or `unique_name`.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Message resource to update.
    Attributes := "Attributes_example" // string | A valid JSON string that contains application-specific data. (optional)
    Body := "Body_example" // string | The message to send to the channel. Can also be an empty string or `null`, which sets the value as an empty string. You can send structured data in the body by serializing it as a string. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateMessage(context.Background(), ServiceSid, ChannelSid, Sid).Attributes(Attributes).Body(Body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMessage`: ChatV1ServiceChannelMessage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to update the resource from.
**ChannelSid** | **string** | The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the message belongs to. Can be the Channel&#39;s &#x60;sid&#x60; or &#x60;unique_name&#x60;.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Message resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateMessageParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **Attributes** | **string** | A valid JSON string that contains application-specific data.
 **Body** | **string** | The message to send to the channel. Can also be an empty string or &#x60;null&#x60;, which sets the value as an empty string. You can send structured data in the body by serializing it as a string.

### Return type

[**ChatV1ServiceChannelMessage**](ChatV1ServiceChannelMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> ChatV1ServiceRole UpdateRole(ctx, ServiceSid, Sid).Permission(Permission).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to update the resource from.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Role resource to update.
    Permission := []string{"Inner_example"} // []string | A permission that you grant to the role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. The values for this parameter depend on the role's `type` and are described in the documentation. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateRole(context.Background(), ServiceSid, Sid).Permission(Permission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRole`: ChatV1ServiceRole
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to update the resource from.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Role resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRoleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Permission** | **[]string** | A permission that you grant to the role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. The values for this parameter depend on the role&#39;s &#x60;type&#x60; and are described in the documentation.

### Return type

[**ChatV1ServiceRole**](ChatV1ServiceRole.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> ChatV1Service UpdateService(ctx, Sid).ConsumptionReportInterval(ConsumptionReportInterval).DefaultChannelCreatorRoleSid(DefaultChannelCreatorRoleSid).DefaultChannelRoleSid(DefaultChannelRoleSid).DefaultServiceRoleSid(DefaultServiceRoleSid).FriendlyName(FriendlyName).LimitsChannelMembers(LimitsChannelMembers).LimitsUserChannels(LimitsUserChannels).NotificationsAddedToChannelEnabled(NotificationsAddedToChannelEnabled).NotificationsAddedToChannelTemplate(NotificationsAddedToChannelTemplate).NotificationsInvitedToChannelEnabled(NotificationsInvitedToChannelEnabled).NotificationsInvitedToChannelTemplate(NotificationsInvitedToChannelTemplate).NotificationsNewMessageEnabled(NotificationsNewMessageEnabled).NotificationsNewMessageTemplate(NotificationsNewMessageTemplate).NotificationsRemovedFromChannelEnabled(NotificationsRemovedFromChannelEnabled).NotificationsRemovedFromChannelTemplate(NotificationsRemovedFromChannelTemplate).PostWebhookUrl(PostWebhookUrl).PreWebhookUrl(PreWebhookUrl).ReachabilityEnabled(ReachabilityEnabled).ReadStatusEnabled(ReadStatusEnabled).TypingIndicatorTimeout(TypingIndicatorTimeout).WebhookFilters(WebhookFilters).WebhookMethod(WebhookMethod).WebhooksOnChannelAddMethod(WebhooksOnChannelAddMethod).WebhooksOnChannelAddUrl(WebhooksOnChannelAddUrl).WebhooksOnChannelAddedMethod(WebhooksOnChannelAddedMethod).WebhooksOnChannelAddedUrl(WebhooksOnChannelAddedUrl).WebhooksOnChannelDestroyMethod(WebhooksOnChannelDestroyMethod).WebhooksOnChannelDestroyUrl(WebhooksOnChannelDestroyUrl).WebhooksOnChannelDestroyedMethod(WebhooksOnChannelDestroyedMethod).WebhooksOnChannelDestroyedUrl(WebhooksOnChannelDestroyedUrl).WebhooksOnChannelUpdateMethod(WebhooksOnChannelUpdateMethod).WebhooksOnChannelUpdateUrl(WebhooksOnChannelUpdateUrl).WebhooksOnChannelUpdatedMethod(WebhooksOnChannelUpdatedMethod).WebhooksOnChannelUpdatedUrl(WebhooksOnChannelUpdatedUrl).WebhooksOnMemberAddMethod(WebhooksOnMemberAddMethod).WebhooksOnMemberAddUrl(WebhooksOnMemberAddUrl).WebhooksOnMemberAddedMethod(WebhooksOnMemberAddedMethod).WebhooksOnMemberAddedUrl(WebhooksOnMemberAddedUrl).WebhooksOnMemberRemoveMethod(WebhooksOnMemberRemoveMethod).WebhooksOnMemberRemoveUrl(WebhooksOnMemberRemoveUrl).WebhooksOnMemberRemovedMethod(WebhooksOnMemberRemovedMethod).WebhooksOnMemberRemovedUrl(WebhooksOnMemberRemovedUrl).WebhooksOnMessageRemoveMethod(WebhooksOnMessageRemoveMethod).WebhooksOnMessageRemoveUrl(WebhooksOnMessageRemoveUrl).WebhooksOnMessageRemovedMethod(WebhooksOnMessageRemovedMethod).WebhooksOnMessageRemovedUrl(WebhooksOnMessageRemovedUrl).WebhooksOnMessageSendMethod(WebhooksOnMessageSendMethod).WebhooksOnMessageSendUrl(WebhooksOnMessageSendUrl).WebhooksOnMessageSentMethod(WebhooksOnMessageSentMethod).WebhooksOnMessageSentUrl(WebhooksOnMessageSentUrl).WebhooksOnMessageUpdateMethod(WebhooksOnMessageUpdateMethod).WebhooksOnMessageUpdateUrl(WebhooksOnMessageUpdateUrl).WebhooksOnMessageUpdatedMethod(WebhooksOnMessageUpdatedMethod).WebhooksOnMessageUpdatedUrl(WebhooksOnMessageUpdatedUrl).Execute()



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
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Service resource to update.
    ConsumptionReportInterval := int32(56) // int32 | DEPRECATED. The interval in seconds between consumption reports submission batches from client endpoints. (optional)
    DefaultChannelCreatorRoleSid := "DefaultChannelCreatorRoleSid_example" // string | The channel role assigned to a channel creator when they join a new channel. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details. (optional)
    DefaultChannelRoleSid := "DefaultChannelRoleSid_example" // string | The channel role assigned to users when they are added to a channel. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details. (optional)
    DefaultServiceRoleSid := "DefaultServiceRoleSid_example" // string | The service role assigned to users when they are added to the service. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)
    LimitsChannelMembers := int32(56) // int32 | The maximum number of Members that can be added to Channels within this Service. Can be up to 1,000. (optional)
    LimitsUserChannels := int32(56) // int32 | The maximum number of Channels Users can be a Member of within this Service. Can be up to 1,000. (optional)
    NotificationsAddedToChannelEnabled := true // bool | Whether to send a notification when a member is added to a channel. Can be: `true` or `false` and the default is `false`. (optional)
    NotificationsAddedToChannelTemplate := "NotificationsAddedToChannelTemplate_example" // string | The template to use to create the notification text displayed when a member is added to a channel and `notifications.added_to_channel.enabled` is `true`. (optional)
    NotificationsInvitedToChannelEnabled := true // bool | Whether to send a notification when a user is invited to a channel. Can be: `true` or `false` and the default is `false`. (optional)
    NotificationsInvitedToChannelTemplate := "NotificationsInvitedToChannelTemplate_example" // string | The template to use to create the notification text displayed when a user is invited to a channel and `notifications.invited_to_channel.enabled` is `true`. (optional)
    NotificationsNewMessageEnabled := true // bool | Whether to send a notification when a new message is added to a channel. Can be: `true` or `false` and the default is `false`. (optional)
    NotificationsNewMessageTemplate := "NotificationsNewMessageTemplate_example" // string | The template to use to create the notification text displayed when a new message is added to a channel and `notifications.new_message.enabled` is `true`. (optional)
    NotificationsRemovedFromChannelEnabled := true // bool | Whether to send a notification to a user when they are removed from a channel. Can be: `true` or `false` and the default is `false`. (optional)
    NotificationsRemovedFromChannelTemplate := "NotificationsRemovedFromChannelTemplate_example" // string | The template to use to create the notification text displayed to a user when they are removed from a channel and `notifications.removed_from_channel.enabled` is `true`. (optional)
    PostWebhookUrl := "PostWebhookUrl_example" // string | The URL for post-event webhooks, which are called by using the `webhook_method`. See [Webhook Events](https://www.twilio.com/docs/api/chat/webhooks) for more details. (optional)
    PreWebhookUrl := "PreWebhookUrl_example" // string | The URL for pre-event webhooks, which are called by using the `webhook_method`. See [Webhook Events](https://www.twilio.com/docs/api/chat/webhooks) for more details. (optional)
    ReachabilityEnabled := true // bool | Whether to enable the [Reachability Indicator](https://www.twilio.com/docs/chat/reachability-indicator) for this Service instance. The default is `false`. (optional)
    ReadStatusEnabled := true // bool | Whether to enable the [Message Consumption Horizon](https://www.twilio.com/docs/chat/consumption-horizon) feature. The default is `true`. (optional)
    TypingIndicatorTimeout := int32(56) // int32 | How long in seconds after a `started typing` event until clients should assume that user is no longer typing, even if no `ended typing` message was received.  The default is 5 seconds. (optional)
    WebhookFilters := []string{"Inner_example"} // []string | The list of WebHook events that are enabled for this Service instance. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details. (optional)
    WebhookMethod := "WebhookMethod_example" // string | The HTTP method to use for calls to the `pre_webhook_url` and `post_webhook_url` webhooks.  Can be: `POST` or `GET` and the default is `POST`. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details. (optional)
    WebhooksOnChannelAddMethod := "WebhooksOnChannelAddMethod_example" // string | The HTTP method to use when calling the `webhooks.on_channel_add.url`. (optional)
    WebhooksOnChannelAddUrl := "WebhooksOnChannelAddUrl_example" // string | The URL of the webhook to call in response to the `on_channel_add` event using the `webhooks.on_channel_add.method` HTTP method. (optional)
    WebhooksOnChannelAddedMethod := "WebhooksOnChannelAddedMethod_example" // string | The URL of the webhook to call in response to the `on_channel_added` event`. (optional)
    WebhooksOnChannelAddedUrl := "WebhooksOnChannelAddedUrl_example" // string | The URL of the webhook to call in response to the `on_channel_added` event using the `webhooks.on_channel_added.method` HTTP method. (optional)
    WebhooksOnChannelDestroyMethod := "WebhooksOnChannelDestroyMethod_example" // string | The HTTP method to use when calling the `webhooks.on_channel_destroy.url`. (optional)
    WebhooksOnChannelDestroyUrl := "WebhooksOnChannelDestroyUrl_example" // string | The URL of the webhook to call in response to the `on_channel_destroy` event using the `webhooks.on_channel_destroy.method` HTTP method. (optional)
    WebhooksOnChannelDestroyedMethod := "WebhooksOnChannelDestroyedMethod_example" // string | The HTTP method to use when calling the `webhooks.on_channel_destroyed.url`. (optional)
    WebhooksOnChannelDestroyedUrl := "WebhooksOnChannelDestroyedUrl_example" // string | The URL of the webhook to call in response to the `on_channel_added` event using the `webhooks.on_channel_destroyed.method` HTTP method. (optional)
    WebhooksOnChannelUpdateMethod := "WebhooksOnChannelUpdateMethod_example" // string | The HTTP method to use when calling the `webhooks.on_channel_update.url`. (optional)
    WebhooksOnChannelUpdateUrl := "WebhooksOnChannelUpdateUrl_example" // string | The URL of the webhook to call in response to the `on_channel_update` event using the `webhooks.on_channel_update.method` HTTP method. (optional)
    WebhooksOnChannelUpdatedMethod := "WebhooksOnChannelUpdatedMethod_example" // string | The HTTP method to use when calling the `webhooks.on_channel_updated.url`. (optional)
    WebhooksOnChannelUpdatedUrl := "WebhooksOnChannelUpdatedUrl_example" // string | The URL of the webhook to call in response to the `on_channel_updated` event using the `webhooks.on_channel_updated.method` HTTP method. (optional)
    WebhooksOnMemberAddMethod := "WebhooksOnMemberAddMethod_example" // string | The HTTP method to use when calling the `webhooks.on_member_add.url`. (optional)
    WebhooksOnMemberAddUrl := "WebhooksOnMemberAddUrl_example" // string | The URL of the webhook to call in response to the `on_member_add` event using the `webhooks.on_member_add.method` HTTP method. (optional)
    WebhooksOnMemberAddedMethod := "WebhooksOnMemberAddedMethod_example" // string | The HTTP method to use when calling the `webhooks.on_channel_updated.url`. (optional)
    WebhooksOnMemberAddedUrl := "WebhooksOnMemberAddedUrl_example" // string | The URL of the webhook to call in response to the `on_channel_updated` event using the `webhooks.on_channel_updated.method` HTTP method. (optional)
    WebhooksOnMemberRemoveMethod := "WebhooksOnMemberRemoveMethod_example" // string | The HTTP method to use when calling the `webhooks.on_member_remove.url`. (optional)
    WebhooksOnMemberRemoveUrl := "WebhooksOnMemberRemoveUrl_example" // string | The URL of the webhook to call in response to the `on_member_remove` event using the `webhooks.on_member_remove.method` HTTP method. (optional)
    WebhooksOnMemberRemovedMethod := "WebhooksOnMemberRemovedMethod_example" // string | The HTTP method to use when calling the `webhooks.on_member_removed.url`. (optional)
    WebhooksOnMemberRemovedUrl := "WebhooksOnMemberRemovedUrl_example" // string | The URL of the webhook to call in response to the `on_member_removed` event using the `webhooks.on_member_removed.method` HTTP method. (optional)
    WebhooksOnMessageRemoveMethod := "WebhooksOnMessageRemoveMethod_example" // string | The HTTP method to use when calling the `webhooks.on_message_remove.url`. (optional)
    WebhooksOnMessageRemoveUrl := "WebhooksOnMessageRemoveUrl_example" // string | The URL of the webhook to call in response to the `on_message_remove` event using the `webhooks.on_message_remove.method` HTTP method. (optional)
    WebhooksOnMessageRemovedMethod := "WebhooksOnMessageRemovedMethod_example" // string | The HTTP method to use when calling the `webhooks.on_message_removed.url`. (optional)
    WebhooksOnMessageRemovedUrl := "WebhooksOnMessageRemovedUrl_example" // string | The URL of the webhook to call in response to the `on_message_removed` event using the `webhooks.on_message_removed.method` HTTP method. (optional)
    WebhooksOnMessageSendMethod := "WebhooksOnMessageSendMethod_example" // string | The HTTP method to use when calling the `webhooks.on_message_send.url`. (optional)
    WebhooksOnMessageSendUrl := "WebhooksOnMessageSendUrl_example" // string | The URL of the webhook to call in response to the `on_message_send` event using the `webhooks.on_message_send.method` HTTP method. (optional)
    WebhooksOnMessageSentMethod := "WebhooksOnMessageSentMethod_example" // string | The URL of the webhook to call in response to the `on_message_sent` event`. (optional)
    WebhooksOnMessageSentUrl := "WebhooksOnMessageSentUrl_example" // string | The URL of the webhook to call in response to the `on_message_sent` event using the `webhooks.on_message_sent.method` HTTP method. (optional)
    WebhooksOnMessageUpdateMethod := "WebhooksOnMessageUpdateMethod_example" // string | The HTTP method to use when calling the `webhooks.on_message_update.url`. (optional)
    WebhooksOnMessageUpdateUrl := "WebhooksOnMessageUpdateUrl_example" // string | The URL of the webhook to call in response to the `on_message_update` event using the `webhooks.on_message_update.method` HTTP method. (optional)
    WebhooksOnMessageUpdatedMethod := "WebhooksOnMessageUpdatedMethod_example" // string | The HTTP method to use when calling the `webhooks.on_message_updated.url`. (optional)
    WebhooksOnMessageUpdatedUrl := "WebhooksOnMessageUpdatedUrl_example" // string | The URL of the webhook to call in response to the `on_message_updated` event using the `webhooks.on_message_updated.method` HTTP method. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateService(context.Background(), Sid).ConsumptionReportInterval(ConsumptionReportInterval).DefaultChannelCreatorRoleSid(DefaultChannelCreatorRoleSid).DefaultChannelRoleSid(DefaultChannelRoleSid).DefaultServiceRoleSid(DefaultServiceRoleSid).FriendlyName(FriendlyName).LimitsChannelMembers(LimitsChannelMembers).LimitsUserChannels(LimitsUserChannels).NotificationsAddedToChannelEnabled(NotificationsAddedToChannelEnabled).NotificationsAddedToChannelTemplate(NotificationsAddedToChannelTemplate).NotificationsInvitedToChannelEnabled(NotificationsInvitedToChannelEnabled).NotificationsInvitedToChannelTemplate(NotificationsInvitedToChannelTemplate).NotificationsNewMessageEnabled(NotificationsNewMessageEnabled).NotificationsNewMessageTemplate(NotificationsNewMessageTemplate).NotificationsRemovedFromChannelEnabled(NotificationsRemovedFromChannelEnabled).NotificationsRemovedFromChannelTemplate(NotificationsRemovedFromChannelTemplate).PostWebhookUrl(PostWebhookUrl).PreWebhookUrl(PreWebhookUrl).ReachabilityEnabled(ReachabilityEnabled).ReadStatusEnabled(ReadStatusEnabled).TypingIndicatorTimeout(TypingIndicatorTimeout).WebhookFilters(WebhookFilters).WebhookMethod(WebhookMethod).WebhooksOnChannelAddMethod(WebhooksOnChannelAddMethod).WebhooksOnChannelAddUrl(WebhooksOnChannelAddUrl).WebhooksOnChannelAddedMethod(WebhooksOnChannelAddedMethod).WebhooksOnChannelAddedUrl(WebhooksOnChannelAddedUrl).WebhooksOnChannelDestroyMethod(WebhooksOnChannelDestroyMethod).WebhooksOnChannelDestroyUrl(WebhooksOnChannelDestroyUrl).WebhooksOnChannelDestroyedMethod(WebhooksOnChannelDestroyedMethod).WebhooksOnChannelDestroyedUrl(WebhooksOnChannelDestroyedUrl).WebhooksOnChannelUpdateMethod(WebhooksOnChannelUpdateMethod).WebhooksOnChannelUpdateUrl(WebhooksOnChannelUpdateUrl).WebhooksOnChannelUpdatedMethod(WebhooksOnChannelUpdatedMethod).WebhooksOnChannelUpdatedUrl(WebhooksOnChannelUpdatedUrl).WebhooksOnMemberAddMethod(WebhooksOnMemberAddMethod).WebhooksOnMemberAddUrl(WebhooksOnMemberAddUrl).WebhooksOnMemberAddedMethod(WebhooksOnMemberAddedMethod).WebhooksOnMemberAddedUrl(WebhooksOnMemberAddedUrl).WebhooksOnMemberRemoveMethod(WebhooksOnMemberRemoveMethod).WebhooksOnMemberRemoveUrl(WebhooksOnMemberRemoveUrl).WebhooksOnMemberRemovedMethod(WebhooksOnMemberRemovedMethod).WebhooksOnMemberRemovedUrl(WebhooksOnMemberRemovedUrl).WebhooksOnMessageRemoveMethod(WebhooksOnMessageRemoveMethod).WebhooksOnMessageRemoveUrl(WebhooksOnMessageRemoveUrl).WebhooksOnMessageRemovedMethod(WebhooksOnMessageRemovedMethod).WebhooksOnMessageRemovedUrl(WebhooksOnMessageRemovedUrl).WebhooksOnMessageSendMethod(WebhooksOnMessageSendMethod).WebhooksOnMessageSendUrl(WebhooksOnMessageSendUrl).WebhooksOnMessageSentMethod(WebhooksOnMessageSentMethod).WebhooksOnMessageSentUrl(WebhooksOnMessageSentUrl).WebhooksOnMessageUpdateMethod(WebhooksOnMessageUpdateMethod).WebhooksOnMessageUpdateUrl(WebhooksOnMessageUpdateUrl).WebhooksOnMessageUpdatedMethod(WebhooksOnMessageUpdatedMethod).WebhooksOnMessageUpdatedUrl(WebhooksOnMessageUpdatedUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateService`: ChatV1Service
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Service resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **ConsumptionReportInterval** | **int32** | DEPRECATED. The interval in seconds between consumption reports submission batches from client endpoints.
 **DefaultChannelCreatorRoleSid** | **string** | The channel role assigned to a channel creator when they join a new channel. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details.
 **DefaultChannelRoleSid** | **string** | The channel role assigned to users when they are added to a channel. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details.
 **DefaultServiceRoleSid** | **string** | The service role assigned to users when they are added to the service. See the [Roles endpoint](https://www.twilio.com/docs/chat/api/roles) for more details.
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
 **LimitsChannelMembers** | **int32** | The maximum number of Members that can be added to Channels within this Service. Can be up to 1,000.
 **LimitsUserChannels** | **int32** | The maximum number of Channels Users can be a Member of within this Service. Can be up to 1,000.
 **NotificationsAddedToChannelEnabled** | **bool** | Whether to send a notification when a member is added to a channel. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **NotificationsAddedToChannelTemplate** | **string** | The template to use to create the notification text displayed when a member is added to a channel and &#x60;notifications.added_to_channel.enabled&#x60; is &#x60;true&#x60;.
 **NotificationsInvitedToChannelEnabled** | **bool** | Whether to send a notification when a user is invited to a channel. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **NotificationsInvitedToChannelTemplate** | **string** | The template to use to create the notification text displayed when a user is invited to a channel and &#x60;notifications.invited_to_channel.enabled&#x60; is &#x60;true&#x60;.
 **NotificationsNewMessageEnabled** | **bool** | Whether to send a notification when a new message is added to a channel. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **NotificationsNewMessageTemplate** | **string** | The template to use to create the notification text displayed when a new message is added to a channel and &#x60;notifications.new_message.enabled&#x60; is &#x60;true&#x60;.
 **NotificationsRemovedFromChannelEnabled** | **bool** | Whether to send a notification to a user when they are removed from a channel. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **NotificationsRemovedFromChannelTemplate** | **string** | The template to use to create the notification text displayed to a user when they are removed from a channel and &#x60;notifications.removed_from_channel.enabled&#x60; is &#x60;true&#x60;.
 **PostWebhookUrl** | **string** | The URL for post-event webhooks, which are called by using the &#x60;webhook_method&#x60;. See [Webhook Events](https://www.twilio.com/docs/api/chat/webhooks) for more details.
 **PreWebhookUrl** | **string** | The URL for pre-event webhooks, which are called by using the &#x60;webhook_method&#x60;. See [Webhook Events](https://www.twilio.com/docs/api/chat/webhooks) for more details.
 **ReachabilityEnabled** | **bool** | Whether to enable the [Reachability Indicator](https://www.twilio.com/docs/chat/reachability-indicator) for this Service instance. The default is &#x60;false&#x60;.
 **ReadStatusEnabled** | **bool** | Whether to enable the [Message Consumption Horizon](https://www.twilio.com/docs/chat/consumption-horizon) feature. The default is &#x60;true&#x60;.
 **TypingIndicatorTimeout** | **int32** | How long in seconds after a &#x60;started typing&#x60; event until clients should assume that user is no longer typing, even if no &#x60;ended typing&#x60; message was received.  The default is 5 seconds.
 **WebhookFilters** | **[]string** | The list of WebHook events that are enabled for this Service instance. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
 **WebhookMethod** | **string** | The HTTP method to use for calls to the &#x60;pre_webhook_url&#x60; and &#x60;post_webhook_url&#x60; webhooks.  Can be: &#x60;POST&#x60; or &#x60;GET&#x60; and the default is &#x60;POST&#x60;. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
 **WebhooksOnChannelAddMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_channel_add.url&#x60;.
 **WebhooksOnChannelAddUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_channel_add&#x60; event using the &#x60;webhooks.on_channel_add.method&#x60; HTTP method.
 **WebhooksOnChannelAddedMethod** | **string** | The URL of the webhook to call in response to the &#x60;on_channel_added&#x60; event&#x60;.
 **WebhooksOnChannelAddedUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_channel_added&#x60; event using the &#x60;webhooks.on_channel_added.method&#x60; HTTP method.
 **WebhooksOnChannelDestroyMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_channel_destroy.url&#x60;.
 **WebhooksOnChannelDestroyUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_channel_destroy&#x60; event using the &#x60;webhooks.on_channel_destroy.method&#x60; HTTP method.
 **WebhooksOnChannelDestroyedMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_channel_destroyed.url&#x60;.
 **WebhooksOnChannelDestroyedUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_channel_added&#x60; event using the &#x60;webhooks.on_channel_destroyed.method&#x60; HTTP method.
 **WebhooksOnChannelUpdateMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_channel_update.url&#x60;.
 **WebhooksOnChannelUpdateUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_channel_update&#x60; event using the &#x60;webhooks.on_channel_update.method&#x60; HTTP method.
 **WebhooksOnChannelUpdatedMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_channel_updated.url&#x60;.
 **WebhooksOnChannelUpdatedUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_channel_updated&#x60; event using the &#x60;webhooks.on_channel_updated.method&#x60; HTTP method.
 **WebhooksOnMemberAddMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_member_add.url&#x60;.
 **WebhooksOnMemberAddUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_member_add&#x60; event using the &#x60;webhooks.on_member_add.method&#x60; HTTP method.
 **WebhooksOnMemberAddedMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_channel_updated.url&#x60;.
 **WebhooksOnMemberAddedUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_channel_updated&#x60; event using the &#x60;webhooks.on_channel_updated.method&#x60; HTTP method.
 **WebhooksOnMemberRemoveMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_member_remove.url&#x60;.
 **WebhooksOnMemberRemoveUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_member_remove&#x60; event using the &#x60;webhooks.on_member_remove.method&#x60; HTTP method.
 **WebhooksOnMemberRemovedMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_member_removed.url&#x60;.
 **WebhooksOnMemberRemovedUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_member_removed&#x60; event using the &#x60;webhooks.on_member_removed.method&#x60; HTTP method.
 **WebhooksOnMessageRemoveMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_message_remove.url&#x60;.
 **WebhooksOnMessageRemoveUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_message_remove&#x60; event using the &#x60;webhooks.on_message_remove.method&#x60; HTTP method.
 **WebhooksOnMessageRemovedMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_message_removed.url&#x60;.
 **WebhooksOnMessageRemovedUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_message_removed&#x60; event using the &#x60;webhooks.on_message_removed.method&#x60; HTTP method.
 **WebhooksOnMessageSendMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_message_send.url&#x60;.
 **WebhooksOnMessageSendUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_message_send&#x60; event using the &#x60;webhooks.on_message_send.method&#x60; HTTP method.
 **WebhooksOnMessageSentMethod** | **string** | The URL of the webhook to call in response to the &#x60;on_message_sent&#x60; event&#x60;.
 **WebhooksOnMessageSentUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_message_sent&#x60; event using the &#x60;webhooks.on_message_sent.method&#x60; HTTP method.
 **WebhooksOnMessageUpdateMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_message_update.url&#x60;.
 **WebhooksOnMessageUpdateUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_message_update&#x60; event using the &#x60;webhooks.on_message_update.method&#x60; HTTP method.
 **WebhooksOnMessageUpdatedMethod** | **string** | The HTTP method to use when calling the &#x60;webhooks.on_message_updated.url&#x60;.
 **WebhooksOnMessageUpdatedUrl** | **string** | The URL of the webhook to call in response to the &#x60;on_message_updated&#x60; event using the &#x60;webhooks.on_message_updated.method&#x60; HTTP method.

### Return type

[**ChatV1Service**](ChatV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> ChatV1ServiceUser UpdateUser(ctx, ServiceSid, Sid).Attributes(Attributes).FriendlyName(FriendlyName).RoleSid(RoleSid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to update the resource from.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the User resource to update.
    Attributes := "Attributes_example" // string | A valid JSON string that contains application-specific data. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It is often used for display purposes. (optional)
    RoleSid := "RoleSid_example" // string | The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) assigned to this user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateUser(context.Background(), ServiceSid, Sid).Attributes(Attributes).FriendlyName(FriendlyName).RoleSid(RoleSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: ChatV1ServiceUser
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) to update the resource from.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the User resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateUserParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Attributes** | **string** | A valid JSON string that contains application-specific data.
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is often used for display purposes.
 **RoleSid** | **string** | The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) assigned to this user.

### Return type

[**ChatV1ServiceUser**](ChatV1ServiceUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

