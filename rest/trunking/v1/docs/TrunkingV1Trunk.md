# TrunkingV1Trunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**DomainName** | **string** | The unique address you reserve on Twilio to which you route your SIP traffic |[optional] 
**DisasterRecoveryMethod** | **string** | The HTTP method we use to call the disaster_recovery_url |[optional] 
**DisasterRecoveryUrl** | **string** | The HTTP URL that we call if an error occurs while sending SIP traffic towards your configured Origination URL |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**Secure** | **bool** | Whether Secure Trunking is enabled for the trunk |[optional] 
**Recording** | Pointer to **interface{}** | The recording settings for the trunk |
**TransferMode** | Pointer to [**string**](TrunkEnumTransferSetting.md) |  |
**TransferCallerId** | Pointer to [**string**](TrunkEnumTransferCallerId.md) |  |
**CnamLookupEnabled** | **bool** | Whether Caller ID Name (CNAM) lookup is enabled for the trunk |[optional] 
**AuthType** | **string** | The types of authentication mapped to the domain |[optional] 
**AuthTypeSet** | **[]string** | Reserved |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**Url** | **string** | The absolute URL of the resource |[optional] 
**Links** | **map[string]interface{}** | The URLs of related resources |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


