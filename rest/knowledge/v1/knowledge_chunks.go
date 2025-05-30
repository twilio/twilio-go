/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Knowledge
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

// Optional parameters for the method 'ListKnowledgeChunks'
type ListKnowledgeChunksParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListKnowledgeChunksParams) SetPageSize(PageSize int) *ListKnowledgeChunksParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListKnowledgeChunksParams) SetLimit(Limit int) *ListKnowledgeChunksParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of KnowledgeChunks records from the API. Request is executed immediately.
func (c *ApiService) PageKnowledgeChunks(Id string, params *ListKnowledgeChunksParams, pageToken, pageNumber string) (*ListKnowledgeChunksResponse, error) {
	path := "/v1/Knowledge/{id}/Chunks"

	path = strings.Replace(path, "{"+"id"+"}", Id, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
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

	ps := &ListKnowledgeChunksResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists KnowledgeChunks records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListKnowledgeChunks(Id string, params *ListKnowledgeChunksParams) ([]KnowledgeV1KnowledgeChunk, error) {
	response, errors := c.StreamKnowledgeChunks(Id, params)

	records := make([]KnowledgeV1KnowledgeChunk, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams KnowledgeChunks records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamKnowledgeChunks(Id string, params *ListKnowledgeChunksParams) (chan KnowledgeV1KnowledgeChunk, chan error) {
	if params == nil {
		params = &ListKnowledgeChunksParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan KnowledgeV1KnowledgeChunk, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageKnowledgeChunks(Id, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamKnowledgeChunks(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamKnowledgeChunks(response *ListKnowledgeChunksResponse, params *ListKnowledgeChunksParams, recordChannel chan KnowledgeV1KnowledgeChunk, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Chunks
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListKnowledgeChunksResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListKnowledgeChunksResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListKnowledgeChunksResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListKnowledgeChunksResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
