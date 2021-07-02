/*
 * Twilio - Api
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

// Optional parameters for the method 'DeleteRecording'
type DeleteRecordingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteRecordingParams) SetPathAccountSid(PathAccountSid string) *DeleteRecordingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete a recording from your account
func (c *ApiService) DeleteRecording(Sid string, params *DeleteRecordingParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/Recordings/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
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

// Optional parameters for the method 'FetchRecording'
type FetchRecordingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchRecordingParams) SetPathAccountSid(PathAccountSid string) *FetchRecordingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch an instance of a recording
func (c *ApiService) FetchRecording(Sid string, params *FetchRecordingParams) (*ApiV2010AccountRecording, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Recordings/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountRecording{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRecording'
type ListRecordingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// Only include recordings that were created on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read recordings that were created on this date. You can also specify an inequality, such as `DateCreated<=YYYY-MM-DD`, to read recordings that were created on or before midnight of this date, and `DateCreated>=YYYY-MM-DD` to read recordings that were created on or after midnight of this date.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// Only include recordings that were created on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read recordings that were created on this date. You can also specify an inequality, such as `DateCreated<=YYYY-MM-DD`, to read recordings that were created on or before midnight of this date, and `DateCreated>=YYYY-MM-DD` to read recordings that were created on or after midnight of this date.
	DateCreatedBefore *time.Time `json:"DateCreated&lt;,omitempty"`
	// Only include recordings that were created on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read recordings that were created on this date. You can also specify an inequality, such as `DateCreated<=YYYY-MM-DD`, to read recordings that were created on or before midnight of this date, and `DateCreated>=YYYY-MM-DD` to read recordings that were created on or after midnight of this date.
	DateCreatedAfter *time.Time `json:"DateCreated&gt;,omitempty"`
	// The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resources to read.
	CallSid *string `json:"CallSid,omitempty"`
	// The Conference SID that identifies the conference associated with the recording to read.
	ConferenceSid *string `json:"ConferenceSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListRecordingParams) SetPathAccountSid(PathAccountSid string) *ListRecordingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListRecordingParams) SetDateCreated(DateCreated time.Time) *ListRecordingParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *ListRecordingParams) SetDateCreatedBefore(DateCreatedBefore time.Time) *ListRecordingParams {
	params.DateCreatedBefore = &DateCreatedBefore
	return params
}
func (params *ListRecordingParams) SetDateCreatedAfter(DateCreatedAfter time.Time) *ListRecordingParams {
	params.DateCreatedAfter = &DateCreatedAfter
	return params
}
func (params *ListRecordingParams) SetCallSid(CallSid string) *ListRecordingParams {
	params.CallSid = &CallSid
	return params
}
func (params *ListRecordingParams) SetConferenceSid(ConferenceSid string) *ListRecordingParams {
	params.ConferenceSid = &ConferenceSid
	return params
}
func (params *ListRecordingParams) SetPageSize(PageSize int) *ListRecordingParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of recordings belonging to the account used to make the request
func (c *ApiService) ListRecording(params *ListRecordingParams) (*ListRecordingResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Recordings.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint((*params.DateCreated).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedBefore != nil {
		data.Set("DateCreated<", fmt.Sprint((*params.DateCreatedBefore).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedAfter != nil {
		data.Set("DateCreated>", fmt.Sprint((*params.DateCreatedAfter).Format(time.RFC3339)))
	}
	if params != nil && params.CallSid != nil {
		data.Set("CallSid", *params.CallSid)
	}
	if params != nil && params.ConferenceSid != nil {
		data.Set("ConferenceSid", *params.ConferenceSid)
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
func (c *ApiService) RecordingPage(params *ListRecordingParams, pageToken string, pageNumber string) (*client.Page, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Recordings.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint((*params.DateCreated).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedBefore != nil {
		data.Set("DateCreated<", fmt.Sprint((*params.DateCreatedBefore).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedAfter != nil {
		data.Set("DateCreated>", fmt.Sprint((*params.DateCreatedAfter).Format(time.RFC3339)))
	}
	if params != nil && params.CallSid != nil {
		data.Set("CallSid", *params.CallSid)
	}
	if params != nil && params.ConferenceSid != nil {
		data.Set("ConferenceSid", *params.ConferenceSid)
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

//Streams Recording records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) RecordingStream(params *ListRecordingParams, limit int) (chan map[string]interface{}, error) {
	if params.PageSize == nil {
		params.SetPageSize(0)
	}
	params.SetPageSize(c.requestHandler.ReadLimits(*params.PageSize, limit))
	page, err := c.RecordingPage(params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.Stream(page, limit, 0), nil
}

//Lists Recording records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) RecordingList(params *ListRecordingParams, limit int) ([]interface{}, error) {
	if params.PageSize == nil {
		params.SetPageSize(0)
	}
	params.SetPageSize(c.requestHandler.ReadLimits(*params.PageSize, limit))
	page, err := c.RecordingPage(params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.List(page, limit, 0), nil
}
