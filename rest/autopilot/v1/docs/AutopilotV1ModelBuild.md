# AutopilotV1ModelBuild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**AssistantSid** | Pointer to **string** | The SID of the Assistant that is the parent of the resource |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Status** | Pointer to [**string**](ModelBuildEnumStatus.md) |  |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource |
**Url** | Pointer to **string** | The absolute URL of the ModelBuild resource |
**BuildDuration** | Pointer to **int** | The time in seconds it took to build the model |
**ErrorCode** | Pointer to **int** | More information about why the model build failed, if `status` is `failed` |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


