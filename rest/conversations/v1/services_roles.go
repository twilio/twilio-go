/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.0
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

// Optional parameters for the method 'CreateServiceRole'
type CreateServiceRoleParams struct {
	// A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// A permission that you grant to the new role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. The values for this parameter depend on the role's `type`.
	Permission *[]string `json:"Permission,omitempty"`
	// The type of role. Can be: `conversation` for [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) roles or `service` for [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) roles.
	Type *string `json:"Type,omitempty"`
}

func (params *CreateServiceRoleParams) SetFriendlyName(FriendlyName string) *CreateServiceRoleParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateServiceRoleParams) SetPermission(Permission []string) *CreateServiceRoleParams {
	params.Permission = &Permission
	return params
}
func (params *CreateServiceRoleParams) SetType(Type string) *CreateServiceRoleParams {
	params.Type = &Type
	return params
}

// Create a new user role in your service
func (c *ApiService) CreateServiceRole(ChatServiceSid string, params *CreateServiceRoleParams) (*ConversationsV1ServiceRole, error) {
	path := "/v1/Services/{ChatServiceSid}/Roles"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Permission != nil {
		for _, item := range *params.Permission {
			data.Add("Permission", item)
		}
	}
	if params != nil && params.Type != nil {
		data.Set("Type", *params.Type)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceRole{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove a user role from your service
func (c *ApiService) DeleteServiceRole(ChatServiceSid string, Sid string) error {
	path := "/v1/Services/{ChatServiceSid}/Roles/{Sid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
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

// Fetch a user role from your service
func (c *ApiService) FetchServiceRole(ChatServiceSid string, Sid string) (*ConversationsV1ServiceRole, error) {
	path := "/v1/Services/{ChatServiceSid}/Roles/{Sid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceRole{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListServiceRole'
type ListServiceRoleParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListServiceRoleParams) SetPageSize(PageSize int) *ListServiceRoleParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListServiceRoleParams) SetLimit(Limit int) *ListServiceRoleParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ServiceRole records from the API. Request is executed immediately.
func (c *ApiService) PageServiceRole(ChatServiceSid string, params *ListServiceRoleParams, pageToken, pageNumber string) (*ListServiceRoleResponse, error) {
	path := "/v1/Services/{ChatServiceSid}/Roles"

	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)

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

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListServiceRoleResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ServiceRole records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListServiceRole(ChatServiceSid string, params *ListServiceRoleParams) ([]ConversationsV1ServiceRole, error) {
	if params == nil {
		params = &ListServiceRoleParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageServiceRole(ChatServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ConversationsV1ServiceRole

	for response != nil {
		records = append(records, response.Roles...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListServiceRoleResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListServiceRoleResponse)
	}

	return records, err
}

// Streams ServiceRole records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamServiceRole(ChatServiceSid string, params *ListServiceRoleParams) (chan ConversationsV1ServiceRole, error) {
	if params == nil {
		params = &ListServiceRoleParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageServiceRole(ChatServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ConversationsV1ServiceRole, 1)

	go func() {
		for response != nil {
			for item := range response.Roles {
				channel <- response.Roles[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListServiceRoleResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListServiceRoleResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListServiceRoleResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListServiceRoleResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateServiceRole'
type UpdateServiceRoleParams struct {
	// A permission that you grant to the role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. Note that the update action replaces all previously assigned permissions with those defined in the update action. To remove a permission, do not include it in the subsequent update action. The values for this parameter depend on the role's `type`.
	Permission *[]string `json:"Permission,omitempty"`
}

func (params *UpdateServiceRoleParams) SetPermission(Permission []string) *UpdateServiceRoleParams {
	params.Permission = &Permission
	return params
}

// Update an existing user role in your service
func (c *ApiService) UpdateServiceRole(ChatServiceSid string, Sid string, params *UpdateServiceRoleParams) (*ConversationsV1ServiceRole, error) {
	path := "/v1/Services/{ChatServiceSid}/Roles/{Sid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Permission != nil {
		for _, item := range *params.Permission {
			data.Add("Permission", item)
		}
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceRole{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
