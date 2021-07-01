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

	"github.com/twilio/twilio-go/client"
)

// Get Video Log Analyzer data for a Room Participant.
func (c *ApiService) FetchVideoParticipantSummary(RoomSid string, ParticipantSid string) (*InsightsV1VideoRoomSummaryVideoParticipantSummary, error) {
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

// Optional parameters for the method 'ListVideoParticipantSummary'
type ListVideoParticipantSummaryParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListVideoParticipantSummaryParams) SetPageSize(PageSize int) *ListVideoParticipantSummaryParams {
	params.PageSize = &PageSize
	return params
}

// Get a list of room participants.
func (c *ApiService) ListVideoParticipantSummary(RoomSid string, params *ListVideoParticipantSummaryParams) (*ListVideoParticipantSummaryResponse, error) {
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

//Retrieve a single page of VideoParticipantSummary records from the API. Request is executed immediately.
func (c *ApiService) VideoParticipantSummaryPage(RoomSid string, params *ListVideoParticipantSummaryParams, pageToken string, pageNumber string, pageSize string) *client.Page {
	path := "/v1/Video/Rooms/{RoomSid}/Participants"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	data.Set("PageToken", pageToken)
	data.Set("PageNumber", pageNumber)
	data.Set("PageSize", pageSize)

	response, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil
	}

	return client.NewPage(c.baseURL, response)
}

//Streams VideoParticipantSummary records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) VideoParticipantSummaryStream(RoomSid string, params *ListVideoParticipantSummaryParams, meta client.PaginationData) chan map[string]interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.VideoParticipantSummaryPage(RoomSid, params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.Stream(page, limits.Limit, limits.PageLimit)
}

//Lists VideoParticipantSummary records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) VideoParticipantSummaryList(RoomSid string, params *ListVideoParticipantSummaryParams, meta client.PaginationData) []interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.VideoParticipantSummaryPage(RoomSid, params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.List(page, limits.Limit, limits.PageLimit)
}
