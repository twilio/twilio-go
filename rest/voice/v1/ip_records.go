/*
 * Twilio - Voice
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
)

// Optional parameters for the method 'CreateIpRecord'
type CreateIpRecordParams struct {
	// An integer representing the length of the [CIDR](https://tools.ietf.org/html/rfc4632) prefix to use with this IP address. By default the entire IP address is used, which for IPv4 is value 32.
	CidrPrefixLength *int `json:"CidrPrefixLength,omitempty"`
	// A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// An IP address in dotted decimal notation, IPv4 only.
	IpAddress *string `json:"IpAddress,omitempty"`
}

func (params *CreateIpRecordParams) SetCidrPrefixLength(CidrPrefixLength int) *CreateIpRecordParams {
	params.CidrPrefixLength = &CidrPrefixLength
	return params
}
func (params *CreateIpRecordParams) SetFriendlyName(FriendlyName string) *CreateIpRecordParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateIpRecordParams) SetIpAddress(IpAddress string) *CreateIpRecordParams {
	params.IpAddress = &IpAddress
	return params
}

func (c *ApiService) CreateIpRecord(params *CreateIpRecordParams) (*VoiceV1IpRecord, error) {
	path := "/v1/IpRecords"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CidrPrefixLength != nil {
		data.Set("CidrPrefixLength", fmt.Sprint(*params.CidrPrefixLength))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.IpAddress != nil {
		data.Set("IpAddress", *params.IpAddress)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1IpRecord{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteIpRecord(Sid string) error {
	path := "/v1/IpRecords/{Sid}"
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

func (c *ApiService) FetchIpRecord(Sid string) (*VoiceV1IpRecord, error) {
	path := "/v1/IpRecords/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1IpRecord{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListIpRecord'
type ListIpRecordParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListIpRecordParams) SetPageSize(PageSize int) *ListIpRecordParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListIpRecord(params *ListIpRecordParams) (*ListIpRecordResponse, error) {
	path := "/v1/IpRecords"

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

	ps := &ListIpRecordResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateIpRecord'
type UpdateIpRecordParams struct {
	// A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateIpRecordParams) SetFriendlyName(FriendlyName string) *UpdateIpRecordParams {
	params.FriendlyName = &FriendlyName
	return params
}

func (c *ApiService) UpdateIpRecord(Sid string, params *UpdateIpRecordParams) (*VoiceV1IpRecord, error) {
	path := "/v1/IpRecords/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1IpRecord{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
