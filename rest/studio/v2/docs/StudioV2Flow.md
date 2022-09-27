# StudioV2Flow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the Flow |
**Definition** | Pointer to **interface{}** | JSON representation of flow definition |
**Status** | Pointer to [**string**](FlowEnumStatus.md) |  |
**Revision** | Pointer to **int** | The latest revision number of the Flow's definition |
**CommitMessage** | Pointer to **string** | Description of change made in the revision |
**Valid** | Pointer to **bool** | Boolean if the flow definition is valid |
**Errors** | Pointer to **[]interface{}** | List of error in the flow definition |
**Warnings** | Pointer to **[]interface{}** | List of warnings in the flow definition |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**WebhookUrl** | Pointer to **string** |  |
**Url** | Pointer to **string** | The absolute URL of the resource |
**Links** | Pointer to **map[string]interface{}** | Nested resource URLs |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


