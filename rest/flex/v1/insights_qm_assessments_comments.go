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


// Optional parameters for the method 'CreateInsightsAssessmentsComment'
type CreateInsightsAssessmentsCommentParams struct {
    // The Token HTTP request header
    Token *string `json:"Token,omitempty"`
    // The ID of the category
    CategoryId *string `json:"CategoryId,omitempty"`
    // The name of the category
    CategoryName *string `json:"CategoryName,omitempty"`
    // The Assessment comment.
    Comment *string `json:"Comment,omitempty"`
    // The id of the segment.
    SegmentId *string `json:"SegmentId,omitempty"`
    // The name of the user.
    UserName *string `json:"UserName,omitempty"`
    // The email id of the user.
    UserEmail *string `json:"UserEmail,omitempty"`
    // The id of the agent.
    AgentId *string `json:"AgentId,omitempty"`
    // The offset
    Offset *float32 `json:"Offset,omitempty"`
}

func (params *CreateInsightsAssessmentsCommentParams) SetToken(Token string) (*CreateInsightsAssessmentsCommentParams){
    params.Token = &Token
    return params
}
func (params *CreateInsightsAssessmentsCommentParams) SetCategoryId(CategoryId string) (*CreateInsightsAssessmentsCommentParams){
    params.CategoryId = &CategoryId
    return params
}
func (params *CreateInsightsAssessmentsCommentParams) SetCategoryName(CategoryName string) (*CreateInsightsAssessmentsCommentParams){
    params.CategoryName = &CategoryName
    return params
}
func (params *CreateInsightsAssessmentsCommentParams) SetComment(Comment string) (*CreateInsightsAssessmentsCommentParams){
    params.Comment = &Comment
    return params
}
func (params *CreateInsightsAssessmentsCommentParams) SetSegmentId(SegmentId string) (*CreateInsightsAssessmentsCommentParams){
    params.SegmentId = &SegmentId
    return params
}
func (params *CreateInsightsAssessmentsCommentParams) SetUserName(UserName string) (*CreateInsightsAssessmentsCommentParams){
    params.UserName = &UserName
    return params
}
func (params *CreateInsightsAssessmentsCommentParams) SetUserEmail(UserEmail string) (*CreateInsightsAssessmentsCommentParams){
    params.UserEmail = &UserEmail
    return params
}
func (params *CreateInsightsAssessmentsCommentParams) SetAgentId(AgentId string) (*CreateInsightsAssessmentsCommentParams){
    params.AgentId = &AgentId
    return params
}
func (params *CreateInsightsAssessmentsCommentParams) SetOffset(Offset float32) (*CreateInsightsAssessmentsCommentParams){
    params.Offset = &Offset
    return params
}

// To create a comment assessment for a conversation
func (c *ApiService) CreateInsightsAssessmentsComment(params *CreateInsightsAssessmentsCommentParams) (*FlexV1InsightsAssessmentsComment, error) {
    path := "/v1/Insights/QM/Assessments/Comments"
    
data := url.Values{}
headers := make(map[string]interface{})

if params != nil && params.CategoryId != nil {
    data.Set("CategoryId", *params.CategoryId)
}
if params != nil && params.CategoryName != nil {
    data.Set("CategoryName", *params.CategoryName)
}
if params != nil && params.Comment != nil {
    data.Set("Comment", *params.Comment)
}
if params != nil && params.SegmentId != nil {
    data.Set("SegmentId", *params.SegmentId)
}
if params != nil && params.UserName != nil {
    data.Set("UserName", *params.UserName)
}
if params != nil && params.UserEmail != nil {
    data.Set("UserEmail", *params.UserEmail)
}
if params != nil && params.AgentId != nil {
    data.Set("AgentId", *params.AgentId)
}
if params != nil && params.Offset != nil {
    data.Set("Offset", fmt.Sprint(*params.Offset))
}


	if params != nil && params.Token != nil {
		headers["Token"] = *params.Token
	}

    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &FlexV1InsightsAssessmentsComment{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'ListInsightsAssessmentsComment'
type ListInsightsAssessmentsCommentParams struct {
    // The Token HTTP request header
    Token *string `json:"Token,omitempty"`
    // The id of the segment.
    SegmentId *string `json:"SegmentId,omitempty"`
    // The id of the agent.
    AgentId *string `json:"AgentId,omitempty"`
    // How many resources to return in each list page. The default is 50, and the maximum is 1000.
    PageSize *int `json:"PageSize,omitempty"`
    // Max number of records to return.
    Limit *int `json:"limit,omitempty"`
}

func (params *ListInsightsAssessmentsCommentParams) SetToken(Token string) (*ListInsightsAssessmentsCommentParams){
    params.Token = &Token
    return params
}
func (params *ListInsightsAssessmentsCommentParams) SetSegmentId(SegmentId string) (*ListInsightsAssessmentsCommentParams){
    params.SegmentId = &SegmentId
    return params
}
func (params *ListInsightsAssessmentsCommentParams) SetAgentId(AgentId string) (*ListInsightsAssessmentsCommentParams){
    params.AgentId = &AgentId
    return params
}
func (params *ListInsightsAssessmentsCommentParams) SetPageSize(PageSize int) (*ListInsightsAssessmentsCommentParams){
    params.PageSize = &PageSize
    return params
}
func (params *ListInsightsAssessmentsCommentParams) SetLimit(Limit int) (*ListInsightsAssessmentsCommentParams){
    params.Limit = &Limit
    return params
}

// Retrieve a single page of InsightsAssessmentsComment records from the API. Request is executed immediately.
func (c *ApiService) PageInsightsAssessmentsComment(params *ListInsightsAssessmentsCommentParams, pageToken, pageNumber string) (*ListInsightsAssessmentsCommentResponse, error) {
    path := "/v1/Insights/QM/Assessments/Comments"

    
data := url.Values{}
headers := make(map[string]interface{})

if params != nil && params.SegmentId != nil {
    data.Set("SegmentId", *params.SegmentId)
}
if params != nil && params.AgentId != nil {
    data.Set("AgentId", *params.AgentId)
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

    ps := &ListInsightsAssessmentsCommentResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Lists InsightsAssessmentsComment records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListInsightsAssessmentsComment(params *ListInsightsAssessmentsCommentParams) ([]FlexV1InsightsAssessmentsComment, error) {
	response, errors := c.StreamInsightsAssessmentsComment(params)

	records := make([]FlexV1InsightsAssessmentsComment, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams InsightsAssessmentsComment records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamInsightsAssessmentsComment(params *ListInsightsAssessmentsCommentParams) (chan FlexV1InsightsAssessmentsComment, chan error) {
	if params == nil {
		params = &ListInsightsAssessmentsCommentParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan FlexV1InsightsAssessmentsComment, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageInsightsAssessmentsComment(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamInsightsAssessmentsComment(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}


func (c *ApiService) streamInsightsAssessmentsComment(response *ListInsightsAssessmentsCommentResponse, params *ListInsightsAssessmentsCommentParams, recordChannel chan FlexV1InsightsAssessmentsComment, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Comments
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListInsightsAssessmentsCommentResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListInsightsAssessmentsCommentResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListInsightsAssessmentsCommentResponse(nextPageUrl string) (interface{}, error) {
    if nextPageUrl == "" {
        return nil, nil
    }
    resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListInsightsAssessmentsCommentResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }
    return ps, nil
}

