# NumbersV1PortingWebhookConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The URL of the webhook configuration request |
**PortInTargetUrl** | Pointer to **string** | The complete webhook url that will be called when a notification event for port in request or port in phone number happens |
**PortOutTargetUrl** | Pointer to **string** | The complete webhook url that will be called when a notification event for a port out phone number happens. |
**NotificationsOf** | Pointer to **[]string** | A list to filter what notification events to receive for this account and its sub accounts. If it is an empty list, then it means that there are no filters for the notifications events to send in each webhook and all events will get sent. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


