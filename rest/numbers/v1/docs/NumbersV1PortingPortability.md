# NumbersV1PortingPortability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | Pointer to **string** | The phone number which portability is to be checked. Phone numbers are in E.164 format (e.g. +16175551212). |
**AccountSid** | Pointer to **string** | The target account sid to which the number will be ported |
**Portable** | Pointer to **bool** | Boolean flag specifying if phone number is portable or not. |
**PinAndAccountNumberRequired** | Pointer to **bool** | Boolean flag specifying if PIN and account number is required for the phone number. |
**NotPortableReason** | Pointer to **string** | Reason why the phone number cannot be ported into Twilio, `null` otherwise. |
**NotPortableReasonCode** | Pointer to **int** | The Portability Reason Code for the phone number if it cannot be ported into Twilio, `null` otherwise. One of `22131`, `22132`, `22130`, `22133`, `22102` or `22135`. |
**NumberType** | Pointer to [**string**](PortingPortabilityEnumNumberType.md) |  |
**Country** | Pointer to **string** | Country the phone number belongs to. |
**Url** | Pointer to **string** | This is the url of the request that you're trying to reach out to locate the resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


