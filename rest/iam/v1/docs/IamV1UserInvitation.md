# IamV1UserInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | Unique Twilio sid of the account that the user was invited to |[optional] 
**Sid** | **string** | Unique Twilio notification sid of the invitation |[optional] 
**Email** | **string** | The recipient's email address |[optional] 
**Roles** | [**[]IamV1UserRole**](IamV1UserRole.md) | The roles assigned to the user in the account |[optional] 
**Status** | **string** | The status of the invitation |[optional] 
**InvitedBy** | **string** | The sender's email address |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date and time when the invitation was created |[optional] 
**DateLastSent** | [**time.Time**](time.Time.md) | The date and time when the invitation was last sent |[optional] 
**TokenExpired** | **bool** | Indicates whether the invitation token has expired |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


