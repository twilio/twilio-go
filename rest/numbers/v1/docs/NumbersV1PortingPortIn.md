# NumbersV1PortingPortIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortInRequestSid** | Pointer to **string** | The SID of the Port In request. This is a unique identifier of the port in request. |
**Url** | Pointer to **string** | The URL of this Port In request |
**AccountSid** | Pointer to **string** | Account Sid or subaccount where the phone number(s) will be Ported |
**NotificationEmails** | Pointer to **[]string** | Additional emails to send a copy of the signed LOA to. |
**TargetPortInDate** | Pointer to **string** | Target date to port the number. We cannot guarantee that this date will be honored by the other carriers, please work with Ops to get a confirmation of the firm order commitment (FOC) date. Expected format is ISO Local Date, example: ‘2011-12-03`. This date must be at least 7 days in the future for US ports and 10 days in the future for Japanese ports. If a start and end range is provided, the date will be converted to its UTC equivalent with the ranges as reference and stored in UTC. We can't guarantee the exact date and time, as this depends on the losing carrier. |
**TargetPortInTimeRangeStart** | Pointer to **string** | The earliest time that the port should occur on the target port in date. Expected format is ISO Offset Time, example: ‘10:15:00-08:00'. We can't guarantee the exact date and time, as this depends on the losing carrier. The time will be stored and returned as UTC standard timezone. |
**TargetPortInTimeRangeEnd** | Pointer to **string** | The latest time that the port should occur on the target port in date. Expected format is ISO Offset Time, example: ‘10:15:00-08:00'. We can't guarantee the exact date and time, as this depends on the losing carrier. The time will be stored and returned as UTC standard timezone. |
**PortInRequestStatus** | Pointer to **string** | The status of the port in request. The possible values are: In progress, Completed, Expired, In review, Waiting for Signature, Action Required, and Canceled. |
**OrderCancellationReason** | Pointer to **string** | If the order is cancelled this field will provide further context on the cause of the cancellation. |
**LosingCarrierInformation** | [**NumbersV1PortingLosingCarrierInformation**](NumbersV1PortingLosingCarrierInformation.md) |  |[optional] 
**PhoneNumbers** | Pointer to [**[]NumbersV1PortingPortInPhoneNumberResult**](NumbersV1PortingPortInPhoneNumberResult.md) |  |
**BundleSid** | Pointer to **string** | The bundle sid is an optional identifier to reference a group of regulatory documents for a port request. |
**PortabilityAdvanceCarrier** | Pointer to **string** | A field only required for Japan port in requests. It is a unique identifier for the donor carrier service the line is being ported from. |
**AutoCancelApprovalNumbers** | Pointer to **string** | Japan specific field, indicates the number of phone numbers to automatically approve for cancellation. |
**Documents** | Pointer to **[]string** | List of document SIDs for all phone numbers included in the port in request. At least one document SID referring to a document of the type Utility Bill is required. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) |  |
**SupportTicketId** | **int** | Unique ID of the request's support ticket |[optional] 
**SignatureRequestUrl** | Pointer to **string** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


