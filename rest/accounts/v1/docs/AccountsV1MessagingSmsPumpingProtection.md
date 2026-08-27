# AccountsV1MessagingSmsPumpingProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionMode** | Pointer to **string** | Protection mode for SMS pumping protection (basic or advanced). |
**DefaultProtectionLevel** | Pointer to **string** | Required when protection_mode is \"advanced\". Default SMS pumping protection level. Note: disable is NOT allowed for default protection level. Options: low, medium, high. Must be omitted when protection_mode is \"basic\". |
**MessageIntent** | Pointer to **string** | Required when protection_mode is \"advanced\". Message intent/purpose for SMS pumping protection. Example: otp, marketing, etc. Options: otp, notifications, fraud, security, customercare, delivery, education, events, polling, announcement, marketing, ai_classification. Must be omitted when protection_mode is \"basic\". |
**CountrySpecificProtectionLevel** | Pointer to [**AccountsV1MessagingSmsPumpingProtectionCountrySpecificProtectionLevel**](AccountsV1MessagingSmsPumpingProtectionCountrySpecificProtectionLevel.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


