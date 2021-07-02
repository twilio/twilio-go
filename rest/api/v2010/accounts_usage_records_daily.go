/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.18.0
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

// Optional parameters for the method 'ListUsageRecordDaily'
type ListUsageRecordDailyParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved.
	Category *string `json:"Category,omitempty"`
	// Only include usage that has occurred on or after this date. Specify the date in GMT and format as `YYYY-MM-DD`. You can also specify offsets from the current date, such as: `-30days`, which will set the start date to be 30 days before the current date.
	StartDate *string `json:"StartDate,omitempty"`
	// Only include usage that occurred on or before this date. Specify the date in GMT and format as `YYYY-MM-DD`.  You can also specify offsets from the current date, such as: `+30days`, which will set the end date to 30 days from the current date.
	EndDate *string `json:"EndDate,omitempty"`
	// Whether to include usage from the master account and all its subaccounts. Can be: `true` (the default) to include usage from the master account and all subaccounts or `false` to retrieve usage from only the specified account.
	IncludeSubaccounts *bool `json:"IncludeSubaccounts,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListUsageRecordDailyParams) SetPathAccountSid(PathAccountSid string) *ListUsageRecordDailyParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListUsageRecordDailyParams) SetCategory(Category string) *ListUsageRecordDailyParams {
	params.Category = &Category
	return params
}
func (params *ListUsageRecordDailyParams) SetStartDate(StartDate string) *ListUsageRecordDailyParams {
	params.StartDate = &StartDate
	return params
}
func (params *ListUsageRecordDailyParams) SetEndDate(EndDate string) *ListUsageRecordDailyParams {
	params.EndDate = &EndDate
	return params
}
func (params *ListUsageRecordDailyParams) SetIncludeSubaccounts(IncludeSubaccounts bool) *ListUsageRecordDailyParams {
	params.IncludeSubaccounts = &IncludeSubaccounts
	return params
}
func (params *ListUsageRecordDailyParams) SetPageSize(PageSize int) *ListUsageRecordDailyParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListUsageRecordDaily(params *ListUsageRecordDailyParams) (*ListUsageRecordDailyResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Usage/Records/Daily.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Category != nil {
		data.Set("Category", *params.Category)
	}
	if params != nil && params.StartDate != nil {
		data.Set("StartDate", fmt.Sprint(*params.StartDate))
	}
	if params != nil && params.EndDate != nil {
		data.Set("EndDate", fmt.Sprint(*params.EndDate))
	}
	if params != nil && params.IncludeSubaccounts != nil {
		data.Set("IncludeSubaccounts", fmt.Sprint(*params.IncludeSubaccounts))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListUsageRecordDailyResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of UsageRecordDaily records from the API. Request is executed immediately.
func (c *ApiService) UsageRecordDailyPage(params *ListUsageRecordDailyParams, pageToken string, pageNumber string) (*client.Page, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Usage/Records/Daily.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Category != nil {
		data.Set("Category", *params.Category)
	}
	if params != nil && params.StartDate != nil {
		data.Set("StartDate", fmt.Sprint(*params.StartDate))
	}
	if params != nil && params.EndDate != nil {
		data.Set("EndDate", fmt.Sprint(*params.EndDate))
	}
	if params != nil && params.IncludeSubaccounts != nil {
		data.Set("IncludeSubaccounts", fmt.Sprint(*params.IncludeSubaccounts))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	data.Set("PageToken", pageToken)
	data.Set("PageNumber", pageNumber)

	response, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	return client.NewPage(c.baseURL, response), nil
}

//Streams UsageRecordDaily records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) UsageRecordDailyStream(params *ListUsageRecordDailyParams, limit int) (chan map[string]interface{}, error) {
	if params.PageSize == nil {
		params.SetPageSize(0)
	}
	params.SetPageSize(c.requestHandler.ReadLimits(*params.PageSize, limit))
	page, err := c.UsageRecordDailyPage(params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.Stream(page, limit, 0), nil
}

//Lists UsageRecordDaily records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) UsageRecordDailyList(params *ListUsageRecordDailyParams, limit int) ([]interface{}, error) {
	if params.PageSize == nil {
		params.SetPageSize(0)
	}
	params.SetPageSize(c.requestHandler.ReadLimits(*params.PageSize, limit))
	page, err := c.UsageRecordDailyPage(params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.List(page, limit, 0), nil
}
