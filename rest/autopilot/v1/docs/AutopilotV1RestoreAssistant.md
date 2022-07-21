# AutopilotV1RestoreAssistant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**NeedsModelBuild** | Pointer to **bool** | Whether model needs to be rebuilt |
**LatestModelBuildSid** | Pointer to **string** | Reserved |
**LogQueries** | Pointer to **bool** | Whether queries should be logged and kept after training |
**DevelopmentStage** | Pointer to **string** | A string describing the state of the assistant. |
**CallbackUrl** | Pointer to **string** | Reserved |
**CallbackEvents** | Pointer to **string** | Reserved |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


