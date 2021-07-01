/*
 * Twilio - Trunking
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

// Optional parameters for the method 'CreateIpAccessControlList'
type CreateIpAccessControlListParams struct {
	// The SID of the [IP Access Control List](https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollist-resource) that you want to associate with the trunk.
	IpAccessControlListSid *string `json:"IpAccessControlListSid,omitempty"`
}

func (params *CreateIpAccessControlListParams) SetIpAccessControlListSid(IpAccessControlListSid string) *CreateIpAccessControlListParams {
	params.IpAccessControlListSid = &IpAccessControlListSid
	return params
}

// Associate an IP Access Control List with a Trunk
func (c *ApiService) CreateIpAccessControlList(TrunkSid string, params *CreateIpAccessControlListParams) (*TrunkingV1TrunkIpAccessControlList, error) {
	path := "/v1/Trunks/{TrunkSid}/IpAccessControlLists"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.IpAccessControlListSid != nil {
		data.Set("IpAccessControlListSid", *params.IpAccessControlListSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1TrunkIpAccessControlList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove an associated IP Access Control List from a Trunk
func (c *ApiService) DeleteIpAccessControlList(TrunkSid string, Sid string) error {
	path := "/v1/Trunks/{TrunkSid}/IpAccessControlLists/{Sid}"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)
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

func (c *ApiService) FetchIpAccessControlList(TrunkSid string, Sid string) (*TrunkingV1TrunkIpAccessControlList, error) {
	path := "/v1/Trunks/{TrunkSid}/IpAccessControlLists/{Sid}"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1TrunkIpAccessControlList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListIpAccessControlList'
type ListIpAccessControlListParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListIpAccessControlListParams) SetPageSize(PageSize int) *ListIpAccessControlListParams {
	params.PageSize = &PageSize
	return params
}

// List all IP Access Control Lists for a Trunk
func (c *ApiService) ListIpAccessControlList(TrunkSid string, params *ListIpAccessControlListParams) (*ListIpAccessControlListResponse, error) {
	path := "/v1/Trunks/{TrunkSid}/IpAccessControlLists"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

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

	ps := &ListIpAccessControlListResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of IpAccessControlList records from the API. Request is executed immediately.
func (c *ApiService) IpAccessControlListPage(TrunkSid string, params *ListIpAccessControlListParams, pageToken string, pageNumber string, pageSize string) *client.Page {
	path := "/v1/Trunks/{TrunkSid}/IpAccessControlLists"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

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

//Streams IpAccessControlList records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) IpAccessControlListStream(TrunkSid string, params *ListIpAccessControlListParams, meta client.PaginationData) chan map[string]interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.IpAccessControlListPage(TrunkSid, params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.Stream(page, limits.Limit, limits.PageLimit)
}

//Lists IpAccessControlList records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) IpAccessControlListList(TrunkSid string, params *ListIpAccessControlListParams, meta client.PaginationData) []interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.IpAccessControlListPage(TrunkSid, params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.List(page, limits.Limit, limits.PageLimit)
}
