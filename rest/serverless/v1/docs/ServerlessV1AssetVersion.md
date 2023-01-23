# ServerlessV1AssetVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the Asset Version resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the Asset Version resource |[optional] 
**ServiceSid** | **string** | The SID of the Service that the Asset Version resource is associated with |[optional] 
**AssetSid** | **string** | The SID of the Asset resource that is the parent of the Asset Version |[optional] 
**Path** | **string** | The URL-friendly string by which the Asset Version can be referenced |[optional] 
**Visibility** | Pointer to [**string**](AssetVersionEnumVisibility.md) |  |
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Asset Version resource was created |[optional] 
**Url** | **string** | The absolute URL of the Asset Version resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


