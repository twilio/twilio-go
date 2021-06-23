# DefaultApi

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

> ProxyV1ServiceSessionParticipantMessageInteraction CreateMessageInteraction(ctx, ServiceSidSessionSidParticipantSidoptional)



Create a new message Interaction to send directly from your system to one [Participant](https://www.twilio.com/docs/proxy/api/participant).  The `inbound` properties for the Interaction will always be empty.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource.
**SessionSid** | **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) resource.
**ParticipantSid** | **string** | The SID of the [Participant](https://www.twilio.com/docs/proxy/api/participant) resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateMessageInteractionParams struct


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

> ProxyV1ServiceSessionParticipant CreateParticipant(ctx, ServiceSidSessionSidoptional)



Add a new Participant to the Session

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource.
**SessionSid** | **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateParticipantParams struct


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

> ProxyV1ServicePhoneNumber CreatePhoneNumber(ctx, ServiceSidoptional)



Add a Phone Number to a Service's Proxy Number Pool.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID parent [Service](https://www.twilio.com/docs/proxy/api/service) resource of the new PhoneNumber resource.

### Other Parameters

Other parameters are passed through a pointer to a CreatePhoneNumberParams struct


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

> ProxyV1Service CreateService(ctx, optional)



Create a new Service for Twilio Proxy

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**CallbackUrl** | **string** | The URL we should call when the interaction status changes.
**ChatInstanceSid** | **string** | The SID of the Chat Service Instance managed by Proxy Service. The Chat Service enables Proxy to forward SMS and channel messages to this chat instance. This is a one-to-one relationship.
**DefaultTtl** | **int** | The default &#x60;ttl&#x60; value to set for Sessions created in the Service. The TTL (time to live) is measured in seconds after the Session&#39;s last create or last Interaction. The default value of &#x60;0&#x60; indicates an unlimited Session length. You can override a Session&#39;s default TTL value by setting its &#x60;ttl&#x60; value.
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

> ProxyV1ServiceSession CreateSession(ctx, ServiceSidoptional)



Create a new Session

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateSessionParams struct


Name | Type | Description
------------- | ------------- | -------------
**DateExpiry** | **time.Time** | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session should expire. If this is value is present, it overrides the &#x60;ttl&#x60; value.
**FailOnParticipantConflict** | **bool** | [Experimental] For accounts with the ProxyAllowParticipantConflict account flag, setting to true enables per-request opt-in to allowing Proxy to reject a Session create (with Participants) request that could cause the same Identifier/ProxyIdentifier pair to be active in multiple Sessions. Depending on the context, this could be a 409 error (Twilio error code 80623) or a 400 error (Twilio error code 80604). If not provided, requests will be allowed to succeed and a Debugger notification (80802) will be emitted. Having multiple, active Participants with the same Identifier/ProxyIdentifier pair causes calls and messages from affected Participants to be routed incorrectly. Please note, the default behavior for accounts without the ProxyAllowParticipantConflict flag is to reject the request as described.  This will eventually be the default for all accounts.
**Mode** | **string** | The Mode of the Session. Can be: &#x60;message-only&#x60;, &#x60;voice-only&#x60;, or &#x60;voice-and-message&#x60; and the default value is &#x60;voice-and-message&#x60;.
**Participants** | **[]map[string]interface{}** | The Participant objects to include in the new session.
**Status** | **string** | The initial status of the Session. Can be: &#x60;open&#x60;, &#x60;in-progress&#x60;, &#x60;closed&#x60;, &#x60;failed&#x60;, or &#x60;unknown&#x60;. The default is &#x60;open&#x60; on create.
**Ttl** | **int** | The time, in seconds, when the session will expire. The time is measured from the last Session create or the Session&#39;s last Interaction.
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

> ProxyV1ServiceShortCode CreateShortCode(ctx, ServiceSidoptional)



Add a Short Code to the Proxy Number Pool for the Service.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateShortCodeParams struct


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

> DeleteInteraction(ctx, ServiceSidSessionSidSid)



Delete a specific Interaction.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to delete.
**SessionSid** | **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resource to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Interaction resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteInteractionParams struct


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

> DeleteParticipant(ctx, ServiceSidSessionSidSid)



Delete a specific Participant. This is a soft-delete. The participant remains associated with the session and cannot be re-added. Participants are only permanently deleted when the [Session](https://www.twilio.com/docs/proxy/api/session) is deleted.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to delete.
**SessionSid** | **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resource to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Participant resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteParticipantParams struct


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

> DeletePhoneNumber(ctx, ServiceSidSid)



Delete a specific Phone Number from a Service.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the PhoneNumber resource to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the PhoneNumber resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeletePhoneNumberParams struct


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

> DeleteService(ctx, Sid)



Delete a specific Service.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Service resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceParams struct


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

> DeleteSession(ctx, ServiceSidSid)



Delete a specific Session.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Session resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSessionParams struct


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

> DeleteShortCode(ctx, ServiceSidSid)



Delete a specific Short Code from a Service.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource to delete the ShortCode resource from.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ShortCode resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteShortCodeParams struct


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

> ProxyV1ServiceSessionInteraction FetchInteraction(ctx, ServiceSidSessionSidSid)



Retrieve a list of Interactions for a given [Session](https://www.twilio.com/docs/proxy/api/session).

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to fetch.
**SessionSid** | **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Interaction resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchInteractionParams struct


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

> ProxyV1ServiceSessionParticipantMessageInteraction FetchMessageInteraction(ctx, ServiceSidSessionSidParticipantSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to fetch.
**SessionSid** | **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resource to fetch.
**ParticipantSid** | **string** | The SID of the [Participant](https://www.twilio.com/docs/proxy/api/participant) resource.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the MessageInteraction resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchMessageInteractionParams struct


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

> ProxyV1ServiceSessionParticipant FetchParticipant(ctx, ServiceSidSessionSidSid)



Fetch a specific Participant.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to fetch.
**SessionSid** | **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Participant resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchParticipantParams struct


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

> ProxyV1ServicePhoneNumber FetchPhoneNumber(ctx, ServiceSidSid)



Fetch a specific Phone Number.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the PhoneNumber resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the PhoneNumber resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchPhoneNumberParams struct


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

> ProxyV1Service FetchService(ctx, Sid)



Fetch a specific Service.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Service resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceParams struct


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

> ProxyV1ServiceSession FetchSession(ctx, ServiceSidSid)



Fetch a specific Session.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Session resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSessionParams struct


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

> ProxyV1ServiceShortCode FetchShortCode(ctx, ServiceSidSid)



Fetch a specific Short Code.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) to fetch the resource from.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ShortCode resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchShortCodeParams struct


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

> ListInteractionResponse ListInteraction(ctx, ServiceSidSessionSidoptional)



Retrieve a list of all Interactions for a Session. A maximum of 100 records will be returned per page.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) to read the resources from.
**SessionSid** | **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) to read the resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListInteractionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListMessageInteractionResponse ListMessageInteraction(ctx, ServiceSidSessionSidParticipantSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) to read the resources from.
**SessionSid** | **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) to read the resources from.
**ParticipantSid** | **string** | The SID of the [Participant](https://www.twilio.com/docs/proxy/api/participant) to read the resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListMessageInteractionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListParticipantResponse ListParticipant(ctx, ServiceSidSessionSidoptional)



Retrieve a list of all Participants in a Session.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resources to read.
**SessionSid** | **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListParticipantParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListPhoneNumberResponse ListPhoneNumber(ctx, ServiceSidoptional)



Retrieve a list of all Phone Numbers in the Proxy Number Pool for a Service. A maximum of 100 records will be returned per page.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the PhoneNumber resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListPhoneNumberParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListServiceResponse ListService(ctx, optional)



Retrieve a list of all Services for Twilio Proxy. A maximum of 100 records will be returned per page.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListSessionResponse ListSession(ctx, ServiceSidoptional)



Retrieve a list of all Sessions for the Service. A maximum of 100 records will be returned per page.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to read.

### Other Parameters

Other parameters are passed through a pointer to a ListSessionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListShortCodeResponse ListShortCode(ctx, ServiceSidoptional)



Retrieve a list of all Short Codes in the Proxy Number Pool for the Service. A maximum of 100 records will be returned per page.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) to read the resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListShortCodeParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ProxyV1ServicePhoneNumber UpdatePhoneNumber(ctx, ServiceSidSidoptional)



Update a specific Proxy Number.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the PhoneNumber resource to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the PhoneNumber resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdatePhoneNumberParams struct


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

> ProxyV1Service UpdateService(ctx, Sidoptional)



Update a specific Service.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Service resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**CallbackUrl** | **string** | The URL we should call when the interaction status changes.
**ChatInstanceSid** | **string** | The SID of the Chat Service Instance managed by Proxy Service. The Chat Service enables Proxy to forward SMS and channel messages to this chat instance. This is a one-to-one relationship.
**DefaultTtl** | **int** | The default &#x60;ttl&#x60; value to set for Sessions created in the Service. The TTL (time to live) is measured in seconds after the Session&#39;s last create or last Interaction. The default value of &#x60;0&#x60; indicates an unlimited Session length. You can override a Session&#39;s default TTL value by setting its &#x60;ttl&#x60; value.
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

> ProxyV1ServiceSession UpdateSession(ctx, ServiceSidSidoptional)



Update a specific Session.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Session resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSessionParams struct


Name | Type | Description
------------- | ------------- | -------------
**DateExpiry** | **time.Time** | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session should expire. If this is value is present, it overrides the &#x60;ttl&#x60; value.
**FailOnParticipantConflict** | **bool** | [Experimental] For accounts with the ProxyAllowParticipantConflict account flag, setting to true enables per-request opt-in to allowing Proxy to return a 400 error (Twilio error code 80604) when a request to set a Session to in-progress would cause Participants with the same Identifier/ProxyIdentifier pair to be active in multiple Sessions. If not provided, requests will be allowed to succeed, and a Debugger notification (80801) will be emitted. Having multiple, active Participants with the same Identifier/ProxyIdentifier pair causes calls and messages from affected Participants to be routed incorrectly. Please note, the default behavior for accounts without the ProxyAllowParticipantConflict flag is to reject the request as described.  This will eventually be the default for all accounts.
**Status** | **string** | The new status of the resource. Can be: &#x60;in-progress&#x60; to re-open a session or &#x60;closed&#x60; to close a session.
**Ttl** | **int** | The time, in seconds, when the session will expire. The time is measured from the last Session create or the Session&#39;s last Interaction.

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

> ProxyV1ServiceShortCode UpdateShortCode(ctx, ServiceSidSidoptional)



Update a specific Short Code.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ShortCode resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateShortCodeParams struct


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

