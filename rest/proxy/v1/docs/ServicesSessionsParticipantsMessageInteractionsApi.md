# ServicesSessionsParticipantsMessageInteractionsApi

All URIs are relative to *https://proxy.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessageInteraction**](ServicesSessionsParticipantsMessageInteractionsApi.md#CreateMessageInteraction) | **Post** /v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants/{ParticipantSid}/MessageInteractions | 
[**FetchMessageInteraction**](ServicesSessionsParticipantsMessageInteractionsApi.md#FetchMessageInteraction) | **Get** /v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants/{ParticipantSid}/MessageInteractions/{Sid} | 
[**ListMessageInteraction**](ServicesSessionsParticipantsMessageInteractionsApi.md#ListMessageInteraction) | **Get** /v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants/{ParticipantSid}/MessageInteractions | 



## CreateMessageInteraction

> ProxyV1MessageInteraction CreateMessageInteraction(ctx, ServiceSidSessionSidParticipantSidoptional)



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

[**ProxyV1MessageInteraction**](ProxyV1MessageInteraction.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMessageInteraction

> ProxyV1MessageInteraction FetchMessageInteraction(ctx, ServiceSidSessionSidParticipantSidSid)





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

[**ProxyV1MessageInteraction**](ProxyV1MessageInteraction.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessageInteraction

> []ProxyV1MessageInteraction ListMessageInteraction(ctx, ServiceSidSessionSidParticipantSidoptional)





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
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ProxyV1MessageInteraction**](ProxyV1MessageInteraction.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

