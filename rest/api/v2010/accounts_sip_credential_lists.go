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

// Optional parameters for the method 'CreateSipCredentialList'
type CreateSipCredentialListParams struct {
	// The unique id of the Account that is responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// A human readable descriptive text that describes the CredentialList, up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateSipCredentialListParams) SetPathAccountSid(PathAccountSid string) *CreateSipCredentialListParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateSipCredentialListParams) SetFriendlyName(FriendlyName string) *CreateSipCredentialListParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Create a Credential List
func (c *ApiService) CreateSipCredentialList(params *CreateSipCredentialListParams) (*ApiV2010AccountSipSipCredentialList, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountSipSipCredentialList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteSipCredentialList'
type DeleteSipCredentialListParams struct {
	// The unique id of the Account that is responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteSipCredentialListParams) SetPathAccountSid(PathAccountSid string) *DeleteSipCredentialListParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete a Credential List
func (c *ApiService) DeleteSipCredentialList(Sid string, params *DeleteSipCredentialListParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{Sid}.json"
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

// Optional parameters for the method 'FetchSipCredentialList'
type FetchSipCredentialListParams struct {
	// The unique id of the Account that is responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchSipCredentialListParams) SetPathAccountSid(PathAccountSid string) *FetchSipCredentialListParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Get a Credential List
func (c *ApiService) FetchSipCredentialList(Sid string, params *FetchSipCredentialListParams) (*ApiV2010AccountSipSipCredentialList, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{Sid}.json"
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

	ps := &ApiV2010AccountSipSipCredentialList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSipCredentialList'
type ListSipCredentialListParams struct {
	// The unique id of the Account that is responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListSipCredentialListParams) SetPathAccountSid(PathAccountSid string) *ListSipCredentialListParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListSipCredentialListParams) SetPageSize(PageSize int) *ListSipCredentialListParams {
	params.PageSize = &PageSize
	return params
}

// Get All Credential Lists
func (c *ApiService) ListSipCredentialList(params *ListSipCredentialListParams) (*ListSipCredentialListResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists.json"
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

	ps := &ListSipCredentialListResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of  records from the API. Request is executed immediately.
func (c *ApiService) AccountsSIPCredentialListsPage(params *ListSipCredentialListParams, pageToken string, pageNumber string) (*ListSipCredentialListResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists.json"
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

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSipCredentialListResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Lists AccountsSIPCredentialLists records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) AccountsSIPCredentialListsList(params *ListSipCredentialListParams, limit int) ([]ListSipCredentialListResponse, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	response, err := c.ListSipCredentialList(params)
	if err != nil {
		return nil, err
	}

	page := client.NewPage(c.baseURL, response)

	resp := c.requestHandler.List(page, limit, 0)
	ret := make([]ListSipCredentialListResponse, len(resp))

	for i := range resp {
		jsonStr, _ := json.Marshal(resp[i])
		ps := ListSipCredentialListResponse{}
		if err := json.Unmarshal(jsonStr, &ps); err != nil {
			return ret, err
		}

		ret[i] = ps
	}

	return ret, nil
}

//Streams AccountsSIPCredentialLists records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) AccountsSIPCredentialListsStream(params *ListSipCredentialListParams, limit int) (chan interface{}, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	response, err := c.ListSipCredentialList(params)
	if err != nil {
		return nil, err
	}

	page := client.NewPage(c.baseURL, response)

	ps := ListSipCredentialListResponse{}
	return c.requestHandler.Stream(page, limit, 0, ps), nil
}

// Optional parameters for the method 'UpdateSipCredentialList'
type UpdateSipCredentialListParams struct {
	// The unique id of the Account that is responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// A human readable descriptive text for a CredentialList, up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateSipCredentialListParams) SetPathAccountSid(PathAccountSid string) *UpdateSipCredentialListParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateSipCredentialListParams) SetFriendlyName(FriendlyName string) *UpdateSipCredentialListParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Update a Credential List
func (c *ApiService) UpdateSipCredentialList(Sid string, params *UpdateSipCredentialListParams) (*ApiV2010AccountSipSipCredentialList, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountSipSipCredentialList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
