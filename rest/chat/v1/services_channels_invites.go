/*
 * Twilio - Chat
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

// Optional parameters for the method 'CreateInvite'
type CreateInviteParams struct {
	// The `identity` value that uniquely identifies the new resource's [User](https://www.twilio.com/docs/api/chat/rest/v1/user) within the [Service](https://www.twilio.com/docs/api/chat/rest/v1/service). See [access tokens](https://www.twilio.com/docs/api/chat/guides/create-tokens) for more info.
	Identity *string `json:"Identity,omitempty"`
	// The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) assigned to the new member.
	RoleSid *string `json:"RoleSid,omitempty"`
}

func (params *CreateInviteParams) SetIdentity(Identity string) *CreateInviteParams {
	params.Identity = &Identity
	return params
}
func (params *CreateInviteParams) SetRoleSid(RoleSid string) *CreateInviteParams {
	params.RoleSid = &RoleSid
	return params
}

func (c *ApiService) CreateInvite(ServiceSid string, ChannelSid string, params *CreateInviteParams) (*ChatV1ServiceChannelInvite, error) {
	path := "/v1/Services/{ServiceSid}/Channels/{ChannelSid}/Invites"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

	data := url.Values{}
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

	ps := &ChatV1ServiceChannelInvite{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteInvite(ServiceSid string, ChannelSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Channels/{ChannelSid}/Invites/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)
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

func (c *ApiService) FetchInvite(ServiceSid string, ChannelSid string, Sid string) (*ChatV1ServiceChannelInvite, error) {
	path := "/v1/Services/{ServiceSid}/Channels/{ChannelSid}/Invites/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ChatV1ServiceChannelInvite{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListInvite'
type ListInviteParams struct {
	// The [User](https://www.twilio.com/docs/api/chat/rest/v1/user)'s `identity` value of the resources to read. See [access tokens](https://www.twilio.com/docs/api/chat/guides/create-tokens) for more details.
	Identity *[]string `json:"Identity,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListInviteParams) SetIdentity(Identity []string) *ListInviteParams {
	params.Identity = &Identity
	return params
}
func (params *ListInviteParams) SetPageSize(PageSize int) *ListInviteParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListInvite(ServiceSid string, ChannelSid string, params *ListInviteParams) (*ListInviteResponse, error) {
	path := "/v1/Services/{ServiceSid}/Channels/{ChannelSid}/Invites"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

	data := url.Values{}
	if params != nil && params.Identity != nil {
		for _, item := range *params.Identity {
			data.Add("Identity", item)
		}
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListInviteResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
