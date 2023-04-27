/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

    "github.com/twilio/twilio-go/client"
)


// Optional parameters for the method 'CreateSipCredential'
type CreateSipCredentialParams struct {
    // The unique id of the Account that is responsible for this resource.
    PathAccountSid *string `json:"PathAccountSid,omitempty"`
    // The username that will be passed when authenticating SIP requests. The username should be sent in response to Twilio's challenge of the initial INVITE. It can be up to 32 characters long.
    Username *string `json:"Username,omitempty"`
    // The password that the username will use when authenticating SIP requests. The password must be a minimum of 12 characters, contain at least 1 digit, and have mixed case. (eg `IWasAtSignal2018`)
    Password *string `json:"Password,omitempty"`
}

func (params *CreateSipCredentialParams) SetPathAccountSid(PathAccountSid string) (*CreateSipCredentialParams){
    params.PathAccountSid = &PathAccountSid
    return params
}
func (params *CreateSipCredentialParams) SetUsername(Username string) (*CreateSipCredentialParams){
    params.Username = &Username
    return params
}
func (params *CreateSipCredentialParams) SetPassword(Password string) (*CreateSipCredentialParams){
    params.Password = &Password
    return params
}

// Create a new credential resource.
func (c *ApiService) CreateSipCredential(CredentialListSid string, params *CreateSipCredentialParams) (*ApiV2010SipCredential, error) {
    path := "/2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials.json"
    if params != nil && params.PathAccountSid != nil {
    path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
} else {
    path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
}
    path = strings.Replace(path, "{"+"CredentialListSid"+"}", CredentialListSid, -1)

data := url.Values{}
headers := make(map[string]interface{})

if params != nil && params.Username != nil {
    data.Set("Username", *params.Username)
}
if params != nil && params.Password != nil {
    data.Set("Password", *params.Password)
}



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ApiV2010SipCredential{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'DeleteSipCredential'
type DeleteSipCredentialParams struct {
    // The unique id of the Account that is responsible for this resource.
    PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteSipCredentialParams) SetPathAccountSid(PathAccountSid string) (*DeleteSipCredentialParams){
    params.PathAccountSid = &PathAccountSid
    return params
}

// Delete a credential resource.
func (c *ApiService) DeleteSipCredential(CredentialListSid string, Sid string, params *DeleteSipCredentialParams) (error) {
    path := "/2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials/{Sid}.json"
    if params != nil && params.PathAccountSid != nil {
    path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
} else {
    path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
}
    path = strings.Replace(path, "{"+"CredentialListSid"+"}", CredentialListSid, -1)
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

// Optional parameters for the method 'FetchSipCredential'
type FetchSipCredentialParams struct {
    // The unique id of the Account that is responsible for this resource.
    PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchSipCredentialParams) SetPathAccountSid(PathAccountSid string) (*FetchSipCredentialParams){
    params.PathAccountSid = &PathAccountSid
    return params
}

// Fetch a single credential.
func (c *ApiService) FetchSipCredential(CredentialListSid string, Sid string, params *FetchSipCredentialParams) (*ApiV2010SipCredential, error) {
    path := "/2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials/{Sid}.json"
    if params != nil && params.PathAccountSid != nil {
    path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
} else {
    path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
}
    path = strings.Replace(path, "{"+"CredentialListSid"+"}", CredentialListSid, -1)
    path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

data := url.Values{}
headers := make(map[string]interface{})




    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ApiV2010SipCredential{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'ListSipCredential'
type ListSipCredentialParams struct {
    // The unique id of the Account that is responsible for this resource.
    PathAccountSid *string `json:"PathAccountSid,omitempty"`
    // How many resources to return in each list page. The default is 50, and the maximum is 1000.
    PageSize *int `json:"PageSize,omitempty"`
    // Max number of records to return.
    Limit *int `json:"limit,omitempty"`
}

func (params *ListSipCredentialParams) SetPathAccountSid(PathAccountSid string) (*ListSipCredentialParams){
    params.PathAccountSid = &PathAccountSid
    return params
}
func (params *ListSipCredentialParams) SetPageSize(PageSize int) (*ListSipCredentialParams){
    params.PageSize = &PageSize
    return params
}
func (params *ListSipCredentialParams) SetLimit(Limit int) (*ListSipCredentialParams){
    params.Limit = &Limit
    return params
}

// Retrieve a single page of SipCredential records from the API. Request is executed immediately.
func (c *ApiService) PageSipCredential(CredentialListSid string, params *ListSipCredentialParams, pageToken, pageNumber string) (*ListSipCredentialResponse, error) {
    path := "/2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials.json"

    if params != nil && params.PathAccountSid != nil {
    path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
} else {
    path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
}
    path = strings.Replace(path, "{"+"CredentialListSid"+"}", CredentialListSid, -1)

data := url.Values{}
headers := make(map[string]interface{})

if params != nil && params.PageSize != nil {
    data.Set("PageSize", fmt.Sprint(*params.PageSize))
}

    if pageToken != "" {
        data.Set("PageToken", pageToken)
    }
    if pageNumber != "" {
        data.Set("Page", pageNumber)
    }

    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListSipCredentialResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Lists SipCredential records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSipCredential(CredentialListSid string, params *ListSipCredentialParams) ([]ApiV2010SipCredential, error) {
	response, errors := c.StreamSipCredential(CredentialListSid, params)

	records := make([]ApiV2010SipCredential, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams SipCredential records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSipCredential(CredentialListSid string, params *ListSipCredentialParams) (chan ApiV2010SipCredential, chan error) {
	if params == nil {
		params = &ListSipCredentialParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ApiV2010SipCredential, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageSipCredential(CredentialListSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamSipCredential(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}


func (c *ApiService) streamSipCredential(response *ListSipCredentialResponse, params *ListSipCredentialParams, recordChannel chan ApiV2010SipCredential, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Credentials
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListSipCredentialResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListSipCredentialResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListSipCredentialResponse(nextPageUrl string) (interface{}, error) {
    if nextPageUrl == "" {
        return nil, nil
    }
    resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListSipCredentialResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }
    return ps, nil
}


// Optional parameters for the method 'UpdateSipCredential'
type UpdateSipCredentialParams struct {
    // The unique id of the Account that is responsible for this resource.
    PathAccountSid *string `json:"PathAccountSid,omitempty"`
    // The password that the username will use when authenticating SIP requests. The password must be a minimum of 12 characters, contain at least 1 digit, and have mixed case. (eg `IWasAtSignal2018`)
    Password *string `json:"Password,omitempty"`
}

func (params *UpdateSipCredentialParams) SetPathAccountSid(PathAccountSid string) (*UpdateSipCredentialParams){
    params.PathAccountSid = &PathAccountSid
    return params
}
func (params *UpdateSipCredentialParams) SetPassword(Password string) (*UpdateSipCredentialParams){
    params.Password = &Password
    return params
}

// Update a credential resource.
func (c *ApiService) UpdateSipCredential(CredentialListSid string, Sid string, params *UpdateSipCredentialParams) (*ApiV2010SipCredential, error) {
    path := "/2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials/{Sid}.json"
    if params != nil && params.PathAccountSid != nil {
    path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
} else {
    path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
}
    path = strings.Replace(path, "{"+"CredentialListSid"+"}", CredentialListSid, -1)
    path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

data := url.Values{}
headers := make(map[string]interface{})

if params != nil && params.Password != nil {
    data.Set("Password", *params.Password)
}



    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ApiV2010SipCredential{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
