# ApiV2010RecordingAddOnResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Status** | Pointer to [**string**](RecordingAddOnResultEnumStatus.md) |  |
**AddOnSid** | Pointer to **string** | The SID of the Add-on to which the result belongs |
**AddOnConfigurationSid** | Pointer to **string** | The SID of the Add-on configuration |
**DateCreated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was created |
**DateUpdated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was last updated |
**DateCompleted** | Pointer to **string** | The date and time in GMT that the result was completed |
**ReferenceSid** | Pointer to **string** | The SID of the recording to which the AddOnResult resource belongs |
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of related resources identified by their relative URIs |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


