/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"fmt"
	twilio "github.com/twilio/twilio-go/client"
	"net/url"
)

type DefaultApiService struct {
	baseURL string
	client  *twilio.Client
}

func NewDefaultApiService(client *twilio.Client) *DefaultApiService {
	return &DefaultApiService {
		client: client,
		baseURL: fmt.Sprintf("https://studio.%s", client.BaseURL),
	}
}

/*
FetchMessagingCountry Method for FetchMessagingCountry
 * @param isoCountry The [ISO country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the pricing information to fetch.
@return PricingV1MessagingMessagingCountryInstance
*/
func (c *DefaultApiService) FetchMessagingCountry(isoCountry string) (*PricingV1MessagingMessagingCountryInstance, error) {
	path := "/v1/Messaging/Countries/{IsoCountry}"
	path = strings.Replace(path, "{"+"IsoCountry"+"}", isoCountry, -1)

	data := 0
	headers := 0



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
FetchPhoneNumberCountry Method for FetchPhoneNumberCountry
 * @param isoCountry The [ISO country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the pricing information to fetch.
@return PricingV1PhoneNumberPhoneNumberCountryInstance
*/
func (c *DefaultApiService) FetchPhoneNumberCountry(isoCountry string) (*PricingV1PhoneNumberPhoneNumberCountryInstance, error) {
	path := "/v1/PhoneNumbers/Countries/{IsoCountry}"
	path = strings.Replace(path, "{"+"IsoCountry"+"}", isoCountry, -1)

	data := 0
	headers := 0



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
FetchVoiceCountry Method for FetchVoiceCountry
 * @param isoCountry The [ISO country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the pricing information to fetch.
@return PricingV1VoiceVoiceCountryInstance
*/
func (c *DefaultApiService) FetchVoiceCountry(isoCountry string) (*PricingV1VoiceVoiceCountryInstance, error) {
	path := "/v1/Voice/Countries/{IsoCountry}"
	path = strings.Replace(path, "{"+"IsoCountry"+"}", isoCountry, -1)

	data := 0
	headers := 0



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
FetchVoiceNumber Method for FetchVoiceNumber
 * @param number The phone number to fetch.
@return PricingV1VoiceVoiceNumber
*/
func (c *DefaultApiService) FetchVoiceNumber(number string) (*PricingV1VoiceVoiceNumber, error) {
	path := "/v1/Voice/Numbers/{Number}"
	path = strings.Replace(path, "{"+"Number"+"}", number, -1)

	data := 0
	headers := 0



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
ListMessagingCountry Method for ListMessagingCountry
 * @param optional nil or *ListMessagingCountryOpts - Optional Parameters:
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return PricingV1MessagingMessagingCountryReadResponse
*/
func (c *DefaultApiService) ListMessagingCountry(params *ListMessagingCountryParams) (*PricingV1MessagingMessagingCountryReadResponse, error) {
	path := "/v1/Messaging/Countries"

	data := url.Values{}
	headers := 0

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", string(*params.PageSize))
	}


	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &PricingV1MessagingMessagingCountryReadResponse{}
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
ListPhoneNumberCountry Method for ListPhoneNumberCountry
 * @param optional nil or *ListPhoneNumberCountryOpts - Optional Parameters:
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return PricingV1PhoneNumberPhoneNumberCountryReadResponse
*/
func (c *DefaultApiService) ListPhoneNumberCountry(params *ListPhoneNumberCountryParams) (*PricingV1PhoneNumberPhoneNumberCountryReadResponse, error) {
	path := "/v1/PhoneNumbers/Countries"

	data := url.Values{}
	headers := 0

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", string(*params.PageSize))
	}


	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &PricingV1PhoneNumberPhoneNumberCountryReadResponse{}
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
ListVoiceCountry Method for ListVoiceCountry
 * @param optional nil or *ListVoiceCountryOpts - Optional Parameters:
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return PricingV1VoiceVoiceCountryReadResponse
*/
func (c *DefaultApiService) ListVoiceCountry(params *ListVoiceCountryParams) (*PricingV1VoiceVoiceCountryReadResponse, error) {
	path := "/v1/Voice/Countries"

	data := url.Values{}
	headers := 0

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", string(*params.PageSize))
	}


	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &PricingV1VoiceVoiceCountryReadResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
