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

// Delete a specific Sync Map Permission.
func (c *ApiService) DeleteSyncMapPermission(ServiceSid string, MapSid string, Identity string) error {
	return c.DeleteSyncMapPermissionWithCtx(context.TODO(), ServiceSid, MapSid, Identity)
}

// Delete a specific Sync Map Permission.
func (c *ApiService) DeleteSyncMapPermissionWithCtx(ctx context.Context, ServiceSid string, MapSid string, Identity string) error {
	path := "/v1/Services/{ServiceSid}/Maps/{MapSid}/Permissions/{Identity}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"MapSid"+"}", MapSid, -1)
	path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch a specific Sync Map Permission.
func (c *ApiService) FetchSyncMapPermission(ServiceSid string, MapSid string, Identity string) (*SyncV1SyncMapPermission, error) {
	return c.FetchSyncMapPermissionWithCtx(context.TODO(), ServiceSid, MapSid, Identity)
}

// Fetch a specific Sync Map Permission.
func (c *ApiService) FetchSyncMapPermissionWithCtx(ctx context.Context, ServiceSid string, MapSid string, Identity string) (*SyncV1SyncMapPermission, error) {
	path := "/v1/Services/{ServiceSid}/Maps/{MapSid}/Permissions/{Identity}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"MapSid"+"}", MapSid, -1)
	path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1SyncMapPermission{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSyncMapPermission'
type ListSyncMapPermissionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSyncMapPermissionParams) SetPageSize(PageSize int) *ListSyncMapPermissionParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSyncMapPermissionParams) SetLimit(Limit int) *ListSyncMapPermissionParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SyncMapPermission records from the API. Request is executed immediately.
func (c *ApiService) PageSyncMapPermission(ServiceSid string, MapSid string, params *ListSyncMapPermissionParams, pageToken, pageNumber string) (*ListSyncMapPermissionResponse, error) {
	return c.PageSyncMapPermissionWithCtx(context.TODO(), ServiceSid, MapSid, params, pageToken, pageNumber)
}

// Retrieve a single page of SyncMapPermission records from the API. Request is executed immediately.
func (c *ApiService) PageSyncMapPermissionWithCtx(ctx context.Context, ServiceSid string, MapSid string, params *ListSyncMapPermissionParams, pageToken, pageNumber string) (*ListSyncMapPermissionResponse, error) {
	path := "/v1/Services/{ServiceSid}/Maps/{MapSid}/Permissions"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"MapSid"+"}", MapSid, -1)

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

	ps := &ListSyncMapPermissionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SyncMapPermission records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSyncMapPermission(ServiceSid string, MapSid string, params *ListSyncMapPermissionParams) ([]SyncV1SyncMapPermission, error) {
	return c.ListSyncMapPermissionWithCtx(context.TODO(), ServiceSid, MapSid, params)
}

// Lists SyncMapPermission records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSyncMapPermissionWithCtx(ctx context.Context, ServiceSid string, MapSid string, params *ListSyncMapPermissionParams) ([]SyncV1SyncMapPermission, error) {
	response, errors := c.StreamSyncMapPermissionWithCtx(ctx, ServiceSid, MapSid, params)

	records := make([]SyncV1SyncMapPermission, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams SyncMapPermission records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSyncMapPermission(ServiceSid string, MapSid string, params *ListSyncMapPermissionParams) (chan SyncV1SyncMapPermission, chan error) {
	return c.StreamSyncMapPermissionWithCtx(context.TODO(), ServiceSid, MapSid, params)
}

// Streams SyncMapPermission records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSyncMapPermissionWithCtx(ctx context.Context, ServiceSid string, MapSid string, params *ListSyncMapPermissionParams) (chan SyncV1SyncMapPermission, chan error) {
	if params == nil {
		params = &ListSyncMapPermissionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan SyncV1SyncMapPermission, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageSyncMapPermissionWithCtx(ctx, ServiceSid, MapSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamSyncMapPermission(ctx, response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamSyncMapPermission(ctx context.Context, response *ListSyncMapPermissionResponse, params *ListSyncMapPermissionParams, recordChannel chan SyncV1SyncMapPermission, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Permissions
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNextWithCtx(ctx, c.baseURL, response, c.getNextListSyncMapPermissionResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListSyncMapPermissionResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListSyncMapPermissionResponse(ctx context.Context, nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(ctx, nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSyncMapPermissionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSyncMapPermission'
type UpdateSyncMapPermissionParams struct {
	// Whether the identity can read the Sync Map and its Items. Default value is `false`.
	Read *bool `json:"Read,omitempty"`
	// Whether the identity can create, update, and delete Items in the Sync Map. Default value is `false`.
	Write *bool `json:"Write,omitempty"`
	// Whether the identity can delete the Sync Map. Default value is `false`.
	Manage *bool `json:"Manage,omitempty"`
}

func (params *UpdateSyncMapPermissionParams) SetRead(Read bool) *UpdateSyncMapPermissionParams {
	params.Read = &Read
	return params
}
func (params *UpdateSyncMapPermissionParams) SetWrite(Write bool) *UpdateSyncMapPermissionParams {
	params.Write = &Write
	return params
}
func (params *UpdateSyncMapPermissionParams) SetManage(Manage bool) *UpdateSyncMapPermissionParams {
	params.Manage = &Manage
	return params
}

// Update an identity&#39;s access to a specific Sync Map.
func (c *ApiService) UpdateSyncMapPermission(ServiceSid string, MapSid string, Identity string, params *UpdateSyncMapPermissionParams) (*SyncV1SyncMapPermission, error) {
	return c.UpdateSyncMapPermissionWithCtx(context.TODO(), ServiceSid, MapSid, Identity, params)
}

// Update an identity&#39;s access to a specific Sync Map.
func (c *ApiService) UpdateSyncMapPermissionWithCtx(ctx context.Context, ServiceSid string, MapSid string, Identity string, params *UpdateSyncMapPermissionParams) (*SyncV1SyncMapPermission, error) {
	path := "/v1/Services/{ServiceSid}/Maps/{MapSid}/Permissions/{Identity}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"MapSid"+"}", MapSid, -1)
	path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Read != nil {
		data.Set("Read", fmt.Sprint(*params.Read))
	}
	if params != nil && params.Write != nil {
		data.Set("Write", fmt.Sprint(*params.Write))
	}
	if params != nil && params.Manage != nil {
		data.Set("Manage", fmt.Sprint(*params.Manage))
	}

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1SyncMapPermission{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
