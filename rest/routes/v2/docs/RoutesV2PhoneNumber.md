# RoutesV2PhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | Pointer to **string** | The phone number in E.164 format |
**Url** | Pointer to **string** | The absolute URL of the resource. |
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies the Inbound Processing Region assignments for this phone number. |
**AccountSid** | Pointer to **string** | The unique SID identifier of the Account. |
**FriendlyName** | Pointer to **string** | A human readable description of the Inbound Processing Region assignments for this phone number, up to 64 characters. |
**VoiceRegion** | Pointer to **string** | The Inbound Processing Region used for this phone number for voice. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this phone number was assigned an Inbound Processing Region, given in ISO 8601 format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that the Inbound Processing Region was updated for this phone number, given in ISO 8601 format. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


