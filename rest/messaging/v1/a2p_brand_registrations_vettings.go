/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Messaging
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

// Optional parameters for the method 'CreateBrandVetting'
type CreateBrandVettingParams struct {
	//
	VettingProvider *string `json:"VettingProvider,omitempty"`
	// The unique ID of the vetting
	VettingId *string `json:"VettingId,omitempty"`
}

func (params *CreateBrandVettingParams) SetVettingProvider(VettingProvider string) *CreateBrandVettingParams {
	params.VettingProvider = &VettingProvider
	return params
}
func (params *CreateBrandVettingParams) SetVettingId(VettingId string) *CreateBrandVettingParams {
	params.VettingId = &VettingId
	return params
}

func (c *ApiService) CreateBrandVetting(BrandSid string, params *CreateBrandVettingParams) (*MessagingV1BrandVetting, error) {
	path := "/v1/a2p/BrandRegistrations/{BrandSid}/Vettings"
	path = strings.Replace(path, "{"+"BrandSid"+"}", BrandSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.VettingProvider != nil {
		data.Set("VettingProvider", *params.VettingProvider)
	}
	if params != nil && params.VettingId != nil {
		data.Set("VettingId", *params.VettingId)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1BrandVetting{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) FetchBrandVetting(BrandSid string, BrandVettingSid string) (*MessagingV1BrandVetting, error) {
	path := "/v1/a2p/BrandRegistrations/{BrandSid}/Vettings/{BrandVettingSid}"
	path = strings.Replace(path, "{"+"BrandSid"+"}", BrandSid, -1)
	path = strings.Replace(path, "{"+"BrandVettingSid"+"}", BrandVettingSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1BrandVetting{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListBrandVetting'
type ListBrandVettingParams struct {
	// The third-party provider of the vettings to read
	VettingProvider *string `json:"VettingProvider,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int64 `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int64 `json:"limit,omitempty"`
}

func (params *ListBrandVettingParams) SetVettingProvider(VettingProvider string) *ListBrandVettingParams {
	params.VettingProvider = &VettingProvider
	return params
}
func (params *ListBrandVettingParams) SetPageSize(PageSize int64) *ListBrandVettingParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListBrandVettingParams) SetLimit(Limit int64) *ListBrandVettingParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of BrandVetting records from the API. Request is executed immediately.
func (c *ApiService) PageBrandVetting(BrandSid string, params *ListBrandVettingParams, pageToken, pageNumber string) (*ListBrandVettingResponse, error) {
	path := "/v1/a2p/BrandRegistrations/{BrandSid}/Vettings"

	path = strings.Replace(path, "{"+"BrandSid"+"}", BrandSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.VettingProvider != nil {
		data.Set("VettingProvider", *params.VettingProvider)
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

	ps := &ListBrandVettingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists BrandVetting records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListBrandVetting(BrandSid string, params *ListBrandVettingParams) ([]MessagingV1BrandVetting, error) {
	response, errors := c.StreamBrandVetting(BrandSid, params)

	records := make([]MessagingV1BrandVetting, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams BrandVetting records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamBrandVetting(BrandSid string, params *ListBrandVettingParams) (chan MessagingV1BrandVetting, chan error) {
	if params == nil {
		params = &ListBrandVettingParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan MessagingV1BrandVetting, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageBrandVetting(BrandSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamBrandVetting(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamBrandVetting(response *ListBrandVettingResponse, params *ListBrandVettingParams, recordChannel chan MessagingV1BrandVetting, errorChannel chan error) {
	var curRecord int64 = 1

	for response != nil {
		responseRecords := response.Data
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListBrandVettingResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListBrandVettingResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListBrandVettingResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListBrandVettingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
