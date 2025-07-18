# NumbersV2Regulation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the Regulation resource. |
**FriendlyName** | Pointer to **string** | A human-readable description that is assigned to describe the Regulation resource. Examples can include Germany: Mobile - Business. |
**IsoCountry** | Pointer to **string** | The ISO country code of the phone number's country. |
**NumberType** | Pointer to **string** | The type of phone number restricted by the regulatory requirement. For example, Germany mobile phone numbers provisioned by businesses require a business name with commercial register proof from the Handelsregisterauszug and a proof of address from Handelsregisterauszug or a trade license by Gewerbeanmeldung. |
**EndUserType** | Pointer to [**string**](RegulationEnumEndUserType.md) |  |
**Requirements** | Pointer to **interface{}** | The SID of an object that holds the regulatory information of the phone number country, phone number type, and end user type. |
**Url** | Pointer to **string** | The absolute URL of the Regulation resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


