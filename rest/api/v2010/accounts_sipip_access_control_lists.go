/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateSipIpAccessControlList'
type CreateSipIpAccessControlListParams struct {
	// The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// A human readable descriptive text that describes the IpAccessControlList, up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateSipIpAccessControlListParams) SetPathAccountSid(PathAccountSid string) *CreateSipIpAccessControlListParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateSipIpAccessControlListParams) SetFriendlyName(FriendlyName string) *CreateSipIpAccessControlListParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Create a new IpAccessControlList resource
func (c *ApiService) CreateSipIpAccessControlList(params *CreateSipIpAccessControlListParams) (*ApiV2010SipIpAccessControlList, error) {
	return c.CreateSipIpAccessControlListWithCtx(context.TODO(), params)
}

// Create a new IpAccessControlList resource
func (c *ApiService) CreateSipIpAccessControlListWithCtx(ctx context.Context, params *CreateSipIpAccessControlListParams) (*ApiV2010SipIpAccessControlList, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists.json"
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

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010SipIpAccessControlList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteSipIpAccessControlList'
type DeleteSipIpAccessControlListParams struct {
	// The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteSipIpAccessControlListParams) SetPathAccountSid(PathAccountSid string) *DeleteSipIpAccessControlListParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete an IpAccessControlList from the requested account
func (c *ApiService) DeleteSipIpAccessControlList(Sid string, params *DeleteSipIpAccessControlListParams) error {
	return c.DeleteSipIpAccessControlListWithCtx(context.TODO(), Sid, params)
}

// Delete an IpAccessControlList from the requested account
func (c *ApiService) DeleteSipIpAccessControlListWithCtx(ctx context.Context, Sid string, params *DeleteSipIpAccessControlListParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Optional parameters for the method 'FetchSipIpAccessControlList'
type FetchSipIpAccessControlListParams struct {
	// The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchSipIpAccessControlListParams) SetPathAccountSid(PathAccountSid string) *FetchSipIpAccessControlListParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch a specific instance of an IpAccessControlList
func (c *ApiService) FetchSipIpAccessControlList(Sid string, params *FetchSipIpAccessControlListParams) (*ApiV2010SipIpAccessControlList, error) {
	return c.FetchSipIpAccessControlListWithCtx(context.TODO(), Sid, params)
}

// Fetch a specific instance of an IpAccessControlList
func (c *ApiService) FetchSipIpAccessControlListWithCtx(ctx context.Context, Sid string, params *FetchSipIpAccessControlListParams) (*ApiV2010SipIpAccessControlList, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010SipIpAccessControlList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSipIpAccessControlList'
type ListSipIpAccessControlListParams struct {
	// The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSipIpAccessControlListParams) SetPathAccountSid(PathAccountSid string) *ListSipIpAccessControlListParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListSipIpAccessControlListParams) SetPageSize(PageSize int) *ListSipIpAccessControlListParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSipIpAccessControlListParams) SetLimit(Limit int) *ListSipIpAccessControlListParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SipIpAccessControlList records from the API. Request is executed immediately.
func (c *ApiService) PageSipIpAccessControlList(params *ListSipIpAccessControlListParams, pageToken, pageNumber string) (*ListSipIpAccessControlListResponse, error) {
	return c.PageSipIpAccessControlListWithCtx(context.TODO(), params, pageToken, pageNumber)
}

// Retrieve a single page of SipIpAccessControlList records from the API. Request is executed immediately.
func (c *ApiService) PageSipIpAccessControlListWithCtx(ctx context.Context, params *ListSipIpAccessControlListParams, pageToken, pageNumber string) (*ListSipIpAccessControlListResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists.json"

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
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSipIpAccessControlListResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SipIpAccessControlList records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSipIpAccessControlList(params *ListSipIpAccessControlListParams) ([]ApiV2010SipIpAccessControlList, error) {
	return c.ListSipIpAccessControlListWithCtx(context.TODO(), params)
}

// Lists SipIpAccessControlList records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSipIpAccessControlListWithCtx(ctx context.Context, params *ListSipIpAccessControlListParams) ([]ApiV2010SipIpAccessControlList, error) {
	response, errors := c.StreamSipIpAccessControlListWithCtx(ctx, params)

	records := make([]ApiV2010SipIpAccessControlList, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams SipIpAccessControlList records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSipIpAccessControlList(params *ListSipIpAccessControlListParams) (chan ApiV2010SipIpAccessControlList, chan error) {
	return c.StreamSipIpAccessControlListWithCtx(context.TODO(), params)
}

// Streams SipIpAccessControlList records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSipIpAccessControlListWithCtx(ctx context.Context, params *ListSipIpAccessControlListParams) (chan ApiV2010SipIpAccessControlList, chan error) {
	if params == nil {
		params = &ListSipIpAccessControlListParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ApiV2010SipIpAccessControlList, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageSipIpAccessControlListWithCtx(ctx, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamSipIpAccessControlList(ctx, response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamSipIpAccessControlList(ctx context.Context, response *ListSipIpAccessControlListResponse, params *ListSipIpAccessControlListParams, recordChannel chan ApiV2010SipIpAccessControlList, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.IpAccessControlLists
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNextWithCtx(ctx, c.baseURL, response, c.getNextListSipIpAccessControlListResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListSipIpAccessControlListResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListSipIpAccessControlListResponse(ctx context.Context, nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(ctx, nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSipIpAccessControlListResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSipIpAccessControlList'
type UpdateSipIpAccessControlListParams struct {
	// The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// A human readable descriptive text, up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateSipIpAccessControlListParams) SetPathAccountSid(PathAccountSid string) *UpdateSipIpAccessControlListParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateSipIpAccessControlListParams) SetFriendlyName(FriendlyName string) *UpdateSipIpAccessControlListParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Rename an IpAccessControlList
func (c *ApiService) UpdateSipIpAccessControlList(Sid string, params *UpdateSipIpAccessControlListParams) (*ApiV2010SipIpAccessControlList, error) {
	return c.UpdateSipIpAccessControlListWithCtx(context.TODO(), Sid, params)
}

// Rename an IpAccessControlList
func (c *ApiService) UpdateSipIpAccessControlListWithCtx(ctx context.Context, Sid string, params *UpdateSipIpAccessControlListParams) (*ApiV2010SipIpAccessControlList, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{Sid}.json"
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

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010SipIpAccessControlList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
