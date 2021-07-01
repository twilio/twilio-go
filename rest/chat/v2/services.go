/*
 * Twilio - Chat
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

// Optional parameters for the method 'CreateService'
type CreateServiceParams struct {
	// A descriptive string that you create to describe the new resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateServiceParams) SetFriendlyName(FriendlyName string) *CreateServiceParams {
	params.FriendlyName = &FriendlyName
	return params
}

func (c *ApiService) CreateService(params *CreateServiceParams) (*ChatV2Service, error) {
	path := "/v2/Services"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ChatV2Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteService(Sid string) error {
	path := "/v2/Services/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ApiService) FetchService(Sid string) (*ChatV2Service, error) {
	path := "/v2/Services/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ChatV2Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListService'
type ListServiceParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListServiceParams) SetPageSize(PageSize int) *ListServiceParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListService(params *ListServiceParams) (*ListServiceResponse, error) {
	path := "/v2/Services"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListServiceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of Service records from the API. Request is executed immediately.
func (c *ApiService) ServicePage(params *ListServiceParams, pageToken string, pageNumber string, pageSize string) *client.Page {
	path := "/v2/Services"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	data.Set("PageToken", pageToken)
	data.Set("PageNumber", pageNumber)
	data.Set("PageSize", pageSize)

	response, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil
	}

	return client.NewPage(c.baseURL, response)
}

//Streams Service records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) ServiceStream(params *ListServiceParams, meta client.PaginationData) chan map[string]interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.ServicePage(params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.Stream(page, limits.Limit, limits.PageLimit)
}

//Lists Service records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) ServiceList(params *ListServiceParams, meta client.PaginationData) []interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.ServicePage(params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.List(page, limits.Limit, limits.PageLimit)
}

// Optional parameters for the method 'UpdateService'
type UpdateServiceParams struct {
	// DEPRECATED. The interval in seconds between consumption reports submission batches from client endpoints.
	ConsumptionReportInterval *int `json:"ConsumptionReportInterval,omitempty"`
	// The channel role assigned to a channel creator when they join a new channel. See the [Role resource](https://www.twilio.com/docs/chat/rest/role-resource) for more info about roles.
	DefaultChannelCreatorRoleSid *string `json:"DefaultChannelCreatorRoleSid,omitempty"`
	// The channel role assigned to users when they are added to a channel. See the [Role resource](https://www.twilio.com/docs/chat/rest/role-resource) for more info about roles.
	DefaultChannelRoleSid *string `json:"DefaultChannelRoleSid,omitempty"`
	// The service role assigned to users when they are added to the service. See the [Role resource](https://www.twilio.com/docs/chat/rest/role-resource) for more info about roles.
	DefaultServiceRoleSid *string `json:"DefaultServiceRoleSid,omitempty"`
	// A descriptive string that you create to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The maximum number of Members that can be added to Channels within this Service. Can be up to 1,000.
	LimitsChannelMembers *int `json:"Limits.ChannelMembers,omitempty"`
	// The maximum number of Channels Users can be a Member of within this Service. Can be up to 1,000.
	LimitsUserChannels *int `json:"Limits.UserChannels,omitempty"`
	// The message to send when a media message has no text. Can be used as placeholder message.
	MediaCompatibilityMessage *string `json:"Media.CompatibilityMessage,omitempty"`
	// Whether to send a notification when a member is added to a channel. The default is `false`.
	NotificationsAddedToChannelEnabled *bool `json:"Notifications.AddedToChannel.Enabled,omitempty"`
	// The name of the sound to play when a member is added to a channel and `notifications.added_to_channel.enabled` is `true`.
	NotificationsAddedToChannelSound *string `json:"Notifications.AddedToChannel.Sound,omitempty"`
	// The template to use to create the notification text displayed when a member is added to a channel and `notifications.added_to_channel.enabled` is `true`.
	NotificationsAddedToChannelTemplate *string `json:"Notifications.AddedToChannel.Template,omitempty"`
	// Whether to send a notification when a user is invited to a channel. The default is `false`.
	NotificationsInvitedToChannelEnabled *bool `json:"Notifications.InvitedToChannel.Enabled,omitempty"`
	// The name of the sound to play when a user is invited to a channel and `notifications.invited_to_channel.enabled` is `true`.
	NotificationsInvitedToChannelSound *string `json:"Notifications.InvitedToChannel.Sound,omitempty"`
	// The template to use to create the notification text displayed when a user is invited to a channel and `notifications.invited_to_channel.enabled` is `true`.
	NotificationsInvitedToChannelTemplate *string `json:"Notifications.InvitedToChannel.Template,omitempty"`
	// Whether to log notifications. The default is `false`.
	NotificationsLogEnabled *bool `json:"Notifications.LogEnabled,omitempty"`
	// Whether the new message badge is enabled. The default is `false`.
	NotificationsNewMessageBadgeCountEnabled *bool `json:"Notifications.NewMessage.BadgeCountEnabled,omitempty"`
	// Whether to send a notification when a new message is added to a channel. The default is `false`.
	NotificationsNewMessageEnabled *bool `json:"Notifications.NewMessage.Enabled,omitempty"`
	// The name of the sound to play when a new message is added to a channel and `notifications.new_message.enabled` is `true`.
	NotificationsNewMessageSound *string `json:"Notifications.NewMessage.Sound,omitempty"`
	// The template to use to create the notification text displayed when a new message is added to a channel and `notifications.new_message.enabled` is `true`.
	NotificationsNewMessageTemplate *string `json:"Notifications.NewMessage.Template,omitempty"`
	// Whether to send a notification to a user when they are removed from a channel. The default is `false`.
	NotificationsRemovedFromChannelEnabled *bool `json:"Notifications.RemovedFromChannel.Enabled,omitempty"`
	// The name of the sound to play to a user when they are removed from a channel and `notifications.removed_from_channel.enabled` is `true`.
	NotificationsRemovedFromChannelSound *string `json:"Notifications.RemovedFromChannel.Sound,omitempty"`
	// The template to use to create the notification text displayed to a user when they are removed from a channel and `notifications.removed_from_channel.enabled` is `true`.
	NotificationsRemovedFromChannelTemplate *string `json:"Notifications.RemovedFromChannel.Template,omitempty"`
	// The number of times to retry a call to the `post_webhook_url` if the request times out (after 5 seconds) or it receives a 429, 503, or 504 HTTP response. The default is 0, which means the call won't be retried.
	PostWebhookRetryCount *int `json:"PostWebhookRetryCount,omitempty"`
	// The URL for post-event webhooks, which are called by using the `webhook_method`. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
	PostWebhookUrl *string `json:"PostWebhookUrl,omitempty"`
	// The number of times to retry a call to the `pre_webhook_url` if the request times out (after 5 seconds) or it receives a 429, 503, or 504 HTTP response. Default retry count is 0 times, which means the call won't be retried.
	PreWebhookRetryCount *int `json:"PreWebhookRetryCount,omitempty"`
	// The URL for pre-event webhooks, which are called by using the `webhook_method`. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
	PreWebhookUrl *string `json:"PreWebhookUrl,omitempty"`
	// Whether to enable the [Reachability Indicator](https://www.twilio.com/docs/chat/reachability-indicator) for this Service instance. The default is `false`.
	ReachabilityEnabled *bool `json:"ReachabilityEnabled,omitempty"`
	// Whether to enable the [Message Consumption Horizon](https://www.twilio.com/docs/chat/consumption-horizon) feature. The default is `true`.
	ReadStatusEnabled *bool `json:"ReadStatusEnabled,omitempty"`
	// How long in seconds after a `started typing` event until clients should assume that user is no longer typing, even if no `ended typing` message was received.  The default is 5 seconds.
	TypingIndicatorTimeout *int `json:"TypingIndicatorTimeout,omitempty"`
	// The list of webhook events that are enabled for this Service instance. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
	WebhookFilters *[]string `json:"WebhookFilters,omitempty"`
	// The HTTP method to use for calls to the `pre_webhook_url` and `post_webhook_url` webhooks.  Can be: `POST` or `GET` and the default is `POST`. See [Webhook Events](https://www.twilio.com/docs/chat/webhook-events) for more details.
	WebhookMethod *string `json:"WebhookMethod,omitempty"`
}

func (params *UpdateServiceParams) SetConsumptionReportInterval(ConsumptionReportInterval int) *UpdateServiceParams {
	params.ConsumptionReportInterval = &ConsumptionReportInterval
	return params
}
func (params *UpdateServiceParams) SetDefaultChannelCreatorRoleSid(DefaultChannelCreatorRoleSid string) *UpdateServiceParams {
	params.DefaultChannelCreatorRoleSid = &DefaultChannelCreatorRoleSid
	return params
}
func (params *UpdateServiceParams) SetDefaultChannelRoleSid(DefaultChannelRoleSid string) *UpdateServiceParams {
	params.DefaultChannelRoleSid = &DefaultChannelRoleSid
	return params
}
func (params *UpdateServiceParams) SetDefaultServiceRoleSid(DefaultServiceRoleSid string) *UpdateServiceParams {
	params.DefaultServiceRoleSid = &DefaultServiceRoleSid
	return params
}
func (params *UpdateServiceParams) SetFriendlyName(FriendlyName string) *UpdateServiceParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateServiceParams) SetLimitsChannelMembers(LimitsChannelMembers int) *UpdateServiceParams {
	params.LimitsChannelMembers = &LimitsChannelMembers
	return params
}
func (params *UpdateServiceParams) SetLimitsUserChannels(LimitsUserChannels int) *UpdateServiceParams {
	params.LimitsUserChannels = &LimitsUserChannels
	return params
}
func (params *UpdateServiceParams) SetMediaCompatibilityMessage(MediaCompatibilityMessage string) *UpdateServiceParams {
	params.MediaCompatibilityMessage = &MediaCompatibilityMessage
	return params
}
func (params *UpdateServiceParams) SetNotificationsAddedToChannelEnabled(NotificationsAddedToChannelEnabled bool) *UpdateServiceParams {
	params.NotificationsAddedToChannelEnabled = &NotificationsAddedToChannelEnabled
	return params
}
func (params *UpdateServiceParams) SetNotificationsAddedToChannelSound(NotificationsAddedToChannelSound string) *UpdateServiceParams {
	params.NotificationsAddedToChannelSound = &NotificationsAddedToChannelSound
	return params
}
func (params *UpdateServiceParams) SetNotificationsAddedToChannelTemplate(NotificationsAddedToChannelTemplate string) *UpdateServiceParams {
	params.NotificationsAddedToChannelTemplate = &NotificationsAddedToChannelTemplate
	return params
}
func (params *UpdateServiceParams) SetNotificationsInvitedToChannelEnabled(NotificationsInvitedToChannelEnabled bool) *UpdateServiceParams {
	params.NotificationsInvitedToChannelEnabled = &NotificationsInvitedToChannelEnabled
	return params
}
func (params *UpdateServiceParams) SetNotificationsInvitedToChannelSound(NotificationsInvitedToChannelSound string) *UpdateServiceParams {
	params.NotificationsInvitedToChannelSound = &NotificationsInvitedToChannelSound
	return params
}
func (params *UpdateServiceParams) SetNotificationsInvitedToChannelTemplate(NotificationsInvitedToChannelTemplate string) *UpdateServiceParams {
	params.NotificationsInvitedToChannelTemplate = &NotificationsInvitedToChannelTemplate
	return params
}
func (params *UpdateServiceParams) SetNotificationsLogEnabled(NotificationsLogEnabled bool) *UpdateServiceParams {
	params.NotificationsLogEnabled = &NotificationsLogEnabled
	return params
}
func (params *UpdateServiceParams) SetNotificationsNewMessageBadgeCountEnabled(NotificationsNewMessageBadgeCountEnabled bool) *UpdateServiceParams {
	params.NotificationsNewMessageBadgeCountEnabled = &NotificationsNewMessageBadgeCountEnabled
	return params
}
func (params *UpdateServiceParams) SetNotificationsNewMessageEnabled(NotificationsNewMessageEnabled bool) *UpdateServiceParams {
	params.NotificationsNewMessageEnabled = &NotificationsNewMessageEnabled
	return params
}
func (params *UpdateServiceParams) SetNotificationsNewMessageSound(NotificationsNewMessageSound string) *UpdateServiceParams {
	params.NotificationsNewMessageSound = &NotificationsNewMessageSound
	return params
}
func (params *UpdateServiceParams) SetNotificationsNewMessageTemplate(NotificationsNewMessageTemplate string) *UpdateServiceParams {
	params.NotificationsNewMessageTemplate = &NotificationsNewMessageTemplate
	return params
}
func (params *UpdateServiceParams) SetNotificationsRemovedFromChannelEnabled(NotificationsRemovedFromChannelEnabled bool) *UpdateServiceParams {
	params.NotificationsRemovedFromChannelEnabled = &NotificationsRemovedFromChannelEnabled
	return params
}
func (params *UpdateServiceParams) SetNotificationsRemovedFromChannelSound(NotificationsRemovedFromChannelSound string) *UpdateServiceParams {
	params.NotificationsRemovedFromChannelSound = &NotificationsRemovedFromChannelSound
	return params
}
func (params *UpdateServiceParams) SetNotificationsRemovedFromChannelTemplate(NotificationsRemovedFromChannelTemplate string) *UpdateServiceParams {
	params.NotificationsRemovedFromChannelTemplate = &NotificationsRemovedFromChannelTemplate
	return params
}
func (params *UpdateServiceParams) SetPostWebhookRetryCount(PostWebhookRetryCount int) *UpdateServiceParams {
	params.PostWebhookRetryCount = &PostWebhookRetryCount
	return params
}
func (params *UpdateServiceParams) SetPostWebhookUrl(PostWebhookUrl string) *UpdateServiceParams {
	params.PostWebhookUrl = &PostWebhookUrl
	return params
}
func (params *UpdateServiceParams) SetPreWebhookRetryCount(PreWebhookRetryCount int) *UpdateServiceParams {
	params.PreWebhookRetryCount = &PreWebhookRetryCount
	return params
}
func (params *UpdateServiceParams) SetPreWebhookUrl(PreWebhookUrl string) *UpdateServiceParams {
	params.PreWebhookUrl = &PreWebhookUrl
	return params
}
func (params *UpdateServiceParams) SetReachabilityEnabled(ReachabilityEnabled bool) *UpdateServiceParams {
	params.ReachabilityEnabled = &ReachabilityEnabled
	return params
}
func (params *UpdateServiceParams) SetReadStatusEnabled(ReadStatusEnabled bool) *UpdateServiceParams {
	params.ReadStatusEnabled = &ReadStatusEnabled
	return params
}
func (params *UpdateServiceParams) SetTypingIndicatorTimeout(TypingIndicatorTimeout int) *UpdateServiceParams {
	params.TypingIndicatorTimeout = &TypingIndicatorTimeout
	return params
}
func (params *UpdateServiceParams) SetWebhookFilters(WebhookFilters []string) *UpdateServiceParams {
	params.WebhookFilters = &WebhookFilters
	return params
}
func (params *UpdateServiceParams) SetWebhookMethod(WebhookMethod string) *UpdateServiceParams {
	params.WebhookMethod = &WebhookMethod
	return params
}

func (c *ApiService) UpdateService(Sid string, params *UpdateServiceParams) (*ChatV2Service, error) {
	path := "/v2/Services/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ConsumptionReportInterval != nil {
		data.Set("ConsumptionReportInterval", fmt.Sprint(*params.ConsumptionReportInterval))
	}
	if params != nil && params.DefaultChannelCreatorRoleSid != nil {
		data.Set("DefaultChannelCreatorRoleSid", *params.DefaultChannelCreatorRoleSid)
	}
	if params != nil && params.DefaultChannelRoleSid != nil {
		data.Set("DefaultChannelRoleSid", *params.DefaultChannelRoleSid)
	}
	if params != nil && params.DefaultServiceRoleSid != nil {
		data.Set("DefaultServiceRoleSid", *params.DefaultServiceRoleSid)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.LimitsChannelMembers != nil {
		data.Set("Limits.ChannelMembers", fmt.Sprint(*params.LimitsChannelMembers))
	}
	if params != nil && params.LimitsUserChannels != nil {
		data.Set("Limits.UserChannels", fmt.Sprint(*params.LimitsUserChannels))
	}
	if params != nil && params.MediaCompatibilityMessage != nil {
		data.Set("Media.CompatibilityMessage", *params.MediaCompatibilityMessage)
	}
	if params != nil && params.NotificationsAddedToChannelEnabled != nil {
		data.Set("Notifications.AddedToChannel.Enabled", fmt.Sprint(*params.NotificationsAddedToChannelEnabled))
	}
	if params != nil && params.NotificationsAddedToChannelSound != nil {
		data.Set("Notifications.AddedToChannel.Sound", *params.NotificationsAddedToChannelSound)
	}
	if params != nil && params.NotificationsAddedToChannelTemplate != nil {
		data.Set("Notifications.AddedToChannel.Template", *params.NotificationsAddedToChannelTemplate)
	}
	if params != nil && params.NotificationsInvitedToChannelEnabled != nil {
		data.Set("Notifications.InvitedToChannel.Enabled", fmt.Sprint(*params.NotificationsInvitedToChannelEnabled))
	}
	if params != nil && params.NotificationsInvitedToChannelSound != nil {
		data.Set("Notifications.InvitedToChannel.Sound", *params.NotificationsInvitedToChannelSound)
	}
	if params != nil && params.NotificationsInvitedToChannelTemplate != nil {
		data.Set("Notifications.InvitedToChannel.Template", *params.NotificationsInvitedToChannelTemplate)
	}
	if params != nil && params.NotificationsLogEnabled != nil {
		data.Set("Notifications.LogEnabled", fmt.Sprint(*params.NotificationsLogEnabled))
	}
	if params != nil && params.NotificationsNewMessageBadgeCountEnabled != nil {
		data.Set("Notifications.NewMessage.BadgeCountEnabled", fmt.Sprint(*params.NotificationsNewMessageBadgeCountEnabled))
	}
	if params != nil && params.NotificationsNewMessageEnabled != nil {
		data.Set("Notifications.NewMessage.Enabled", fmt.Sprint(*params.NotificationsNewMessageEnabled))
	}
	if params != nil && params.NotificationsNewMessageSound != nil {
		data.Set("Notifications.NewMessage.Sound", *params.NotificationsNewMessageSound)
	}
	if params != nil && params.NotificationsNewMessageTemplate != nil {
		data.Set("Notifications.NewMessage.Template", *params.NotificationsNewMessageTemplate)
	}
	if params != nil && params.NotificationsRemovedFromChannelEnabled != nil {
		data.Set("Notifications.RemovedFromChannel.Enabled", fmt.Sprint(*params.NotificationsRemovedFromChannelEnabled))
	}
	if params != nil && params.NotificationsRemovedFromChannelSound != nil {
		data.Set("Notifications.RemovedFromChannel.Sound", *params.NotificationsRemovedFromChannelSound)
	}
	if params != nil && params.NotificationsRemovedFromChannelTemplate != nil {
		data.Set("Notifications.RemovedFromChannel.Template", *params.NotificationsRemovedFromChannelTemplate)
	}
	if params != nil && params.PostWebhookRetryCount != nil {
		data.Set("PostWebhookRetryCount", fmt.Sprint(*params.PostWebhookRetryCount))
	}
	if params != nil && params.PostWebhookUrl != nil {
		data.Set("PostWebhookUrl", *params.PostWebhookUrl)
	}
	if params != nil && params.PreWebhookRetryCount != nil {
		data.Set("PreWebhookRetryCount", fmt.Sprint(*params.PreWebhookRetryCount))
	}
	if params != nil && params.PreWebhookUrl != nil {
		data.Set("PreWebhookUrl", *params.PreWebhookUrl)
	}
	if params != nil && params.ReachabilityEnabled != nil {
		data.Set("ReachabilityEnabled", fmt.Sprint(*params.ReachabilityEnabled))
	}
	if params != nil && params.ReadStatusEnabled != nil {
		data.Set("ReadStatusEnabled", fmt.Sprint(*params.ReadStatusEnabled))
	}
	if params != nil && params.TypingIndicatorTimeout != nil {
		data.Set("TypingIndicatorTimeout", fmt.Sprint(*params.TypingIndicatorTimeout))
	}
	if params != nil && params.WebhookFilters != nil {
		for _, item := range *params.WebhookFilters {
			data.Add("WebhookFilters", item)
		}
	}
	if params != nil && params.WebhookMethod != nil {
		data.Set("WebhookMethod", *params.WebhookMethod)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ChatV2Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
