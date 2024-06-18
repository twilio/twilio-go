# NumbersV1PortingPortIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortInRequestSid** | Pointer to **string** | The SID of the Port In request. This is a unique identifier of the port in request. |
**Url** | Pointer to **string** | The URL of this Port In request |
**AccountSid** | Pointer to **string** | The Account SID that the numbers will be added to after they are ported into Twilio. |
**NotificationEmails** | Pointer to **[]string** | List of emails for getting notifications about the LOA signing process. Allowed Max 10 emails. |
**TargetPortInDate** | Pointer to **string** | Minimum number of days in the future (at least 2 days) needs to be established with the Ops team for validation. |
**TargetPortInTimeRangeStart** | Pointer to **string** | Minimum hour in the future needs to be established with the Ops team for validation. |
**TargetPortInTimeRangeEnd** | Pointer to **string** | Maximum hour in the future needs to be established with the Ops team for validation. |
**PortInRequestStatus** | Pointer to **string** | The status of the port in request. The possible values are: In progress, Completed, Expired, In review, Waiting for Signature, Action Required, and Canceled. |
**LosingCarrierInformation** | Pointer to **interface{}** | The information for the losing carrier.  |
**PhoneNumbers** | Pointer to **[]interface{}** | The list of phone numbers to Port in. Phone numbers are in E.164 format (e.g. +16175551212). |
**Documents** | Pointer to **[]string** | The list of documents SID referencing a utility bills |
**DateCreated** | Pointer to **string** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


