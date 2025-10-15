# NumbersV1PortingPortInPhoneNumberResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotPortabilityReason** | Pointer to **string** | The not portability reason code description. This field may be null if the number is portable or if the portability for a number has not yet been evaluated. |
**NotPortabilityReasonCode** | Pointer to **int** | The not portability reason code. This field may be null if the number is portable or if the portability for a number has not yet been evaluated. |
**NumberType** | Pointer to **string** | The number type of the phone number. This can be: toll-free, local, mobile or unknown. This field may be null if the number is not portable or if the portability for a number has not yet been evaluated. |
**PhoneNumber** | **string** | Phone number to be ported. This will be in the E164 Format. |[optional] 
**PortDate** | Pointer to [**time.Time**](time.Time.md) | The timestamp the phone number will be ported. This will only be set once a port date has been confirmed. Not all carriers can guarantee a specific time on the port date. Twilio will try its best to get the port completed by this time on the port date. Please subscribe to webhooks for confirmation on when a port has actually been completed. |
**PortInPhoneNumberSid** | **string** | The SID of the Phone number. This is a unique identifier of the phone number. |[optional] 
**PortInPhoneNumberStatus** | **string** | The status of the port in phone number. |[optional] 
**Portable** | Pointer to **bool** | Whether the number is portable by Twilio or not. This field may be null if the number portability has not yet been evaluated. If a number is not portable reference the `not_portability_reason_code` and `not_portability_reason` fields for more details |
**RejectionReason** | Pointer to **string** | The description of the rejection reason provided by the losing carrier. This field may be null if the number has not been rejected by the losing carrier. |
**RejectionReasonCode** | Pointer to **string** | The code for the rejection reason provided by the losing carrier. This field may be null if the number has not been rejected by the losing carrier. |
**StatusLastTimeUpdatedTimestamp** | Pointer to **string** | Timestamp indicating when the Port In Phone Number resource was last modified. |
**ExternalPortingVendorPhoneNumberId** | Pointer to **string** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


