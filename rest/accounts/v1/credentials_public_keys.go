/*
 * Twilio - Accounts
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

// Optional parameters for the method 'CreateCredentialPublicKey'
type CreateCredentialPublicKeyParams struct {
	// The SID of the Subaccount that this Credential should be associated with. Must be a valid Subaccount of the account issuing the request
	AccountSid *string `json:"AccountSid,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// A URL encoded representation of the public key. For example, `-----BEGIN PUBLIC KEY-----MIIBIjANB.pa9xQIDAQAB-----END PUBLIC KEY-----`
	PublicKey *string `json:"PublicKey,omitempty"`
}

func (params *CreateCredentialPublicKeyParams) SetAccountSid(AccountSid string) *CreateCredentialPublicKeyParams {
	params.AccountSid = &AccountSid
	return params
}
func (params *CreateCredentialPublicKeyParams) SetFriendlyName(FriendlyName string) *CreateCredentialPublicKeyParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateCredentialPublicKeyParams) SetPublicKey(PublicKey string) *CreateCredentialPublicKeyParams {
	params.PublicKey = &PublicKey
	return params
}

// Create a new Public Key Credential
func (c *ApiService) CreateCredentialPublicKey(params *CreateCredentialPublicKeyParams) (*AccountsV1CredentialCredentialPublicKey, error) {
	path := "/v1/Credentials/PublicKeys"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AccountSid != nil {
		data.Set("AccountSid", *params.AccountSid)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.PublicKey != nil {
		data.Set("PublicKey", *params.PublicKey)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AccountsV1CredentialCredentialPublicKey{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a Credential from your account
func (c *ApiService) DeleteCredentialPublicKey(Sid string) error {
	path := "/v1/Credentials/PublicKeys/{Sid}"
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

// Fetch the public key specified by the provided Credential Sid
func (c *ApiService) FetchCredentialPublicKey(Sid string) (*AccountsV1CredentialCredentialPublicKey, error) {
	path := "/v1/Credentials/PublicKeys/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AccountsV1CredentialCredentialPublicKey{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListCredentialPublicKey'
type ListCredentialPublicKeyParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListCredentialPublicKeyParams) SetPageSize(PageSize int) *ListCredentialPublicKeyParams {
	params.PageSize = &PageSize
	return params
}

// Retrieves a collection of Public Key Credentials belonging to the account used to make the request
func (c *ApiService) ListCredentialPublicKey(params *ListCredentialPublicKeyParams) (*ListCredentialPublicKeyResponse, error) {
	path := "/v1/Credentials/PublicKeys"

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

	ps := &ListCredentialPublicKeyResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateCredentialPublicKey'
type UpdateCredentialPublicKeyParams struct {
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateCredentialPublicKeyParams) SetFriendlyName(FriendlyName string) *UpdateCredentialPublicKeyParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Modify the properties of a given Account
func (c *ApiService) UpdateCredentialPublicKey(Sid string, params *UpdateCredentialPublicKeyParams) (*AccountsV1CredentialCredentialPublicKey, error) {
	path := "/v1/Credentials/PublicKeys/{Sid}"
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

	ps := &AccountsV1CredentialCredentialPublicKey{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
