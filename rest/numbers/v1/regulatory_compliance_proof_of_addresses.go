/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Numbers
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

// Optional parameters for the method 'CreateProofOfAddress'
type CreateProofOfAddressParams struct {
	// The SID of the [Address](https://www.twilio.com/docs/usage/api/address) resource that is associated with the ProofOfAddress resource to update.
	AddressSid *string `json:"AddressSid,omitempty"`
	// The SID of the Identity resource that is associated with the ProofOfAddress resource.
	IdentitySid *string `json:"IdentitySid,omitempty"`
}

func (params *CreateProofOfAddressParams) SetAddressSid(AddressSid string) *CreateProofOfAddressParams {
	params.AddressSid = &AddressSid
	return params
}
func (params *CreateProofOfAddressParams) SetIdentitySid(IdentitySid string) *CreateProofOfAddressParams {
	params.IdentitySid = &IdentitySid
	return params
}

// Create an ProofOfAddress for authorizing the hosting of phone number capabilities on our platform.
func (c *ApiService) CreateProofOfAddress(params *CreateProofOfAddressParams) (*NumbersV1ProofOfAddress, error) {
	path := "/v1/RegulatoryCompliance/ProofOfAddresses"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AddressSid != nil {
		data.Set("AddressSid", *params.AddressSid)
	}
	if params != nil && params.IdentitySid != nil {
		data.Set("IdentitySid", *params.IdentitySid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV1ProofOfAddress{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Fetch a specific ProofOfAddress.
func (c *ApiService) FetchProofOfAddress(Sid string) (*NumbersV1ProofOfAddress, error) {
	path := "/v1/RegulatoryCompliance/ProofOfAddresses/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV1ProofOfAddress{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListProofOfAddress'
type ListProofOfAddressParams struct {
	// The SID of the [Address](https://www.twilio.com/docs/usage/api/address) resource that is associated with the ProofOfAddress resources to read.
	AddressSid *string `json:"AddressSid,omitempty"`
	// The SID of the Identity resource that is associated with the ProofOfAddress resources to read.
	IdentitySid *string `json:"IdentitySid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListProofOfAddressParams) SetAddressSid(AddressSid string) *ListProofOfAddressParams {
	params.AddressSid = &AddressSid
	return params
}
func (params *ListProofOfAddressParams) SetIdentitySid(IdentitySid string) *ListProofOfAddressParams {
	params.IdentitySid = &IdentitySid
	return params
}
func (params *ListProofOfAddressParams) SetPageSize(PageSize int) *ListProofOfAddressParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListProofOfAddressParams) SetLimit(Limit int) *ListProofOfAddressParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ProofOfAddress records from the API. Request is executed immediately.
func (c *ApiService) PageProofOfAddress(params *ListProofOfAddressParams, pageToken, pageNumber string) (*ListProofOfAddressResponse, error) {
	path := "/v1/RegulatoryCompliance/ProofOfAddresses"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AddressSid != nil {
		data.Set("AddressSid", *params.AddressSid)
	}
	if params != nil && params.IdentitySid != nil {
		data.Set("IdentitySid", *params.IdentitySid)
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

	ps := &ListProofOfAddressResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ProofOfAddress records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListProofOfAddress(params *ListProofOfAddressParams) ([]NumbersV1ProofOfAddress, error) {
	response, errors := c.StreamProofOfAddress(params)

	records := make([]NumbersV1ProofOfAddress, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams ProofOfAddress records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamProofOfAddress(params *ListProofOfAddressParams) (chan NumbersV1ProofOfAddress, chan error) {
	if params == nil {
		params = &ListProofOfAddressParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan NumbersV1ProofOfAddress, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageProofOfAddress(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamProofOfAddress(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamProofOfAddress(response *ListProofOfAddressResponse, params *ListProofOfAddressParams, recordChannel chan NumbersV1ProofOfAddress, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Items
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListProofOfAddressResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListProofOfAddressResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListProofOfAddressResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListProofOfAddressResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
