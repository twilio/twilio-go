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
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/twilio/twilio-go/client"
)

// Delete a Recording resource identified by a Recording SID.
func (c *ApiService) DeleteRecording(Sid string) error {
	return c.DeleteRecordingWithCtx(context.TODO(), Sid)
}

// Delete a Recording resource identified by a Recording SID.
func (c *ApiService) DeleteRecordingWithCtx(ctx context.Context, Sid string) error {
	path := "/v1/Recordings/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Returns a single Recording resource identified by a Recording SID.
func (c *ApiService) FetchRecording(Sid string) (*VideoV1Recording, error) {
	return c.FetchRecordingWithCtx(context.TODO(), Sid)
}

// Returns a single Recording resource identified by a Recording SID.
func (c *ApiService) FetchRecordingWithCtx(ctx context.Context, Sid string) (*VideoV1Recording, error) {
	path := "/v1/Recordings/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
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
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
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
func (params *ListRecordingParams) SetLimit(Limit int) *ListRecordingParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Recording records from the API. Request is executed immediately.
func (c *ApiService) PageRecording(params *ListRecordingParams, pageToken, pageNumber string) (*ListRecordingResponse, error) {
	return c.PageRecordingWithCtx(context.TODO(), params, pageToken, pageNumber)
}

// Retrieve a single page of Recording records from the API. Request is executed immediately.
func (c *ApiService) PageRecordingWithCtx(ctx context.Context, params *ListRecordingParams, pageToken, pageNumber string) (*ListRecordingResponse, error) {
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

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
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

// Lists Recording records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListRecording(params *ListRecordingParams) ([]VideoV1Recording, error) {
	return c.ListRecordingWithCtx(context.TODO(), params)
}

// Lists Recording records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListRecordingWithCtx(ctx context.Context, params *ListRecordingParams) ([]VideoV1Recording, error) {
	response, errors := c.StreamRecordingWithCtx(ctx, params)

	records := make([]VideoV1Recording, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Recording records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamRecording(params *ListRecordingParams) (chan VideoV1Recording, chan error) {
	return c.StreamRecordingWithCtx(context.TODO(), params)
}

// Streams Recording records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamRecordingWithCtx(ctx context.Context, params *ListRecordingParams) (chan VideoV1Recording, chan error) {
	if params == nil {
		params = &ListRecordingParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan VideoV1Recording, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageRecordingWithCtx(ctx, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamRecording(ctx, response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamRecording(ctx context.Context, response *ListRecordingResponse, params *ListRecordingParams, recordChannel chan VideoV1Recording, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Recordings
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNextWithCtx(ctx, c.baseURL, response, c.getNextListRecordingResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListRecordingResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListRecordingResponse(ctx context.Context, nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(ctx, nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRecordingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
