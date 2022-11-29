# WorkspacesWorkersReservationsApi

All URIs are relative to *https://taskrouter.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchWorkerReservation**](WorkspacesWorkersReservationsApi.md#FetchWorkerReservation) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Reservations/{Sid} | 
[**ListWorkerReservation**](WorkspacesWorkersReservationsApi.md#ListWorkerReservation) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Reservations | 
[**UpdateWorkerReservation**](WorkspacesWorkersReservationsApi.md#UpdateWorkerReservation) | **Post** /v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Reservations/{Sid} | 



## FetchWorkerReservation

> TaskrouterV1WorkerReservation FetchWorkerReservation(ctx, WorkspaceSidWorkerSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the WorkerReservation resource to fetch.
**WorkerSid** | **string** | The SID of the reserved Worker resource with the WorkerReservation resource to fetch.
**Sid** | **string** | The SID of the WorkerReservation resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWorkerReservationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TaskrouterV1WorkerReservation**](TaskrouterV1WorkerReservation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkerReservation

> []TaskrouterV1WorkerReservation ListWorkerReservation(ctx, WorkspaceSidWorkerSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the WorkerReservation resources to read.
**WorkerSid** | **string** | The SID of the reserved Worker resource with the WorkerReservation resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListWorkerReservationParams struct


Name | Type | Description
------------- | ------------- | -------------
**ReservationStatus** | **string** | Returns the list of reservations for a worker with a specified ReservationStatus. Can be: `pending`, `accepted`, `rejected`, `timeout`, `canceled`, or `rescinded`.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TaskrouterV1WorkerReservation**](TaskrouterV1WorkerReservation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkerReservation

> TaskrouterV1WorkerReservation UpdateWorkerReservation(ctx, WorkspaceSidWorkerSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the WorkerReservation resources to update.
**WorkerSid** | **string** | The SID of the reserved Worker resource with the WorkerReservation resources to update.
**Sid** | **string** | The SID of the WorkerReservation resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateWorkerReservationParams struct


Name | Type | Description
------------- | ------------- | -------------
**IfMatch** | **string** | The If-Match HTTP request header
**ReservationStatus** | **string** | 
**WorkerActivitySid** | **string** | The new worker activity SID if rejecting a reservation.
**Instruction** | **string** | The assignment instruction for the reservation.
**DequeuePostWorkActivitySid** | **string** | The SID of the Activity resource to start after executing a Dequeue instruction.
**DequeueFrom** | **string** | The caller ID of the call to the worker when executing a Dequeue instruction.
**DequeueRecord** | **string** | Whether to record both legs of a call when executing a Dequeue instruction or which leg to record.
**DequeueTimeout** | **int** | The timeout for call when executing a Dequeue instruction.
**DequeueTo** | **string** | The contact URI of the worker when executing a Dequeue instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination.
**DequeueStatusCallbackUrl** | **string** | The callback URL for completed call event when executing a Dequeue instruction.
**CallFrom** | **string** | The Caller ID of the outbound call when executing a Call instruction.
**CallRecord** | **string** | Whether to record both legs of a call when executing a Call instruction.
**CallTimeout** | **int** | The timeout for a call when executing a Call instruction.
**CallTo** | **string** | The contact URI of the worker when executing a Call instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination.
**CallUrl** | **string** | TwiML URI executed on answering the worker's leg as a result of the Call instruction.
**CallStatusCallbackUrl** | **string** | The URL to call for the completed call event when executing a Call instruction.
**CallAccept** | **bool** | Whether to accept a reservation when executing a Call instruction.
**RedirectCallSid** | **string** | The Call SID of the call parked in the queue when executing a Redirect instruction.
**RedirectAccept** | **bool** | Whether the reservation should be accepted when executing a Redirect instruction.
**RedirectUrl** | **string** | TwiML URI to redirect the call to when executing the Redirect instruction.
**To** | **string** | The Contact URI of the worker when executing a Conference instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination.
**From** | **string** | The caller ID of the call to the worker when executing a Conference instruction.
**StatusCallback** | **string** | The URL we should call using the `status_callback_method` to send status information to your application.
**StatusCallbackMethod** | **string** | The HTTP method we should use to call `status_callback`. Can be: `POST` or `GET` and the default is `POST`.
**StatusCallbackEvent** | [**[]WorkerReservationEnumCallStatus**](WorkerReservationEnumCallStatus.md) | The call progress events that we will send to `status_callback`. Can be: `initiated`, `ringing`, `answered`, or `completed`.
**Timeout** | **int** | The timeout for a call when executing a Conference instruction.
**Record** | **bool** | Whether to record the participant and their conferences, including the time between conferences. Can be `true` or `false` and the default is `false`.
**Muted** | **bool** | Whether the agent is muted in the conference. Defaults to `false`.
**Beep** | **string** | Whether to play a notification beep when the participant joins or when to play a beep. Can be: `true`, `false`, `onEnter`, or `onExit`. The default value is `true`.
**StartConferenceOnEnter** | **bool** | Whether to start the conference when the participant joins, if it has not already started. Can be: `true` or `false` and the default is `true`. If `false` and the conference has not started, the participant is muted and hears background music until another participant starts the conference.
**EndConferenceOnExit** | **bool** | Whether to end the conference when the agent leaves.
**WaitUrl** | **string** | The URL we should call using the `wait_method` for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic).
**WaitMethod** | **string** | The HTTP method we should use to call `wait_url`. Can be `GET` or `POST` and the default is `POST`. When using a static audio file, this should be `GET` so that we can cache the file.
**EarlyMedia** | **bool** | Whether to allow an agent to hear the state of the outbound call, including ringing or disconnect messages. The default is `true`.
**MaxParticipants** | **int** | The maximum number of participants allowed in the conference. Can be a positive integer from `2` to `250`. The default value is `250`.
**ConferenceStatusCallback** | **string** | The URL we should call using the `conference_status_callback_method` when the conference events in `conference_status_callback_event` occur. Only the value set by the first participant to join the conference is used. Subsequent `conference_status_callback` values are ignored.
**ConferenceStatusCallbackMethod** | **string** | The HTTP method we should use to call `conference_status_callback`. Can be: `GET` or `POST` and defaults to `POST`.
**ConferenceStatusCallbackEvent** | [**[]WorkerReservationEnumConferenceEvent**](WorkerReservationEnumConferenceEvent.md) | The conference status events that we will send to `conference_status_callback`. Can be: `start`, `end`, `join`, `leave`, `mute`, `hold`, `speaker`.
**ConferenceRecord** | **string** | Whether to record the conference the participant is joining or when to record the conference. Can be: `true`, `false`, `record-from-start`, and `do-not-record`. The default value is `false`.
**ConferenceTrim** | **string** | Whether to trim leading and trailing silence from your recorded conference audio files. Can be: `trim-silence` or `do-not-trim` and defaults to `trim-silence`.
**RecordingChannels** | **string** | The recording channels for the final recording. Can be: `mono` or `dual` and the default is `mono`.
**RecordingStatusCallback** | **string** | The URL that we should call using the `recording_status_callback_method` when the recording status changes.
**RecordingStatusCallbackMethod** | **string** | The HTTP method we should use when we call `recording_status_callback`. Can be: `GET` or `POST` and defaults to `POST`.
**ConferenceRecordingStatusCallback** | **string** | The URL we should call using the `conference_recording_status_callback_method` when the conference recording is available.
**ConferenceRecordingStatusCallbackMethod** | **string** | The HTTP method we should use to call `conference_recording_status_callback`. Can be: `GET` or `POST` and defaults to `POST`.
**Region** | **string** | The [region](https://support.twilio.com/hc/en-us/articles/223132167-How-global-low-latency-routing-and-region-selection-work-for-conferences-and-Client-calls) where we should mix the recorded audio. Can be:`us1`, `ie1`, `de1`, `sg1`, `br1`, `au1`, or `jp1`.
**SipAuthUsername** | **string** | The SIP username used for authentication.
**SipAuthPassword** | **string** | The SIP password for authentication.
**DequeueStatusCallbackEvent** | **[]string** | The call progress events sent via webhooks as a result of a Dequeue instruction.
**PostWorkActivitySid** | **string** | The new worker activity SID after executing a Conference instruction.
**EndConferenceOnCustomerExit** | **bool** | Whether to end the conference when the customer leaves.
**BeepOnCustomerEntrance** | **bool** | Whether to play a notification beep when the customer joins.

### Return type

[**TaskrouterV1WorkerReservation**](TaskrouterV1WorkerReservation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

