# AutopilotV1ModelBuild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |[optional] 
**AssistantSid** | **string** | The SID of the Assistant that is the parent of the resource |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**Status** | Pointer to [**string**](ModelBuildEnumStatus.md) |  |
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource |[optional] 
**Url** | **string** | The absolute URL of the ModelBuild resource |[optional] 
**BuildDuration** | Pointer to **int** | The time in seconds it took to build the model |
**ErrorCode** | Pointer to **int** | More information about why the model build failed, if `status` is `failed` |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


