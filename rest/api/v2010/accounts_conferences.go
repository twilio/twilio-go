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

// Optional parameters for the method 'FetchConference'
type FetchConferenceParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference resource(s) to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchConferenceParams) SetPathAccountSid(PathAccountSid string) *FetchConferenceParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch an instance of a conference
func (c *ApiService) FetchConference(Sid string, params *FetchConferenceParams) (*ApiV2010AccountConference, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Conferences/{Sid}.json"
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

	ps := &ApiV2010AccountConference{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListConference'
type ListConferenceParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference resource(s) to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. To read conferences that started on or before midnight on a date, use `<=YYYY-MM-DD`, and to specify  conferences that started on or after midnight on a date, use `>=YYYY-MM-DD`.
	DateCreated *string `json:"DateCreated,omitempty"`
	// The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. To read conferences that started on or before midnight on a date, use `<=YYYY-MM-DD`, and to specify  conferences that started on or after midnight on a date, use `>=YYYY-MM-DD`.
	DateCreatedBefore *string `json:"DateCreated&lt;,omitempty"`
	// The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. To read conferences that started on or before midnight on a date, use `<=YYYY-MM-DD`, and to specify  conferences that started on or after midnight on a date, use `>=YYYY-MM-DD`.
	DateCreatedAfter *string `json:"DateCreated&gt;,omitempty"`
	// The `date_updated` value, specified as `YYYY-MM-DD`, of the resources to read. To read conferences that were last updated on or before midnight on a date, use `<=YYYY-MM-DD`, and to specify conferences that were last updated on or after midnight on a given date, use  `>=YYYY-MM-DD`.
	DateUpdated *string `json:"DateUpdated,omitempty"`
	// The `date_updated` value, specified as `YYYY-MM-DD`, of the resources to read. To read conferences that were last updated on or before midnight on a date, use `<=YYYY-MM-DD`, and to specify conferences that were last updated on or after midnight on a given date, use  `>=YYYY-MM-DD`.
	DateUpdatedBefore *string `json:"DateUpdated&lt;,omitempty"`
	// The `date_updated` value, specified as `YYYY-MM-DD`, of the resources to read. To read conferences that were last updated on or before midnight on a date, use `<=YYYY-MM-DD`, and to specify conferences that were last updated on or after midnight on a given date, use  `>=YYYY-MM-DD`.
	DateUpdatedAfter *string `json:"DateUpdated&gt;,omitempty"`
	// The string that identifies the Conference resources to read.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The status of the resources to read. Can be: `init`, `in-progress`, or `completed`.
	Status *string `json:"Status,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListConferenceParams) SetPathAccountSid(PathAccountSid string) *ListConferenceParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListConferenceParams) SetDateCreated(DateCreated string) *ListConferenceParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *ListConferenceParams) SetDateCreatedBefore(DateCreatedBefore string) *ListConferenceParams {
	params.DateCreatedBefore = &DateCreatedBefore
	return params
}
func (params *ListConferenceParams) SetDateCreatedAfter(DateCreatedAfter string) *ListConferenceParams {
	params.DateCreatedAfter = &DateCreatedAfter
	return params
}
func (params *ListConferenceParams) SetDateUpdated(DateUpdated string) *ListConferenceParams {
	params.DateUpdated = &DateUpdated
	return params
}
func (params *ListConferenceParams) SetDateUpdatedBefore(DateUpdatedBefore string) *ListConferenceParams {
	params.DateUpdatedBefore = &DateUpdatedBefore
	return params
}
func (params *ListConferenceParams) SetDateUpdatedAfter(DateUpdatedAfter string) *ListConferenceParams {
	params.DateUpdatedAfter = &DateUpdatedAfter
	return params
}
func (params *ListConferenceParams) SetFriendlyName(FriendlyName string) *ListConferenceParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListConferenceParams) SetStatus(Status string) *ListConferenceParams {
	params.Status = &Status
	return params
}
func (params *ListConferenceParams) SetPageSize(PageSize int) *ListConferenceParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of conferences belonging to the account used to make the request
func (c *ApiService) ListConference(params *ListConferenceParams) (*ListConferenceResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Conferences.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint(*params.DateCreated))
	}
	if params != nil && params.DateCreatedBefore != nil {
		data.Set("DateCreated<", fmt.Sprint(*params.DateCreatedBefore))
	}
	if params != nil && params.DateCreatedAfter != nil {
		data.Set("DateCreated>", fmt.Sprint(*params.DateCreatedAfter))
	}
	if params != nil && params.DateUpdated != nil {
		data.Set("DateUpdated", fmt.Sprint(*params.DateUpdated))
	}
	if params != nil && params.DateUpdatedBefore != nil {
		data.Set("DateUpdated<", fmt.Sprint(*params.DateUpdatedBefore))
	}
	if params != nil && params.DateUpdatedAfter != nil {
		data.Set("DateUpdated>", fmt.Sprint(*params.DateUpdatedAfter))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListConferenceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of Conference records from the API. Request is executed immediately.
func (c *ApiService) ConferencePage(params *ListConferenceParams, pageToken string, pageNumber string) (*client.Page, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Conferences.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint(*params.DateCreated))
	}
	if params != nil && params.DateCreatedBefore != nil {
		data.Set("DateCreated<", fmt.Sprint(*params.DateCreatedBefore))
	}
	if params != nil && params.DateCreatedAfter != nil {
		data.Set("DateCreated>", fmt.Sprint(*params.DateCreatedAfter))
	}
	if params != nil && params.DateUpdated != nil {
		data.Set("DateUpdated", fmt.Sprint(*params.DateUpdated))
	}
	if params != nil && params.DateUpdatedBefore != nil {
		data.Set("DateUpdated<", fmt.Sprint(*params.DateUpdatedBefore))
	}
	if params != nil && params.DateUpdatedAfter != nil {
		data.Set("DateUpdated>", fmt.Sprint(*params.DateUpdatedAfter))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
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

//Streams Conference records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) ConferenceStream(params *ListConferenceParams, limit int) (chan map[string]interface{}, error) {
	if params.PageSize == nil {
		params.SetPageSize(0)
	}
	params.SetPageSize(c.requestHandler.ReadLimits(*params.PageSize, limit))
	page, err := c.ConferencePage(params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.Stream(page, limit, 0), nil
}

//Lists Conference records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) ConferenceList(params *ListConferenceParams, limit int) ([]interface{}, error) {
	if params.PageSize == nil {
		params.SetPageSize(0)
	}
	params.SetPageSize(c.requestHandler.ReadLimits(*params.PageSize, limit))
	page, err := c.ConferencePage(params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.List(page, limit, 0), nil
}

// Optional parameters for the method 'UpdateConference'
type UpdateConferenceParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference resource(s) to update.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The HTTP method used to call `announce_url`. Can be: `GET` or `POST` and the default is `POST`
	AnnounceMethod *string `json:"AnnounceMethod,omitempty"`
	// The URL we should call to announce something into the conference. The URL can return an MP3, a WAV, or a TwiML document with `<Play>` or `<Say>`.
	AnnounceUrl *string `json:"AnnounceUrl,omitempty"`
	// The new status of the resource. Can be:  Can be: `init`, `in-progress`, or `completed`. Specifying `completed` will end the conference and hang up all participants
	Status *string `json:"Status,omitempty"`
}

func (params *UpdateConferenceParams) SetPathAccountSid(PathAccountSid string) *UpdateConferenceParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateConferenceParams) SetAnnounceMethod(AnnounceMethod string) *UpdateConferenceParams {
	params.AnnounceMethod = &AnnounceMethod
	return params
}
func (params *UpdateConferenceParams) SetAnnounceUrl(AnnounceUrl string) *UpdateConferenceParams {
	params.AnnounceUrl = &AnnounceUrl
	return params
}
func (params *UpdateConferenceParams) SetStatus(Status string) *UpdateConferenceParams {
	params.Status = &Status
	return params
}

func (c *ApiService) UpdateConference(Sid string, params *UpdateConferenceParams) (*ApiV2010AccountConference, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Conferences/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AnnounceMethod != nil {
		data.Set("AnnounceMethod", *params.AnnounceMethod)
	}
	if params != nil && params.AnnounceUrl != nil {
		data.Set("AnnounceUrl", *params.AnnounceUrl)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountConference{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
