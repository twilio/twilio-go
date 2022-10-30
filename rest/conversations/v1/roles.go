/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Conversations
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

// Optional parameters for the method 'CreateRole'
type CreateRoleParams struct {
	// A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	//
	Type *string `json:"Type,omitempty"`
	// A permission that you grant to the new role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. The values for this parameter depend on the role's `type`.
	Permission *[]string `json:"Permission,omitempty"`
}

func (params *CreateRoleParams) SetFriendlyName(FriendlyName string) *CreateRoleParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateRoleParams) SetType(Type string) *CreateRoleParams {
	params.Type = &Type
	return params
}
func (params *CreateRoleParams) SetPermission(Permission []string) *CreateRoleParams {
	params.Permission = &Permission
	return params
}

// Create a new user role in your account&#39;s default service
func (c *ApiService) CreateRole(params *CreateRoleParams) (*ConversationsV1Role, error) {
	return c.CreateRoleWithCtx(context.TODO(), params)
}

// Create a new user role in your account&#39;s default service
func (c *ApiService) CreateRoleWithCtx(ctx context.Context, params *CreateRoleParams) (*ConversationsV1Role, error) {
	path := "/v1/Roles"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Type != nil {
		data.Set("Type", *params.Type)
	}
	if params != nil && params.Permission != nil {
		for _, item := range *params.Permission {
			data.Add("Permission", item)
		}
	}

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1Role{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove a user role from your account&#39;s default service
func (c *ApiService) DeleteRole(Sid string) error {
	return c.DeleteRoleWithCtx(context.TODO(), Sid)
}

// Remove a user role from your account&#39;s default service
func (c *ApiService) DeleteRoleWithCtx(ctx context.Context, Sid string) error {
	path := "/v1/Roles/{Sid}"
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

// Fetch a user role from your account&#39;s default service
func (c *ApiService) FetchRole(Sid string) (*ConversationsV1Role, error) {
	return c.FetchRoleWithCtx(context.TODO(), Sid)
}

// Fetch a user role from your account&#39;s default service
func (c *ApiService) FetchRoleWithCtx(ctx context.Context, Sid string) (*ConversationsV1Role, error) {
	path := "/v1/Roles/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1Role{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRole'
type ListRoleParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListRoleParams) SetPageSize(PageSize int) *ListRoleParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListRoleParams) SetLimit(Limit int) *ListRoleParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Role records from the API. Request is executed immediately.
func (c *ApiService) PageRole(params *ListRoleParams, pageToken, pageNumber string) (*ListRoleResponse, error) {
	return c.PageRoleWithCtx(context.TODO(), params, pageToken, pageNumber)
}

// Retrieve a single page of Role records from the API. Request is executed immediately.
func (c *ApiService) PageRoleWithCtx(ctx context.Context, params *ListRoleParams, pageToken, pageNumber string) (*ListRoleResponse, error) {
	path := "/v1/Roles"

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

	ps := &ListRoleResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Role records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListRole(params *ListRoleParams) ([]ConversationsV1Role, error) {
	return c.ListRoleWithCtx(context.TODO(), params)
}

// Lists Role records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListRoleWithCtx(ctx context.Context, params *ListRoleParams) ([]ConversationsV1Role, error) {
	response, errors := c.StreamRoleWithCtx(ctx, params)

	records := make([]ConversationsV1Role, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Role records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamRole(params *ListRoleParams) (chan ConversationsV1Role, chan error) {
	return c.StreamRoleWithCtx(context.TODO(), params)
}

// Streams Role records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamRoleWithCtx(ctx context.Context, params *ListRoleParams) (chan ConversationsV1Role, chan error) {
	if params == nil {
		params = &ListRoleParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ConversationsV1Role, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageRoleWithCtx(ctx, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamRole(ctx, response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamRole(ctx context.Context, response *ListRoleResponse, params *ListRoleParams, recordChannel chan ConversationsV1Role, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Roles
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNextWithCtx(ctx, c.baseURL, response, c.getNextListRoleResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListRoleResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListRoleResponse(ctx context.Context, nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(ctx, nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRoleResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateRole'
type UpdateRoleParams struct {
	// A permission that you grant to the role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. Note that the update action replaces all previously assigned permissions with those defined in the update action. To remove a permission, do not include it in the subsequent update action. The values for this parameter depend on the role's `type`.
	Permission *[]string `json:"Permission,omitempty"`
}

func (params *UpdateRoleParams) SetPermission(Permission []string) *UpdateRoleParams {
	params.Permission = &Permission
	return params
}

// Update an existing user role in your account&#39;s default service
func (c *ApiService) UpdateRole(Sid string, params *UpdateRoleParams) (*ConversationsV1Role, error) {
	return c.UpdateRoleWithCtx(context.TODO(), Sid, params)
}

// Update an existing user role in your account&#39;s default service
func (c *ApiService) UpdateRoleWithCtx(ctx context.Context, Sid string, params *UpdateRoleParams) (*ConversationsV1Role, error) {
	path := "/v1/Roles/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Permission != nil {
		for _, item := range *params.Permission {
			data.Add("Permission", item)
		}
	}

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1Role{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
