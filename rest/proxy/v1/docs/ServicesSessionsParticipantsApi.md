# ServicesSessionsParticipantsApi

All URIs are relative to *https://proxy.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateParticipant**](ServicesSessionsParticipantsApi.md#CreateParticipant) | **Post** /v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants | 
[**DeleteParticipant**](ServicesSessionsParticipantsApi.md#DeleteParticipant) | **Delete** /v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants/{Sid} | 
[**FetchParticipant**](ServicesSessionsParticipantsApi.md#FetchParticipant) | **Get** /v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants/{Sid} | 
[**ListParticipant**](ServicesSessionsParticipantsApi.md#ListParticipant) | **Get** /v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants | 



## CreateParticipant

> ProxyV1Participant CreateParticipant(ctx, ServiceSidSessionSidoptional)



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
**Identifier** | **string** | The phone number of the Participant.
**FriendlyName** | **string** | The string that you assigned to describe the participant. This value must be 255 characters or fewer. **This value should not have PII.**
**ProxyIdentifier** | **string** | The proxy phone number to use for the Participant. If not specified, Proxy will select a number from the pool.
**ProxyIdentifierSid** | **string** | The SID of the Proxy Identifier to assign to the Participant.

### Return type

[**ProxyV1Participant**](ProxyV1Participant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

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


## FetchParticipant

> ProxyV1Participant FetchParticipant(ctx, ServiceSidSessionSidSid)



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

[**ProxyV1Participant**](ProxyV1Participant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListParticipant

> []ProxyV1Participant ListParticipant(ctx, ServiceSidSessionSidoptional)



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
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ProxyV1Participant**](ProxyV1Participant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

