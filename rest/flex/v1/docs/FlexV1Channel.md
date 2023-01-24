# FlexV1Channel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Channel resource and owns this Workflow. |
**FlexFlowSid** | Pointer to **string** | The SID of the Flex Flow. |
**Sid** | Pointer to **string** | The unique string that we created to identify the Channel resource. |
**UserSid** | Pointer to **string** | The SID of the chat user. |
**TaskSid** | Pointer to **string** | The SID of the TaskRouter Task. Only valid when integration type is `task`. `null` for integration types `studio` & `external` |
**Url** | Pointer to **string** | The absolute URL of the Flex chat channel resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the Flex chat channel was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the Flex chat channel was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


