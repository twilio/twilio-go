/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Conversations
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


// Optional parameters for the method 'CreateServiceConversationMessage'
type CreateServiceConversationMessageParams struct {
    // The X-Twilio-Webhook-Enabled HTTP request header
    XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
    // The channel specific identifier of the message's author. Defaults to `system`.
    Author *string `json:"Author,omitempty"`
    // The content of the message, can be up to 1,600 characters long.
    Body *string `json:"Body,omitempty"`
    // The date that this resource was created.
    DateCreated *time.Time `json:"DateCreated,omitempty"`
    // The date that this resource was last updated. `null` if the message has not been edited.
    DateUpdated *time.Time `json:"DateUpdated,omitempty"`
    // A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
    Attributes *string `json:"Attributes,omitempty"`
    // The Media SID to be attached to the new Message.
    MediaSid *string `json:"MediaSid,omitempty"`
    // The unique ID of the multi-channel [Rich Content](https://www.twilio.com/docs/content-api) template, required for template-generated messages.  **Note** that if this field is set, `Body` and `MediaSid` parameters are ignored.
    ContentSid *string `json:"ContentSid,omitempty"`
    // A structurally valid JSON string that contains values to resolve Rich Content template variables.
    ContentVariables *string `json:"ContentVariables,omitempty"`
}

func (params *CreateServiceConversationMessageParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) (*CreateServiceConversationMessageParams){
    params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
    return params
}
func (params *CreateServiceConversationMessageParams) SetAuthor(Author string) (*CreateServiceConversationMessageParams){
    params.Author = &Author
    return params
}
func (params *CreateServiceConversationMessageParams) SetBody(Body string) (*CreateServiceConversationMessageParams){
    params.Body = &Body
    return params
}
func (params *CreateServiceConversationMessageParams) SetDateCreated(DateCreated time.Time) (*CreateServiceConversationMessageParams){
    params.DateCreated = &DateCreated
    return params
}
func (params *CreateServiceConversationMessageParams) SetDateUpdated(DateUpdated time.Time) (*CreateServiceConversationMessageParams){
    params.DateUpdated = &DateUpdated
    return params
}
func (params *CreateServiceConversationMessageParams) SetAttributes(Attributes string) (*CreateServiceConversationMessageParams){
    params.Attributes = &Attributes
    return params
}
func (params *CreateServiceConversationMessageParams) SetMediaSid(MediaSid string) (*CreateServiceConversationMessageParams){
    params.MediaSid = &MediaSid
    return params
}
func (params *CreateServiceConversationMessageParams) SetContentSid(ContentSid string) (*CreateServiceConversationMessageParams){
    params.ContentSid = &ContentSid
    return params
}
func (params *CreateServiceConversationMessageParams) SetContentVariables(ContentVariables string) (*CreateServiceConversationMessageParams){
    params.ContentVariables = &ContentVariables
    return params
}

// Add a new message to the conversation in a specific service
func (c *ApiService) CreateServiceConversationMessage(ChatServiceSid string, ConversationSid string, params *CreateServiceConversationMessageParams) (*ConversationsV1ServiceConversationMessage, error) {
    path := "/v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages"
        path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
    path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

data := url.Values{}
headers := make(map[string]interface{})

if params != nil && params.Author != nil {
    data.Set("Author", *params.Author)
}
if params != nil && params.Body != nil {
    data.Set("Body", *params.Body)
}
if params != nil && params.DateCreated != nil {
    data.Set("DateCreated", fmt.Sprint((*params.DateCreated).Format(time.RFC3339)))
}
if params != nil && params.DateUpdated != nil {
    data.Set("DateUpdated", fmt.Sprint((*params.DateUpdated).Format(time.RFC3339)))
}
if params != nil && params.Attributes != nil {
    data.Set("Attributes", *params.Attributes)
}
if params != nil && params.MediaSid != nil {
    data.Set("MediaSid", *params.MediaSid)
}
if params != nil && params.ContentSid != nil {
    data.Set("ContentSid", *params.ContentSid)
}
if params != nil && params.ContentVariables != nil {
    data.Set("ContentVariables", *params.ContentVariables)
}


	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ConversationsV1ServiceConversationMessage{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'DeleteServiceConversationMessage'
type DeleteServiceConversationMessageParams struct {
    // The X-Twilio-Webhook-Enabled HTTP request header
    XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
}

func (params *DeleteServiceConversationMessageParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) (*DeleteServiceConversationMessageParams){
    params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
    return params
}

// Remove a message from the conversation
func (c *ApiService) DeleteServiceConversationMessage(ChatServiceSid string, ConversationSid string, Sid string, params *DeleteServiceConversationMessageParams) (error) {
    path := "/v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{Sid}"
        path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
    path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
    path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

data := url.Values{}
headers := make(map[string]interface{})



	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

    resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
    if err != nil {
        return err
    }

    defer resp.Body.Close()

    return nil
}

// Fetch a message from the conversation
func (c *ApiService) FetchServiceConversationMessage(ChatServiceSid string, ConversationSid string, Sid string, ) (*ConversationsV1ServiceConversationMessage, error) {
    path := "/v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{Sid}"
        path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
    path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
    path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

data := url.Values{}
headers := make(map[string]interface{})




    resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ConversationsV1ServiceConversationMessage{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Optional parameters for the method 'ListServiceConversationMessage'
type ListServiceConversationMessageParams struct {
    // The sort order of the returned messages. Can be: `asc` (ascending) or `desc` (descending), with `asc` as the default.
    Order *string `json:"Order,omitempty"`
    // How many resources to return in each list page. The default is 50, and the maximum is 1000.
    PageSize *int `json:"PageSize,omitempty"`
    // Max number of records to return.
    Limit *int `json:"limit,omitempty"`
}

func (params *ListServiceConversationMessageParams) SetOrder(Order string) (*ListServiceConversationMessageParams){
    params.Order = &Order
    return params
}
func (params *ListServiceConversationMessageParams) SetPageSize(PageSize int) (*ListServiceConversationMessageParams){
    params.PageSize = &PageSize
    return params
}
func (params *ListServiceConversationMessageParams) SetLimit(Limit int) (*ListServiceConversationMessageParams){
    params.Limit = &Limit
    return params
}

// Retrieve a single page of ServiceConversationMessage records from the API. Request is executed immediately.
func (c *ApiService) PageServiceConversationMessage(ChatServiceSid string, ConversationSid string, params *ListServiceConversationMessageParams, pageToken, pageNumber string) (*ListServiceConversationMessageResponse, error) {
    path := "/v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages"

        path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
    path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

data := url.Values{}
headers := make(map[string]interface{})

if params != nil && params.Order != nil {
    data.Set("Order", *params.Order)
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

    ps := &ListServiceConversationMessageResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

// Lists ServiceConversationMessage records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListServiceConversationMessage(ChatServiceSid string, ConversationSid string, params *ListServiceConversationMessageParams) ([]ConversationsV1ServiceConversationMessage, error) {
	response, errors := c.StreamServiceConversationMessage(ChatServiceSid, ConversationSid, params)

	records := make([]ConversationsV1ServiceConversationMessage, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams ServiceConversationMessage records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamServiceConversationMessage(ChatServiceSid string, ConversationSid string, params *ListServiceConversationMessageParams) (chan ConversationsV1ServiceConversationMessage, chan error) {
	if params == nil {
		params = &ListServiceConversationMessageParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ConversationsV1ServiceConversationMessage, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageServiceConversationMessage(ChatServiceSid, ConversationSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamServiceConversationMessage(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}


func (c *ApiService) streamServiceConversationMessage(response *ListServiceConversationMessageResponse, params *ListServiceConversationMessageParams, recordChannel chan ConversationsV1ServiceConversationMessage, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Messages
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListServiceConversationMessageResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListServiceConversationMessageResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListServiceConversationMessageResponse(nextPageUrl string) (interface{}, error) {
    if nextPageUrl == "" {
        return nil, nil
    }
    resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ListServiceConversationMessageResponse{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }
    return ps, nil
}


// Optional parameters for the method 'UpdateServiceConversationMessage'
type UpdateServiceConversationMessageParams struct {
    // The X-Twilio-Webhook-Enabled HTTP request header
    XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
    // The channel specific identifier of the message's author. Defaults to `system`.
    Author *string `json:"Author,omitempty"`
    // The content of the message, can be up to 1,600 characters long.
    Body *string `json:"Body,omitempty"`
    // The date that this resource was created.
    DateCreated *time.Time `json:"DateCreated,omitempty"`
    // The date that this resource was last updated. `null` if the message has not been edited.
    DateUpdated *time.Time `json:"DateUpdated,omitempty"`
    // A string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \\\"{}\\\" will be returned.
    Attributes *string `json:"Attributes,omitempty"`
}

func (params *UpdateServiceConversationMessageParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) (*UpdateServiceConversationMessageParams){
    params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
    return params
}
func (params *UpdateServiceConversationMessageParams) SetAuthor(Author string) (*UpdateServiceConversationMessageParams){
    params.Author = &Author
    return params
}
func (params *UpdateServiceConversationMessageParams) SetBody(Body string) (*UpdateServiceConversationMessageParams){
    params.Body = &Body
    return params
}
func (params *UpdateServiceConversationMessageParams) SetDateCreated(DateCreated time.Time) (*UpdateServiceConversationMessageParams){
    params.DateCreated = &DateCreated
    return params
}
func (params *UpdateServiceConversationMessageParams) SetDateUpdated(DateUpdated time.Time) (*UpdateServiceConversationMessageParams){
    params.DateUpdated = &DateUpdated
    return params
}
func (params *UpdateServiceConversationMessageParams) SetAttributes(Attributes string) (*UpdateServiceConversationMessageParams){
    params.Attributes = &Attributes
    return params
}

// Update an existing message in the conversation
func (c *ApiService) UpdateServiceConversationMessage(ChatServiceSid string, ConversationSid string, Sid string, params *UpdateServiceConversationMessageParams) (*ConversationsV1ServiceConversationMessage, error) {
    path := "/v1/Services/{ChatServiceSid}/Conversations/{ConversationSid}/Messages/{Sid}"
        path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
    path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
    path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

data := url.Values{}
headers := make(map[string]interface{})

if params != nil && params.Author != nil {
    data.Set("Author", *params.Author)
}
if params != nil && params.Body != nil {
    data.Set("Body", *params.Body)
}
if params != nil && params.DateCreated != nil {
    data.Set("DateCreated", fmt.Sprint((*params.DateCreated).Format(time.RFC3339)))
}
if params != nil && params.DateUpdated != nil {
    data.Set("DateUpdated", fmt.Sprint((*params.DateUpdated).Format(time.RFC3339)))
}
if params != nil && params.Attributes != nil {
    data.Set("Attributes", *params.Attributes)
}


	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

    resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &ConversationsV1ServiceConversationMessage{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
