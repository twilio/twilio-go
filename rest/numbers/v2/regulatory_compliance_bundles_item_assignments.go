/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
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

// Optional parameters for the method 'CreateItemAssignment'
type CreateItemAssignmentParams struct {
	// The SID of an object bag that holds information of the different items.
	ObjectSid *string `json:"ObjectSid,omitempty"`
}

func (params *CreateItemAssignmentParams) SetObjectSid(ObjectSid string) *CreateItemAssignmentParams {
	params.ObjectSid = &ObjectSid
	return params
}

// Create a new Assigned Item.
func (c *ApiService) CreateItemAssignment(BundleSid string, params *CreateItemAssignmentParams) (*NumbersV2RegulatoryComplianceBundleItemAssignment, error) {
	path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ObjectSid != nil {
		data.Set("ObjectSid", *params.ObjectSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2RegulatoryComplianceBundleItemAssignment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove an Assignment Item Instance.
func (c *ApiService) DeleteItemAssignment(BundleSid string, Sid string) error {
	path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments/{Sid}"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)
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

// Fetch specific Assigned Item Instance.
func (c *ApiService) FetchItemAssignment(BundleSid string, Sid string) (*NumbersV2RegulatoryComplianceBundleItemAssignment, error) {
	path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments/{Sid}"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2RegulatoryComplianceBundleItemAssignment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListItemAssignment'
type ListItemAssignmentParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListItemAssignmentParams) SetPageSize(PageSize int) *ListItemAssignmentParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListItemAssignmentParams) SetLimit(Limit int) *ListItemAssignmentParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ItemAssignment records from the API. Request is executed immediately.
func (c *ApiService) PageItemAssignment(BundleSid string, params *ListItemAssignmentParams, pageToken string, pageNumber string) (*ListItemAssignmentResponse, error) {
	path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments"

	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)

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

	ps := &ListItemAssignmentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ItemAssignment records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListItemAssignment(BundleSid string, params *ListItemAssignmentParams) ([]NumbersV2RegulatoryComplianceBundleItemAssignment, error) {
	if params == nil {
		params = &ListItemAssignmentParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageItemAssignment(BundleSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []NumbersV2RegulatoryComplianceBundleItemAssignment

	for response != nil {
		records = append(records, response.Results...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListItemAssignmentResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListItemAssignmentResponse)
	}

	return records, err
}

// Streams ItemAssignment records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamItemAssignment(BundleSid string, params *ListItemAssignmentParams) (chan NumbersV2RegulatoryComplianceBundleItemAssignment, error) {
	if params == nil {
		params = &ListItemAssignmentParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageItemAssignment(BundleSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan NumbersV2RegulatoryComplianceBundleItemAssignment, 1)

	go func() {
		for response != nil {
			for item := range response.Results {
				channel <- response.Results[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListItemAssignmentResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListItemAssignmentResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListItemAssignmentResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListItemAssignmentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
