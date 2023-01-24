# ServerlessV1AssetVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Asset Version resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Asset Version resource. |
**ServiceSid** | Pointer to **string** | The SID of the Service that the Asset Version resource is associated with. |
**AssetSid** | Pointer to **string** | The SID of the Asset resource that is the parent of the Asset Version. |
**Path** | Pointer to **string** | The URL-friendly string by which the Asset Version can be referenced. It can be a maximum of 255 characters. All paths begin with a forward slash ('/'). If an Asset Version creation request is submitted with a path not containing a leading slash, the path will automatically be prepended with one. |
**Visibility** | Pointer to [**string**](AssetVersionEnumVisibility.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the Asset Version resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The absolute URL of the Asset Version resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


