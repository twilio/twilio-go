# StudioV2Flow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**CommitMessage** | Pointer to **string** | Description of change made in the revision |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Definition** | Pointer to **map[string]interface{}** | JSON representation of flow definition |
**Errors** | Pointer to **[]map[string]interface{}** | List of error in the flow definition |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the Flow |
**Links** | Pointer to **map[string]interface{}** | Nested resource URLs |
**Revision** | Pointer to **int32** | The latest revision number of the Flow's definition |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Status** | Pointer to **string** | The status of the Flow |
**Url** | Pointer to **string** | The absolute URL of the resource |
**Valid** | Pointer to **bool** | Boolean if the flow definition is valid |
**Warnings** | Pointer to **[]map[string]interface{}** | List of warnings in the flow definition |
**WebhookUrl** | Pointer to **string** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


