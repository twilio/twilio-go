# NumbersV1PortingBulkPortability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this Portability check. |
**Status** | Pointer to [**string**](PortingBulkPortabilityEnumStatus.md) |  |
**DatetimeCreated** | Pointer to [**time.Time**](time.Time.md) | The date that the Portability check was created, given in ISO 8601 format. |
**PhoneNumbers** | Pointer to **[]interface{}** | Contains a list with all the information of the requested phone numbers. Each phone number contains the following properties: `phone_number`: The phone number which portability is to be checked. `portable`: Boolean flag specifying if phone number is portable or not. `not_portable_reason`: Reason why the phone number cannot be ported into Twilio, `null` otherwise. `not_portable_reason_code`: The Portability Reason Code for the phone number if it cannot be ported in Twilio, `null` otherwise. `pin_and_account_number_required`: Boolean flag specifying if PIN and account number is required for the phone number. `number_type`: The type of the requested phone number. `country` Country the phone number belongs to. `messaging_carrier` Current messaging carrier of the phone number. `voice_carrier` Current voice carrier of the phone number. |
**Url** | Pointer to **string** | This is the url of the request that you're trying to reach out to locate the resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


