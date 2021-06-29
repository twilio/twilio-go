/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
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
)

// Optional parameters for the method 'DeleteMedia'
type DeleteMediaParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Media resource(s) to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteMediaParams) SetPathAccountSid(PathAccountSid string) *DeleteMediaParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete media from your account. Once delete, you will no longer be billed
func (c *ApiService) DeleteMedia(MessageSid string, Sid string, params *DeleteMediaParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/Messages/{MessageSid}/Media/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"MessageSid"+"}", MessageSid, -1)
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

// Optional parameters for the method 'FetchMedia'
type FetchMediaParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Media resource(s) to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchMediaParams) SetPathAccountSid(PathAccountSid string) *FetchMediaParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch a single media instance belonging to the account used to make the request
func (c *ApiService) FetchMedia(MessageSid string, Sid string, params *FetchMediaParams) (*ApiV2010AccountMessageMedia, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Messages/{MessageSid}/Media/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"MessageSid"+"}", MessageSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountMessageMedia{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListMedia'
type ListMediaParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Media resource(s) to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// Only include media that was created on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read media that was created on this date. You can also specify an inequality, such as `StartTime<=YYYY-MM-DD`, to read media that was created on or before midnight of this date, and `StartTime>=YYYY-MM-DD` to read media that was created on or after midnight of this date.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// Only include media that was created on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read media that was created on this date. You can also specify an inequality, such as `StartTime<=YYYY-MM-DD`, to read media that was created on or before midnight of this date, and `StartTime>=YYYY-MM-DD` to read media that was created on or after midnight of this date.
	DateCreatedBefore *time.Time `json:"DateCreated&lt;,omitempty"`
	// Only include media that was created on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read media that was created on this date. You can also specify an inequality, such as `StartTime<=YYYY-MM-DD`, to read media that was created on or before midnight of this date, and `StartTime>=YYYY-MM-DD` to read media that was created on or after midnight of this date.
	DateCreatedAfter *time.Time `json:"DateCreated&gt;,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListMediaParams) SetPathAccountSid(PathAccountSid string) *ListMediaParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListMediaParams) SetDateCreated(DateCreated time.Time) *ListMediaParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *ListMediaParams) SetDateCreatedBefore(DateCreatedBefore time.Time) *ListMediaParams {
	params.DateCreatedBefore = &DateCreatedBefore
	return params
}
func (params *ListMediaParams) SetDateCreatedAfter(DateCreatedAfter time.Time) *ListMediaParams {
	params.DateCreatedAfter = &DateCreatedAfter
	return params
}
func (params *ListMediaParams) SetPageSize(PageSize int) *ListMediaParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of Media resources belonging to the account used to make the request
func (c *ApiService) ListMedia(MessageSid string, params *ListMediaParams) (*ListMediaResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Messages/{MessageSid}/Media.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"MessageSid"+"}", MessageSid, -1)

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
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMediaResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}