/*
 * Twilio - Studio
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

// Optional parameters for the method 'CreateEngagement'
type CreateEngagementParams struct {
	// The Twilio phone number to send messages or initiate calls from during the Flow Engagement. Available as variable `{{flow.channel.address}}`
	From *string `json:"From,omitempty"`
	// A JSON string we will add to your flow's context and that you can access as variables inside your flow. For example, if you pass in `Parameters={'name':'Zeke'}` then inside a widget you can reference the variable `{{flow.data.name}}` which will return the string 'Zeke'. Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode your JSON string.
	Parameters *interface{} `json:"Parameters,omitempty"`
	// The Contact phone number to start a Studio Flow Engagement, available as variable `{{contact.channel.address}}`.
	To *string `json:"To,omitempty"`
}

func (params *CreateEngagementParams) SetFrom(From string) *CreateEngagementParams {
	params.From = &From
	return params
}
func (params *CreateEngagementParams) SetParameters(Parameters interface{}) *CreateEngagementParams {
	params.Parameters = &Parameters
	return params
}
func (params *CreateEngagementParams) SetTo(To string) *CreateEngagementParams {
	params.To = &To
	return params
}

// Triggers a new Engagement for the Flow
func (c *ApiService) CreateEngagement(FlowSid string, params *CreateEngagementParams) (*StudioV1Engagement, error) {
	path := "/v1/Flows/{FlowSid}/Engagements"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.Parameters != nil {
		v, err := json.Marshal(params.Parameters)

		if err != nil {
			return nil, err
		}

		data.Set("Parameters", string(v))
	}
	if params != nil && params.To != nil {
		data.Set("To", *params.To)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV1Engagement{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete this Engagement and all Steps relating to it.
func (c *ApiService) DeleteEngagement(FlowSid string, Sid string) error {
	path := "/v1/Flows/{FlowSid}/Engagements/{Sid}"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
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

// Retrieve an Engagement
func (c *ApiService) FetchEngagement(FlowSid string, Sid string) (*StudioV1Engagement, error) {
	path := "/v1/Flows/{FlowSid}/Engagements/{Sid}"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV1Engagement{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListEngagement'
type ListEngagementParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListEngagementParams) SetPageSize(PageSize int) *ListEngagementParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListEngagementParams) SetLimit(Limit int) *ListEngagementParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Engagement records from the API. Request is executed immediately.
func (c *ApiService) PageEngagement(FlowSid string, params *ListEngagementParams, pageToken, pageNumber string) (*ListEngagementResponse, error) {
	path := "/v1/Flows/{FlowSid}/Engagements"

	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)

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

	ps := &ListEngagementResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Engagement records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListEngagement(FlowSid string, params *ListEngagementParams) ([]StudioV1Engagement, error) {
	response, errors := c.StreamEngagement(FlowSid, params)

	records := make([]StudioV1Engagement, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Engagement records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamEngagement(FlowSid string, params *ListEngagementParams) (chan StudioV1Engagement, chan error) {
	if params == nil {
		params = &ListEngagementParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan StudioV1Engagement, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageEngagement(FlowSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamEngagement(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamEngagement(response *ListEngagementResponse, params *ListEngagementParams, recordChannel chan StudioV1Engagement, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Engagements
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListEngagementResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListEngagementResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListEngagementResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListEngagementResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
