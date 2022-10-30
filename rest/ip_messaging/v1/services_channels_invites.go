/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Ip_messaging
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

// Optional parameters for the method 'CreateInvite'
type CreateInviteParams struct {
	//
	Identity *string `json:"Identity,omitempty"`
	//
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

//
func (c *ApiService) CreateInvite(ServiceSid string, ChannelSid string, params *CreateInviteParams) (*IpMessagingV1Invite, error) {
	return c.CreateInviteWithCtx(context.TODO(), ServiceSid, ChannelSid, params)
}

//
func (c *ApiService) CreateInviteWithCtx(ctx context.Context, ServiceSid string, ChannelSid string, params *CreateInviteParams) (*IpMessagingV1Invite, error) {
	path := "/v1/Services/{ServiceSid}/Channels/{ChannelSid}/Invites"
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

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV1Invite{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//
func (c *ApiService) DeleteInvite(ServiceSid string, ChannelSid string, Sid string) error {
	return c.DeleteInviteWithCtx(context.TODO(), ServiceSid, ChannelSid, Sid)
}

//
func (c *ApiService) DeleteInviteWithCtx(ctx context.Context, ServiceSid string, ChannelSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Channels/{ChannelSid}/Invites/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)
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
func (c *ApiService) FetchInvite(ServiceSid string, ChannelSid string, Sid string) (*IpMessagingV1Invite, error) {
	return c.FetchInviteWithCtx(context.TODO(), ServiceSid, ChannelSid, Sid)
}

//
func (c *ApiService) FetchInviteWithCtx(ctx context.Context, ServiceSid string, ChannelSid string, Sid string) (*IpMessagingV1Invite, error) {
	path := "/v1/Services/{ServiceSid}/Channels/{ChannelSid}/Invites/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV1Invite{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListInvite'
type ListInviteParams struct {
	//
	Identity *[]string `json:"Identity,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListInviteParams) SetIdentity(Identity []string) *ListInviteParams {
	params.Identity = &Identity
	return params
}
func (params *ListInviteParams) SetPageSize(PageSize int) *ListInviteParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListInviteParams) SetLimit(Limit int) *ListInviteParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Invite records from the API. Request is executed immediately.
func (c *ApiService) PageInvite(ServiceSid string, ChannelSid string, params *ListInviteParams, pageToken, pageNumber string) (*ListInviteResponse, error) {
	return c.PageInviteWithCtx(context.TODO(), ServiceSid, ChannelSid, params, pageToken, pageNumber)
}

// Retrieve a single page of Invite records from the API. Request is executed immediately.
func (c *ApiService) PageInviteWithCtx(ctx context.Context, ServiceSid string, ChannelSid string, params *ListInviteParams, pageToken, pageNumber string) (*ListInviteResponse, error) {
	path := "/v1/Services/{ServiceSid}/Channels/{ChannelSid}/Invites"

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

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
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
func (c *ApiService) ListInvite(ServiceSid string, ChannelSid string, params *ListInviteParams) ([]IpMessagingV1Invite, error) {
	return c.ListInviteWithCtx(context.TODO(), ServiceSid, ChannelSid, params)
}

// Lists Invite records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListInviteWithCtx(ctx context.Context, ServiceSid string, ChannelSid string, params *ListInviteParams) ([]IpMessagingV1Invite, error) {
	response, errors := c.StreamInviteWithCtx(ctx, ServiceSid, ChannelSid, params)

	records := make([]IpMessagingV1Invite, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Invite records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamInvite(ServiceSid string, ChannelSid string, params *ListInviteParams) (chan IpMessagingV1Invite, chan error) {
	return c.StreamInviteWithCtx(context.TODO(), ServiceSid, ChannelSid, params)
}

// Streams Invite records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamInviteWithCtx(ctx context.Context, ServiceSid string, ChannelSid string, params *ListInviteParams) (chan IpMessagingV1Invite, chan error) {
	if params == nil {
		params = &ListInviteParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan IpMessagingV1Invite, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageInviteWithCtx(ctx, ServiceSid, ChannelSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamInvite(ctx, response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamInvite(ctx context.Context, response *ListInviteResponse, params *ListInviteParams, recordChannel chan IpMessagingV1Invite, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Invites
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNextWithCtx(ctx, c.baseURL, response, c.getNextListInviteResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListInviteResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListInviteResponse(ctx context.Context, nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(ctx, nextPageUrl, nil, nil)
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
