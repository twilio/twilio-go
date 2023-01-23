# ApiV2010Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource. |
**City** | Pointer to **string** | The city in which the address is located. |
**CustomerName** | Pointer to **string** | The name associated with the address.This property has a maximum length of 16 4-byte characters, or 21 3-byte characters. |
**DateCreated** | Pointer to **string** | The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to **string** | The date and time in GMT that the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**IsoCountry** | Pointer to **string** | The ISO country code of the address. |
**PostalCode** | Pointer to **string** | The postal code of the address. |
**Region** | Pointer to **string** | The state or region of the address. |
**Sid** | Pointer to **string** | The unique string that that we created to identify the Address resource. |
**Street** | Pointer to **string** | The number and street address of the address. |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com`. |
**EmergencyEnabled** | Pointer to **bool** | Whether emergency calling has been enabled on this number. |
**Validated** | Pointer to **bool** | Whether the address has been validated to comply with local regulation. In countries that require valid addresses, an invalid address will not be accepted. `true` indicates the Address has been validated. `false` indicate the country doesn't require validation or the Address is not valid. |
**Verified** | Pointer to **bool** | Whether the address has been verified to comply with regulation. In countries that require valid addresses, an invalid address will not be accepted. `true` indicates the Address has been verified. `false` indicate the country doesn't require verified or the Address is not valid. |
**StreetSecondary** | Pointer to **string** | The additional number and street address of the address. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


