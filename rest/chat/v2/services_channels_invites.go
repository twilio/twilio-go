/*
 * Twilio - Chat
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

// Optional parameters for the method 'CreateInvite'
type CreateInviteParams struct {
	// The `identity` value that uniquely identifies the new resource's [User](https://www.twilio.com/docs/chat/rest/user-resource) within the [Service](https://www.twilio.com/docs/chat/rest/service-resource). See [access tokens](https://www.twilio.com/docs/chat/create-tokens) for more info.
	Identity *string `json:"Identity,omitempty"`
	// The SID of the [Role](https://www.twilio.com/docs/chat/rest/role-resource) assigned to the new member.
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

func (c *ApiService) CreateInvite(ServiceSid string, ChannelSid string, params *CreateInviteParams) (*ChatV2ServiceChannelInvite, error) {
	path := "/v2/Services/{ServiceSid}/Channels/{ChannelSid}/Invites"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Identity != nil {
		data.Set("Identity", *params.Identity)
	}
	if params != nil && params.RoleSid != nil {
		data.Set("RoleSid", *params.RoleSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ChatV2ServiceChannelInvite{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteInvite(ServiceSid string, ChannelSid string, Sid string) error {
	path := "/v2/Services/{ServiceSid}/Channels/{ChannelSid}/Invites/{Sid}"
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

func (c *ApiService) FetchInvite(ServiceSid string, ChannelSid string, Sid string) (*ChatV2ServiceChannelInvite, error) {
	path := "/v2/Services/{ServiceSid}/Channels/{ChannelSid}/Invites/{Sid}"
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

	ps := &ChatV2ServiceChannelInvite{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListInvite'
type ListInviteParams struct {
	// The [User](https://www.twilio.com/docs/chat/rest/user-resource)'s `identity` value of the resources to read. See [access tokens](https://www.twilio.com/docs/chat/create-tokens) for more details.
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

// Retrieve a single page of Invite records from the API. Request is executed immediately.
func (c *ApiService) PageInvite(ServiceSid string, ChannelSid string, params *ListInviteParams, pageToken string, pageNumber string) (*ListInviteResponse, error) {
	path := "/v2/Services/{ServiceSid}/Channels/{ChannelSid}/Invites"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Identity != nil {
		for _, item := range *params.Identity {
			data.Add("Identity", item)
		}
	}
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

	ps := &ListInviteResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Invite records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListInvite(ServiceSid string, ChannelSid string, params *ListInviteParams, limit int) ([]ChatV2ServiceChannelInvite, error) {
	if params == nil {
		params = &ListInviteParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageInvite(ServiceSid, ChannelSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ChatV2ServiceChannelInvite

	for response != nil {
		records = append(records, response.Invites...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, limit, c.getNextListInviteResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListInviteResponse)
	}

	return records, err
}

// Streams Invite records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamInvite(ServiceSid string, ChannelSid string, params *ListInviteParams, limit int) (chan ChatV2ServiceChannelInvite, error) {
	if params == nil {
		params = &ListInviteParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageInvite(ServiceSid, ChannelSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ChatV2ServiceChannelInvite, 1)

	go func() {
		for response != nil {
			for item := range response.Invites {
				channel <- response.Invites[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, limit, c.getNextListInviteResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListInviteResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListInviteResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListInviteResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
