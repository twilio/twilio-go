/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Microvisor
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

// Optional parameters for the method 'CreateAccountSecret'
type CreateAccountSecretParams struct {
	// The secret key; up to 100 characters.
	Key *string `json:"Key,omitempty"`
	// The secret value; up to 4096 characters.
	Value *string `json:"Value,omitempty"`
}

func (params *CreateAccountSecretParams) SetKey(Key string) *CreateAccountSecretParams {
	params.Key = &Key
	return params
}
func (params *CreateAccountSecretParams) SetValue(Value string) *CreateAccountSecretParams {
	params.Value = &Value
	return params
}

// Create a secret for an Account.
func (c *ApiService) CreateAccountSecret(params *CreateAccountSecretParams) (*MicrovisorV1AccountSecret, error) {
	path := "/v1/Secrets"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Key != nil {
		data.Set("Key", *params.Key)
	}
	if params != nil && params.Value != nil {
		data.Set("Value", *params.Value)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MicrovisorV1AccountSecret{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a secret for an Account.
func (c *ApiService) DeleteAccountSecret(Key string) error {
	path := "/v1/Secrets/{Key}"
	path = strings.Replace(path, "{"+"Key"+"}", Key, -1)

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

// Retrieve a Secret for an Account.
func (c *ApiService) FetchAccountSecret(Key string) (*MicrovisorV1AccountSecret, error) {
	path := "/v1/Secrets/{Key}"
	path = strings.Replace(path, "{"+"Key"+"}", Key, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MicrovisorV1AccountSecret{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListAccountSecret'
type ListAccountSecretParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int64 `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int64 `json:"limit,omitempty"`
}

func (params *ListAccountSecretParams) SetPageSize(PageSize int64) *ListAccountSecretParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListAccountSecretParams) SetLimit(Limit int64) *ListAccountSecretParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of AccountSecret records from the API. Request is executed immediately.
func (c *ApiService) PageAccountSecret(params *ListAccountSecretParams, pageToken, pageNumber string) (*ListAccountSecretResponse, error) {
	path := "/v1/Secrets"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
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

	ps := &ListAccountSecretResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists AccountSecret records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListAccountSecret(params *ListAccountSecretParams) ([]MicrovisorV1AccountSecret, error) {
	response, errors := c.StreamAccountSecret(params)

	records := make([]MicrovisorV1AccountSecret, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams AccountSecret records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamAccountSecret(params *ListAccountSecretParams) (chan MicrovisorV1AccountSecret, chan error) {
	if params == nil {
		params = &ListAccountSecretParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan MicrovisorV1AccountSecret, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageAccountSecret(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamAccountSecret(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamAccountSecret(response *ListAccountSecretResponse, params *ListAccountSecretParams, recordChannel chan MicrovisorV1AccountSecret, errorChannel chan error) {
	var curRecord int64 = 1

	for response != nil {
		responseRecords := response.Secrets
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListAccountSecretResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListAccountSecretResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListAccountSecretResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAccountSecretResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateAccountSecret'
type UpdateAccountSecretParams struct {
	// The secret value; up to 4096 characters.
	Value *string `json:"Value,omitempty"`
}

func (params *UpdateAccountSecretParams) SetValue(Value string) *UpdateAccountSecretParams {
	params.Value = &Value
	return params
}

// Update a secret for an Account.
func (c *ApiService) UpdateAccountSecret(Key string, params *UpdateAccountSecretParams) (*MicrovisorV1AccountSecret, error) {
	path := "/v1/Secrets/{Key}"
	path = strings.Replace(path, "{"+"Key"+"}", Key, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Value != nil {
		data.Set("Value", *params.Value)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MicrovisorV1AccountSecret{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
