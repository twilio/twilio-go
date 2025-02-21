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
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateInsightsQuestionnairesCategory'
type CreateInsightsQuestionnairesCategoryParams struct {
	// The Authorization HTTP request header
	Authorization *string `json:"Authorization,omitempty"`
	// The name of this category.
	Name *string `json:"Name,omitempty"`
}

func (params *CreateInsightsQuestionnairesCategoryParams) SetAuthorization(Authorization string) *CreateInsightsQuestionnairesCategoryParams {
	params.Authorization = &Authorization
	return params
}
func (params *CreateInsightsQuestionnairesCategoryParams) SetName(Name string) *CreateInsightsQuestionnairesCategoryParams {
	params.Name = &Name
	return params
}

// To create a category for Questions
func (c *ApiService) CreateInsightsQuestionnairesCategory(params *CreateInsightsQuestionnairesCategoryParams) (*FlexV1InsightsQuestionnairesCategory, error) {
	path := "/v1/Insights/QualityManagement/Categories"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Name != nil {
		data.Set("Name", *params.Name)
	}

	if params != nil && params.Authorization != nil {
		headers["Authorization"] = *params.Authorization
	}
	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1InsightsQuestionnairesCategory{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteInsightsQuestionnairesCategory'
type DeleteInsightsQuestionnairesCategoryParams struct {
	// The Authorization HTTP request header
	Authorization *string `json:"Authorization,omitempty"`
}

func (params *DeleteInsightsQuestionnairesCategoryParams) SetAuthorization(Authorization string) *DeleteInsightsQuestionnairesCategoryParams {
	params.Authorization = &Authorization
	return params
}

func (c *ApiService) DeleteInsightsQuestionnairesCategory(CategorySid string, params *DeleteInsightsQuestionnairesCategoryParams) error {
	path := "/v1/Insights/QualityManagement/Categories/{CategorySid}"
	path = strings.Replace(path, "{"+"CategorySid"+"}", CategorySid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Authorization != nil {
		headers["Authorization"] = *params.Authorization
	}
	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Optional parameters for the method 'ListInsightsQuestionnairesCategory'
type ListInsightsQuestionnairesCategoryParams struct {
	// The Authorization HTTP request header
	Authorization *string `json:"Authorization,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int64 `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int64 `json:"limit,omitempty"`
}

func (params *ListInsightsQuestionnairesCategoryParams) SetAuthorization(Authorization string) *ListInsightsQuestionnairesCategoryParams {
	params.Authorization = &Authorization
	return params
}
func (params *ListInsightsQuestionnairesCategoryParams) SetPageSize(PageSize int64) *ListInsightsQuestionnairesCategoryParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListInsightsQuestionnairesCategoryParams) SetLimit(Limit int64) *ListInsightsQuestionnairesCategoryParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of InsightsQuestionnairesCategory records from the API. Request is executed immediately.
func (c *ApiService) PageInsightsQuestionnairesCategory(params *ListInsightsQuestionnairesCategoryParams, pageToken, pageNumber string) (*ListInsightsQuestionnairesCategoryResponse, error) {
	path := "/v1/Insights/QualityManagement/Categories"

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

	ps := &ListInsightsQuestionnairesCategoryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists InsightsQuestionnairesCategory records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListInsightsQuestionnairesCategory(params *ListInsightsQuestionnairesCategoryParams) ([]FlexV1InsightsQuestionnairesCategory, error) {
	response, errors := c.StreamInsightsQuestionnairesCategory(params)

	records := make([]FlexV1InsightsQuestionnairesCategory, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams InsightsQuestionnairesCategory records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamInsightsQuestionnairesCategory(params *ListInsightsQuestionnairesCategoryParams) (chan FlexV1InsightsQuestionnairesCategory, chan error) {
	if params == nil {
		params = &ListInsightsQuestionnairesCategoryParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan FlexV1InsightsQuestionnairesCategory, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageInsightsQuestionnairesCategory(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamInsightsQuestionnairesCategory(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamInsightsQuestionnairesCategory(response *ListInsightsQuestionnairesCategoryResponse, params *ListInsightsQuestionnairesCategoryParams, recordChannel chan FlexV1InsightsQuestionnairesCategory, errorChannel chan error) {
	var curRecord int64 = 1

	for response != nil {
		responseRecords := response.Categories
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListInsightsQuestionnairesCategoryResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListInsightsQuestionnairesCategoryResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListInsightsQuestionnairesCategoryResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListInsightsQuestionnairesCategoryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateInsightsQuestionnairesCategory'
type UpdateInsightsQuestionnairesCategoryParams struct {
	// The Authorization HTTP request header
	Authorization *string `json:"Authorization,omitempty"`
	// The name of this category.
	Name *string `json:"Name,omitempty"`
}

func (params *UpdateInsightsQuestionnairesCategoryParams) SetAuthorization(Authorization string) *UpdateInsightsQuestionnairesCategoryParams {
	params.Authorization = &Authorization
	return params
}
func (params *UpdateInsightsQuestionnairesCategoryParams) SetName(Name string) *UpdateInsightsQuestionnairesCategoryParams {
	params.Name = &Name
	return params
}

// To update the category for Questions
func (c *ApiService) UpdateInsightsQuestionnairesCategory(CategorySid string, params *UpdateInsightsQuestionnairesCategoryParams) (*FlexV1InsightsQuestionnairesCategory, error) {
	path := "/v1/Insights/QualityManagement/Categories/{CategorySid}"
	path = strings.Replace(path, "{"+"CategorySid"+"}", CategorySid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Name != nil {
		data.Set("Name", *params.Name)
	}

	if params != nil && params.Authorization != nil {
		headers["Authorization"] = *params.Authorization
	}
	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1InsightsQuestionnairesCategory{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
