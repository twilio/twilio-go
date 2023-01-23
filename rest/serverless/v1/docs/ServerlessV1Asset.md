# ServerlessV1Asset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Asset resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Asset resource. |
**ServiceSid** | Pointer to **string** | The SID of the Service that the Asset resource is associated with. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the Asset resource. It can be a maximum of 255 characters. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the Asset resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the Asset resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The absolute URL of the Asset resource. |
**Links** | Pointer to **map[string]interface{}** | The URLs of the Asset resource's nested resources. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


