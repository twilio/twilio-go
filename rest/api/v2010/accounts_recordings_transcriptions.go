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
)

// Optional parameters for the method 'DeleteRecordingTranscription'
type DeleteRecordingTranscriptionParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteRecordingTranscriptionParams) SetPathAccountSid(PathAccountSid string) *DeleteRecordingTranscriptionParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

func (c *ApiService) DeleteRecordingTranscription(RecordingSid string, Sid string, params *DeleteRecordingTranscriptionParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/Recordings/{RecordingSid}/Transcriptions/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"RecordingSid"+"}", RecordingSid, -1)
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

// Optional parameters for the method 'FetchRecordingTranscription'
type FetchRecordingTranscriptionParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchRecordingTranscriptionParams) SetPathAccountSid(PathAccountSid string) *FetchRecordingTranscriptionParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

func (c *ApiService) FetchRecordingTranscription(RecordingSid string, Sid string, params *FetchRecordingTranscriptionParams) (*ApiV2010AccountRecordingRecordingTranscription, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Recordings/{RecordingSid}/Transcriptions/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"RecordingSid"+"}", RecordingSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountRecordingRecordingTranscription{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRecordingTranscription'
type ListRecordingTranscriptionParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListRecordingTranscriptionParams) SetPathAccountSid(PathAccountSid string) *ListRecordingTranscriptionParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListRecordingTranscriptionParams) SetPageSize(PageSize int) *ListRecordingTranscriptionParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListRecordingTranscription(RecordingSid string, params *ListRecordingTranscriptionParams) (*ListRecordingTranscriptionResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Recordings/{RecordingSid}/Transcriptions.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"RecordingSid"+"}", RecordingSid, -1)

	data := url.Values{}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRecordingTranscriptionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
