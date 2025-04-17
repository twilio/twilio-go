# NumbersV1PortingPortInPhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortInRequestSid** | Pointer to **string** | The unique identifier for the port in request that this phone number is associated with. |
**PhoneNumberSid** | Pointer to **string** | The unique identifier for this phone number associated with this port in request. |
**Url** | Pointer to **string** | URL reference for this resource. |
**AccountSid** | Pointer to **string** | Account Sid or subaccount where the phone number(s) will be Ported. |
**PhoneNumberType** | Pointer to **string** | The number type of the phone number. This can be: toll-free, local, mobile or unknown. This field may be null if the number is not portable or if the portability for a number has not yet been evaluated. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The timestamp for when this port in phone number was created. |
**Country** | Pointer to **string** | The ISO country code that this number is associated with. This field may be null if the number is not portable or if the portability for a number has not yet been evaluated. |
**MissingRequiredFields** | Pointer to **bool** | Indicates if the phone number is missing required fields such as a PIN or account number. This field may be null if the number is not portable or if the portability for a number has not yet been evaluated. |
**LastUpdated** | Pointer to [**time.Time**](time.Time.md) | Timestamp indicating when the Port In Phone Number resource was last modified. |
**PhoneNumber** | Pointer to **string** | Phone number to be ported. This will be in the E164 Format. |
**Portable** | Pointer to **bool** | If the number is portable by Twilio or not. This field may be null if the number portability has not yet been evaluated. If a number is not portable reference the `not_portability_reason_code` and `not_portability_reason` fields for more details |
**NotPortabilityReason** | Pointer to **string** | The not portability reason code description. This field may be null if the number is portable or if the portability for a number has not yet been evaluated. |
**NotPortabilityReasonCode** | Pointer to **int** | The not portability reason code. This field may be null if the number is portable or if the portability for a number has not yet been evaluated. |
**PortInPhoneNumberStatus** | Pointer to **string** | The status of the port in phone number. |
**PortOutPin** | Pointer to **int** | The pin required by the losing carrier to do the port out. |
**RejectionReason** | Pointer to **string** | The description of the rejection reason provided by the losing carrier. This field may be null if the number has not been rejected by the losing carrier. |
**RejectionReasonCode** | Pointer to **int** | The code for the rejection reason provided by the losing carrier. This field may be null if the number has not been rejected by the losing carrier. |
**PortDate** | Pointer to [**time.Time**](time.Time.md) | The timestamp the phone number will be ported. This will only be set once a port date has been confirmed. Not all carriers can guarantee a specific time on the port date. Twilio will try its best to get the port completed by this time on the port date. Please subscribe to webhooks for confirmation on when a port has actually been completed. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


