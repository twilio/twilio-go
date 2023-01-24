# TaskrouterV1TaskChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Task Channel resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**Sid** | Pointer to **string** | The unique string that we created to identify the Task Channel resource. |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the Task Channel, such as `voice` or `sms`. |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the Task Channel. |
**ChannelOptimizedRouting** | Pointer to **bool** | Whether the Task Channel will prioritize Workers that have been idle. When `true`, Workers that have been idle the longest are prioritized. |
**Url** | Pointer to **string** | The absolute URL of the Task Channel resource. |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


