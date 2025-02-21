/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Supersim
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

// Optional parameters for the method 'CreateEsimProfile'
type CreateEsimProfileParams struct {
	// The URL we should call using the `callback_method` when the status of the eSIM Profile changes. At this stage of the eSIM Profile pilot, the a request to the URL will only be called when the ESimProfile resource changes from `reserving` to `available`.
	CallbackUrl *string `json:"CallbackUrl,omitempty"`
	// The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is POST.
	CallbackMethod *string `json:"CallbackMethod,omitempty"`
	// When set to `true`, a value for `Eid` does not need to be provided. Instead, when the eSIM profile is reserved, a matching ID will be generated and returned via the `matching_id` property. This identifies the specific eSIM profile that can be used by any capable device to claim and download the profile.
	GenerateMatchingId *bool `json:"GenerateMatchingId,omitempty"`
	// Identifier of the eUICC that will claim the eSIM Profile.
	Eid *string `json:"Eid,omitempty"`
}

func (params *CreateEsimProfileParams) SetCallbackUrl(CallbackUrl string) *CreateEsimProfileParams {
	params.CallbackUrl = &CallbackUrl
	return params
}
func (params *CreateEsimProfileParams) SetCallbackMethod(CallbackMethod string) *CreateEsimProfileParams {
	params.CallbackMethod = &CallbackMethod
	return params
}
func (params *CreateEsimProfileParams) SetGenerateMatchingId(GenerateMatchingId bool) *CreateEsimProfileParams {
	params.GenerateMatchingId = &GenerateMatchingId
	return params
}
func (params *CreateEsimProfileParams) SetEid(Eid string) *CreateEsimProfileParams {
	params.Eid = &Eid
	return params
}

// Order an eSIM Profile.
func (c *ApiService) CreateEsimProfile(params *CreateEsimProfileParams) (*SupersimV1EsimProfile, error) {
	path := "/v1/ESimProfiles"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.CallbackUrl != nil {
		data.Set("CallbackUrl", *params.CallbackUrl)
	}
	if params != nil && params.CallbackMethod != nil {
		data.Set("CallbackMethod", *params.CallbackMethod)
	}
	if params != nil && params.GenerateMatchingId != nil {
		data.Set("GenerateMatchingId", fmt.Sprint(*params.GenerateMatchingId))
	}
	if params != nil && params.Eid != nil {
		data.Set("Eid", *params.Eid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1EsimProfile{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Fetch an eSIM Profile.
func (c *ApiService) FetchEsimProfile(Sid string) (*SupersimV1EsimProfile, error) {
	path := "/v1/ESimProfiles/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1EsimProfile{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListEsimProfile'
type ListEsimProfileParams struct {
	// List the eSIM Profiles that have been associated with an EId.
	Eid *string `json:"Eid,omitempty"`
	// Find the eSIM Profile resource related to a [Sim](https://www.twilio.com/docs/iot/supersim/api/sim-resource) resource by providing the SIM SID. Will always return an array with either 1 or 0 records.
	SimSid *string `json:"SimSid,omitempty"`
	// List the eSIM Profiles that are in a given status.
	Status *string `json:"Status,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int64 `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int64 `json:"limit,omitempty"`
}

func (params *ListEsimProfileParams) SetEid(Eid string) *ListEsimProfileParams {
	params.Eid = &Eid
	return params
}
func (params *ListEsimProfileParams) SetSimSid(SimSid string) *ListEsimProfileParams {
	params.SimSid = &SimSid
	return params
}
func (params *ListEsimProfileParams) SetStatus(Status string) *ListEsimProfileParams {
	params.Status = &Status
	return params
}
func (params *ListEsimProfileParams) SetPageSize(PageSize int64) *ListEsimProfileParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListEsimProfileParams) SetLimit(Limit int64) *ListEsimProfileParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of EsimProfile records from the API. Request is executed immediately.
func (c *ApiService) PageEsimProfile(params *ListEsimProfileParams, pageToken, pageNumber string) (*ListEsimProfileResponse, error) {
	path := "/v1/ESimProfiles"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Eid != nil {
		data.Set("Eid", *params.Eid)
	}
	if params != nil && params.SimSid != nil {
		data.Set("SimSid", *params.SimSid)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
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

	ps := &ListEsimProfileResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists EsimProfile records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListEsimProfile(params *ListEsimProfileParams) ([]SupersimV1EsimProfile, error) {
	response, errors := c.StreamEsimProfile(params)

	records := make([]SupersimV1EsimProfile, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams EsimProfile records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamEsimProfile(params *ListEsimProfileParams) (chan SupersimV1EsimProfile, chan error) {
	if params == nil {
		params = &ListEsimProfileParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan SupersimV1EsimProfile, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageEsimProfile(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamEsimProfile(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamEsimProfile(response *ListEsimProfileResponse, params *ListEsimProfileParams, recordChannel chan SupersimV1EsimProfile, errorChannel chan error) {
	var curRecord int64 = 1

	for response != nil {
		responseRecords := response.EsimProfiles
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListEsimProfileResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListEsimProfileResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListEsimProfileResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListEsimProfileResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
