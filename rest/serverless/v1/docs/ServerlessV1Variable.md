# ServerlessV1Variable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the Variable resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the Variable resource |
**ServiceSid** | Pointer to **string** | The SID of the Service that the Variable resource is associated with |
**EnvironmentSid** | Pointer to **string** | The SID of the Environment in which the Variable exists |
**Key** | Pointer to **string** | A string by which the Variable resource can be referenced |
**Value** | Pointer to **string** | A string that contains the actual value of the Variable |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Variable resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Variable resource was last updated |
**Url** | Pointer to **string** | The absolute URL of the Variable resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


