# NumbersV1VoiceVerificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | **string** | The phone number to verify in E.164 format. |
**FriendlyName** | **string** | A human-readable name for the caller ID. |[optional] 
**CustomCode** | **string** | A custom verification code to use instead of a generated one. |[optional] 
**CallDelay** | **string** | Number of seconds to delay the verification call (0-60). |[optional] 
**StatusCallback** | **string** | URL to receive status callback events. |[optional] 
**StatusCallbackMethod** | **string** | HTTP method for status callback requests. |[optional] 
**Extension** | **string** | Phone extension to dial after connecting. |[optional] 
**CustomSuccessMessage** | **string** | Custom message to play after successful verification. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


