/*
 * Twilio - Ip_messaging
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
)

// Optional parameters for the method 'CreateUser'
type CreateUserParams struct {
	//
	Attributes *string `json:"Attributes,omitempty"`
	//
	FriendlyName *string `json:"FriendlyName,omitempty"`
	//
	Identity *string `json:"Identity,omitempty"`
	//
	RoleSid *string `json:"RoleSid,omitempty"`
}

func (params *CreateUserParams) SetAttributes(Attributes string) *CreateUserParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateUserParams) SetFriendlyName(FriendlyName string) *CreateUserParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateUserParams) SetIdentity(Identity string) *CreateUserParams {
	params.Identity = &Identity
	return params
}
func (params *CreateUserParams) SetRoleSid(RoleSid string) *CreateUserParams {
	params.RoleSid = &RoleSid
	return params
}

func (c *ApiService) CreateUser(ServiceSid string, params *CreateUserParams) (*IpMessagingV1ServiceUser, error) {
	path := "/v1/Services/{ServiceSid}/Users"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Identity != nil {
		data.Set("Identity", *params.Identity)
	}
	if params != nil && params.RoleSid != nil {
		data.Set("RoleSid", *params.RoleSid)
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV1ServiceUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteUser(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Users/{Sid}"
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

func (c *ApiService) FetchUser(ServiceSid string, Sid string) (*IpMessagingV1ServiceUser, error) {
	path := "/v1/Services/{ServiceSid}/Users/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV1ServiceUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListUser'
type ListUserParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListUserParams) SetPageSize(PageSize int) *ListUserParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListUser(ServiceSid string, params *ListUserParams) (*ListUserResponse, error) {
	path := "/v1/Services/{ServiceSid}/Users"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateUser'
type UpdateUserParams struct {
	//
	Attributes *string `json:"Attributes,omitempty"`
	//
	FriendlyName *string `json:"FriendlyName,omitempty"`
	//
	RoleSid *string `json:"RoleSid,omitempty"`
}

func (params *UpdateUserParams) SetAttributes(Attributes string) *UpdateUserParams {
	params.Attributes = &Attributes
	return params
}
func (params *UpdateUserParams) SetFriendlyName(FriendlyName string) *UpdateUserParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateUserParams) SetRoleSid(RoleSid string) *UpdateUserParams {
	params.RoleSid = &RoleSid
	return params
}

func (c *ApiService) UpdateUser(ServiceSid string, Sid string, params *UpdateUserParams) (*IpMessagingV1ServiceUser, error) {
	path := "/v1/Services/{ServiceSid}/Users/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.RoleSid != nil {
		data.Set("RoleSid", *params.RoleSid)
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV1ServiceUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
