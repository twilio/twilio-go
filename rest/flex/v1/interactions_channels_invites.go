/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Flex
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateInteractionChannelInvite'
type CreateInteractionChannelInviteParams struct {
	// The Interaction's routing logic.
	Routing *map[string]interface{} `json:"Routing,omitempty"`
}

func (params *CreateInteractionChannelInviteParams) SetRouting(Routing map[string]interface{}) *CreateInteractionChannelInviteParams {
	params.Routing = &Routing
	return params
}

// Invite an Agent or a TaskQueue to a Channel.
func (c *ApiService) CreateInteractionChannelInvite(InteractionSid string, ChannelSid string, params *CreateInteractionChannelInviteParams) (*FlexV1InteractionChannelInvite, error) {
	path := "/v1/Interactions/{InteractionSid}/Channels/{ChannelSid}/Invites"
	path = strings.Replace(path, "{"+"InteractionSid"+"}", InteractionSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Routing != nil {
		v, err := json.Marshal(params.Routing)

		if err != nil {
			return nil, err
		}

		data.Set("Routing", string(v))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1InteractionChannelInvite{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListInteractionChannelInvite'
type ListInteractionChannelInviteParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListInteractionChannelInviteParams) SetPageSize(PageSize int) *ListInteractionChannelInviteParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListInteractionChannelInviteParams) SetLimit(Limit int) *ListInteractionChannelInviteParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of InteractionChannelInvite records from the API. Request is executed immediately.
func (c *ApiService) PageInteractionChannelInvite(InteractionSid string, ChannelSid string, params *ListInteractionChannelInviteParams, pageToken, pageNumber string) (*ListInteractionChannelInviteResponse, error) {
	path := "/v1/Interactions/{InteractionSid}/Channels/{ChannelSid}/Invites"

	path = strings.Replace(path, "{"+"InteractionSid"+"}", InteractionSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
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

	ps := &ListInteractionChannelInviteResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists InteractionChannelInvite records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListInteractionChannelInvite(InteractionSid string, ChannelSid string, params *ListInteractionChannelInviteParams) ([]FlexV1InteractionChannelInvite, error) {
	response, errors := c.StreamInteractionChannelInvite(InteractionSid, ChannelSid, params)

	records := make([]FlexV1InteractionChannelInvite, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams InteractionChannelInvite records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamInteractionChannelInvite(InteractionSid string, ChannelSid string, params *ListInteractionChannelInviteParams) (chan FlexV1InteractionChannelInvite, chan error) {
	if params == nil {
		params = &ListInteractionChannelInviteParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan FlexV1InteractionChannelInvite, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageInteractionChannelInvite(InteractionSid, ChannelSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamInteractionChannelInvite(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamInteractionChannelInvite(response *ListInteractionChannelInviteResponse, params *ListInteractionChannelInviteParams, recordChannel chan FlexV1InteractionChannelInvite, errorChannel chan error) {
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

		record, err := client.GetNext(c.baseURL, response, c.getNextListInteractionChannelInviteResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListInteractionChannelInviteResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListInteractionChannelInviteResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListInteractionChannelInviteResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
