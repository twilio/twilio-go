/*
 * Twilio - Bulkexports
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	twilio "github.com/twilio/twilio-go/client"
)

type DefaultApiService struct {
	baseURL string
	client  *twilio.Client
}

func NewDefaultApiService(client *twilio.Client) *DefaultApiService {
	return &DefaultApiService{
		client:  client,
		baseURL: "https://bulkexports.twilio.com",
	}
}

// CreateExportCustomJobParams Optional parameters for the method 'CreateExportCustomJob'
type CreateExportCustomJobParams struct {
	Email         *string `json:"Email,omitempty"`
	EndDay        *string `json:"EndDay,omitempty"`
	FriendlyName  *string `json:"FriendlyName,omitempty"`
	StartDay      *string `json:"StartDay,omitempty"`
	WebhookMethod *string `json:"WebhookMethod,omitempty"`
	WebhookUrl    *string `json:"WebhookUrl,omitempty"`
}

/*
* CreateExportCustomJob Method for CreateExportCustomJob
* @param ResourceType The type of communication – Messages or Calls, Conferences, and Participants
* @param optional nil or *CreateExportCustomJobParams - Optional Parameters:
* @param "Email" (string) - The optional email to send the completion notification to. You can set both webhook, and email, or one or the other. If you set neither, the job will run but you will have to query to determine your job's status.
* @param "EndDay" (string) - The end day for the custom export specified as a string in the format of yyyy-mm-dd. End day is inclusive and must be 2 days earlier than the current UTC day.
* @param "FriendlyName" (string) - The friendly name specified when creating the job
* @param "StartDay" (string) - The start day for the custom export specified as a string in the format of yyyy-mm-dd
* @param "WebhookMethod" (string) - This is the method used to call the webhook on completion of the job. If this is supplied, `WebhookUrl` must also be supplied.
* @param "WebhookUrl" (string) - The optional webhook url called on completion of the job. If this is supplied, `WebhookMethod` must also be supplied. If you set neither webhook nor email, you will have to check your job's status manually.
* @return BulkexportsV1ExportExportCustomJob
 */
func (c *DefaultApiService) CreateExportCustomJob(ResourceType string, params *CreateExportCustomJobParams) (*BulkexportsV1ExportExportCustomJob, error) {
	path := "/v1/Exports/{ResourceType}/Jobs"
	path = strings.Replace(path, "{"+"ResourceType"+"}", ResourceType, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.Email != nil {
		data.Set("Email", *params.Email)
	}
	if params != nil && params.EndDay != nil {
		data.Set("EndDay", *params.EndDay)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.StartDay != nil {
		data.Set("StartDay", *params.StartDay)
	}
	if params != nil && params.WebhookMethod != nil {
		data.Set("WebhookMethod", *params.WebhookMethod)
	}
	if params != nil && params.WebhookUrl != nil {
		data.Set("WebhookUrl", *params.WebhookUrl)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &BulkexportsV1ExportExportCustomJob{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
* DeleteJob Method for DeleteJob
* @param JobSid The unique string that that we created to identify the Bulk Export job
 */
func (c *DefaultApiService) DeleteJob(JobSid string) error {
	path := "/v1/Exports/Jobs/{JobSid}"
	path = strings.Replace(path, "{"+"JobSid"+"}", JobSid, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

/*
* FetchDay Method for FetchDay
* Fetch a specific Day.
* @param ResourceType The type of communication – Messages, Calls, Conferences, and Participants
* @param Day The ISO 8601 format date of the resources in the file, for a UTC day
 */
func (c *DefaultApiService) FetchDay(ResourceType string, Day string) error {
	path := "/v1/Exports/{ResourceType}/Days/{Day}"
	path = strings.Replace(path, "{"+"ResourceType"+"}", ResourceType, -1)
	path = strings.Replace(path, "{"+"Day"+"}", Day, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

/*
* FetchExport Method for FetchExport
* Fetch a specific Export.
* @param ResourceType The type of communication – Messages, Calls, Conferences, and Participants
* @return BulkexportsV1Export
 */
func (c *DefaultApiService) FetchExport(ResourceType string) (*BulkexportsV1Export, error) {
	path := "/v1/Exports/{ResourceType}"
	path = strings.Replace(path, "{"+"ResourceType"+"}", ResourceType, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &BulkexportsV1Export{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
* FetchExportConfiguration Method for FetchExportConfiguration
* Fetch a specific Export Configuration.
* @param ResourceType The type of communication – Messages, Calls, Conferences, and Participants
* @return BulkexportsV1ExportConfiguration
 */
func (c *DefaultApiService) FetchExportConfiguration(ResourceType string) (*BulkexportsV1ExportConfiguration, error) {
	path := "/v1/Exports/{ResourceType}/Configuration"
	path = strings.Replace(path, "{"+"ResourceType"+"}", ResourceType, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &BulkexportsV1ExportConfiguration{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
* FetchJob Method for FetchJob
* @param JobSid The unique string that that we created to identify the Bulk Export job
* @return BulkexportsV1ExportJob
 */
func (c *DefaultApiService) FetchJob(JobSid string) (*BulkexportsV1ExportJob, error) {
	path := "/v1/Exports/Jobs/{JobSid}"
	path = strings.Replace(path, "{"+"JobSid"+"}", JobSid, -1)

	data := url.Values{}
	headers := 0

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &BulkexportsV1ExportJob{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListDayParams Optional parameters for the method 'ListDay'
type ListDayParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListDay Method for ListDay
* Retrieve a list of all Days for a resource.
* @param ResourceType The type of communication – Messages, Calls, Conferences, and Participants
* @param optional nil or *ListDayParams - Optional Parameters:
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListDayResponse
 */
func (c *DefaultApiService) ListDay(ResourceType string, params *ListDayParams) (*ListDayResponse, error) {
	path := "/v1/Exports/{ResourceType}/Days"
	path = strings.Replace(path, "{"+"ResourceType"+"}", ResourceType, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListDayResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListExportCustomJobParams Optional parameters for the method 'ListExportCustomJob'
type ListExportCustomJobParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListExportCustomJob Method for ListExportCustomJob
* @param ResourceType The type of communication – Messages, Calls, Conferences, and Participants
* @param optional nil or *ListExportCustomJobParams - Optional Parameters:
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListExportCustomJobResponse
 */
func (c *DefaultApiService) ListExportCustomJob(ResourceType string, params *ListExportCustomJobParams) (*ListExportCustomJobResponse, error) {
	path := "/v1/Exports/{ResourceType}/Jobs"
	path = strings.Replace(path, "{"+"ResourceType"+"}", ResourceType, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListExportCustomJobResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// UpdateExportConfigurationParams Optional parameters for the method 'UpdateExportConfiguration'
type UpdateExportConfigurationParams struct {
	Enabled       *bool   `json:"Enabled,omitempty"`
	WebhookMethod *string `json:"WebhookMethod,omitempty"`
	WebhookUrl    *string `json:"WebhookUrl,omitempty"`
}

/*
* UpdateExportConfiguration Method for UpdateExportConfiguration
* Update a specific Export Configuration.
* @param ResourceType The type of communication – Messages, Calls, Conferences, and Participants
* @param optional nil or *UpdateExportConfigurationParams - Optional Parameters:
* @param "Enabled" (bool) - If true, Twilio will automatically generate every day's file when the day is over.
* @param "WebhookMethod" (string) - Sets whether Twilio should call a webhook URL when the automatic generation is complete, using GET or POST. The actual destination is set in the webhook_url
* @param "WebhookUrl" (string) - Stores the URL destination for the method specified in webhook_method.
* @return BulkexportsV1ExportConfiguration
 */
func (c *DefaultApiService) UpdateExportConfiguration(ResourceType string, params *UpdateExportConfigurationParams) (*BulkexportsV1ExportConfiguration, error) {
	path := "/v1/Exports/{ResourceType}/Configuration"
	path = strings.Replace(path, "{"+"ResourceType"+"}", ResourceType, -1)

	data := url.Values{}
	headers := 0

	if params != nil && params.Enabled != nil {
		data.Set("Enabled", fmt.Sprint(*params.Enabled))
	}
	if params != nil && params.WebhookMethod != nil {
		data.Set("WebhookMethod", *params.WebhookMethod)
	}
	if params != nil && params.WebhookUrl != nil {
		data.Set("WebhookUrl", *params.WebhookUrl)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &BulkexportsV1ExportConfiguration{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
