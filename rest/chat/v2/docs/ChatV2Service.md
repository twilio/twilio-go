# ChatV2Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |[optional] 
**DefaultServiceRoleSid** | **string** | The service role assigned to users when they are added to the service |[optional] 
**DefaultChannelRoleSid** | **string** | The channel role assigned to users when they are added to a channel |[optional] 
**DefaultChannelCreatorRoleSid** | **string** | The channel role assigned to a channel creator when they join a new channel |[optional] 
**ReadStatusEnabled** | **bool** | Whether the Message Consumption Horizon feature is enabled |[optional] 
**ReachabilityEnabled** | **bool** | Whether the Reachability Indicator feature is enabled for this Service instance |[optional] 
**TypingIndicatorTimeout** | **int** | How long in seconds to wait before assuming the user is no longer typing |[optional] 
**ConsumptionReportInterval** | **int** | DEPRECATED |[optional] 
**Limits** | Pointer to **interface{}** | An object that describes the limits of the service instance |
**PreWebhookUrl** | **string** | The webhook URL for pre-event webhooks |[optional] 
**PostWebhookUrl** | **string** | The URL for post-event webhooks |[optional] 
**WebhookMethod** | **string** | The HTTP method  to use for both PRE and POST webhooks |[optional] 
**WebhookFilters** | **[]string** | The list of webhook events that are enabled for this Service instance |[optional] 
**PreWebhookRetryCount** | **int** | Count of times webhook will be retried in case of timeout or 429/503/504 HTTP responses |[optional] 
**PostWebhookRetryCount** | **int** | The number of times calls to the `post_webhook_url` will be retried |[optional] 
**Notifications** | Pointer to **interface{}** | The notification configuration for the Service instance |
**Media** | Pointer to **interface{}** | The properties of the media that the service supports |
**Url** | **string** | The absolute URL of the Service resource |[optional] 
**Links** | **map[string]interface{}** | The absolute URLs of the Service's Channels, Roles, and Users |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


