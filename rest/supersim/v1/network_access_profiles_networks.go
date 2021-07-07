/*
 * Twilio - Supersim
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
}

func (params *ListNetworkAccessProfileNetworkParams) SetPageSize(PageSize int) *ListNetworkAccessProfileNetworkParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of Network Access Profile resource&#39;s Network resource.
func (c *ApiService) ListNetworkAccessProfileNetwork(NetworkAccessProfileSid string, params *ListNetworkAccessProfileNetworkParams) (*ListNetworkAccessProfileNetworkResponse, error) {
	path := "/v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks"
	path = strings.Replace(path, "{"+"NetworkAccessProfileSid"+"}", NetworkAccessProfileSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
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

//Retrieve a single page of  records from the API. Request is executed immediately.
func (c *ApiService) NetworkAccessProfilesNetworksPage(NetworkAccessProfileSid string, params *ListNetworkAccessProfileNetworkParams, pageToken string, pageNumber string) (*ListNetworkAccessProfileNetworkResponse, error) {
	path := "/v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks"
	path = strings.Replace(path, "{"+"NetworkAccessProfileSid"+"}", NetworkAccessProfileSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	data.Set("PageToken", pageToken)
	data.Set("PageNumber", pageNumber)

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

//Lists NetworkAccessProfilesNetworks records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) NetworkAccessProfilesNetworksList(NetworkAccessProfileSid string, params *ListNetworkAccessProfileNetworkParams, limit int) ([]ListNetworkAccessProfileNetworkResponse, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	response, err := c.ListNetworkAccessProfileNetwork(NetworkAccessProfileSid, params)
	if err != nil {
		return nil, err
	}

	page := client.NewPage(c.baseURL, response)

	resp := c.requestHandler.List(page, limit, 0)
	ret := make([]ListNetworkAccessProfileNetworkResponse, len(resp))

	for i := range resp {
		jsonStr, _ := json.Marshal(resp[i])
		ps := ListNetworkAccessProfileNetworkResponse{}
		if err := json.Unmarshal(jsonStr, &ps); err != nil {
			return ret, err
		}

		ret[i] = ps
	}

	return ret, nil
}

//Streams NetworkAccessProfilesNetworks records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) NetworkAccessProfilesNetworksStream(NetworkAccessProfileSid string, params *ListNetworkAccessProfileNetworkParams, limit int) (chan interface{}, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	response, err := c.ListNetworkAccessProfileNetwork(NetworkAccessProfileSid, params)
	if err != nil {
		return nil, err
	}

	page := client.NewPage(c.baseURL, response)

	ps := ListNetworkAccessProfileNetworkResponse{}
	return c.requestHandler.Stream(page, limit, 0, ps), nil
}
