/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Numbers
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

// Creates an evaluation for a bundle
func (c *ApiService) CreateEvaluation(BundleSid string) (*NumbersV2Evaluation, error) {
	path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/Evaluations"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2Evaluation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Fetch specific Evaluation Instance.
func (c *ApiService) FetchEvaluation(BundleSid string, Sid string) (*NumbersV2Evaluation, error) {
	path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/Evaluations/{Sid}"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)
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

	ps := &NumbersV2Evaluation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListEvaluation'
type ListEvaluationParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListEvaluationParams) SetPageSize(PageSize int) *ListEvaluationParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListEvaluationParams) SetLimit(Limit int) *ListEvaluationParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Evaluation records from the API. Request is executed immediately.
func (c *ApiService) PageEvaluation(BundleSid string, params *ListEvaluationParams, pageToken, pageNumber string) (*ListEvaluationResponse, error) {
	path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/Evaluations"

	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)

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

	ps := &ListEvaluationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Evaluation records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListEvaluation(BundleSid string, params *ListEvaluationParams) ([]NumbersV2Evaluation, error) {
	response, errors := c.StreamEvaluation(BundleSid, params)

	records := make([]NumbersV2Evaluation, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Evaluation records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamEvaluation(BundleSid string, params *ListEvaluationParams) (chan NumbersV2Evaluation, chan error) {
	if params == nil {
		params = &ListEvaluationParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan NumbersV2Evaluation, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageEvaluation(BundleSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamEvaluation(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamEvaluation(response *ListEvaluationResponse, params *ListEvaluationParams, recordChannel chan NumbersV2Evaluation, errorChannel chan error) {
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

		record, err := client.GetNext(c.baseURL, response, c.getNextListEvaluationResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListEvaluationResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListEvaluationResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListEvaluationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
