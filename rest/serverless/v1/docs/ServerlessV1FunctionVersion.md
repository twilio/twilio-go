# ServerlessV1FunctionVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the Function Version resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the Function Version resource |
**ServiceSid** | Pointer to **string** | The SID of the Service that the Function Version resource is associated with |
**FunctionSid** | Pointer to **string** | The SID of the Function resource that is the parent of the Function Version resource |
**Path** | Pointer to **string** | The URL-friendly string by which the Function Version resource can be referenced |
**Visibility** | Pointer to [**string**](FunctionVersionEnumVisibility.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Function Version resource was created |
**Url** | Pointer to **string** | The absolute URL of the Function Version resource |
**Links** | Pointer to **map[string]interface{}** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


