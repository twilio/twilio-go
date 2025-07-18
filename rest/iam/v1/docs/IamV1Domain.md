# IamV1Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | Unique Twilio domain sid |[optional] 
**Name** | **string** | Name of the domain |[optional] 
**Token** | **string** | Verification token for the domain |[optional] 
**Status** | **string** | Status of the domain |[optional] 
**VerificationType** | **string** | Type of verification used for the domain |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | Creation date of the domain |[optional] 
**DateLastVerification** | [**time.Time**](time.Time.md) | Date where the domain was last verified |[optional] 
**DaysLeftToReverify** | **int** | Days left to re-verify until this domain gets back into 'unverified' state. Only applicable to domains in 'verified-uncertain' state |[optional] 
**VerifiedByAnotherResource** | **bool** | Whether the domain is verified by another resource. Only applicable to domains in 'verified' state |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


