# ServicesSessionsApi

All URIs are relative to *https://proxy.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSession**](ServicesSessionsApi.md#CreateSession) | **Post** /v1/Services/{ServiceSid}/Sessions | 
[**DeleteSession**](ServicesSessionsApi.md#DeleteSession) | **Delete** /v1/Services/{ServiceSid}/Sessions/{Sid} | 
[**FetchSession**](ServicesSessionsApi.md#FetchSession) | **Get** /v1/Services/{ServiceSid}/Sessions/{Sid} | 
[**ListSession**](ServicesSessionsApi.md#ListSession) | **Get** /v1/Services/{ServiceSid}/Sessions | 
[**UpdateSession**](ServicesSessionsApi.md#UpdateSession) | **Post** /v1/Services/{ServiceSid}/Sessions/{Sid} | 



## CreateSession

> ProxyV1Session CreateSession(ctx, ServiceSidoptional)



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

[**ProxyV1Session**](ProxyV1Session.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

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


## FetchSession

> ProxyV1Session FetchSession(ctx, ServiceSidSid)



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

[**ProxyV1Session**](ProxyV1Session.md)

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
**Limit** | **int** | Max number of records to return.

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


## UpdateSession

> ProxyV1Session UpdateSession(ctx, ServiceSidSidoptional)



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

[**ProxyV1Session**](ProxyV1Session.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

