/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Flex
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

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'ListInsightsConversations'
type ListInsightsConversationsParams struct {
	// The Authorization HTTP request header
	Authorization *string `json:"Authorization,omitempty"`
	// Unique Id of the segment for which conversation details needs to be fetched
	SegmentId *string `json:"SegmentId,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int64 `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int64 `json:"limit,omitempty"`
}

func (params *ListInsightsConversationsParams) SetAuthorization(Authorization string) *ListInsightsConversationsParams {
	params.Authorization = &Authorization
	return params
}
func (params *ListInsightsConversationsParams) SetSegmentId(SegmentId string) *ListInsightsConversationsParams {
	params.SegmentId = &SegmentId
	return params
}
func (params *ListInsightsConversationsParams) SetPageSize(PageSize int64) *ListInsightsConversationsParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListInsightsConversationsParams) SetLimit(Limit int64) *ListInsightsConversationsParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of InsightsConversations records from the API. Request is executed immediately.
func (c *ApiService) PageInsightsConversations(params *ListInsightsConversationsParams, pageToken, pageNumber string) (*ListInsightsConversationsResponse, error) {
	path := "/v1/Insights/Conversations"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.SegmentId != nil {
		data.Set("SegmentId", *params.SegmentId)
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

	ps := &ListInsightsConversationsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists InsightsConversations records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListInsightsConversations(params *ListInsightsConversationsParams) ([]FlexV1InsightsConversations, error) {
	response, errors := c.StreamInsightsConversations(params)

	records := make([]FlexV1InsightsConversations, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams InsightsConversations records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamInsightsConversations(params *ListInsightsConversationsParams) (chan FlexV1InsightsConversations, chan error) {
	if params == nil {
		params = &ListInsightsConversationsParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan FlexV1InsightsConversations, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageInsightsConversations(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamInsightsConversations(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamInsightsConversations(response *ListInsightsConversationsResponse, params *ListInsightsConversationsParams, recordChannel chan FlexV1InsightsConversations, errorChannel chan error) {
	var curRecord int64 = 1

	for response != nil {
		responseRecords := response.Conversations
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListInsightsConversationsResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListInsightsConversationsResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListInsightsConversationsResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListInsightsConversationsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
