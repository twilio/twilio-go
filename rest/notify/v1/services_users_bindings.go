/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Notify
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

// Optional parameters for the method 'CreateUserBinding'
type CreateUserBindingParams struct {
	//
	BindingType *string `json:"BindingType,omitempty"`
	//
	Address *string `json:"Address,omitempty"`
	//
	Tag *[]string `json:"Tag,omitempty"`
	//
	NotificationProtocolVersion *string `json:"NotificationProtocolVersion,omitempty"`
	//
	CredentialSid *string `json:"CredentialSid,omitempty"`
	//
	Endpoint *string `json:"Endpoint,omitempty"`
}

func (params *CreateUserBindingParams) SetBindingType(BindingType string) *CreateUserBindingParams {
	params.BindingType = &BindingType
	return params
}
func (params *CreateUserBindingParams) SetAddress(Address string) *CreateUserBindingParams {
	params.Address = &Address
	return params
}
func (params *CreateUserBindingParams) SetTag(Tag []string) *CreateUserBindingParams {
	params.Tag = &Tag
	return params
}
func (params *CreateUserBindingParams) SetNotificationProtocolVersion(NotificationProtocolVersion string) *CreateUserBindingParams {
	params.NotificationProtocolVersion = &NotificationProtocolVersion
	return params
}
func (params *CreateUserBindingParams) SetCredentialSid(CredentialSid string) *CreateUserBindingParams {
	params.CredentialSid = &CredentialSid
	return params
}
func (params *CreateUserBindingParams) SetEndpoint(Endpoint string) *CreateUserBindingParams {
	params.Endpoint = &Endpoint
	return params
}

//
func (c *ApiService) CreateUserBinding(ServiceSid string, Identity string, params *CreateUserBindingParams) (*NotifyV1UserBinding, error) {
	path := "/v1/Services/{ServiceSid}/Users/{Identity}/Bindings"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.BindingType != nil {
		data.Set("BindingType", *params.BindingType)
	}
	if params != nil && params.Address != nil {
		data.Set("Address", *params.Address)
	}
	if params != nil && params.Tag != nil {
		for _, item := range *params.Tag {
			data.Add("Tag", item)
		}
	}
	if params != nil && params.NotificationProtocolVersion != nil {
		data.Set("NotificationProtocolVersion", *params.NotificationProtocolVersion)
	}
	if params != nil && params.CredentialSid != nil {
		data.Set("CredentialSid", *params.CredentialSid)
	}
	if params != nil && params.Endpoint != nil {
		data.Set("Endpoint", *params.Endpoint)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NotifyV1UserBinding{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//
func (c *ApiService) DeleteUserBinding(ServiceSid string, Identity string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Users/{Identity}/Bindings/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)
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

//
func (c *ApiService) FetchUserBinding(ServiceSid string, Identity string, Sid string) (*NotifyV1UserBinding, error) {
	path := "/v1/Services/{ServiceSid}/Users/{Identity}/Bindings/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NotifyV1UserBinding{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListUserBinding'
type ListUserBindingParams struct {
	//
	StartDate *string `json:"StartDate,omitempty"`
	//
	EndDate *string `json:"EndDate,omitempty"`
	//
	Tag *[]string `json:"Tag,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListUserBindingParams) SetStartDate(StartDate string) *ListUserBindingParams {
	params.StartDate = &StartDate
	return params
}
func (params *ListUserBindingParams) SetEndDate(EndDate string) *ListUserBindingParams {
	params.EndDate = &EndDate
	return params
}
func (params *ListUserBindingParams) SetTag(Tag []string) *ListUserBindingParams {
	params.Tag = &Tag
	return params
}
func (params *ListUserBindingParams) SetPageSize(PageSize int) *ListUserBindingParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListUserBindingParams) SetLimit(Limit int) *ListUserBindingParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of UserBinding records from the API. Request is executed immediately.
func (c *ApiService) PageUserBinding(ServiceSid string, Identity string, params *ListUserBindingParams, pageToken, pageNumber string) (*ListUserBindingResponse, error) {
	path := "/v1/Services/{ServiceSid}/Users/{Identity}/Bindings"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.StartDate != nil {
		data.Set("StartDate", fmt.Sprint(*params.StartDate))
	}
	if params != nil && params.EndDate != nil {
		data.Set("EndDate", fmt.Sprint(*params.EndDate))
	}
	if params != nil && params.Tag != nil {
		for _, item := range *params.Tag {
			data.Add("Tag", item)
		}
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

	ps := &ListUserBindingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists UserBinding records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListUserBinding(ServiceSid string, Identity string, params *ListUserBindingParams) ([]NotifyV1UserBinding, error) {
	response, errors := c.StreamUserBinding(ServiceSid, Identity, params)

	records := make([]NotifyV1UserBinding, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams UserBinding records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamUserBinding(ServiceSid string, Identity string, params *ListUserBindingParams) (chan NotifyV1UserBinding, chan error) {
	if params == nil {
		params = &ListUserBindingParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan NotifyV1UserBinding, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageUserBinding(ServiceSid, Identity, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamUserBinding(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamUserBinding(response *ListUserBindingResponse, params *ListUserBindingParams, recordChannel chan NotifyV1UserBinding, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Bindings
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListUserBindingResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListUserBindingResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListUserBindingResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListUserBindingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
