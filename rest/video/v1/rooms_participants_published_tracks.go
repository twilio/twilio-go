/*
 * Twilio - Video
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

// Returns a single Track resource represented by TrackName or SID.
func (c *ApiService) FetchRoomParticipantPublishedTrack(RoomSid string, ParticipantSid string, Sid string) (*VideoV1RoomRoomParticipantRoomParticipantPublishedTrack, error) {
	path := "/v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/PublishedTracks/{Sid}"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)
	path = strings.Replace(path, "{"+"ParticipantSid"+"}", ParticipantSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1RoomRoomParticipantRoomParticipantPublishedTrack{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRoomParticipantPublishedTrack'
type ListRoomParticipantPublishedTrackParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListRoomParticipantPublishedTrackParams) SetPageSize(PageSize int) *ListRoomParticipantPublishedTrackParams {
	params.PageSize = &PageSize
	return params
}

// Returns a list of tracks associated with a given Participant. Only &#x60;currently&#x60; Published Tracks are in the list resource.
func (c *ApiService) ListRoomParticipantPublishedTrack(RoomSid string, ParticipantSid string, params *ListRoomParticipantPublishedTrackParams) (*ListRoomParticipantPublishedTrackResponse, error) {
	path := "/v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/PublishedTracks"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)
	path = strings.Replace(path, "{"+"ParticipantSid"+"}", ParticipantSid, -1)

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

	ps := &ListRoomParticipantPublishedTrackResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of  records from the API. Request is executed immediately.
func (c *ApiService) RoomsParticipantsPublishedTracksPage(RoomSid string, ParticipantSid string, params *ListRoomParticipantPublishedTrackParams, pageToken string, pageNumber string) (*ListRoomParticipantPublishedTrackResponse, error) {
	path := "/v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/PublishedTracks"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)
	path = strings.Replace(path, "{"+"ParticipantSid"+"}", ParticipantSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	data.Set("PageToken", pageToken)
	data.Set("PageNumber", pageNumber)

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRoomParticipantPublishedTrackResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Lists RoomsParticipantsPublishedTracks records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) RoomsParticipantsPublishedTracksList(RoomSid string, ParticipantSid string, params *ListRoomParticipantPublishedTrackParams, limit int) ([]ListRoomParticipantPublishedTrackResponse, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	response, err := c.ListRoomParticipantPublishedTrack(RoomSid, ParticipantSid, params)
	if err != nil {
		return nil, err
	}

	page := client.NewPage(c.baseURL, response)

	resp := c.requestHandler.List(page, limit, 0)
	ret := make([]ListRoomParticipantPublishedTrackResponse, len(resp))

	for i := range resp {
		jsonStr, _ := json.Marshal(resp[i])
		ps := ListRoomParticipantPublishedTrackResponse{}
		if err := json.Unmarshal(jsonStr, &ps); err != nil {
			return ret, err
		}

		ret[i] = ps
	}

	return ret, nil
}

//Streams RoomsParticipantsPublishedTracks records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) RoomsParticipantsPublishedTracksStream(RoomSid string, ParticipantSid string, params *ListRoomParticipantPublishedTrackParams, limit int) (chan interface{}, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	response, err := c.ListRoomParticipantPublishedTrack(RoomSid, ParticipantSid, params)
	if err != nil {
		return nil, err
	}

	page := client.NewPage(c.baseURL, response)

	ps := ListRoomParticipantPublishedTrackResponse{}
	return c.requestHandler.Stream(page, limit, 0, ps), nil
}
