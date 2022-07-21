# TrunkingV1Trunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**DomainName** | Pointer to **string** | The unique address you reserve on Twilio to which you route your SIP traffic |
**DisasterRecoveryMethod** | Pointer to **string** | The HTTP method we use to call the disaster_recovery_url |
**DisasterRecoveryUrl** | Pointer to **string** | The HTTP URL that we call if an error occurs while sending SIP traffic towards your configured Origination URL |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**Secure** | Pointer to **bool** | Whether Secure Trunking is enabled for the trunk |
**Recording** | Pointer to **interface{}** | The recording settings for the trunk |
**TransferMode** | Pointer to [**string**](TrunkEnumTransferSetting.md) |  |
**TransferCallerId** | Pointer to [**string**](TrunkEnumTransferCallerId.md) |  |
**CnamLookupEnabled** | Pointer to **bool** | Whether Caller ID Name (CNAM) lookup is enabled for the trunk |
**AuthType** | Pointer to **string** | The types of authentication mapped to the domain |
**AuthTypeSet** | Pointer to **[]string** | Reserved |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Url** | Pointer to **string** | The absolute URL of the resource |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


