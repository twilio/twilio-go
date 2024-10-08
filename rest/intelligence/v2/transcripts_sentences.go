/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Intelligence
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

// Optional parameters for the method 'ListSentence'
type ListSentenceParams struct {
	// Grant access to PII Redacted/Unredacted Sentences. If redaction is enabled, the default is `true` to access redacted sentences.
	Redacted *bool `json:"Redacted,omitempty"`
	// Returns word level timestamps information, if word_timestamps is enabled. The default is `false`.
	WordTimestamps *bool `json:"WordTimestamps,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSentenceParams) SetRedacted(Redacted bool) *ListSentenceParams {
	params.Redacted = &Redacted
	return params
}
func (params *ListSentenceParams) SetWordTimestamps(WordTimestamps bool) *ListSentenceParams {
	params.WordTimestamps = &WordTimestamps
	return params
}
func (params *ListSentenceParams) SetPageSize(PageSize int) *ListSentenceParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSentenceParams) SetLimit(Limit int) *ListSentenceParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Sentence records from the API. Request is executed immediately.
func (c *ApiService) PageSentence(TranscriptSid string, params *ListSentenceParams, pageToken, pageNumber string) (*ListSentenceResponse, error) {
	path := "/v2/Transcripts/{TranscriptSid}/Sentences"

	path = strings.Replace(path, "{"+"TranscriptSid"+"}", TranscriptSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Redacted != nil {
		data.Set("Redacted", fmt.Sprint(*params.Redacted))
	}
	if params != nil && params.WordTimestamps != nil {
		data.Set("WordTimestamps", fmt.Sprint(*params.WordTimestamps))
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

	ps := &ListSentenceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Sentence records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSentence(TranscriptSid string, params *ListSentenceParams) ([]IntelligenceV2Sentence, error) {
	response, errors := c.StreamSentence(TranscriptSid, params)

	records := make([]IntelligenceV2Sentence, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Sentence records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSentence(TranscriptSid string, params *ListSentenceParams) (chan IntelligenceV2Sentence, chan error) {
	if params == nil {
		params = &ListSentenceParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan IntelligenceV2Sentence, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageSentence(TranscriptSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamSentence(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamSentence(response *ListSentenceResponse, params *ListSentenceParams, recordChannel chan IntelligenceV2Sentence, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Sentences
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListSentenceResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListSentenceResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListSentenceResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSentenceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
