/*
 * Twilio - Proxy
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

// Delete a specific Interaction.
func (c *ApiService) DeleteInteraction(ServiceSid string, SessionSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Sessions/{SessionSid}/Interactions/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"SessionSid"+"}", SessionSid, -1)
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

// Retrieve a list of Interactions for a given [Session](https://www.twilio.com/docs/proxy/api/session).
func (c *ApiService) FetchInteraction(ServiceSid string, SessionSid string, Sid string) (*ProxyV1ServiceSessionInteraction, error) {
	path := "/v1/Services/{ServiceSid}/Sessions/{SessionSid}/Interactions/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"SessionSid"+"}", SessionSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyV1ServiceSessionInteraction{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListInteraction'
type ListInteractionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListInteractionParams) SetPageSize(PageSize int) *ListInteractionParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all Interactions for a Session. A maximum of 100 records will be returned per page.
func (c *ApiService) ListInteraction(ServiceSid string, SessionSid string, params *ListInteractionParams) (*ListInteractionResponse, error) {
	path := "/v1/Services/{ServiceSid}/Sessions/{SessionSid}/Interactions"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"SessionSid"+"}", SessionSid, -1)

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

	ps := &ListInteractionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of Interaction records from the API. Request is executed immediately.
func (c *ApiService) InteractionPage(ServiceSid string, SessionSid string, params *ListInteractionParams, pageToken string, pageNumber string, pageSize string) *client.Page {
	path := "/v1/Services/{ServiceSid}/Sessions/{SessionSid}/Interactions"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"SessionSid"+"}", SessionSid, -1)

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

//Streams Interaction records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) InteractionStream(ServiceSid string, SessionSid string, params *ListInteractionParams, meta client.PaginationData) chan map[string]interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.InteractionPage(ServiceSid, SessionSid, params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.Stream(page, limits.Limit, limits.PageLimit)
}

//Lists Interaction records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) InteractionList(ServiceSid string, SessionSid string, params *ListInteractionParams, meta client.PaginationData) []interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.InteractionPage(ServiceSid, SessionSid, params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.List(page, limits.Limit, limits.PageLimit)
}
