/*
 * Twilio - Api
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

// Optional parameters for the method 'CreateSipAuthCallsCredentialListMapping'
type CreateSipAuthCallsCredentialListMappingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The SID of the CredentialList resource to map to the SIP domain.
	CredentialListSid *string `json:"CredentialListSid,omitempty"`
}

func (params *CreateSipAuthCallsCredentialListMappingParams) SetPathAccountSid(PathAccountSid string) *CreateSipAuthCallsCredentialListMappingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateSipAuthCallsCredentialListMappingParams) SetCredentialListSid(CredentialListSid string) *CreateSipAuthCallsCredentialListMappingParams {
	params.CredentialListSid = &CredentialListSid
	return params
}

// Create a new credential list mapping resource
func (c *ApiService) CreateSipAuthCallsCredentialListMapping(DomainSid string, params *CreateSipAuthCallsCredentialListMappingParams) (*ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/CredentialListMappings.json"
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

	ps := &ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteSipAuthCallsCredentialListMapping'
type DeleteSipAuthCallsCredentialListMappingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteSipAuthCallsCredentialListMappingParams) SetPathAccountSid(PathAccountSid string) *DeleteSipAuthCallsCredentialListMappingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete a credential list mapping from the requested domain
func (c *ApiService) DeleteSipAuthCallsCredentialListMapping(DomainSid string, Sid string, params *DeleteSipAuthCallsCredentialListMappingParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/CredentialListMappings/{Sid}.json"
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

// Optional parameters for the method 'FetchSipAuthCallsCredentialListMapping'
type FetchSipAuthCallsCredentialListMappingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchSipAuthCallsCredentialListMappingParams) SetPathAccountSid(PathAccountSid string) *FetchSipAuthCallsCredentialListMappingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch a specific instance of a credential list mapping
func (c *ApiService) FetchSipAuthCallsCredentialListMapping(DomainSid string, Sid string, params *FetchSipAuthCallsCredentialListMappingParams) (*ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/CredentialListMappings/{Sid}.json"
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

	ps := &ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSipAuthCallsCredentialListMapping'
type ListSipAuthCallsCredentialListMappingParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListSipAuthCallsCredentialListMappingParams) SetPathAccountSid(PathAccountSid string) *ListSipAuthCallsCredentialListMappingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListSipAuthCallsCredentialListMappingParams) SetPageSize(PageSize int) *ListSipAuthCallsCredentialListMappingParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a single page of SipAuthCallsCredentialListMapping records from the API. Request is executed immediately.
func (c *ApiService) PageSipAuthCallsCredentialListMapping(DomainSid string, params *ListSipAuthCallsCredentialListMappingParams, pageToken string, pageNumber string) (*ListSipAuthCallsCredentialListMappingResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/CredentialListMappings.json"

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

	ps := &ListSipAuthCallsCredentialListMappingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SipAuthCallsCredentialListMapping records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSipAuthCallsCredentialListMapping(DomainSid string, params *ListSipAuthCallsCredentialListMappingParams, limit int) ([]ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping, error) {
	if params == nil {
		params = &ListSipAuthCallsCredentialListMappingParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageSipAuthCallsCredentialListMapping(DomainSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping

	for response != nil {
		records = append(records, response.Contents...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, limit, c.getNextListSipAuthCallsCredentialListMappingResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListSipAuthCallsCredentialListMappingResponse)
	}

	return records, err
}

// Streams SipAuthCallsCredentialListMapping records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSipAuthCallsCredentialListMapping(DomainSid string, params *ListSipAuthCallsCredentialListMappingParams, limit int) (chan ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping, error) {
	if params == nil {
		params = &ListSipAuthCallsCredentialListMappingParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageSipAuthCallsCredentialListMapping(DomainSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping, 1)

	go func() {
		for response != nil {
			for item := range response.Contents {
				channel <- response.Contents[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, limit, c.getNextListSipAuthCallsCredentialListMappingResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListSipAuthCallsCredentialListMappingResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListSipAuthCallsCredentialListMappingResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSipAuthCallsCredentialListMappingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
