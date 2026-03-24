# NumbersV1PortingBulkUpdateStatusRequestPortInPhoneNumberRequests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortInPhoneNumberSid** | **string** | The SID of the Port In Phone Number resource that is being updated. |
**PortDate** | Pointer to [**time.Time**](time.Time.md) | The timestamp the phone number will be ported. This will only be set once a port date has been confirmed. Not all carriers can guarantee a specific time on the port date. Twilio will try its best to get the port completed by this time on the port date. |
**RejectionReason** | Pointer to **string** | The description of the rejection reason provided by the losing carrier. This field may be null if the number has not been rejected by the losing carrier. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


