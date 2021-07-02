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

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'DeleteRecordingAddOnResultPayload'
type DeleteRecordingAddOnResultPayloadParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resources to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteRecordingAddOnResultPayloadParams) SetPathAccountSid(PathAccountSid string) *DeleteRecordingAddOnResultPayloadParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete a payload from the result along with all associated Data
func (c *ApiService) DeleteRecordingAddOnResultPayload(ReferenceSid string, AddOnResultSid string, Sid string, params *DeleteRecordingAddOnResultPayloadParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{AddOnResultSid}/Payloads/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ReferenceSid"+"}", ReferenceSid, -1)
	path = strings.Replace(path, "{"+"AddOnResultSid"+"}", AddOnResultSid, -1)
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

// Optional parameters for the method 'FetchRecordingAddOnResultPayload'
type FetchRecordingAddOnResultPayloadParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchRecordingAddOnResultPayloadParams) SetPathAccountSid(PathAccountSid string) *FetchRecordingAddOnResultPayloadParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch an instance of a result payload
func (c *ApiService) FetchRecordingAddOnResultPayload(ReferenceSid string, AddOnResultSid string, Sid string, params *FetchRecordingAddOnResultPayloadParams) (*ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{AddOnResultSid}/Payloads/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ReferenceSid"+"}", ReferenceSid, -1)
	path = strings.Replace(path, "{"+"AddOnResultSid"+"}", AddOnResultSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRecordingAddOnResultPayload'
type ListRecordingAddOnResultPayloadParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListRecordingAddOnResultPayloadParams) SetPathAccountSid(PathAccountSid string) *ListRecordingAddOnResultPayloadParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListRecordingAddOnResultPayloadParams) SetPageSize(PageSize int) *ListRecordingAddOnResultPayloadParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of payloads belonging to the AddOnResult
func (c *ApiService) ListRecordingAddOnResultPayload(ReferenceSid string, AddOnResultSid string, params *ListRecordingAddOnResultPayloadParams) (*ListRecordingAddOnResultPayloadResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{AddOnResultSid}/Payloads.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ReferenceSid"+"}", ReferenceSid, -1)
	path = strings.Replace(path, "{"+"AddOnResultSid"+"}", AddOnResultSid, -1)

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

	ps := &ListRecordingAddOnResultPayloadResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of RecordingAddOnResultPayload records from the API. Request is executed immediately.
func (c *ApiService) RecordingAddOnResultPayloadPage(ReferenceSid string, AddOnResultSid string, params *ListRecordingAddOnResultPayloadParams, pageToken string, pageNumber string) (*client.Page, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{AddOnResultSid}/Payloads.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ReferenceSid"+"}", ReferenceSid, -1)
	path = strings.Replace(path, "{"+"AddOnResultSid"+"}", AddOnResultSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

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

//Streams RecordingAddOnResultPayload records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) RecordingAddOnResultPayloadStream(ReferenceSid string, AddOnResultSid string, params *ListRecordingAddOnResultPayloadParams, limit int) (chan map[string]interface{}, error) {
	if params.PageSize == nil {
		params.SetPageSize(0)
	}
	params.SetPageSize(c.requestHandler.ReadLimits(*params.PageSize, limit))
	page, err := c.RecordingAddOnResultPayloadPage(ReferenceSid, AddOnResultSid, params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.Stream(page, limit, 0), nil
}

//Lists RecordingAddOnResultPayload records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) RecordingAddOnResultPayloadList(ReferenceSid string, AddOnResultSid string, params *ListRecordingAddOnResultPayloadParams, limit int) ([]interface{}, error) {
	if params.PageSize == nil {
		params.SetPageSize(0)
	}
	params.SetPageSize(c.requestHandler.ReadLimits(*params.PageSize, limit))
	page, err := c.RecordingAddOnResultPayloadPage(ReferenceSid, AddOnResultSid, params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.List(page, limit, 0), nil
}
