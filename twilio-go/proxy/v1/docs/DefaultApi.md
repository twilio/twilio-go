# \DefaultApi

All URIs are relative to *http://localhost*

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

> ProxyV1ServiceSessionParticipantMessageInteraction CreateMessageInteraction(ctx, serviceSid, sessionSid, participantSid, optional)



Create a new message Interaction to send directly from your system to one [Participant](https://www.twilio.com/docs/proxy/api/participant).  The `inbound` properties for the Interaction will always be empty.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource. | 
**sessionSid** | **string**| The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) resource. | 
**participantSid** | **string**| The SID of the [Participant](https://www.twilio.com/docs/proxy/api/participant) resource. | 
 **optional** | ***CreateMessageInteractionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateMessageInteractionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **optional.String**| The message to send to the participant | 
 **mediaUrl** | [**optional.Interface of []string**](string.md)| Reserved. Not currently supported. | 

### Return type

[**ProxyV1ServiceSessionParticipantMessageInteraction**](proxy.v1.service.session.participant.message_interaction.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateParticipant

> ProxyV1ServiceSessionParticipant CreateParticipant(ctx, serviceSid, sessionSid, optional)



Add a new Participant to the Session

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource. | 
**sessionSid** | **string**| The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) resource. | 
 **optional** | ***CreateParticipantOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateParticipantOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **failOnParticipantConflict** | **optional.Bool**| [Experimental] For accounts with the ProxyAllowParticipantConflict account flag, setting to true enables per-request opt-in to allowing Proxy to reject a Participant create request that could cause the same Identifier/ProxyIdentifier pair to be active in multiple Sessions. Depending on the context, this could be a 409 error (Twilio error code 80623) or a 400 error (Twilio error code 80604). If not provided, requests will be allowed to succeed and a Debugger notification (80802) will be emitted. Having multiple, active Participants with the same Identifier/ProxyIdentifier pair causes calls and messages from affected Participants to be routed incorrectly. Please note, the default behavior for accounts without the ProxyAllowParticipantConflict flag is to reject the request as described.  This will eventually be the default for all accounts. | 
 **friendlyName** | **optional.String**| The string that you assigned to describe the participant. This value must be 255 characters or fewer. **This value should not have PII.** | 
 **identifier** | **optional.String**| The phone number of the Participant. | 
 **proxyIdentifier** | **optional.String**| The proxy phone number to use for the Participant. If not specified, Proxy will select a number from the pool. | 
 **proxyIdentifierSid** | **optional.String**| The SID of the Proxy Identifier to assign to the Participant. | 

### Return type

[**ProxyV1ServiceSessionParticipant**](proxy.v1.service.session.participant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePhoneNumber

> ProxyV1ServicePhoneNumber CreatePhoneNumber(ctx, serviceSid, optional)



Add a Phone Number to a Service's Proxy Number Pool.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID parent [Service](https://www.twilio.com/docs/proxy/api/service) resource of the new PhoneNumber resource. | 
 **optional** | ***CreatePhoneNumberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreatePhoneNumberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isReserved** | **optional.Bool**| Whether the new phone number should be reserved and not be assigned to a participant using proxy pool logic. See [Reserved Phone Numbers](https://www.twilio.com/docs/proxy/reserved-phone-numbers) for more information. | 
 **phoneNumber** | **optional.String**| The phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234. | 
 **sid** | **optional.String**| The SID of a Twilio [IncomingPhoneNumber](https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource) resource that represents the Twilio Number you would like to assign to your Proxy Service. | 

### Return type

[**ProxyV1ServicePhoneNumber**](proxy.v1.service.phone_number.md)

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

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **callbackUrl** | **optional.String**| The URL we should call when the interaction status changes. | 
 **chatInstanceSid** | **optional.String**| The SID of the Chat Service Instance managed by Proxy Service. The Chat Service enables Proxy to forward SMS and channel messages to this chat instance. This is a one-to-one relationship. | 
 **defaultTtl** | **optional.Int32**| The default &#x60;ttl&#x60; value to set for Sessions created in the Service. The TTL (time to live) is measured in seconds after the Session&#39;s last create or last Interaction. The default value of &#x60;0&#x60; indicates an unlimited Session length. You can override a Session&#39;s default TTL value by setting its &#x60;ttl&#x60; value. | 
 **geoMatchLevel** | **optional.String**| Where a proxy number must be located relative to the participant identifier. Can be: &#x60;country&#x60;, &#x60;area-code&#x60;, or &#x60;extended-area-code&#x60;. The default value is &#x60;country&#x60; and more specific areas than &#x60;country&#x60; are only available in North America. | 
 **interceptCallbackUrl** | **optional.String**| The URL we call on each interaction. If we receive a 403 status, we block the interaction; otherwise the interaction continues. | 
 **numberSelectionBehavior** | **optional.String**| The preference for Proxy Number selection in the Service instance. Can be: &#x60;prefer-sticky&#x60; or &#x60;avoid-sticky&#x60; and the default is &#x60;prefer-sticky&#x60;. &#x60;prefer-sticky&#x60; means that we will try and select the same Proxy Number for a given participant if they have previous [Sessions](https://www.twilio.com/docs/proxy/api/session), but we will not fail if that Proxy Number cannot be used.  &#x60;avoid-sticky&#x60; means that we will try to use different Proxy Numbers as long as that is possible within a given pool rather than try and use a previously assigned number. | 
 **outOfSessionCallbackUrl** | **optional.String**| The URL we should call when an inbound call or SMS action occurs on a closed or non-existent Session. If your server (or a Twilio [function](https://www.twilio.com/functions)) responds with valid [TwiML](https://www.twilio.com/docs/voice/twiml), we will process it. This means it is possible, for example, to play a message for a call, send an automated text message response, or redirect a call to another Phone Number. See [Out-of-Session Callback Response Guide](https://www.twilio.com/docs/proxy/out-session-callback-response-guide) for more information. | 
 **uniqueName** | **optional.String**| An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. **This value should not have PII.** | 

### Return type

[**ProxyV1Service**](proxy.v1.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSession

> ProxyV1ServiceSession CreateSession(ctx, serviceSid, optional)



Create a new Session

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource. | 
 **optional** | ***CreateSessionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSessionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dateExpiry** | **optional.Time**| The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session should expire. If this is value is present, it overrides the &#x60;ttl&#x60; value. | 
 **failOnParticipantConflict** | **optional.Bool**| [Experimental] For accounts with the ProxyAllowParticipantConflict account flag, setting to true enables per-request opt-in to allowing Proxy to reject a Session create (with Participants) request that could cause the same Identifier/ProxyIdentifier pair to be active in multiple Sessions. Depending on the context, this could be a 409 error (Twilio error code 80623) or a 400 error (Twilio error code 80604). If not provided, requests will be allowed to succeed and a Debugger notification (80802) will be emitted. Having multiple, active Participants with the same Identifier/ProxyIdentifier pair causes calls and messages from affected Participants to be routed incorrectly. Please note, the default behavior for accounts without the ProxyAllowParticipantConflict flag is to reject the request as described.  This will eventually be the default for all accounts. | 
 **mode** | **optional.String**| The Mode of the Session. Can be: &#x60;message-only&#x60;, &#x60;voice-only&#x60;, or &#x60;voice-and-message&#x60; and the default value is &#x60;voice-and-message&#x60;. | 
 **participants** | [**optional.Interface of []map[string]interface{}**](map[string]interface{}.md)| The Participant objects to include in the new session. | 
 **status** | **optional.String**| The initial status of the Session. Can be: &#x60;open&#x60;, &#x60;in-progress&#x60;, &#x60;closed&#x60;, &#x60;failed&#x60;, or &#x60;unknown&#x60;. The default is &#x60;open&#x60; on create. | 
 **ttl** | **optional.Int32**| The time, in seconds, when the session will expire. The time is measured from the last Session create or the Session&#39;s last Interaction. | 
 **uniqueName** | **optional.String**| An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. **This value should not have PII.** | 

### Return type

[**ProxyV1ServiceSession**](proxy.v1.service.session.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateShortCode

> ProxyV1ServiceShortCode CreateShortCode(ctx, serviceSid, optional)



Add a Short Code to the Proxy Number Pool for the Service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource. | 
 **optional** | ***CreateShortCodeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateShortCodeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sid** | **optional.String**| The SID of a Twilio [ShortCode](https://www.twilio.com/docs/sms/api/short-code) resource that represents the short code you would like to assign to your Proxy Service. | 

### Return type

[**ProxyV1ServiceShortCode**](proxy.v1.service.short_code.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInteraction

> DeleteInteraction(ctx, serviceSid, sessionSid, sid)



Delete a specific Interaction.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to delete. | 
**sessionSid** | **string**| The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resource to delete. | 
**sid** | **string**| The Twilio-provided string that uniquely identifies the Interaction resource to delete. | 

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

> DeleteParticipant(ctx, serviceSid, sessionSid, sid)



Delete a specific Participant. This is a soft-delete. The participant remains associated with the session and cannot be re-added. Participants are only permanently deleted when the [Session](https://www.twilio.com/docs/proxy/api/session) is deleted.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to delete. | 
**sessionSid** | **string**| The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resource to delete. | 
**sid** | **string**| The Twilio-provided string that uniquely identifies the Participant resource to delete. | 

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

> DeletePhoneNumber(ctx, serviceSid, sid)



Delete a specific Phone Number from a Service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the PhoneNumber resource to delete. | 
**sid** | **string**| The Twilio-provided string that uniquely identifies the PhoneNumber resource to delete. | 

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

> DeleteService(ctx, sid)



Delete a specific Service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The Twilio-provided string that uniquely identifies the Service resource to delete. | 

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

> DeleteSession(ctx, serviceSid, sid)



Delete a specific Session.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to delete. | 
**sid** | **string**| The Twilio-provided string that uniquely identifies the Session resource to delete. | 

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

> DeleteShortCode(ctx, serviceSid, sid)



Delete a specific Short Code from a Service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource to delete the ShortCode resource from. | 
**sid** | **string**| The Twilio-provided string that uniquely identifies the ShortCode resource to delete. | 

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

> ProxyV1ServiceSessionInteraction FetchInteraction(ctx, serviceSid, sessionSid, sid)



Retrieve a list of Interactions for a given [Session](https://www.twilio.com/docs/proxy/api/session).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to fetch. | 
**sessionSid** | **string**| The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resource to fetch. | 
**sid** | **string**| The Twilio-provided string that uniquely identifies the Interaction resource to fetch. | 

### Return type

[**ProxyV1ServiceSessionInteraction**](proxy.v1.service.session.interaction.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMessageInteraction

> ProxyV1ServiceSessionParticipantMessageInteraction FetchMessageInteraction(ctx, serviceSid, sessionSid, participantSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to fetch. | 
**sessionSid** | **string**| The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resource to fetch. | 
**participantSid** | **string**| The SID of the [Participant](https://www.twilio.com/docs/proxy/api/participant) resource. | 
**sid** | **string**| The Twilio-provided string that uniquely identifies the MessageInteraction resource to fetch. | 

### Return type

[**ProxyV1ServiceSessionParticipantMessageInteraction**](proxy.v1.service.session.participant.message_interaction.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchParticipant

> ProxyV1ServiceSessionParticipant FetchParticipant(ctx, serviceSid, sessionSid, sid)



Fetch a specific Participant.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to fetch. | 
**sessionSid** | **string**| The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resource to fetch. | 
**sid** | **string**| The Twilio-provided string that uniquely identifies the Participant resource to fetch. | 

### Return type

[**ProxyV1ServiceSessionParticipant**](proxy.v1.service.session.participant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPhoneNumber

> ProxyV1ServicePhoneNumber FetchPhoneNumber(ctx, serviceSid, sid)



Fetch a specific Phone Number.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the PhoneNumber resource to fetch. | 
**sid** | **string**| The Twilio-provided string that uniquely identifies the PhoneNumber resource to fetch. | 

### Return type

[**ProxyV1ServicePhoneNumber**](proxy.v1.service.phone_number.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> ProxyV1Service FetchService(ctx, sid)



Fetch a specific Service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The Twilio-provided string that uniquely identifies the Service resource to fetch. | 

### Return type

[**ProxyV1Service**](proxy.v1.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSession

> ProxyV1ServiceSession FetchSession(ctx, serviceSid, sid)



Fetch a specific Session.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to fetch. | 
**sid** | **string**| The Twilio-provided string that uniquely identifies the Session resource to fetch. | 

### Return type

[**ProxyV1ServiceSession**](proxy.v1.service.session.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchShortCode

> ProxyV1ServiceShortCode FetchShortCode(ctx, serviceSid, sid)



Fetch a specific Short Code.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) to fetch the resource from. | 
**sid** | **string**| The Twilio-provided string that uniquely identifies the ShortCode resource to fetch. | 

### Return type

[**ProxyV1ServiceShortCode**](proxy.v1.service.short_code.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInteraction

> ProxyV1ServiceSessionInteractionReadResponse ListInteraction(ctx, serviceSid, sessionSid, optional)



Retrieve a list of all Interactions for a Session. A maximum of 100 records will be returned per page.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) to read the resources from. | 
**sessionSid** | **string**| The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) to read the resources from. | 
 **optional** | ***ListInteractionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListInteractionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ProxyV1ServiceSessionInteractionReadResponse**](proxy_v1_service_session_interactionReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessageInteraction

> ProxyV1ServiceSessionParticipantMessageInteractionReadResponse ListMessageInteraction(ctx, serviceSid, sessionSid, participantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) to read the resources from. | 
**sessionSid** | **string**| The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) to read the resources from. | 
**participantSid** | **string**| The SID of the [Participant](https://www.twilio.com/docs/proxy/api/participant) to read the resources from. | 
 **optional** | ***ListMessageInteractionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMessageInteractionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ProxyV1ServiceSessionParticipantMessageInteractionReadResponse**](proxy_v1_service_session_participant_message_interactionReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListParticipant

> ProxyV1ServiceSessionParticipantReadResponse ListParticipant(ctx, serviceSid, sessionSid, optional)



Retrieve a list of all Participants in a Session.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resources to read. | 
**sessionSid** | **string**| The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resources to read. | 
 **optional** | ***ListParticipantOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListParticipantOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ProxyV1ServiceSessionParticipantReadResponse**](proxy_v1_service_session_participantReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPhoneNumber

> ProxyV1ServicePhoneNumberReadResponse ListPhoneNumber(ctx, serviceSid, optional)



Retrieve a list of all Phone Numbers in the Proxy Number Pool for a Service. A maximum of 100 records will be returned per page.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the PhoneNumber resources to read. | 
 **optional** | ***ListPhoneNumberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPhoneNumberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ProxyV1ServicePhoneNumberReadResponse**](proxy_v1_service_phone_numberReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListService

> ProxyV1ServiceReadResponse ListService(ctx, optional)



Retrieve a list of all Services for Twilio Proxy. A maximum of 100 records will be returned per page.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ProxyV1ServiceReadResponse**](proxy_v1_serviceReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSession

> ProxyV1ServiceSessionReadResponse ListSession(ctx, serviceSid, optional)



Retrieve a list of all Sessions for the Service. A maximum of 100 records will be returned per page.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to read. | 
 **optional** | ***ListSessionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSessionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ProxyV1ServiceSessionReadResponse**](proxy_v1_service_sessionReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListShortCode

> ProxyV1ServiceShortCodeReadResponse ListShortCode(ctx, serviceSid, optional)



Retrieve a list of all Short Codes in the Proxy Number Pool for the Service. A maximum of 100 records will be returned per page.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) to read the resources from. | 
 **optional** | ***ListShortCodeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListShortCodeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ProxyV1ServiceShortCodeReadResponse**](proxy_v1_service_short_codeReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePhoneNumber

> ProxyV1ServicePhoneNumber UpdatePhoneNumber(ctx, serviceSid, sid, optional)



Update a specific Proxy Number.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the PhoneNumber resource to update. | 
**sid** | **string**| The Twilio-provided string that uniquely identifies the PhoneNumber resource to update. | 
 **optional** | ***UpdatePhoneNumberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdatePhoneNumberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isReserved** | **optional.Bool**| Whether the phone number should be reserved and not be assigned to a participant using proxy pool logic. See [Reserved Phone Numbers](https://www.twilio.com/docs/proxy/reserved-phone-numbers) for more information. | 

### Return type

[**ProxyV1ServicePhoneNumber**](proxy.v1.service.phone_number.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> ProxyV1Service UpdateService(ctx, sid, optional)



Update a specific Service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The Twilio-provided string that uniquely identifies the Service resource to update. | 
 **optional** | ***UpdateServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **callbackUrl** | **optional.String**| The URL we should call when the interaction status changes. | 
 **chatInstanceSid** | **optional.String**| The SID of the Chat Service Instance managed by Proxy Service. The Chat Service enables Proxy to forward SMS and channel messages to this chat instance. This is a one-to-one relationship. | 
 **defaultTtl** | **optional.Int32**| The default &#x60;ttl&#x60; value to set for Sessions created in the Service. The TTL (time to live) is measured in seconds after the Session&#39;s last create or last Interaction. The default value of &#x60;0&#x60; indicates an unlimited Session length. You can override a Session&#39;s default TTL value by setting its &#x60;ttl&#x60; value. | 
 **geoMatchLevel** | **optional.String**| Where a proxy number must be located relative to the participant identifier. Can be: &#x60;country&#x60;, &#x60;area-code&#x60;, or &#x60;extended-area-code&#x60;. The default value is &#x60;country&#x60; and more specific areas than &#x60;country&#x60; are only available in North America. | 
 **interceptCallbackUrl** | **optional.String**| The URL we call on each interaction. If we receive a 403 status, we block the interaction; otherwise the interaction continues. | 
 **numberSelectionBehavior** | **optional.String**| The preference for Proxy Number selection in the Service instance. Can be: &#x60;prefer-sticky&#x60; or &#x60;avoid-sticky&#x60; and the default is &#x60;prefer-sticky&#x60;. &#x60;prefer-sticky&#x60; means that we will try and select the same Proxy Number for a given participant if they have previous [Sessions](https://www.twilio.com/docs/proxy/api/session), but we will not fail if that Proxy Number cannot be used.  &#x60;avoid-sticky&#x60; means that we will try to use different Proxy Numbers as long as that is possible within a given pool rather than try and use a previously assigned number. | 
 **outOfSessionCallbackUrl** | **optional.String**| The URL we should call when an inbound call or SMS action occurs on a closed or non-existent Session. If your server (or a Twilio [function](https://www.twilio.com/functions)) responds with valid [TwiML](https://www.twilio.com/docs/voice/twiml), we will process it. This means it is possible, for example, to play a message for a call, send an automated text message response, or redirect a call to another Phone Number. See [Out-of-Session Callback Response Guide](https://www.twilio.com/docs/proxy/out-session-callback-response-guide) for more information. | 
 **uniqueName** | **optional.String**| An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. **This value should not have PII.** | 

### Return type

[**ProxyV1Service**](proxy.v1.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSession

> ProxyV1ServiceSession UpdateSession(ctx, serviceSid, sid, optional)



Update a specific Session.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to update. | 
**sid** | **string**| The Twilio-provided string that uniquely identifies the Session resource to update. | 
 **optional** | ***UpdateSessionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSessionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dateExpiry** | **optional.Time**| The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session should expire. If this is value is present, it overrides the &#x60;ttl&#x60; value. | 
 **failOnParticipantConflict** | **optional.Bool**| [Experimental] For accounts with the ProxyAllowParticipantConflict account flag, setting to true enables per-request opt-in to allowing Proxy to return a 400 error (Twilio error code 80604) when a request to set a Session to in-progress would cause Participants with the same Identifier/ProxyIdentifier pair to be active in multiple Sessions. If not provided, requests will be allowed to succeed, and a Debugger notification (80801) will be emitted. Having multiple, active Participants with the same Identifier/ProxyIdentifier pair causes calls and messages from affected Participants to be routed incorrectly. Please note, the default behavior for accounts without the ProxyAllowParticipantConflict flag is to reject the request as described.  This will eventually be the default for all accounts. | 
 **status** | **optional.String**| The new status of the resource. Can be: &#x60;in-progress&#x60; to re-open a session or &#x60;closed&#x60; to close a session. | 
 **ttl** | **optional.Int32**| The time, in seconds, when the session will expire. The time is measured from the last Session create or the Session&#39;s last Interaction. | 

### Return type

[**ProxyV1ServiceSession**](proxy.v1.service.session.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShortCode

> ProxyV1ServiceShortCode UpdateShortCode(ctx, serviceSid, sid, optional)



Update a specific Short Code.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to update. | 
**sid** | **string**| The Twilio-provided string that uniquely identifies the ShortCode resource to update. | 
 **optional** | ***UpdateShortCodeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateShortCodeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isReserved** | **optional.Bool**| Whether the short code should be reserved and not be assigned to a participant using proxy pool logic. See [Reserved Phone Numbers](https://www.twilio.com/docs/proxy/reserved-phone-numbers) for more information. | 

### Return type

[**ProxyV1ServiceShortCode**](proxy.v1.service.short_code.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

