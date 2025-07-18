# IamV1InvitedAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationSid** | **string** | Unique Twilio notification sid |[optional] 
**AccountSid** | **string** | Unique Twilio account sid |[optional] 
**SenderFirstName** | **string** | First name of the sender who invited the account |[optional] 
**SenderLastName** | **string** | Last name of the sender who invited the account |[optional] 
**SenderEmail** | **string** | Email of the sender who invited the account |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date and time of the invitation creation |[optional] 
**TotalSends** | **int** | Total number of invitations sent |[optional] 
**DateLastSent** | [**time.Time**](time.Time.md) | The date and time of the last invitation sent |[optional] 
**TokenExpired** | **bool** | Whether the invitation token has expired |[optional] 
**CanResend** | **bool** | Whether the invitation can be resent |[optional] 
**CanResendReason** | **string** | Reason why the invitation cannot be resent, if applicable |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


