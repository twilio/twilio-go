# StudioV2Flow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the Flow |[optional] 
**Definition** | Pointer to **interface{}** | JSON representation of flow definition |
**Status** | Pointer to [**string**](FlowEnumStatus.md) |  |
**Revision** | **int** | The latest revision number of the Flow's definition |[optional] 
**CommitMessage** | **string** | Description of change made in the revision |[optional] 
**Valid** | **bool** | Boolean if the flow definition is valid |[optional] 
**Errors** | **[]interface{}** | List of error in the flow definition |[optional] 
**Warnings** | **[]interface{}** | List of warnings in the flow definition |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**WebhookUrl** | **string** |  |[optional] 
**Url** | **string** | The absolute URL of the resource |[optional] 
**Links** | **map[string]interface{}** | Nested resource URLs |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


