/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Video
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateRoom'
type CreateRoomParams struct {
	// Deprecated, now always considered to be true.
	EnableTurn *bool `json:"EnableTurn,omitempty"`
	//
	Type *string `json:"Type,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used as a `room_sid` in place of the resource's `sid` in the URL to address the resource, assuming it does not contain any [reserved characters](https://tools.ietf.org/html/rfc3986#section-2.2) that would need to be URL encoded. This value is unique for `in-progress` rooms. SDK clients can use this name to connect to the room. REST API clients can use this name in place of the Room SID to interact with the room as long as the room is `in-progress`.
	UniqueName *string `json:"UniqueName,omitempty"`
	// The URL Twilio should call using the `status_callback_method` to send status information to your application on every room event. See [Status Callbacks](https://www.twilio.com/docs/video/api/status-callbacks) for more info.
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// The HTTP method Twilio should use to call `status_callback`. Can be `POST` or `GET`.
	StatusCallbackMethod *string `json:"StatusCallbackMethod,omitempty"`
	// The maximum number of concurrent Participants allowed in the room. The maximum allowed value is 50.
	MaxParticipants *int `json:"MaxParticipants,omitempty"`
	// Whether to start recording when Participants connect.
	RecordParticipantsOnConnect *bool `json:"RecordParticipantsOnConnect,omitempty"`
	// Whether to start transcriptions when Participants connect. If TranscriptionsConfiguration is not provided, default settings will be used.
	TranscribeParticipantsOnConnect *bool `json:"TranscribeParticipantsOnConnect,omitempty"`
	// An array of the video codecs that are supported when publishing a track in the room.  Can be: `VP8` and `H264`.
	VideoCodecs *[]string `json:"VideoCodecs,omitempty"`
	// The region for the Room's media server.  Can be one of the [available Media Regions](https://www.twilio.com/docs/video/ip-addresses#group-rooms-media-servers).
	MediaRegion *string `json:"MediaRegion,omitempty"`
	// A collection of Recording Rules that describe how to include or exclude matching tracks for recording
	RecordingRules *interface{} `json:"RecordingRules,omitempty"`
	// A collection of properties that describe transcription behaviour. If TranscribeParticipantsOnConnect is set to true and TranscriptionsConfiguration is not provided, default settings will be used.
	TranscriptionsConfiguration *map[string]interface{} `json:"TranscriptionsConfiguration,omitempty"`
	// When set to true, indicates that the participants in the room will only publish audio. No video tracks will be allowed.
	AudioOnly *bool `json:"AudioOnly,omitempty"`
	// The maximum number of seconds a Participant can be connected to the room. The maximum possible value is 86400 seconds (24 hours). The default is 14400 seconds (4 hours).
	MaxParticipantDuration *int `json:"MaxParticipantDuration,omitempty"`
	// Configures how long (in minutes) a room will remain active after last participant leaves. Valid values range from 1 to 60 minutes (no fractions).
	EmptyRoomTimeout *int `json:"EmptyRoomTimeout,omitempty"`
	// Configures how long (in minutes) a room will remain active if no one joins. Valid values range from 1 to 60 minutes (no fractions).
	UnusedRoomTimeout *int `json:"UnusedRoomTimeout,omitempty"`
	// When set to true, indicated that this is the large room.
	LargeRoom *bool `json:"LargeRoom,omitempty"`
}

func (params *CreateRoomParams) SetEnableTurn(EnableTurn bool) *CreateRoomParams {
	params.EnableTurn = &EnableTurn
	return params
}
func (params *CreateRoomParams) SetType(Type string) *CreateRoomParams {
	params.Type = &Type
	return params
}
func (params *CreateRoomParams) SetUniqueName(UniqueName string) *CreateRoomParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *CreateRoomParams) SetStatusCallback(StatusCallback string) *CreateRoomParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *CreateRoomParams) SetStatusCallbackMethod(StatusCallbackMethod string) *CreateRoomParams {
	params.StatusCallbackMethod = &StatusCallbackMethod
	return params
}
func (params *CreateRoomParams) SetMaxParticipants(MaxParticipants int) *CreateRoomParams {
	params.MaxParticipants = &MaxParticipants
	return params
}
func (params *CreateRoomParams) SetRecordParticipantsOnConnect(RecordParticipantsOnConnect bool) *CreateRoomParams {
	params.RecordParticipantsOnConnect = &RecordParticipantsOnConnect
	return params
}
func (params *CreateRoomParams) SetTranscribeParticipantsOnConnect(TranscribeParticipantsOnConnect bool) *CreateRoomParams {
	params.TranscribeParticipantsOnConnect = &TranscribeParticipantsOnConnect
	return params
}
func (params *CreateRoomParams) SetVideoCodecs(VideoCodecs []string) *CreateRoomParams {
	params.VideoCodecs = &VideoCodecs
	return params
}
func (params *CreateRoomParams) SetMediaRegion(MediaRegion string) *CreateRoomParams {
	params.MediaRegion = &MediaRegion
	return params
}
func (params *CreateRoomParams) SetRecordingRules(RecordingRules interface{}) *CreateRoomParams {
	params.RecordingRules = &RecordingRules
	return params
}
func (params *CreateRoomParams) SetTranscriptionsConfiguration(TranscriptionsConfiguration map[string]interface{}) *CreateRoomParams {
	params.TranscriptionsConfiguration = &TranscriptionsConfiguration
	return params
}
func (params *CreateRoomParams) SetAudioOnly(AudioOnly bool) *CreateRoomParams {
	params.AudioOnly = &AudioOnly
	return params
}
func (params *CreateRoomParams) SetMaxParticipantDuration(MaxParticipantDuration int) *CreateRoomParams {
	params.MaxParticipantDuration = &MaxParticipantDuration
	return params
}
func (params *CreateRoomParams) SetEmptyRoomTimeout(EmptyRoomTimeout int) *CreateRoomParams {
	params.EmptyRoomTimeout = &EmptyRoomTimeout
	return params
}
func (params *CreateRoomParams) SetUnusedRoomTimeout(UnusedRoomTimeout int) *CreateRoomParams {
	params.UnusedRoomTimeout = &UnusedRoomTimeout
	return params
}
func (params *CreateRoomParams) SetLargeRoom(LargeRoom bool) *CreateRoomParams {
	params.LargeRoom = &LargeRoom
	return params
}

//
func (c *ApiService) CreateRoom(params *CreateRoomParams) (*VideoV1Room, error) {
	path := "/v1/Rooms"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.EnableTurn != nil {
		data.Set("EnableTurn", fmt.Sprint(*params.EnableTurn))
	}
	if params != nil && params.Type != nil {
		data.Set("Type", fmt.Sprint(*params.Type))
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.StatusCallbackMethod != nil {
		data.Set("StatusCallbackMethod", *params.StatusCallbackMethod)
	}
	if params != nil && params.MaxParticipants != nil {
		data.Set("MaxParticipants", fmt.Sprint(*params.MaxParticipants))
	}
	if params != nil && params.RecordParticipantsOnConnect != nil {
		data.Set("RecordParticipantsOnConnect", fmt.Sprint(*params.RecordParticipantsOnConnect))
	}
	if params != nil && params.TranscribeParticipantsOnConnect != nil {
		data.Set("TranscribeParticipantsOnConnect", fmt.Sprint(*params.TranscribeParticipantsOnConnect))
	}
	if params != nil && params.VideoCodecs != nil {
		for _, item := range *params.VideoCodecs {
			data.Add("VideoCodecs", item)
		}
	}
	if params != nil && params.MediaRegion != nil {
		data.Set("MediaRegion", *params.MediaRegion)
	}
	if params != nil && params.RecordingRules != nil {
		v, err := json.Marshal(params.RecordingRules)

		if err != nil {
			return nil, err
		}

		data.Set("RecordingRules", string(v))
	}
	if params != nil && params.TranscriptionsConfiguration != nil {
		v, err := json.Marshal(params.TranscriptionsConfiguration)

		if err != nil {
			return nil, err
		}

		data.Set("TranscriptionsConfiguration", string(v))
	}
	if params != nil && params.AudioOnly != nil {
		data.Set("AudioOnly", fmt.Sprint(*params.AudioOnly))
	}
	if params != nil && params.MaxParticipantDuration != nil {
		data.Set("MaxParticipantDuration", fmt.Sprint(*params.MaxParticipantDuration))
	}
	if params != nil && params.EmptyRoomTimeout != nil {
		data.Set("EmptyRoomTimeout", fmt.Sprint(*params.EmptyRoomTimeout))
	}
	if params != nil && params.UnusedRoomTimeout != nil {
		data.Set("UnusedRoomTimeout", fmt.Sprint(*params.UnusedRoomTimeout))
	}
	if params != nil && params.LargeRoom != nil {
		data.Set("LargeRoom", fmt.Sprint(*params.LargeRoom))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1Room{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//
func (c *ApiService) FetchRoom(Sid string) (*VideoV1Room, error) {
	path := "/v1/Rooms/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1Room{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRoom'
type ListRoomParams struct {
	// Read only the rooms with this status. Can be: `in-progress` (default) or `completed`
	Status *string `json:"Status,omitempty"`
	// Read only rooms with the this `unique_name`.
	UniqueName *string `json:"UniqueName,omitempty"`
	// Read only rooms that started on or after this date, given as `YYYY-MM-DD`.
	DateCreatedAfter *time.Time `json:"DateCreatedAfter,omitempty"`
	// Read only rooms that started before this date, given as `YYYY-MM-DD`.
	DateCreatedBefore *time.Time `json:"DateCreatedBefore,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListRoomParams) SetStatus(Status string) *ListRoomParams {
	params.Status = &Status
	return params
}
func (params *ListRoomParams) SetUniqueName(UniqueName string) *ListRoomParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *ListRoomParams) SetDateCreatedAfter(DateCreatedAfter time.Time) *ListRoomParams {
	params.DateCreatedAfter = &DateCreatedAfter
	return params
}
func (params *ListRoomParams) SetDateCreatedBefore(DateCreatedBefore time.Time) *ListRoomParams {
	params.DateCreatedBefore = &DateCreatedBefore
	return params
}
func (params *ListRoomParams) SetPageSize(PageSize int) *ListRoomParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListRoomParams) SetLimit(Limit int) *ListRoomParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Room records from the API. Request is executed immediately.
func (c *ApiService) PageRoom(params *ListRoomParams, pageToken, pageNumber string) (*ListRoomResponse, error) {
	path := "/v1/Rooms"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Status != nil {
		data.Set("Status", fmt.Sprint(*params.Status))
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.DateCreatedAfter != nil {
		data.Set("DateCreatedAfter", fmt.Sprint((*params.DateCreatedAfter).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedBefore != nil {
		data.Set("DateCreatedBefore", fmt.Sprint((*params.DateCreatedBefore).Format(time.RFC3339)))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRoomResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Room records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListRoom(params *ListRoomParams) ([]VideoV1Room, error) {
	response, errors := c.StreamRoom(params)

	records := make([]VideoV1Room, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Room records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamRoom(params *ListRoomParams) (chan VideoV1Room, chan error) {
	if params == nil {
		params = &ListRoomParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan VideoV1Room, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageRoom(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamRoom(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamRoom(response *ListRoomResponse, params *ListRoomParams, recordChannel chan VideoV1Room, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Rooms
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListRoomResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListRoomResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListRoomResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRoomResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateRoom'
type UpdateRoomParams struct {
	//
	Status *string `json:"Status,omitempty"`
}

func (params *UpdateRoomParams) SetStatus(Status string) *UpdateRoomParams {
	params.Status = &Status
	return params
}

//
func (c *ApiService) UpdateRoom(Sid string, params *UpdateRoomParams) (*VideoV1Room, error) {
	path := "/v1/Rooms/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Status != nil {
		data.Set("Status", fmt.Sprint(*params.Status))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1Room{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
