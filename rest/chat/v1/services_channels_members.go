/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.2
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

// Optional parameters for the method 'CreateMember'
type CreateMemberParams struct {
	// The `identity` value that uniquely identifies the new resource's [User](https://www.twilio.com/docs/api/chat/rest/v1/user) within the [Service](https://www.twilio.com/docs/api/chat/rest/services). See [access tokens](https://www.twilio.com/docs/api/chat/guides/create-tokens) for more details.
	Identity *string `json:"Identity,omitempty"`
	// The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) to assign to the member. The default roles are those specified on the [Service](https://www.twilio.com/docs/chat/api/services).
	RoleSid *string `json:"RoleSid,omitempty"`
}

func (params *CreateMemberParams) SetIdentity(Identity string) *CreateMemberParams {
	params.Identity = &Identity
	return params
}
func (params *CreateMemberParams) SetRoleSid(RoleSid string) *CreateMemberParams {
	params.RoleSid = &RoleSid
	return params
}

func (c *ApiService) CreateMember(ServiceSid string, ChannelSid string, params *CreateMemberParams) (*ChatV1Member, error) {
	path := "/v1/Services/{ServiceSid}/Channels/{ChannelSid}/Members"
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

	ps := &ChatV1Member{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteMember(ServiceSid string, ChannelSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Channels/{ChannelSid}/Members/{Sid}"
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

func (c *ApiService) FetchMember(ServiceSid string, ChannelSid string, Sid string) (*ChatV1Member, error) {
	path := "/v1/Services/{ServiceSid}/Channels/{ChannelSid}/Members/{Sid}"
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

	ps := &ChatV1Member{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListMember'
type ListMemberParams struct {
	// The [User](https://www.twilio.com/docs/api/chat/rest/v1/user)'s `identity` value of the resources to read. See [access tokens](https://www.twilio.com/docs/api/chat/guides/create-tokens) for more details.
	Identity *[]string `json:"Identity,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListMemberParams) SetIdentity(Identity []string) *ListMemberParams {
	params.Identity = &Identity
	return params
}
func (params *ListMemberParams) SetPageSize(PageSize int) *ListMemberParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListMemberParams) SetLimit(Limit int) *ListMemberParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Member records from the API. Request is executed immediately.
func (c *ApiService) PageMember(ServiceSid string, ChannelSid string, params *ListMemberParams, pageToken, pageNumber string) (*ListMemberResponse, error) {
	path := "/v1/Services/{ServiceSid}/Channels/{ChannelSid}/Members"

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
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMemberResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Member records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListMember(ServiceSid string, ChannelSid string, params *ListMemberParams) ([]ChatV1Member, error) {
	response, errors := c.StreamMember(ServiceSid, ChannelSid, params)

	records := make([]ChatV1Member, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Member records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamMember(ServiceSid string, ChannelSid string, params *ListMemberParams) (chan ChatV1Member, chan error) {
	if params == nil {
		params = &ListMemberParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ChatV1Member, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageMember(ServiceSid, ChannelSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamMember(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamMember(response *ListMemberResponse, params *ListMemberParams, recordChannel chan ChatV1Member, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Members
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListMemberResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListMemberResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListMemberResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMemberResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateMember'
type UpdateMemberParams struct {
	// The index of the last [Message](https://www.twilio.com/docs/api/chat/rest/messages) that the Member has read within the [Channel](https://www.twilio.com/docs/api/chat/rest/channels).
	LastConsumedMessageIndex *int `json:"LastConsumedMessageIndex,omitempty"`
	// The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) to assign to the member. The default roles are those specified on the [Service](https://www.twilio.com/docs/chat/api/services).
	RoleSid *string `json:"RoleSid,omitempty"`
}

func (params *UpdateMemberParams) SetLastConsumedMessageIndex(LastConsumedMessageIndex int) *UpdateMemberParams {
	params.LastConsumedMessageIndex = &LastConsumedMessageIndex
	return params
}
func (params *UpdateMemberParams) SetRoleSid(RoleSid string) *UpdateMemberParams {
	params.RoleSid = &RoleSid
	return params
}

func (c *ApiService) UpdateMember(ServiceSid string, ChannelSid string, Sid string, params *UpdateMemberParams) (*ChatV1Member, error) {
	path := "/v1/Services/{ServiceSid}/Channels/{ChannelSid}/Members/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.LastConsumedMessageIndex != nil {
		data.Set("LastConsumedMessageIndex", fmt.Sprint(*params.LastConsumedMessageIndex))
	}
	if params != nil && params.RoleSid != nil {
		data.Set("RoleSid", *params.RoleSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ChatV1Member{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
