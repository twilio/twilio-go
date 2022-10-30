/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Autopilot
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

// Optional parameters for the method 'CreateWebhook'
type CreateWebhookParams struct {
	// An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the `sid` in the URL path to address the resource. This value must be unique and 64 characters or less in length.
	UniqueName *string `json:"UniqueName,omitempty"`
	// The list of space-separated events that this Webhook will subscribe to.
	Events *string `json:"Events,omitempty"`
	// The URL associated with this Webhook.
	WebhookUrl *string `json:"WebhookUrl,omitempty"`
	// The method to be used when calling the webhook's URL.
	WebhookMethod *string `json:"WebhookMethod,omitempty"`
}

func (params *CreateWebhookParams) SetUniqueName(UniqueName string) *CreateWebhookParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *CreateWebhookParams) SetEvents(Events string) *CreateWebhookParams {
	params.Events = &Events
	return params
}
func (params *CreateWebhookParams) SetWebhookUrl(WebhookUrl string) *CreateWebhookParams {
	params.WebhookUrl = &WebhookUrl
	return params
}
func (params *CreateWebhookParams) SetWebhookMethod(WebhookMethod string) *CreateWebhookParams {
	params.WebhookMethod = &WebhookMethod
	return params
}

//
func (c *ApiService) CreateWebhook(AssistantSid string, params *CreateWebhookParams) (*AutopilotV1Webhook, error) {
	return c.CreateWebhookWithCtx(context.TODO(), AssistantSid, params)
}

//
func (c *ApiService) CreateWebhookWithCtx(ctx context.Context, AssistantSid string, params *CreateWebhookParams) (*AutopilotV1Webhook, error) {
	path := "/v1/Assistants/{AssistantSid}/Webhooks"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.Events != nil {
		data.Set("Events", *params.Events)
	}
	if params != nil && params.WebhookUrl != nil {
		data.Set("WebhookUrl", *params.WebhookUrl)
	}
	if params != nil && params.WebhookMethod != nil {
		data.Set("WebhookMethod", *params.WebhookMethod)
	}

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1Webhook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//
func (c *ApiService) DeleteWebhook(AssistantSid string, Sid string) error {
	return c.DeleteWebhookWithCtx(context.TODO(), AssistantSid, Sid)
}

//
func (c *ApiService) DeleteWebhookWithCtx(ctx context.Context, AssistantSid string, Sid string) error {
	path := "/v1/Assistants/{AssistantSid}/Webhooks/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
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

//
func (c *ApiService) FetchWebhook(AssistantSid string, Sid string) (*AutopilotV1Webhook, error) {
	return c.FetchWebhookWithCtx(context.TODO(), AssistantSid, Sid)
}

//
func (c *ApiService) FetchWebhookWithCtx(ctx context.Context, AssistantSid string, Sid string) (*AutopilotV1Webhook, error) {
	path := "/v1/Assistants/{AssistantSid}/Webhooks/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1Webhook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListWebhook'
type ListWebhookParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListWebhookParams) SetPageSize(PageSize int) *ListWebhookParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListWebhookParams) SetLimit(Limit int) *ListWebhookParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Webhook records from the API. Request is executed immediately.
func (c *ApiService) PageWebhook(AssistantSid string, params *ListWebhookParams, pageToken, pageNumber string) (*ListWebhookResponse, error) {
	return c.PageWebhookWithCtx(context.TODO(), AssistantSid, params, pageToken, pageNumber)
}

// Retrieve a single page of Webhook records from the API. Request is executed immediately.
func (c *ApiService) PageWebhookWithCtx(ctx context.Context, AssistantSid string, params *ListWebhookParams, pageToken, pageNumber string) (*ListWebhookResponse, error) {
	path := "/v1/Assistants/{AssistantSid}/Webhooks"

	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)

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

	ps := &ListWebhookResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Webhook records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListWebhook(AssistantSid string, params *ListWebhookParams) ([]AutopilotV1Webhook, error) {
	return c.ListWebhookWithCtx(context.TODO(), AssistantSid, params)
}

// Lists Webhook records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListWebhookWithCtx(ctx context.Context, AssistantSid string, params *ListWebhookParams) ([]AutopilotV1Webhook, error) {
	response, errors := c.StreamWebhookWithCtx(ctx, AssistantSid, params)

	records := make([]AutopilotV1Webhook, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Webhook records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamWebhook(AssistantSid string, params *ListWebhookParams) (chan AutopilotV1Webhook, chan error) {
	return c.StreamWebhookWithCtx(context.TODO(), AssistantSid, params)
}

// Streams Webhook records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamWebhookWithCtx(ctx context.Context, AssistantSid string, params *ListWebhookParams) (chan AutopilotV1Webhook, chan error) {
	if params == nil {
		params = &ListWebhookParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan AutopilotV1Webhook, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageWebhookWithCtx(ctx, AssistantSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamWebhook(ctx, response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamWebhook(ctx context.Context, response *ListWebhookResponse, params *ListWebhookParams, recordChannel chan AutopilotV1Webhook, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Webhooks
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNextWithCtx(ctx, c.baseURL, response, c.getNextListWebhookResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListWebhookResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListWebhookResponse(ctx context.Context, nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(ctx, nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListWebhookResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateWebhook'
type UpdateWebhookParams struct {
	// An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the `sid` in the URL path to address the resource. This value must be unique and 64 characters or less in length.
	UniqueName *string `json:"UniqueName,omitempty"`
	// The list of space-separated events that this Webhook will subscribe to.
	Events *string `json:"Events,omitempty"`
	// The URL associated with this Webhook.
	WebhookUrl *string `json:"WebhookUrl,omitempty"`
	// The method to be used when calling the webhook's URL.
	WebhookMethod *string `json:"WebhookMethod,omitempty"`
}

func (params *UpdateWebhookParams) SetUniqueName(UniqueName string) *UpdateWebhookParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *UpdateWebhookParams) SetEvents(Events string) *UpdateWebhookParams {
	params.Events = &Events
	return params
}
func (params *UpdateWebhookParams) SetWebhookUrl(WebhookUrl string) *UpdateWebhookParams {
	params.WebhookUrl = &WebhookUrl
	return params
}
func (params *UpdateWebhookParams) SetWebhookMethod(WebhookMethod string) *UpdateWebhookParams {
	params.WebhookMethod = &WebhookMethod
	return params
}

//
func (c *ApiService) UpdateWebhook(AssistantSid string, Sid string, params *UpdateWebhookParams) (*AutopilotV1Webhook, error) {
	return c.UpdateWebhookWithCtx(context.TODO(), AssistantSid, Sid, params)
}

//
func (c *ApiService) UpdateWebhookWithCtx(ctx context.Context, AssistantSid string, Sid string, params *UpdateWebhookParams) (*AutopilotV1Webhook, error) {
	path := "/v1/Assistants/{AssistantSid}/Webhooks/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.Events != nil {
		data.Set("Events", *params.Events)
	}
	if params != nil && params.WebhookUrl != nil {
		data.Set("WebhookUrl", *params.WebhookUrl)
	}
	if params != nil && params.WebhookMethod != nil {
		data.Set("WebhookMethod", *params.WebhookMethod)
	}

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1Webhook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
