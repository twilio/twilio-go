# IamV1PostLoginActionUserContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | Whitney's own identifier for the user. This is not a Twilio user SID - the Twilio user record is resolved by email instead |[optional] 
**Emails** | **[]string** | The user's email addresses. The first entry is used to resolve their Twilio unified user record |[optional] 
**Phones** | **[]string** | The user's phone numbers, if any are on file |[optional] 
**IsFirstLogin** | **bool** | Whether this is the user's first successful login |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


