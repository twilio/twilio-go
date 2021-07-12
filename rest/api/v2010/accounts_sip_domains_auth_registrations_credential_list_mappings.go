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

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateSipAuthRegistrationsCredentialListMapping'
type CreateSipAuthRegistrationsCredentialListMappingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The SID of the CredentialList resource to map to the SIP domain.
	CredentialListSid *string `json:"CredentialListSid,omitempty"`
}

func (params *CreateSipAuthRegistrationsCredentialListMappingParams) SetPathAccountSid(PathAccountSid string) *CreateSipAuthRegistrationsCredentialListMappingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateSipAuthRegistrationsCredentialListMappingParams) SetCredentialListSid(CredentialListSid string) *CreateSipAuthRegistrationsCredentialListMappingParams {
	params.CredentialListSid = &CredentialListSid
	return params
}

// Create a new credential list mapping resource
func (c *ApiService) CreateSipAuthRegistrationsCredentialListMapping(DomainSid string, params *CreateSipAuthRegistrationsCredentialListMappingParams) (*ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Registrations/CredentialListMappings.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"DomainSid"+"}", DomainSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CredentialListSid != nil {
		data.Set("CredentialListSid", *params.CredentialListSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteSipAuthRegistrationsCredentialListMapping'
type DeleteSipAuthRegistrationsCredentialListMappingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteSipAuthRegistrationsCredentialListMappingParams) SetPathAccountSid(PathAccountSid string) *DeleteSipAuthRegistrationsCredentialListMappingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete a credential list mapping from the requested domain
func (c *ApiService) DeleteSipAuthRegistrationsCredentialListMapping(DomainSid string, Sid string, params *DeleteSipAuthRegistrationsCredentialListMappingParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Registrations/CredentialListMappings/{Sid}.json"
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

// Optional parameters for the method 'FetchSipAuthRegistrationsCredentialListMapping'
type FetchSipAuthRegistrationsCredentialListMappingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchSipAuthRegistrationsCredentialListMappingParams) SetPathAccountSid(PathAccountSid string) *FetchSipAuthRegistrationsCredentialListMappingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch a specific instance of a credential list mapping
func (c *ApiService) FetchSipAuthRegistrationsCredentialListMapping(DomainSid string, Sid string, params *FetchSipAuthRegistrationsCredentialListMappingParams) (*ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Registrations/CredentialListMappings/{Sid}.json"
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

	ps := &ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSipAuthRegistrationsCredentialListMapping'
type ListSipAuthRegistrationsCredentialListMappingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListSipAuthRegistrationsCredentialListMappingParams) SetPathAccountSid(PathAccountSid string) *ListSipAuthRegistrationsCredentialListMappingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListSipAuthRegistrationsCredentialListMappingParams) SetPageSize(PageSize int) *ListSipAuthRegistrationsCredentialListMappingParams {
	params.PageSize = &PageSize
	return params
}

//Retrieve a single page of SipAuthRegistrationsCredentialListMapping records from the API. Request is executed immediately.
func (c *ApiService) PageSipAuthRegistrationsCredentialListMapping(DomainSid string, params *ListSipAuthRegistrationsCredentialListMappingParams, pageToken string, pageNumber string) (*ListSipAuthRegistrationsCredentialListMappingResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Registrations/CredentialListMappings.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"DomainSid"+"}", DomainSid, -1)

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

	ps := &ListSipAuthRegistrationsCredentialListMappingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Lists SipAuthRegistrationsCredentialListMapping records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSipAuthRegistrationsCredentialListMapping(DomainSid string, params *ListSipAuthRegistrationsCredentialListMappingParams, limit *int) ([]*ListSipAuthRegistrationsCredentialListMappingResponse, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageSipAuthRegistrationsCredentialListMapping(DomainSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []*ListSipAuthRegistrationsCredentialListMappingResponse

	for response != nil {
		records = append(records, response)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, limit, c.getNextListSipAuthRegistrationsCredentialListMappingResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListSipAuthRegistrationsCredentialListMappingResponse)
	}

	return records, err
}

//Streams SipAuthRegistrationsCredentialListMapping records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSipAuthRegistrationsCredentialListMapping(DomainSid string, params *ListSipAuthRegistrationsCredentialListMappingParams, limit *int) (chan *ListSipAuthRegistrationsCredentialListMappingResponse, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageSipAuthRegistrationsCredentialListMapping(DomainSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan *ListSipAuthRegistrationsCredentialListMappingResponse, 1)

	go func() {
		for response != nil {
			channel <- response

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, limit, c.getNextListSipAuthRegistrationsCredentialListMappingResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListSipAuthRegistrationsCredentialListMappingResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListSipAuthRegistrationsCredentialListMappingResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSipAuthRegistrationsCredentialListMappingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
