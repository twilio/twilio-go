# NumbersV2AvailablePhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Did** | Pointer to **string** | The phone number in E.164 format. |
**InventoryDidSid** | Pointer to **string** | The unique string that identifies the inventory DID resource. |
**FriendlyName** | Pointer to **string** | A human-readable phone number in national format. |
**Type** | Pointer to **string** | The type of phone number. Can be Local, Mobile, TollFree, etc. |
**Npa** | Pointer to **string** | The North American Numbering Plan (NANP) area code of the phone number. |
**Nxx** | Pointer to **string** | The three-digit exchange code of the phone number. |
**Locked** | Pointer to **bool** | Whether the phone number is locked for purchase. |
**LockedUntil** | Pointer to **int** | The Unix timestamp when the phone number lock expires. |
**Capabilities** | Pointer to [**NumbersV2AvailablePhoneNumberCapabilities**](NumbersV2AvailablePhoneNumberCapabilities.md) |  |
**Geography** | Pointer to [**NumbersV2AvailablePhoneNumberGeography**](NumbersV2AvailablePhoneNumberGeography.md) |  |
**AddressRequirements** | Pointer to **string** | The type of Address resource the phone number requires. |
**Certifications** | Pointer to [**NumbersV2AvailablePhoneNumberCertifications**](NumbersV2AvailablePhoneNumberCertifications.md) |  |
**Billing** | Pointer to [**NumbersV2AvailablePhoneNumberBilling**](NumbersV2AvailablePhoneNumberBilling.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in ISO 8601 format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in ISO 8601 format. |
**Beta** | Pointer to **bool** | Whether the phone number is in beta. |
**VoiceEmergencyCapable** | Pointer to **bool** | Whether the phone number can handle emergency calls. |
**Flags** | Pointer to [**NumbersV2AvailablePhoneNumberFlags**](NumbersV2AvailablePhoneNumberFlags.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


