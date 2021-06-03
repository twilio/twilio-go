/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"
	"time"

	twilio "github.com/twilio/twilio-go/client"
)

type DefaultApiService struct {
	baseURL        string
	requestHandler *twilio.RequestHandler
}

func NewDefaultApiService(requestHandler *twilio.RequestHandler) *DefaultApiService {
	return &DefaultApiService{
		requestHandler: requestHandler,
		baseURL:        "https://insights.twilio.com",
	}
}

func NewDefaultApiServiceWithClient(client twilio.BaseClient) *DefaultApiService {
	return NewDefaultApiService(twilio.NewRequestHandler(client))
}

func (c *DefaultApiService) FetchCall(Sid string) (*InsightsV1Call, error) {
	path := "/v1/Voice/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &InsightsV1Call{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'FetchSummary'
type FetchSummaryParams struct {
	//
	ProcessingState *string `json:"ProcessingState,omitempty"`
}

func (params *FetchSummaryParams) SetProcessingState(ProcessingState string) *FetchSummaryParams {
	params.ProcessingState = &ProcessingState
	return params
}

func (c *DefaultApiService) FetchSummary(CallSid string, params *FetchSummaryParams) (*InsightsV1CallSummary, error) {
	path := "/v1/Voice/{CallSid}/Summary"
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ProcessingState != nil {
		data.Set("ProcessingState", *params.ProcessingState)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &InsightsV1CallSummary{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Get Video Log Analyzer data for a Room Participant.
func (c *DefaultApiService) FetchVideoParticipantSummary(RoomSid string, ParticipantSid string) (*InsightsV1VideoRoomSummaryVideoParticipantSummary, error) {
	path := "/v1/Video/Rooms/{RoomSid}/Participants/{ParticipantSid}"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)
	path = strings.Replace(path, "{"+"ParticipantSid"+"}", ParticipantSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &InsightsV1VideoRoomSummaryVideoParticipantSummary{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Get Video Log Analyzer data for a Room.
func (c *DefaultApiService) FetchVideoRoomSummary(RoomSid string) (*InsightsV1VideoRoomSummary, error) {
	path := "/v1/Video/Rooms/{RoomSid}"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &InsightsV1VideoRoomSummary{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListEvent'
type ListEventParams struct {
	//
	Edge *string `json:"Edge,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListEventParams) SetEdge(Edge string) *ListEventParams {
	params.Edge = &Edge
	return params
}
func (params *ListEventParams) SetPageSize(PageSize int32) *ListEventParams {
	params.PageSize = &PageSize
	return params
}

func (c *DefaultApiService) ListEvent(CallSid string, params *ListEventParams) (*ListEventResponse, error) {
	path := "/v1/Voice/{CallSid}/Events"
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Edge != nil {
		data.Set("Edge", *params.Edge)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListEventResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListMetric'
type ListMetricParams struct {
	//
	Edge *string `json:"Edge,omitempty"`
	//
	Direction *string `json:"Direction,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListMetricParams) SetEdge(Edge string) *ListMetricParams {
	params.Edge = &Edge
	return params
}
func (params *ListMetricParams) SetDirection(Direction string) *ListMetricParams {
	params.Direction = &Direction
	return params
}
func (params *ListMetricParams) SetPageSize(PageSize int32) *ListMetricParams {
	params.PageSize = &PageSize
	return params
}

func (c *DefaultApiService) ListMetric(CallSid string, params *ListMetricParams) (*ListMetricResponse, error) {
	path := "/v1/Voice/{CallSid}/Metrics"
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Edge != nil {
		data.Set("Edge", *params.Edge)
	}
	if params != nil && params.Direction != nil {
		data.Set("Direction", *params.Direction)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMetricResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListVideoParticipantSummary'
type ListVideoParticipantSummaryParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListVideoParticipantSummaryParams) SetPageSize(PageSize int32) *ListVideoParticipantSummaryParams {
	params.PageSize = &PageSize
	return params
}

// Get a list of room participants.
func (c *DefaultApiService) ListVideoParticipantSummary(RoomSid string, params *ListVideoParticipantSummaryParams) (*ListVideoParticipantSummaryResponse, error) {
	path := "/v1/Video/Rooms/{RoomSid}/Participants"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListVideoParticipantSummaryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListVideoRoomSummary'
type ListVideoRoomSummaryParams struct {
	// Type of room. Can be `go`, `peer_to_peer`, `group`, or `group_small`.
	RoomType *[]string `json:"RoomType,omitempty"`
	// Codecs used by participants in the room. Can be `VP8`, `H264`, or `VP9`.
	Codec *[]string `json:"Codec,omitempty"`
	// Room friendly name.
	RoomName *string `json:"RoomName,omitempty"`
	// Only read rooms that started on or after this ISO 8601 timestamp.
	CreatedAfter *time.Time `json:"CreatedAfter,omitempty"`
	// Only read rooms that started before this ISO 8601 timestamp.
	CreatedBefore *time.Time `json:"CreatedBefore,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListVideoRoomSummaryParams) SetRoomType(RoomType []string) *ListVideoRoomSummaryParams {
	params.RoomType = &RoomType
	return params
}
func (params *ListVideoRoomSummaryParams) SetCodec(Codec []string) *ListVideoRoomSummaryParams {
	params.Codec = &Codec
	return params
}
func (params *ListVideoRoomSummaryParams) SetRoomName(RoomName string) *ListVideoRoomSummaryParams {
	params.RoomName = &RoomName
	return params
}
func (params *ListVideoRoomSummaryParams) SetCreatedAfter(CreatedAfter time.Time) *ListVideoRoomSummaryParams {
	params.CreatedAfter = &CreatedAfter
	return params
}
func (params *ListVideoRoomSummaryParams) SetCreatedBefore(CreatedBefore time.Time) *ListVideoRoomSummaryParams {
	params.CreatedBefore = &CreatedBefore
	return params
}
func (params *ListVideoRoomSummaryParams) SetPageSize(PageSize int32) *ListVideoRoomSummaryParams {
	params.PageSize = &PageSize
	return params
}

// Get a list of Programmable Video Rooms.
func (c *DefaultApiService) ListVideoRoomSummary(params *ListVideoRoomSummaryParams) (*ListVideoRoomSummaryResponse, error) {
	path := "/v1/Video/Rooms"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.RoomType != nil {
		for _, item := range *params.RoomType {
			data.Add("RoomType", item)
		}
	}
	if params != nil && params.Codec != nil {
		for _, item := range *params.Codec {
			data.Add("Codec", item)
		}
	}
	if params != nil && params.RoomName != nil {
		data.Set("RoomName", *params.RoomName)
	}
	if params != nil && params.CreatedAfter != nil {
		data.Set("CreatedAfter", fmt.Sprint((*params.CreatedAfter).Format(time.RFC3339)))
	}
	if params != nil && params.CreatedBefore != nil {
		data.Set("CreatedBefore", fmt.Sprint((*params.CreatedBefore).Format(time.RFC3339)))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListVideoRoomSummaryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
