# StudioV2Flow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Flow resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Flow resource. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the Flow. |
**Definition** | Pointer to **interface{}** | JSON representation of flow definition. |
**Status** | Pointer to [**string**](FlowEnumStatus.md) |  |
**Revision** | Pointer to **int** | The latest revision number of the Flow's definition. |
**CommitMessage** | Pointer to **string** | Description of change made in the revision. |
**Valid** | Pointer to **bool** | Boolean if the flow definition is valid. |
**Errors** | Pointer to **[]interface{}** | List of error in the flow definition. |
**Warnings** | Pointer to **[]interface{}** | List of warnings in the flow definition. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**WebhookUrl** | Pointer to **string** |  |
**Url** | Pointer to **string** | The absolute URL of the resource. |
**Links** | Pointer to **map[string]interface{}** | The URLs of the Flow's nested resources. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


