# UpdateHostedNumberOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallDelay** | **int32** | The number of seconds, between 0 and 60, to delay before initiating the verification call. Defaults to 0. | [optional] 
**CcEmails** | **[]string** | Optional. A list of emails that LOA document for this HostedNumberOrder will be carbon copied to. | [optional] 
**Email** | **string** | Email of the owner of this phone number that is being hosted. | [optional] 
**Extension** | **string** | Digits to dial after connecting the verification call. | [optional] 
**FriendlyName** | **string** | A 64 character string that is a human readable text that describes this resource. | [optional] 
**Status** | **string** | User can only post to &#x60;pending-verification&#x60; status to transition the HostedNumberOrder to initiate a verification call or verification of ownership with a copy of a phone bill. | [optional] 
**UniqueName** | **string** | Provides a unique and addressable name to be assigned to this HostedNumberOrder, assigned by the developer, to be optionally used in addition to SID. | [optional] 
**VerificationCode** | **string** | A verification code that is given to the user via a phone call to the phone number that is being hosted. | [optional] 
**VerificationDocumentSid** | **string** | Optional. The unique sid identifier of the Identity Document that represents the document for verifying ownership of the number to be hosted. Required when VerificationType is phone-bill. | [optional] 
**VerificationType** | **string** | Optional. The method used for verifying ownership of the number to be hosted. One of phone-call (default) or phone-bill. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


