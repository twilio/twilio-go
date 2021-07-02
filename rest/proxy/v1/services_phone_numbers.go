/*
 * Twilio - Proxy
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

// Optional parameters for the method 'CreatePhoneNumber'
type CreatePhoneNumberParams struct {
	// Whether the new phone number should be reserved and not be assigned to a participant using proxy pool logic. See [Reserved Phone Numbers](https://www.twilio.com/docs/proxy/reserved-phone-numbers) for more information.
	IsReserved *bool `json:"IsReserved,omitempty"`
	// The phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234.
	PhoneNumber *string `json:"PhoneNumber,omitempty"`
	// The SID of a Twilio [IncomingPhoneNumber](https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource) resource that represents the Twilio Number you would like to assign to your Proxy Service.
	Sid *string `json:"Sid,omitempty"`
}

func (params *CreatePhoneNumberParams) SetIsReserved(IsReserved bool) *CreatePhoneNumberParams {
	params.IsReserved = &IsReserved
	return params
}
func (params *CreatePhoneNumberParams) SetPhoneNumber(PhoneNumber string) *CreatePhoneNumberParams {
	params.PhoneNumber = &PhoneNumber
	return params
}
func (params *CreatePhoneNumberParams) SetSid(Sid string) *CreatePhoneNumberParams {
	params.Sid = &Sid
	return params
}

// Add a Phone Number to a Service&#39;s Proxy Number Pool.
func (c *ApiService) CreatePhoneNumber(ServiceSid string, params *CreatePhoneNumberParams) (*ProxyV1ServicePhoneNumber, error) {
	path := "/v1/Services/{ServiceSid}/PhoneNumbers"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.IsReserved != nil {
		data.Set("IsReserved", fmt.Sprint(*params.IsReserved))
	}
	if params != nil && params.PhoneNumber != nil {
		data.Set("PhoneNumber", *params.PhoneNumber)
	}
	if params != nil && params.Sid != nil {
		data.Set("Sid", *params.Sid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyV1ServicePhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Phone Number from a Service.
func (c *ApiService) DeletePhoneNumber(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/PhoneNumbers/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
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

// Fetch a specific Phone Number.
func (c *ApiService) FetchPhoneNumber(ServiceSid string, Sid string) (*ProxyV1ServicePhoneNumber, error) {
	path := "/v1/Services/{ServiceSid}/PhoneNumbers/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyV1ServicePhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListPhoneNumber'
type ListPhoneNumberParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListPhoneNumberParams) SetPageSize(PageSize int) *ListPhoneNumberParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all Phone Numbers in the Proxy Number Pool for a Service. A maximum of 100 records will be returned per page.
func (c *ApiService) ListPhoneNumber(ServiceSid string, params *ListPhoneNumberParams) (*ListPhoneNumberResponse, error) {
	path := "/v1/Services/{ServiceSid}/PhoneNumbers"
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

	ps := &ListPhoneNumberResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of PhoneNumber records from the API. Request is executed immediately.
func (c *ApiService) PhoneNumberPage(ServiceSid string, params *ListPhoneNumberParams, pageToken string, pageNumber string) (*client.Page, error) {
	path := "/v1/Services/{ServiceSid}/PhoneNumbers"
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

//Streams PhoneNumber records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) PhoneNumberStream(ServiceSid string, params *ListPhoneNumberParams, limit int) (chan map[string]interface{}, error) {
	if params.PageSize == nil {
		params.SetPageSize(0)
	}
	params.SetPageSize(c.requestHandler.ReadLimits(*params.PageSize, limit))
	page, err := c.PhoneNumberPage(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.Stream(page, limit, 0), nil
}

//Lists PhoneNumber records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) PhoneNumberList(ServiceSid string, params *ListPhoneNumberParams, limit int) ([]interface{}, error) {
	if params.PageSize == nil {
		params.SetPageSize(0)
	}
	params.SetPageSize(c.requestHandler.ReadLimits(*params.PageSize, limit))
	page, err := c.PhoneNumberPage(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.List(page, limit, 0), nil
}

// Optional parameters for the method 'UpdatePhoneNumber'
type UpdatePhoneNumberParams struct {
	// Whether the phone number should be reserved and not be assigned to a participant using proxy pool logic. See [Reserved Phone Numbers](https://www.twilio.com/docs/proxy/reserved-phone-numbers) for more information.
	IsReserved *bool `json:"IsReserved,omitempty"`
}

func (params *UpdatePhoneNumberParams) SetIsReserved(IsReserved bool) *UpdatePhoneNumberParams {
	params.IsReserved = &IsReserved
	return params
}

// Update a specific Proxy Number.
func (c *ApiService) UpdatePhoneNumber(ServiceSid string, Sid string, params *UpdatePhoneNumberParams) (*ProxyV1ServicePhoneNumber, error) {
	path := "/v1/Services/{ServiceSid}/PhoneNumbers/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.IsReserved != nil {
		data.Set("IsReserved", fmt.Sprint(*params.IsReserved))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyV1ServicePhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
