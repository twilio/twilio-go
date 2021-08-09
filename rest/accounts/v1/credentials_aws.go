/*
 * Twilio - Accounts
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.1
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

// Optional parameters for the method 'CreateCredentialAws'
type CreateCredentialAwsParams struct {
	// The SID of the Subaccount that this Credential should be associated with. Must be a valid Subaccount of the account issuing the request.
	AccountSid *string `json:"AccountSid,omitempty"`
	// A string that contains the AWS access credentials in the format `<AWS_ACCESS_KEY_ID>:<AWS_SECRET_ACCESS_KEY>`. For example, `AKIAIOSFODNN7EXAMPLE:wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY`
	Credentials *string `json:"Credentials,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateCredentialAwsParams) SetAccountSid(AccountSid string) *CreateCredentialAwsParams {
	params.AccountSid = &AccountSid
	return params
}
func (params *CreateCredentialAwsParams) SetCredentials(Credentials string) *CreateCredentialAwsParams {
	params.Credentials = &Credentials
	return params
}
func (params *CreateCredentialAwsParams) SetFriendlyName(FriendlyName string) *CreateCredentialAwsParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Create a new AWS Credential
func (c *ApiService) CreateCredentialAws(params *CreateCredentialAwsParams) (*AccountsV1CredentialAws, error) {
	path := "/v1/Credentials/AWS"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AccountSid != nil {
		data.Set("AccountSid", *params.AccountSid)
	}
	if params != nil && params.Credentials != nil {
		data.Set("Credentials", *params.Credentials)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AccountsV1CredentialAws{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a Credential from your account
func (c *ApiService) DeleteCredentialAws(Sid string) error {
	path := "/v1/Credentials/AWS/{Sid}"
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

// Fetch the AWS credentials specified by the provided Credential Sid
func (c *ApiService) FetchCredentialAws(Sid string) (*AccountsV1CredentialAws, error) {
	path := "/v1/Credentials/AWS/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AccountsV1CredentialAws{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListCredentialAws'
type ListCredentialAwsParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListCredentialAwsParams) SetPageSize(PageSize int) *ListCredentialAwsParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListCredentialAwsParams) SetLimit(Limit int) *ListCredentialAwsParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of CredentialAws records from the API. Request is executed immediately.
func (c *ApiService) PageCredentialAws(params *ListCredentialAwsParams, pageToken string, pageNumber string) (*ListCredentialAwsResponse, error) {
	path := "/v1/Credentials/AWS"

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

	ps := &ListCredentialAwsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists CredentialAws records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListCredentialAws(params *ListCredentialAwsParams) ([]AccountsV1CredentialAws, error) {
	if params == nil {
		params = &ListCredentialAwsParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageCredentialAws(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []AccountsV1CredentialAws

	for response != nil {
		records = append(records, response.Credentials...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListCredentialAwsResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListCredentialAwsResponse)
	}

	return records, err
}

// Streams CredentialAws records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamCredentialAws(params *ListCredentialAwsParams) (chan AccountsV1CredentialAws, error) {
	if params == nil {
		params = &ListCredentialAwsParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageCredentialAws(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan AccountsV1CredentialAws, 1)

	go func() {
		for response != nil {
			for item := range response.Credentials {
				channel <- response.Credentials[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListCredentialAwsResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListCredentialAwsResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListCredentialAwsResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCredentialAwsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateCredentialAws'
type UpdateCredentialAwsParams struct {
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateCredentialAwsParams) SetFriendlyName(FriendlyName string) *UpdateCredentialAwsParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Modify the properties of a given Account
func (c *ApiService) UpdateCredentialAws(Sid string, params *UpdateCredentialAwsParams) (*AccountsV1CredentialAws, error) {
	path := "/v1/Credentials/AWS/{Sid}"
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

	ps := &AccountsV1CredentialAws{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
