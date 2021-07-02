/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.18.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"

	"github.com/twilio/twilio-go/client"
)

func (c *ApiService) FetchTaskReservation(WorkspaceSid string, TaskSid string, Sid string) (*TaskrouterV1WorkspaceTaskTaskReservation, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Tasks/{TaskSid}/Reservations/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"TaskSid"+"}", TaskSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1WorkspaceTaskTaskReservation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListTaskReservation'
type ListTaskReservationParams struct {
	// Returns the list of reservations for a task with a specified ReservationStatus.  Can be: `pending`, `accepted`, `rejected`, or `timeout`.
	ReservationStatus *string `json:"ReservationStatus,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListTaskReservationParams) SetReservationStatus(ReservationStatus string) *ListTaskReservationParams {
	params.ReservationStatus = &ReservationStatus
	return params
}
func (params *ListTaskReservationParams) SetPageSize(PageSize int) *ListTaskReservationParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListTaskReservation(WorkspaceSid string, TaskSid string, params *ListTaskReservationParams) (*ListTaskReservationResponse, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Tasks/{TaskSid}/Reservations"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"TaskSid"+"}", TaskSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ReservationStatus != nil {
		data.Set("ReservationStatus", *params.ReservationStatus)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTaskReservationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of TaskReservation records from the API. Request is executed immediately.
func (c *ApiService) TaskReservationPage(WorkspaceSid string, TaskSid string, params *ListTaskReservationParams, pageToken string, pageNumber string) (*client.Page, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Tasks/{TaskSid}/Reservations"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"TaskSid"+"}", TaskSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ReservationStatus != nil {
		data.Set("ReservationStatus", *params.ReservationStatus)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	data.Set("PageToken", pageToken)
	data.Set("PageNumber", pageNumber)

	response, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	return client.NewPage(c.baseURL, response), nil
}

//Streams TaskReservation records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) TaskReservationStream(WorkspaceSid string, TaskSid string, params *ListTaskReservationParams, limit int) (chan map[string]interface{}, error) {
	if params.PageSize == nil {
		params.SetPageSize(0)
	}
	params.SetPageSize(c.requestHandler.ReadLimits(*params.PageSize, limit))
	page, err := c.TaskReservationPage(WorkspaceSid, TaskSid, params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.Stream(page, limit, 0), nil
}

//Lists TaskReservation records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) TaskReservationList(WorkspaceSid string, TaskSid string, params *ListTaskReservationParams, limit int) ([]interface{}, error) {
	if params.PageSize == nil {
		params.SetPageSize(0)
	}
	params.SetPageSize(c.requestHandler.ReadLimits(*params.PageSize, limit))
	page, err := c.TaskReservationPage(WorkspaceSid, TaskSid, params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.List(page, limit, 0), nil
}

// Optional parameters for the method 'UpdateTaskReservation'
type UpdateTaskReservationParams struct {
	// Whether to play a notification beep when the participant joins or when to play a beep. Can be: `true`, `false`, `onEnter`, or `onExit`. The default value is `true`.
	Beep *string `json:"Beep,omitempty"`
	// Whether to play a notification beep when the customer joins.
	BeepOnCustomerEntrance *bool `json:"BeepOnCustomerEntrance,omitempty"`
	// Whether to accept a reservation when executing a Call instruction.
	CallAccept *bool `json:"CallAccept,omitempty"`
	// The Caller ID of the outbound call when executing a Call instruction.
	CallFrom *string `json:"CallFrom,omitempty"`
	// Whether to record both legs of a call when executing a Call instruction or which leg to record.
	CallRecord *string `json:"CallRecord,omitempty"`
	// The URL to call  for the completed call event when executing a Call instruction.
	CallStatusCallbackUrl *string `json:"CallStatusCallbackUrl,omitempty"`
	// Timeout for call when executing a Call instruction.
	CallTimeout *int `json:"CallTimeout,omitempty"`
	// The Contact URI of the worker when executing a Call instruction.  Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination.
	CallTo *string `json:"CallTo,omitempty"`
	// TwiML URI executed on answering the worker's leg as a result of the Call instruction.
	CallUrl *string `json:"CallUrl,omitempty"`
	// Whether to record the conference the participant is joining or when to record the conference. Can be: `true`, `false`, `record-from-start`, and `do-not-record`. The default value is `false`.
	ConferenceRecord *string `json:"ConferenceRecord,omitempty"`
	// The URL we should call using the `conference_recording_status_callback_method` when the conference recording is available.
	ConferenceRecordingStatusCallback *string `json:"ConferenceRecordingStatusCallback,omitempty"`
	// The HTTP method we should use to call `conference_recording_status_callback`. Can be: `GET` or `POST` and defaults to `POST`.
	ConferenceRecordingStatusCallbackMethod *string `json:"ConferenceRecordingStatusCallbackMethod,omitempty"`
	// The URL we should call using the `conference_status_callback_method` when the conference events in `conference_status_callback_event` occur. Only the value set by the first participant to join the conference is used. Subsequent `conference_status_callback` values are ignored.
	ConferenceStatusCallback *string `json:"ConferenceStatusCallback,omitempty"`
	// The conference status events that we will send to `conference_status_callback`. Can be: `start`, `end`, `join`, `leave`, `mute`, `hold`, `speaker`.
	ConferenceStatusCallbackEvent *[]string `json:"ConferenceStatusCallbackEvent,omitempty"`
	// The HTTP method we should use to call `conference_status_callback`. Can be: `GET` or `POST` and defaults to `POST`.
	ConferenceStatusCallbackMethod *string `json:"ConferenceStatusCallbackMethod,omitempty"`
	// How to trim the leading and trailing silence from your recorded conference audio files. Can be: `trim-silence` or `do-not-trim` and defaults to `trim-silence`.
	ConferenceTrim *string `json:"ConferenceTrim,omitempty"`
	// The Caller ID of the call to the worker when executing a Dequeue instruction.
	DequeueFrom *string `json:"DequeueFrom,omitempty"`
	// The SID of the Activity resource to start after executing a Dequeue instruction.
	DequeuePostWorkActivitySid *string `json:"DequeuePostWorkActivitySid,omitempty"`
	// Whether to record both legs of a call when executing a Dequeue instruction or which leg to record.
	DequeueRecord *string `json:"DequeueRecord,omitempty"`
	// The Call progress events sent via webhooks as a result of a Dequeue instruction.
	DequeueStatusCallbackEvent *[]string `json:"DequeueStatusCallbackEvent,omitempty"`
	// The Callback URL for completed call event when executing a Dequeue instruction.
	DequeueStatusCallbackUrl *string `json:"DequeueStatusCallbackUrl,omitempty"`
	// Timeout for call when executing a Dequeue instruction.
	DequeueTimeout *int `json:"DequeueTimeout,omitempty"`
	// The Contact URI of the worker when executing a Dequeue instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination.
	DequeueTo *string `json:"DequeueTo,omitempty"`
	// Whether to allow an agent to hear the state of the outbound call, including ringing or disconnect messages. The default is `true`.
	EarlyMedia *bool `json:"EarlyMedia,omitempty"`
	// Whether to end the conference when the customer leaves.
	EndConferenceOnCustomerExit *bool `json:"EndConferenceOnCustomerExit,omitempty"`
	// Whether to end the conference when the agent leaves.
	EndConferenceOnExit *bool `json:"EndConferenceOnExit,omitempty"`
	// The Caller ID of the call to the worker when executing a Conference instruction.
	From *string `json:"From,omitempty"`
	// The assignment instruction for reservation.
	Instruction *string `json:"Instruction,omitempty"`
	// The maximum number of participants in the conference. Can be a positive integer from `2` to `250`. The default value is `250`.
	MaxParticipants *int `json:"MaxParticipants,omitempty"`
	// Whether the agent is muted in the conference. The default is `false`.
	Muted *bool `json:"Muted,omitempty"`
	// The new worker activity SID after executing a Conference instruction.
	PostWorkActivitySid *string `json:"PostWorkActivitySid,omitempty"`
	// Whether to record the participant and their conferences, including the time between conferences. The default is `false`.
	Record *bool `json:"Record,omitempty"`
	// The recording channels for the final recording. Can be: `mono` or `dual` and the default is `mono`.
	RecordingChannels *string `json:"RecordingChannels,omitempty"`
	// The URL that we should call using the `recording_status_callback_method` when the recording status changes.
	RecordingStatusCallback *string `json:"RecordingStatusCallback,omitempty"`
	// The HTTP method we should use when we call `recording_status_callback`. Can be: `GET` or `POST` and defaults to `POST`.
	RecordingStatusCallbackMethod *string `json:"RecordingStatusCallbackMethod,omitempty"`
	// Whether the reservation should be accepted when executing a Redirect instruction.
	RedirectAccept *bool `json:"RedirectAccept,omitempty"`
	// The Call SID of the call parked in the queue when executing a Redirect instruction.
	RedirectCallSid *string `json:"RedirectCallSid,omitempty"`
	// TwiML URI to redirect the call to when executing the Redirect instruction.
	RedirectUrl *string `json:"RedirectUrl,omitempty"`
	// The [region](https://support.twilio.com/hc/en-us/articles/223132167-How-global-low-latency-routing-and-region-selection-work-for-conferences-and-Client-calls) where we should mix the recorded audio. Can be:`us1`, `ie1`, `de1`, `sg1`, `br1`, `au1`, or `jp1`.
	Region *string `json:"Region,omitempty"`
	// The new status of the reservation. Can be: `pending`, `accepted`, `rejected`, or `timeout`.
	ReservationStatus *string `json:"ReservationStatus,omitempty"`
	// The SIP password for authentication.
	SipAuthPassword *string `json:"SipAuthPassword,omitempty"`
	// The SIP username used for authentication.
	SipAuthUsername *string `json:"SipAuthUsername,omitempty"`
	// Whether to start the conference when the participant joins, if it has not already started. The default is `true`. If `false` and the conference has not started, the participant is muted and hears background music until another participant starts the conference.
	StartConferenceOnEnter *bool `json:"StartConferenceOnEnter,omitempty"`
	// The URL we should call using the `status_callback_method` to send status information to your application.
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// The call progress events that we will send to `status_callback`. Can be: `initiated`, `ringing`, `answered`, or `completed`.
	StatusCallbackEvent *[]string `json:"StatusCallbackEvent,omitempty"`
	// The HTTP method we should use to call `status_callback`. Can be: `POST` or `GET` and the default is `POST`.
	StatusCallbackMethod *string `json:"StatusCallbackMethod,omitempty"`
	// The Supervisor SID/URI when executing the Supervise instruction.
	Supervisor *string `json:"Supervisor,omitempty"`
	// The Supervisor mode when executing the Supervise instruction.
	SupervisorMode *string `json:"SupervisorMode,omitempty"`
	// Timeout for call when executing a Conference instruction.
	Timeout *int `json:"Timeout,omitempty"`
	// The Contact URI of the worker when executing a Conference instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination.
	To *string `json:"To,omitempty"`
	// The HTTP method we should use to call `wait_url`. Can be `GET` or `POST` and the default is `POST`. When using a static audio file, this should be `GET` so that we can cache the file.
	WaitMethod *string `json:"WaitMethod,omitempty"`
	// The URL we should call using the `wait_method` for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic).
	WaitUrl *string `json:"WaitUrl,omitempty"`
	// The new worker activity SID if rejecting a reservation.
	WorkerActivitySid *string `json:"WorkerActivitySid,omitempty"`
}

func (params *UpdateTaskReservationParams) SetBeep(Beep string) *UpdateTaskReservationParams {
	params.Beep = &Beep
	return params
}
func (params *UpdateTaskReservationParams) SetBeepOnCustomerEntrance(BeepOnCustomerEntrance bool) *UpdateTaskReservationParams {
	params.BeepOnCustomerEntrance = &BeepOnCustomerEntrance
	return params
}
func (params *UpdateTaskReservationParams) SetCallAccept(CallAccept bool) *UpdateTaskReservationParams {
	params.CallAccept = &CallAccept
	return params
}
func (params *UpdateTaskReservationParams) SetCallFrom(CallFrom string) *UpdateTaskReservationParams {
	params.CallFrom = &CallFrom
	return params
}
func (params *UpdateTaskReservationParams) SetCallRecord(CallRecord string) *UpdateTaskReservationParams {
	params.CallRecord = &CallRecord
	return params
}
func (params *UpdateTaskReservationParams) SetCallStatusCallbackUrl(CallStatusCallbackUrl string) *UpdateTaskReservationParams {
	params.CallStatusCallbackUrl = &CallStatusCallbackUrl
	return params
}
func (params *UpdateTaskReservationParams) SetCallTimeout(CallTimeout int) *UpdateTaskReservationParams {
	params.CallTimeout = &CallTimeout
	return params
}
func (params *UpdateTaskReservationParams) SetCallTo(CallTo string) *UpdateTaskReservationParams {
	params.CallTo = &CallTo
	return params
}
func (params *UpdateTaskReservationParams) SetCallUrl(CallUrl string) *UpdateTaskReservationParams {
	params.CallUrl = &CallUrl
	return params
}
func (params *UpdateTaskReservationParams) SetConferenceRecord(ConferenceRecord string) *UpdateTaskReservationParams {
	params.ConferenceRecord = &ConferenceRecord
	return params
}
func (params *UpdateTaskReservationParams) SetConferenceRecordingStatusCallback(ConferenceRecordingStatusCallback string) *UpdateTaskReservationParams {
	params.ConferenceRecordingStatusCallback = &ConferenceRecordingStatusCallback
	return params
}
func (params *UpdateTaskReservationParams) SetConferenceRecordingStatusCallbackMethod(ConferenceRecordingStatusCallbackMethod string) *UpdateTaskReservationParams {
	params.ConferenceRecordingStatusCallbackMethod = &ConferenceRecordingStatusCallbackMethod
	return params
}
func (params *UpdateTaskReservationParams) SetConferenceStatusCallback(ConferenceStatusCallback string) *UpdateTaskReservationParams {
	params.ConferenceStatusCallback = &ConferenceStatusCallback
	return params
}
func (params *UpdateTaskReservationParams) SetConferenceStatusCallbackEvent(ConferenceStatusCallbackEvent []string) *UpdateTaskReservationParams {
	params.ConferenceStatusCallbackEvent = &ConferenceStatusCallbackEvent
	return params
}
func (params *UpdateTaskReservationParams) SetConferenceStatusCallbackMethod(ConferenceStatusCallbackMethod string) *UpdateTaskReservationParams {
	params.ConferenceStatusCallbackMethod = &ConferenceStatusCallbackMethod
	return params
}
func (params *UpdateTaskReservationParams) SetConferenceTrim(ConferenceTrim string) *UpdateTaskReservationParams {
	params.ConferenceTrim = &ConferenceTrim
	return params
}
func (params *UpdateTaskReservationParams) SetDequeueFrom(DequeueFrom string) *UpdateTaskReservationParams {
	params.DequeueFrom = &DequeueFrom
	return params
}
func (params *UpdateTaskReservationParams) SetDequeuePostWorkActivitySid(DequeuePostWorkActivitySid string) *UpdateTaskReservationParams {
	params.DequeuePostWorkActivitySid = &DequeuePostWorkActivitySid
	return params
}
func (params *UpdateTaskReservationParams) SetDequeueRecord(DequeueRecord string) *UpdateTaskReservationParams {
	params.DequeueRecord = &DequeueRecord
	return params
}
func (params *UpdateTaskReservationParams) SetDequeueStatusCallbackEvent(DequeueStatusCallbackEvent []string) *UpdateTaskReservationParams {
	params.DequeueStatusCallbackEvent = &DequeueStatusCallbackEvent
	return params
}
func (params *UpdateTaskReservationParams) SetDequeueStatusCallbackUrl(DequeueStatusCallbackUrl string) *UpdateTaskReservationParams {
	params.DequeueStatusCallbackUrl = &DequeueStatusCallbackUrl
	return params
}
func (params *UpdateTaskReservationParams) SetDequeueTimeout(DequeueTimeout int) *UpdateTaskReservationParams {
	params.DequeueTimeout = &DequeueTimeout
	return params
}
func (params *UpdateTaskReservationParams) SetDequeueTo(DequeueTo string) *UpdateTaskReservationParams {
	params.DequeueTo = &DequeueTo
	return params
}
func (params *UpdateTaskReservationParams) SetEarlyMedia(EarlyMedia bool) *UpdateTaskReservationParams {
	params.EarlyMedia = &EarlyMedia
	return params
}
func (params *UpdateTaskReservationParams) SetEndConferenceOnCustomerExit(EndConferenceOnCustomerExit bool) *UpdateTaskReservationParams {
	params.EndConferenceOnCustomerExit = &EndConferenceOnCustomerExit
	return params
}
func (params *UpdateTaskReservationParams) SetEndConferenceOnExit(EndConferenceOnExit bool) *UpdateTaskReservationParams {
	params.EndConferenceOnExit = &EndConferenceOnExit
	return params
}
func (params *UpdateTaskReservationParams) SetFrom(From string) *UpdateTaskReservationParams {
	params.From = &From
	return params
}
func (params *UpdateTaskReservationParams) SetInstruction(Instruction string) *UpdateTaskReservationParams {
	params.Instruction = &Instruction
	return params
}
func (params *UpdateTaskReservationParams) SetMaxParticipants(MaxParticipants int) *UpdateTaskReservationParams {
	params.MaxParticipants = &MaxParticipants
	return params
}
func (params *UpdateTaskReservationParams) SetMuted(Muted bool) *UpdateTaskReservationParams {
	params.Muted = &Muted
	return params
}
func (params *UpdateTaskReservationParams) SetPostWorkActivitySid(PostWorkActivitySid string) *UpdateTaskReservationParams {
	params.PostWorkActivitySid = &PostWorkActivitySid
	return params
}
func (params *UpdateTaskReservationParams) SetRecord(Record bool) *UpdateTaskReservationParams {
	params.Record = &Record
	return params
}
func (params *UpdateTaskReservationParams) SetRecordingChannels(RecordingChannels string) *UpdateTaskReservationParams {
	params.RecordingChannels = &RecordingChannels
	return params
}
func (params *UpdateTaskReservationParams) SetRecordingStatusCallback(RecordingStatusCallback string) *UpdateTaskReservationParams {
	params.RecordingStatusCallback = &RecordingStatusCallback
	return params
}
func (params *UpdateTaskReservationParams) SetRecordingStatusCallbackMethod(RecordingStatusCallbackMethod string) *UpdateTaskReservationParams {
	params.RecordingStatusCallbackMethod = &RecordingStatusCallbackMethod
	return params
}
func (params *UpdateTaskReservationParams) SetRedirectAccept(RedirectAccept bool) *UpdateTaskReservationParams {
	params.RedirectAccept = &RedirectAccept
	return params
}
func (params *UpdateTaskReservationParams) SetRedirectCallSid(RedirectCallSid string) *UpdateTaskReservationParams {
	params.RedirectCallSid = &RedirectCallSid
	return params
}
func (params *UpdateTaskReservationParams) SetRedirectUrl(RedirectUrl string) *UpdateTaskReservationParams {
	params.RedirectUrl = &RedirectUrl
	return params
}
func (params *UpdateTaskReservationParams) SetRegion(Region string) *UpdateTaskReservationParams {
	params.Region = &Region
	return params
}
func (params *UpdateTaskReservationParams) SetReservationStatus(ReservationStatus string) *UpdateTaskReservationParams {
	params.ReservationStatus = &ReservationStatus
	return params
}
func (params *UpdateTaskReservationParams) SetSipAuthPassword(SipAuthPassword string) *UpdateTaskReservationParams {
	params.SipAuthPassword = &SipAuthPassword
	return params
}
func (params *UpdateTaskReservationParams) SetSipAuthUsername(SipAuthUsername string) *UpdateTaskReservationParams {
	params.SipAuthUsername = &SipAuthUsername
	return params
}
func (params *UpdateTaskReservationParams) SetStartConferenceOnEnter(StartConferenceOnEnter bool) *UpdateTaskReservationParams {
	params.StartConferenceOnEnter = &StartConferenceOnEnter
	return params
}
func (params *UpdateTaskReservationParams) SetStatusCallback(StatusCallback string) *UpdateTaskReservationParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *UpdateTaskReservationParams) SetStatusCallbackEvent(StatusCallbackEvent []string) *UpdateTaskReservationParams {
	params.StatusCallbackEvent = &StatusCallbackEvent
	return params
}
func (params *UpdateTaskReservationParams) SetStatusCallbackMethod(StatusCallbackMethod string) *UpdateTaskReservationParams {
	params.StatusCallbackMethod = &StatusCallbackMethod
	return params
}
func (params *UpdateTaskReservationParams) SetSupervisor(Supervisor string) *UpdateTaskReservationParams {
	params.Supervisor = &Supervisor
	return params
}
func (params *UpdateTaskReservationParams) SetSupervisorMode(SupervisorMode string) *UpdateTaskReservationParams {
	params.SupervisorMode = &SupervisorMode
	return params
}
func (params *UpdateTaskReservationParams) SetTimeout(Timeout int) *UpdateTaskReservationParams {
	params.Timeout = &Timeout
	return params
}
func (params *UpdateTaskReservationParams) SetTo(To string) *UpdateTaskReservationParams {
	params.To = &To
	return params
}
func (params *UpdateTaskReservationParams) SetWaitMethod(WaitMethod string) *UpdateTaskReservationParams {
	params.WaitMethod = &WaitMethod
	return params
}
func (params *UpdateTaskReservationParams) SetWaitUrl(WaitUrl string) *UpdateTaskReservationParams {
	params.WaitUrl = &WaitUrl
	return params
}
func (params *UpdateTaskReservationParams) SetWorkerActivitySid(WorkerActivitySid string) *UpdateTaskReservationParams {
	params.WorkerActivitySid = &WorkerActivitySid
	return params
}

func (c *ApiService) UpdateTaskReservation(WorkspaceSid string, TaskSid string, Sid string, params *UpdateTaskReservationParams) (*TaskrouterV1WorkspaceTaskTaskReservation, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Tasks/{TaskSid}/Reservations/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"TaskSid"+"}", TaskSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Beep != nil {
		data.Set("Beep", *params.Beep)
	}
	if params != nil && params.BeepOnCustomerEntrance != nil {
		data.Set("BeepOnCustomerEntrance", fmt.Sprint(*params.BeepOnCustomerEntrance))
	}
	if params != nil && params.CallAccept != nil {
		data.Set("CallAccept", fmt.Sprint(*params.CallAccept))
	}
	if params != nil && params.CallFrom != nil {
		data.Set("CallFrom", *params.CallFrom)
	}
	if params != nil && params.CallRecord != nil {
		data.Set("CallRecord", *params.CallRecord)
	}
	if params != nil && params.CallStatusCallbackUrl != nil {
		data.Set("CallStatusCallbackUrl", *params.CallStatusCallbackUrl)
	}
	if params != nil && params.CallTimeout != nil {
		data.Set("CallTimeout", fmt.Sprint(*params.CallTimeout))
	}
	if params != nil && params.CallTo != nil {
		data.Set("CallTo", *params.CallTo)
	}
	if params != nil && params.CallUrl != nil {
		data.Set("CallUrl", *params.CallUrl)
	}
	if params != nil && params.ConferenceRecord != nil {
		data.Set("ConferenceRecord", *params.ConferenceRecord)
	}
	if params != nil && params.ConferenceRecordingStatusCallback != nil {
		data.Set("ConferenceRecordingStatusCallback", *params.ConferenceRecordingStatusCallback)
	}
	if params != nil && params.ConferenceRecordingStatusCallbackMethod != nil {
		data.Set("ConferenceRecordingStatusCallbackMethod", *params.ConferenceRecordingStatusCallbackMethod)
	}
	if params != nil && params.ConferenceStatusCallback != nil {
		data.Set("ConferenceStatusCallback", *params.ConferenceStatusCallback)
	}
	if params != nil && params.ConferenceStatusCallbackEvent != nil {
		for _, item := range *params.ConferenceStatusCallbackEvent {
			data.Add("ConferenceStatusCallbackEvent", item)
		}
	}
	if params != nil && params.ConferenceStatusCallbackMethod != nil {
		data.Set("ConferenceStatusCallbackMethod", *params.ConferenceStatusCallbackMethod)
	}
	if params != nil && params.ConferenceTrim != nil {
		data.Set("ConferenceTrim", *params.ConferenceTrim)
	}
	if params != nil && params.DequeueFrom != nil {
		data.Set("DequeueFrom", *params.DequeueFrom)
	}
	if params != nil && params.DequeuePostWorkActivitySid != nil {
		data.Set("DequeuePostWorkActivitySid", *params.DequeuePostWorkActivitySid)
	}
	if params != nil && params.DequeueRecord != nil {
		data.Set("DequeueRecord", *params.DequeueRecord)
	}
	if params != nil && params.DequeueStatusCallbackEvent != nil {
		for _, item := range *params.DequeueStatusCallbackEvent {
			data.Add("DequeueStatusCallbackEvent", item)
		}
	}
	if params != nil && params.DequeueStatusCallbackUrl != nil {
		data.Set("DequeueStatusCallbackUrl", *params.DequeueStatusCallbackUrl)
	}
	if params != nil && params.DequeueTimeout != nil {
		data.Set("DequeueTimeout", fmt.Sprint(*params.DequeueTimeout))
	}
	if params != nil && params.DequeueTo != nil {
		data.Set("DequeueTo", *params.DequeueTo)
	}
	if params != nil && params.EarlyMedia != nil {
		data.Set("EarlyMedia", fmt.Sprint(*params.EarlyMedia))
	}
	if params != nil && params.EndConferenceOnCustomerExit != nil {
		data.Set("EndConferenceOnCustomerExit", fmt.Sprint(*params.EndConferenceOnCustomerExit))
	}
	if params != nil && params.EndConferenceOnExit != nil {
		data.Set("EndConferenceOnExit", fmt.Sprint(*params.EndConferenceOnExit))
	}
	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.Instruction != nil {
		data.Set("Instruction", *params.Instruction)
	}
	if params != nil && params.MaxParticipants != nil {
		data.Set("MaxParticipants", fmt.Sprint(*params.MaxParticipants))
	}
	if params != nil && params.Muted != nil {
		data.Set("Muted", fmt.Sprint(*params.Muted))
	}
	if params != nil && params.PostWorkActivitySid != nil {
		data.Set("PostWorkActivitySid", *params.PostWorkActivitySid)
	}
	if params != nil && params.Record != nil {
		data.Set("Record", fmt.Sprint(*params.Record))
	}
	if params != nil && params.RecordingChannels != nil {
		data.Set("RecordingChannels", *params.RecordingChannels)
	}
	if params != nil && params.RecordingStatusCallback != nil {
		data.Set("RecordingStatusCallback", *params.RecordingStatusCallback)
	}
	if params != nil && params.RecordingStatusCallbackMethod != nil {
		data.Set("RecordingStatusCallbackMethod", *params.RecordingStatusCallbackMethod)
	}
	if params != nil && params.RedirectAccept != nil {
		data.Set("RedirectAccept", fmt.Sprint(*params.RedirectAccept))
	}
	if params != nil && params.RedirectCallSid != nil {
		data.Set("RedirectCallSid", *params.RedirectCallSid)
	}
	if params != nil && params.RedirectUrl != nil {
		data.Set("RedirectUrl", *params.RedirectUrl)
	}
	if params != nil && params.Region != nil {
		data.Set("Region", *params.Region)
	}
	if params != nil && params.ReservationStatus != nil {
		data.Set("ReservationStatus", *params.ReservationStatus)
	}
	if params != nil && params.SipAuthPassword != nil {
		data.Set("SipAuthPassword", *params.SipAuthPassword)
	}
	if params != nil && params.SipAuthUsername != nil {
		data.Set("SipAuthUsername", *params.SipAuthUsername)
	}
	if params != nil && params.StartConferenceOnEnter != nil {
		data.Set("StartConferenceOnEnter", fmt.Sprint(*params.StartConferenceOnEnter))
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.StatusCallbackEvent != nil {
		for _, item := range *params.StatusCallbackEvent {
			data.Add("StatusCallbackEvent", item)
		}
	}
	if params != nil && params.StatusCallbackMethod != nil {
		data.Set("StatusCallbackMethod", *params.StatusCallbackMethod)
	}
	if params != nil && params.Supervisor != nil {
		data.Set("Supervisor", *params.Supervisor)
	}
	if params != nil && params.SupervisorMode != nil {
		data.Set("SupervisorMode", *params.SupervisorMode)
	}
	if params != nil && params.Timeout != nil {
		data.Set("Timeout", fmt.Sprint(*params.Timeout))
	}
	if params != nil && params.To != nil {
		data.Set("To", *params.To)
	}
	if params != nil && params.WaitMethod != nil {
		data.Set("WaitMethod", *params.WaitMethod)
	}
	if params != nil && params.WaitUrl != nil {
		data.Set("WaitUrl", *params.WaitUrl)
	}
	if params != nil && params.WorkerActivitySid != nil {
		data.Set("WorkerActivitySid", *params.WorkerActivitySid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1WorkspaceTaskTaskReservation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
