/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Sync
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

// Optional parameters for the method 'CreateSyncMap'
type CreateSyncMapParams struct {
	// An application-defined string that uniquely identifies the resource. It can be used as an alternative to the `sid` in the URL path to address the resource.
	UniqueName *string `json:"UniqueName,omitempty"`
	// An alias for `collection_ttl`. If both parameters are provided, this value is ignored.
	Ttl *int `json:"Ttl,omitempty"`
	// How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync Map expires (time-to-live) and is deleted.
	CollectionTtl *int `json:"CollectionTtl,omitempty"`
}

func (params *CreateSyncMapParams) SetUniqueName(UniqueName string) *CreateSyncMapParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *CreateSyncMapParams) SetTtl(Ttl int) *CreateSyncMapParams {
	params.Ttl = &Ttl
	return params
}
func (params *CreateSyncMapParams) SetCollectionTtl(CollectionTtl int) *CreateSyncMapParams {
	params.CollectionTtl = &CollectionTtl
	return params
}

//
func (c *ApiService) CreateSyncMap(ServiceSid string, params *CreateSyncMapParams) (*SyncV1SyncMap, error) {
	return c.CreateSyncMapWithCtx(context.TODO(), ServiceSid, params)
}

//
func (c *ApiService) CreateSyncMapWithCtx(ctx context.Context, ServiceSid string, params *CreateSyncMapParams) (*SyncV1SyncMap, error) {
	path := "/v1/Services/{ServiceSid}/Maps"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.Ttl != nil {
		data.Set("Ttl", fmt.Sprint(*params.Ttl))
	}
	if params != nil && params.CollectionTtl != nil {
		data.Set("CollectionTtl", fmt.Sprint(*params.CollectionTtl))
	}

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1SyncMap{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//
func (c *ApiService) DeleteSyncMap(ServiceSid string, Sid string) error {
	return c.DeleteSyncMapWithCtx(context.TODO(), ServiceSid, Sid)
}

//
func (c *ApiService) DeleteSyncMapWithCtx(ctx context.Context, ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Maps/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
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
func (c *ApiService) FetchSyncMap(ServiceSid string, Sid string) (*SyncV1SyncMap, error) {
	return c.FetchSyncMapWithCtx(context.TODO(), ServiceSid, Sid)
}

//
func (c *ApiService) FetchSyncMapWithCtx(ctx context.Context, ServiceSid string, Sid string) (*SyncV1SyncMap, error) {
	path := "/v1/Services/{ServiceSid}/Maps/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1SyncMap{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSyncMap'
type ListSyncMapParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSyncMapParams) SetPageSize(PageSize int) *ListSyncMapParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSyncMapParams) SetLimit(Limit int) *ListSyncMapParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SyncMap records from the API. Request is executed immediately.
func (c *ApiService) PageSyncMap(ServiceSid string, params *ListSyncMapParams, pageToken, pageNumber string) (*ListSyncMapResponse, error) {
	return c.PageSyncMapWithCtx(context.TODO(), ServiceSid, params, pageToken, pageNumber)
}

// Retrieve a single page of SyncMap records from the API. Request is executed immediately.
func (c *ApiService) PageSyncMapWithCtx(ctx context.Context, ServiceSid string, params *ListSyncMapParams, pageToken, pageNumber string) (*ListSyncMapResponse, error) {
	path := "/v1/Services/{ServiceSid}/Maps"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

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

	ps := &ListSyncMapResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SyncMap records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSyncMap(ServiceSid string, params *ListSyncMapParams) ([]SyncV1SyncMap, error) {
	return c.ListSyncMapWithCtx(context.TODO(), ServiceSid, params)
}

// Lists SyncMap records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSyncMapWithCtx(ctx context.Context, ServiceSid string, params *ListSyncMapParams) ([]SyncV1SyncMap, error) {
	response, errors := c.StreamSyncMapWithCtx(ctx, ServiceSid, params)

	records := make([]SyncV1SyncMap, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams SyncMap records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSyncMap(ServiceSid string, params *ListSyncMapParams) (chan SyncV1SyncMap, chan error) {
	return c.StreamSyncMapWithCtx(context.TODO(), ServiceSid, params)
}

// Streams SyncMap records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSyncMapWithCtx(ctx context.Context, ServiceSid string, params *ListSyncMapParams) (chan SyncV1SyncMap, chan error) {
	if params == nil {
		params = &ListSyncMapParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan SyncV1SyncMap, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageSyncMapWithCtx(ctx, ServiceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamSyncMap(ctx, response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamSyncMap(ctx context.Context, response *ListSyncMapResponse, params *ListSyncMapParams, recordChannel chan SyncV1SyncMap, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Maps
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNextWithCtx(ctx, c.baseURL, response, c.getNextListSyncMapResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListSyncMapResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListSyncMapResponse(ctx context.Context, nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(ctx, nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSyncMapResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSyncMap'
type UpdateSyncMapParams struct {
	// An alias for `collection_ttl`. If both parameters are provided, this value is ignored.
	Ttl *int `json:"Ttl,omitempty"`
	// How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync Map expires (time-to-live) and is deleted.
	CollectionTtl *int `json:"CollectionTtl,omitempty"`
}

func (params *UpdateSyncMapParams) SetTtl(Ttl int) *UpdateSyncMapParams {
	params.Ttl = &Ttl
	return params
}
func (params *UpdateSyncMapParams) SetCollectionTtl(CollectionTtl int) *UpdateSyncMapParams {
	params.CollectionTtl = &CollectionTtl
	return params
}

//
func (c *ApiService) UpdateSyncMap(ServiceSid string, Sid string, params *UpdateSyncMapParams) (*SyncV1SyncMap, error) {
	return c.UpdateSyncMapWithCtx(context.TODO(), ServiceSid, Sid, params)
}

//
func (c *ApiService) UpdateSyncMapWithCtx(ctx context.Context, ServiceSid string, Sid string, params *UpdateSyncMapParams) (*SyncV1SyncMap, error) {
	path := "/v1/Services/{ServiceSid}/Maps/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Ttl != nil {
		data.Set("Ttl", fmt.Sprint(*params.Ttl))
	}
	if params != nil && params.CollectionTtl != nil {
		data.Set("CollectionTtl", fmt.Sprint(*params.CollectionTtl))
	}

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1SyncMap{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
