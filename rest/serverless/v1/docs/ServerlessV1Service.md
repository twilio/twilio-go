# ServerlessV1Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the Service resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the Service resource |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the Service resource |
**UniqueName** | Pointer to **string** | A user-defined string that uniquely identifies the Service resource |
**IncludeCredentials** | Pointer to **bool** | Whether to inject Account credentials into a function invocation context |
**UiEditable** | Pointer to **bool** | Whether the Service resource's properties and subresources can be edited via the UI |
**DomainBase** | Pointer to **string** | The base domain name for this Service, which is a combination of the unique name and a randomly generated string |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Service resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Service resource was last updated |
**Url** | Pointer to **string** | The absolute URL of the Service resource |
**Links** | Pointer to **map[string]interface{}** | The URLs of the Service's nested resources |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


