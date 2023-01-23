# ApiV2010RecordingAddOnResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**Status** | Pointer to [**string**](RecordingAddOnResultEnumStatus.md) |  |
**AddOnSid** | **string** | The SID of the Add-on to which the result belongs |[optional] 
**AddOnConfigurationSid** | **string** | The SID of the Add-on configuration |[optional] 
**DateCreated** | **string** | The RFC 2822 date and time in GMT that the resource was created |[optional] 
**DateUpdated** | **string** | The RFC 2822 date and time in GMT that the resource was last updated |[optional] 
**DateCompleted** | **string** | The date and time in GMT that the result was completed |[optional] 
**ReferenceSid** | **string** | The SID of the recording to which the AddOnResult resource belongs |[optional] 
**SubresourceUris** | **map[string]interface{}** | A list of related resources identified by their relative URIs |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


