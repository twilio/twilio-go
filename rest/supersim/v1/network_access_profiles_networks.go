/*
 * Twilio - Supersim
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

// Optional parameters for the method 'CreateNetworkAccessProfileNetwork'
type CreateNetworkAccessProfileNetworkParams struct {
	// The SID of the Network resource to be added to the Network Access Profile resource.
	Network *string `json:"Network,omitempty"`
}

func (params *CreateNetworkAccessProfileNetworkParams) SetNetwork(Network string) *CreateNetworkAccessProfileNetworkParams {
	params.Network = &Network
	return params
}

// Add a Network resource to the Network Access Profile resource.
func (c *ApiService) CreateNetworkAccessProfileNetwork(NetworkAccessProfileSid string, params *CreateNetworkAccessProfileNetworkParams) (*SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork, error) {
	path := "/v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks"
	path = strings.Replace(path, "{"+"NetworkAccessProfileSid"+"}", NetworkAccessProfileSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Network != nil {
		data.Set("Network", *params.Network)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove a Network resource from the Network Access Profile resource&#39;s.
func (c *ApiService) DeleteNetworkAccessProfileNetwork(NetworkAccessProfileSid string, Sid string) error {
	path := "/v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks/{Sid}"
	path = strings.Replace(path, "{"+"NetworkAccessProfileSid"+"}", NetworkAccessProfileSid, -1)
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

// Fetch a Network Access Profile resource&#39;s Network resource.
func (c *ApiService) FetchNetworkAccessProfileNetwork(NetworkAccessProfileSid string, Sid string) (*SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork, error) {
	path := "/v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks/{Sid}"
	path = strings.Replace(path, "{"+"NetworkAccessProfileSid"+"}", NetworkAccessProfileSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListNetworkAccessProfileNetwork'
type ListNetworkAccessProfileNetworkParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListNetworkAccessProfileNetworkParams) SetPageSize(PageSize int) *ListNetworkAccessProfileNetworkParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListNetworkAccessProfileNetworkParams) SetLimit(Limit int) *ListNetworkAccessProfileNetworkParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of NetworkAccessProfileNetwork records from the API. Request is executed immediately.
func (c *ApiService) PageNetworkAccessProfileNetwork(NetworkAccessProfileSid string, params *ListNetworkAccessProfileNetworkParams, pageToken string, pageNumber string) (*ListNetworkAccessProfileNetworkResponse, error) {
	path := "/v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks"

	path = strings.Replace(path, "{"+"NetworkAccessProfileSid"+"}", NetworkAccessProfileSid, -1)

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

	ps := &ListNetworkAccessProfileNetworkResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists NetworkAccessProfileNetwork records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListNetworkAccessProfileNetwork(NetworkAccessProfileSid string, params *ListNetworkAccessProfileNetworkParams) ([]SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork, error) {
	if params == nil {
		params = &ListNetworkAccessProfileNetworkParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageNetworkAccessProfileNetwork(NetworkAccessProfileSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork

	for response != nil {
		records = append(records, response.Networks...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListNetworkAccessProfileNetworkResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListNetworkAccessProfileNetworkResponse)
	}

	return records, err
}

// Streams NetworkAccessProfileNetwork records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamNetworkAccessProfileNetwork(NetworkAccessProfileSid string, params *ListNetworkAccessProfileNetworkParams) (chan SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork, error) {
	if params == nil {
		params = &ListNetworkAccessProfileNetworkParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageNetworkAccessProfileNetwork(NetworkAccessProfileSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork, 1)

	go func() {
		for response != nil {
			for item := range response.Networks {
				channel <- response.Networks[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListNetworkAccessProfileNetworkResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListNetworkAccessProfileNetworkResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListNetworkAccessProfileNetworkResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListNetworkAccessProfileNetworkResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
