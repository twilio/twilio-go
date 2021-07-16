/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
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

// Optional parameters for the method 'CreateNewKey'
type CreateNewKeyParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Key resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateNewKeyParams) SetPathAccountSid(PathAccountSid string) *CreateNewKeyParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateNewKeyParams) SetFriendlyName(FriendlyName string) *CreateNewKeyParams {
	params.FriendlyName = &FriendlyName
	return params
}

func (c *ApiService) CreateNewKey(params *CreateNewKeyParams) (*ApiV2010AccountNewKey, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Keys.json"
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

	ps := &ApiV2010AccountNewKey{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteKey'
type DeleteKeyParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resources to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteKeyParams) SetPathAccountSid(PathAccountSid string) *DeleteKeyParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

func (c *ApiService) DeleteKey(Sid string, params *DeleteKeyParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/Keys/{Sid}.json"
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

// Optional parameters for the method 'FetchKey'
type FetchKeyParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchKeyParams) SetPathAccountSid(PathAccountSid string) *FetchKeyParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

func (c *ApiService) FetchKey(Sid string, params *FetchKeyParams) (*ApiV2010AccountKey, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Keys/{Sid}.json"
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

	ps := &ApiV2010AccountKey{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListKey'
type ListKeyParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListKeyParams) SetPathAccountSid(PathAccountSid string) *ListKeyParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListKeyParams) SetPageSize(PageSize int) *ListKeyParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a single page of Key records from the API. Request is executed immediately.
func (c *ApiService) PageKey(params *ListKeyParams, pageToken string, pageNumber string) (*ListKeyResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Keys.json"

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

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageToken != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListKeyResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Key records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListKey(params *ListKeyParams, limit int) ([]ApiV2010AccountKey, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageKey(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010AccountKey

	for response != nil {
		records = append(records, response.Keys...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, limit, c.getNextListKeyResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListKeyResponse)
	}

	return records, err
}

// Streams Key records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamKey(params *ListKeyParams, limit int) (chan ApiV2010AccountKey, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageKey(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010AccountKey, 1)

	go func() {
		for response != nil {
			for item := range response.Keys {
				channel <- response.Keys[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, limit, c.getNextListKeyResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListKeyResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListKeyResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListKeyResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateKey'
type UpdateKeyParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resources to update.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateKeyParams) SetPathAccountSid(PathAccountSid string) *UpdateKeyParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateKeyParams) SetFriendlyName(FriendlyName string) *UpdateKeyParams {
	params.FriendlyName = &FriendlyName
	return params
}

func (c *ApiService) UpdateKey(Sid string, params *UpdateKeyParams) (*ApiV2010AccountKey, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Keys/{Sid}.json"
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

	ps := &ApiV2010AccountKey{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
