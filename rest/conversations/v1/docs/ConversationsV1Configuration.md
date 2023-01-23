# ConversationsV1Configuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this configuration. |
**DefaultChatServiceSid** | Pointer to **string** | The SID of the default [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) used when creating a conversation. |
**DefaultMessagingServiceSid** | Pointer to **string** | The SID of the default [Messaging Service](https://www.twilio.com/docs/sms/services/api) used when creating a conversation. |
**DefaultInactiveTimer** | Pointer to **string** | Default ISO8601 duration when conversation will be switched to `inactive` state. Minimum value for this timer is 1 minute. |
**DefaultClosedTimer** | Pointer to **string** | Default ISO8601 duration when conversation will be switched to `closed` state. Minimum value for this timer is 10 minutes. |
**Url** | Pointer to **string** | An absolute API resource URL for this global configuration. |
**Links** | Pointer to **map[string]interface{}** | Contains absolute API resource URLs to access the webhook and default service configurations. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


