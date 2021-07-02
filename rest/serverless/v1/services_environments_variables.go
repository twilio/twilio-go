/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.18.0
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

// Optional parameters for the method 'CreateVariable'
type CreateVariableParams struct {
	// A string by which the Variable resource can be referenced. It can be a maximum of 128 characters.
	Key *string `json:"Key,omitempty"`
	// A string that contains the actual value of the Variable. It can be a maximum of 450 bytes in size.
	Value *string `json:"Value,omitempty"`
}

func (params *CreateVariableParams) SetKey(Key string) *CreateVariableParams {
	params.Key = &Key
	return params
}
func (params *CreateVariableParams) SetValue(Value string) *CreateVariableParams {
	params.Value = &Value
	return params
}

// Create a new Variable.
func (c *ApiService) CreateVariable(ServiceSid string, EnvironmentSid string, params *CreateVariableParams) (*ServerlessV1ServiceEnvironmentVariable, error) {
	path := "/v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Variables"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"EnvironmentSid"+"}", EnvironmentSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

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

	ps := &ServerlessV1ServiceEnvironmentVariable{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Variable.
func (c *ApiService) DeleteVariable(ServiceSid string, EnvironmentSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Variables/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"EnvironmentSid"+"}", EnvironmentSid, -1)
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

// Retrieve a specific Variable.
func (c *ApiService) FetchVariable(ServiceSid string, EnvironmentSid string, Sid string) (*ServerlessV1ServiceEnvironmentVariable, error) {
	path := "/v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Variables/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"EnvironmentSid"+"}", EnvironmentSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ServerlessV1ServiceEnvironmentVariable{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListVariable'
type ListVariableParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListVariableParams) SetPageSize(PageSize int) *ListVariableParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all Variables.
func (c *ApiService) ListVariable(ServiceSid string, EnvironmentSid string, params *ListVariableParams) (*ListVariableResponse, error) {
	path := "/v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Variables"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"EnvironmentSid"+"}", EnvironmentSid, -1)

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

	ps := &ListVariableResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of Variable records from the API. Request is executed immediately.
func (c *ApiService) VariablePage(ServiceSid string, EnvironmentSid string, params *ListVariableParams, pageToken string, pageNumber string) (*client.Page, error) {
	path := "/v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Variables"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"EnvironmentSid"+"}", EnvironmentSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	data.Set("PageToken", pageToken)
	data.Set("PageNumber", pageNumber)

	response, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	return client.NewPage(c.baseURL, response), nil
}

//Streams Variable records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) VariableStream(ServiceSid string, EnvironmentSid string, params *ListVariableParams, limit int) (chan map[string]interface{}, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	page, err := c.VariablePage(ServiceSid, EnvironmentSid, params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.Stream(page, limit, 0), nil
}

//Lists Variable records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) VariableList(ServiceSid string, EnvironmentSid string, params *ListVariableParams, limit int) ([]interface{}, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	page, err := c.VariablePage(ServiceSid, EnvironmentSid, params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.List(page, limit, 0), nil
}

// Optional parameters for the method 'UpdateVariable'
type UpdateVariableParams struct {
	// A string by which the Variable resource can be referenced. It can be a maximum of 128 characters.
	Key *string `json:"Key,omitempty"`
	// A string that contains the actual value of the Variable. It can be a maximum of 450 bytes in size.
	Value *string `json:"Value,omitempty"`
}

func (params *UpdateVariableParams) SetKey(Key string) *UpdateVariableParams {
	params.Key = &Key
	return params
}
func (params *UpdateVariableParams) SetValue(Value string) *UpdateVariableParams {
	params.Value = &Value
	return params
}

// Update a specific Variable.
func (c *ApiService) UpdateVariable(ServiceSid string, EnvironmentSid string, Sid string, params *UpdateVariableParams) (*ServerlessV1ServiceEnvironmentVariable, error) {
	path := "/v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Variables/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"EnvironmentSid"+"}", EnvironmentSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

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

	ps := &ServerlessV1ServiceEnvironmentVariable{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
