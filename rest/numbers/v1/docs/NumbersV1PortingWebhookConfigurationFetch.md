# NumbersV1PortingWebhookConfigurationFetch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The URL of the webhook configuration request |
**PortInTargetUrl** | Pointer to **string** | The complete webhook url that will be called when a notification event for port in request or port in phone number happens |
**PortOutTargetUrl** | Pointer to **string** | The complete webhook url that will be called when a notification event for a port out phone number happens. |
**NotificationsOf** | Pointer to **[]string** | A list to filter what notification events to receive for this account and its sub accounts. If it is an empty list, then it means that there are no filters for the notifications events to send in each webhook and all events will get sent. |
**PortInTargetDateCreated** | Pointer to [**time.Time**](time.Time.md) | Creation date for the port in webhook configuration |
**PortOutTargetDateCreated** | Pointer to [**time.Time**](time.Time.md) | Creation date for the port out webhook configuration |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


