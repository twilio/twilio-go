# AutopilotV1Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The absolute URL of the Webhook resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Webhook resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**AssistantSid** | Pointer to **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource. |
**Sid** | Pointer to **string** | The unique string that we created to identify the Webhook resource. |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource. |
**Events** | Pointer to **string** | The list of space-separated events that this Webhook is subscribed to. |
**WebhookUrl** | Pointer to **string** | The URL associated with this Webhook. |
**WebhookMethod** | Pointer to **string** | The method used when calling the webhook's URL. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


