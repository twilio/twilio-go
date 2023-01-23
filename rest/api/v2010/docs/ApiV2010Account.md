# ApiV2010Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthToken** | **string** | The authorization token for this account |[optional] 
**DateCreated** | **string** | The date this account was created |[optional] 
**DateUpdated** | **string** | The date this account was last updated |[optional] 
**FriendlyName** | **string** | A human readable description of this account |[optional] 
**OwnerAccountSid** | **string** | The unique 34 character id representing the parent of this account |[optional] 
**Sid** | **string** | A 34 character string that uniquely identifies this resource. |[optional] 
**Status** | Pointer to [**string**](AccountEnumStatus.md) |  |
**SubresourceUris** | **map[string]interface{}** | Account Instance Subresources |[optional] 
**Type** | Pointer to [**string**](AccountEnumType.md) |  |
**Uri** | **string** | The URI for this resource, relative to `https://api.twilio.com` |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


