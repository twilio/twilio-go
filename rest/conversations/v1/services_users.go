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
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateServiceUser'
type CreateServiceUserParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	// The application-defined string that uniquely identifies the resource's User within the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource). This value is often a username or an email address, and is case-sensitive.
	Identity *string `json:"Identity,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The JSON Object string that stores application-specific data. If attributes have not been set, `{}` is returned.
	Attributes *string `json:"Attributes,omitempty"`
	// The SID of a service-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the user.
	RoleSid *string `json:"RoleSid,omitempty"`
}

func (params *CreateServiceUserParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *CreateServiceUserParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *CreateServiceUserParams) SetIdentity(Identity string) *CreateServiceUserParams {
	params.Identity = &Identity
	return params
}
func (params *CreateServiceUserParams) SetFriendlyName(FriendlyName string) *CreateServiceUserParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateServiceUserParams) SetAttributes(Attributes string) *CreateServiceUserParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateServiceUserParams) SetRoleSid(RoleSid string) *CreateServiceUserParams {
	params.RoleSid = &RoleSid
	return params
}

// Add a new conversation user to your service
func (c *ApiService) CreateServiceUser(ChatServiceSid string, params *CreateServiceUserParams) (*ConversationsV1ServiceUser, error) {
	return c.CreateServiceUserWithCtx(context.TODO(), ChatServiceSid, params)
}

// Add a new conversation user to your service
func (c *ApiService) CreateServiceUserWithCtx(ctx context.Context, ChatServiceSid string, params *CreateServiceUserParams) (*ConversationsV1ServiceUser, error) {
	path := "/v1/Services/{ChatServiceSid}/Users"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Identity != nil {
		data.Set("Identity", *params.Identity)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.RoleSid != nil {
		data.Set("RoleSid", *params.RoleSid)
	}

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteServiceUser'
type DeleteServiceUserParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
}

func (params *DeleteServiceUserParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *DeleteServiceUserParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}

// Remove a conversation user from your service
func (c *ApiService) DeleteServiceUser(ChatServiceSid string, Sid string, params *DeleteServiceUserParams) error {
	return c.DeleteServiceUserWithCtx(context.TODO(), ChatServiceSid, Sid, params)
}

// Remove a conversation user from your service
func (c *ApiService) DeleteServiceUserWithCtx(ctx context.Context, ChatServiceSid string, Sid string, params *DeleteServiceUserParams) error {
	path := "/v1/Services/{ChatServiceSid}/Users/{Sid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Delete(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch a conversation user from your service
func (c *ApiService) FetchServiceUser(ChatServiceSid string, Sid string) (*ConversationsV1ServiceUser, error) {
	return c.FetchServiceUserWithCtx(context.TODO(), ChatServiceSid, Sid)
}

// Fetch a conversation user from your service
func (c *ApiService) FetchServiceUserWithCtx(ctx context.Context, ChatServiceSid string, Sid string) (*ConversationsV1ServiceUser, error) {
	path := "/v1/Services/{ChatServiceSid}/Users/{Sid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListServiceUser'
type ListServiceUserParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListServiceUserParams) SetPageSize(PageSize int) *ListServiceUserParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListServiceUserParams) SetLimit(Limit int) *ListServiceUserParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ServiceUser records from the API. Request is executed immediately.
func (c *ApiService) PageServiceUser(ChatServiceSid string, params *ListServiceUserParams, pageToken, pageNumber string) (*ListServiceUserResponse, error) {
	return c.PageServiceUserWithCtx(context.TODO(), ChatServiceSid, params, pageToken, pageNumber)
}

// Retrieve a single page of ServiceUser records from the API. Request is executed immediately.
func (c *ApiService) PageServiceUserWithCtx(ctx context.Context, ChatServiceSid string, params *ListServiceUserParams, pageToken, pageNumber string) (*ListServiceUserResponse, error) {
	path := "/v1/Services/{ChatServiceSid}/Users"

	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListServiceUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ServiceUser records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListServiceUser(ChatServiceSid string, params *ListServiceUserParams) ([]ConversationsV1ServiceUser, error) {
	return c.ListServiceUserWithCtx(context.TODO(), ChatServiceSid, params)
}

// Lists ServiceUser records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListServiceUserWithCtx(ctx context.Context, ChatServiceSid string, params *ListServiceUserParams) ([]ConversationsV1ServiceUser, error) {
	response, errors := c.StreamServiceUserWithCtx(ctx, ChatServiceSid, params)

	records := make([]ConversationsV1ServiceUser, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams ServiceUser records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamServiceUser(ChatServiceSid string, params *ListServiceUserParams) (chan ConversationsV1ServiceUser, chan error) {
	return c.StreamServiceUserWithCtx(context.TODO(), ChatServiceSid, params)
}

// Streams ServiceUser records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamServiceUserWithCtx(ctx context.Context, ChatServiceSid string, params *ListServiceUserParams) (chan ConversationsV1ServiceUser, chan error) {
	if params == nil {
		params = &ListServiceUserParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ConversationsV1ServiceUser, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageServiceUserWithCtx(ctx, ChatServiceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamServiceUser(ctx, response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamServiceUser(ctx context.Context, response *ListServiceUserResponse, params *ListServiceUserParams, recordChannel chan ConversationsV1ServiceUser, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Users
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNextWithCtx(ctx, c.baseURL, response, c.getNextListServiceUserResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListServiceUserResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListServiceUserResponse(ctx context.Context, nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(ctx, nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListServiceUserResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateServiceUser'
type UpdateServiceUserParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The JSON Object string that stores application-specific data. If attributes have not been set, `{}` is returned.
	Attributes *string `json:"Attributes,omitempty"`
	// The SID of a service-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the user.
	RoleSid *string `json:"RoleSid,omitempty"`
}

func (params *UpdateServiceUserParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *UpdateServiceUserParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *UpdateServiceUserParams) SetFriendlyName(FriendlyName string) *UpdateServiceUserParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateServiceUserParams) SetAttributes(Attributes string) *UpdateServiceUserParams {
	params.Attributes = &Attributes
	return params
}
func (params *UpdateServiceUserParams) SetRoleSid(RoleSid string) *UpdateServiceUserParams {
	params.RoleSid = &RoleSid
	return params
}

// Update an existing conversation user in your service
func (c *ApiService) UpdateServiceUser(ChatServiceSid string, Sid string, params *UpdateServiceUserParams) (*ConversationsV1ServiceUser, error) {
	return c.UpdateServiceUserWithCtx(context.TODO(), ChatServiceSid, Sid, params)
}

// Update an existing conversation user in your service
func (c *ApiService) UpdateServiceUserWithCtx(ctx context.Context, ChatServiceSid string, Sid string, params *UpdateServiceUserParams) (*ConversationsV1ServiceUser, error) {
	path := "/v1/Services/{ChatServiceSid}/Users/{Sid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.RoleSid != nil {
		data.Set("RoleSid", *params.RoleSid)
	}

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceUser{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
