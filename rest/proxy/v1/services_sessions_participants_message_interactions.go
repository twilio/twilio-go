/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.1
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

// Optional parameters for the method 'CreateMessageInteraction'
type CreateMessageInteractionParams struct {
	// The message to send to the participant
	Body *string `json:"Body,omitempty"`
	// Reserved. Not currently supported.
	MediaUrl *[]string `json:"MediaUrl,omitempty"`
}

func (params *CreateMessageInteractionParams) SetBody(Body string) *CreateMessageInteractionParams {
	params.Body = &Body
	return params
}
func (params *CreateMessageInteractionParams) SetMediaUrl(MediaUrl []string) *CreateMessageInteractionParams {
	params.MediaUrl = &MediaUrl
	return params
}

// Create a new message Interaction to send directly from your system to one [Participant](https://www.twilio.com/docs/proxy/api/participant).  The &#x60;inbound&#x60; properties for the Interaction will always be empty.
func (c *ApiService) CreateMessageInteraction(ServiceSid string, SessionSid string, ParticipantSid string, params *CreateMessageInteractionParams) (*ProxyV1ServiceSessionParticipantMessageInteraction, error) {
	path := "/v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants/{ParticipantSid}/MessageInteractions"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"SessionSid"+"}", SessionSid, -1)
	path = strings.Replace(path, "{"+"ParticipantSid"+"}", ParticipantSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Body != nil {
		data.Set("Body", *params.Body)
	}
	if params != nil && params.MediaUrl != nil {
		for _, item := range *params.MediaUrl {
			data.Add("MediaUrl", item)
		}
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyV1ServiceSessionParticipantMessageInteraction{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) FetchMessageInteraction(ServiceSid string, SessionSid string, ParticipantSid string, Sid string) (*ProxyV1ServiceSessionParticipantMessageInteraction, error) {
	path := "/v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants/{ParticipantSid}/MessageInteractions/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"SessionSid"+"}", SessionSid, -1)
	path = strings.Replace(path, "{"+"ParticipantSid"+"}", ParticipantSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyV1ServiceSessionParticipantMessageInteraction{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListMessageInteraction'
type ListMessageInteractionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListMessageInteractionParams) SetPageSize(PageSize int) *ListMessageInteractionParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListMessageInteractionParams) SetLimit(Limit int) *ListMessageInteractionParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of MessageInteraction records from the API. Request is executed immediately.
func (c *ApiService) PageMessageInteraction(ServiceSid string, SessionSid string, ParticipantSid string, params *ListMessageInteractionParams, pageToken string, pageNumber string) (*ListMessageInteractionResponse, error) {
	path := "/v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants/{ParticipantSid}/MessageInteractions"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"SessionSid"+"}", SessionSid, -1)
	path = strings.Replace(path, "{"+"ParticipantSid"+"}", ParticipantSid, -1)

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

	ps := &ListMessageInteractionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists MessageInteraction records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListMessageInteraction(ServiceSid string, SessionSid string, ParticipantSid string, params *ListMessageInteractionParams) ([]ProxyV1ServiceSessionParticipantMessageInteraction, error) {
	if params == nil {
		params = &ListMessageInteractionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageMessageInteraction(ServiceSid, SessionSid, ParticipantSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ProxyV1ServiceSessionParticipantMessageInteraction

	for response != nil {
		records = append(records, response.Interactions...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListMessageInteractionResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListMessageInteractionResponse)
	}

	return records, err
}

// Streams MessageInteraction records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamMessageInteraction(ServiceSid string, SessionSid string, ParticipantSid string, params *ListMessageInteractionParams) (chan ProxyV1ServiceSessionParticipantMessageInteraction, error) {
	if params == nil {
		params = &ListMessageInteractionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageMessageInteraction(ServiceSid, SessionSid, ParticipantSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ProxyV1ServiceSessionParticipantMessageInteraction, 1)

	go func() {
		for response != nil {
			for item := range response.Interactions {
				channel <- response.Interactions[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListMessageInteractionResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListMessageInteractionResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListMessageInteractionResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMessageInteractionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
