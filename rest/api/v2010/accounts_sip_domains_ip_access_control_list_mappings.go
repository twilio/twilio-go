/*
 * Twilio - Api
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
)

// Optional parameters for the method 'CreateSipIpAccessControlListMapping'
type CreateSipIpAccessControlListMappingParams struct {
	// The unique id of the Account that is responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The unique id of the IP access control list to map to the SIP domain.
	IpAccessControlListSid *string `json:"IpAccessControlListSid,omitempty"`
}

func (params *CreateSipIpAccessControlListMappingParams) SetPathAccountSid(PathAccountSid string) *CreateSipIpAccessControlListMappingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateSipIpAccessControlListMappingParams) SetIpAccessControlListSid(IpAccessControlListSid string) *CreateSipIpAccessControlListMappingParams {
	params.IpAccessControlListSid = &IpAccessControlListSid
	return params
}

// Create a new IpAccessControlListMapping resource.
func (c *ApiService) CreateSipIpAccessControlListMapping(DomainSid string, params *CreateSipIpAccessControlListMappingParams) (*ApiV2010AccountSipSipDomainSipIpAccessControlListMapping, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/IpAccessControlListMappings.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"DomainSid"+"}", DomainSid, -1)

	data := url.Values{}
	if params != nil && params.IpAccessControlListSid != nil {
		data.Set("IpAccessControlListSid", *params.IpAccessControlListSid)
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountSipSipDomainSipIpAccessControlListMapping{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteSipIpAccessControlListMapping'
type DeleteSipIpAccessControlListMappingParams struct {
	// The unique id of the Account that is responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteSipIpAccessControlListMappingParams) SetPathAccountSid(PathAccountSid string) *DeleteSipIpAccessControlListMappingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete an IpAccessControlListMapping resource.
func (c *ApiService) DeleteSipIpAccessControlListMapping(DomainSid string, Sid string, params *DeleteSipIpAccessControlListMappingParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/IpAccessControlListMappings/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"DomainSid"+"}", DomainSid, -1)
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

// Optional parameters for the method 'FetchSipIpAccessControlListMapping'
type FetchSipIpAccessControlListMappingParams struct {
	// The unique id of the Account that is responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchSipIpAccessControlListMappingParams) SetPathAccountSid(PathAccountSid string) *FetchSipIpAccessControlListMappingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch an IpAccessControlListMapping resource.
func (c *ApiService) FetchSipIpAccessControlListMapping(DomainSid string, Sid string, params *FetchSipIpAccessControlListMappingParams) (*ApiV2010AccountSipSipDomainSipIpAccessControlListMapping, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/IpAccessControlListMappings/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"DomainSid"+"}", DomainSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountSipSipDomainSipIpAccessControlListMapping{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSipIpAccessControlListMapping'
type ListSipIpAccessControlListMappingParams struct {
	// The unique id of the Account that is responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListSipIpAccessControlListMappingParams) SetPathAccountSid(PathAccountSid string) *ListSipIpAccessControlListMappingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListSipIpAccessControlListMappingParams) SetPageSize(PageSize int) *ListSipIpAccessControlListMappingParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of IpAccessControlListMapping resources.
func (c *ApiService) ListSipIpAccessControlListMapping(DomainSid string, params *ListSipIpAccessControlListMappingParams) (*ListSipIpAccessControlListMappingResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/IpAccessControlListMappings.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"DomainSid"+"}", DomainSid, -1)

	data := url.Values{}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSipIpAccessControlListMappingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
