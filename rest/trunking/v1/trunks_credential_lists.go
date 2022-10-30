/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Trunking
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

// Optional parameters for the method 'CreateCredentialList'
type CreateCredentialListParams struct {
	// The SID of the [Credential List](https://www.twilio.com/docs/voice/sip/api/sip-credentiallist-resource) that you want to associate with the trunk. Once associated, we will authenticate access to the trunk against this list.
	CredentialListSid *string `json:"CredentialListSid,omitempty"`
}

func (params *CreateCredentialListParams) SetCredentialListSid(CredentialListSid string) *CreateCredentialListParams {
	params.CredentialListSid = &CredentialListSid
	return params
}

//
func (c *ApiService) CreateCredentialList(TrunkSid string, params *CreateCredentialListParams) (*TrunkingV1CredentialList, error) {
	return c.CreateCredentialListWithCtx(context.TODO(), TrunkSid, params)
}

//
func (c *ApiService) CreateCredentialListWithCtx(ctx context.Context, TrunkSid string, params *CreateCredentialListParams) (*TrunkingV1CredentialList, error) {
	path := "/v1/Trunks/{TrunkSid}/CredentialLists"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CredentialListSid != nil {
		data.Set("CredentialListSid", *params.CredentialListSid)
	}

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1CredentialList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//
func (c *ApiService) DeleteCredentialList(TrunkSid string, Sid string) error {
	return c.DeleteCredentialListWithCtx(context.TODO(), TrunkSid, Sid)
}

//
func (c *ApiService) DeleteCredentialListWithCtx(ctx context.Context, TrunkSid string, Sid string) error {
	path := "/v1/Trunks/{TrunkSid}/CredentialLists/{Sid}"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)
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

//
func (c *ApiService) FetchCredentialList(TrunkSid string, Sid string) (*TrunkingV1CredentialList, error) {
	return c.FetchCredentialListWithCtx(context.TODO(), TrunkSid, Sid)
}

//
func (c *ApiService) FetchCredentialListWithCtx(ctx context.Context, TrunkSid string, Sid string) (*TrunkingV1CredentialList, error) {
	path := "/v1/Trunks/{TrunkSid}/CredentialLists/{Sid}"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1CredentialList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListCredentialList'
type ListCredentialListParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListCredentialListParams) SetPageSize(PageSize int) *ListCredentialListParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListCredentialListParams) SetLimit(Limit int) *ListCredentialListParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of CredentialList records from the API. Request is executed immediately.
func (c *ApiService) PageCredentialList(TrunkSid string, params *ListCredentialListParams, pageToken, pageNumber string) (*ListCredentialListResponse, error) {
	return c.PageCredentialListWithCtx(context.TODO(), TrunkSid, params, pageToken, pageNumber)
}

// Retrieve a single page of CredentialList records from the API. Request is executed immediately.
func (c *ApiService) PageCredentialListWithCtx(ctx context.Context, TrunkSid string, params *ListCredentialListParams, pageToken, pageNumber string) (*ListCredentialListResponse, error) {
	path := "/v1/Trunks/{TrunkSid}/CredentialLists"

	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

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

	ps := &ListCredentialListResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists CredentialList records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListCredentialList(TrunkSid string, params *ListCredentialListParams) ([]TrunkingV1CredentialList, error) {
	return c.ListCredentialListWithCtx(context.TODO(), TrunkSid, params)
}

// Lists CredentialList records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListCredentialListWithCtx(ctx context.Context, TrunkSid string, params *ListCredentialListParams) ([]TrunkingV1CredentialList, error) {
	response, errors := c.StreamCredentialListWithCtx(ctx, TrunkSid, params)

	records := make([]TrunkingV1CredentialList, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams CredentialList records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamCredentialList(TrunkSid string, params *ListCredentialListParams) (chan TrunkingV1CredentialList, chan error) {
	return c.StreamCredentialListWithCtx(context.TODO(), TrunkSid, params)
}

// Streams CredentialList records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamCredentialListWithCtx(ctx context.Context, TrunkSid string, params *ListCredentialListParams) (chan TrunkingV1CredentialList, chan error) {
	if params == nil {
		params = &ListCredentialListParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan TrunkingV1CredentialList, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageCredentialListWithCtx(ctx, TrunkSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamCredentialList(ctx, response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamCredentialList(ctx context.Context, response *ListCredentialListResponse, params *ListCredentialListParams, recordChannel chan TrunkingV1CredentialList, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.CredentialLists
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNextWithCtx(ctx, c.baseURL, response, c.getNextListCredentialListResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListCredentialListResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListCredentialListResponse(ctx context.Context, nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(ctx, nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCredentialListResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
