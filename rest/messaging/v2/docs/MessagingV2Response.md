# MessagingV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsParentAccount** | **bool** | True if this account has no parent account. |[optional] 
**ParentAccount** | **string** | The account SID whose terms-of-service acceptance and account type this reflects: this account's own SID if isParentAccount is true, otherwise the parent's SID. |[optional] 
**HasAcceptedTos** | **bool** | Whether parentAccount has accepted the WhatsApp terms of service. |[optional] 
**AccountType** | **string** | parentAccount's declared account type, or NOT_SET_YET if never set. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


