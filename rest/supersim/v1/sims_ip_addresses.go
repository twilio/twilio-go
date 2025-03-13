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

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'ListSimIpAddress'
type ListSimIpAddressParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSimIpAddressParams) SetPageSize(PageSize int) *ListSimIpAddressParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSimIpAddressParams) SetLimit(Limit int) *ListSimIpAddressParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SimIpAddress records from the API. Request is executed immediately.
func (c *ApiService) PageSimIpAddress(SimSid string, params *ListSimIpAddressParams, pageToken, pageNumber string) (*ListSimIpAddressResponse, error) {
	path := "/v1/Sims/{SimSid}/IpAddresses"

	path = strings.Replace(path, "{"+"SimSid"+"}", SimSid, -1)

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

	ps := &ListSimIpAddressResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SimIpAddress records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSimIpAddress(SimSid string, params *ListSimIpAddressParams) ([]SupersimV1SimIpAddress, error) {
	response, errors := c.StreamSimIpAddress(SimSid, params)

	records := make([]SupersimV1SimIpAddress, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams SimIpAddress records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSimIpAddress(SimSid string, params *ListSimIpAddressParams) (chan SupersimV1SimIpAddress, chan error) {
	if params == nil {
		params = &ListSimIpAddressParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan SupersimV1SimIpAddress, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageSimIpAddress(SimSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamSimIpAddress(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamSimIpAddress(response *ListSimIpAddressResponse, params *ListSimIpAddressParams, recordChannel chan SupersimV1SimIpAddress, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.IpAddresses
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListSimIpAddressResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListSimIpAddressResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListSimIpAddressResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSimIpAddressResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
