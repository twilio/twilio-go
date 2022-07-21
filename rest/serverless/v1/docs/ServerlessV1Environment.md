# ServerlessV1Environment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the Environment resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the Environment resource |
**ServiceSid** | Pointer to **string** | The SID of the Service that the Environment resource is associated with |
**BuildSid** | Pointer to **string** | The SID of the build deployed in the environment |
**UniqueName** | Pointer to **string** | A user-defined string that uniquely identifies the Environment resource |
**DomainSuffix** | Pointer to **string** | A URL-friendly name that represents the environment |
**DomainName** | Pointer to **string** | The domain name for all Functions and Assets deployed in the Environment |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Environment resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Environment resource was last updated |
**Url** | Pointer to **string** | The absolute URL of the Environment resource |
**Links** | Pointer to **map[string]interface{}** | The URLs of the Environment resource's nested resources |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


