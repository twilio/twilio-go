# AutopilotV1Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**AssistantSid** | Pointer to **string** | The SID of the Assistant that is the parent of the resource |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**Events** | Pointer to **string** | The list of space-separated events that this Webhook is subscribed to. |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource |
**Url** | Pointer to **string** | The absolute URL of the Webhook resource |
**WebhookMethod** | Pointer to **string** | The method used when calling the webhook's URL. |
**WebhookUrl** | Pointer to **string** | The URL associated with this Webhook. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


