/*
 * Twilio - Ip_messaging
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

// Optional parameters for the method 'CreateRole'
type CreateRoleParams struct {
	//
	FriendlyName *string `json:"FriendlyName,omitempty"`
	//
	Permission *[]string `json:"Permission,omitempty"`
	//
	Type *string `json:"Type,omitempty"`
}

func (params *CreateRoleParams) SetFriendlyName(FriendlyName string) *CreateRoleParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateRoleParams) SetPermission(Permission []string) *CreateRoleParams {
	params.Permission = &Permission
	return params
}
func (params *CreateRoleParams) SetType(Type string) *CreateRoleParams {
	params.Type = &Type
	return params
}

func (c *ApiService) CreateRole(ServiceSid string, params *CreateRoleParams) (*IpMessagingV2ServiceRole, error) {
	path := "/v2/Services/{ServiceSid}/Roles"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

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

	ps := &IpMessagingV2ServiceRole{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteRole(ServiceSid string, Sid string) error {
	path := "/v2/Services/{ServiceSid}/Roles/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
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

func (c *ApiService) FetchRole(ServiceSid string, Sid string) (*IpMessagingV2ServiceRole, error) {
	path := "/v2/Services/{ServiceSid}/Roles/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV2ServiceRole{}
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
func (c *ApiService) PageRole(ServiceSid string, params *ListRoleParams, pageToken string, pageNumber string) (*ListRoleResponse, error) {
	path := "/v2/Services/{ServiceSid}/Roles"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

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

	ps := &ListRoleResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Role records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListRole(ServiceSid string, params *ListRoleParams) ([]IpMessagingV2ServiceRole, error) {
	if params == nil {
		params = &ListRoleParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageRole(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []IpMessagingV2ServiceRole

	for response != nil {
		records = append(records, response.Roles...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListRoleResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListRoleResponse)
	}

	return records, err
}

// Streams Role records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamRole(ServiceSid string, params *ListRoleParams) (chan IpMessagingV2ServiceRole, error) {
	if params == nil {
		params = &ListRoleParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageRole(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan IpMessagingV2ServiceRole, 1)

	go func() {
		for response != nil {
			for item := range response.Roles {
				channel <- response.Roles[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListRoleResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListRoleResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListRoleResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
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
	//
	Permission *[]string `json:"Permission,omitempty"`
}

func (params *UpdateRoleParams) SetPermission(Permission []string) *UpdateRoleParams {
	params.Permission = &Permission
	return params
}

func (c *ApiService) UpdateRole(ServiceSid string, Sid string, params *UpdateRoleParams) (*IpMessagingV2ServiceRole, error) {
	path := "/v2/Services/{ServiceSid}/Roles/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
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

	ps := &IpMessagingV2ServiceRole{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
