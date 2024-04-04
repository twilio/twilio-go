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
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreatePlugin'
type CreatePluginParams struct {
	// The Flex-Metadata HTTP request header
	FlexMetadata *string `json:"Flex-Metadata,omitempty"`
	// The Flex Plugin's unique name.
	UniqueName *string `json:"UniqueName,omitempty"`
	// The Flex Plugin's friendly name.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// A descriptive string that you create to describe the plugin resource. It can be up to 500 characters long
	Description *string `json:"Description,omitempty"`
	// The version of Flex Plugins CLI used to create this plugin
	CliVersion *string `json:"CliVersion,omitempty"`
	// The validation status of the plugin, indicating whether it has been validated
	ValidateStatus *string `json:"ValidateStatus,omitempty"`
}

func (params *CreatePluginParams) SetFlexMetadata(FlexMetadata string) *CreatePluginParams {
	params.FlexMetadata = &FlexMetadata
	return params
}
func (params *CreatePluginParams) SetUniqueName(UniqueName string) *CreatePluginParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *CreatePluginParams) SetFriendlyName(FriendlyName string) *CreatePluginParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreatePluginParams) SetDescription(Description string) *CreatePluginParams {
	params.Description = &Description
	return params
}
func (params *CreatePluginParams) SetCliVersion(CliVersion string) *CreatePluginParams {
	params.CliVersion = &CliVersion
	return params
}
func (params *CreatePluginParams) SetValidateStatus(ValidateStatus string) *CreatePluginParams {
	params.ValidateStatus = &ValidateStatus
	return params
}

func (c *ApiService) CreatePlugin(params *CreatePluginParams) (*FlexV1Plugin, error) {
	path := "/v1/PluginService/Plugins"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}
	if params != nil && params.CliVersion != nil {
		data.Set("CliVersion", *params.CliVersion)
	}
	if params != nil && params.ValidateStatus != nil {
		data.Set("ValidateStatus", *params.ValidateStatus)
	}

	if params != nil && params.FlexMetadata != nil {
		headers["Flex-Metadata"] = *params.FlexMetadata
	}
	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1Plugin{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'FetchPlugin'
type FetchPluginParams struct {
	// The Flex-Metadata HTTP request header
	FlexMetadata *string `json:"Flex-Metadata,omitempty"`
}

func (params *FetchPluginParams) SetFlexMetadata(FlexMetadata string) *FetchPluginParams {
	params.FlexMetadata = &FlexMetadata
	return params
}

func (c *ApiService) FetchPlugin(Sid string, params *FetchPluginParams) (*FlexV1Plugin, error) {
	path := "/v1/PluginService/Plugins/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FlexMetadata != nil {
		headers["Flex-Metadata"] = *params.FlexMetadata
	}
	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1Plugin{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListPlugin'
type ListPluginParams struct {
	// The Flex-Metadata HTTP request header
	FlexMetadata *string `json:"Flex-Metadata,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListPluginParams) SetFlexMetadata(FlexMetadata string) *ListPluginParams {
	params.FlexMetadata = &FlexMetadata
	return params
}
func (params *ListPluginParams) SetPageSize(PageSize int) *ListPluginParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListPluginParams) SetLimit(Limit int) *ListPluginParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Plugin records from the API. Request is executed immediately.
func (c *ApiService) PagePlugin(params *ListPluginParams, pageToken, pageNumber string) (*ListPluginResponse, error) {
	path := "/v1/PluginService/Plugins"

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

	ps := &ListPluginResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Plugin records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListPlugin(params *ListPluginParams) ([]FlexV1Plugin, error) {
	response, errors := c.StreamPlugin(params)

	records := make([]FlexV1Plugin, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Plugin records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamPlugin(params *ListPluginParams) (chan FlexV1Plugin, chan error) {
	if params == nil {
		params = &ListPluginParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan FlexV1Plugin, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PagePlugin(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamPlugin(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamPlugin(response *ListPluginResponse, params *ListPluginParams, recordChannel chan FlexV1Plugin, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Plugins
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListPluginResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListPluginResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListPluginResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListPluginResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdatePlugin'
type UpdatePluginParams struct {
	// The Flex-Metadata HTTP request header
	FlexMetadata *string `json:"Flex-Metadata,omitempty"`
	// The Flex Plugin's friendly name.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// A descriptive string that you update to describe the plugin resource. It can be up to 500 characters long
	Description *string `json:"Description,omitempty"`
}

func (params *UpdatePluginParams) SetFlexMetadata(FlexMetadata string) *UpdatePluginParams {
	params.FlexMetadata = &FlexMetadata
	return params
}
func (params *UpdatePluginParams) SetFriendlyName(FriendlyName string) *UpdatePluginParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdatePluginParams) SetDescription(Description string) *UpdatePluginParams {
	params.Description = &Description
	return params
}

func (c *ApiService) UpdatePlugin(Sid string, params *UpdatePluginParams) (*FlexV1Plugin, error) {
	path := "/v1/PluginService/Plugins/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}

	if params != nil && params.FlexMetadata != nil {
		headers["Flex-Metadata"] = *params.FlexMetadata
	}
	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1Plugin{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
