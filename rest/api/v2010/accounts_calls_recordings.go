/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
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

// Optional parameters for the method 'CreateCallRecording'
type CreateCallRecordingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The number of channels used in the recording. Can be: `mono` or `dual` and the default is `mono`. `mono` records all parties of the call into one channel. `dual` records each party of a 2-party call into separate channels.
	RecordingChannels *string `json:"RecordingChannels,omitempty"`
	// The URL we should call using the `recording_status_callback_method` on each recording event specified in  `recording_status_callback_event`. For more information, see [RecordingStatusCallback parameters](https://www.twilio.com/docs/voice/api/recording#recordingstatuscallback).
	RecordingStatusCallback *string `json:"RecordingStatusCallback,omitempty"`
	// The recording status events on which we should call the `recording_status_callback` URL. Can be: `in-progress`, `completed` and `absent` and the default is `completed`. Separate multiple event values with a space.
	RecordingStatusCallbackEvent *[]string `json:"RecordingStatusCallbackEvent,omitempty"`
	// The HTTP method we should use to call `recording_status_callback`. Can be: `GET` or `POST` and the default is `POST`.
	RecordingStatusCallbackMethod *string `json:"RecordingStatusCallbackMethod,omitempty"`
	// The audio track to record for the call. Can be: `inbound`, `outbound` or `both`. The default is `both`. `inbound` records the audio that is received by Twilio. `outbound` records the audio that is generated from Twilio. `both` records the audio that is received and generated by Twilio.
	RecordingTrack *string `json:"RecordingTrack,omitempty"`
	// Whether to trim any leading and trailing silence in the recording. Can be: `trim-silence` or `do-not-trim` and the default is `do-not-trim`. `trim-silence` trims the silence from the beginning and end of the recording and `do-not-trim` does not.
	Trim *string `json:"Trim,omitempty"`
}

func (params *CreateCallRecordingParams) SetPathAccountSid(PathAccountSid string) *CreateCallRecordingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateCallRecordingParams) SetRecordingChannels(RecordingChannels string) *CreateCallRecordingParams {
	params.RecordingChannels = &RecordingChannels
	return params
}
func (params *CreateCallRecordingParams) SetRecordingStatusCallback(RecordingStatusCallback string) *CreateCallRecordingParams {
	params.RecordingStatusCallback = &RecordingStatusCallback
	return params
}
func (params *CreateCallRecordingParams) SetRecordingStatusCallbackEvent(RecordingStatusCallbackEvent []string) *CreateCallRecordingParams {
	params.RecordingStatusCallbackEvent = &RecordingStatusCallbackEvent
	return params
}
func (params *CreateCallRecordingParams) SetRecordingStatusCallbackMethod(RecordingStatusCallbackMethod string) *CreateCallRecordingParams {
	params.RecordingStatusCallbackMethod = &RecordingStatusCallbackMethod
	return params
}
func (params *CreateCallRecordingParams) SetRecordingTrack(RecordingTrack string) *CreateCallRecordingParams {
	params.RecordingTrack = &RecordingTrack
	return params
}
func (params *CreateCallRecordingParams) SetTrim(Trim string) *CreateCallRecordingParams {
	params.Trim = &Trim
	return params
}

// Create a recording for the call
func (c *ApiService) CreateCallRecording(CallSid string, params *CreateCallRecordingParams) (*ApiV2010AccountCallCallRecording, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	if params != nil && params.RecordingChannels != nil {
		data.Set("RecordingChannels", *params.RecordingChannels)
	}
	if params != nil && params.RecordingStatusCallback != nil {
		data.Set("RecordingStatusCallback", *params.RecordingStatusCallback)
	}
	if params != nil && params.RecordingStatusCallbackEvent != nil {
		for _, item := range *params.RecordingStatusCallbackEvent {
			data.Add("RecordingStatusCallbackEvent", item)
		}
	}
	if params != nil && params.RecordingStatusCallbackMethod != nil {
		data.Set("RecordingStatusCallbackMethod", *params.RecordingStatusCallbackMethod)
	}
	if params != nil && params.RecordingTrack != nil {
		data.Set("RecordingTrack", *params.RecordingTrack)
	}
	if params != nil && params.Trim != nil {
		data.Set("Trim", *params.Trim)
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountCallCallRecording{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteCallRecording'
type DeleteCallRecordingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteCallRecordingParams) SetPathAccountSid(PathAccountSid string) *DeleteCallRecordingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete a recording from your account
func (c *ApiService) DeleteCallRecording(CallSid string, Sid string, params *DeleteCallRecordingParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Optional parameters for the method 'FetchCallRecording'
type FetchCallRecordingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchCallRecordingParams) SetPathAccountSid(PathAccountSid string) *FetchCallRecordingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch an instance of a recording for a call
func (c *ApiService) FetchCallRecording(CallSid string, Sid string, params *FetchCallRecordingParams) (*ApiV2010AccountCallCallRecording, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountCallCallRecording{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListCallRecording'
type ListCallRecordingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. You can also specify inequality: `DateCreated<=YYYY-MM-DD` will return recordings generated at or before midnight on a given date, and `DateCreated>=YYYY-MM-DD` returns recordings generated at or after midnight on a date.
	DateCreated *string `json:"DateCreated,omitempty"`
	// The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. You can also specify inequality: `DateCreated<=YYYY-MM-DD` will return recordings generated at or before midnight on a given date, and `DateCreated>=YYYY-MM-DD` returns recordings generated at or after midnight on a date.
	DateCreatedBefore *string `json:"DateCreated<,omitempty"`
	// The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. You can also specify inequality: `DateCreated<=YYYY-MM-DD` will return recordings generated at or before midnight on a given date, and `DateCreated>=YYYY-MM-DD` returns recordings generated at or after midnight on a date.
	DateCreatedAfter *string `json:"DateCreated>,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListCallRecordingParams) SetPathAccountSid(PathAccountSid string) *ListCallRecordingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListCallRecordingParams) SetDateCreated(DateCreated string) *ListCallRecordingParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *ListCallRecordingParams) SetDateCreatedBefore(DateCreatedBefore string) *ListCallRecordingParams {
	params.DateCreatedBefore = &DateCreatedBefore
	return params
}
func (params *ListCallRecordingParams) SetDateCreatedAfter(DateCreatedAfter string) *ListCallRecordingParams {
	params.DateCreatedAfter = &DateCreatedAfter
	return params
}
func (params *ListCallRecordingParams) SetPageSize(PageSize int) *ListCallRecordingParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListCallRecordingParams) SetLimit(Limit int) *ListCallRecordingParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of CallRecording records from the API. Request is executed immediately.
func (c *ApiService) PageCallRecording(CallSid string, params *ListCallRecordingParams, pageToken string, pageNumber string) (*ListCallRecordingResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint(*params.DateCreated))
	}
	if params != nil && params.DateCreatedBefore != nil {
		data.Set("DateCreated<", fmt.Sprint(*params.DateCreatedBefore))
	}
	if params != nil && params.DateCreatedAfter != nil {
		data.Set("DateCreated>", fmt.Sprint(*params.DateCreatedAfter))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}
	headers := make(map[string]interface{})

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageToken != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCallRecordingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists CallRecording records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListCallRecording(CallSid string, params *ListCallRecordingParams) ([]ApiV2010AccountCallCallRecording, error) {
	if params == nil {
		params = &ListCallRecordingParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageCallRecording(CallSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010AccountCallCallRecording

	for response != nil {
		records = append(records, response.Recordings...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListCallRecordingResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListCallRecordingResponse)
	}

	return records, err
}

// Streams CallRecording records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamCallRecording(CallSid string, params *ListCallRecordingParams) (chan ApiV2010AccountCallCallRecording, error) {
	if params == nil {
		params = &ListCallRecordingParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageCallRecording(CallSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010AccountCallCallRecording, 1)

	go func() {
		for response != nil {
			for item := range response.Recordings {
				channel <- response.Recordings[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListCallRecordingResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListCallRecordingResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListCallRecordingResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCallRecordingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateCallRecording'
type UpdateCallRecordingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resource to update.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// Whether to record during a pause. Can be: `skip` or `silence` and the default is `silence`. `skip` does not record during the pause period, while `silence` will replace the actual audio of the call with silence during the pause period. This parameter only applies when setting `status` is set to `paused`.
	PauseBehavior *string `json:"PauseBehavior,omitempty"`
	// The new status of the recording. Can be: `stopped`, `paused`, `in-progress`.
	Status *string `json:"Status,omitempty"`
}

func (params *UpdateCallRecordingParams) SetPathAccountSid(PathAccountSid string) *UpdateCallRecordingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateCallRecordingParams) SetPauseBehavior(PauseBehavior string) *UpdateCallRecordingParams {
	params.PauseBehavior = &PauseBehavior
	return params
}
func (params *UpdateCallRecordingParams) SetStatus(Status string) *UpdateCallRecordingParams {
	params.Status = &Status
	return params
}

// Changes the status of the recording to paused, stopped, or in-progress. Note: Pass &#x60;Twilio.CURRENT&#x60; instead of recording sid to reference current active recording.
func (c *ApiService) UpdateCallRecording(CallSid string, Sid string, params *UpdateCallRecordingParams) (*ApiV2010AccountCallCallRecording, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	if params != nil && params.PauseBehavior != nil {
		data.Set("PauseBehavior", *params.PauseBehavior)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountCallCallRecording{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
