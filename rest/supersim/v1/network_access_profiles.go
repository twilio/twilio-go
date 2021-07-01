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

// Optional parameters for the method 'CreateNetworkAccessProfile'
type CreateNetworkAccessProfileParams struct {
	// List of Network SIDs that this Network Access Profile will allow connections to.
	Networks *[]string `json:"Networks,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *CreateNetworkAccessProfileParams) SetNetworks(Networks []string) *CreateNetworkAccessProfileParams {
	params.Networks = &Networks
	return params
}
func (params *CreateNetworkAccessProfileParams) SetUniqueName(UniqueName string) *CreateNetworkAccessProfileParams {
	params.UniqueName = &UniqueName
	return params
}

// Create a new Network Access Profile
func (c *ApiService) CreateNetworkAccessProfile(params *CreateNetworkAccessProfileParams) (*SupersimV1NetworkAccessProfile, error) {
	path := "/v1/NetworkAccessProfiles"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Networks != nil {
		for _, item := range *params.Networks {
			data.Add("Networks", item)
		}
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1NetworkAccessProfile{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Fetch a Network Access Profile instance from your account.
func (c *ApiService) FetchNetworkAccessProfile(Sid string) (*SupersimV1NetworkAccessProfile, error) {
	path := "/v1/NetworkAccessProfiles/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1NetworkAccessProfile{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListNetworkAccessProfile'
type ListNetworkAccessProfileParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListNetworkAccessProfileParams) SetPageSize(PageSize int) *ListNetworkAccessProfileParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of Network Access Profiles from your account.
func (c *ApiService) ListNetworkAccessProfile(params *ListNetworkAccessProfileParams) (*ListNetworkAccessProfileResponse, error) {
	path := "/v1/NetworkAccessProfiles"

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

	ps := &ListNetworkAccessProfileResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of NetworkAccessProfile records from the API. Request is executed immediately.
func (c *ApiService) NetworkAccessProfilePage(params *ListNetworkAccessProfileParams, pageToken string, pageNumber string, pageSize string) *client.Page {
	path := "/v1/NetworkAccessProfiles"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	data.Set("PageToken", pageToken)
	data.Set("PageNumber", pageNumber)
	data.Set("PageSize", pageSize)

	response, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil
	}

	return client.NewPage(c.baseURL, response)
}

//Streams NetworkAccessProfile records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) NetworkAccessProfileStream(params *ListNetworkAccessProfileParams, meta client.PaginationData) chan map[string]interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.NetworkAccessProfilePage(params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.Stream(page, limits.Limit, limits.PageLimit)
}

//Lists NetworkAccessProfile records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) NetworkAccessProfileList(params *ListNetworkAccessProfileParams, meta client.PaginationData) []interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.NetworkAccessProfilePage(params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.List(page, limits.Limit, limits.PageLimit)
}

// Optional parameters for the method 'UpdateNetworkAccessProfile'
type UpdateNetworkAccessProfileParams struct {
	// The new unique name of the Network Access Profile.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *UpdateNetworkAccessProfileParams) SetUniqueName(UniqueName string) *UpdateNetworkAccessProfileParams {
	params.UniqueName = &UniqueName
	return params
}

// Updates the given properties of a Network Access Profile in your account.
func (c *ApiService) UpdateNetworkAccessProfile(Sid string, params *UpdateNetworkAccessProfileParams) (*SupersimV1NetworkAccessProfile, error) {
	path := "/v1/NetworkAccessProfiles/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1NetworkAccessProfile{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
