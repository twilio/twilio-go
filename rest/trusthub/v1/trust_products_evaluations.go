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
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateTrustProductEvaluation'
type CreateTrustProductEvaluationParams struct {
	// The unique string of a policy that is associated to the customer_profile resource.
	PolicySid *string `json:"PolicySid,omitempty"`
}

func (params *CreateTrustProductEvaluationParams) SetPolicySid(PolicySid string) *CreateTrustProductEvaluationParams {
	params.PolicySid = &PolicySid
	return params
}

// Create a new Evaluation
func (c *ApiService) CreateTrustProductEvaluation(TrustProductSid string, params *CreateTrustProductEvaluationParams) (*TrusthubV1TrustProductEvaluation, error) {
	return c.CreateTrustProductEvaluationWithCtx(context.TODO(), TrustProductSid, params)
}

// Create a new Evaluation
func (c *ApiService) CreateTrustProductEvaluationWithCtx(ctx context.Context, TrustProductSid string, params *CreateTrustProductEvaluationParams) (*TrusthubV1TrustProductEvaluation, error) {
	path := "/v1/TrustProducts/{TrustProductSid}/Evaluations"
	path = strings.Replace(path, "{"+"TrustProductSid"+"}", TrustProductSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PolicySid != nil {
		data.Set("PolicySid", *params.PolicySid)
	}

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1TrustProductEvaluation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Fetch specific Evaluation Instance.
func (c *ApiService) FetchTrustProductEvaluation(TrustProductSid string, Sid string) (*TrusthubV1TrustProductEvaluation, error) {
	return c.FetchTrustProductEvaluationWithCtx(context.TODO(), TrustProductSid, Sid)
}

// Fetch specific Evaluation Instance.
func (c *ApiService) FetchTrustProductEvaluationWithCtx(ctx context.Context, TrustProductSid string, Sid string) (*TrusthubV1TrustProductEvaluation, error) {
	path := "/v1/TrustProducts/{TrustProductSid}/Evaluations/{Sid}"
	path = strings.Replace(path, "{"+"TrustProductSid"+"}", TrustProductSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1TrustProductEvaluation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListTrustProductEvaluation'
type ListTrustProductEvaluationParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListTrustProductEvaluationParams) SetPageSize(PageSize int) *ListTrustProductEvaluationParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListTrustProductEvaluationParams) SetLimit(Limit int) *ListTrustProductEvaluationParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of TrustProductEvaluation records from the API. Request is executed immediately.
func (c *ApiService) PageTrustProductEvaluation(TrustProductSid string, params *ListTrustProductEvaluationParams, pageToken, pageNumber string) (*ListTrustProductEvaluationResponse, error) {
	return c.PageTrustProductEvaluationWithCtx(context.TODO(), TrustProductSid, params, pageToken, pageNumber)
}

// Retrieve a single page of TrustProductEvaluation records from the API. Request is executed immediately.
func (c *ApiService) PageTrustProductEvaluationWithCtx(ctx context.Context, TrustProductSid string, params *ListTrustProductEvaluationParams, pageToken, pageNumber string) (*ListTrustProductEvaluationResponse, error) {
	path := "/v1/TrustProducts/{TrustProductSid}/Evaluations"

	path = strings.Replace(path, "{"+"TrustProductSid"+"}", TrustProductSid, -1)

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

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTrustProductEvaluationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists TrustProductEvaluation records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListTrustProductEvaluation(TrustProductSid string, params *ListTrustProductEvaluationParams) ([]TrusthubV1TrustProductEvaluation, error) {
	return c.ListTrustProductEvaluationWithCtx(context.TODO(), TrustProductSid, params)
}

// Lists TrustProductEvaluation records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListTrustProductEvaluationWithCtx(ctx context.Context, TrustProductSid string, params *ListTrustProductEvaluationParams) ([]TrusthubV1TrustProductEvaluation, error) {
	response, errors := c.StreamTrustProductEvaluationWithCtx(ctx, TrustProductSid, params)

	records := make([]TrusthubV1TrustProductEvaluation, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams TrustProductEvaluation records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamTrustProductEvaluation(TrustProductSid string, params *ListTrustProductEvaluationParams) (chan TrusthubV1TrustProductEvaluation, chan error) {
	return c.StreamTrustProductEvaluationWithCtx(context.TODO(), TrustProductSid, params)
}

// Streams TrustProductEvaluation records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamTrustProductEvaluationWithCtx(ctx context.Context, TrustProductSid string, params *ListTrustProductEvaluationParams) (chan TrusthubV1TrustProductEvaluation, chan error) {
	if params == nil {
		params = &ListTrustProductEvaluationParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan TrusthubV1TrustProductEvaluation, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageTrustProductEvaluationWithCtx(ctx, TrustProductSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamTrustProductEvaluation(ctx, response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamTrustProductEvaluation(ctx context.Context, response *ListTrustProductEvaluationResponse, params *ListTrustProductEvaluationParams, recordChannel chan TrusthubV1TrustProductEvaluation, errorChannel chan error) {
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

		record, err := client.GetNextWithCtx(ctx, c.baseURL, response, c.getNextListTrustProductEvaluationResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListTrustProductEvaluationResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListTrustProductEvaluationResponse(ctx context.Context, nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(ctx, nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTrustProductEvaluationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
