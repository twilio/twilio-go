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

// Optional parameters for the method 'FetchAuthorizedConnectApp'
type FetchAuthorizedConnectAppParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the AuthorizedConnectApp resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchAuthorizedConnectAppParams) SetPathAccountSid(PathAccountSid string) *FetchAuthorizedConnectAppParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch an instance of an authorized-connect-app
func (c *ApiService) FetchAuthorizedConnectApp(ConnectAppSid string, params *FetchAuthorizedConnectAppParams) (*ApiV2010AccountAuthorizedConnectApp, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/AuthorizedConnectApps/{ConnectAppSid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ConnectAppSid"+"}", ConnectAppSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountAuthorizedConnectApp{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListAuthorizedConnectApp'
type ListAuthorizedConnectAppParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the AuthorizedConnectApp resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListAuthorizedConnectAppParams) SetPathAccountSid(PathAccountSid string) *ListAuthorizedConnectAppParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListAuthorizedConnectAppParams) SetPageSize(PageSize int) *ListAuthorizedConnectAppParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of authorized-connect-apps belonging to the account used to make the request
func (c *ApiService) ListAuthorizedConnectApp(params *ListAuthorizedConnectAppParams) (*ListAuthorizedConnectAppResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/AuthorizedConnectApps.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

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

	ps := &ListAuthorizedConnectAppResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of AuthorizedConnectApp records from the API. Request is executed immediately.
func (c *ApiService) AuthorizedConnectAppPage(params *ListAuthorizedConnectAppParams, pageToken string, pageNumber string) (*client.Page, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/AuthorizedConnectApps.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

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

//Streams AuthorizedConnectApp records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) AuthorizedConnectAppStream(params *ListAuthorizedConnectAppParams, limit int) (chan map[string]interface{}, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	page, err := c.AuthorizedConnectAppPage(params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.Stream(page, limit, 0), nil
}

//Lists AuthorizedConnectApp records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) AuthorizedConnectAppList(params *ListAuthorizedConnectAppParams, limit int) ([]interface{}, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	page, err := c.AuthorizedConnectAppPage(params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.List(page, limit, 0), nil
}
