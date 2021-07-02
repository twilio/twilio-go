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

// Retrieve a list of all Deployments.
func (c *ApiService) ListDeployment(ServiceSid string, EnvironmentSid string, params *ListDeploymentParams) (*ListDeploymentResponse, error) {
	path := "/v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Deployments"
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

	ps := &ListDeploymentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of Deployment records from the API. Request is executed immediately.
func (c *ApiService) DeploymentPage(ServiceSid string, EnvironmentSid string, params *ListDeploymentParams, pageToken string, pageNumber string) (*client.Page, error) {
	path := "/v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Deployments"
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

//Streams Deployment records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) DeploymentStream(ServiceSid string, EnvironmentSid string, params *ListDeploymentParams, limit int) (chan map[string]interface{}, error) {
	if params.PageSize == nil {
		params.SetPageSize(0)
	}
	params.SetPageSize(c.requestHandler.ReadLimits(*params.PageSize, limit))
	page, err := c.DeploymentPage(ServiceSid, EnvironmentSid, params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.Stream(page, limit, 0), nil
}

//Lists Deployment records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) DeploymentList(ServiceSid string, EnvironmentSid string, params *ListDeploymentParams, limit int) ([]interface{}, error) {
	if params.PageSize == nil {
		params.SetPageSize(0)
	}
	params.SetPageSize(c.requestHandler.ReadLimits(*params.PageSize, limit))
	page, err := c.DeploymentPage(ServiceSid, EnvironmentSid, params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.List(page, limit, 0), nil
}
