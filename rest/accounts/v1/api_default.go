/*
 * Twilio - Accounts
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"fmt"
    twilio "github.com/twilio/twilio-go/client"
    "net/url"
    "strings"
    ""
)

type DefaultApiService struct {
    baseURL string
    client  *twilio.Client
}

func NewDefaultApiService(client *twilio.Client) *DefaultApiService {
    return &DefaultApiService{
        client: client,
        baseURL: fmt.Sprintf("https://studio.%s", client.BaseURL),
    }
}
// CreateCredentialAwsParams Optional parameters for the method 'CreateCredentialAws'
type CreateCredentialAwsParams struct {
    AccountSid *string `json:"AccountSid,omitempty"`
    Credentials *string `json:"Credentials,omitempty"`
    FriendlyName *string `json:"FriendlyName,omitempty"`
}

/*
CreateCredentialAws Method for CreateCredentialAws
Create a new AWS Credential
 * @param optional nil or *CreateCredentialAwsOpts - Optional Parameters:
 * @param "AccountSid" (string) - The SID of the Subaccount that this Credential should be associated with. Must be a valid Subaccount of the account issuing the request.
 * @param "Credentials" (string) - A string that contains the AWS access credentials in the format `<AWS_ACCESS_KEY_ID>:<AWS_SECRET_ACCESS_KEY>`. For example, `AKIAIOSFODNN7EXAMPLE:wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY`
 * @param "FriendlyName" (string) - A descriptive string that you create to describe the resource. It can be up to 64 characters long.
@return AccountsV1CredentialCredentialAws
*/
func (c *DefaultApiService) CreateCredentialAws(params *CreateCredentialAwsParams) (*AccountsV1CredentialCredentialAws, error) {
    path := "/v1/Credentials/AWS"

    data := url.Values{}
    headers := 0

    if params != nil && params.AccountSid != nil {
        data.Set("AccountSid", *params.AccountSid)
    }
    if params != nil && params.Credentials != nil {
        data.Set("Credentials", *params.Credentials)
    }
    if params != nil && params.FriendlyName != nil {
        data.Set("FriendlyName", *params.FriendlyName)
    }


    resp, err := c.client.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &AccountsV1CredentialCredentialAws{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// CreateCredentialPublicKeyParams Optional parameters for the method 'CreateCredentialPublicKey'
type CreateCredentialPublicKeyParams struct {
    AccountSid *string `json:"AccountSid,omitempty"`
    FriendlyName *string `json:"FriendlyName,omitempty"`
    PublicKey *string `json:"PublicKey,omitempty"`
}

/*
CreateCredentialPublicKey Method for CreateCredentialPublicKey
Create a new Public Key Credential
 * @param optional nil or *CreateCredentialPublicKeyOpts - Optional Parameters:
 * @param "AccountSid" (string) - The SID of the Subaccount that this Credential should be associated with. Must be a valid Subaccount of the account issuing the request
 * @param "FriendlyName" (string) - A descriptive string that you create to describe the resource. It can be up to 64 characters long.
 * @param "PublicKey" (string) - A URL encoded representation of the public key. For example, `-----BEGIN PUBLIC KEY-----MIIBIjANB.pa9xQIDAQAB-----END PUBLIC KEY-----`
@return AccountsV1CredentialCredentialPublicKey
*/
func (c *DefaultApiService) CreateCredentialPublicKey(params *CreateCredentialPublicKeyParams) (*AccountsV1CredentialCredentialPublicKey, error) {
    path := "/v1/Credentials/PublicKeys"

    data := url.Values{}
    headers := 0

    if params != nil && params.AccountSid != nil {
        data.Set("AccountSid", *params.AccountSid)
    }
    if params != nil && params.FriendlyName != nil {
        data.Set("FriendlyName", *params.FriendlyName)
    }
    if params != nil && params.PublicKey != nil {
        data.Set("PublicKey", *params.PublicKey)
    }


    resp, err := c.client.Post(c.baseURL+path, data, headers)
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

/*
CreateSecondaryAuthToken Method for CreateSecondaryAuthToken
Create a new secondary Auth Token
@return AccountsV1SecondaryAuthToken
*/
func (c *DefaultApiService) CreateSecondaryAuthToken() (*AccountsV1SecondaryAuthToken, error) {
    path := "/v1/AuthTokens/Secondary"

    data := 0
    headers := 0



    resp, err := c.client.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &AccountsV1SecondaryAuthToken{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

/*
DeleteCredentialAws Method for DeleteCredentialAws
Delete a Credential from your account
 * @param sid The Twilio-provided string that uniquely identifies the AWS resource to delete.
*/
func (c *DefaultApiService) DeleteCredentialAws(sid string) (error) {
    path := "/v1/Credentials/AWS/{Sid}"
    path = strings.Replace(path, "{"+"Sid"+"}", sid, -1)

    data := 0
    headers := 0



    resp, err := c.client.Delete(c.baseURL+path, data, headers)
    if err != nil {
        return err
    }

    defer resp.Body.Close()

    return nil
}

/*
DeleteCredentialPublicKey Method for DeleteCredentialPublicKey
Delete a Credential from your account
 * @param sid The Twilio-provided string that uniquely identifies the PublicKey resource to delete.
*/
func (c *DefaultApiService) DeleteCredentialPublicKey(sid string) (error) {
    path := "/v1/Credentials/PublicKeys/{Sid}"
    path = strings.Replace(path, "{"+"Sid"+"}", sid, -1)

    data := 0
    headers := 0



    resp, err := c.client.Delete(c.baseURL+path, data, headers)
    if err != nil {
        return err
    }

    defer resp.Body.Close()

    return nil
}

/*
DeleteSecondaryAuthToken Method for DeleteSecondaryAuthToken
Delete the secondary Auth Token from your account
*/
func (c *DefaultApiService) DeleteSecondaryAuthToken() (error) {
    path := "/v1/AuthTokens/Secondary"

    data := 0
    headers := 0



    resp, err := c.client.Delete(c.baseURL+path, data, headers)
    if err != nil {
        return err
    }

    defer resp.Body.Close()

    return nil
}

/*
FetchCredentialAws Method for FetchCredentialAws
Fetch the AWS credentials specified by the provided Credential Sid
 * @param sid The Twilio-provided string that uniquely identifies the AWS resource to fetch.
@return AccountsV1CredentialCredentialAws
*/
func (c *DefaultApiService) FetchCredentialAws(sid string) (*AccountsV1CredentialCredentialAws, error) {
    path := "/v1/Credentials/AWS/{Sid}"
    path = strings.Replace(path, "{"+"Sid"+"}", sid, -1)

    data := 0
    headers := 0



    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &AccountsV1CredentialCredentialAws{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

/*
FetchCredentialPublicKey Method for FetchCredentialPublicKey
Fetch the public key specified by the provided Credential Sid
 * @param sid The Twilio-provided string that uniquely identifies the PublicKey resource to fetch.
@return AccountsV1CredentialCredentialPublicKey
*/
func (c *DefaultApiService) FetchCredentialPublicKey(sid string) (*AccountsV1CredentialCredentialPublicKey, error) {
    path := "/v1/Credentials/PublicKeys/{Sid}"
    path = strings.Replace(path, "{"+"Sid"+"}", sid, -1)

    data := 0
    headers := 0



    resp, err := c.client.Get(c.baseURL+path, data, headers)
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
// ListCredentialAwsParams Optional parameters for the method 'ListCredentialAws'
type ListCredentialAwsParams struct {
    PageSize *int32 `json:"PageSize,omitempty"`
}

/*
ListCredentialAws Method for ListCredentialAws
Retrieves a collection of AWS Credentials belonging to the account used to make the request
 * @param optional nil or *ListCredentialAwsOpts - Optional Parameters:
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return AccountsV1CredentialCredentialAwsReadResponse
*/
func (c *DefaultApiService) ListCredentialAws(params *ListCredentialAwsParams) (*AccountsV1CredentialCredentialAwsReadResponse, error) {
    path := "/v1/Credentials/AWS"

    data := url.Values{}
    headers := 0

    if params != nil && params.PageSize != nil {
        data.Set("PageSize", string(*params.PageSize))
    }


    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &AccountsV1CredentialCredentialAwsReadResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// ListCredentialPublicKeyParams Optional parameters for the method 'ListCredentialPublicKey'
type ListCredentialPublicKeyParams struct {
    PageSize *int32 `json:"PageSize,omitempty"`
}

/*
ListCredentialPublicKey Method for ListCredentialPublicKey
Retrieves a collection of Public Key Credentials belonging to the account used to make the request
 * @param optional nil or *ListCredentialPublicKeyOpts - Optional Parameters:
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return AccountsV1CredentialCredentialPublicKeyReadResponse
*/
func (c *DefaultApiService) ListCredentialPublicKey(params *ListCredentialPublicKeyParams) (*AccountsV1CredentialCredentialPublicKeyReadResponse, error) {
    path := "/v1/Credentials/PublicKeys"

    data := url.Values{}
    headers := 0

    if params != nil && params.PageSize != nil {
        data.Set("PageSize", string(*params.PageSize))
    }


    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &AccountsV1CredentialCredentialPublicKeyReadResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

/*
UpdateAuthTokenPromotion Method for UpdateAuthTokenPromotion
Promote the secondary Auth Token to primary. After promoting the new token, all requests to Twilio using your old primary Auth Token will result in an error.
@return AccountsV1AuthTokenPromotion
*/
func (c *DefaultApiService) UpdateAuthTokenPromotion() (*AccountsV1AuthTokenPromotion, error) {
    path := "/v1/AuthTokens/Promote"

    data := 0
    headers := 0



    resp, err := c.client.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &AccountsV1AuthTokenPromotion{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// UpdateCredentialAwsParams Optional parameters for the method 'UpdateCredentialAws'
type UpdateCredentialAwsParams struct {
    FriendlyName *string `json:"FriendlyName,omitempty"`
}

/*
UpdateCredentialAws Method for UpdateCredentialAws
Modify the properties of a given Account
 * @param sid The Twilio-provided string that uniquely identifies the AWS resource to update.
 * @param optional nil or *UpdateCredentialAwsOpts - Optional Parameters:
 * @param "FriendlyName" (string) - A descriptive string that you create to describe the resource. It can be up to 64 characters long.
@return AccountsV1CredentialCredentialAws
*/
func (c *DefaultApiService) UpdateCredentialAws(sid string, params *UpdateCredentialAwsParams) (*AccountsV1CredentialCredentialAws, error) {
    path := "/v1/Credentials/AWS/{Sid}"
    path = strings.Replace(path, "{"+"Sid"+"}", sid, -1)

    data := url.Values{}
    headers := 0

    if params != nil && params.FriendlyName != nil {
        data.Set("FriendlyName", *params.FriendlyName)
    }


    resp, err := c.client.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &AccountsV1CredentialCredentialAws{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// UpdateCredentialPublicKeyParams Optional parameters for the method 'UpdateCredentialPublicKey'
type UpdateCredentialPublicKeyParams struct {
    FriendlyName *string `json:"FriendlyName,omitempty"`
}

/*
UpdateCredentialPublicKey Method for UpdateCredentialPublicKey
Modify the properties of a given Account
 * @param sid The Twilio-provided string that uniquely identifies the PublicKey resource to update.
 * @param optional nil or *UpdateCredentialPublicKeyOpts - Optional Parameters:
 * @param "FriendlyName" (string) - A descriptive string that you create to describe the resource. It can be up to 64 characters long.
@return AccountsV1CredentialCredentialPublicKey
*/
func (c *DefaultApiService) UpdateCredentialPublicKey(sid string, params *UpdateCredentialPublicKeyParams) (*AccountsV1CredentialCredentialPublicKey, error) {
    path := "/v1/Credentials/PublicKeys/{Sid}"
    path = strings.Replace(path, "{"+"Sid"+"}", sid, -1)

    data := url.Values{}
    headers := 0

    if params != nil && params.FriendlyName != nil {
        data.Set("FriendlyName", *params.FriendlyName)
    }


    resp, err := c.client.Post(c.baseURL+path, data, headers)
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
