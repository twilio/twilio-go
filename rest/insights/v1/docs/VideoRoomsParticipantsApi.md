# VideoRoomsParticipantsApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchVideoParticipantSummary**](VideoRoomsParticipantsApi.md#FetchVideoParticipantSummary) | **Get** /v1/Video/Rooms/{RoomSid}/Participants/{ParticipantSid} | 
[**ListVideoParticipantSummary**](VideoRoomsParticipantsApi.md#ListVideoParticipantSummary) | **Get** /v1/Video/Rooms/{RoomSid}/Participants | 



## FetchVideoParticipantSummary

> InsightsV1VideoParticipantSummary FetchVideoParticipantSummary(ctx, RoomSidParticipantSid)



Get Video Log Analyzer data for a Room Participant.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room resource.
**ParticipantSid** | **string** | The SID of the Participant resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchVideoParticipantSummaryParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**InsightsV1VideoParticipantSummary**](InsightsV1VideoParticipantSummary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVideoParticipantSummary

> []InsightsV1VideoParticipantSummary ListVideoParticipantSummary(ctx, RoomSidoptional)



Get a list of room participants.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room resource.

### Other Parameters

Other parameters are passed through a pointer to a ListVideoParticipantSummaryParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]InsightsV1VideoParticipantSummary**](InsightsV1VideoParticipantSummary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

