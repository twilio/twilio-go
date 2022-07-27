# ServerlessV1AssetVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the Asset Version resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the Asset Version resource |
**ServiceSid** | Pointer to **string** | The SID of the Service that the Asset Version resource is associated with |
**AssetSid** | Pointer to **string** | The SID of the Asset resource that is the parent of the Asset Version |
**Path** | Pointer to **string** | The URL-friendly string by which the Asset Version can be referenced |
**Visibility** | Pointer to [**string**](AssetVersionEnumVisibility.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Asset Version resource was created |
**Url** | Pointer to **string** | The absolute URL of the Asset Version resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


