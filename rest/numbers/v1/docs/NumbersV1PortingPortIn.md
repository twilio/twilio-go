# NumbersV1PortingPortIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortInRequestSid** | Pointer to **string** | The SID of the Port In request. This is a unique identifier of the port in request. |
**Url** | Pointer to **string** | The URL of this Port In request |
**AccountSid** | Pointer to **string** | Account Sid or subaccount where the phone number(s) will be Ported |
**NotificationEmails** | Pointer to **[]string** | Additional emails to send a copy of the signed LOA to. |
**TargetPortInDate** | Pointer to **string** | Target date to port the number. We cannot guarantee that this date will be honored by the other carriers, please work with Ops to get a confirmation of the firm order commitment (FOC) date. Expected format is ISO Local Date, example: ‘2011-12-03`. This date must be at least 7 days in the future for US ports and 10 days in the future for Japanese ports. (This value is only available for custom porting customers.) |
**TargetPortInTimeRangeStart** | Pointer to **string** | The earliest time that the port should occur on the target port in date. Expected format is ISO Offset Time, example: ‘10:15:00-08:00'. (This value is only available for custom porting customers.) |
**TargetPortInTimeRangeEnd** | Pointer to **string** | The latest time that the port should occur on the target port in date. Expected format is ISO Offset Time, example: ‘10:15:00-08:00'.  (This value is only available for custom porting customers.) |
**PortInRequestStatus** | Pointer to **string** | The status of the port in request. The possible values are: In progress, Completed, Expired, In review, Waiting for Signature, Action Required, and Canceled. |
**LosingCarrierInformation** | Pointer to **interface{}** | Details regarding the customer’s information with the losing carrier. These values will be used to generate the letter of authorization and should match the losing carrier’s data as closely as possible to ensure the port is accepted. |
**PhoneNumbers** | Pointer to **[]interface{}** |  |
**Documents** | Pointer to **[]string** | List of document SIDs for all phone numbers included in the port in request. At least one document SID referring to a document of the type Utility Bill is required. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


