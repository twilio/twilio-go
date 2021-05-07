# ChatV1Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ConsumptionReportInterval** | Pointer to **int32** | DEPRECATED |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**DefaultChannelCreatorRoleSid** | Pointer to **string** | The channel role assigned to a channel creator when they join a new channel |
**DefaultChannelRoleSid** | Pointer to **string** | The channel role assigned to users when they are added to a channel |
**DefaultServiceRoleSid** | Pointer to **string** | The service role assigned to users when they are added to the service |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**Limits** | Pointer to **map[string]interface{}** | An object that describes the limits of the service instance |
**Links** | Pointer to **map[string]interface{}** | The absolute URLs of the Service's Channels, Roles, and Users |
**Notifications** | Pointer to **map[string]interface{}** | The notification configuration for the Service instance |
**PostWebhookUrl** | Pointer to **string** | The URL for post-event webhooks |
**PreWebhookUrl** | Pointer to **string** | The webhook URL for pre-event webhooks |
**ReachabilityEnabled** | Pointer to **bool** | Whether the Reachability Indicator feature is enabled for this Service instance |
**ReadStatusEnabled** | Pointer to **bool** | Whether the Message Consumption Horizon feature is enabled |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**TypingIndicatorTimeout** | Pointer to **int32** | How long in seconds to wait before assuming the user is no longer typing |
**Url** | Pointer to **string** | The absolute URL of the Service resource |
**WebhookFilters** | Pointer to **[]string** | The list of WebHook events that are enabled for this Service instance |
**WebhookMethod** | Pointer to **string** | The HTTP method  to use for both PRE and POST webhooks |
**Webhooks** | Pointer to **map[string]interface{}** | An object that contains information about the webhooks configured for this service |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


