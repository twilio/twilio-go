/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Trusthub
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

// Optional parameters for the method 'CreateTrustProduct'
type CreateTrustProductParams struct {
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The email address that will receive updates when the Trust Product resource changes status.
	Email *string `json:"Email,omitempty"`
	// The unique string of a policy that is associated to the Trust Product resource.
	PolicySid *string `json:"PolicySid,omitempty"`
	// The URL we call to inform your application of status changes.
	StatusCallback *string `json:"StatusCallback,omitempty"`
}

func (params *CreateTrustProductParams) SetFriendlyName(FriendlyName string) *CreateTrustProductParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateTrustProductParams) SetEmail(Email string) *CreateTrustProductParams {
	params.Email = &Email
	return params
}
func (params *CreateTrustProductParams) SetPolicySid(PolicySid string) *CreateTrustProductParams {
	params.PolicySid = &PolicySid
	return params
}
func (params *CreateTrustProductParams) SetStatusCallback(StatusCallback string) *CreateTrustProductParams {
	params.StatusCallback = &StatusCallback
	return params
}

// Create a new Trust Product.
func (c *ApiService) CreateTrustProduct(params *CreateTrustProductParams) (*TrusthubV1TrustProduct, error) {
	path := "/v1/TrustProducts"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Email != nil {
		data.Set("Email", *params.Email)
	}
	if params != nil && params.PolicySid != nil {
		data.Set("PolicySid", *params.PolicySid)
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1TrustProduct{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Trust Product.
func (c *ApiService) DeleteTrustProduct(Sid string) error {
	path := "/v1/TrustProducts/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch a specific Trust Product instance.
func (c *ApiService) FetchTrustProduct(Sid string) (*TrusthubV1TrustProduct, error) {
	path := "/v1/TrustProducts/{Sid}"
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

	ps := &TrusthubV1TrustProduct{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListTrustProduct'
type ListTrustProductParams struct {
	// The verification status of the Trust Product resource.
	Status *string `json:"Status,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The unique string of a policy that is associated to the Trust Product resource.
	PolicySid *string `json:"PolicySid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListTrustProductParams) SetStatus(Status string) *ListTrustProductParams {
	params.Status = &Status
	return params
}
func (params *ListTrustProductParams) SetFriendlyName(FriendlyName string) *ListTrustProductParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListTrustProductParams) SetPolicySid(PolicySid string) *ListTrustProductParams {
	params.PolicySid = &PolicySid
	return params
}
func (params *ListTrustProductParams) SetPageSize(PageSize int) *ListTrustProductParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListTrustProductParams) SetLimit(Limit int) *ListTrustProductParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of TrustProduct records from the API. Request is executed immediately.
func (c *ApiService) PageTrustProduct(params *ListTrustProductParams, pageToken, pageNumber string) (*ListTrustProductResponse, error) {
	path := "/v1/TrustProducts"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.PolicySid != nil {
		data.Set("PolicySid", *params.PolicySid)
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

	ps := &ListTrustProductResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists TrustProduct records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListTrustProduct(params *ListTrustProductParams) ([]TrusthubV1TrustProduct, error) {
	response, errors := c.StreamTrustProduct(params)

	records := make([]TrusthubV1TrustProduct, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams TrustProduct records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamTrustProduct(params *ListTrustProductParams) (chan TrusthubV1TrustProduct, chan error) {
	if params == nil {
		params = &ListTrustProductParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan TrusthubV1TrustProduct, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageTrustProduct(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamTrustProduct(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamTrustProduct(response *ListTrustProductResponse, params *ListTrustProductParams, recordChannel chan TrusthubV1TrustProduct, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Results
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListTrustProductResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListTrustProductResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListTrustProductResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTrustProductResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateTrustProduct'
type UpdateTrustProductParams struct {
	//
	Status *string `json:"Status,omitempty"`
	// The URL we call to inform your application of status changes.
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The email address that will receive updates when the Trust Product resource changes status.
	Email *string `json:"Email,omitempty"`
}

func (params *UpdateTrustProductParams) SetStatus(Status string) *UpdateTrustProductParams {
	params.Status = &Status
	return params
}
func (params *UpdateTrustProductParams) SetStatusCallback(StatusCallback string) *UpdateTrustProductParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *UpdateTrustProductParams) SetFriendlyName(FriendlyName string) *UpdateTrustProductParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateTrustProductParams) SetEmail(Email string) *UpdateTrustProductParams {
	params.Email = &Email
	return params
}

// Updates a Trust Product in an account.
func (c *ApiService) UpdateTrustProduct(Sid string, params *UpdateTrustProductParams) (*TrusthubV1TrustProduct, error) {
	path := "/v1/TrustProducts/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Email != nil {
		data.Set("Email", *params.Email)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1TrustProduct{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
