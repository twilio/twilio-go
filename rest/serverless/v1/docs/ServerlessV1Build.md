# ServerlessV1Build

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Build resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Build resource. |
**ServiceSid** | Pointer to **string** | The SID of the Service that the Build resource is associated with. |
**Status** | Pointer to [**string**](BuildEnumStatus.md) |  |
**AssetVersions** | Pointer to **[]interface{}** | The list of Asset Version resource SIDs that are included in the Build. |
**FunctionVersions** | Pointer to **[]interface{}** | The list of Function Version resource SIDs that are included in the Build. |
**Dependencies** | Pointer to **[]interface{}** | A list of objects that describe the Dependencies included in the Build. Each object contains the `name` and `version` of the dependency. |
**Runtime** | Pointer to [**string**](BuildEnumRuntime.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the Build resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the Build resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The absolute URL of the Build resource. |
**Links** | Pointer to **map[string]interface{}** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


