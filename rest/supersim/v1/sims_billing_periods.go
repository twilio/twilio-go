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

// Optional parameters for the method 'ListBillingPeriod'
type ListBillingPeriodParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListBillingPeriodParams) SetPageSize(PageSize int) *ListBillingPeriodParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of Billing Periods for a Super SIM.
func (c *ApiService) ListBillingPeriod(SimSid string, params *ListBillingPeriodParams) (*ListBillingPeriodResponse, error) {
	path := "/v1/Sims/{SimSid}/BillingPeriods"
	path = strings.Replace(path, "{"+"SimSid"+"}", SimSid, -1)

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

	ps := &ListBillingPeriodResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of BillingPeriod records from the API. Request is executed immediately.
func (c *ApiService) BillingPeriodPage(SimSid string, params *ListBillingPeriodParams, pageToken string, pageNumber string, pageSize string) *client.Page {
	path := "/v1/Sims/{SimSid}/BillingPeriods"
	path = strings.Replace(path, "{"+"SimSid"+"}", SimSid, -1)

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

//Streams BillingPeriod records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) BillingPeriodStream(SimSid string, params *ListBillingPeriodParams, meta client.PaginationData) chan map[string]interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.BillingPeriodPage(SimSid, params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.Stream(page, limits.Limit, limits.PageLimit)
}

//Lists BillingPeriod records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) BillingPeriodList(SimSid string, params *ListBillingPeriodParams, meta client.PaginationData) []interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.BillingPeriodPage(SimSid, params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.List(page, limits.Limit, limits.PageLimit)
}
