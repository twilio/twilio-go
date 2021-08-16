/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.0
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

// Optional parameters for the method 'CreateAddress'
type CreateAddressParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Address resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// Whether we should automatically correct the address. Can be: `true` or `false` and the default is `true`. If empty or `true`, we will correct the address you provide if necessary. If `false`, we won't alter the address you provide.
	AutoCorrectAddress *bool `json:"AutoCorrectAddress,omitempty"`
	// The city of the new address.
	City *string `json:"City,omitempty"`
	// The name to associate with the new address.
	CustomerName *string `json:"CustomerName,omitempty"`
	// Whether to enable emergency calling on the new address. Can be: `true` or `false`.
	EmergencyEnabled *bool `json:"EmergencyEnabled,omitempty"`
	// A descriptive string that you create to describe the new address. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The ISO country code of the new address.
	IsoCountry *string `json:"IsoCountry,omitempty"`
	// The postal code of the new address.
	PostalCode *string `json:"PostalCode,omitempty"`
	// The state or region of the new address.
	Region *string `json:"Region,omitempty"`
	// The number and street address of the new address.
	Street *string `json:"Street,omitempty"`
}

func (params *CreateAddressParams) SetPathAccountSid(PathAccountSid string) *CreateAddressParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateAddressParams) SetAutoCorrectAddress(AutoCorrectAddress bool) *CreateAddressParams {
	params.AutoCorrectAddress = &AutoCorrectAddress
	return params
}
func (params *CreateAddressParams) SetCity(City string) *CreateAddressParams {
	params.City = &City
	return params
}
func (params *CreateAddressParams) SetCustomerName(CustomerName string) *CreateAddressParams {
	params.CustomerName = &CustomerName
	return params
}
func (params *CreateAddressParams) SetEmergencyEnabled(EmergencyEnabled bool) *CreateAddressParams {
	params.EmergencyEnabled = &EmergencyEnabled
	return params
}
func (params *CreateAddressParams) SetFriendlyName(FriendlyName string) *CreateAddressParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateAddressParams) SetIsoCountry(IsoCountry string) *CreateAddressParams {
	params.IsoCountry = &IsoCountry
	return params
}
func (params *CreateAddressParams) SetPostalCode(PostalCode string) *CreateAddressParams {
	params.PostalCode = &PostalCode
	return params
}
func (params *CreateAddressParams) SetRegion(Region string) *CreateAddressParams {
	params.Region = &Region
	return params
}
func (params *CreateAddressParams) SetStreet(Street string) *CreateAddressParams {
	params.Street = &Street
	return params
}

func (c *ApiService) CreateAddress(params *CreateAddressParams) (*ApiV2010Address, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Addresses.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	if params != nil && params.AutoCorrectAddress != nil {
		data.Set("AutoCorrectAddress", fmt.Sprint(*params.AutoCorrectAddress))
	}
	if params != nil && params.City != nil {
		data.Set("City", *params.City)
	}
	if params != nil && params.CustomerName != nil {
		data.Set("CustomerName", *params.CustomerName)
	}
	if params != nil && params.EmergencyEnabled != nil {
		data.Set("EmergencyEnabled", fmt.Sprint(*params.EmergencyEnabled))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.IsoCountry != nil {
		data.Set("IsoCountry", *params.IsoCountry)
	}
	if params != nil && params.PostalCode != nil {
		data.Set("PostalCode", *params.PostalCode)
	}
	if params != nil && params.Region != nil {
		data.Set("Region", *params.Region)
	}
	if params != nil && params.Street != nil {
		data.Set("Street", *params.Street)
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010Address{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteAddress'
type DeleteAddressParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteAddressParams) SetPathAccountSid(PathAccountSid string) *DeleteAddressParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

func (c *ApiService) DeleteAddress(Sid string, params *DeleteAddressParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/Addresses/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
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

// Optional parameters for the method 'FetchAddress'
type FetchAddressParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchAddressParams) SetPathAccountSid(PathAccountSid string) *FetchAddressParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

func (c *ApiService) FetchAddress(Sid string, params *FetchAddressParams) (*ApiV2010Address, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Addresses/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010Address{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListAddress'
type ListAddressParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The `customer_name` of the Address resources to read.
	CustomerName *string `json:"CustomerName,omitempty"`
	// The string that identifies the Address resources to read.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The ISO country code of the Address resources to read.
	IsoCountry *string `json:"IsoCountry,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListAddressParams) SetPathAccountSid(PathAccountSid string) *ListAddressParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListAddressParams) SetCustomerName(CustomerName string) *ListAddressParams {
	params.CustomerName = &CustomerName
	return params
}
func (params *ListAddressParams) SetFriendlyName(FriendlyName string) *ListAddressParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListAddressParams) SetIsoCountry(IsoCountry string) *ListAddressParams {
	params.IsoCountry = &IsoCountry
	return params
}
func (params *ListAddressParams) SetPageSize(PageSize int) *ListAddressParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListAddressParams) SetLimit(Limit int) *ListAddressParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Address records from the API. Request is executed immediately.
func (c *ApiService) PageAddress(params *ListAddressParams, pageToken string, pageNumber string) (*ListAddressResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Addresses.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	if params != nil && params.CustomerName != nil {
		data.Set("CustomerName", *params.CustomerName)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.IsoCountry != nil {
		data.Set("IsoCountry", *params.IsoCountry)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}
	headers := make(map[string]interface{})

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageToken != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAddressResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Address records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListAddress(params *ListAddressParams) ([]ApiV2010Address, error) {
	if params == nil {
		params = &ListAddressParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageAddress(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010Address

	for response != nil {
		records = append(records, response.Addresses...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListAddressResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListAddressResponse)
	}

	return records, err
}

// Streams Address records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamAddress(params *ListAddressParams) (chan ApiV2010Address, error) {
	if params == nil {
		params = &ListAddressParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageAddress(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010Address, 1)

	go func() {
		for response != nil {
			for item := range response.Addresses {
				channel <- response.Addresses[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListAddressResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListAddressResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListAddressResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAddressResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateAddress'
type UpdateAddressParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to update.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// Whether we should automatically correct the address. Can be: `true` or `false` and the default is `true`. If empty or `true`, we will correct the address you provide if necessary. If `false`, we won't alter the address you provide.
	AutoCorrectAddress *bool `json:"AutoCorrectAddress,omitempty"`
	// The city of the address.
	City *string `json:"City,omitempty"`
	// The name to associate with the address.
	CustomerName *string `json:"CustomerName,omitempty"`
	// Whether to enable emergency calling on the address. Can be: `true` or `false`.
	EmergencyEnabled *bool `json:"EmergencyEnabled,omitempty"`
	// A descriptive string that you create to describe the address. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The postal code of the address.
	PostalCode *string `json:"PostalCode,omitempty"`
	// The state or region of the address.
	Region *string `json:"Region,omitempty"`
	// The number and street address of the address.
	Street *string `json:"Street,omitempty"`
}

func (params *UpdateAddressParams) SetPathAccountSid(PathAccountSid string) *UpdateAddressParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateAddressParams) SetAutoCorrectAddress(AutoCorrectAddress bool) *UpdateAddressParams {
	params.AutoCorrectAddress = &AutoCorrectAddress
	return params
}
func (params *UpdateAddressParams) SetCity(City string) *UpdateAddressParams {
	params.City = &City
	return params
}
func (params *UpdateAddressParams) SetCustomerName(CustomerName string) *UpdateAddressParams {
	params.CustomerName = &CustomerName
	return params
}
func (params *UpdateAddressParams) SetEmergencyEnabled(EmergencyEnabled bool) *UpdateAddressParams {
	params.EmergencyEnabled = &EmergencyEnabled
	return params
}
func (params *UpdateAddressParams) SetFriendlyName(FriendlyName string) *UpdateAddressParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateAddressParams) SetPostalCode(PostalCode string) *UpdateAddressParams {
	params.PostalCode = &PostalCode
	return params
}
func (params *UpdateAddressParams) SetRegion(Region string) *UpdateAddressParams {
	params.Region = &Region
	return params
}
func (params *UpdateAddressParams) SetStreet(Street string) *UpdateAddressParams {
	params.Street = &Street
	return params
}

func (c *ApiService) UpdateAddress(Sid string, params *UpdateAddressParams) (*ApiV2010Address, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Addresses/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	if params != nil && params.AutoCorrectAddress != nil {
		data.Set("AutoCorrectAddress", fmt.Sprint(*params.AutoCorrectAddress))
	}
	if params != nil && params.City != nil {
		data.Set("City", *params.City)
	}
	if params != nil && params.CustomerName != nil {
		data.Set("CustomerName", *params.CustomerName)
	}
	if params != nil && params.EmergencyEnabled != nil {
		data.Set("EmergencyEnabled", fmt.Sprint(*params.EmergencyEnabled))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.PostalCode != nil {
		data.Set("PostalCode", *params.PostalCode)
	}
	if params != nil && params.Region != nil {
		data.Set("Region", *params.Region)
	}
	if params != nil && params.Street != nil {
		data.Set("Street", *params.Street)
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010Address{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
