# NumbersV1PortingPortInPhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortInRequestSid** | Pointer to **string** | The SID of the Port In request. This is a unique identifier of the port in request. |
**PhoneNumberSid** | Pointer to **string** | The SID of the Port In request phone number. This is a unique identifier of the phone number. |
**Url** | Pointer to **string** |  |
**AccountSid** | Pointer to **string** | The SID of the account that the phone number belongs to. |
**PhoneNumberType** | Pointer to **string** | The type of the phone number. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date when the phone number was created. |
**Country** | Pointer to **string** | The country of the phone number. |
**MissingRequiredFields** | Pointer to **bool** | The phone number is missing required fields. |
**LastUpdated** | Pointer to [**time.Time**](time.Time.md) | The timestamp when the status was last updated. |
**PhoneNumber** | Pointer to **string** | The phone number. |
**Portable** | Pointer to **bool** | The phone number is portable. |
**NotPortabilityReason** | Pointer to **string** | The reason why the phone number is not portable. |
**NotPortabilityReasonCode** | Pointer to **int** | The code of the reason why the phone number is not portable. |
**PortInPhoneNumberStatus** | Pointer to **string** | The status of the phone number in the port in request. |
**PortOutPin** | Pointer to **int** | The pin required for the losing carrier to port out the phone number. |
**RejectionReason** | Pointer to **string** | The rejection reason returned by the vendor. |
**RejectionReasonCode** | Pointer to **int** | The rejection reason code returned by the vendor. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


