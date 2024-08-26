# NumbersV1PortingPortability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | Pointer to **string** | The phone number which portability is to be checked. Phone numbers are in E.164 format (e.g. +16175551212). |
**AccountSid** | Pointer to **string** | Account Sid that the phone number belongs to in Twilio. This is only returned for phone numbers that already exist in Twilio’s inventory and belong to your account or sub account. |
**Portable** | Pointer to **bool** | Boolean flag indicates if the phone number can be ported into Twilio through the Porting API or not. |
**PinAndAccountNumberRequired** | Pointer to **bool** | Indicates if the port in process will require a personal identification number (PIN) and an account number for this phone number. If this is true you will be required to submit both a PIN and account number from the losing carrier for this number when opening a port in request. These fields will be required in order to complete the port in process to Twilio. |
**NotPortableReason** | Pointer to **string** | Reason why the phone number cannot be ported into Twilio, `null` otherwise. |
**NotPortableReasonCode** | Pointer to **int** | The Portability Reason Code for the phone number if it cannot be ported into Twilio, `null` otherwise. |
**NumberType** | Pointer to [**string**](PortingPortabilityEnumNumberType.md) |  |
**Country** | Pointer to **string** | Country the phone number belongs to. |
**Url** | Pointer to **string** | This is the url of the request that you're trying to reach out to locate the resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


