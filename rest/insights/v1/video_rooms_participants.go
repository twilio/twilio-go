/*
 * Twilio - Insights
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
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListVideoParticipantSummaryParams) SetPageSize(PageSize int) *ListVideoParticipantSummaryParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListVideoParticipantSummaryParams) SetLimit(Limit int) *ListVideoParticipantSummaryParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of VideoParticipantSummary records from the API. Request is executed immediately.
func (c *ApiService) PageVideoParticipantSummary(RoomSid string, params *ListVideoParticipantSummaryParams, pageToken string, pageNumber string) (*ListVideoParticipantSummaryResponse, error) {
	path := "/v1/Video/Rooms/{RoomSid}/Participants"

	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)

	data := url.Values{}
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

	ps := &ListVideoParticipantSummaryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists VideoParticipantSummary records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListVideoParticipantSummary(RoomSid string, params *ListVideoParticipantSummaryParams) ([]InsightsV1VideoRoomSummaryVideoParticipantSummary, error) {
	if params == nil {
		params = &ListVideoParticipantSummaryParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageVideoParticipantSummary(RoomSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []InsightsV1VideoRoomSummaryVideoParticipantSummary

	for response != nil {
		records = append(records, response.Participants...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListVideoParticipantSummaryResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListVideoParticipantSummaryResponse)
	}

	return records, err
}

// Streams VideoParticipantSummary records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamVideoParticipantSummary(RoomSid string, params *ListVideoParticipantSummaryParams) (chan InsightsV1VideoRoomSummaryVideoParticipantSummary, error) {
	if params == nil {
		params = &ListVideoParticipantSummaryParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageVideoParticipantSummary(RoomSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan InsightsV1VideoRoomSummaryVideoParticipantSummary, 1)

	go func() {
		for response != nil {
			for item := range response.Participants {
				channel <- response.Participants[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListVideoParticipantSummaryResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListVideoParticipantSummaryResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListVideoParticipantSummaryResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListVideoParticipantSummaryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
