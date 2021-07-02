/*
 * Twilio - Insights
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
	"time"

	"github.com/twilio/twilio-go/client"
)

// Get Video Log Analyzer data for a Room.
func (c *ApiService) FetchVideoRoomSummary(RoomSid string) (*InsightsV1VideoRoomSummary, error) {
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
	PageSize *int `json:"PageSize,omitempty"`
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
func (params *ListVideoRoomSummaryParams) SetPageSize(PageSize int) *ListVideoRoomSummaryParams {
	params.PageSize = &PageSize
	return params
}

// Get a list of Programmable Video Rooms.
func (c *ApiService) ListVideoRoomSummary(params *ListVideoRoomSummaryParams) (*ListVideoRoomSummaryResponse, error) {
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

//Retrieve a single page of VideoRoomSummary records from the API. Request is executed immediately.
func (c *ApiService) VideoRoomSummaryPage(params *ListVideoRoomSummaryParams, pageToken string, pageNumber string) (*client.Page, error) {
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

	data.Set("PageToken", pageToken)
	data.Set("PageNumber", pageNumber)

	response, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	return client.NewPage(c.baseURL, response), nil
}

//Streams VideoRoomSummary records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) VideoRoomSummaryStream(params *ListVideoRoomSummaryParams, limit int) (chan map[string]interface{}, error) {
	if params.PageSize == nil {
		params.SetPageSize(0)
	}
	params.SetPageSize(c.requestHandler.ReadLimits(*params.PageSize, limit))
	page, err := c.VideoRoomSummaryPage(params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.Stream(page, limit, 0), nil
}

//Lists VideoRoomSummary records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) VideoRoomSummaryList(params *ListVideoRoomSummaryParams, limit int) ([]interface{}, error) {
	if params.PageSize == nil {
		params.SetPageSize(0)
	}
	params.SetPageSize(c.requestHandler.ReadLimits(*params.PageSize, limit))
	page, err := c.VideoRoomSummaryPage(params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.List(page, limit, 0), nil
}
