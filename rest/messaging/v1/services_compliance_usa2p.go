/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Messaging
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

// Optional parameters for the method 'CreateUsAppToPerson'
type CreateUsAppToPersonParams struct {
	// A2P Brand Registration SID
	BrandRegistrationSid *string `json:"BrandRegistrationSid,omitempty"`
	// A short description of what this SMS campaign does.
	Description *string `json:"Description,omitempty"`
	// Message samples, at least 2 and up to 5 sample messages, <=1024 chars each.
	MessageSamples *[]string `json:"MessageSamples,omitempty"`
	// A2P Campaign Use Case. Examples: [ 2FA, EMERGENCY, MARKETING..]
	UsAppToPersonUsecase *string `json:"UsAppToPersonUsecase,omitempty"`
	// Indicates that this SMS campaign will send messages that contain links.
	HasEmbeddedLinks *bool `json:"HasEmbeddedLinks,omitempty"`
	// Indicates that this SMS campaign will send messages that contain phone numbers.
	HasEmbeddedPhone *bool `json:"HasEmbeddedPhone,omitempty"`
}

func (params *CreateUsAppToPersonParams) SetBrandRegistrationSid(BrandRegistrationSid string) *CreateUsAppToPersonParams {
	params.BrandRegistrationSid = &BrandRegistrationSid
	return params
}
func (params *CreateUsAppToPersonParams) SetDescription(Description string) *CreateUsAppToPersonParams {
	params.Description = &Description
	return params
}
func (params *CreateUsAppToPersonParams) SetMessageSamples(MessageSamples []string) *CreateUsAppToPersonParams {
	params.MessageSamples = &MessageSamples
	return params
}
func (params *CreateUsAppToPersonParams) SetUsAppToPersonUsecase(UsAppToPersonUsecase string) *CreateUsAppToPersonParams {
	params.UsAppToPersonUsecase = &UsAppToPersonUsecase
	return params
}
func (params *CreateUsAppToPersonParams) SetHasEmbeddedLinks(HasEmbeddedLinks bool) *CreateUsAppToPersonParams {
	params.HasEmbeddedLinks = &HasEmbeddedLinks
	return params
}
func (params *CreateUsAppToPersonParams) SetHasEmbeddedPhone(HasEmbeddedPhone bool) *CreateUsAppToPersonParams {
	params.HasEmbeddedPhone = &HasEmbeddedPhone
	return params
}

func (c *ApiService) CreateUsAppToPerson(MessagingServiceSid string, params *CreateUsAppToPersonParams) (*MessagingV1UsAppToPerson, error) {
	return c.CreateUsAppToPersonWithCtx(context.TODO(), MessagingServiceSid, params)
}

func (c *ApiService) CreateUsAppToPersonWithCtx(ctx context.Context, MessagingServiceSid string, params *CreateUsAppToPersonParams) (*MessagingV1UsAppToPerson, error) {
	path := "/v1/Services/{MessagingServiceSid}/Compliance/Usa2p"
	path = strings.Replace(path, "{"+"MessagingServiceSid"+"}", MessagingServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.BrandRegistrationSid != nil {
		data.Set("BrandRegistrationSid", *params.BrandRegistrationSid)
	}
	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}
	if params != nil && params.MessageSamples != nil {
		for _, item := range *params.MessageSamples {
			data.Add("MessageSamples", item)
		}
	}
	if params != nil && params.UsAppToPersonUsecase != nil {
		data.Set("UsAppToPersonUsecase", *params.UsAppToPersonUsecase)
	}
	if params != nil && params.HasEmbeddedLinks != nil {
		data.Set("HasEmbeddedLinks", fmt.Sprint(*params.HasEmbeddedLinks))
	}
	if params != nil && params.HasEmbeddedPhone != nil {
		data.Set("HasEmbeddedPhone", fmt.Sprint(*params.HasEmbeddedPhone))
	}

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1UsAppToPerson{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteUsAppToPerson(MessagingServiceSid string, Sid string) error {
	return c.DeleteUsAppToPersonWithCtx(context.TODO(), MessagingServiceSid, Sid)
}

func (c *ApiService) DeleteUsAppToPersonWithCtx(ctx context.Context, MessagingServiceSid string, Sid string) error {
	path := "/v1/Services/{MessagingServiceSid}/Compliance/Usa2p/{Sid}"
	path = strings.Replace(path, "{"+"MessagingServiceSid"+"}", MessagingServiceSid, -1)
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

func (c *ApiService) FetchUsAppToPerson(MessagingServiceSid string, Sid string) (*MessagingV1UsAppToPerson, error) {
	return c.FetchUsAppToPersonWithCtx(context.TODO(), MessagingServiceSid, Sid)
}

func (c *ApiService) FetchUsAppToPersonWithCtx(ctx context.Context, MessagingServiceSid string, Sid string) (*MessagingV1UsAppToPerson, error) {
	path := "/v1/Services/{MessagingServiceSid}/Compliance/Usa2p/{Sid}"
	path = strings.Replace(path, "{"+"MessagingServiceSid"+"}", MessagingServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1UsAppToPerson{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListUsAppToPerson'
type ListUsAppToPersonParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListUsAppToPersonParams) SetPageSize(PageSize int) *ListUsAppToPersonParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListUsAppToPersonParams) SetLimit(Limit int) *ListUsAppToPersonParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of UsAppToPerson records from the API. Request is executed immediately.
func (c *ApiService) PageUsAppToPerson(MessagingServiceSid string, params *ListUsAppToPersonParams, pageToken, pageNumber string) (*ListUsAppToPersonResponse, error) {
	return c.PageUsAppToPersonWithCtx(context.TODO(), MessagingServiceSid, params, pageToken, pageNumber)
}

// Retrieve a single page of UsAppToPerson records from the API. Request is executed immediately.
func (c *ApiService) PageUsAppToPersonWithCtx(ctx context.Context, MessagingServiceSid string, params *ListUsAppToPersonParams, pageToken, pageNumber string) (*ListUsAppToPersonResponse, error) {
	path := "/v1/Services/{MessagingServiceSid}/Compliance/Usa2p"

	path = strings.Replace(path, "{"+"MessagingServiceSid"+"}", MessagingServiceSid, -1)

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

	ps := &ListUsAppToPersonResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists UsAppToPerson records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListUsAppToPerson(MessagingServiceSid string, params *ListUsAppToPersonParams) ([]MessagingV1UsAppToPerson, error) {
	return c.ListUsAppToPersonWithCtx(context.TODO(), MessagingServiceSid, params)
}

// Lists UsAppToPerson records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListUsAppToPersonWithCtx(ctx context.Context, MessagingServiceSid string, params *ListUsAppToPersonParams) ([]MessagingV1UsAppToPerson, error) {
	response, errors := c.StreamUsAppToPersonWithCtx(ctx, MessagingServiceSid, params)

	records := make([]MessagingV1UsAppToPerson, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams UsAppToPerson records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamUsAppToPerson(MessagingServiceSid string, params *ListUsAppToPersonParams) (chan MessagingV1UsAppToPerson, chan error) {
	return c.StreamUsAppToPersonWithCtx(context.TODO(), MessagingServiceSid, params)
}

// Streams UsAppToPerson records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamUsAppToPersonWithCtx(ctx context.Context, MessagingServiceSid string, params *ListUsAppToPersonParams) (chan MessagingV1UsAppToPerson, chan error) {
	if params == nil {
		params = &ListUsAppToPersonParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan MessagingV1UsAppToPerson, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageUsAppToPersonWithCtx(ctx, MessagingServiceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamUsAppToPerson(ctx, response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamUsAppToPerson(ctx context.Context, response *ListUsAppToPersonResponse, params *ListUsAppToPersonParams, recordChannel chan MessagingV1UsAppToPerson, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Compliance
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNextWithCtx(ctx, c.baseURL, response, c.getNextListUsAppToPersonResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListUsAppToPersonResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListUsAppToPersonResponse(ctx context.Context, nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(ctx, nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListUsAppToPersonResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
