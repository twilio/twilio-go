# ConferencesParticipantsApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchConferenceParticipant**](ConferencesParticipantsApi.md#FetchConferenceParticipant) | **Get** /v1/Conferences/{ConferenceSid}/Participants/{ParticipantSid} | 
[**ListConferenceParticipant**](ConferencesParticipantsApi.md#ListConferenceParticipant) | **Get** /v1/Conferences/{ConferenceSid}/Participants | 



## FetchConferenceParticipant

> InsightsV1ConferenceParticipant FetchConferenceParticipant(ctx, ConferenceSidParticipantSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConferenceSid** | **string** | 
**ParticipantSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchConferenceParticipantParams struct


Name | Type | Description
------------- | ------------- | -------------
**Events** | **string** | 
**Metrics** | **string** | 

### Return type

[**InsightsV1ConferenceParticipant**](InsightsV1ConferenceParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConferenceParticipant

> ListConferenceParticipantResponse ListConferenceParticipant(ctx, ConferenceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConferenceSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListConferenceParticipantParams struct


Name | Type | Description
------------- | ------------- | -------------
**ParticipantSid** | **string** | 
**Label** | **string** | 
**Events** | **string** | 
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**ListConferenceParticipantResponse**](ListConferenceParticipantResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

