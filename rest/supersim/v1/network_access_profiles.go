/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Supersim
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
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListNetworkAccessProfileParams) SetPageSize(PageSize int) *ListNetworkAccessProfileParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListNetworkAccessProfileParams) SetLimit(Limit int) *ListNetworkAccessProfileParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of NetworkAccessProfile records from the API. Request is executed immediately.
func (c *ApiService) PageNetworkAccessProfile(params *ListNetworkAccessProfileParams, pageToken, pageNumber string) (*ListNetworkAccessProfileResponse, error) {
	path := "/v1/NetworkAccessProfiles"

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

	ps := &ListNetworkAccessProfileResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists NetworkAccessProfile records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListNetworkAccessProfile(params *ListNetworkAccessProfileParams) ([]SupersimV1NetworkAccessProfile, error) {
	response, errors := c.StreamNetworkAccessProfile(params)

	records := make([]SupersimV1NetworkAccessProfile, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams NetworkAccessProfile records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamNetworkAccessProfile(params *ListNetworkAccessProfileParams) (chan SupersimV1NetworkAccessProfile, chan error) {
	if params == nil {
		params = &ListNetworkAccessProfileParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan SupersimV1NetworkAccessProfile, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageNetworkAccessProfile(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamNetworkAccessProfile(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamNetworkAccessProfile(response *ListNetworkAccessProfileResponse, params *ListNetworkAccessProfileParams, recordChannel chan SupersimV1NetworkAccessProfile, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.NetworkAccessProfiles
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListNetworkAccessProfileResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListNetworkAccessProfileResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListNetworkAccessProfileResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListNetworkAccessProfileResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
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
