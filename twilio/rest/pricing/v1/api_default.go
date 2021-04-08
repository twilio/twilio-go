/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.13.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	twilio "github.com/twilio/twilio-go/client"
)

type DefaultApiService struct {
	baseURL string
	client  twilio.BaseClient
}

func NewDefaultApiService(client twilio.BaseClient) *DefaultApiService {
	return &DefaultApiService{
		client:  client,
		baseURL: "https://pricing.twilio.com",
	}
}

/*
* FetchMessagingCountry Method for FetchMessagingCountry
* @param IsoCountry The [ISO country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the pricing information to fetch.
* @return PricingV1MessagingMessagingCountryInstance
 */
func (c *DefaultApiService) FetchMessagingCountry(IsoCountry string) (*PricingV1MessagingMessagingCountryInstance, error) {
	path := "/v1/Messaging/Countries/{IsoCountry}"
	path = strings.Replace(path, "{"+"IsoCountry"+"}", IsoCountry, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &PricingV1MessagingMessagingCountryInstance{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
* FetchPhoneNumberCountry Method for FetchPhoneNumberCountry
* @param IsoCountry The [ISO country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the pricing information to fetch.
* @return PricingV1PhoneNumberPhoneNumberCountryInstance
 */
func (c *DefaultApiService) FetchPhoneNumberCountry(IsoCountry string) (*PricingV1PhoneNumberPhoneNumberCountryInstance, error) {
	path := "/v1/PhoneNumbers/Countries/{IsoCountry}"
	path = strings.Replace(path, "{"+"IsoCountry"+"}", IsoCountry, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &PricingV1PhoneNumberPhoneNumberCountryInstance{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
* FetchVoiceCountry Method for FetchVoiceCountry
* @param IsoCountry The [ISO country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the pricing information to fetch.
* @return PricingV1VoiceVoiceCountryInstance
 */
func (c *DefaultApiService) FetchVoiceCountry(IsoCountry string) (*PricingV1VoiceVoiceCountryInstance, error) {
	path := "/v1/Voice/Countries/{IsoCountry}"
	path = strings.Replace(path, "{"+"IsoCountry"+"}", IsoCountry, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &PricingV1VoiceVoiceCountryInstance{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
* FetchVoiceNumber Method for FetchVoiceNumber
* @param Number The phone number to fetch.
* @return PricingV1VoiceVoiceNumber
 */
func (c *DefaultApiService) FetchVoiceNumber(Number string) (*PricingV1VoiceVoiceNumber, error) {
	path := "/v1/Voice/Numbers/{Number}"
	path = strings.Replace(path, "{"+"Number"+"}", Number, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &PricingV1VoiceVoiceNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListMessagingCountryParams Optional parameters for the method 'ListMessagingCountry'
type ListMessagingCountryParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListMessagingCountry Method for ListMessagingCountry
* @param optional nil or *ListMessagingCountryParams - Optional Parameters:
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListMessagingCountryResponse
 */
func (c *DefaultApiService) ListMessagingCountry(params *ListMessagingCountryParams) (*ListMessagingCountryResponse, error) {
	path := "/v1/Messaging/Countries"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMessagingCountryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListPhoneNumberCountryParams Optional parameters for the method 'ListPhoneNumberCountry'
type ListPhoneNumberCountryParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListPhoneNumberCountry Method for ListPhoneNumberCountry
* @param optional nil or *ListPhoneNumberCountryParams - Optional Parameters:
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListPhoneNumberCountryResponse
 */
func (c *DefaultApiService) ListPhoneNumberCountry(params *ListPhoneNumberCountryParams) (*ListPhoneNumberCountryResponse, error) {
	path := "/v1/PhoneNumbers/Countries"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListPhoneNumberCountryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListVoiceCountryParams Optional parameters for the method 'ListVoiceCountry'
type ListVoiceCountryParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListVoiceCountry Method for ListVoiceCountry
* @param optional nil or *ListVoiceCountryParams - Optional Parameters:
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListVoiceCountryResponse
 */
func (c *DefaultApiService) ListVoiceCountry(params *ListVoiceCountryParams) (*ListVoiceCountryResponse, error) {
	path := "/v1/Voice/Countries"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListVoiceCountryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
