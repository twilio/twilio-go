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

// Optional parameters for the method 'CreateDeployment'
type CreateDeploymentParams struct {
	// The SID of the Build for the Deployment.
	BuildSid *string `json:"BuildSid,omitempty"`
}

func (params *CreateDeploymentParams) SetBuildSid(BuildSid string) *CreateDeploymentParams {
	params.BuildSid = &BuildSid
	return params
}

// Create a new Deployment.
func (c *ApiService) CreateDeployment(ServiceSid string, EnvironmentSid string, params *CreateDeploymentParams) (*ServerlessV1ServiceEnvironmentDeployment, error) {
	path := "/v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Deployments"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"EnvironmentSid"+"}", EnvironmentSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.BuildSid != nil {
		data.Set("BuildSid", *params.BuildSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ServerlessV1ServiceEnvironmentDeployment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Retrieve a specific Deployment.
func (c *ApiService) FetchDeployment(ServiceSid string, EnvironmentSid string, Sid string) (*ServerlessV1ServiceEnvironmentDeployment, error) {
	path := "/v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Deployments/{Sid}"
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

	ps := &ServerlessV1ServiceEnvironmentDeployment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListDeployment'
type ListDeploymentParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListDeploymentParams) SetPageSize(PageSize int) *ListDeploymentParams {
	params.PageSize = &PageSize
	return params
}

//Retrieve a single page of Deployment records from the API. Request is executed immediately.
func (c *ApiService) PageDeployment(ServiceSid string, EnvironmentSid string, params *ListDeploymentParams, pageToken string, pageNumber string) (*ListDeploymentResponse, error) {
	path := "/v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Deployments"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"EnvironmentSid"+"}", EnvironmentSid, -1)

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

	ps := &ListDeploymentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Lists Deployment records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListDeployment(ServiceSid string, EnvironmentSid string, params *ListDeploymentParams, limit *int) ([]*ListDeploymentResponse, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageDeployment(ServiceSid, EnvironmentSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []*ListDeploymentResponse

	for response != nil {
		records = append(records, response)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, limit, c.getNextListDeploymentResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListDeploymentResponse)
	}

	return records, err
}

//Streams Deployment records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamDeployment(ServiceSid string, EnvironmentSid string, params *ListDeploymentParams, limit *int) (chan *ListDeploymentResponse, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageDeployment(ServiceSid, EnvironmentSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan *ListDeploymentResponse, 1)

	go func() {
		for response != nil {
			channel <- response

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, limit, c.getNextListDeploymentResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListDeploymentResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListDeploymentResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListDeploymentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
