# ServerlessV1Build

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the Build resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the Build resource |[optional] 
**ServiceSid** | **string** | The SID of the Service that the Build resource is associated with |[optional] 
**Status** | Pointer to [**string**](BuildEnumStatus.md) |  |
**AssetVersions** | **[]interface{}** | The list of Asset Version resource SIDs that are included in the Build |[optional] 
**FunctionVersions** | **[]interface{}** | The list of Function Version resource SIDs that are included in the Build |[optional] 
**Dependencies** | **[]interface{}** | A list of objects that describe the Dependencies included in the Build |[optional] 
**Runtime** | Pointer to [**string**](BuildEnumRuntime.md) |  |
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Build resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Build resource was last updated |[optional] 
**Url** | **string** | The absolute URL of the Build resource |[optional] 
**Links** | **map[string]interface{}** |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


