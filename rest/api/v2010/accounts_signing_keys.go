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

// Optional parameters for the method 'CreateNewSigningKey'
type CreateNewSigningKeyParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Key resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateNewSigningKeyParams) SetPathAccountSid(PathAccountSid string) *CreateNewSigningKeyParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateNewSigningKeyParams) SetFriendlyName(FriendlyName string) *CreateNewSigningKeyParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Create a new Signing Key for the account making the request.
func (c *ApiService) CreateNewSigningKey(params *CreateNewSigningKeyParams) (*ApiV2010AccountNewSigningKey, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SigningKeys.json"
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

	ps := &ApiV2010AccountNewSigningKey{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteSigningKey'
type DeleteSigningKeyParams struct {
	//
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteSigningKeyParams) SetPathAccountSid(PathAccountSid string) *DeleteSigningKeyParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

func (c *ApiService) DeleteSigningKey(Sid string, params *DeleteSigningKeyParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/SigningKeys/{Sid}.json"
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

// Optional parameters for the method 'FetchSigningKey'
type FetchSigningKeyParams struct {
	//
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchSigningKeyParams) SetPathAccountSid(PathAccountSid string) *FetchSigningKeyParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

func (c *ApiService) FetchSigningKey(Sid string, params *FetchSigningKeyParams) (*ApiV2010AccountSigningKey, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SigningKeys/{Sid}.json"
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

	ps := &ApiV2010AccountSigningKey{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSigningKey'
type ListSigningKeyParams struct {
	//
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListSigningKeyParams) SetPathAccountSid(PathAccountSid string) *ListSigningKeyParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListSigningKeyParams) SetPageSize(PageSize int) *ListSigningKeyParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListSigningKey(params *ListSigningKeyParams) (*ListSigningKeyResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SigningKeys.json"
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

	ps := &ListSigningKeyResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of  records from the API. Request is executed immediately.
func (c *ApiService) AccountsSigningKeysPage(params *ListSigningKeyParams, pageToken string, pageNumber string) (*ListSigningKeyResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SigningKeys.json"
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

	ps := &ListSigningKeyResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Lists AccountsSigningKeys records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) AccountsSigningKeysList(params *ListSigningKeyParams, limit int) ([]ListSigningKeyResponse, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	response, err := c.ListSigningKey(params)
	if err != nil {
		return nil, err
	}

	page := client.NewPage(c.baseURL, response)

	resp := c.requestHandler.List(page, limit, 0)
	ret := make([]ListSigningKeyResponse, len(resp))

	for i := range resp {
		jsonStr, _ := json.Marshal(resp[i])
		ps := ListSigningKeyResponse{}
		if err := json.Unmarshal(jsonStr, &ps); err != nil {
			return ret, err
		}

		ret[i] = ps
	}

	return ret, nil
}

//Streams AccountsSigningKeys records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) AccountsSigningKeysStream(params *ListSigningKeyParams, limit int) (chan interface{}, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	response, err := c.ListSigningKey(params)
	if err != nil {
		return nil, err
	}

	page := client.NewPage(c.baseURL, response)

	ps := ListSigningKeyResponse{}
	return c.requestHandler.Stream(page, limit, 0, ps), nil
}

// Optional parameters for the method 'UpdateSigningKey'
type UpdateSigningKeyParams struct {
	//
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	//
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateSigningKeyParams) SetPathAccountSid(PathAccountSid string) *UpdateSigningKeyParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateSigningKeyParams) SetFriendlyName(FriendlyName string) *UpdateSigningKeyParams {
	params.FriendlyName = &FriendlyName
	return params
}

func (c *ApiService) UpdateSigningKey(Sid string, params *UpdateSigningKeyParams) (*ApiV2010AccountSigningKey, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SigningKeys/{Sid}.json"
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

	ps := &ApiV2010AccountSigningKey{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
