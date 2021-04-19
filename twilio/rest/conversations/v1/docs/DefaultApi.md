# \DefaultApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConversation**](DefaultApi.md#CreateConversation) | **Post** /v1/Conversations | 
[**CreateConversationMessage**](DefaultApi.md#CreateConversationMessage) | **Post** /v1/Conversations/{ConversationSid}/Messages | 
[**CreateConversationParticipant**](DefaultApi.md#CreateConversationParticipant) | **Post** /v1/Conversations/{ConversationSid}/Participants | 
[**CreateConversationScopedWebhook**](DefaultApi.md#CreateConversationScopedWebhook) | **Post** /v1/Conversations/{ConversationSid}/Webhooks | 
[**CreateCredential**](DefaultApi.md#CreateCredential) | **Post** /v1/Credentials | 
[**CreateRole**](DefaultApi.md#CreateRole) | **Post** /v1/Roles | 
[**CreateService**](DefaultApi.md#CreateService) | **Post** /v1/Services | 
[**CreateServiceConversation**](DefaultApi.md#CreateServiceConversation) | **Post** /v1/Services/{ChatServiceSid}/Conversations | 
[**CreateServiceConversationMessage**](DefaultApi.md#CreateServiceConversationMessage) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages | 
[**CreateServiceConversationParticipant**](DefaultApi.md#CreateServiceConversationParticipant) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Participants | 
[**CreateServiceConversationScopedWebhook**](DefaultApi.md#CreateServiceConversationScopedWebhook) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks | 
[**CreateServiceRole**](DefaultApi.md#CreateServiceRole) | **Post** /v1/Services/{ChatServiceSid}/Roles | 
[**CreateServiceUser**](DefaultApi.md#CreateServiceUser) | **Post** /v1/Services/{ChatServiceSid}/Users | 
[**CreateUser**](DefaultApi.md#CreateUser) | **Post** /v1/Users | 
[**DeleteConversation**](DefaultApi.md#DeleteConversation) | **Delete** /v1/Conversations/{Sid} | 
[**DeleteConversationMessage**](DefaultApi.md#DeleteConversationMessage) | **Delete** /v1/Conversations/{ConversationSid}/Messages/{Sid} | 
[**DeleteConversationParticipant**](DefaultApi.md#DeleteConversationParticipant) | **Delete** /v1/Conversations/{ConversationSid}/Participants/{Sid} | 
[**DeleteConversationScopedWebhook**](DefaultApi.md#DeleteConversationScopedWebhook) | **Delete** /v1/Conversations/{ConversationSid}/Webhooks/{Sid} | 
[**DeleteCredential**](DefaultApi.md#DeleteCredential) | **Delete** /v1/Credentials/{Sid} | 
[**DeleteRole**](DefaultApi.md#DeleteRole) | **Delete** /v1/Roles/{Sid} | 
[**DeleteService**](DefaultApi.md#DeleteService) | **Delete** /v1/Services/{Sid} | 
[**DeleteServiceBinding**](DefaultApi.md#DeleteServiceBinding) | **Delete** /v1/Services/{ChatServiceSid}/Bindings/{Sid} | 
[**DeleteServiceConversation**](DefaultApi.md#DeleteServiceConversation) | **Delete** /v1/Services/{ChatServiceSid}/Conversations/{Sid} | 
[**DeleteServiceConversationMessage**](DefaultApi.md#DeleteServiceConversationMessage) | **Delete** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{Sid} | 
[**DeleteServiceConversationParticipant**](DefaultApi.md#DeleteServiceConversationParticipant) | **Delete** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Participants/{Sid} | 
[**DeleteServiceConversationScopedWebhook**](DefaultApi.md#DeleteServiceConversationScopedWebhook) | **Delete** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks/{Sid} | 
[**DeleteServiceRole**](DefaultApi.md#DeleteServiceRole) | **Delete** /v1/Services/{ChatServiceSid}/Roles/{Sid} | 
[**DeleteServiceUser**](DefaultApi.md#DeleteServiceUser) | **Delete** /v1/Services/{ChatServiceSid}/Users/{Sid} | 
[**DeleteUser**](DefaultApi.md#DeleteUser) | **Delete** /v1/Users/{Sid} | 
[**FetchConfiguration**](DefaultApi.md#FetchConfiguration) | **Get** /v1/Configuration | 
[**FetchConfigurationWebhook**](DefaultApi.md#FetchConfigurationWebhook) | **Get** /v1/Configuration/Webhooks | 
[**FetchConversation**](DefaultApi.md#FetchConversation) | **Get** /v1/Conversations/{Sid} | 
[**FetchConversationMessage**](DefaultApi.md#FetchConversationMessage) | **Get** /v1/Conversations/{ConversationSid}/Messages/{Sid} | 
[**FetchConversationMessageReceipt**](DefaultApi.md#FetchConversationMessageReceipt) | **Get** /v1/Conversations/{ConversationSid}/Messages/{MessageSid}/Receipts/{Sid} | 
[**FetchConversationParticipant**](DefaultApi.md#FetchConversationParticipant) | **Get** /v1/Conversations/{ConversationSid}/Participants/{Sid} | 
[**FetchConversationScopedWebhook**](DefaultApi.md#FetchConversationScopedWebhook) | **Get** /v1/Conversations/{ConversationSid}/Webhooks/{Sid} | 
[**FetchCredential**](DefaultApi.md#FetchCredential) | **Get** /v1/Credentials/{Sid} | 
[**FetchRole**](DefaultApi.md#FetchRole) | **Get** /v1/Roles/{Sid} | 
[**FetchService**](DefaultApi.md#FetchService) | **Get** /v1/Services/{Sid} | 
[**FetchServiceBinding**](DefaultApi.md#FetchServiceBinding) | **Get** /v1/Services/{ChatServiceSid}/Bindings/{Sid} | 
[**FetchServiceConfiguration**](DefaultApi.md#FetchServiceConfiguration) | **Get** /v1/Services/{ChatServiceSid}/Configuration | 
[**FetchServiceConversation**](DefaultApi.md#FetchServiceConversation) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{Sid} | 
[**FetchServiceConversationMessage**](DefaultApi.md#FetchServiceConversationMessage) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{Sid} | 
[**FetchServiceConversationMessageReceipt**](DefaultApi.md#FetchServiceConversationMessageReceipt) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{MessageSid}/Receipts/{Sid} | 
[**FetchServiceConversationParticipant**](DefaultApi.md#FetchServiceConversationParticipant) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Participants/{Sid} | 
[**FetchServiceConversationScopedWebhook**](DefaultApi.md#FetchServiceConversationScopedWebhook) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks/{Sid} | 
[**FetchServiceNotification**](DefaultApi.md#FetchServiceNotification) | **Get** /v1/Services/{ChatServiceSid}/Configuration/Notifications | 
[**FetchServiceRole**](DefaultApi.md#FetchServiceRole) | **Get** /v1/Services/{ChatServiceSid}/Roles/{Sid} | 
[**FetchServiceUser**](DefaultApi.md#FetchServiceUser) | **Get** /v1/Services/{ChatServiceSid}/Users/{Sid} | 
[**FetchUser**](DefaultApi.md#FetchUser) | **Get** /v1/Users/{Sid} | 
[**ListConversation**](DefaultApi.md#ListConversation) | **Get** /v1/Conversations | 
[**ListConversationMessage**](DefaultApi.md#ListConversationMessage) | **Get** /v1/Conversations/{ConversationSid}/Messages | 
[**ListConversationMessageReceipt**](DefaultApi.md#ListConversationMessageReceipt) | **Get** /v1/Conversations/{ConversationSid}/Messages/{MessageSid}/Receipts | 
[**ListConversationParticipant**](DefaultApi.md#ListConversationParticipant) | **Get** /v1/Conversations/{ConversationSid}/Participants | 
[**ListConversationScopedWebhook**](DefaultApi.md#ListConversationScopedWebhook) | **Get** /v1/Conversations/{ConversationSid}/Webhooks | 
[**ListCredential**](DefaultApi.md#ListCredential) | **Get** /v1/Credentials | 
[**ListRole**](DefaultApi.md#ListRole) | **Get** /v1/Roles | 
[**ListService**](DefaultApi.md#ListService) | **Get** /v1/Services | 
[**ListServiceBinding**](DefaultApi.md#ListServiceBinding) | **Get** /v1/Services/{ChatServiceSid}/Bindings | 
[**ListServiceConversation**](DefaultApi.md#ListServiceConversation) | **Get** /v1/Services/{ChatServiceSid}/Conversations | 
[**ListServiceConversationMessage**](DefaultApi.md#ListServiceConversationMessage) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages | 
[**ListServiceConversationMessageReceipt**](DefaultApi.md#ListServiceConversationMessageReceipt) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{MessageSid}/Receipts | 
[**ListServiceConversationParticipant**](DefaultApi.md#ListServiceConversationParticipant) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Participants | 
[**ListServiceConversationScopedWebhook**](DefaultApi.md#ListServiceConversationScopedWebhook) | **Get** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks | 
[**ListServiceRole**](DefaultApi.md#ListServiceRole) | **Get** /v1/Services/{ChatServiceSid}/Roles | 
[**ListServiceUser**](DefaultApi.md#ListServiceUser) | **Get** /v1/Services/{ChatServiceSid}/Users | 
[**ListUser**](DefaultApi.md#ListUser) | **Get** /v1/Users | 
[**UpdateConfiguration**](DefaultApi.md#UpdateConfiguration) | **Post** /v1/Configuration | 
[**UpdateConfigurationWebhook**](DefaultApi.md#UpdateConfigurationWebhook) | **Post** /v1/Configuration/Webhooks | 
[**UpdateConversation**](DefaultApi.md#UpdateConversation) | **Post** /v1/Conversations/{Sid} | 
[**UpdateConversationMessage**](DefaultApi.md#UpdateConversationMessage) | **Post** /v1/Conversations/{ConversationSid}/Messages/{Sid} | 
[**UpdateConversationParticipant**](DefaultApi.md#UpdateConversationParticipant) | **Post** /v1/Conversations/{ConversationSid}/Participants/{Sid} | 
[**UpdateConversationScopedWebhook**](DefaultApi.md#UpdateConversationScopedWebhook) | **Post** /v1/Conversations/{ConversationSid}/Webhooks/{Sid} | 
[**UpdateCredential**](DefaultApi.md#UpdateCredential) | **Post** /v1/Credentials/{Sid} | 
[**UpdateRole**](DefaultApi.md#UpdateRole) | **Post** /v1/Roles/{Sid} | 
[**UpdateServiceConfiguration**](DefaultApi.md#UpdateServiceConfiguration) | **Post** /v1/Services/{ChatServiceSid}/Configuration | 
[**UpdateServiceConversation**](DefaultApi.md#UpdateServiceConversation) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{Sid} | 
[**UpdateServiceConversationMessage**](DefaultApi.md#UpdateServiceConversationMessage) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{Sid} | 
[**UpdateServiceConversationParticipant**](DefaultApi.md#UpdateServiceConversationParticipant) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Participants/{Sid} | 
[**UpdateServiceConversationScopedWebhook**](DefaultApi.md#UpdateServiceConversationScopedWebhook) | **Post** /v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Webhooks/{Sid} | 
[**UpdateServiceNotification**](DefaultApi.md#UpdateServiceNotification) | **Post** /v1/Services/{ChatServiceSid}/Configuration/Notifications | 
[**UpdateServiceRole**](DefaultApi.md#UpdateServiceRole) | **Post** /v1/Services/{ChatServiceSid}/Roles/{Sid} | 
[**UpdateServiceUser**](DefaultApi.md#UpdateServiceUser) | **Post** /v1/Services/{ChatServiceSid}/Users/{Sid} | 
[**UpdateUser**](DefaultApi.md#UpdateUser) | **Post** /v1/Users/{Sid} | 



## CreateConversation

> ConversationsV1Conversation CreateConversation(ctx).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).DateCreated(DateCreated).DateUpdated(DateUpdated).FriendlyName(FriendlyName).MessagingServiceSid(MessagingServiceSid).State(State).TimersClosed(TimersClosed).TimersInactive(TimersInactive).UniqueName(UniqueName).Execute()





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
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)
    Attributes := "Attributes_example" // string | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned. (optional)
    DateCreated := time.Now() // time.Time | The date that this resource was created. (optional)
    DateUpdated := time.Now() // time.Time | The date that this resource was last updated. (optional)
    FriendlyName := "FriendlyName_example" // string | The human-readable name of this conversation, limited to 256 characters. Optional. (optional)
    MessagingServiceSid := "MessagingServiceSid_example" // string | The unique ID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) this conversation belongs to. (optional)
    State := "State_example" // string | Current state of this conversation. Can be either `active`, `inactive` or `closed` and defaults to `active` (optional)
    TimersClosed := "TimersClosed_example" // string | ISO8601 duration when conversation will be switched to `closed` state. Minimum value for this timer is 10 minutes. (optional)
    TimersInactive := "TimersInactive_example" // string | ISO8601 duration when conversation will be switched to `inactive` state. Minimum value for this timer is 1 minute. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateConversation(context.Background()).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).DateCreated(DateCreated).DateUpdated(DateUpdated).FriendlyName(FriendlyName).MessagingServiceSid(MessagingServiceSid).State(State).TimersClosed(TimersClosed).TimersInactive(TimersInactive).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateConversation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConversation`: ConversationsV1Conversation
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateConversation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateConversationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
 **Attributes** | **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\&quot;{}\\\&quot; will be returned.
 **DateCreated** | **time.Time** | The date that this resource was created.
 **DateUpdated** | **time.Time** | The date that this resource was last updated.
 **FriendlyName** | **string** | The human-readable name of this conversation, limited to 256 characters. Optional.
 **MessagingServiceSid** | **string** | The unique ID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) this conversation belongs to.
 **State** | **string** | Current state of this conversation. Can be either &#x60;active&#x60;, &#x60;inactive&#x60; or &#x60;closed&#x60; and defaults to &#x60;active&#x60;
 **TimersClosed** | **string** | ISO8601 duration when conversation will be switched to &#x60;closed&#x60; state. Minimum value for this timer is 10 minutes.
 **TimersInactive** | **string** | ISO8601 duration when conversation will be switched to &#x60;inactive&#x60; state. Minimum value for this timer is 1 minute.
 **UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource&#39;s &#x60;sid&#x60; in the URL.

### Return type

[**ConversationsV1Conversation**](ConversationsV1Conversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConversationMessage

> ConversationsV1ConversationConversationMessage CreateConversationMessage(ctx, ConversationSid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).Author(Author).Body(Body).DateCreated(DateCreated).DateUpdated(DateUpdated).MediaSid(MediaSid).Execute()





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
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)
    Attributes := "Attributes_example" // string | A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned. (optional)
    Author := "Author_example" // string | The channel specific identifier of the message's author. Defaults to `system`. (optional)
    Body := "Body_example" // string | The content of the message, can be up to 1,600 characters long. (optional)
    DateCreated := time.Now() // time.Time | The date that this resource was created. (optional)
    DateUpdated := time.Now() // time.Time | The date that this resource was last updated. `null` if the message has not been edited. (optional)
    MediaSid := "MediaSid_example" // string | The Media SID to be attached to the new Message. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateConversationMessage(context.Background(), ConversationSid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).Author(Author).Body(Body).DateCreated(DateCreated).DateUpdated(DateUpdated).MediaSid(MediaSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateConversationMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConversationMessage`: ConversationsV1ConversationConversationMessage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateConversationMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.

### Other Parameters

Other parameters are passed through a pointer to a CreateConversationMessageParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
 **Attributes** | **string** | A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\&quot;{}\\\&quot; will be returned.
 **Author** | **string** | The channel specific identifier of the message&#39;s author. Defaults to &#x60;system&#x60;.
 **Body** | **string** | The content of the message, can be up to 1,600 characters long.
 **DateCreated** | **time.Time** | The date that this resource was created.
 **DateUpdated** | **time.Time** | The date that this resource was last updated. &#x60;null&#x60; if the message has not been edited.
 **MediaSid** | **string** | The Media SID to be attached to the new Message.

### Return type

[**ConversationsV1ConversationConversationMessage**](ConversationsV1ConversationConversationMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConversationParticipant

> ConversationsV1ConversationConversationParticipant CreateConversationParticipant(ctx, ConversationSid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).DateCreated(DateCreated).DateUpdated(DateUpdated).Identity(Identity).MessagingBindingAddress(MessagingBindingAddress).MessagingBindingProjectedAddress(MessagingBindingProjectedAddress).MessagingBindingProxyAddress(MessagingBindingProxyAddress).RoleSid(RoleSid).Execute()





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
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)
    Attributes := "Attributes_example" // string | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned. (optional)
    DateCreated := time.Now() // time.Time | The date that this resource was created. (optional)
    DateUpdated := time.Now() // time.Time | The date that this resource was last updated. (optional)
    Identity := "Identity_example" // string | A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversations SDK to communicate. Limited to 256 characters. (optional)
    MessagingBindingAddress := "MessagingBindingAddress_example" // string | The address of the participant's device, e.g. a phone or WhatsApp number. Together with the Proxy address, this determines a participant uniquely. This field (with proxy_address) is only null when the participant is interacting from an SDK endpoint (see the 'identity' field). (optional)
    MessagingBindingProjectedAddress := "MessagingBindingProjectedAddress_example" // string | The address of the Twilio phone number that is used in Group MMS. Communication mask for the Conversation participant with Identity. (optional)
    MessagingBindingProxyAddress := "MessagingBindingProxyAddress_example" // string | The address of the Twilio phone number (or WhatsApp number) that the participant is in contact with. This field, together with participant address, is only null when the participant is interacting from an SDK endpoint (see the 'identity' field). (optional)
    RoleSid := "RoleSid_example" // string | The SID of a conversation-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the participant. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateConversationParticipant(context.Background(), ConversationSid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).DateCreated(DateCreated).DateUpdated(DateUpdated).Identity(Identity).MessagingBindingAddress(MessagingBindingAddress).MessagingBindingProjectedAddress(MessagingBindingProjectedAddress).MessagingBindingProxyAddress(MessagingBindingProxyAddress).RoleSid(RoleSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateConversationParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConversationParticipant`: ConversationsV1ConversationConversationParticipant
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateConversationParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.

### Other Parameters

Other parameters are passed through a pointer to a CreateConversationParticipantParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
 **Attributes** | **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\&quot;{}\\\&quot; will be returned.
 **DateCreated** | **time.Time** | The date that this resource was created.
 **DateUpdated** | **time.Time** | The date that this resource was last updated.
 **Identity** | **string** | A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversations SDK to communicate. Limited to 256 characters.
 **MessagingBindingAddress** | **string** | The address of the participant&#39;s device, e.g. a phone or WhatsApp number. Together with the Proxy address, this determines a participant uniquely. This field (with proxy_address) is only null when the participant is interacting from an SDK endpoint (see the &#39;identity&#39; field).
 **MessagingBindingProjectedAddress** | **string** | The address of the Twilio phone number that is used in Group MMS. Communication mask for the Conversation participant with Identity.
 **MessagingBindingProxyAddress** | **string** | The address of the Twilio phone number (or WhatsApp number) that the participant is in contact with. This field, together with participant address, is only null when the participant is interacting from an SDK endpoint (see the &#39;identity&#39; field).
 **RoleSid** | **string** | The SID of a conversation-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the participant.

### Return type

[**ConversationsV1ConversationConversationParticipant**](ConversationsV1ConversationConversationParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConversationScopedWebhook

> ConversationsV1ConversationConversationScopedWebhook CreateConversationScopedWebhook(ctx, ConversationSid).ConfigurationFilters(ConfigurationFilters).ConfigurationFlowSid(ConfigurationFlowSid).ConfigurationMethod(ConfigurationMethod).ConfigurationReplayAfter(ConfigurationReplayAfter).ConfigurationTriggers(ConfigurationTriggers).ConfigurationUrl(ConfigurationUrl).Target(Target).Execute()





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
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
    ConfigurationFilters := []string{"Inner_example"} // []string | The list of events, firing webhook event for this Conversation. (optional)
    ConfigurationFlowSid := "ConfigurationFlowSid_example" // string | The studio flow SID, where the webhook should be sent to. (optional)
    ConfigurationMethod := "ConfigurationMethod_example" // string | The HTTP method to be used when sending a webhook request. (optional)
    ConfigurationReplayAfter := int32(56) // int32 | The message index for which and it's successors the webhook will be replayed. Not set by default (optional)
    ConfigurationTriggers := []string{"Inner_example"} // []string | The list of keywords, firing webhook event for this Conversation. (optional)
    ConfigurationUrl := "ConfigurationUrl_example" // string | The absolute url the webhook request should be sent to. (optional)
    Target := "Target_example" // string | The target of this webhook: `webhook`, `studio`, `trigger` (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateConversationScopedWebhook(context.Background(), ConversationSid).ConfigurationFilters(ConfigurationFilters).ConfigurationFlowSid(ConfigurationFlowSid).ConfigurationMethod(ConfigurationMethod).ConfigurationReplayAfter(ConfigurationReplayAfter).ConfigurationTriggers(ConfigurationTriggers).ConfigurationUrl(ConfigurationUrl).Target(Target).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateConversationScopedWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConversationScopedWebhook`: ConversationsV1ConversationConversationScopedWebhook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateConversationScopedWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.

### Other Parameters

Other parameters are passed through a pointer to a CreateConversationScopedWebhookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **ConfigurationFilters** | **[]string** | The list of events, firing webhook event for this Conversation.
 **ConfigurationFlowSid** | **string** | The studio flow SID, where the webhook should be sent to.
 **ConfigurationMethod** | **string** | The HTTP method to be used when sending a webhook request.
 **ConfigurationReplayAfter** | **int32** | The message index for which and it&#39;s successors the webhook will be replayed. Not set by default
 **ConfigurationTriggers** | **[]string** | The list of keywords, firing webhook event for this Conversation.
 **ConfigurationUrl** | **string** | The absolute url the webhook request should be sent to.
 **Target** | **string** | The target of this webhook: &#x60;webhook&#x60;, &#x60;studio&#x60;, &#x60;trigger&#x60;

### Return type

[**ConversationsV1ConversationConversationScopedWebhook**](ConversationsV1ConversationConversationScopedWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCredential

> ConversationsV1Credential CreateCredential(ctx).ApiKey(ApiKey).Certificate(Certificate).FriendlyName(FriendlyName).PrivateKey(PrivateKey).Sandbox(Sandbox).Secret(Secret).Type(Type).Execute()





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
    Certificate := "Certificate_example" // string | [APN only] The URL encoded representation of the certificate. For example,  `-----BEGIN CERTIFICATE----- MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEF.....A== -----END CERTIFICATE-----`. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the new resource. It can be up to 64 characters long. (optional)
    PrivateKey := "PrivateKey_example" // string | [APN only] The URL encoded representation of the private key. For example, `-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fG... -----END RSA PRIVATE KEY-----`. (optional)
    Sandbox := true // bool | [APN only] Whether to send the credential to sandbox APNs. Can be `true` to send to sandbox APNs or `false` to send to production. (optional)
    Secret := "Secret_example" // string | [FCM only] The **Server key** of your project from the Firebase console, found under Settings / Cloud messaging. (optional)
    Type := "Type_example" // string | The type of push-notification service the credential is for. Can be: `fcm`, `gcm`, or `apn`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateCredential(context.Background()).ApiKey(ApiKey).Certificate(Certificate).FriendlyName(FriendlyName).PrivateKey(PrivateKey).Sandbox(Sandbox).Secret(Secret).Type(Type).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCredential`: ConversationsV1Credential
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
 **Certificate** | **string** | [APN only] The URL encoded representation of the certificate. For example,  &#x60;-----BEGIN CERTIFICATE----- MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEF.....A&#x3D;&#x3D; -----END CERTIFICATE-----&#x60;.
 **FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
 **PrivateKey** | **string** | [APN only] The URL encoded representation of the private key. For example, &#x60;-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fG... -----END RSA PRIVATE KEY-----&#x60;.
 **Sandbox** | **bool** | [APN only] Whether to send the credential to sandbox APNs. Can be &#x60;true&#x60; to send to sandbox APNs or &#x60;false&#x60; to send to production.
 **Secret** | **string** | [FCM only] The **Server key** of your project from the Firebase console, found under Settings / Cloud messaging.
 **Type** | **string** | The type of push-notification service the credential is for. Can be: &#x60;fcm&#x60;, &#x60;gcm&#x60;, or &#x60;apn&#x60;.

### Return type

[**ConversationsV1Credential**](ConversationsV1Credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> ConversationsV1Role CreateRole(ctx).FriendlyName(FriendlyName).Permission(Permission).Type(Type).Execute()





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
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the new resource. It can be up to 64 characters long. (optional)
    Permission := []string{"Inner_example"} // []string | A permission that you grant to the new role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. The values for this parameter depend on the role's `type`. (optional)
    Type := "Type_example" // string | The type of role. Can be: `conversation` for [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) roles or `service` for [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) roles. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateRole(context.Background()).FriendlyName(FriendlyName).Permission(Permission).Type(Type).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRole`: ConversationsV1Role
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateRole`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateRoleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
 **Permission** | **[]string** | A permission that you grant to the new role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. The values for this parameter depend on the role&#39;s &#x60;type&#x60;.
 **Type** | **string** | The type of role. Can be: &#x60;conversation&#x60; for [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) roles or &#x60;service&#x60; for [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) roles.

### Return type

[**ConversationsV1Role**](ConversationsV1Role.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateService

> ConversationsV1Service CreateService(ctx).FriendlyName(FriendlyName).Execute()





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
    FriendlyName := "FriendlyName_example" // string | The human-readable name of this service, limited to 256 characters. Optional. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateService(context.Background()).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateService`: ConversationsV1Service
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateService`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **FriendlyName** | **string** | The human-readable name of this service, limited to 256 characters. Optional.

### Return type

[**ConversationsV1Service**](ConversationsV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServiceConversation

> ConversationsV1ServiceServiceConversation CreateServiceConversation(ctx, ChatServiceSid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).DateCreated(DateCreated).DateUpdated(DateUpdated).FriendlyName(FriendlyName).MessagingServiceSid(MessagingServiceSid).State(State).TimersClosed(TimersClosed).TimersInactive(TimersInactive).UniqueName(UniqueName).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)
    Attributes := "Attributes_example" // string | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned. (optional)
    DateCreated := time.Now() // time.Time | The date that this resource was created. (optional)
    DateUpdated := time.Now() // time.Time | The date that this resource was last updated. (optional)
    FriendlyName := "FriendlyName_example" // string | The human-readable name of this conversation, limited to 256 characters. Optional. (optional)
    MessagingServiceSid := "MessagingServiceSid_example" // string | The unique ID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) this conversation belongs to. (optional)
    State := "State_example" // string | Current state of this conversation. Can be either `active`, `inactive` or `closed` and defaults to `active` (optional)
    TimersClosed := "TimersClosed_example" // string | ISO8601 duration when conversation will be switched to `closed` state. Minimum value for this timer is 10 minutes. (optional)
    TimersInactive := "TimersInactive_example" // string | ISO8601 duration when conversation will be switched to `inactive` state. Minimum value for this timer is 1 minute. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateServiceConversation(context.Background(), ChatServiceSid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).DateCreated(DateCreated).DateUpdated(DateUpdated).FriendlyName(FriendlyName).MessagingServiceSid(MessagingServiceSid).State(State).TimersClosed(TimersClosed).TimersInactive(TimersInactive).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateServiceConversation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceConversation`: ConversationsV1ServiceServiceConversation
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateServiceConversation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceConversationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
 **Attributes** | **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\&quot;{}\\\&quot; will be returned.
 **DateCreated** | **time.Time** | The date that this resource was created.
 **DateUpdated** | **time.Time** | The date that this resource was last updated.
 **FriendlyName** | **string** | The human-readable name of this conversation, limited to 256 characters. Optional.
 **MessagingServiceSid** | **string** | The unique ID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) this conversation belongs to.
 **State** | **string** | Current state of this conversation. Can be either &#x60;active&#x60;, &#x60;inactive&#x60; or &#x60;closed&#x60; and defaults to &#x60;active&#x60;
 **TimersClosed** | **string** | ISO8601 duration when conversation will be switched to &#x60;closed&#x60; state. Minimum value for this timer is 10 minutes.
 **TimersInactive** | **string** | ISO8601 duration when conversation will be switched to &#x60;inactive&#x60; state. Minimum value for this timer is 1 minute.
 **UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource&#39;s &#x60;sid&#x60; in the URL.

### Return type

[**ConversationsV1ServiceServiceConversation**](ConversationsV1ServiceServiceConversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServiceConversationMessage

> ConversationsV1ServiceServiceConversationServiceConversationMessage CreateServiceConversationMessage(ctx, ChatServiceSid, ConversationSid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).Author(Author).Body(Body).DateCreated(DateCreated).DateUpdated(DateUpdated).MediaSid(MediaSid).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)
    Attributes := "Attributes_example" // string | A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned. (optional)
    Author := "Author_example" // string | The channel specific identifier of the message's author. Defaults to `system`. (optional)
    Body := "Body_example" // string | The content of the message, can be up to 1,600 characters long. (optional)
    DateCreated := time.Now() // time.Time | The date that this resource was created. (optional)
    DateUpdated := time.Now() // time.Time | The date that this resource was last updated. `null` if the message has not been edited. (optional)
    MediaSid := "MediaSid_example" // string | The Media SID to be attached to the new Message. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateServiceConversationMessage(context.Background(), ChatServiceSid, ConversationSid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).Author(Author).Body(Body).DateCreated(DateCreated).DateUpdated(DateUpdated).MediaSid(MediaSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateServiceConversationMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceConversationMessage`: ConversationsV1ServiceServiceConversationServiceConversationMessage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateServiceConversationMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceConversationMessageParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
 **Attributes** | **string** | A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\&quot;{}\\\&quot; will be returned.
 **Author** | **string** | The channel specific identifier of the message&#39;s author. Defaults to &#x60;system&#x60;.
 **Body** | **string** | The content of the message, can be up to 1,600 characters long.
 **DateCreated** | **time.Time** | The date that this resource was created.
 **DateUpdated** | **time.Time** | The date that this resource was last updated. &#x60;null&#x60; if the message has not been edited.
 **MediaSid** | **string** | The Media SID to be attached to the new Message.

### Return type

[**ConversationsV1ServiceServiceConversationServiceConversationMessage**](ConversationsV1ServiceServiceConversationServiceConversationMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServiceConversationParticipant

> ConversationsV1ServiceServiceConversationServiceConversationParticipant CreateServiceConversationParticipant(ctx, ChatServiceSid, ConversationSid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).DateCreated(DateCreated).DateUpdated(DateUpdated).Identity(Identity).MessagingBindingAddress(MessagingBindingAddress).MessagingBindingProjectedAddress(MessagingBindingProjectedAddress).MessagingBindingProxyAddress(MessagingBindingProxyAddress).RoleSid(RoleSid).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)
    Attributes := "Attributes_example" // string | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned. (optional)
    DateCreated := time.Now() // time.Time | The date that this resource was created. (optional)
    DateUpdated := time.Now() // time.Time | The date that this resource was last updated. (optional)
    Identity := "Identity_example" // string | A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversation SDK to communicate. Limited to 256 characters. (optional)
    MessagingBindingAddress := "MessagingBindingAddress_example" // string | The address of the participant's device, e.g. a phone or WhatsApp number. Together with the Proxy address, this determines a participant uniquely. This field (with proxy_address) is only null when the participant is interacting from an SDK endpoint (see the 'identity' field). (optional)
    MessagingBindingProjectedAddress := "MessagingBindingProjectedAddress_example" // string | The address of the Twilio phone number that is used in Group MMS. Communication mask for the Conversation participant with Identity. (optional)
    MessagingBindingProxyAddress := "MessagingBindingProxyAddress_example" // string | The address of the Twilio phone number (or WhatsApp number) that the participant is in contact with. This field, together with participant address, is only null when the participant is interacting from an SDK endpoint (see the 'identity' field). (optional)
    RoleSid := "RoleSid_example" // string | The SID of a conversation-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the participant. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateServiceConversationParticipant(context.Background(), ChatServiceSid, ConversationSid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).DateCreated(DateCreated).DateUpdated(DateUpdated).Identity(Identity).MessagingBindingAddress(MessagingBindingAddress).MessagingBindingProjectedAddress(MessagingBindingProjectedAddress).MessagingBindingProxyAddress(MessagingBindingProxyAddress).RoleSid(RoleSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateServiceConversationParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceConversationParticipant`: ConversationsV1ServiceServiceConversationServiceConversationParticipant
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateServiceConversationParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceConversationParticipantParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
 **Attributes** | **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\&quot;{}\\\&quot; will be returned.
 **DateCreated** | **time.Time** | The date that this resource was created.
 **DateUpdated** | **time.Time** | The date that this resource was last updated.
 **Identity** | **string** | A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversation SDK to communicate. Limited to 256 characters.
 **MessagingBindingAddress** | **string** | The address of the participant&#39;s device, e.g. a phone or WhatsApp number. Together with the Proxy address, this determines a participant uniquely. This field (with proxy_address) is only null when the participant is interacting from an SDK endpoint (see the &#39;identity&#39; field).
 **MessagingBindingProjectedAddress** | **string** | The address of the Twilio phone number that is used in Group MMS. Communication mask for the Conversation participant with Identity.
 **MessagingBindingProxyAddress** | **string** | The address of the Twilio phone number (or WhatsApp number) that the participant is in contact with. This field, together with participant address, is only null when the participant is interacting from an SDK endpoint (see the &#39;identity&#39; field).
 **RoleSid** | **string** | The SID of a conversation-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the participant.

### Return type

[**ConversationsV1ServiceServiceConversationServiceConversationParticipant**](ConversationsV1ServiceServiceConversationServiceConversationParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServiceConversationScopedWebhook

> ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook CreateServiceConversationScopedWebhook(ctx, ChatServiceSid, ConversationSid).ConfigurationFilters(ConfigurationFilters).ConfigurationFlowSid(ConfigurationFlowSid).ConfigurationMethod(ConfigurationMethod).ConfigurationReplayAfter(ConfigurationReplayAfter).ConfigurationTriggers(ConfigurationTriggers).ConfigurationUrl(ConfigurationUrl).Target(Target).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
    ConfigurationFilters := []string{"Inner_example"} // []string | The list of events, firing webhook event for this Conversation. (optional)
    ConfigurationFlowSid := "ConfigurationFlowSid_example" // string | The studio flow SID, where the webhook should be sent to. (optional)
    ConfigurationMethod := "ConfigurationMethod_example" // string | The HTTP method to be used when sending a webhook request. (optional)
    ConfigurationReplayAfter := int32(56) // int32 | The message index for which and it's successors the webhook will be replayed. Not set by default (optional)
    ConfigurationTriggers := []string{"Inner_example"} // []string | The list of keywords, firing webhook event for this Conversation. (optional)
    ConfigurationUrl := "ConfigurationUrl_example" // string | The absolute url the webhook request should be sent to. (optional)
    Target := "Target_example" // string | The target of this webhook: `webhook`, `studio`, `trigger` (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateServiceConversationScopedWebhook(context.Background(), ChatServiceSid, ConversationSid).ConfigurationFilters(ConfigurationFilters).ConfigurationFlowSid(ConfigurationFlowSid).ConfigurationMethod(ConfigurationMethod).ConfigurationReplayAfter(ConfigurationReplayAfter).ConfigurationTriggers(ConfigurationTriggers).ConfigurationUrl(ConfigurationUrl).Target(Target).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateServiceConversationScopedWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceConversationScopedWebhook`: ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateServiceConversationScopedWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceConversationScopedWebhookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **ConfigurationFilters** | **[]string** | The list of events, firing webhook event for this Conversation.
 **ConfigurationFlowSid** | **string** | The studio flow SID, where the webhook should be sent to.
 **ConfigurationMethod** | **string** | The HTTP method to be used when sending a webhook request.
 **ConfigurationReplayAfter** | **int32** | The message index for which and it&#39;s successors the webhook will be replayed. Not set by default
 **ConfigurationTriggers** | **[]string** | The list of keywords, firing webhook event for this Conversation.
 **ConfigurationUrl** | **string** | The absolute url the webhook request should be sent to.
 **Target** | **string** | The target of this webhook: &#x60;webhook&#x60;, &#x60;studio&#x60;, &#x60;trigger&#x60;

### Return type

[**ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook**](ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServiceRole

> ConversationsV1ServiceServiceRole CreateServiceRole(ctx, ChatServiceSid).FriendlyName(FriendlyName).Permission(Permission).Type(Type).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to create the Role resource under.
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the new resource. It can be up to 64 characters long. (optional)
    Permission := []string{"Inner_example"} // []string | A permission that you grant to the new role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. The values for this parameter depend on the role's `type`. (optional)
    Type := "Type_example" // string | The type of role. Can be: `conversation` for [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) roles or `service` for [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) roles. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateServiceRole(context.Background(), ChatServiceSid).FriendlyName(FriendlyName).Permission(Permission).Type(Type).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateServiceRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceRole`: ConversationsV1ServiceServiceRole
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateServiceRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to create the Role resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceRoleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
 **Permission** | **[]string** | A permission that you grant to the new role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. The values for this parameter depend on the role&#39;s &#x60;type&#x60;.
 **Type** | **string** | The type of role. Can be: &#x60;conversation&#x60; for [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) roles or &#x60;service&#x60; for [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) roles.

### Return type

[**ConversationsV1ServiceServiceRole**](ConversationsV1ServiceServiceRole.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServiceUser

> ConversationsV1ServiceServiceUser CreateServiceUser(ctx, ChatServiceSid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).FriendlyName(FriendlyName).Identity(Identity).RoleSid(RoleSid).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the User resource is associated with.
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)
    Attributes := "Attributes_example" // string | The JSON Object string that stores application-specific data. If attributes have not been set, `{}` is returned. (optional)
    FriendlyName := "FriendlyName_example" // string | The string that you assigned to describe the resource. (optional)
    Identity := "Identity_example" // string | The application-defined string that uniquely identifies the resource's User within the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource). This value is often a username or an email address, and is case-sensitive. (optional)
    RoleSid := "RoleSid_example" // string | The SID of a service-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateServiceUser(context.Background(), ChatServiceSid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).FriendlyName(FriendlyName).Identity(Identity).RoleSid(RoleSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateServiceUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceUser`: ConversationsV1ServiceServiceUser
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateServiceUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the User resource is associated with.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceUserParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
 **Attributes** | **string** | The JSON Object string that stores application-specific data. If attributes have not been set, &#x60;{}&#x60; is returned.
 **FriendlyName** | **string** | The string that you assigned to describe the resource.
 **Identity** | **string** | The application-defined string that uniquely identifies the resource&#39;s User within the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource). This value is often a username or an email address, and is case-sensitive.
 **RoleSid** | **string** | The SID of a service-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the user.

### Return type

[**ConversationsV1ServiceServiceUser**](ConversationsV1ServiceServiceUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> ConversationsV1User CreateUser(ctx).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).FriendlyName(FriendlyName).Identity(Identity).RoleSid(RoleSid).Execute()





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
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)
    Attributes := "Attributes_example" // string | The JSON Object string that stores application-specific data. If attributes have not been set, `{}` is returned. (optional)
    FriendlyName := "FriendlyName_example" // string | The string that you assigned to describe the resource. (optional)
    Identity := "Identity_example" // string | The application-defined string that uniquely identifies the resource's User within the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource). This value is often a username or an email address, and is case-sensitive. (optional)
    RoleSid := "RoleSid_example" // string | The SID of a service-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateUser(context.Background()).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).FriendlyName(FriendlyName).Identity(Identity).RoleSid(RoleSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: ConversationsV1User
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateUserParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
 **Attributes** | **string** | The JSON Object string that stores application-specific data. If attributes have not been set, &#x60;{}&#x60; is returned.
 **FriendlyName** | **string** | The string that you assigned to describe the resource.
 **Identity** | **string** | The application-defined string that uniquely identifies the resource&#39;s User within the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource). This value is often a username or an email address, and is case-sensitive.
 **RoleSid** | **string** | The SID of a service-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the user.

### Return type

[**ConversationsV1User**](ConversationsV1User.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConversation

> DeleteConversation(ctx, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Execute()





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
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource. Can also be the `unique_name` of the Conversation.
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteConversation(context.Background(), Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteConversation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this resource. Can also be the &#x60;unique_name&#x60; of the Conversation.

### Other Parameters

Other parameters are passed through a pointer to a DeleteConversationParams struct via the builder pattern


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


## DeleteConversationMessage

> DeleteConversationMessage(ctx, ConversationSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Execute()





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
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteConversationMessage(context.Background(), ConversationSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteConversationMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteConversationMessageParams struct via the builder pattern


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


## DeleteConversationParticipant

> DeleteConversationParticipant(ctx, ConversationSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Execute()





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
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteConversationParticipant(context.Background(), ConversationSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteConversationParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteConversationParticipantParams struct via the builder pattern


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


## DeleteConversationScopedWebhook

> DeleteConversationScopedWebhook(ctx, ConversationSid, Sid).Execute()





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
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteConversationScopedWebhook(context.Background(), ConversationSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteConversationScopedWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteConversationScopedWebhookParams struct via the builder pattern


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
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.

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
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

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


## DeleteRole

> DeleteRole(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | The SID of the Role resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteRole(context.Background(), Sid).Execute()
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
**Sid** | **string** | The SID of the Role resource to delete.

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
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.

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
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

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


## DeleteServiceBinding

> DeleteServiceBinding(ctx, ChatServiceSid, Sid).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to delete the Binding resource from.
    Sid := "Sid_example" // string | The SID of the Binding resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteServiceBinding(context.Background(), ChatServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteServiceBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to delete the Binding resource from.
**Sid** | **string** | The SID of the Binding resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceBindingParams struct via the builder pattern


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


## DeleteServiceConversation

> DeleteServiceConversation(ctx, ChatServiceSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource. Can also be the `unique_name` of the Conversation.
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteServiceConversation(context.Background(), ChatServiceSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteServiceConversation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.
**Sid** | **string** | A 34 character string that uniquely identifies this resource. Can also be the &#x60;unique_name&#x60; of the Conversation.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceConversationParams struct via the builder pattern


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


## DeleteServiceConversationMessage

> DeleteServiceConversationMessage(ctx, ChatServiceSid, ConversationSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteServiceConversationMessage(context.Background(), ChatServiceSid, ConversationSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteServiceConversationMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceConversationMessageParams struct via the builder pattern


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


## DeleteServiceConversationParticipant

> DeleteServiceConversationParticipant(ctx, ChatServiceSid, ConversationSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteServiceConversationParticipant(context.Background(), ChatServiceSid, ConversationSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteServiceConversationParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceConversationParticipantParams struct via the builder pattern


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


## DeleteServiceConversationScopedWebhook

> DeleteServiceConversationScopedWebhook(ctx, ChatServiceSid, ConversationSid, Sid).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteServiceConversationScopedWebhook(context.Background(), ChatServiceSid, ConversationSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteServiceConversationScopedWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceConversationScopedWebhookParams struct via the builder pattern


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


## DeleteServiceRole

> DeleteServiceRole(ctx, ChatServiceSid, Sid).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to delete the Role resource from.
    Sid := "Sid_example" // string | The SID of the Role resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteServiceRole(context.Background(), ChatServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteServiceRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to delete the Role resource from.
**Sid** | **string** | The SID of the Role resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceRoleParams struct via the builder pattern


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


## DeleteServiceUser

> DeleteServiceUser(ctx, ChatServiceSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to delete the User resource from.
    Sid := "Sid_example" // string | The SID of the User resource to delete. This value can be either the `sid` or the `identity` of the User resource to delete.
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteServiceUser(context.Background(), ChatServiceSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteServiceUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to delete the User resource from.
**Sid** | **string** | The SID of the User resource to delete. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceUserParams struct via the builder pattern


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


## DeleteUser

> DeleteUser(ctx, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Execute()





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
    Sid := "Sid_example" // string | The SID of the User resource to delete. This value can be either the `sid` or the `identity` of the User resource to delete.
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteUser(context.Background(), Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Execute()
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
**Sid** | **string** | The SID of the User resource to delete. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteUserParams struct via the builder pattern


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


## FetchConfiguration

> ConversationsV1Configuration FetchConfiguration(ctx).Execute()





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
    resp, r, err := api_client.DefaultApi.FetchConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchConfiguration`: ConversationsV1Configuration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchConfigurationParams struct via the builder pattern


### Return type

[**ConversationsV1Configuration**](ConversationsV1Configuration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConfigurationWebhook

> ConversationsV1ConfigurationConfigurationWebhook FetchConfigurationWebhook(ctx).Execute()



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
    resp, r, err := api_client.DefaultApi.FetchConfigurationWebhook(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchConfigurationWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchConfigurationWebhook`: ConversationsV1ConfigurationConfigurationWebhook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchConfigurationWebhook`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchConfigurationWebhookParams struct via the builder pattern


### Return type

[**ConversationsV1ConfigurationConfigurationWebhook**](ConversationsV1ConfigurationConfigurationWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConversation

> ConversationsV1Conversation FetchConversation(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource. Can also be the `unique_name` of the Conversation.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchConversation(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchConversation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchConversation`: ConversationsV1Conversation
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchConversation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this resource. Can also be the &#x60;unique_name&#x60; of the Conversation.

### Other Parameters

Other parameters are passed through a pointer to a FetchConversationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**ConversationsV1Conversation**](ConversationsV1Conversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConversationMessage

> ConversationsV1ConversationConversationMessage FetchConversationMessage(ctx, ConversationSid, Sid).Execute()





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
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchConversationMessage(context.Background(), ConversationSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchConversationMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchConversationMessage`: ConversationsV1ConversationConversationMessage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchConversationMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchConversationMessageParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ConversationsV1ConversationConversationMessage**](ConversationsV1ConversationConversationMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConversationMessageReceipt

> ConversationsV1ConversationConversationMessageConversationMessageReceipt FetchConversationMessageReceipt(ctx, ConversationSid, MessageSid, Sid).Execute()





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
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
    MessageSid := "MessageSid_example" // string | The SID of the message within a [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) the delivery receipt belongs to.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchConversationMessageReceipt(context.Background(), ConversationSid, MessageSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchConversationMessageReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchConversationMessageReceipt`: ConversationsV1ConversationConversationMessageConversationMessageReceipt
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchConversationMessageReceipt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**MessageSid** | **string** | The SID of the message within a [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) the delivery receipt belongs to.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchConversationMessageReceiptParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ConversationsV1ConversationConversationMessageConversationMessageReceipt**](ConversationsV1ConversationConversationMessageConversationMessageReceipt.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConversationParticipant

> ConversationsV1ConversationConversationParticipant FetchConversationParticipant(ctx, ConversationSid, Sid).Execute()





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
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchConversationParticipant(context.Background(), ConversationSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchConversationParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchConversationParticipant`: ConversationsV1ConversationConversationParticipant
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchConversationParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchConversationParticipantParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ConversationsV1ConversationConversationParticipant**](ConversationsV1ConversationConversationParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConversationScopedWebhook

> ConversationsV1ConversationConversationScopedWebhook FetchConversationScopedWebhook(ctx, ConversationSid, Sid).Execute()





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
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchConversationScopedWebhook(context.Background(), ConversationSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchConversationScopedWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchConversationScopedWebhook`: ConversationsV1ConversationConversationScopedWebhook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchConversationScopedWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchConversationScopedWebhookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ConversationsV1ConversationConversationScopedWebhook**](ConversationsV1ConversationConversationScopedWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCredential

> ConversationsV1Credential FetchCredential(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchCredential(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCredential`: ConversationsV1Credential
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchCredentialParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**ConversationsV1Credential**](ConversationsV1Credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRole

> ConversationsV1Role FetchRole(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | The SID of the Role resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchRole(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRole`: ConversationsV1Role
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Role resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRoleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**ConversationsV1Role**](ConversationsV1Role.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> ConversationsV1Service FetchService(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchService(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchService`: ConversationsV1Service
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**ConversationsV1Service**](ConversationsV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchServiceBinding

> ConversationsV1ServiceServiceBinding FetchServiceBinding(ctx, ChatServiceSid, Sid).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Binding resource is associated with.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchServiceBinding(context.Background(), ChatServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchServiceBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchServiceBinding`: ConversationsV1ServiceServiceBinding
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchServiceBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Binding resource is associated with.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceBindingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ConversationsV1ServiceServiceBinding**](ConversationsV1ServiceServiceBinding.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchServiceConfiguration

> ConversationsV1ServiceServiceConfiguration FetchServiceConfiguration(ctx, ChatServiceSid).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the Service configuration resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchServiceConfiguration(context.Background(), ChatServiceSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchServiceConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchServiceConfiguration`: ConversationsV1ServiceServiceConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchServiceConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the Service configuration resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceConfigurationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**ConversationsV1ServiceServiceConfiguration**](ConversationsV1ServiceServiceConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchServiceConversation

> ConversationsV1ServiceServiceConversation FetchServiceConversation(ctx, ChatServiceSid, Sid).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource. Can also be the `unique_name` of the Conversation.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchServiceConversation(context.Background(), ChatServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchServiceConversation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchServiceConversation`: ConversationsV1ServiceServiceConversation
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchServiceConversation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.
**Sid** | **string** | A 34 character string that uniquely identifies this resource. Can also be the &#x60;unique_name&#x60; of the Conversation.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceConversationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ConversationsV1ServiceServiceConversation**](ConversationsV1ServiceServiceConversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchServiceConversationMessage

> ConversationsV1ServiceServiceConversationServiceConversationMessage FetchServiceConversationMessage(ctx, ChatServiceSid, ConversationSid, Sid).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchServiceConversationMessage(context.Background(), ChatServiceSid, ConversationSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchServiceConversationMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchServiceConversationMessage`: ConversationsV1ServiceServiceConversationServiceConversationMessage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchServiceConversationMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceConversationMessageParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ConversationsV1ServiceServiceConversationServiceConversationMessage**](ConversationsV1ServiceServiceConversationServiceConversationMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchServiceConversationMessageReceipt

> ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt FetchServiceConversationMessageReceipt(ctx, ChatServiceSid, ConversationSid, MessageSid, Sid).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Message resource is associated with.
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
    MessageSid := "MessageSid_example" // string | The SID of the message within a [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) the delivery receipt belongs to.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchServiceConversationMessageReceipt(context.Background(), ChatServiceSid, ConversationSid, MessageSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchServiceConversationMessageReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchServiceConversationMessageReceipt`: ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchServiceConversationMessageReceipt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Message resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**MessageSid** | **string** | The SID of the message within a [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) the delivery receipt belongs to.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceConversationMessageReceiptParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------





### Return type

[**ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt**](ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchServiceConversationParticipant

> ConversationsV1ServiceServiceConversationServiceConversationParticipant FetchServiceConversationParticipant(ctx, ChatServiceSid, ConversationSid, Sid).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchServiceConversationParticipant(context.Background(), ChatServiceSid, ConversationSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchServiceConversationParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchServiceConversationParticipant`: ConversationsV1ServiceServiceConversationServiceConversationParticipant
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchServiceConversationParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceConversationParticipantParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ConversationsV1ServiceServiceConversationServiceConversationParticipant**](ConversationsV1ServiceServiceConversationServiceConversationParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchServiceConversationScopedWebhook

> ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook FetchServiceConversationScopedWebhook(ctx, ChatServiceSid, ConversationSid, Sid).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchServiceConversationScopedWebhook(context.Background(), ChatServiceSid, ConversationSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchServiceConversationScopedWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchServiceConversationScopedWebhook`: ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchServiceConversationScopedWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceConversationScopedWebhookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook**](ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchServiceNotification

> ConversationsV1ServiceServiceConfigurationServiceNotification FetchServiceNotification(ctx, ChatServiceSid).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Configuration applies to.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchServiceNotification(context.Background(), ChatServiceSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchServiceNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchServiceNotification`: ConversationsV1ServiceServiceConfigurationServiceNotification
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchServiceNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Configuration applies to.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceNotificationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**ConversationsV1ServiceServiceConfigurationServiceNotification**](ConversationsV1ServiceServiceConfigurationServiceNotification.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchServiceRole

> ConversationsV1ServiceServiceRole FetchServiceRole(ctx, ChatServiceSid, Sid).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to fetch the Role resource from.
    Sid := "Sid_example" // string | The SID of the Role resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchServiceRole(context.Background(), ChatServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchServiceRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchServiceRole`: ConversationsV1ServiceServiceRole
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchServiceRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to fetch the Role resource from.
**Sid** | **string** | The SID of the Role resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceRoleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ConversationsV1ServiceServiceRole**](ConversationsV1ServiceServiceRole.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchServiceUser

> ConversationsV1ServiceServiceUser FetchServiceUser(ctx, ChatServiceSid, Sid).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to fetch the User resource from.
    Sid := "Sid_example" // string | The SID of the User resource to fetch. This value can be either the `sid` or the `identity` of the User resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchServiceUser(context.Background(), ChatServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchServiceUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchServiceUser`: ConversationsV1ServiceServiceUser
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchServiceUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to fetch the User resource from.
**Sid** | **string** | The SID of the User resource to fetch. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceUserParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ConversationsV1ServiceServiceUser**](ConversationsV1ServiceServiceUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUser

> ConversationsV1User FetchUser(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | The SID of the User resource to fetch. This value can be either the `sid` or the `identity` of the User resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchUser(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchUser`: ConversationsV1User
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the User resource to fetch. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchUserParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**ConversationsV1User**](ConversationsV1User.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConversation

> ListConversationResponse ListConversation(ctx).PageSize(PageSize).Execute()





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
    resp, r, err := api_client.DefaultApi.ListConversation(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListConversation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConversation`: ListConversationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListConversation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListConversationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListConversationResponse**](ListConversationResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConversationMessage

> ListConversationMessageResponse ListConversationMessage(ctx, ConversationSid).PageSize(PageSize).Execute()





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
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for messages.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListConversationMessage(context.Background(), ConversationSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListConversationMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConversationMessage`: ListConversationMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListConversationMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for messages.

### Other Parameters

Other parameters are passed through a pointer to a ListConversationMessageParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListConversationMessageResponse**](ListConversationMessageResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConversationMessageReceipt

> ListConversationMessageReceiptResponse ListConversationMessageReceipt(ctx, ConversationSid, MessageSid).PageSize(PageSize).Execute()





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
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
    MessageSid := "MessageSid_example" // string | The SID of the message within a [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) the delivery receipt belongs to.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListConversationMessageReceipt(context.Background(), ConversationSid, MessageSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListConversationMessageReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConversationMessageReceipt`: ListConversationMessageReceiptResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListConversationMessageReceipt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**MessageSid** | **string** | The SID of the message within a [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) the delivery receipt belongs to.

### Other Parameters

Other parameters are passed through a pointer to a ListConversationMessageReceiptParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListConversationMessageReceiptResponse**](ListConversationMessageReceiptResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConversationParticipant

> ListConversationParticipantResponse ListConversationParticipant(ctx, ConversationSid).PageSize(PageSize).Execute()





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
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for participants.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListConversationParticipant(context.Background(), ConversationSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListConversationParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConversationParticipant`: ListConversationParticipantResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListConversationParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for participants.

### Other Parameters

Other parameters are passed through a pointer to a ListConversationParticipantParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListConversationParticipantResponse**](ListConversationParticipantResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConversationScopedWebhook

> ListConversationScopedWebhookResponse ListConversationScopedWebhook(ctx, ConversationSid).PageSize(PageSize).Execute()





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
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListConversationScopedWebhook(context.Background(), ConversationSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListConversationScopedWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConversationScopedWebhook`: ListConversationScopedWebhookResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListConversationScopedWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.

### Other Parameters

Other parameters are passed through a pointer to a ListConversationScopedWebhookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListConversationScopedWebhookResponse**](ListConversationScopedWebhookResponse.md)

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


## ListRole

> ListRoleResponse ListRole(ctx).PageSize(PageSize).Execute()





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
    resp, r, err := api_client.DefaultApi.ListRole(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRole`: ListRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRole`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

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


## ListServiceBinding

> ListServiceBindingResponse ListServiceBinding(ctx, ChatServiceSid).BindingType(BindingType).Identity(Identity).PageSize(PageSize).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Binding resource is associated with.
    BindingType := []string{"BindingType_example"} // []string | The push technology used by the Binding resources to read.  Can be: `apn`, `gcm`, or `fcm`.  See [push notification configuration](https://www.twilio.com/docs/chat/push-notification-configuration) for more info. (optional)
    Identity := []string{"Inner_example"} // []string | The identity of a [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource) this binding belongs to. See [access tokens](https://www.twilio.com/docs/conversations/create-tokens) for more details. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListServiceBinding(context.Background(), ChatServiceSid).BindingType(BindingType).Identity(Identity).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListServiceBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceBinding`: ListServiceBindingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListServiceBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Binding resource is associated with.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceBindingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **BindingType** | **[]string** | The push technology used by the Binding resources to read.  Can be: &#x60;apn&#x60;, &#x60;gcm&#x60;, or &#x60;fcm&#x60;.  See [push notification configuration](https://www.twilio.com/docs/chat/push-notification-configuration) for more info.
 **Identity** | **[]string** | The identity of a [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource) this binding belongs to. See [access tokens](https://www.twilio.com/docs/conversations/create-tokens) for more details.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListServiceBindingResponse**](ListServiceBindingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceConversation

> ListServiceConversationResponse ListServiceConversation(ctx, ChatServiceSid).PageSize(PageSize).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListServiceConversation(context.Background(), ChatServiceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListServiceConversation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceConversation`: ListServiceConversationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListServiceConversation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceConversationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListServiceConversationResponse**](ListServiceConversationResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceConversationMessage

> ListServiceConversationMessageResponse ListServiceConversationMessage(ctx, ChatServiceSid, ConversationSid).PageSize(PageSize).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for messages.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListServiceConversationMessage(context.Background(), ChatServiceSid, ConversationSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListServiceConversationMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceConversationMessage`: ListServiceConversationMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListServiceConversationMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for messages.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceConversationMessageParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListServiceConversationMessageResponse**](ListServiceConversationMessageResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceConversationMessageReceipt

> ListServiceConversationMessageReceiptResponse ListServiceConversationMessageReceipt(ctx, ChatServiceSid, ConversationSid, MessageSid).PageSize(PageSize).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Message resource is associated with.
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
    MessageSid := "MessageSid_example" // string | The SID of the message within a [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) the delivery receipt belongs to.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListServiceConversationMessageReceipt(context.Background(), ChatServiceSid, ConversationSid, MessageSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListServiceConversationMessageReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceConversationMessageReceipt`: ListServiceConversationMessageReceiptResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListServiceConversationMessageReceipt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Message resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**MessageSid** | **string** | The SID of the message within a [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) the delivery receipt belongs to.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceConversationMessageReceiptParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListServiceConversationMessageReceiptResponse**](ListServiceConversationMessageReceiptResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceConversationParticipant

> ListServiceConversationParticipantResponse ListServiceConversationParticipant(ctx, ChatServiceSid, ConversationSid).PageSize(PageSize).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for participants.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListServiceConversationParticipant(context.Background(), ChatServiceSid, ConversationSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListServiceConversationParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceConversationParticipant`: ListServiceConversationParticipantResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListServiceConversationParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for participants.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceConversationParticipantParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListServiceConversationParticipantResponse**](ListServiceConversationParticipantResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceConversationScopedWebhook

> ListServiceConversationScopedWebhookResponse ListServiceConversationScopedWebhook(ctx, ChatServiceSid, ConversationSid).PageSize(PageSize).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListServiceConversationScopedWebhook(context.Background(), ChatServiceSid, ConversationSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListServiceConversationScopedWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceConversationScopedWebhook`: ListServiceConversationScopedWebhookResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListServiceConversationScopedWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceConversationScopedWebhookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListServiceConversationScopedWebhookResponse**](ListServiceConversationScopedWebhookResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceRole

> ListServiceRoleResponse ListServiceRole(ctx, ChatServiceSid).PageSize(PageSize).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to read the Role resources from.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListServiceRole(context.Background(), ChatServiceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListServiceRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceRole`: ListServiceRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListServiceRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to read the Role resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceRoleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListServiceRoleResponse**](ListServiceRoleResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceUser

> ListServiceUserResponse ListServiceUser(ctx, ChatServiceSid).PageSize(PageSize).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to read the User resources from.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListServiceUser(context.Background(), ChatServiceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListServiceUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceUser`: ListServiceUserResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListServiceUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to read the User resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceUserParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListServiceUserResponse**](ListServiceUserResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUser

> ListUserResponse ListUser(ctx).PageSize(PageSize).Execute()





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
    resp, r, err := api_client.DefaultApi.ListUser(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUser`: ListUserResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

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


## UpdateConfiguration

> ConversationsV1Configuration UpdateConfiguration(ctx).DefaultChatServiceSid(DefaultChatServiceSid).DefaultClosedTimer(DefaultClosedTimer).DefaultInactiveTimer(DefaultInactiveTimer).DefaultMessagingServiceSid(DefaultMessagingServiceSid).Execute()





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
    DefaultChatServiceSid := "DefaultChatServiceSid_example" // string | The SID of the default [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to use when creating a conversation. (optional)
    DefaultClosedTimer := "DefaultClosedTimer_example" // string | Default ISO8601 duration when conversation will be switched to `closed` state. Minimum value for this timer is 10 minutes. (optional)
    DefaultInactiveTimer := "DefaultInactiveTimer_example" // string | Default ISO8601 duration when conversation will be switched to `inactive` state. Minimum value for this timer is 1 minute. (optional)
    DefaultMessagingServiceSid := "DefaultMessagingServiceSid_example" // string | The SID of the default [Messaging Service](https://www.twilio.com/docs/sms/services/api) to use when creating a conversation. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateConfiguration(context.Background()).DefaultChatServiceSid(DefaultChatServiceSid).DefaultClosedTimer(DefaultClosedTimer).DefaultInactiveTimer(DefaultInactiveTimer).DefaultMessagingServiceSid(DefaultMessagingServiceSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfiguration`: ConversationsV1Configuration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConfigurationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **DefaultChatServiceSid** | **string** | The SID of the default [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to use when creating a conversation.
 **DefaultClosedTimer** | **string** | Default ISO8601 duration when conversation will be switched to &#x60;closed&#x60; state. Minimum value for this timer is 10 minutes.
 **DefaultInactiveTimer** | **string** | Default ISO8601 duration when conversation will be switched to &#x60;inactive&#x60; state. Minimum value for this timer is 1 minute.
 **DefaultMessagingServiceSid** | **string** | The SID of the default [Messaging Service](https://www.twilio.com/docs/sms/services/api) to use when creating a conversation.

### Return type

[**ConversationsV1Configuration**](ConversationsV1Configuration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigurationWebhook

> ConversationsV1ConfigurationConfigurationWebhook UpdateConfigurationWebhook(ctx).Filters(Filters).Method(Method).PostWebhookUrl(PostWebhookUrl).PreWebhookUrl(PreWebhookUrl).Target(Target).Execute()



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
    Filters := []string{"Inner_example"} // []string | The list of webhook event triggers that are enabled for this Service: `onMessageAdded`, `onMessageUpdated`, `onMessageRemoved`, `onConversationUpdated`, `onConversationRemoved`, `onParticipantAdded`, `onParticipantUpdated`, `onParticipantRemoved` (optional)
    Method := "Method_example" // string | The HTTP method to be used when sending a webhook request. (optional)
    PostWebhookUrl := "PostWebhookUrl_example" // string | The absolute url the post-event webhook request should be sent to. (optional)
    PreWebhookUrl := "PreWebhookUrl_example" // string | The absolute url the pre-event webhook request should be sent to. (optional)
    Target := "Target_example" // string | The routing target of the webhook. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateConfigurationWebhook(context.Background()).Filters(Filters).Method(Method).PostWebhookUrl(PostWebhookUrl).PreWebhookUrl(PreWebhookUrl).Target(Target).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateConfigurationWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfigurationWebhook`: ConversationsV1ConfigurationConfigurationWebhook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateConfigurationWebhook`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConfigurationWebhookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **Filters** | **[]string** | The list of webhook event triggers that are enabled for this Service: &#x60;onMessageAdded&#x60;, &#x60;onMessageUpdated&#x60;, &#x60;onMessageRemoved&#x60;, &#x60;onConversationUpdated&#x60;, &#x60;onConversationRemoved&#x60;, &#x60;onParticipantAdded&#x60;, &#x60;onParticipantUpdated&#x60;, &#x60;onParticipantRemoved&#x60;
 **Method** | **string** | The HTTP method to be used when sending a webhook request.
 **PostWebhookUrl** | **string** | The absolute url the post-event webhook request should be sent to.
 **PreWebhookUrl** | **string** | The absolute url the pre-event webhook request should be sent to.
 **Target** | **string** | The routing target of the webhook.

### Return type

[**ConversationsV1ConfigurationConfigurationWebhook**](ConversationsV1ConfigurationConfigurationWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConversation

> ConversationsV1Conversation UpdateConversation(ctx, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).DateCreated(DateCreated).DateUpdated(DateUpdated).FriendlyName(FriendlyName).MessagingServiceSid(MessagingServiceSid).State(State).TimersClosed(TimersClosed).TimersInactive(TimersInactive).UniqueName(UniqueName).Execute()





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
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource. Can also be the `unique_name` of the Conversation.
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)
    Attributes := "Attributes_example" // string | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned. (optional)
    DateCreated := time.Now() // time.Time | The date that this resource was created. (optional)
    DateUpdated := time.Now() // time.Time | The date that this resource was last updated. (optional)
    FriendlyName := "FriendlyName_example" // string | The human-readable name of this conversation, limited to 256 characters. Optional. (optional)
    MessagingServiceSid := "MessagingServiceSid_example" // string | The unique ID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) this conversation belongs to. (optional)
    State := "State_example" // string | Current state of this conversation. Can be either `active`, `inactive` or `closed` and defaults to `active` (optional)
    TimersClosed := "TimersClosed_example" // string | ISO8601 duration when conversation will be switched to `closed` state. Minimum value for this timer is 10 minutes. (optional)
    TimersInactive := "TimersInactive_example" // string | ISO8601 duration when conversation will be switched to `inactive` state. Minimum value for this timer is 1 minute. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateConversation(context.Background(), Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).DateCreated(DateCreated).DateUpdated(DateUpdated).FriendlyName(FriendlyName).MessagingServiceSid(MessagingServiceSid).State(State).TimersClosed(TimersClosed).TimersInactive(TimersInactive).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateConversation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConversation`: ConversationsV1Conversation
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateConversation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this resource. Can also be the &#x60;unique_name&#x60; of the Conversation.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConversationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
 **Attributes** | **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\&quot;{}\\\&quot; will be returned.
 **DateCreated** | **time.Time** | The date that this resource was created.
 **DateUpdated** | **time.Time** | The date that this resource was last updated.
 **FriendlyName** | **string** | The human-readable name of this conversation, limited to 256 characters. Optional.
 **MessagingServiceSid** | **string** | The unique ID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) this conversation belongs to.
 **State** | **string** | Current state of this conversation. Can be either &#x60;active&#x60;, &#x60;inactive&#x60; or &#x60;closed&#x60; and defaults to &#x60;active&#x60;
 **TimersClosed** | **string** | ISO8601 duration when conversation will be switched to &#x60;closed&#x60; state. Minimum value for this timer is 10 minutes.
 **TimersInactive** | **string** | ISO8601 duration when conversation will be switched to &#x60;inactive&#x60; state. Minimum value for this timer is 1 minute.
 **UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource&#39;s &#x60;sid&#x60; in the URL.

### Return type

[**ConversationsV1Conversation**](ConversationsV1Conversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConversationMessage

> ConversationsV1ConversationConversationMessage UpdateConversationMessage(ctx, ConversationSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).Author(Author).Body(Body).DateCreated(DateCreated).DateUpdated(DateUpdated).Execute()





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
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)
    Attributes := "Attributes_example" // string | A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned. (optional)
    Author := "Author_example" // string | The channel specific identifier of the message's author. Defaults to `system`. (optional)
    Body := "Body_example" // string | The content of the message, can be up to 1,600 characters long. (optional)
    DateCreated := time.Now() // time.Time | The date that this resource was created. (optional)
    DateUpdated := time.Now() // time.Time | The date that this resource was last updated. `null` if the message has not been edited. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateConversationMessage(context.Background(), ConversationSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).Author(Author).Body(Body).DateCreated(DateCreated).DateUpdated(DateUpdated).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateConversationMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConversationMessage`: ConversationsV1ConversationConversationMessage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateConversationMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConversationMessageParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
 **Attributes** | **string** | A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\&quot;{}\\\&quot; will be returned.
 **Author** | **string** | The channel specific identifier of the message&#39;s author. Defaults to &#x60;system&#x60;.
 **Body** | **string** | The content of the message, can be up to 1,600 characters long.
 **DateCreated** | **time.Time** | The date that this resource was created.
 **DateUpdated** | **time.Time** | The date that this resource was last updated. &#x60;null&#x60; if the message has not been edited.

### Return type

[**ConversationsV1ConversationConversationMessage**](ConversationsV1ConversationConversationMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConversationParticipant

> ConversationsV1ConversationConversationParticipant UpdateConversationParticipant(ctx, ConversationSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).DateCreated(DateCreated).DateUpdated(DateUpdated).Identity(Identity).LastReadMessageIndex(LastReadMessageIndex).LastReadTimestamp(LastReadTimestamp).MessagingBindingProjectedAddress(MessagingBindingProjectedAddress).MessagingBindingProxyAddress(MessagingBindingProxyAddress).RoleSid(RoleSid).Execute()





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
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)
    Attributes := "Attributes_example" // string | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned. (optional)
    DateCreated := time.Now() // time.Time | The date that this resource was created. (optional)
    DateUpdated := time.Now() // time.Time | The date that this resource was last updated. (optional)
    Identity := "Identity_example" // string | A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversations SDK to communicate. Limited to 256 characters. (optional)
    LastReadMessageIndex := int32(56) // int32 | Index of last “read” message in the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for the Participant. (optional)
    LastReadTimestamp := "LastReadTimestamp_example" // string | Timestamp of last “read” message in the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for the Participant. (optional)
    MessagingBindingProjectedAddress := "MessagingBindingProjectedAddress_example" // string | The address of the Twilio phone number that is used in Group MMS. 'null' value will remove it. (optional)
    MessagingBindingProxyAddress := "MessagingBindingProxyAddress_example" // string | The address of the Twilio phone number that the participant is in contact with. 'null' value will remove it. (optional)
    RoleSid := "RoleSid_example" // string | The SID of a conversation-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the participant. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateConversationParticipant(context.Background(), ConversationSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).DateCreated(DateCreated).DateUpdated(DateUpdated).Identity(Identity).LastReadMessageIndex(LastReadMessageIndex).LastReadTimestamp(LastReadTimestamp).MessagingBindingProjectedAddress(MessagingBindingProjectedAddress).MessagingBindingProxyAddress(MessagingBindingProxyAddress).RoleSid(RoleSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateConversationParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConversationParticipant`: ConversationsV1ConversationConversationParticipant
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateConversationParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConversationParticipantParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
 **Attributes** | **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\&quot;{}\\\&quot; will be returned.
 **DateCreated** | **time.Time** | The date that this resource was created.
 **DateUpdated** | **time.Time** | The date that this resource was last updated.
 **Identity** | **string** | A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversations SDK to communicate. Limited to 256 characters.
 **LastReadMessageIndex** | **int32** | Index of last “read” message in the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for the Participant.
 **LastReadTimestamp** | **string** | Timestamp of last “read” message in the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for the Participant.
 **MessagingBindingProjectedAddress** | **string** | The address of the Twilio phone number that is used in Group MMS. &#39;null&#39; value will remove it.
 **MessagingBindingProxyAddress** | **string** | The address of the Twilio phone number that the participant is in contact with. &#39;null&#39; value will remove it.
 **RoleSid** | **string** | The SID of a conversation-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the participant.

### Return type

[**ConversationsV1ConversationConversationParticipant**](ConversationsV1ConversationConversationParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConversationScopedWebhook

> ConversationsV1ConversationConversationScopedWebhook UpdateConversationScopedWebhook(ctx, ConversationSid, Sid).ConfigurationFilters(ConfigurationFilters).ConfigurationFlowSid(ConfigurationFlowSid).ConfigurationMethod(ConfigurationMethod).ConfigurationTriggers(ConfigurationTriggers).ConfigurationUrl(ConfigurationUrl).Execute()





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
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.
    ConfigurationFilters := []string{"Inner_example"} // []string | The list of events, firing webhook event for this Conversation. (optional)
    ConfigurationFlowSid := "ConfigurationFlowSid_example" // string | The studio flow SID, where the webhook should be sent to. (optional)
    ConfigurationMethod := "ConfigurationMethod_example" // string | The HTTP method to be used when sending a webhook request. (optional)
    ConfigurationTriggers := []string{"Inner_example"} // []string | The list of keywords, firing webhook event for this Conversation. (optional)
    ConfigurationUrl := "ConfigurationUrl_example" // string | The absolute url the webhook request should be sent to. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateConversationScopedWebhook(context.Background(), ConversationSid, Sid).ConfigurationFilters(ConfigurationFilters).ConfigurationFlowSid(ConfigurationFlowSid).ConfigurationMethod(ConfigurationMethod).ConfigurationTriggers(ConfigurationTriggers).ConfigurationUrl(ConfigurationUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateConversationScopedWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConversationScopedWebhook`: ConversationsV1ConversationConversationScopedWebhook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateConversationScopedWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConversationScopedWebhookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **ConfigurationFilters** | **[]string** | The list of events, firing webhook event for this Conversation.
 **ConfigurationFlowSid** | **string** | The studio flow SID, where the webhook should be sent to.
 **ConfigurationMethod** | **string** | The HTTP method to be used when sending a webhook request.
 **ConfigurationTriggers** | **[]string** | The list of keywords, firing webhook event for this Conversation.
 **ConfigurationUrl** | **string** | The absolute url the webhook request should be sent to.

### Return type

[**ConversationsV1ConversationConversationScopedWebhook**](ConversationsV1ConversationConversationScopedWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredential

> ConversationsV1Credential UpdateCredential(ctx, Sid).ApiKey(ApiKey).Certificate(Certificate).FriendlyName(FriendlyName).PrivateKey(PrivateKey).Sandbox(Sandbox).Secret(Secret).Type(Type).Execute()





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
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.
    ApiKey := "ApiKey_example" // string | [GCM only] The API key for the project that was obtained from the Google Developer console for your GCM Service application credential. (optional)
    Certificate := "Certificate_example" // string | [APN only] The URL encoded representation of the certificate. For example,  `-----BEGIN CERTIFICATE----- MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEF.....A== -----END CERTIFICATE-----`. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the new resource. It can be up to 64 characters long. (optional)
    PrivateKey := "PrivateKey_example" // string | [APN only] The URL encoded representation of the private key. For example, `-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fG... -----END RSA PRIVATE KEY-----`. (optional)
    Sandbox := true // bool | [APN only] Whether to send the credential to sandbox APNs. Can be `true` to send to sandbox APNs or `false` to send to production. (optional)
    Secret := "Secret_example" // string | [FCM only] The **Server key** of your project from the Firebase console, found under Settings / Cloud messaging. (optional)
    Type := "Type_example" // string | The type of push-notification service the credential is for. Can be: `fcm`, `gcm`, or `apn`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateCredential(context.Background(), Sid).ApiKey(ApiKey).Certificate(Certificate).FriendlyName(FriendlyName).PrivateKey(PrivateKey).Sandbox(Sandbox).Secret(Secret).Type(Type).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCredential`: ConversationsV1Credential
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateCredentialParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **ApiKey** | **string** | [GCM only] The API key for the project that was obtained from the Google Developer console for your GCM Service application credential.
 **Certificate** | **string** | [APN only] The URL encoded representation of the certificate. For example,  &#x60;-----BEGIN CERTIFICATE----- MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEF.....A&#x3D;&#x3D; -----END CERTIFICATE-----&#x60;.
 **FriendlyName** | **string** | A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
 **PrivateKey** | **string** | [APN only] The URL encoded representation of the private key. For example, &#x60;-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fG... -----END RSA PRIVATE KEY-----&#x60;.
 **Sandbox** | **bool** | [APN only] Whether to send the credential to sandbox APNs. Can be &#x60;true&#x60; to send to sandbox APNs or &#x60;false&#x60; to send to production.
 **Secret** | **string** | [FCM only] The **Server key** of your project from the Firebase console, found under Settings / Cloud messaging.
 **Type** | **string** | The type of push-notification service the credential is for. Can be: &#x60;fcm&#x60;, &#x60;gcm&#x60;, or &#x60;apn&#x60;.

### Return type

[**ConversationsV1Credential**](ConversationsV1Credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> ConversationsV1Role UpdateRole(ctx, Sid).Permission(Permission).Execute()





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
    Sid := "Sid_example" // string | The SID of the Role resource to update.
    Permission := []string{"Inner_example"} // []string | A permission that you grant to the role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. Note that the update action replaces all previously assigned permissions with those defined in the update action. To remove a permission, do not include it in the subsequent update action. The values for this parameter depend on the role's `type`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateRole(context.Background(), Sid).Permission(Permission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRole`: ConversationsV1Role
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Role resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRoleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Permission** | **[]string** | A permission that you grant to the role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. Note that the update action replaces all previously assigned permissions with those defined in the update action. To remove a permission, do not include it in the subsequent update action. The values for this parameter depend on the role&#39;s &#x60;type&#x60;.

### Return type

[**ConversationsV1Role**](ConversationsV1Role.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceConfiguration

> ConversationsV1ServiceServiceConfiguration UpdateServiceConfiguration(ctx, ChatServiceSid).DefaultChatServiceRoleSid(DefaultChatServiceRoleSid).DefaultConversationCreatorRoleSid(DefaultConversationCreatorRoleSid).DefaultConversationRoleSid(DefaultConversationRoleSid).ReachabilityEnabled(ReachabilityEnabled).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the Service configuration resource to update.
    DefaultChatServiceRoleSid := "DefaultChatServiceRoleSid_example" // string | The service-level role assigned to users when they are added to the service. See the [Conversation Role](https://www.twilio.com/docs/conversations/api/role-resource) for more info about roles. (optional)
    DefaultConversationCreatorRoleSid := "DefaultConversationCreatorRoleSid_example" // string | The conversation-level role assigned to a conversation creator when they join a new conversation. See the [Conversation Role](https://www.twilio.com/docs/conversations/api/role-resource) for more info about roles. (optional)
    DefaultConversationRoleSid := "DefaultConversationRoleSid_example" // string | The conversation-level role assigned to users when they are added to a conversation. See the [Conversation Role](https://www.twilio.com/docs/conversations/api/role-resource) for more info about roles. (optional)
    ReachabilityEnabled := true // bool | Whether the [Reachability Indicator](https://www.twilio.com/docs/chat/reachability-indicator) is enabled for this Conversations Service. The default is `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateServiceConfiguration(context.Background(), ChatServiceSid).DefaultChatServiceRoleSid(DefaultChatServiceRoleSid).DefaultConversationCreatorRoleSid(DefaultConversationCreatorRoleSid).DefaultConversationRoleSid(DefaultConversationRoleSid).ReachabilityEnabled(ReachabilityEnabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateServiceConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServiceConfiguration`: ConversationsV1ServiceServiceConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateServiceConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the Service configuration resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceConfigurationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **DefaultChatServiceRoleSid** | **string** | The service-level role assigned to users when they are added to the service. See the [Conversation Role](https://www.twilio.com/docs/conversations/api/role-resource) for more info about roles.
 **DefaultConversationCreatorRoleSid** | **string** | The conversation-level role assigned to a conversation creator when they join a new conversation. See the [Conversation Role](https://www.twilio.com/docs/conversations/api/role-resource) for more info about roles.
 **DefaultConversationRoleSid** | **string** | The conversation-level role assigned to users when they are added to a conversation. See the [Conversation Role](https://www.twilio.com/docs/conversations/api/role-resource) for more info about roles.
 **ReachabilityEnabled** | **bool** | Whether the [Reachability Indicator](https://www.twilio.com/docs/chat/reachability-indicator) is enabled for this Conversations Service. The default is &#x60;false&#x60;.

### Return type

[**ConversationsV1ServiceServiceConfiguration**](ConversationsV1ServiceServiceConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceConversation

> ConversationsV1ServiceServiceConversation UpdateServiceConversation(ctx, ChatServiceSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).DateCreated(DateCreated).DateUpdated(DateUpdated).FriendlyName(FriendlyName).MessagingServiceSid(MessagingServiceSid).State(State).TimersClosed(TimersClosed).TimersInactive(TimersInactive).UniqueName(UniqueName).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource. Can also be the `unique_name` of the Conversation.
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)
    Attributes := "Attributes_example" // string | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned. (optional)
    DateCreated := time.Now() // time.Time | The date that this resource was created. (optional)
    DateUpdated := time.Now() // time.Time | The date that this resource was last updated. (optional)
    FriendlyName := "FriendlyName_example" // string | The human-readable name of this conversation, limited to 256 characters. Optional. (optional)
    MessagingServiceSid := "MessagingServiceSid_example" // string | The unique ID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) this conversation belongs to. (optional)
    State := "State_example" // string | Current state of this conversation. Can be either `active`, `inactive` or `closed` and defaults to `active` (optional)
    TimersClosed := "TimersClosed_example" // string | ISO8601 duration when conversation will be switched to `closed` state. Minimum value for this timer is 10 minutes. (optional)
    TimersInactive := "TimersInactive_example" // string | ISO8601 duration when conversation will be switched to `inactive` state. Minimum value for this timer is 1 minute. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateServiceConversation(context.Background(), ChatServiceSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).DateCreated(DateCreated).DateUpdated(DateUpdated).FriendlyName(FriendlyName).MessagingServiceSid(MessagingServiceSid).State(State).TimersClosed(TimersClosed).TimersInactive(TimersInactive).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateServiceConversation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServiceConversation`: ConversationsV1ServiceServiceConversation
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateServiceConversation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Conversation resource is associated with.
**Sid** | **string** | A 34 character string that uniquely identifies this resource. Can also be the &#x60;unique_name&#x60; of the Conversation.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceConversationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
 **Attributes** | **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\&quot;{}\\\&quot; will be returned.
 **DateCreated** | **time.Time** | The date that this resource was created.
 **DateUpdated** | **time.Time** | The date that this resource was last updated.
 **FriendlyName** | **string** | The human-readable name of this conversation, limited to 256 characters. Optional.
 **MessagingServiceSid** | **string** | The unique ID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) this conversation belongs to.
 **State** | **string** | Current state of this conversation. Can be either &#x60;active&#x60;, &#x60;inactive&#x60; or &#x60;closed&#x60; and defaults to &#x60;active&#x60;
 **TimersClosed** | **string** | ISO8601 duration when conversation will be switched to &#x60;closed&#x60; state. Minimum value for this timer is 10 minutes.
 **TimersInactive** | **string** | ISO8601 duration when conversation will be switched to &#x60;inactive&#x60; state. Minimum value for this timer is 1 minute.
 **UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource&#39;s &#x60;sid&#x60; in the URL.

### Return type

[**ConversationsV1ServiceServiceConversation**](ConversationsV1ServiceServiceConversation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceConversationMessage

> ConversationsV1ServiceServiceConversationServiceConversationMessage UpdateServiceConversationMessage(ctx, ChatServiceSid, ConversationSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).Author(Author).Body(Body).DateCreated(DateCreated).DateUpdated(DateUpdated).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)
    Attributes := "Attributes_example" // string | A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned. (optional)
    Author := "Author_example" // string | The channel specific identifier of the message's author. Defaults to `system`. (optional)
    Body := "Body_example" // string | The content of the message, can be up to 1,600 characters long. (optional)
    DateCreated := time.Now() // time.Time | The date that this resource was created. (optional)
    DateUpdated := time.Now() // time.Time | The date that this resource was last updated. `null` if the message has not been edited. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateServiceConversationMessage(context.Background(), ChatServiceSid, ConversationSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).Author(Author).Body(Body).DateCreated(DateCreated).DateUpdated(DateUpdated).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateServiceConversationMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServiceConversationMessage`: ConversationsV1ServiceServiceConversationServiceConversationMessage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateServiceConversationMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this message.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceConversationMessageParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
 **Attributes** | **string** | A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\&quot;{}\\\&quot; will be returned.
 **Author** | **string** | The channel specific identifier of the message&#39;s author. Defaults to &#x60;system&#x60;.
 **Body** | **string** | The content of the message, can be up to 1,600 characters long.
 **DateCreated** | **time.Time** | The date that this resource was created.
 **DateUpdated** | **time.Time** | The date that this resource was last updated. &#x60;null&#x60; if the message has not been edited.

### Return type

[**ConversationsV1ServiceServiceConversationServiceConversationMessage**](ConversationsV1ServiceServiceConversationServiceConversationMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceConversationParticipant

> ConversationsV1ServiceServiceConversationServiceConversationParticipant UpdateServiceConversationParticipant(ctx, ChatServiceSid, ConversationSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).DateCreated(DateCreated).DateUpdated(DateUpdated).Identity(Identity).LastReadMessageIndex(LastReadMessageIndex).LastReadTimestamp(LastReadTimestamp).MessagingBindingProjectedAddress(MessagingBindingProjectedAddress).MessagingBindingProxyAddress(MessagingBindingProxyAddress).RoleSid(RoleSid).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)
    Attributes := "Attributes_example" // string | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned. (optional)
    DateCreated := time.Now() // time.Time | The date that this resource was created. (optional)
    DateUpdated := time.Now() // time.Time | The date that this resource was last updated. (optional)
    Identity := "Identity_example" // string | A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversation SDK to communicate. Limited to 256 characters. (optional)
    LastReadMessageIndex := int32(56) // int32 | Index of last “read” message in the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for the Participant. (optional)
    LastReadTimestamp := "LastReadTimestamp_example" // string | Timestamp of last “read” message in the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for the Participant. (optional)
    MessagingBindingProjectedAddress := "MessagingBindingProjectedAddress_example" // string | The address of the Twilio phone number that is used in Group MMS. 'null' value will remove it. (optional)
    MessagingBindingProxyAddress := "MessagingBindingProxyAddress_example" // string | The address of the Twilio phone number that the participant is in contact with. 'null' value will remove it. (optional)
    RoleSid := "RoleSid_example" // string | The SID of a conversation-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the participant. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateServiceConversationParticipant(context.Background(), ChatServiceSid, ConversationSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).DateCreated(DateCreated).DateUpdated(DateUpdated).Identity(Identity).LastReadMessageIndex(LastReadMessageIndex).LastReadTimestamp(LastReadTimestamp).MessagingBindingProjectedAddress(MessagingBindingProjectedAddress).MessagingBindingProxyAddress(MessagingBindingProxyAddress).RoleSid(RoleSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateServiceConversationParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServiceConversationParticipant`: ConversationsV1ServiceServiceConversationServiceConversationParticipant
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateServiceConversationParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this participant.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceConversationParticipantParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
 **Attributes** | **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\&quot;{}\\\&quot; will be returned.
 **DateCreated** | **time.Time** | The date that this resource was created.
 **DateUpdated** | **time.Time** | The date that this resource was last updated.
 **Identity** | **string** | A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversation SDK to communicate. Limited to 256 characters.
 **LastReadMessageIndex** | **int32** | Index of last “read” message in the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for the Participant.
 **LastReadTimestamp** | **string** | Timestamp of last “read” message in the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for the Participant.
 **MessagingBindingProjectedAddress** | **string** | The address of the Twilio phone number that is used in Group MMS. &#39;null&#39; value will remove it.
 **MessagingBindingProxyAddress** | **string** | The address of the Twilio phone number that the participant is in contact with. &#39;null&#39; value will remove it.
 **RoleSid** | **string** | The SID of a conversation-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the participant.

### Return type

[**ConversationsV1ServiceServiceConversationServiceConversationParticipant**](ConversationsV1ServiceServiceConversationServiceConversationParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceConversationScopedWebhook

> ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook UpdateServiceConversationScopedWebhook(ctx, ChatServiceSid, ConversationSid, Sid).ConfigurationFilters(ConfigurationFilters).ConfigurationFlowSid(ConfigurationFlowSid).ConfigurationMethod(ConfigurationMethod).ConfigurationTriggers(ConfigurationTriggers).ConfigurationUrl(ConfigurationUrl).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
    ConversationSid := "ConversationSid_example" // string | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.
    ConfigurationFilters := []string{"Inner_example"} // []string | The list of events, firing webhook event for this Conversation. (optional)
    ConfigurationFlowSid := "ConfigurationFlowSid_example" // string | The studio flow SID, where the webhook should be sent to. (optional)
    ConfigurationMethod := "ConfigurationMethod_example" // string | The HTTP method to be used when sending a webhook request. (optional)
    ConfigurationTriggers := []string{"Inner_example"} // []string | The list of keywords, firing webhook event for this Conversation. (optional)
    ConfigurationUrl := "ConfigurationUrl_example" // string | The absolute url the webhook request should be sent to. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateServiceConversationScopedWebhook(context.Background(), ChatServiceSid, ConversationSid, Sid).ConfigurationFilters(ConfigurationFilters).ConfigurationFlowSid(ConfigurationFlowSid).ConfigurationMethod(ConfigurationMethod).ConfigurationTriggers(ConfigurationTriggers).ConfigurationUrl(ConfigurationUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateServiceConversationScopedWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServiceConversationScopedWebhook`: ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateServiceConversationScopedWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Participant resource is associated with.
**ConversationSid** | **string** | The unique ID of the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for this webhook.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceConversationScopedWebhookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **ConfigurationFilters** | **[]string** | The list of events, firing webhook event for this Conversation.
 **ConfigurationFlowSid** | **string** | The studio flow SID, where the webhook should be sent to.
 **ConfigurationMethod** | **string** | The HTTP method to be used when sending a webhook request.
 **ConfigurationTriggers** | **[]string** | The list of keywords, firing webhook event for this Conversation.
 **ConfigurationUrl** | **string** | The absolute url the webhook request should be sent to.

### Return type

[**ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook**](ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceNotification

> ConversationsV1ServiceServiceConfigurationServiceNotification UpdateServiceNotification(ctx, ChatServiceSid).AddedToConversationEnabled(AddedToConversationEnabled).AddedToConversationSound(AddedToConversationSound).AddedToConversationTemplate(AddedToConversationTemplate).LogEnabled(LogEnabled).NewMessageBadgeCountEnabled(NewMessageBadgeCountEnabled).NewMessageEnabled(NewMessageEnabled).NewMessageSound(NewMessageSound).NewMessageTemplate(NewMessageTemplate).RemovedFromConversationEnabled(RemovedFromConversationEnabled).RemovedFromConversationSound(RemovedFromConversationSound).RemovedFromConversationTemplate(RemovedFromConversationTemplate).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Configuration applies to.
    AddedToConversationEnabled := true // bool | Whether to send a notification when a participant is added to a conversation. The default is `false`. (optional)
    AddedToConversationSound := "AddedToConversationSound_example" // string | The name of the sound to play when a participant is added to a conversation and `added_to_conversation.enabled` is `true`. (optional)
    AddedToConversationTemplate := "AddedToConversationTemplate_example" // string | The template to use to create the notification text displayed when a participant is added to a conversation and `added_to_conversation.enabled` is `true`. (optional)
    LogEnabled := true // bool | Weather the notification logging is enabled. (optional)
    NewMessageBadgeCountEnabled := true // bool | Whether the new message badge is enabled. The default is `false`. (optional)
    NewMessageEnabled := true // bool | Whether to send a notification when a new message is added to a conversation. The default is `false`. (optional)
    NewMessageSound := "NewMessageSound_example" // string | The name of the sound to play when a new message is added to a conversation and `new_message.enabled` is `true`. (optional)
    NewMessageTemplate := "NewMessageTemplate_example" // string | The template to use to create the notification text displayed when a new message is added to a conversation and `new_message.enabled` is `true`. (optional)
    RemovedFromConversationEnabled := true // bool | Whether to send a notification to a user when they are removed from a conversation. The default is `false`. (optional)
    RemovedFromConversationSound := "RemovedFromConversationSound_example" // string | The name of the sound to play to a user when they are removed from a conversation and `removed_from_conversation.enabled` is `true`. (optional)
    RemovedFromConversationTemplate := "RemovedFromConversationTemplate_example" // string | The template to use to create the notification text displayed to a user when they are removed from a conversation and `removed_from_conversation.enabled` is `true`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateServiceNotification(context.Background(), ChatServiceSid).AddedToConversationEnabled(AddedToConversationEnabled).AddedToConversationSound(AddedToConversationSound).AddedToConversationTemplate(AddedToConversationTemplate).LogEnabled(LogEnabled).NewMessageBadgeCountEnabled(NewMessageBadgeCountEnabled).NewMessageEnabled(NewMessageEnabled).NewMessageSound(NewMessageSound).NewMessageTemplate(NewMessageTemplate).RemovedFromConversationEnabled(RemovedFromConversationEnabled).RemovedFromConversationSound(RemovedFromConversationSound).RemovedFromConversationTemplate(RemovedFromConversationTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateServiceNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServiceNotification`: ConversationsV1ServiceServiceConfigurationServiceNotification
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateServiceNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the Configuration applies to.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceNotificationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **AddedToConversationEnabled** | **bool** | Whether to send a notification when a participant is added to a conversation. The default is &#x60;false&#x60;.
 **AddedToConversationSound** | **string** | The name of the sound to play when a participant is added to a conversation and &#x60;added_to_conversation.enabled&#x60; is &#x60;true&#x60;.
 **AddedToConversationTemplate** | **string** | The template to use to create the notification text displayed when a participant is added to a conversation and &#x60;added_to_conversation.enabled&#x60; is &#x60;true&#x60;.
 **LogEnabled** | **bool** | Weather the notification logging is enabled.
 **NewMessageBadgeCountEnabled** | **bool** | Whether the new message badge is enabled. The default is &#x60;false&#x60;.
 **NewMessageEnabled** | **bool** | Whether to send a notification when a new message is added to a conversation. The default is &#x60;false&#x60;.
 **NewMessageSound** | **string** | The name of the sound to play when a new message is added to a conversation and &#x60;new_message.enabled&#x60; is &#x60;true&#x60;.
 **NewMessageTemplate** | **string** | The template to use to create the notification text displayed when a new message is added to a conversation and &#x60;new_message.enabled&#x60; is &#x60;true&#x60;.
 **RemovedFromConversationEnabled** | **bool** | Whether to send a notification to a user when they are removed from a conversation. The default is &#x60;false&#x60;.
 **RemovedFromConversationSound** | **string** | The name of the sound to play to a user when they are removed from a conversation and &#x60;removed_from_conversation.enabled&#x60; is &#x60;true&#x60;.
 **RemovedFromConversationTemplate** | **string** | The template to use to create the notification text displayed to a user when they are removed from a conversation and &#x60;removed_from_conversation.enabled&#x60; is &#x60;true&#x60;.

### Return type

[**ConversationsV1ServiceServiceConfigurationServiceNotification**](ConversationsV1ServiceServiceConfigurationServiceNotification.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceRole

> ConversationsV1ServiceServiceRole UpdateServiceRole(ctx, ChatServiceSid, Sid).Permission(Permission).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to update the Role resource in.
    Sid := "Sid_example" // string | The SID of the Role resource to update.
    Permission := []string{"Inner_example"} // []string | A permission that you grant to the role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. Note that the update action replaces all previously assigned permissions with those defined in the update action. To remove a permission, do not include it in the subsequent update action. The values for this parameter depend on the role's `type`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateServiceRole(context.Background(), ChatServiceSid, Sid).Permission(Permission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateServiceRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServiceRole`: ConversationsV1ServiceServiceRole
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateServiceRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) to update the Role resource in.
**Sid** | **string** | The SID of the Role resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceRoleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Permission** | **[]string** | A permission that you grant to the role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. Note that the update action replaces all previously assigned permissions with those defined in the update action. To remove a permission, do not include it in the subsequent update action. The values for this parameter depend on the role&#39;s &#x60;type&#x60;.

### Return type

[**ConversationsV1ServiceServiceRole**](ConversationsV1ServiceServiceRole.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceUser

> ConversationsV1ServiceServiceUser UpdateServiceUser(ctx, ChatServiceSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).FriendlyName(FriendlyName).RoleSid(RoleSid).Execute()





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
    ChatServiceSid := "ChatServiceSid_example" // string | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the User resource is associated with.
    Sid := "Sid_example" // string | The SID of the User resource to update. This value can be either the `sid` or the `identity` of the User resource to update.
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)
    Attributes := "Attributes_example" // string | The JSON Object string that stores application-specific data. If attributes have not been set, `{}` is returned. (optional)
    FriendlyName := "FriendlyName_example" // string | The string that you assigned to describe the resource. (optional)
    RoleSid := "RoleSid_example" // string | The SID of a service-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateServiceUser(context.Background(), ChatServiceSid, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).FriendlyName(FriendlyName).RoleSid(RoleSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateServiceUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServiceUser`: ConversationsV1ServiceServiceUser
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateServiceUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ChatServiceSid** | **string** | The SID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) the User resource is associated with.
**Sid** | **string** | The SID of the User resource to update. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceUserParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
 **Attributes** | **string** | The JSON Object string that stores application-specific data. If attributes have not been set, &#x60;{}&#x60; is returned.
 **FriendlyName** | **string** | The string that you assigned to describe the resource.
 **RoleSid** | **string** | The SID of a service-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the user.

### Return type

[**ConversationsV1ServiceServiceUser**](ConversationsV1ServiceServiceUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> ConversationsV1User UpdateUser(ctx, Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).FriendlyName(FriendlyName).RoleSid(RoleSid).Execute()





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
    Sid := "Sid_example" // string | The SID of the User resource to update. This value can be either the `sid` or the `identity` of the User resource to update.
    XTwilioWebhookEnabled := "XTwilioWebhookEnabled_example" // string | The X-Twilio-Webhook-Enabled HTTP request header (optional)
    Attributes := "Attributes_example" // string | The JSON Object string that stores application-specific data. If attributes have not been set, `{}` is returned. (optional)
    FriendlyName := "FriendlyName_example" // string | The string that you assigned to describe the resource. (optional)
    RoleSid := "RoleSid_example" // string | The SID of a service-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateUser(context.Background(), Sid).XTwilioWebhookEnabled(XTwilioWebhookEnabled).Attributes(Attributes).FriendlyName(FriendlyName).RoleSid(RoleSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: ConversationsV1User
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the User resource to update. This value can be either the &#x60;sid&#x60; or the &#x60;identity&#x60; of the User resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateUserParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **XTwilioWebhookEnabled** | **string** | The X-Twilio-Webhook-Enabled HTTP request header
 **Attributes** | **string** | The JSON Object string that stores application-specific data. If attributes have not been set, &#x60;{}&#x60; is returned.
 **FriendlyName** | **string** | The string that you assigned to describe the resource.
 **RoleSid** | **string** | The SID of a service-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the user.

### Return type

[**ConversationsV1User**](ConversationsV1User.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

