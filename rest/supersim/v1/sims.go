/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.2
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

// Optional parameters for the method 'CreateSim'
type CreateSimParams struct {
	// The [ICCID](https://en.wikipedia.org/wiki/Subscriber_identity_module#ICCID) of the Super SIM to be added to your Account.
	Iccid *string `json:"Iccid,omitempty"`
	// The 10-digit code required to claim the Super SIM for your Account.
	RegistrationCode *string `json:"RegistrationCode,omitempty"`
}

func (params *CreateSimParams) SetIccid(Iccid string) *CreateSimParams {
	params.Iccid = &Iccid
	return params
}
func (params *CreateSimParams) SetRegistrationCode(RegistrationCode string) *CreateSimParams {
	params.RegistrationCode = &RegistrationCode
	return params
}

// Register a Super SIM to your Account
func (c *ApiService) CreateSim(params *CreateSimParams) (*SupersimV1Sim, error) {
	path := "/v1/Sims"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Iccid != nil {
		data.Set("Iccid", *params.Iccid)
	}
	if params != nil && params.RegistrationCode != nil {
		data.Set("RegistrationCode", *params.RegistrationCode)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1Sim{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Fetch a Super SIM instance from your account.
func (c *ApiService) FetchSim(Sid string) (*SupersimV1Sim, error) {
	path := "/v1/Sims/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1Sim{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSim'
type ListSimParams struct {
	// The status of the Sim resources to read. Can be `new`, `ready`, `active`, `inactive`, or `scheduled`.
	Status *string `json:"Status,omitempty"`
	// The SID or unique name of the Fleet to which a list of Sims are assigned.
	Fleet *string `json:"Fleet,omitempty"`
	// The [ICCID](https://en.wikipedia.org/wiki/Subscriber_identity_module#ICCID) associated with a Super SIM to filter the list by. Passing this parameter will always return a list containing zero or one SIMs.
	Iccid *string `json:"Iccid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSimParams) SetStatus(Status string) *ListSimParams {
	params.Status = &Status
	return params
}
func (params *ListSimParams) SetFleet(Fleet string) *ListSimParams {
	params.Fleet = &Fleet
	return params
}
func (params *ListSimParams) SetIccid(Iccid string) *ListSimParams {
	params.Iccid = &Iccid
	return params
}
func (params *ListSimParams) SetPageSize(PageSize int) *ListSimParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSimParams) SetLimit(Limit int) *ListSimParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Sim records from the API. Request is executed immediately.
func (c *ApiService) PageSim(params *ListSimParams, pageToken, pageNumber string) (*ListSimResponse, error) {
	path := "/v1/Sims"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.Fleet != nil {
		data.Set("Fleet", *params.Fleet)
	}
	if params != nil && params.Iccid != nil {
		data.Set("Iccid", *params.Iccid)
	}
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

	ps := &ListSimResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Sim records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSim(params *ListSimParams) ([]SupersimV1Sim, error) {
	response, errors := c.StreamSim(params)

	records := make([]SupersimV1Sim, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Sim records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSim(params *ListSimParams) (chan SupersimV1Sim, chan error) {
	if params == nil {
		params = &ListSimParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan SupersimV1Sim, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageSim(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamSim(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamSim(response *ListSimResponse, params *ListSimParams, recordChannel chan SupersimV1Sim, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Sims
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListSimResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListSimResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListSimResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSimResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSim'
type UpdateSimParams struct {
	// The SID of the Account to which the Sim resource should belong. The Account SID can only be that of the requesting Account or that of a Subaccount of the requesting Account. Only valid when the Sim resource's status is new.
	AccountSid *string `json:"AccountSid,omitempty"`
	// The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is POST.
	CallbackMethod *string `json:"CallbackMethod,omitempty"`
	// The URL we should call using the `callback_method` after an asynchronous update has finished.
	CallbackUrl *string `json:"CallbackUrl,omitempty"`
	// The SID or unique name of the Fleet to which the SIM resource should be assigned.
	Fleet *string `json:"Fleet,omitempty"`
	// The new status of the resource. Can be: `ready`, `active`, or `inactive`. See the [Super SIM Status Values](https://www.twilio.com/docs/iot/supersim/api/sim-resource#status-values) for more info.
	Status *string `json:"Status,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *UpdateSimParams) SetAccountSid(AccountSid string) *UpdateSimParams {
	params.AccountSid = &AccountSid
	return params
}
func (params *UpdateSimParams) SetCallbackMethod(CallbackMethod string) *UpdateSimParams {
	params.CallbackMethod = &CallbackMethod
	return params
}
func (params *UpdateSimParams) SetCallbackUrl(CallbackUrl string) *UpdateSimParams {
	params.CallbackUrl = &CallbackUrl
	return params
}
func (params *UpdateSimParams) SetFleet(Fleet string) *UpdateSimParams {
	params.Fleet = &Fleet
	return params
}
func (params *UpdateSimParams) SetStatus(Status string) *UpdateSimParams {
	params.Status = &Status
	return params
}
func (params *UpdateSimParams) SetUniqueName(UniqueName string) *UpdateSimParams {
	params.UniqueName = &UniqueName
	return params
}

// Updates the given properties of a Super SIM instance from your account.
func (c *ApiService) UpdateSim(Sid string, params *UpdateSimParams) (*SupersimV1Sim, error) {
	path := "/v1/Sims/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AccountSid != nil {
		data.Set("AccountSid", *params.AccountSid)
	}
	if params != nil && params.CallbackMethod != nil {
		data.Set("CallbackMethod", *params.CallbackMethod)
	}
	if params != nil && params.CallbackUrl != nil {
		data.Set("CallbackUrl", *params.CallbackUrl)
	}
	if params != nil && params.Fleet != nil {
		data.Set("Fleet", *params.Fleet)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1Sim{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
