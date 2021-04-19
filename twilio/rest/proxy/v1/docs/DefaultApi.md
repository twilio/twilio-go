# \DefaultApi

All URIs are relative to *https://proxy.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessageInteraction**](DefaultApi.md#CreateMessageInteraction) | **Post** /v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants/{ParticipantSid}/MessageInteractions | 
[**CreateParticipant**](DefaultApi.md#CreateParticipant) | **Post** /v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants | 
[**CreatePhoneNumber**](DefaultApi.md#CreatePhoneNumber) | **Post** /v1/Services/{ServiceSid}/PhoneNumbers | 
[**CreateService**](DefaultApi.md#CreateService) | **Post** /v1/Services | 
[**CreateSession**](DefaultApi.md#CreateSession) | **Post** /v1/Services/{ServiceSid}/Sessions | 
[**CreateShortCode**](DefaultApi.md#CreateShortCode) | **Post** /v1/Services/{ServiceSid}/ShortCodes | 
[**DeleteInteraction**](DefaultApi.md#DeleteInteraction) | **Delete** /v1/Services/{ServiceSid}/Sessions/{SessionSid}/Interactions/{Sid} | 
[**DeleteParticipant**](DefaultApi.md#DeleteParticipant) | **Delete** /v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants/{Sid} | 
[**DeletePhoneNumber**](DefaultApi.md#DeletePhoneNumber) | **Delete** /v1/Services/{ServiceSid}/PhoneNumbers/{Sid} | 
[**DeleteService**](DefaultApi.md#DeleteService) | **Delete** /v1/Services/{Sid} | 
[**DeleteSession**](DefaultApi.md#DeleteSession) | **Delete** /v1/Services/{ServiceSid}/Sessions/{Sid} | 
[**DeleteShortCode**](DefaultApi.md#DeleteShortCode) | **Delete** /v1/Services/{ServiceSid}/ShortCodes/{Sid} | 
[**FetchInteraction**](DefaultApi.md#FetchInteraction) | **Get** /v1/Services/{ServiceSid}/Sessions/{SessionSid}/Interactions/{Sid} | 
[**FetchMessageInteraction**](DefaultApi.md#FetchMessageInteraction) | **Get** /v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants/{ParticipantSid}/MessageInteractions/{Sid} | 
[**FetchParticipant**](DefaultApi.md#FetchParticipant) | **Get** /v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants/{Sid} | 
[**FetchPhoneNumber**](DefaultApi.md#FetchPhoneNumber) | **Get** /v1/Services/{ServiceSid}/PhoneNumbers/{Sid} | 
[**FetchService**](DefaultApi.md#FetchService) | **Get** /v1/Services/{Sid} | 
[**FetchSession**](DefaultApi.md#FetchSession) | **Get** /v1/Services/{ServiceSid}/Sessions/{Sid} | 
[**FetchShortCode**](DefaultApi.md#FetchShortCode) | **Get** /v1/Services/{ServiceSid}/ShortCodes/{Sid} | 
[**ListInteraction**](DefaultApi.md#ListInteraction) | **Get** /v1/Services/{ServiceSid}/Sessions/{SessionSid}/Interactions | 
[**ListMessageInteraction**](DefaultApi.md#ListMessageInteraction) | **Get** /v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants/{ParticipantSid}/MessageInteractions | 
[**ListParticipant**](DefaultApi.md#ListParticipant) | **Get** /v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants | 
[**ListPhoneNumber**](DefaultApi.md#ListPhoneNumber) | **Get** /v1/Services/{ServiceSid}/PhoneNumbers | 
[**ListService**](DefaultApi.md#ListService) | **Get** /v1/Services | 
[**ListSession**](DefaultApi.md#ListSession) | **Get** /v1/Services/{ServiceSid}/Sessions | 
[**ListShortCode**](DefaultApi.md#ListShortCode) | **Get** /v1/Services/{ServiceSid}/ShortCodes | 
[**UpdatePhoneNumber**](DefaultApi.md#UpdatePhoneNumber) | **Post** /v1/Services/{ServiceSid}/PhoneNumbers/{Sid} | 
[**UpdateService**](DefaultApi.md#UpdateService) | **Post** /v1/Services/{Sid} | 
[**UpdateSession**](DefaultApi.md#UpdateSession) | **Post** /v1/Services/{ServiceSid}/Sessions/{Sid} | 
[**UpdateShortCode**](DefaultApi.md#UpdateShortCode) | **Post** /v1/Services/{ServiceSid}/ShortCodes/{Sid} | 



## CreateMessageInteraction

> ProxyV1ServiceSessionParticipantMessageInteraction CreateMessageInteraction(ctx, ServiceSid, SessionSid, ParticipantSid).Body(Body).MediaUrl(MediaUrl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource.
    SessionSid := "SessionSid_example" // string | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) resource.
    ParticipantSid := "ParticipantSid_example" // string | The SID of the [Participant](https://www.twilio.com/docs/proxy/api/participant) resource.
    Body := "Body_example" // string | The message to send to the participant (optional)
    MediaUrl := []string{"Inner_example"} // []string | Reserved. Not currently supported. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateMessageInteraction(context.Background(), ServiceSid, SessionSid, ParticipantSid).Body(Body).MediaUrl(MediaUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateMessageInteraction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMessageInteraction`: ProxyV1ServiceSessionParticipantMessageInteraction
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateMessageInteraction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource.
**SessionSid** | **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) resource.
**ParticipantSid** | **string** | The SID of the [Participant](https://www.twilio.com/docs/proxy/api/participant) resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateMessageInteractionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **Body** | **string** | The message to send to the participant
 **MediaUrl** | **[]string** | Reserved. Not currently supported.

### Return type

[**ProxyV1ServiceSessionParticipantMessageInteraction**](ProxyV1ServiceSessionParticipantMessageInteraction.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateParticipant

> ProxyV1ServiceSessionParticipant CreateParticipant(ctx, ServiceSid, SessionSid).FailOnParticipantConflict(FailOnParticipantConflict).FriendlyName(FriendlyName).Identifier(Identifier).ProxyIdentifier(ProxyIdentifier).ProxyIdentifierSid(ProxyIdentifierSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource.
    SessionSid := "SessionSid_example" // string | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) resource.
    FailOnParticipantConflict := true // bool | [Experimental] For accounts with the ProxyAllowParticipantConflict account flag, setting to true enables per-request opt-in to allowing Proxy to reject a Participant create request that could cause the same Identifier/ProxyIdentifier pair to be active in multiple Sessions. Depending on the context, this could be a 409 error (Twilio error code 80623) or a 400 error (Twilio error code 80604). If not provided, requests will be allowed to succeed and a Debugger notification (80802) will be emitted. Having multiple, active Participants with the same Identifier/ProxyIdentifier pair causes calls and messages from affected Participants to be routed incorrectly. Please note, the default behavior for accounts without the ProxyAllowParticipantConflict flag is to reject the request as described.  This will eventually be the default for all accounts. (optional)
    FriendlyName := "FriendlyName_example" // string | The string that you assigned to describe the participant. This value must be 255 characters or fewer. **This value should not have PII.** (optional)
    Identifier := "Identifier_example" // string | The phone number of the Participant. (optional)
    ProxyIdentifier := "ProxyIdentifier_example" // string | The proxy phone number to use for the Participant. If not specified, Proxy will select a number from the pool. (optional)
    ProxyIdentifierSid := "ProxyIdentifierSid_example" // string | The SID of the Proxy Identifier to assign to the Participant. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateParticipant(context.Background(), ServiceSid, SessionSid).FailOnParticipantConflict(FailOnParticipantConflict).FriendlyName(FriendlyName).Identifier(Identifier).ProxyIdentifier(ProxyIdentifier).ProxyIdentifierSid(ProxyIdentifierSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateParticipant`: ProxyV1ServiceSessionParticipant
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource.
**SessionSid** | **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateParticipantParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **FailOnParticipantConflict** | **bool** | [Experimental] For accounts with the ProxyAllowParticipantConflict account flag, setting to true enables per-request opt-in to allowing Proxy to reject a Participant create request that could cause the same Identifier/ProxyIdentifier pair to be active in multiple Sessions. Depending on the context, this could be a 409 error (Twilio error code 80623) or a 400 error (Twilio error code 80604). If not provided, requests will be allowed to succeed and a Debugger notification (80802) will be emitted. Having multiple, active Participants with the same Identifier/ProxyIdentifier pair causes calls and messages from affected Participants to be routed incorrectly. Please note, the default behavior for accounts without the ProxyAllowParticipantConflict flag is to reject the request as described.  This will eventually be the default for all accounts.
 **FriendlyName** | **string** | The string that you assigned to describe the participant. This value must be 255 characters or fewer. **This value should not have PII.**
 **Identifier** | **string** | The phone number of the Participant.
 **ProxyIdentifier** | **string** | The proxy phone number to use for the Participant. If not specified, Proxy will select a number from the pool.
 **ProxyIdentifierSid** | **string** | The SID of the Proxy Identifier to assign to the Participant.

### Return type

[**ProxyV1ServiceSessionParticipant**](ProxyV1ServiceSessionParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePhoneNumber

> ProxyV1ServicePhoneNumber CreatePhoneNumber(ctx, ServiceSid).IsReserved(IsReserved).PhoneNumber(PhoneNumber).Sid(Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID parent [Service](https://www.twilio.com/docs/proxy/api/service) resource of the new PhoneNumber resource.
    IsReserved := true // bool | Whether the new phone number should be reserved and not be assigned to a participant using proxy pool logic. See [Reserved Phone Numbers](https://www.twilio.com/docs/proxy/reserved-phone-numbers) for more information. (optional)
    PhoneNumber := "PhoneNumber_example" // string | The phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234. (optional)
    Sid := "Sid_example" // string | The SID of a Twilio [IncomingPhoneNumber](https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource) resource that represents the Twilio Number you would like to assign to your Proxy Service. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreatePhoneNumber(context.Background(), ServiceSid).IsReserved(IsReserved).PhoneNumber(PhoneNumber).Sid(Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreatePhoneNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePhoneNumber`: ProxyV1ServicePhoneNumber
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreatePhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID parent [Service](https://www.twilio.com/docs/proxy/api/service) resource of the new PhoneNumber resource.

### Other Parameters

Other parameters are passed through a pointer to a CreatePhoneNumberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **IsReserved** | **bool** | Whether the new phone number should be reserved and not be assigned to a participant using proxy pool logic. See [Reserved Phone Numbers](https://www.twilio.com/docs/proxy/reserved-phone-numbers) for more information.
 **PhoneNumber** | **string** | The phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234.
 **Sid** | **string** | The SID of a Twilio [IncomingPhoneNumber](https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource) resource that represents the Twilio Number you would like to assign to your Proxy Service.

### Return type

[**ProxyV1ServicePhoneNumber**](ProxyV1ServicePhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateService

> ProxyV1Service CreateService(ctx).CallbackUrl(CallbackUrl).ChatInstanceSid(ChatInstanceSid).DefaultTtl(DefaultTtl).GeoMatchLevel(GeoMatchLevel).InterceptCallbackUrl(InterceptCallbackUrl).NumberSelectionBehavior(NumberSelectionBehavior).OutOfSessionCallbackUrl(OutOfSessionCallbackUrl).UniqueName(UniqueName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    CallbackUrl := "CallbackUrl_example" // string | The URL we should call when the interaction status changes. (optional)
    ChatInstanceSid := "ChatInstanceSid_example" // string | The SID of the Chat Service Instance managed by Proxy Service. The Chat Service enables Proxy to forward SMS and channel messages to this chat instance. This is a one-to-one relationship. (optional)
    DefaultTtl := int32(56) // int32 | The default `ttl` value to set for Sessions created in the Service. The TTL (time to live) is measured in seconds after the Session's last create or last Interaction. The default value of `0` indicates an unlimited Session length. You can override a Session's default TTL value by setting its `ttl` value. (optional)
    GeoMatchLevel := "GeoMatchLevel_example" // string | Where a proxy number must be located relative to the participant identifier. Can be: `country`, `area-code`, or `extended-area-code`. The default value is `country` and more specific areas than `country` are only available in North America. (optional)
    InterceptCallbackUrl := "InterceptCallbackUrl_example" // string | The URL we call on each interaction. If we receive a 403 status, we block the interaction; otherwise the interaction continues. (optional)
    NumberSelectionBehavior := "NumberSelectionBehavior_example" // string | The preference for Proxy Number selection in the Service instance. Can be: `prefer-sticky` or `avoid-sticky` and the default is `prefer-sticky`. `prefer-sticky` means that we will try and select the same Proxy Number for a given participant if they have previous [Sessions](https://www.twilio.com/docs/proxy/api/session), but we will not fail if that Proxy Number cannot be used.  `avoid-sticky` means that we will try to use different Proxy Numbers as long as that is possible within a given pool rather than try and use a previously assigned number. (optional)
    OutOfSessionCallbackUrl := "OutOfSessionCallbackUrl_example" // string | The URL we should call when an inbound call or SMS action occurs on a closed or non-existent Session. If your server (or a Twilio [function](https://www.twilio.com/functions)) responds with valid [TwiML](https://www.twilio.com/docs/voice/twiml), we will process it. This means it is possible, for example, to play a message for a call, send an automated text message response, or redirect a call to another Phone Number. See [Out-of-Session Callback Response Guide](https://www.twilio.com/docs/proxy/out-session-callback-response-guide) for more information. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. **This value should not have PII.** (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateService(context.Background()).CallbackUrl(CallbackUrl).ChatInstanceSid(ChatInstanceSid).DefaultTtl(DefaultTtl).GeoMatchLevel(GeoMatchLevel).InterceptCallbackUrl(InterceptCallbackUrl).NumberSelectionBehavior(NumberSelectionBehavior).OutOfSessionCallbackUrl(OutOfSessionCallbackUrl).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateService`: ProxyV1Service
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateService`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **CallbackUrl** | **string** | The URL we should call when the interaction status changes.
 **ChatInstanceSid** | **string** | The SID of the Chat Service Instance managed by Proxy Service. The Chat Service enables Proxy to forward SMS and channel messages to this chat instance. This is a one-to-one relationship.
 **DefaultTtl** | **int32** | The default &#x60;ttl&#x60; value to set for Sessions created in the Service. The TTL (time to live) is measured in seconds after the Session&#39;s last create or last Interaction. The default value of &#x60;0&#x60; indicates an unlimited Session length. You can override a Session&#39;s default TTL value by setting its &#x60;ttl&#x60; value.
 **GeoMatchLevel** | **string** | Where a proxy number must be located relative to the participant identifier. Can be: &#x60;country&#x60;, &#x60;area-code&#x60;, or &#x60;extended-area-code&#x60;. The default value is &#x60;country&#x60; and more specific areas than &#x60;country&#x60; are only available in North America.
 **InterceptCallbackUrl** | **string** | The URL we call on each interaction. If we receive a 403 status, we block the interaction; otherwise the interaction continues.
 **NumberSelectionBehavior** | **string** | The preference for Proxy Number selection in the Service instance. Can be: &#x60;prefer-sticky&#x60; or &#x60;avoid-sticky&#x60; and the default is &#x60;prefer-sticky&#x60;. &#x60;prefer-sticky&#x60; means that we will try and select the same Proxy Number for a given participant if they have previous [Sessions](https://www.twilio.com/docs/proxy/api/session), but we will not fail if that Proxy Number cannot be used.  &#x60;avoid-sticky&#x60; means that we will try to use different Proxy Numbers as long as that is possible within a given pool rather than try and use a previously assigned number.
 **OutOfSessionCallbackUrl** | **string** | The URL we should call when an inbound call or SMS action occurs on a closed or non-existent Session. If your server (or a Twilio [function](https://www.twilio.com/functions)) responds with valid [TwiML](https://www.twilio.com/docs/voice/twiml), we will process it. This means it is possible, for example, to play a message for a call, send an automated text message response, or redirect a call to another Phone Number. See [Out-of-Session Callback Response Guide](https://www.twilio.com/docs/proxy/out-session-callback-response-guide) for more information.
 **UniqueName** | **string** | An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. **This value should not have PII.**

### Return type

[**ProxyV1Service**](ProxyV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSession

> ProxyV1ServiceSession CreateSession(ctx, ServiceSid).DateExpiry(DateExpiry).FailOnParticipantConflict(FailOnParticipantConflict).Mode(Mode).Participants(Participants).Status(Status).Ttl(Ttl).UniqueName(UniqueName).Execute()





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
    ServiceSid := "ServiceSid_example" // string | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource.
    DateExpiry := time.Now() // time.Time | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session should expire. If this is value is present, it overrides the `ttl` value. (optional)
    FailOnParticipantConflict := true // bool | [Experimental] For accounts with the ProxyAllowParticipantConflict account flag, setting to true enables per-request opt-in to allowing Proxy to reject a Session create (with Participants) request that could cause the same Identifier/ProxyIdentifier pair to be active in multiple Sessions. Depending on the context, this could be a 409 error (Twilio error code 80623) or a 400 error (Twilio error code 80604). If not provided, requests will be allowed to succeed and a Debugger notification (80802) will be emitted. Having multiple, active Participants with the same Identifier/ProxyIdentifier pair causes calls and messages from affected Participants to be routed incorrectly. Please note, the default behavior for accounts without the ProxyAllowParticipantConflict flag is to reject the request as described.  This will eventually be the default for all accounts. (optional)
    Mode := "Mode_example" // string | The Mode of the Session. Can be: `message-only`, `voice-only`, or `voice-and-message` and the default value is `voice-and-message`. (optional)
    Participants := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | The Participant objects to include in the new session. (optional)
    Status := "Status_example" // string | The initial status of the Session. Can be: `open`, `in-progress`, `closed`, `failed`, or `unknown`. The default is `open` on create. (optional)
    Ttl := int32(56) // int32 | The time, in seconds, when the session will expire. The time is measured from the last Session create or the Session's last Interaction. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. **This value should not have PII.** (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSession(context.Background(), ServiceSid).DateExpiry(DateExpiry).FailOnParticipantConflict(FailOnParticipantConflict).Mode(Mode).Participants(Participants).Status(Status).Ttl(Ttl).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSession`: ProxyV1ServiceSession
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateSessionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **DateExpiry** | **time.Time** | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session should expire. If this is value is present, it overrides the &#x60;ttl&#x60; value.
 **FailOnParticipantConflict** | **bool** | [Experimental] For accounts with the ProxyAllowParticipantConflict account flag, setting to true enables per-request opt-in to allowing Proxy to reject a Session create (with Participants) request that could cause the same Identifier/ProxyIdentifier pair to be active in multiple Sessions. Depending on the context, this could be a 409 error (Twilio error code 80623) or a 400 error (Twilio error code 80604). If not provided, requests will be allowed to succeed and a Debugger notification (80802) will be emitted. Having multiple, active Participants with the same Identifier/ProxyIdentifier pair causes calls and messages from affected Participants to be routed incorrectly. Please note, the default behavior for accounts without the ProxyAllowParticipantConflict flag is to reject the request as described.  This will eventually be the default for all accounts.
 **Mode** | **string** | The Mode of the Session. Can be: &#x60;message-only&#x60;, &#x60;voice-only&#x60;, or &#x60;voice-and-message&#x60; and the default value is &#x60;voice-and-message&#x60;.
 **Participants** | **[]map[string]interface{}** | The Participant objects to include in the new session.
 **Status** | **string** | The initial status of the Session. Can be: &#x60;open&#x60;, &#x60;in-progress&#x60;, &#x60;closed&#x60;, &#x60;failed&#x60;, or &#x60;unknown&#x60;. The default is &#x60;open&#x60; on create.
 **Ttl** | **int32** | The time, in seconds, when the session will expire. The time is measured from the last Session create or the Session&#39;s last Interaction.
 **UniqueName** | **string** | An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. **This value should not have PII.**

### Return type

[**ProxyV1ServiceSession**](ProxyV1ServiceSession.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateShortCode

> ProxyV1ServiceShortCode CreateShortCode(ctx, ServiceSid).Sid(Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource.
    Sid := "Sid_example" // string | The SID of a Twilio [ShortCode](https://www.twilio.com/docs/sms/api/short-code) resource that represents the short code you would like to assign to your Proxy Service. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateShortCode(context.Background(), ServiceSid).Sid(Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateShortCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateShortCode`: ProxyV1ServiceShortCode
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateShortCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateShortCodeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Sid** | **string** | The SID of a Twilio [ShortCode](https://www.twilio.com/docs/sms/api/short-code) resource that represents the short code you would like to assign to your Proxy Service.

### Return type

[**ProxyV1ServiceShortCode**](ProxyV1ServiceShortCode.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInteraction

> DeleteInteraction(ctx, ServiceSid, SessionSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to delete.
    SessionSid := "SessionSid_example" // string | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resource to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Interaction resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteInteraction(context.Background(), ServiceSid, SessionSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteInteraction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to delete.
**SessionSid** | **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resource to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Interaction resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteInteractionParams struct via the builder pattern


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


## DeleteParticipant

> DeleteParticipant(ctx, ServiceSid, SessionSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to delete.
    SessionSid := "SessionSid_example" // string | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resource to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Participant resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteParticipant(context.Background(), ServiceSid, SessionSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to delete.
**SessionSid** | **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resource to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Participant resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteParticipantParams struct via the builder pattern


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


## DeletePhoneNumber

> DeletePhoneNumber(ctx, ServiceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the PhoneNumber resource to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the PhoneNumber resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeletePhoneNumber(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeletePhoneNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the PhoneNumber resource to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the PhoneNumber resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeletePhoneNumberParams struct via the builder pattern


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


## DeleteSession

> DeleteSession(ctx, ServiceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Session resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSession(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Session resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSessionParams struct via the builder pattern


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


## DeleteShortCode

> DeleteShortCode(ctx, ServiceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource to delete the ShortCode resource from.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the ShortCode resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteShortCode(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteShortCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource to delete the ShortCode resource from.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ShortCode resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteShortCodeParams struct via the builder pattern


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


## FetchInteraction

> ProxyV1ServiceSessionInteraction FetchInteraction(ctx, ServiceSid, SessionSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to fetch.
    SessionSid := "SessionSid_example" // string | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Interaction resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchInteraction(context.Background(), ServiceSid, SessionSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchInteraction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchInteraction`: ProxyV1ServiceSessionInteraction
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchInteraction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to fetch.
**SessionSid** | **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Interaction resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchInteractionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ProxyV1ServiceSessionInteraction**](ProxyV1ServiceSessionInteraction.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMessageInteraction

> ProxyV1ServiceSessionParticipantMessageInteraction FetchMessageInteraction(ctx, ServiceSid, SessionSid, ParticipantSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to fetch.
    SessionSid := "SessionSid_example" // string | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resource to fetch.
    ParticipantSid := "ParticipantSid_example" // string | The SID of the [Participant](https://www.twilio.com/docs/proxy/api/participant) resource.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the MessageInteraction resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchMessageInteraction(context.Background(), ServiceSid, SessionSid, ParticipantSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchMessageInteraction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchMessageInteraction`: ProxyV1ServiceSessionParticipantMessageInteraction
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchMessageInteraction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to fetch.
**SessionSid** | **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resource to fetch.
**ParticipantSid** | **string** | The SID of the [Participant](https://www.twilio.com/docs/proxy/api/participant) resource.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the MessageInteraction resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchMessageInteractionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------





### Return type

[**ProxyV1ServiceSessionParticipantMessageInteraction**](ProxyV1ServiceSessionParticipantMessageInteraction.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchParticipant

> ProxyV1ServiceSessionParticipant FetchParticipant(ctx, ServiceSid, SessionSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to fetch.
    SessionSid := "SessionSid_example" // string | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Participant resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchParticipant(context.Background(), ServiceSid, SessionSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchParticipant`: ProxyV1ServiceSessionParticipant
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to fetch.
**SessionSid** | **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Participant resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchParticipantParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ProxyV1ServiceSessionParticipant**](ProxyV1ServiceSessionParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPhoneNumber

> ProxyV1ServicePhoneNumber FetchPhoneNumber(ctx, ServiceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the PhoneNumber resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the PhoneNumber resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchPhoneNumber(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchPhoneNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchPhoneNumber`: ProxyV1ServicePhoneNumber
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchPhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the PhoneNumber resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the PhoneNumber resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchPhoneNumberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ProxyV1ServicePhoneNumber**](ProxyV1ServicePhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> ProxyV1Service FetchService(ctx, Sid).Execute()





### Example

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
    // response from `FetchService`: ProxyV1Service
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

[**ProxyV1Service**](ProxyV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSession

> ProxyV1ServiceSession FetchSession(ctx, ServiceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Session resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSession(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSession`: ProxyV1ServiceSession
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Session resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSessionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ProxyV1ServiceSession**](ProxyV1ServiceSession.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchShortCode

> ProxyV1ServiceShortCode FetchShortCode(ctx, ServiceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) to fetch the resource from.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the ShortCode resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchShortCode(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchShortCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchShortCode`: ProxyV1ServiceShortCode
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchShortCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) to fetch the resource from.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ShortCode resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchShortCodeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ProxyV1ServiceShortCode**](ProxyV1ServiceShortCode.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInteraction

> ListInteractionResponse ListInteraction(ctx, ServiceSid, SessionSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) to read the resources from.
    SessionSid := "SessionSid_example" // string | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) to read the resources from.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListInteraction(context.Background(), ServiceSid, SessionSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListInteraction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInteraction`: ListInteractionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListInteraction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) to read the resources from.
**SessionSid** | **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) to read the resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListInteractionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListInteractionResponse**](ListInteractionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessageInteraction

> ListMessageInteractionResponse ListMessageInteraction(ctx, ServiceSid, SessionSid, ParticipantSid).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) to read the resources from.
    SessionSid := "SessionSid_example" // string | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) to read the resources from.
    ParticipantSid := "ParticipantSid_example" // string | The SID of the [Participant](https://www.twilio.com/docs/proxy/api/participant) to read the resources from.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListMessageInteraction(context.Background(), ServiceSid, SessionSid, ParticipantSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListMessageInteraction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMessageInteraction`: ListMessageInteractionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListMessageInteraction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) to read the resources from.
**SessionSid** | **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) to read the resources from.
**ParticipantSid** | **string** | The SID of the [Participant](https://www.twilio.com/docs/proxy/api/participant) to read the resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListMessageInteractionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListMessageInteractionResponse**](ListMessageInteractionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListParticipant

> ListParticipantResponse ListParticipant(ctx, ServiceSid, SessionSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resources to read.
    SessionSid := "SessionSid_example" // string | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListParticipant(context.Background(), ServiceSid, SessionSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListParticipant`: ListParticipantResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resources to read.
**SessionSid** | **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListParticipantParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListParticipantResponse**](ListParticipantResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPhoneNumber

> ListPhoneNumberResponse ListPhoneNumber(ctx, ServiceSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the PhoneNumber resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListPhoneNumber(context.Background(), ServiceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListPhoneNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPhoneNumber`: ListPhoneNumberResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListPhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the PhoneNumber resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListPhoneNumberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListPhoneNumberResponse**](ListPhoneNumberResponse.md)

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


## ListSession

> ListSessionResponse ListSession(ctx, ServiceSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSession(context.Background(), ServiceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSession`: ListSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to read.

### Other Parameters

Other parameters are passed through a pointer to a ListSessionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSessionResponse**](ListSessionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListShortCode

> ListShortCodeResponse ListShortCode(ctx, ServiceSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) to read the resources from.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListShortCode(context.Background(), ServiceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListShortCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListShortCode`: ListShortCodeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListShortCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) to read the resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListShortCodeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListShortCodeResponse**](ListShortCodeResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePhoneNumber

> ProxyV1ServicePhoneNumber UpdatePhoneNumber(ctx, ServiceSid, Sid).IsReserved(IsReserved).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the PhoneNumber resource to update.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the PhoneNumber resource to update.
    IsReserved := true // bool | Whether the phone number should be reserved and not be assigned to a participant using proxy pool logic. See [Reserved Phone Numbers](https://www.twilio.com/docs/proxy/reserved-phone-numbers) for more information. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdatePhoneNumber(context.Background(), ServiceSid, Sid).IsReserved(IsReserved).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdatePhoneNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePhoneNumber`: ProxyV1ServicePhoneNumber
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdatePhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the PhoneNumber resource to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the PhoneNumber resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdatePhoneNumberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **IsReserved** | **bool** | Whether the phone number should be reserved and not be assigned to a participant using proxy pool logic. See [Reserved Phone Numbers](https://www.twilio.com/docs/proxy/reserved-phone-numbers) for more information.

### Return type

[**ProxyV1ServicePhoneNumber**](ProxyV1ServicePhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> ProxyV1Service UpdateService(ctx, Sid).CallbackUrl(CallbackUrl).ChatInstanceSid(ChatInstanceSid).DefaultTtl(DefaultTtl).GeoMatchLevel(GeoMatchLevel).InterceptCallbackUrl(InterceptCallbackUrl).NumberSelectionBehavior(NumberSelectionBehavior).OutOfSessionCallbackUrl(OutOfSessionCallbackUrl).UniqueName(UniqueName).Execute()





### Example

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
    CallbackUrl := "CallbackUrl_example" // string | The URL we should call when the interaction status changes. (optional)
    ChatInstanceSid := "ChatInstanceSid_example" // string | The SID of the Chat Service Instance managed by Proxy Service. The Chat Service enables Proxy to forward SMS and channel messages to this chat instance. This is a one-to-one relationship. (optional)
    DefaultTtl := int32(56) // int32 | The default `ttl` value to set for Sessions created in the Service. The TTL (time to live) is measured in seconds after the Session's last create or last Interaction. The default value of `0` indicates an unlimited Session length. You can override a Session's default TTL value by setting its `ttl` value. (optional)
    GeoMatchLevel := "GeoMatchLevel_example" // string | Where a proxy number must be located relative to the participant identifier. Can be: `country`, `area-code`, or `extended-area-code`. The default value is `country` and more specific areas than `country` are only available in North America. (optional)
    InterceptCallbackUrl := "InterceptCallbackUrl_example" // string | The URL we call on each interaction. If we receive a 403 status, we block the interaction; otherwise the interaction continues. (optional)
    NumberSelectionBehavior := "NumberSelectionBehavior_example" // string | The preference for Proxy Number selection in the Service instance. Can be: `prefer-sticky` or `avoid-sticky` and the default is `prefer-sticky`. `prefer-sticky` means that we will try and select the same Proxy Number for a given participant if they have previous [Sessions](https://www.twilio.com/docs/proxy/api/session), but we will not fail if that Proxy Number cannot be used.  `avoid-sticky` means that we will try to use different Proxy Numbers as long as that is possible within a given pool rather than try and use a previously assigned number. (optional)
    OutOfSessionCallbackUrl := "OutOfSessionCallbackUrl_example" // string | The URL we should call when an inbound call or SMS action occurs on a closed or non-existent Session. If your server (or a Twilio [function](https://www.twilio.com/functions)) responds with valid [TwiML](https://www.twilio.com/docs/voice/twiml), we will process it. This means it is possible, for example, to play a message for a call, send an automated text message response, or redirect a call to another Phone Number. See [Out-of-Session Callback Response Guide](https://www.twilio.com/docs/proxy/out-session-callback-response-guide) for more information. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. **This value should not have PII.** (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateService(context.Background(), Sid).CallbackUrl(CallbackUrl).ChatInstanceSid(ChatInstanceSid).DefaultTtl(DefaultTtl).GeoMatchLevel(GeoMatchLevel).InterceptCallbackUrl(InterceptCallbackUrl).NumberSelectionBehavior(NumberSelectionBehavior).OutOfSessionCallbackUrl(OutOfSessionCallbackUrl).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateService`: ProxyV1Service
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

 **CallbackUrl** | **string** | The URL we should call when the interaction status changes.
 **ChatInstanceSid** | **string** | The SID of the Chat Service Instance managed by Proxy Service. The Chat Service enables Proxy to forward SMS and channel messages to this chat instance. This is a one-to-one relationship.
 **DefaultTtl** | **int32** | The default &#x60;ttl&#x60; value to set for Sessions created in the Service. The TTL (time to live) is measured in seconds after the Session&#39;s last create or last Interaction. The default value of &#x60;0&#x60; indicates an unlimited Session length. You can override a Session&#39;s default TTL value by setting its &#x60;ttl&#x60; value.
 **GeoMatchLevel** | **string** | Where a proxy number must be located relative to the participant identifier. Can be: &#x60;country&#x60;, &#x60;area-code&#x60;, or &#x60;extended-area-code&#x60;. The default value is &#x60;country&#x60; and more specific areas than &#x60;country&#x60; are only available in North America.
 **InterceptCallbackUrl** | **string** | The URL we call on each interaction. If we receive a 403 status, we block the interaction; otherwise the interaction continues.
 **NumberSelectionBehavior** | **string** | The preference for Proxy Number selection in the Service instance. Can be: &#x60;prefer-sticky&#x60; or &#x60;avoid-sticky&#x60; and the default is &#x60;prefer-sticky&#x60;. &#x60;prefer-sticky&#x60; means that we will try and select the same Proxy Number for a given participant if they have previous [Sessions](https://www.twilio.com/docs/proxy/api/session), but we will not fail if that Proxy Number cannot be used.  &#x60;avoid-sticky&#x60; means that we will try to use different Proxy Numbers as long as that is possible within a given pool rather than try and use a previously assigned number.
 **OutOfSessionCallbackUrl** | **string** | The URL we should call when an inbound call or SMS action occurs on a closed or non-existent Session. If your server (or a Twilio [function](https://www.twilio.com/functions)) responds with valid [TwiML](https://www.twilio.com/docs/voice/twiml), we will process it. This means it is possible, for example, to play a message for a call, send an automated text message response, or redirect a call to another Phone Number. See [Out-of-Session Callback Response Guide](https://www.twilio.com/docs/proxy/out-session-callback-response-guide) for more information.
 **UniqueName** | **string** | An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. **This value should not have PII.**

### Return type

[**ProxyV1Service**](ProxyV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSession

> ProxyV1ServiceSession UpdateSession(ctx, ServiceSid, Sid).DateExpiry(DateExpiry).FailOnParticipantConflict(FailOnParticipantConflict).Status(Status).Ttl(Ttl).Execute()





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
    ServiceSid := "ServiceSid_example" // string | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to update.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Session resource to update.
    DateExpiry := time.Now() // time.Time | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session should expire. If this is value is present, it overrides the `ttl` value. (optional)
    FailOnParticipantConflict := true // bool | [Experimental] For accounts with the ProxyAllowParticipantConflict account flag, setting to true enables per-request opt-in to allowing Proxy to return a 400 error (Twilio error code 80604) when a request to set a Session to in-progress would cause Participants with the same Identifier/ProxyIdentifier pair to be active in multiple Sessions. If not provided, requests will be allowed to succeed, and a Debugger notification (80801) will be emitted. Having multiple, active Participants with the same Identifier/ProxyIdentifier pair causes calls and messages from affected Participants to be routed incorrectly. Please note, the default behavior for accounts without the ProxyAllowParticipantConflict flag is to reject the request as described.  This will eventually be the default for all accounts. (optional)
    Status := "Status_example" // string | The new status of the resource. Can be: `in-progress` to re-open a session or `closed` to close a session. (optional)
    Ttl := int32(56) // int32 | The time, in seconds, when the session will expire. The time is measured from the last Session create or the Session's last Interaction. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSession(context.Background(), ServiceSid, Sid).DateExpiry(DateExpiry).FailOnParticipantConflict(FailOnParticipantConflict).Status(Status).Ttl(Ttl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSession`: ProxyV1ServiceSession
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Session resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSessionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **DateExpiry** | **time.Time** | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session should expire. If this is value is present, it overrides the &#x60;ttl&#x60; value.
 **FailOnParticipantConflict** | **bool** | [Experimental] For accounts with the ProxyAllowParticipantConflict account flag, setting to true enables per-request opt-in to allowing Proxy to return a 400 error (Twilio error code 80604) when a request to set a Session to in-progress would cause Participants with the same Identifier/ProxyIdentifier pair to be active in multiple Sessions. If not provided, requests will be allowed to succeed, and a Debugger notification (80801) will be emitted. Having multiple, active Participants with the same Identifier/ProxyIdentifier pair causes calls and messages from affected Participants to be routed incorrectly. Please note, the default behavior for accounts without the ProxyAllowParticipantConflict flag is to reject the request as described.  This will eventually be the default for all accounts.
 **Status** | **string** | The new status of the resource. Can be: &#x60;in-progress&#x60; to re-open a session or &#x60;closed&#x60; to close a session.
 **Ttl** | **int32** | The time, in seconds, when the session will expire. The time is measured from the last Session create or the Session&#39;s last Interaction.

### Return type

[**ProxyV1ServiceSession**](ProxyV1ServiceSession.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShortCode

> ProxyV1ServiceShortCode UpdateShortCode(ctx, ServiceSid, Sid).IsReserved(IsReserved).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to update.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the ShortCode resource to update.
    IsReserved := true // bool | Whether the short code should be reserved and not be assigned to a participant using proxy pool logic. See [Reserved Phone Numbers](https://www.twilio.com/docs/proxy/reserved-phone-numbers) for more information. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateShortCode(context.Background(), ServiceSid, Sid).IsReserved(IsReserved).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateShortCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateShortCode`: ProxyV1ServiceShortCode
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateShortCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ShortCode resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateShortCodeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **IsReserved** | **bool** | Whether the short code should be reserved and not be assigned to a participant using proxy pool logic. See [Reserved Phone Numbers](https://www.twilio.com/docs/proxy/reserved-phone-numbers) for more information.

### Return type

[**ProxyV1ServiceShortCode**](ProxyV1ServiceShortCode.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

