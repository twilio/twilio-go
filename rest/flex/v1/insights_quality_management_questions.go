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

// Optional parameters for the method 'CreateInsightsQuestionnairesQuestion'
type CreateInsightsQuestionnairesQuestionParams struct {
	// The Authorization HTTP request header
	Authorization *string `json:"Authorization,omitempty"`
	// The SID of the category
	CategorySid *string `json:"CategorySid,omitempty"`
	// The question.
	Question *string `json:"Question,omitempty"`
	// The answer_set for the question.
	AnswerSetId *string `json:"AnswerSetId,omitempty"`
	// The flag to enable for disable NA for answer.
	AllowNa *bool `json:"AllowNa,omitempty"`
	// The description for the question.
	Description *string `json:"Description,omitempty"`
}

func (params *CreateInsightsQuestionnairesQuestionParams) SetAuthorization(Authorization string) *CreateInsightsQuestionnairesQuestionParams {
	params.Authorization = &Authorization
	return params
}
func (params *CreateInsightsQuestionnairesQuestionParams) SetCategorySid(CategorySid string) *CreateInsightsQuestionnairesQuestionParams {
	params.CategorySid = &CategorySid
	return params
}
func (params *CreateInsightsQuestionnairesQuestionParams) SetQuestion(Question string) *CreateInsightsQuestionnairesQuestionParams {
	params.Question = &Question
	return params
}
func (params *CreateInsightsQuestionnairesQuestionParams) SetAnswerSetId(AnswerSetId string) *CreateInsightsQuestionnairesQuestionParams {
	params.AnswerSetId = &AnswerSetId
	return params
}
func (params *CreateInsightsQuestionnairesQuestionParams) SetAllowNa(AllowNa bool) *CreateInsightsQuestionnairesQuestionParams {
	params.AllowNa = &AllowNa
	return params
}
func (params *CreateInsightsQuestionnairesQuestionParams) SetDescription(Description string) *CreateInsightsQuestionnairesQuestionParams {
	params.Description = &Description
	return params
}

// To create a question for a Category
func (c *ApiService) CreateInsightsQuestionnairesQuestion(params *CreateInsightsQuestionnairesQuestionParams) (*FlexV1InsightsQuestionnairesQuestion, error) {
	path := "/v1/Insights/QualityManagement/Questions"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.CategorySid != nil {
		data.Set("CategorySid", *params.CategorySid)
	}
	if params != nil && params.Question != nil {
		data.Set("Question", *params.Question)
	}
	if params != nil && params.AnswerSetId != nil {
		data.Set("AnswerSetId", *params.AnswerSetId)
	}
	if params != nil && params.AllowNa != nil {
		data.Set("AllowNa", fmt.Sprint(*params.AllowNa))
	}
	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}

	if params != nil && params.Authorization != nil {
		headers["Authorization"] = *params.Authorization
	}
	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1InsightsQuestionnairesQuestion{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteInsightsQuestionnairesQuestion'
type DeleteInsightsQuestionnairesQuestionParams struct {
	// The Authorization HTTP request header
	Authorization *string `json:"Authorization,omitempty"`
}

func (params *DeleteInsightsQuestionnairesQuestionParams) SetAuthorization(Authorization string) *DeleteInsightsQuestionnairesQuestionParams {
	params.Authorization = &Authorization
	return params
}

func (c *ApiService) DeleteInsightsQuestionnairesQuestion(QuestionSid string, params *DeleteInsightsQuestionnairesQuestionParams) error {
	path := "/v1/Insights/QualityManagement/Questions/{QuestionSid}"
	path = strings.Replace(path, "{"+"QuestionSid"+"}", QuestionSid, -1)

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

// Optional parameters for the method 'ListInsightsQuestionnairesQuestion'
type ListInsightsQuestionnairesQuestionParams struct {
	// The Authorization HTTP request header
	Authorization *string `json:"Authorization,omitempty"`
	// The list of category SIDs
	CategorySid *[]string `json:"CategorySid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int64 `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int64 `json:"limit,omitempty"`
}

func (params *ListInsightsQuestionnairesQuestionParams) SetAuthorization(Authorization string) *ListInsightsQuestionnairesQuestionParams {
	params.Authorization = &Authorization
	return params
}
func (params *ListInsightsQuestionnairesQuestionParams) SetCategorySid(CategorySid []string) *ListInsightsQuestionnairesQuestionParams {
	params.CategorySid = &CategorySid
	return params
}
func (params *ListInsightsQuestionnairesQuestionParams) SetPageSize(PageSize int64) *ListInsightsQuestionnairesQuestionParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListInsightsQuestionnairesQuestionParams) SetLimit(Limit int64) *ListInsightsQuestionnairesQuestionParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of InsightsQuestionnairesQuestion records from the API. Request is executed immediately.
func (c *ApiService) PageInsightsQuestionnairesQuestion(params *ListInsightsQuestionnairesQuestionParams, pageToken, pageNumber string) (*ListInsightsQuestionnairesQuestionResponse, error) {
	path := "/v1/Insights/QualityManagement/Questions"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.CategorySid != nil {
		for _, item := range *params.CategorySid {
			data.Add("CategorySid", item)
		}
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

	ps := &ListInsightsQuestionnairesQuestionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists InsightsQuestionnairesQuestion records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListInsightsQuestionnairesQuestion(params *ListInsightsQuestionnairesQuestionParams) ([]FlexV1InsightsQuestionnairesQuestion, error) {
	response, errors := c.StreamInsightsQuestionnairesQuestion(params)

	records := make([]FlexV1InsightsQuestionnairesQuestion, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams InsightsQuestionnairesQuestion records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamInsightsQuestionnairesQuestion(params *ListInsightsQuestionnairesQuestionParams) (chan FlexV1InsightsQuestionnairesQuestion, chan error) {
	if params == nil {
		params = &ListInsightsQuestionnairesQuestionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan FlexV1InsightsQuestionnairesQuestion, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageInsightsQuestionnairesQuestion(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamInsightsQuestionnairesQuestion(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamInsightsQuestionnairesQuestion(response *ListInsightsQuestionnairesQuestionResponse, params *ListInsightsQuestionnairesQuestionParams, recordChannel chan FlexV1InsightsQuestionnairesQuestion, errorChannel chan error) {
	var curRecord int64 = 1

	for response != nil {
		responseRecords := response.Questions
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListInsightsQuestionnairesQuestionResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListInsightsQuestionnairesQuestionResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListInsightsQuestionnairesQuestionResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListInsightsQuestionnairesQuestionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateInsightsQuestionnairesQuestion'
type UpdateInsightsQuestionnairesQuestionParams struct {
	// The Authorization HTTP request header
	Authorization *string `json:"Authorization,omitempty"`
	// The flag to enable for disable NA for answer.
	AllowNa *bool `json:"AllowNa,omitempty"`
	// The SID of the category
	CategorySid *string `json:"CategorySid,omitempty"`
	// The question.
	Question *string `json:"Question,omitempty"`
	// The description for the question.
	Description *string `json:"Description,omitempty"`
	// The answer_set for the question.
	AnswerSetId *string `json:"AnswerSetId,omitempty"`
}

func (params *UpdateInsightsQuestionnairesQuestionParams) SetAuthorization(Authorization string) *UpdateInsightsQuestionnairesQuestionParams {
	params.Authorization = &Authorization
	return params
}
func (params *UpdateInsightsQuestionnairesQuestionParams) SetAllowNa(AllowNa bool) *UpdateInsightsQuestionnairesQuestionParams {
	params.AllowNa = &AllowNa
	return params
}
func (params *UpdateInsightsQuestionnairesQuestionParams) SetCategorySid(CategorySid string) *UpdateInsightsQuestionnairesQuestionParams {
	params.CategorySid = &CategorySid
	return params
}
func (params *UpdateInsightsQuestionnairesQuestionParams) SetQuestion(Question string) *UpdateInsightsQuestionnairesQuestionParams {
	params.Question = &Question
	return params
}
func (params *UpdateInsightsQuestionnairesQuestionParams) SetDescription(Description string) *UpdateInsightsQuestionnairesQuestionParams {
	params.Description = &Description
	return params
}
func (params *UpdateInsightsQuestionnairesQuestionParams) SetAnswerSetId(AnswerSetId string) *UpdateInsightsQuestionnairesQuestionParams {
	params.AnswerSetId = &AnswerSetId
	return params
}

// To update the question
func (c *ApiService) UpdateInsightsQuestionnairesQuestion(QuestionSid string, params *UpdateInsightsQuestionnairesQuestionParams) (*FlexV1InsightsQuestionnairesQuestion, error) {
	path := "/v1/Insights/QualityManagement/Questions/{QuestionSid}"
	path = strings.Replace(path, "{"+"QuestionSid"+"}", QuestionSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.AllowNa != nil {
		data.Set("AllowNa", fmt.Sprint(*params.AllowNa))
	}
	if params != nil && params.CategorySid != nil {
		data.Set("CategorySid", *params.CategorySid)
	}
	if params != nil && params.Question != nil {
		data.Set("Question", *params.Question)
	}
	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}
	if params != nil && params.AnswerSetId != nil {
		data.Set("AnswerSetId", *params.AnswerSetId)
	}

	if params != nil && params.Authorization != nil {
		headers["Authorization"] = *params.Authorization
	}
	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1InsightsQuestionnairesQuestion{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
