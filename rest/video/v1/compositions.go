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

// Optional parameters for the method 'CreateComposition'
type CreateCompositionParams struct {
	// The SID of the Group Room with the media tracks to be used as composition sources.
	RoomSid *string `json:"RoomSid,omitempty"`
	// An object that describes the video layout of the composition in terms of regions. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info. Please, be aware that either video_layout or audio_sources have to be provided to get a valid creation request
	VideoLayout *map[string]interface{} `json:"VideoLayout,omitempty"`
	// An array of track names from the same group room to merge into the new composition. Can include zero or more track names. The new composition includes all audio sources specified in `audio_sources` except for those specified in `audio_sources_excluded`. The track names in this parameter can include an asterisk as a wild card character, which will match zero or more characters in a track name. For example, `student*` includes `student` as well as `studentTeam`. Please, be aware that either video_layout or audio_sources have to be provided to get a valid creation request
	AudioSources *[]string `json:"AudioSources,omitempty"`
	// An array of track names to exclude. The new composition includes all audio sources specified in `audio_sources` except for those specified in `audio_sources_excluded`. The track names in this parameter can include an asterisk as a wild card character, which will match zero or more characters in a track name. For example, `student*` excludes `student` as well as `studentTeam`. This parameter can also be empty.
	AudioSourcesExcluded *[]string `json:"AudioSourcesExcluded,omitempty"`
	// A string that describes the columns (width) and rows (height) of the generated composed video in pixels. Defaults to `640x480`.  The string's format is `{width}x{height}` where:   * 16 <= `{width}` <= 1280 * 16 <= `{height}` <= 1280 * `{width}` * `{height}` <= 921,600  Typical values are:   * HD = `1280x720` * PAL = `1024x576` * VGA = `640x480` * CIF = `320x240`  Note that the `resolution` imposes an aspect ratio to the resulting composition. When the original video tracks are constrained by the aspect ratio, they are scaled to fit. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info.
	Resolution *string `json:"Resolution,omitempty"`
	//
	Format *string `json:"Format,omitempty"`
	// The URL we should call using the `status_callback_method` to send status information to your application on every composition event. If not provided, status callback events will not be dispatched.
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// The HTTP method we should use to call `status_callback`. Can be: `POST` or `GET` and the default is `POST`.
	StatusCallbackMethod *string `json:"StatusCallbackMethod,omitempty"`
	// Whether to clip the intervals where there is no active media in the composition. The default is `true`. Compositions with `trim` enabled are shorter when the Room is created and no Participant joins for a while as well as if all the Participants leave the room and join later, because those gaps will be removed. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info.
	Trim *bool `json:"Trim,omitempty"`
}

func (params *CreateCompositionParams) SetRoomSid(RoomSid string) *CreateCompositionParams {
	params.RoomSid = &RoomSid
	return params
}
func (params *CreateCompositionParams) SetVideoLayout(VideoLayout map[string]interface{}) *CreateCompositionParams {
	params.VideoLayout = &VideoLayout
	return params
}
func (params *CreateCompositionParams) SetAudioSources(AudioSources []string) *CreateCompositionParams {
	params.AudioSources = &AudioSources
	return params
}
func (params *CreateCompositionParams) SetAudioSourcesExcluded(AudioSourcesExcluded []string) *CreateCompositionParams {
	params.AudioSourcesExcluded = &AudioSourcesExcluded
	return params
}
func (params *CreateCompositionParams) SetResolution(Resolution string) *CreateCompositionParams {
	params.Resolution = &Resolution
	return params
}
func (params *CreateCompositionParams) SetFormat(Format string) *CreateCompositionParams {
	params.Format = &Format
	return params
}
func (params *CreateCompositionParams) SetStatusCallback(StatusCallback string) *CreateCompositionParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *CreateCompositionParams) SetStatusCallbackMethod(StatusCallbackMethod string) *CreateCompositionParams {
	params.StatusCallbackMethod = &StatusCallbackMethod
	return params
}
func (params *CreateCompositionParams) SetTrim(Trim bool) *CreateCompositionParams {
	params.Trim = &Trim
	return params
}

//
func (c *ApiService) CreateComposition(params *CreateCompositionParams) (*VideoV1Composition, error) {
	path := "/v1/Compositions"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.RoomSid != nil {
		data.Set("RoomSid", *params.RoomSid)
	}
	if params != nil && params.VideoLayout != nil {
		v, err := json.Marshal(params.VideoLayout)

		if err != nil {
			return nil, err
		}

		data.Set("VideoLayout", string(v))
	}
	if params != nil && params.AudioSources != nil {
		for _, item := range *params.AudioSources {
			data.Add("AudioSources", item)
		}
	}
	if params != nil && params.AudioSourcesExcluded != nil {
		for _, item := range *params.AudioSourcesExcluded {
			data.Add("AudioSourcesExcluded", item)
		}
	}
	if params != nil && params.Resolution != nil {
		data.Set("Resolution", *params.Resolution)
	}
	if params != nil && params.Format != nil {
		data.Set("Format", fmt.Sprint(*params.Format))
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.StatusCallbackMethod != nil {
		data.Set("StatusCallbackMethod", *params.StatusCallbackMethod)
	}
	if params != nil && params.Trim != nil {
		data.Set("Trim", fmt.Sprint(*params.Trim))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1Composition{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a Recording Composition resource identified by a Composition SID.
func (c *ApiService) DeleteComposition(Sid string) error {
	path := "/v1/Compositions/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Returns a single Composition resource identified by a Composition SID.
func (c *ApiService) FetchComposition(Sid string) (*VideoV1Composition, error) {
	path := "/v1/Compositions/{Sid}"
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

	ps := &VideoV1Composition{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListComposition'
type ListCompositionParams struct {
	// Read only Composition resources with this status. Can be: `enqueued`, `processing`, `completed`, `deleted`, or `failed`.
	Status *string `json:"Status,omitempty"`
	// Read only Composition resources created on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time with time zone.
	DateCreatedAfter *time.Time `json:"DateCreatedAfter,omitempty"`
	// Read only Composition resources created before this ISO 8601 date-time with time zone.
	DateCreatedBefore *time.Time `json:"DateCreatedBefore,omitempty"`
	// Read only Composition resources with this Room SID.
	RoomSid *string `json:"RoomSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListCompositionParams) SetStatus(Status string) *ListCompositionParams {
	params.Status = &Status
	return params
}
func (params *ListCompositionParams) SetDateCreatedAfter(DateCreatedAfter time.Time) *ListCompositionParams {
	params.DateCreatedAfter = &DateCreatedAfter
	return params
}
func (params *ListCompositionParams) SetDateCreatedBefore(DateCreatedBefore time.Time) *ListCompositionParams {
	params.DateCreatedBefore = &DateCreatedBefore
	return params
}
func (params *ListCompositionParams) SetRoomSid(RoomSid string) *ListCompositionParams {
	params.RoomSid = &RoomSid
	return params
}
func (params *ListCompositionParams) SetPageSize(PageSize int) *ListCompositionParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListCompositionParams) SetLimit(Limit int) *ListCompositionParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Composition records from the API. Request is executed immediately.
func (c *ApiService) PageComposition(params *ListCompositionParams, pageToken, pageNumber string) (*ListCompositionResponse, error) {
	path := "/v1/Compositions"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Status != nil {
		data.Set("Status", fmt.Sprint(*params.Status))
	}
	if params != nil && params.DateCreatedAfter != nil {
		data.Set("DateCreatedAfter", fmt.Sprint((*params.DateCreatedAfter).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedBefore != nil {
		data.Set("DateCreatedBefore", fmt.Sprint((*params.DateCreatedBefore).Format(time.RFC3339)))
	}
	if params != nil && params.RoomSid != nil {
		data.Set("RoomSid", *params.RoomSid)
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

	ps := &ListCompositionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Composition records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListComposition(params *ListCompositionParams) ([]VideoV1Composition, error) {
	response, errors := c.StreamComposition(params)

	records := make([]VideoV1Composition, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Composition records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamComposition(params *ListCompositionParams) (chan VideoV1Composition, chan error) {
	if params == nil {
		params = &ListCompositionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan VideoV1Composition, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageComposition(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamComposition(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamComposition(response *ListCompositionResponse, params *ListCompositionParams, recordChannel chan VideoV1Composition, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Compositions
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListCompositionResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListCompositionResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListCompositionResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCompositionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
