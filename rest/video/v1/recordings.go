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
	"time"

	"github.com/twilio/twilio-go/client"
)

// Delete a Recording resource identified by a Recording SID.
func (c *ApiService) DeleteRecording(Sid string) error {
	path := "/v1/Recordings/{Sid}"
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

// Returns a single Recording resource identified by a Recording SID.
func (c *ApiService) FetchRecording(Sid string) (*VideoV1Recording, error) {
	path := "/v1/Recordings/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1Recording{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRecording'
type ListRecordingParams struct {
	// Read only the recordings that have this status. Can be: `processing`, `completed`, or `deleted`.
	Status *string `json:"Status,omitempty"`
	// Read only the recordings that have this `source_sid`.
	SourceSid *string `json:"SourceSid,omitempty"`
	// Read only recordings with this `grouping_sid`, which may include a `participant_sid` and/or a `room_sid`.
	GroupingSid *[]string `json:"GroupingSid,omitempty"`
	// Read only recordings that started on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time with time zone.
	DateCreatedAfter *time.Time `json:"DateCreatedAfter,omitempty"`
	// Read only recordings that started before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time with time zone, given as `YYYY-MM-DDThh:mm:ss+|-hh:mm` or `YYYY-MM-DDThh:mm:ssZ`.
	DateCreatedBefore *time.Time `json:"DateCreatedBefore,omitempty"`
	// Read only recordings that have this media type. Can be either `audio` or `video`.
	MediaType *string `json:"MediaType,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListRecordingParams) SetStatus(Status string) *ListRecordingParams {
	params.Status = &Status
	return params
}
func (params *ListRecordingParams) SetSourceSid(SourceSid string) *ListRecordingParams {
	params.SourceSid = &SourceSid
	return params
}
func (params *ListRecordingParams) SetGroupingSid(GroupingSid []string) *ListRecordingParams {
	params.GroupingSid = &GroupingSid
	return params
}
func (params *ListRecordingParams) SetDateCreatedAfter(DateCreatedAfter time.Time) *ListRecordingParams {
	params.DateCreatedAfter = &DateCreatedAfter
	return params
}
func (params *ListRecordingParams) SetDateCreatedBefore(DateCreatedBefore time.Time) *ListRecordingParams {
	params.DateCreatedBefore = &DateCreatedBefore
	return params
}
func (params *ListRecordingParams) SetMediaType(MediaType string) *ListRecordingParams {
	params.MediaType = &MediaType
	return params
}
func (params *ListRecordingParams) SetPageSize(PageSize int) *ListRecordingParams {
	params.PageSize = &PageSize
	return params
}

// List of all Track recordings.
func (c *ApiService) ListRecording(params *ListRecordingParams) (*ListRecordingResponse, error) {
	path := "/v1/Recordings"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.SourceSid != nil {
		data.Set("SourceSid", *params.SourceSid)
	}
	if params != nil && params.GroupingSid != nil {
		for _, item := range *params.GroupingSid {
			data.Add("GroupingSid", item)
		}
	}
	if params != nil && params.DateCreatedAfter != nil {
		data.Set("DateCreatedAfter", fmt.Sprint((*params.DateCreatedAfter).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedBefore != nil {
		data.Set("DateCreatedBefore", fmt.Sprint((*params.DateCreatedBefore).Format(time.RFC3339)))
	}
	if params != nil && params.MediaType != nil {
		data.Set("MediaType", *params.MediaType)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRecordingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of Recording records from the API. Request is executed immediately.
func (c *ApiService) RecordingPage(params *ListRecordingParams, pageToken string, pageNumber string, pageSize string) *client.Page {
	path := "/v1/Recordings"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.SourceSid != nil {
		data.Set("SourceSid", *params.SourceSid)
	}
	if params != nil && params.GroupingSid != nil {
		for _, item := range *params.GroupingSid {
			data.Add("GroupingSid", item)
		}
	}
	if params != nil && params.DateCreatedAfter != nil {
		data.Set("DateCreatedAfter", fmt.Sprint((*params.DateCreatedAfter).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedBefore != nil {
		data.Set("DateCreatedBefore", fmt.Sprint((*params.DateCreatedBefore).Format(time.RFC3339)))
	}
	if params != nil && params.MediaType != nil {
		data.Set("MediaType", *params.MediaType)
	}
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

//Streams Recording records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) RecordingStream(params *ListRecordingParams, meta client.PaginationData) chan map[string]interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.RecordingPage(params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.Stream(page, limits.Limit, limits.PageLimit)
}

//Lists Recording records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) RecordingList(params *ListRecordingParams, meta client.PaginationData) []interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.RecordingPage(params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.List(page, limits.Limit, limits.PageLimit)
}
