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

// Optional parameters for the method 'CreateQueue'
type CreateQueueParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// A descriptive string that you created to describe this resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The maximum number of calls allowed to be in the queue. The default is 100. The maximum is 5000.
	MaxSize *int `json:"MaxSize,omitempty"`
}

func (params *CreateQueueParams) SetPathAccountSid(PathAccountSid string) *CreateQueueParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateQueueParams) SetFriendlyName(FriendlyName string) *CreateQueueParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateQueueParams) SetMaxSize(MaxSize int) *CreateQueueParams {
	params.MaxSize = &MaxSize
	return params
}

// Create a queue
func (c *ApiService) CreateQueue(params *CreateQueueParams) (*ApiV2010AccountQueue, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Queues.json"
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
	if params != nil && params.MaxSize != nil {
		data.Set("MaxSize", fmt.Sprint(*params.MaxSize))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountQueue{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteQueue'
type DeleteQueueParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Queue resource to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteQueueParams) SetPathAccountSid(PathAccountSid string) *DeleteQueueParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Remove an empty queue
func (c *ApiService) DeleteQueue(Sid string, params *DeleteQueueParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/Queues/{Sid}.json"
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

// Optional parameters for the method 'FetchQueue'
type FetchQueueParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Queue resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchQueueParams) SetPathAccountSid(PathAccountSid string) *FetchQueueParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch an instance of a queue identified by the QueueSid
func (c *ApiService) FetchQueue(Sid string, params *FetchQueueParams) (*ApiV2010AccountQueue, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Queues/{Sid}.json"
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

	ps := &ApiV2010AccountQueue{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListQueue'
type ListQueueParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Queue resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListQueueParams) SetPathAccountSid(PathAccountSid string) *ListQueueParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListQueueParams) SetPageSize(PageSize int) *ListQueueParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of queues belonging to the account used to make the request
func (c *ApiService) ListQueue(params *ListQueueParams) (*ListQueueResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Queues.json"
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

	ps := &ListQueueResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of Queue records from the API. Request is executed immediately.
func (c *ApiService) QueuePage(params *ListQueueParams, pageToken string, pageNumber string, pageSize string) *client.Page {
	path := "/2010-04-01/Accounts/{AccountSid}/Queues.json"
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
	data.Set("PageSize", pageSize)

	response, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil
	}

	return client.NewPage(c.baseURL, response)
}

//Streams Queue records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) QueueStream(params *ListQueueParams, meta client.PaginationData) chan map[string]interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.QueuePage(params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.Stream(page, limits.Limit, limits.PageLimit)
}

//Lists Queue records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) QueueList(params *ListQueueParams, meta client.PaginationData) []interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.QueuePage(params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.List(page, limits.Limit, limits.PageLimit)
}

// Optional parameters for the method 'UpdateQueue'
type UpdateQueueParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Queue resource to update.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// A descriptive string that you created to describe this resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The maximum number of calls allowed to be in the queue. The default is 100. The maximum is 5000.
	MaxSize *int `json:"MaxSize,omitempty"`
}

func (params *UpdateQueueParams) SetPathAccountSid(PathAccountSid string) *UpdateQueueParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateQueueParams) SetFriendlyName(FriendlyName string) *UpdateQueueParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateQueueParams) SetMaxSize(MaxSize int) *UpdateQueueParams {
	params.MaxSize = &MaxSize
	return params
}

// Update the queue with the new parameters
func (c *ApiService) UpdateQueue(Sid string, params *UpdateQueueParams) (*ApiV2010AccountQueue, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Queues/{Sid}.json"
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
	if params != nil && params.MaxSize != nil {
		data.Set("MaxSize", fmt.Sprint(*params.MaxSize))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountQueue{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
