/*
 * Twilio - Verify
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

// Optional parameters for the method 'CreateMessagingConfiguration'
type CreateMessagingConfigurationParams struct {
	// The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country this configuration will be applied to. If this is a global configuration, Country will take the value `all`.
	Country *string `json:"Country,omitempty"`
	// The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) to be used to send SMS to the country of this configuration.
	MessagingServiceSid *string `json:"MessagingServiceSid,omitempty"`
}

func (params *CreateMessagingConfigurationParams) SetCountry(Country string) *CreateMessagingConfigurationParams {
	params.Country = &Country
	return params
}
func (params *CreateMessagingConfigurationParams) SetMessagingServiceSid(MessagingServiceSid string) *CreateMessagingConfigurationParams {
	params.MessagingServiceSid = &MessagingServiceSid
	return params
}

// Create a new MessagingConfiguration for a service.
func (c *ApiService) CreateMessagingConfiguration(ServiceSid string, params *CreateMessagingConfigurationParams) (*VerifyV2ServiceMessagingConfiguration, error) {
	path := "/v2/Services/{ServiceSid}/MessagingConfigurations"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Country != nil {
		data.Set("Country", *params.Country)
	}
	if params != nil && params.MessagingServiceSid != nil {
		data.Set("MessagingServiceSid", *params.MessagingServiceSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2ServiceMessagingConfiguration{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific MessagingConfiguration.
func (c *ApiService) DeleteMessagingConfiguration(ServiceSid string, Country string) error {
	path := "/v2/Services/{ServiceSid}/MessagingConfigurations/{Country}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Country"+"}", Country, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch a specific MessagingConfiguration.
func (c *ApiService) FetchMessagingConfiguration(ServiceSid string, Country string) (*VerifyV2ServiceMessagingConfiguration, error) {
	path := "/v2/Services/{ServiceSid}/MessagingConfigurations/{Country}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Country"+"}", Country, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2ServiceMessagingConfiguration{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListMessagingConfiguration'
type ListMessagingConfigurationParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListMessagingConfigurationParams) SetPageSize(PageSize int) *ListMessagingConfigurationParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all Messaging Configurations for a Service.
func (c *ApiService) ListMessagingConfiguration(ServiceSid string, params *ListMessagingConfigurationParams) (*ListMessagingConfigurationResponse, error) {
	path := "/v2/Services/{ServiceSid}/MessagingConfigurations"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

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

	ps := &ListMessagingConfigurationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of MessagingConfiguration records from the API. Request is executed immediately.
func (c *ApiService) MessagingConfigurationPage(ServiceSid string, params *ListMessagingConfigurationParams, pageToken string, pageNumber string) (*client.Page, error) {
	path := "/v2/Services/{ServiceSid}/MessagingConfigurations"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

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

//Streams MessagingConfiguration records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) MessagingConfigurationStream(ServiceSid string, params *ListMessagingConfigurationParams, limit int) (chan map[string]interface{}, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	page, err := c.MessagingConfigurationPage(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.Stream(page, limit, 0), nil
}

//Lists MessagingConfiguration records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) MessagingConfigurationList(ServiceSid string, params *ListMessagingConfigurationParams, limit int) ([]interface{}, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	page, err := c.MessagingConfigurationPage(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.List(page, limit, 0), nil
}

// Optional parameters for the method 'UpdateMessagingConfiguration'
type UpdateMessagingConfigurationParams struct {
	// The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) to be used to send SMS to the country of this configuration.
	MessagingServiceSid *string `json:"MessagingServiceSid,omitempty"`
}

func (params *UpdateMessagingConfigurationParams) SetMessagingServiceSid(MessagingServiceSid string) *UpdateMessagingConfigurationParams {
	params.MessagingServiceSid = &MessagingServiceSid
	return params
}

// Update a specific MessagingConfiguration
func (c *ApiService) UpdateMessagingConfiguration(ServiceSid string, Country string, params *UpdateMessagingConfigurationParams) (*VerifyV2ServiceMessagingConfiguration, error) {
	path := "/v2/Services/{ServiceSid}/MessagingConfigurations/{Country}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Country"+"}", Country, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.MessagingServiceSid != nil {
		data.Set("MessagingServiceSid", *params.MessagingServiceSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2ServiceMessagingConfiguration{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
