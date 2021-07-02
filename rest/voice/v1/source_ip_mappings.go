/*
 * Twilio - Voice
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

// Optional parameters for the method 'CreateSourceIpMapping'
type CreateSourceIpMappingParams struct {
	// The Twilio-provided string that uniquely identifies the IP Record resource to map from.
	IpRecordSid *string `json:"IpRecordSid,omitempty"`
	// The SID of the SIP Domain that the IP Record should be mapped to.
	SipDomainSid *string `json:"SipDomainSid,omitempty"`
}

func (params *CreateSourceIpMappingParams) SetIpRecordSid(IpRecordSid string) *CreateSourceIpMappingParams {
	params.IpRecordSid = &IpRecordSid
	return params
}
func (params *CreateSourceIpMappingParams) SetSipDomainSid(SipDomainSid string) *CreateSourceIpMappingParams {
	params.SipDomainSid = &SipDomainSid
	return params
}

func (c *ApiService) CreateSourceIpMapping(params *CreateSourceIpMappingParams) (*VoiceV1SourceIpMapping, error) {
	path := "/v1/SourceIpMappings"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.IpRecordSid != nil {
		data.Set("IpRecordSid", *params.IpRecordSid)
	}
	if params != nil && params.SipDomainSid != nil {
		data.Set("SipDomainSid", *params.SipDomainSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1SourceIpMapping{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteSourceIpMapping(Sid string) error {
	path := "/v1/SourceIpMappings/{Sid}"
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

func (c *ApiService) FetchSourceIpMapping(Sid string) (*VoiceV1SourceIpMapping, error) {
	path := "/v1/SourceIpMappings/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1SourceIpMapping{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSourceIpMapping'
type ListSourceIpMappingParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListSourceIpMappingParams) SetPageSize(PageSize int) *ListSourceIpMappingParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListSourceIpMapping(params *ListSourceIpMappingParams) (*ListSourceIpMappingResponse, error) {
	path := "/v1/SourceIpMappings"

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

	ps := &ListSourceIpMappingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of SourceIpMapping records from the API. Request is executed immediately.
func (c *ApiService) SourceIpMappingPage(params *ListSourceIpMappingParams, pageToken string, pageNumber string) (*client.Page, error) {
	path := "/v1/SourceIpMappings"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	data.Set("PageToken", pageToken)
	data.Set("PageNumber", pageNumber)

	response, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	return client.NewPage(c.baseURL, response), nil
}

//Streams SourceIpMapping records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) SourceIpMappingStream(params *ListSourceIpMappingParams, limit int) (chan map[string]interface{}, error) {
	if params.PageSize == nil {
		params.SetPageSize(0)
	}
	params.SetPageSize(c.requestHandler.ReadLimits(*params.PageSize, limit))
	page, err := c.SourceIpMappingPage(params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.Stream(page, limit, 0), nil
}

//Lists SourceIpMapping records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) SourceIpMappingList(params *ListSourceIpMappingParams, limit int) ([]interface{}, error) {
	if params.PageSize == nil {
		params.SetPageSize(0)
	}
	params.SetPageSize(c.requestHandler.ReadLimits(*params.PageSize, limit))
	page, err := c.SourceIpMappingPage(params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.List(page, limit, 0), nil
}

// Optional parameters for the method 'UpdateSourceIpMapping'
type UpdateSourceIpMappingParams struct {
	// The SID of the SIP Domain that the IP Record should be mapped to.
	SipDomainSid *string `json:"SipDomainSid,omitempty"`
}

func (params *UpdateSourceIpMappingParams) SetSipDomainSid(SipDomainSid string) *UpdateSourceIpMappingParams {
	params.SipDomainSid = &SipDomainSid
	return params
}

func (c *ApiService) UpdateSourceIpMapping(Sid string, params *UpdateSourceIpMappingParams) (*VoiceV1SourceIpMapping, error) {
	path := "/v1/SourceIpMappings/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.SipDomainSid != nil {
		data.Set("SipDomainSid", *params.SipDomainSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1SourceIpMapping{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
