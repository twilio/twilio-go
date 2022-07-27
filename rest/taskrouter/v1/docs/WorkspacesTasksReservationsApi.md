# WorkspacesTasksReservationsApi

All URIs are relative to *https://taskrouter.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchTaskReservation**](WorkspacesTasksReservationsApi.md#FetchTaskReservation) | **Get** /v1/Workspaces/{WorkspaceSid}/Tasks/{TaskSid}/Reservations/{Sid} | 
[**ListTaskReservation**](WorkspacesTasksReservationsApi.md#ListTaskReservation) | **Get** /v1/Workspaces/{WorkspaceSid}/Tasks/{TaskSid}/Reservations | 
[**UpdateTaskReservation**](WorkspacesTasksReservationsApi.md#UpdateTaskReservation) | **Post** /v1/Workspaces/{WorkspaceSid}/Tasks/{TaskSid}/Reservations/{Sid} | 



## FetchTaskReservation

> TaskrouterV1TaskReservation FetchTaskReservation(ctx, WorkspaceSidTaskSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the TaskReservation resource to fetch.
**TaskSid** | **string** | The SID of the reserved Task resource with the TaskReservation resource to fetch.
**Sid** | **string** | The SID of the TaskReservation resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchTaskReservationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TaskrouterV1TaskReservation**](TaskrouterV1TaskReservation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTaskReservation

> []TaskrouterV1TaskReservation ListTaskReservation(ctx, WorkspaceSidTaskSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the TaskReservation resources to read.
**TaskSid** | **string** | The SID of the reserved Task resource with the TaskReservation resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListTaskReservationParams struct


Name | Type | Description
------------- | ------------- | -------------
**ReservationStatus** | **string** | Returns the list of reservations for a task with a specified ReservationStatus.  Can be: &#x60;pending&#x60;, &#x60;accepted&#x60;, &#x60;rejected&#x60;, or &#x60;timeout&#x60;.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TaskrouterV1TaskReservation**](TaskrouterV1TaskReservation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaskReservation

> TaskrouterV1TaskReservation UpdateTaskReservation(ctx, WorkspaceSidTaskSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the TaskReservation resources to update.
**TaskSid** | **string** | The SID of the reserved Task resource with the TaskReservation resources to update.
**Sid** | **string** | The SID of the TaskReservation resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateTaskReservationParams struct


Name | Type | Description
------------- | ------------- | -------------
**IfMatch** | **string** | The If-Match HTTP request header
**ReservationStatus** | **string** | 
**WorkerActivitySid** | **string** | The new worker activity SID if rejecting a reservation.
**Instruction** | **string** | The assignment instruction for reservation.
**DequeuePostWorkActivitySid** | **string** | The SID of the Activity resource to start after executing a Dequeue instruction.
**DequeueFrom** | **string** | The Caller ID of the call to the worker when executing a Dequeue instruction.
**DequeueRecord** | **string** | Whether to record both legs of a call when executing a Dequeue instruction or which leg to record.
**DequeueTimeout** | **int** | Timeout for call when executing a Dequeue instruction.
**DequeueTo** | **string** | The Contact URI of the worker when executing a Dequeue instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination.
**DequeueStatusCallbackUrl** | **string** | The Callback URL for completed call event when executing a Dequeue instruction.
**CallFrom** | **string** | The Caller ID of the outbound call when executing a Call instruction.
**CallRecord** | **string** | Whether to record both legs of a call when executing a Call instruction or which leg to record.
**CallTimeout** | **int** | Timeout for call when executing a Call instruction.
**CallTo** | **string** | The Contact URI of the worker when executing a Call instruction.  Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination.
**CallUrl** | **string** | TwiML URI executed on answering the worker&#39;s leg as a result of the Call instruction.
**CallStatusCallbackUrl** | **string** | The URL to call  for the completed call event when executing a Call instruction.
**CallAccept** | **bool** | Whether to accept a reservation when executing a Call instruction.
**RedirectCallSid** | **string** | The Call SID of the call parked in the queue when executing a Redirect instruction.
**RedirectAccept** | **bool** | Whether the reservation should be accepted when executing a Redirect instruction.
**RedirectUrl** | **string** | TwiML URI to redirect the call to when executing the Redirect instruction.
**To** | **string** | The Contact URI of the worker when executing a Conference instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination.
**From** | **string** | The Caller ID of the call to the worker when executing a Conference instruction.
**StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application.
**StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;POST&#x60; or &#x60;GET&#x60; and the default is &#x60;POST&#x60;.
**StatusCallbackEvent** | [**[]TaskReservationEnumCallStatus**](TaskReservationEnumCallStatus.md) | The call progress events that we will send to &#x60;status_callback&#x60;. Can be: &#x60;initiated&#x60;, &#x60;ringing&#x60;, &#x60;answered&#x60;, or &#x60;completed&#x60;.
**Timeout** | **int** | Timeout for call when executing a Conference instruction.
**Record** | **bool** | Whether to record the participant and their conferences, including the time between conferences. The default is &#x60;false&#x60;.
**Muted** | **bool** | Whether the agent is muted in the conference. The default is &#x60;false&#x60;.
**Beep** | **string** | Whether to play a notification beep when the participant joins or when to play a beep. Can be: &#x60;true&#x60;, &#x60;false&#x60;, &#x60;onEnter&#x60;, or &#x60;onExit&#x60;. The default value is &#x60;true&#x60;.
**StartConferenceOnEnter** | **bool** | Whether to start the conference when the participant joins, if it has not already started. The default is &#x60;true&#x60;. If &#x60;false&#x60; and the conference has not started, the participant is muted and hears background music until another participant starts the conference.
**EndConferenceOnExit** | **bool** | Whether to end the conference when the agent leaves.
**WaitUrl** | **string** | The URL we should call using the &#x60;wait_method&#x60; for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic).
**WaitMethod** | **string** | The HTTP method we should use to call &#x60;wait_url&#x60;. Can be &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. When using a static audio file, this should be &#x60;GET&#x60; so that we can cache the file.
**EarlyMedia** | **bool** | Whether to allow an agent to hear the state of the outbound call, including ringing or disconnect messages. The default is &#x60;true&#x60;.
**MaxParticipants** | **int** | The maximum number of participants in the conference. Can be a positive integer from &#x60;2&#x60; to &#x60;250&#x60;. The default value is &#x60;250&#x60;.
**ConferenceStatusCallback** | **string** | The URL we should call using the &#x60;conference_status_callback_method&#x60; when the conference events in &#x60;conference_status_callback_event&#x60; occur. Only the value set by the first participant to join the conference is used. Subsequent &#x60;conference_status_callback&#x60; values are ignored.
**ConferenceStatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;conference_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**ConferenceStatusCallbackEvent** | [**[]TaskReservationEnumConferenceEvent**](TaskReservationEnumConferenceEvent.md) | The conference status events that we will send to &#x60;conference_status_callback&#x60;. Can be: &#x60;start&#x60;, &#x60;end&#x60;, &#x60;join&#x60;, &#x60;leave&#x60;, &#x60;mute&#x60;, &#x60;hold&#x60;, &#x60;speaker&#x60;.
**ConferenceRecord** | **string** | Whether to record the conference the participant is joining or when to record the conference. Can be: &#x60;true&#x60;, &#x60;false&#x60;, &#x60;record-from-start&#x60;, and &#x60;do-not-record&#x60;. The default value is &#x60;false&#x60;.
**ConferenceTrim** | **string** | How to trim the leading and trailing silence from your recorded conference audio files. Can be: &#x60;trim-silence&#x60; or &#x60;do-not-trim&#x60; and defaults to &#x60;trim-silence&#x60;.
**RecordingChannels** | **string** | The recording channels for the final recording. Can be: &#x60;mono&#x60; or &#x60;dual&#x60; and the default is &#x60;mono&#x60;.
**RecordingStatusCallback** | **string** | The URL that we should call using the &#x60;recording_status_callback_method&#x60; when the recording status changes.
**RecordingStatusCallbackMethod** | **string** | The HTTP method we should use when we call &#x60;recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**ConferenceRecordingStatusCallback** | **string** | The URL we should call using the &#x60;conference_recording_status_callback_method&#x60; when the conference recording is available.
**ConferenceRecordingStatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;conference_recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**Region** | **string** | The [region](https://support.twilio.com/hc/en-us/articles/223132167-How-global-low-latency-routing-and-region-selection-work-for-conferences-and-Client-calls) where we should mix the recorded audio. Can be:&#x60;us1&#x60;, &#x60;ie1&#x60;, &#x60;de1&#x60;, &#x60;sg1&#x60;, &#x60;br1&#x60;, &#x60;au1&#x60;, or &#x60;jp1&#x60;.
**SipAuthUsername** | **string** | The SIP username used for authentication.
**SipAuthPassword** | **string** | The SIP password for authentication.
**DequeueStatusCallbackEvent** | **[]string** | The Call progress events sent via webhooks as a result of a Dequeue instruction.
**PostWorkActivitySid** | **string** | The new worker activity SID after executing a Conference instruction.
**SupervisorMode** | **string** | 
**Supervisor** | **string** | The Supervisor SID/URI when executing the Supervise instruction.
**EndConferenceOnCustomerExit** | **bool** | Whether to end the conference when the customer leaves.
**BeepOnCustomerEntrance** | **bool** | Whether to play a notification beep when the customer joins.

### Return type

[**TaskrouterV1TaskReservation**](TaskrouterV1TaskReservation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

