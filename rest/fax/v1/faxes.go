/*
 * Twilio - Fax
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/twilio/twilio-go/client"
)

// Delete a specific fax and its associated media.
func (c *ApiService) DeleteFax(Sid string) error {
	path := "/v1/Faxes/{Sid}"
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

// Fetch a specific fax.
func (c *ApiService) FetchFax(Sid string) (*FaxV1Fax, error) {
	path := "/v1/Faxes/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FaxV1Fax{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListFax'
type ListFaxParams struct {
	// Retrieve only those faxes sent from this phone number, specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.
	From *string `json:"From,omitempty"`
	// Retrieve only those faxes sent to this phone number, specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.
	To *string `json:"To,omitempty"`
	// Retrieve only those faxes with a `date_created` that is before or equal to this value, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreatedOnOrBefore *time.Time `json:"DateCreatedOnOrBefore,omitempty"`
	// Retrieve only those faxes with a `date_created` that is later than this value, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreatedAfter *time.Time `json:"DateCreatedAfter,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListFaxParams) SetFrom(From string) *ListFaxParams {
	params.From = &From
	return params
}
func (params *ListFaxParams) SetTo(To string) *ListFaxParams {
	params.To = &To
	return params
}
func (params *ListFaxParams) SetDateCreatedOnOrBefore(DateCreatedOnOrBefore time.Time) *ListFaxParams {
	params.DateCreatedOnOrBefore = &DateCreatedOnOrBefore
	return params
}
func (params *ListFaxParams) SetDateCreatedAfter(DateCreatedAfter time.Time) *ListFaxParams {
	params.DateCreatedAfter = &DateCreatedAfter
	return params
}
func (params *ListFaxParams) SetPageSize(PageSize int) *ListFaxParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListFaxParams) SetLimit(Limit int) *ListFaxParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Fax records from the API. Request is executed immediately.
func (c *ApiService) PageFax(params *ListFaxParams, pageToken, pageNumber string) (*ListFaxResponse, error) {
	path := "/v1/Faxes"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.To != nil {
		data.Set("To", *params.To)
	}
	if params != nil && params.DateCreatedOnOrBefore != nil {
		data.Set("DateCreatedOnOrBefore", fmt.Sprint((*params.DateCreatedOnOrBefore).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedAfter != nil {
		data.Set("DateCreatedAfter", fmt.Sprint((*params.DateCreatedAfter).Format(time.RFC3339)))
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

	ps := &ListFaxResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Fax records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListFax(params *ListFaxParams) ([]FaxV1Fax, error) {
	response, errors := c.StreamFax(params)

	records := make([]FaxV1Fax, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Fax records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamFax(params *ListFaxParams) (chan FaxV1Fax, chan error) {
	if params == nil {
		params = &ListFaxParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan FaxV1Fax, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageFax(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamFax(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamFax(response *ListFaxResponse, params *ListFaxParams, recordChannel chan FaxV1Fax, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Faxes
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListFaxResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListFaxResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListFaxResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFaxResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
