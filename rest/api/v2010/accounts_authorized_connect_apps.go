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
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'FetchAuthorizedConnectApp'
type FetchAuthorizedConnectAppParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the AuthorizedConnectApp resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchAuthorizedConnectAppParams) SetPathAccountSid(PathAccountSid string) *FetchAuthorizedConnectAppParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch an instance of an authorized-connect-app
func (c *ApiService) FetchAuthorizedConnectApp(ConnectAppSid string, params *FetchAuthorizedConnectAppParams) (*ApiV2010AuthorizedConnectApp, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/AuthorizedConnectApps/{ConnectAppSid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ConnectAppSid"+"}", ConnectAppSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AuthorizedConnectApp{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListAuthorizedConnectApp'
type ListAuthorizedConnectAppParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the AuthorizedConnectApp resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListAuthorizedConnectAppParams) SetPathAccountSid(PathAccountSid string) *ListAuthorizedConnectAppParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListAuthorizedConnectAppParams) SetPageSize(PageSize int) *ListAuthorizedConnectAppParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListAuthorizedConnectAppParams) SetLimit(Limit int) *ListAuthorizedConnectAppParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of AuthorizedConnectApp records from the API. Request is executed immediately.
func (c *ApiService) PageAuthorizedConnectApp(params *ListAuthorizedConnectAppParams, pageToken, pageNumber string) (*ListAuthorizedConnectApp200Response, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/AuthorizedConnectApps.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

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

	ps := &ListAuthorizedConnectApp200Response{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists AuthorizedConnectApp records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListAuthorizedConnectApp(params *ListAuthorizedConnectAppParams) ([]ApiV2010AuthorizedConnectApp, error) {
	response, errors := c.StreamAuthorizedConnectApp(params)

	records := make([]ApiV2010AuthorizedConnectApp, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams AuthorizedConnectApp records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamAuthorizedConnectApp(params *ListAuthorizedConnectAppParams) (chan ApiV2010AuthorizedConnectApp, chan error) {
	if params == nil {
		params = &ListAuthorizedConnectAppParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ApiV2010AuthorizedConnectApp, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageAuthorizedConnectApp(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamAuthorizedConnectApp(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamAuthorizedConnectApp(response *ListAuthorizedConnectApp200Response, params *ListAuthorizedConnectAppParams, recordChannel chan ApiV2010AuthorizedConnectApp, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.AuthorizedConnectApps
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListAuthorizedConnectApp200Response)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListAuthorizedConnectApp200Response)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListAuthorizedConnectApp200Response(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAuthorizedConnectApp200Response{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
