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

// Optional parameters for the method 'CreateIncomingPhoneNumberAssignedAddOn'
type CreateIncomingPhoneNumberAssignedAddOnParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The SID that identifies the Add-on installation.
	InstalledAddOnSid *string `json:"InstalledAddOnSid,omitempty"`
}

func (params *CreateIncomingPhoneNumberAssignedAddOnParams) SetPathAccountSid(PathAccountSid string) *CreateIncomingPhoneNumberAssignedAddOnParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateIncomingPhoneNumberAssignedAddOnParams) SetInstalledAddOnSid(InstalledAddOnSid string) *CreateIncomingPhoneNumberAssignedAddOnParams {
	params.InstalledAddOnSid = &InstalledAddOnSid
	return params
}

// Assign an Add-on installation to the Number specified.
func (c *ApiService) CreateIncomingPhoneNumberAssignedAddOn(ResourceSid string, params *CreateIncomingPhoneNumberAssignedAddOnParams) (*ApiV2010IncomingPhoneNumberAssignedAddOn, error) {
	return c.CreateIncomingPhoneNumberAssignedAddOnWithCtx(context.TODO(), ResourceSid, params)
}

// Assign an Add-on installation to the Number specified.
func (c *ApiService) CreateIncomingPhoneNumberAssignedAddOnWithCtx(ctx context.Context, ResourceSid string, params *CreateIncomingPhoneNumberAssignedAddOnParams) (*ApiV2010IncomingPhoneNumberAssignedAddOn, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ResourceSid"+"}", ResourceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.InstalledAddOnSid != nil {
		data.Set("InstalledAddOnSid", *params.InstalledAddOnSid)
	}

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010IncomingPhoneNumberAssignedAddOn{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteIncomingPhoneNumberAssignedAddOn'
type DeleteIncomingPhoneNumberAssignedAddOnParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteIncomingPhoneNumberAssignedAddOnParams) SetPathAccountSid(PathAccountSid string) *DeleteIncomingPhoneNumberAssignedAddOnParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Remove the assignment of an Add-on installation from the Number specified.
func (c *ApiService) DeleteIncomingPhoneNumberAssignedAddOn(ResourceSid string, Sid string, params *DeleteIncomingPhoneNumberAssignedAddOnParams) error {
	return c.DeleteIncomingPhoneNumberAssignedAddOnWithCtx(context.TODO(), ResourceSid, Sid, params)
}

// Remove the assignment of an Add-on installation from the Number specified.
func (c *ApiService) DeleteIncomingPhoneNumberAssignedAddOnWithCtx(ctx context.Context, ResourceSid string, Sid string, params *DeleteIncomingPhoneNumberAssignedAddOnParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ResourceSid"+"}", ResourceSid, -1)
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

// Optional parameters for the method 'FetchIncomingPhoneNumberAssignedAddOn'
type FetchIncomingPhoneNumberAssignedAddOnParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchIncomingPhoneNumberAssignedAddOnParams) SetPathAccountSid(PathAccountSid string) *FetchIncomingPhoneNumberAssignedAddOnParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch an instance of an Add-on installation currently assigned to this Number.
func (c *ApiService) FetchIncomingPhoneNumberAssignedAddOn(ResourceSid string, Sid string, params *FetchIncomingPhoneNumberAssignedAddOnParams) (*ApiV2010IncomingPhoneNumberAssignedAddOn, error) {
	return c.FetchIncomingPhoneNumberAssignedAddOnWithCtx(context.TODO(), ResourceSid, Sid, params)
}

// Fetch an instance of an Add-on installation currently assigned to this Number.
func (c *ApiService) FetchIncomingPhoneNumberAssignedAddOnWithCtx(ctx context.Context, ResourceSid string, Sid string, params *FetchIncomingPhoneNumberAssignedAddOnParams) (*ApiV2010IncomingPhoneNumberAssignedAddOn, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ResourceSid"+"}", ResourceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010IncomingPhoneNumberAssignedAddOn{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListIncomingPhoneNumberAssignedAddOn'
type ListIncomingPhoneNumberAssignedAddOnParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListIncomingPhoneNumberAssignedAddOnParams) SetPathAccountSid(PathAccountSid string) *ListIncomingPhoneNumberAssignedAddOnParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListIncomingPhoneNumberAssignedAddOnParams) SetPageSize(PageSize int) *ListIncomingPhoneNumberAssignedAddOnParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListIncomingPhoneNumberAssignedAddOnParams) SetLimit(Limit int) *ListIncomingPhoneNumberAssignedAddOnParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of IncomingPhoneNumberAssignedAddOn records from the API. Request is executed immediately.
func (c *ApiService) PageIncomingPhoneNumberAssignedAddOn(ResourceSid string, params *ListIncomingPhoneNumberAssignedAddOnParams, pageToken, pageNumber string) (*ListIncomingPhoneNumberAssignedAddOnResponse, error) {
	return c.PageIncomingPhoneNumberAssignedAddOnWithCtx(context.TODO(), ResourceSid, params, pageToken, pageNumber)
}

// Retrieve a single page of IncomingPhoneNumberAssignedAddOn records from the API. Request is executed immediately.
func (c *ApiService) PageIncomingPhoneNumberAssignedAddOnWithCtx(ctx context.Context, ResourceSid string, params *ListIncomingPhoneNumberAssignedAddOnParams, pageToken, pageNumber string) (*ListIncomingPhoneNumberAssignedAddOnResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ResourceSid"+"}", ResourceSid, -1)

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

	ps := &ListIncomingPhoneNumberAssignedAddOnResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists IncomingPhoneNumberAssignedAddOn records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListIncomingPhoneNumberAssignedAddOn(ResourceSid string, params *ListIncomingPhoneNumberAssignedAddOnParams) ([]ApiV2010IncomingPhoneNumberAssignedAddOn, error) {
	return c.ListIncomingPhoneNumberAssignedAddOnWithCtx(context.TODO(), ResourceSid, params)
}

// Lists IncomingPhoneNumberAssignedAddOn records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListIncomingPhoneNumberAssignedAddOnWithCtx(ctx context.Context, ResourceSid string, params *ListIncomingPhoneNumberAssignedAddOnParams) ([]ApiV2010IncomingPhoneNumberAssignedAddOn, error) {
	response, errors := c.StreamIncomingPhoneNumberAssignedAddOnWithCtx(ctx, ResourceSid, params)

	records := make([]ApiV2010IncomingPhoneNumberAssignedAddOn, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams IncomingPhoneNumberAssignedAddOn records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamIncomingPhoneNumberAssignedAddOn(ResourceSid string, params *ListIncomingPhoneNumberAssignedAddOnParams) (chan ApiV2010IncomingPhoneNumberAssignedAddOn, chan error) {
	return c.StreamIncomingPhoneNumberAssignedAddOnWithCtx(context.TODO(), ResourceSid, params)
}

// Streams IncomingPhoneNumberAssignedAddOn records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamIncomingPhoneNumberAssignedAddOnWithCtx(ctx context.Context, ResourceSid string, params *ListIncomingPhoneNumberAssignedAddOnParams) (chan ApiV2010IncomingPhoneNumberAssignedAddOn, chan error) {
	if params == nil {
		params = &ListIncomingPhoneNumberAssignedAddOnParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ApiV2010IncomingPhoneNumberAssignedAddOn, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageIncomingPhoneNumberAssignedAddOnWithCtx(ctx, ResourceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamIncomingPhoneNumberAssignedAddOn(ctx, response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamIncomingPhoneNumberAssignedAddOn(ctx context.Context, response *ListIncomingPhoneNumberAssignedAddOnResponse, params *ListIncomingPhoneNumberAssignedAddOnParams, recordChannel chan ApiV2010IncomingPhoneNumberAssignedAddOn, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.AssignedAddOns
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNextWithCtx(ctx, c.baseURL, response, c.getNextListIncomingPhoneNumberAssignedAddOnResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListIncomingPhoneNumberAssignedAddOnResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListIncomingPhoneNumberAssignedAddOnResponse(ctx context.Context, nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(ctx, nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListIncomingPhoneNumberAssignedAddOnResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
