/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Serverless
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

// Optional parameters for the method 'CreateDeployment'
type CreateDeploymentParams struct {
	// The SID of the Build for the Deployment.
	BuildSid *string `json:"BuildSid,omitempty"`
	// Whether the Deployment is a plugin.
	IsPlugin *bool `json:"IsPlugin,omitempty"`
}

func (params *CreateDeploymentParams) SetBuildSid(BuildSid string) *CreateDeploymentParams {
	params.BuildSid = &BuildSid
	return params
}
func (params *CreateDeploymentParams) SetIsPlugin(IsPlugin bool) *CreateDeploymentParams {
	params.IsPlugin = &IsPlugin
	return params
}

// Create a new Deployment.
func (c *ApiService) CreateDeployment(ServiceSid string, EnvironmentSid string, params *CreateDeploymentParams) (*ServerlessV1Deployment, error) {
	path := "/v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Deployments"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"EnvironmentSid"+"}", EnvironmentSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.BuildSid != nil {
		data.Set("BuildSid", *params.BuildSid)
	}
	if params != nil && params.IsPlugin != nil {
		data.Set("IsPlugin", fmt.Sprint(*params.IsPlugin))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ServerlessV1Deployment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Retrieve a specific Deployment.
func (c *ApiService) FetchDeployment(ServiceSid string, EnvironmentSid string, Sid string) (*ServerlessV1Deployment, error) {
	path := "/v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Deployments/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"EnvironmentSid"+"}", EnvironmentSid, -1)
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

	ps := &ServerlessV1Deployment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListDeployment'
type ListDeploymentParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListDeploymentParams) SetPageSize(PageSize int) *ListDeploymentParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListDeploymentParams) SetLimit(Limit int) *ListDeploymentParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Deployment records from the API. Request is executed immediately.
func (c *ApiService) PageDeployment(ServiceSid string, EnvironmentSid string, params *ListDeploymentParams, pageToken, pageNumber string) (*ListDeploymentResponse, error) {
	path := "/v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Deployments"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"EnvironmentSid"+"}", EnvironmentSid, -1)

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

	ps := &ListDeploymentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Deployment records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListDeployment(ServiceSid string, EnvironmentSid string, params *ListDeploymentParams) ([]ServerlessV1Deployment, error) {
	response, errors := c.StreamDeployment(ServiceSid, EnvironmentSid, params)

	records := make([]ServerlessV1Deployment, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Deployment records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamDeployment(ServiceSid string, EnvironmentSid string, params *ListDeploymentParams) (chan ServerlessV1Deployment, chan error) {
	if params == nil {
		params = &ListDeploymentParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ServerlessV1Deployment, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageDeployment(ServiceSid, EnvironmentSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamDeployment(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamDeployment(response *ListDeploymentResponse, params *ListDeploymentParams, recordChannel chan ServerlessV1Deployment, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Deployments
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListDeploymentResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListDeploymentResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListDeploymentResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
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
