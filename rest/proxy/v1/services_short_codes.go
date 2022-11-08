/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Proxy
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

// Optional parameters for the method 'CreateShortCode'
type CreateShortCodeParams struct {
	// The SID of a Twilio [ShortCode](https://www.twilio.com/docs/sms/api/short-code) resource that represents the short code you would like to assign to your Proxy Service.
	Sid *string `json:"Sid,omitempty"`
}

func (params *CreateShortCodeParams) SetSid(Sid string) *CreateShortCodeParams {
	params.Sid = &Sid
	return params
}

// Add a Short Code to the Proxy Number Pool for the Service.
func (c *ApiService) CreateShortCode(ServiceSid string, params *CreateShortCodeParams) (*ProxyV1ShortCode, error) {
	return c.CreateShortCodeWithCtx(context.TODO(), ServiceSid, params)
}

// Add a Short Code to the Proxy Number Pool for the Service.
func (c *ApiService) CreateShortCodeWithCtx(ctx context.Context, ServiceSid string, params *CreateShortCodeParams) (*ProxyV1ShortCode, error) {
	path := "/v1/Services/{ServiceSid}/ShortCodes"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Sid != nil {
		data.Set("Sid", *params.Sid)
	}

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyV1ShortCode{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Short Code from a Service.
func (c *ApiService) DeleteShortCode(ServiceSid string, Sid string) error {
	return c.DeleteShortCodeWithCtx(context.TODO(), ServiceSid, Sid)
}

// Delete a specific Short Code from a Service.
func (c *ApiService) DeleteShortCodeWithCtx(ctx context.Context, ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/ShortCodes/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch a specific Short Code.
func (c *ApiService) FetchShortCode(ServiceSid string, Sid string) (*ProxyV1ShortCode, error) {
	return c.FetchShortCodeWithCtx(context.TODO(), ServiceSid, Sid)
}

// Fetch a specific Short Code.
func (c *ApiService) FetchShortCodeWithCtx(ctx context.Context, ServiceSid string, Sid string) (*ProxyV1ShortCode, error) {
	path := "/v1/Services/{ServiceSid}/ShortCodes/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyV1ShortCode{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListShortCode'
type ListShortCodeParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListShortCodeParams) SetPageSize(PageSize int) *ListShortCodeParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListShortCodeParams) SetLimit(Limit int) *ListShortCodeParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ShortCode records from the API. Request is executed immediately.
func (c *ApiService) PageShortCode(ServiceSid string, params *ListShortCodeParams, pageToken, pageNumber string) (*ListShortCodeResponse, error) {
	return c.PageShortCodeWithCtx(context.TODO(), ServiceSid, params, pageToken, pageNumber)
}

// Retrieve a single page of ShortCode records from the API. Request is executed immediately.
func (c *ApiService) PageShortCodeWithCtx(ctx context.Context, ServiceSid string, params *ListShortCodeParams, pageToken, pageNumber string) (*ListShortCodeResponse, error) {
	path := "/v1/Services/{ServiceSid}/ShortCodes"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

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

	ps := &ListShortCodeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ShortCode records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListShortCode(ServiceSid string, params *ListShortCodeParams) ([]ProxyV1ShortCode, error) {
	return c.ListShortCodeWithCtx(context.TODO(), ServiceSid, params)
}

// Lists ShortCode records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListShortCodeWithCtx(ctx context.Context, ServiceSid string, params *ListShortCodeParams) ([]ProxyV1ShortCode, error) {
	response, errors := c.StreamShortCodeWithCtx(ctx, ServiceSid, params)

	records := make([]ProxyV1ShortCode, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams ShortCode records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamShortCode(ServiceSid string, params *ListShortCodeParams) (chan ProxyV1ShortCode, chan error) {
	return c.StreamShortCodeWithCtx(context.TODO(), ServiceSid, params)
}

// Streams ShortCode records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamShortCodeWithCtx(ctx context.Context, ServiceSid string, params *ListShortCodeParams) (chan ProxyV1ShortCode, chan error) {
	if params == nil {
		params = &ListShortCodeParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ProxyV1ShortCode, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageShortCodeWithCtx(ctx, ServiceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamShortCode(ctx, response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamShortCode(ctx context.Context, response *ListShortCodeResponse, params *ListShortCodeParams, recordChannel chan ProxyV1ShortCode, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.ShortCodes
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNextWithCtx(ctx, c.baseURL, response, c.getNextListShortCodeResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListShortCodeResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListShortCodeResponse(ctx context.Context, nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(ctx, nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListShortCodeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateShortCode'
type UpdateShortCodeParams struct {
	// Whether the short code should be reserved and not be assigned to a participant using proxy pool logic. See [Reserved Phone Numbers](https://www.twilio.com/docs/proxy/reserved-phone-numbers) for more information.
	IsReserved *bool `json:"IsReserved,omitempty"`
}

func (params *UpdateShortCodeParams) SetIsReserved(IsReserved bool) *UpdateShortCodeParams {
	params.IsReserved = &IsReserved
	return params
}

// Update a specific Short Code.
func (c *ApiService) UpdateShortCode(ServiceSid string, Sid string, params *UpdateShortCodeParams) (*ProxyV1ShortCode, error) {
	return c.UpdateShortCodeWithCtx(context.TODO(), ServiceSid, Sid, params)
}

// Update a specific Short Code.
func (c *ApiService) UpdateShortCodeWithCtx(ctx context.Context, ServiceSid string, Sid string, params *UpdateShortCodeParams) (*ProxyV1ShortCode, error) {
	path := "/v1/Services/{ServiceSid}/ShortCodes/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.IsReserved != nil {
		data.Set("IsReserved", fmt.Sprint(*params.IsReserved))
	}

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyV1ShortCode{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
